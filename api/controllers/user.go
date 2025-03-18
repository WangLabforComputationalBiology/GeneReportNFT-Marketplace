package controllers

import (
	"GeneReport_platform/api/dto/user_dto"
	"GeneReport_platform/internal/dao/global"
	"GeneReport_platform/internal/services"
	"GeneReport_platform/pkg/auth"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"io"
	"log"
	"net/http"
)

type User struct{}

var UserController = User{}

// GetNonce
// @Summary 处理用户获取签名nonce请求，用于防重放
// @Description 检查当前redis中nonce是否过期：未过期则更新后返回，过期则重新生成
// @Tags 用户认证
// @Produce json
// @Param        user_address  path  string  true  "User address"
// @Success      200  {object} map[string]string "成功响应nonce"
// @Failure      400  {object} map[string]string "地址非法或无效"
// @Failure      503  {object} map[string]string "redis服务不可用"
// @Router       /user_dto/nonce/{user_address} [get]
func (u *User) GetNonce(ctx *gin.Context) {
	address := ctx.Param("user_address")
	//参数校验
	if !auth.IsValidAddress(address) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "地址非法或无效"})
		return
	}
	//获取nonce
	if nonce, err := services.UserService.GetNonce(address); err != nil {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{"error": "redis服务不可用"})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{"nonce": nonce})
	}

}

// Login
// @Summary      用户登录
// @Description  通过用户地址和签名生成 JWT 令牌进行登录
// @Tags         用户认证
// @Accept       json
// @Produce      json
// @Param        loginReq  body      user_dto.LoginReq  true  "登录请求参数，包括用户地址和签名"
// @Success      200       {object}  map[string]string "成功响应 JWT 令牌"
// @Failure      400       {object}  map[string]string "请求体格式错误"/"地址非法或无效"
// @Failure      401       {object}  map[string]string "签名验证失败"
// @Failure      503       {object}  map[string]string "redis服务不可用"
// @Router       /user_dto/login [post]
func (u *User) Login(ctx *gin.Context) {
	log.Println("进入登录接口！")

	var json user_dto.LoginReq
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "请求体格式错误"})
		return
	}

	address := json.UserAddress
	signature := json.Signature

	if !auth.IsValidAddress(address) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "地址非法或无效"})
		return
	}
	if err := services.UserService.EnsureUserExists(address); err != nil {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{"error": "mysql服务不可用！"})
	}
	//获取nonce
	nonce, err := services.UserService.GetNonce(address)
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{"error": "redis服务不可用"})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{"nonce": nonce})
	}
	//执行验签
	if isAccept, err := auth.VerifySignature(address, nonce, signature); !isAccept && err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "签名验证失败"})
	}

	jwt, _ := auth.GenerateToken(address)
	ctx.JSON(200, gin.H{"access_token": jwt})
	log.Printf("签名验证成功！\n用户地址: %v;用户签名: %v\n", address, signature)
}

// Logout
// @Summary 登出
// @Description 用户退出登录，将当前 JWT 加入黑名单
// @Tags 用户认证
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} map[string]string "成功登出"
// @Failure 503 {object} map[string]string "Redis服务不可用"
// @Router /user_dto/logout [post]
func (u *User) Logout(ctx *gin.Context) {
	jti, _ := ctx.Get("jti")
	//将jti加入redis黑名单
	err := global.RedisClient.SetEX(ctx, "blacklist:"+jti.(string), "1", auth.TokenExpireDuration).Err()
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{"error": "Redis服务不可用"})
		return
	}
	ctx.JSON(200, gin.H{"msg:": "成功登出"})
}

// EditUserName
// @Summary      用户编辑用户名
// @Description 根据用户地址更新用户名
// @Tags 用户管理
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param object body user_dto.UpdateUser true "请求体，包含新用户名"
// @Success 200 {object} map[string]string "用户名更新成功的响应"
// @Failure 400 {object} map[string]string "请求体格式错误"
// @Failure 503 {object} map[string]string "mysql不可用"
// @Router /user_dto/edit/name [post]
func (u *User) EditUserName(ctx *gin.Context) {
	log.Println("进入编辑用户名接口！")
	var json user_dto.UpdateUser
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "请求体格式错误"})
		return
	}

	newName := json.Name
	log.Println(" 来自post请求体的json的new_name:" + newName)
	toUpdate := user_dto.UpdateUser{Name: newName}
	if err := services.UserService.UpdateUser(toUpdate); err != nil {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{"error": "mysql不可用"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"new_name": newName, "msg": "用户名更新成功"})
}

// UploadProfile 上传用户头像
func (u *User) UploadProfile(ctx *gin.Context) {
	// 获取名为"profile"的文件
	file, header, err := ctx.Request.FormFile("profile")
	if err != nil {
		ctx.String(http.StatusBadRequest, "获取文件失败: %v", err)
		return
	}
	defer file.Close()

	// 获取其他表单字段的值
	address := ctx.PostForm("user_address")

	//图片保存在minio的名字
	pictureName := address + ".png"

	//先删除原本在minio用户对应的头像图片文件,Assignment count mismatch: 2 = 1表明在赋值语句中，左侧的变量数量与右侧提供的值数量不匹配。
	err = global.MinioClient.RemoveObject(context.Background(), "test", pictureName, minio.RemoveObjectOptions{})

	if err != nil {
		log.Println("上传失败！！\n", err)
	}

	fileSize := header.Size
	// todo 处理文件，例如保存到minio服务器,桶可以考虑用地址哈希！
	// 上传文件到 MinIO
	_, err = global.MinioClient.PutObject(
		context.Background(),
		"test", // 替换为你的桶名称
		pictureName,
		file,
		fileSize,
		minio.PutObjectOptions{ContentType: "application/octet-stream"},
	)
	if err != nil {
		log.Println("上传文件到 MinIO 出错！", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "无法上传文件"})
		return
	} else {
		log.Println("from Minio :上传头像成功！")
	}

	log.Println("用户:", address, "正在更改头像！")
	//更改数据库的picture字段

	var toUpdate = user_dto.UpdateUser{Avatar: "/test/" + pictureName}
	// 修改数据库的内容
	if err := services.UserService.UpdateUser(toUpdate); err != nil {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{"error": "mysql不可用,上传头像失败"})
		return
	}

	//todo 下面是把对象返回，测试用的
	//读取对象数据
	data, err := io.ReadAll(file)
	if err != nil {
		log.Println("读取对象数据出错！", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "无法读取对象数据"})
		return
	}

	log.Println("上传‘" + address + "’的头像！")
	// 设置响应头
	ctx.Header("Content-Type", "image/png")
	ctx.Header("Content-Disposition", "inline; filename=1.png")
	ctx.Data(http.StatusOK, "image/png", data)
	//ctx.String(http.StatusOK, "文件上传成功，其他字段值: %s", address)
}

// GetInfo
// @Summary      用户获取基本信息
// @Description 根据用户地址获取用户基本信息
// @Tags 用户管理
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} map[string]string "响应用户基本信息"
// @Failure 503 {object} map[string]string "mysql不可用"
// @Router /user_dto/edit/name [post]
func (u *User) GetInfo(ctx *gin.Context) {
	// 获取请求头中的 Authorization 值
	address := ctx.GetString("user_address")
	userInfo, err := services.UserService.GetUserInfo(address)
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{"error": err})
	}
	ctx.JSON(http.StatusOK, userInfo)

	// 打印查询到的用户信息
	log.Printf("找到用户: %+v\n", userInfo)

}

// GetGNFTList 返回用户藏品信息
func (u *User) GetGNFTList(ctx *gin.Context) {

}
