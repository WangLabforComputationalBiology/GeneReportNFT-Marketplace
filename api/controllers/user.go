package controllers

import (
	"GeneReport_platform/api/dto"
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
// @Summary 用户获取nonce
// @Description 检查当前redis中nonce是否过期：未过期则更新后返回，过期则重新生成
// @Tags 用户认证
// @Produce json
// @Param        user_address  path  string  true  "User address"
// @Success      200  {object} dto.Response "成功响应nonce"
// @Failure      400  {object} dto.ErrResponse "地址非法或无效"
// @Failure      503  {object} dto.ErrResponse "redis服务不可用"
// @Router       /user/nonce/{user_address} [get]
func (u *User) GetNonce(ctx *gin.Context) {
	address := ctx.Param("user_address")
	//地址校验
	if !auth.IsValidAddress(address) {
		ctx.JSON(http.StatusBadRequest, dto.ErrResponse{
			Code:    http.StatusBadRequest,
			Message: "地址非法或无效",
		})
		return
	}
	//获取nonce
	if nonce, err := services.UserService.GetNonce(address); err != nil {
		ctx.JSON(http.StatusServiceUnavailable, err.ToErrResponse())
		return
	} else {
		ctx.JSON(http.StatusOK, dto.Response{
			Code:    http.StatusOK,
			Message: "获取nonce成功",
			Data:    gin.H{"nonce": nonce},
		})
	}

}

// Login
// @Summary      用户登录
// @Description  校验签名并返回生成的 JWT 令牌
// @Tags         用户认证
// @Accept       json
// @Produce      json
// @Param        loginReq  body      user_dto.LoginReq  true  "登录请求参数，包括用户地址和签名"
// @Success      200       {object}  dto.Response "成功响应 JWT 令牌"
// @Failure      400       {object}  dto.ErrResponse "请求体格式错误"/"地址非法或无效"
// @Failure      401       {object}  dto.ErrResponse "签名验证失败"
// @Failure      503       {object}  dto.ErrResponse "redis服务不可用"
// @Router       /user/login [post]
func (u *User) Login(ctx *gin.Context) {
	log.Println("进入登录接口！")

	var json user_dto.LoginReq
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrResponse{
			Code:    http.StatusBadRequest,
			Message: "请求体格式错误",
		})
		return
	}

	address := json.UserAddress
	signature := json.Signature
	//地址校验
	if !auth.IsValidAddress(address) {
		ctx.JSON(http.StatusBadRequest, dto.ErrResponse{
			Code:    http.StatusBadRequest,
			Message: "地址非法或无效",
		})
		return
	}
	//确保用户存在，不存在执行创建
	if err := services.UserService.EnsureUserExists(address); err != nil {
		ctx.JSON(http.StatusServiceUnavailable, err.ToErrResponse())
		return
	}
	//获取nonce
	nonce, err := services.UserService.GetNonce(address)
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, err.ToErrResponse())
		return
	}

	//执行验签
	if isAccept, err := auth.VerifySignature(address, nonce, signature); !isAccept && err != nil {
		ctx.JSON(http.StatusUnauthorized, dto.ErrResponse{
			Code:    http.StatusUnauthorized,
			Message: "签名验证失败," + err.Error(),
		})
		return
	}

	jwt, _ := auth.GenerateToken(address)
	ctx.JSON(200, dto.Response{
		Code:    http.StatusOK,
		Message: "成功登录",
		Data: gin.H{
			"access_token": jwt,
		},
	})
	log.Printf("签名验证成功！\n用户地址: %v;用户签名: %v\n", address, signature)
}

// Logout
// @Summary 用户登出
// @Description 用户退出登录，将当前 JWT 加入黑名单
// @Tags 用户认证
// @Produce json
// @Security JwtAuth
// @Success 200 {object} dto.Response "成功登出"
// @Failure 503 {object} dto.ErrResponse "Redis服务不可用"
// @Router /user/logout [post]
func (u *User) Logout(ctx *gin.Context) {
	jti, _ := ctx.Get("jti")
	//将jti加入redis黑名单
	err := global.RedisClient.SetEX(ctx, "blacklist:"+jti.(string), "1", auth.TokenExpireDuration).Err()
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, dto.ErrResponse{
			Code:    http.StatusServiceUnavailable,
			Message: "服务器内部错误",
		})
		return
	}
	ctx.JSON(200, dto.Response{
		Code:    http.StatusOK,
		Message: "成功登出",
		Data:    "",
	})
}

// EditUserName
// @Summary      用户编辑用户名
// @Description 根据用户地址更新用户名
// @Tags 用户管理
// @Accept json
// @Produce json
// @Security JwtAuth
// @Param object body user_dto.UpdateUser true "请求体，包含新用户名"
// @Param Authorization header string true "JWT"
// @Success 200 {object} dto.Response "用户名更新成功"
// @Failure 400 {object} dto.ErrResponse "请求体格式错误"
// @Failure 503 {object} dto.ErrResponse "mysql异常"
// @Router /user/edit/name [post]
func (u *User) EditUserName(ctx *gin.Context) {
	log.Println("进入编辑用户名接口！")
	var json user_dto.UpdateUser
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrResponse{
			Code:    http.StatusBadRequest,
			Message: "请求体格式错误",
		})
		return
	}

	newName := json.Name
	log.Println(" 来自post请求体的json的new_name:" + newName)
	toUpdate := user_dto.UpdateUser{Name: newName}
	if err := services.UserService.UpdateUser(toUpdate); err != nil {
		ctx.JSON(http.StatusServiceUnavailable, err.ToErrResponse())
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Code:    http.StatusOK,
		Message: "用户名更新成功",
		Data: gin.H{
			"new_name": newName, // 返回新用户名
		},
	})
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
// @Summary      获取用户基本信息
// @Description 根据用户地址获取用户基本信息
// @Tags 用户管理
// @Produce json
// @Security JwtAuth
// @Param Authorization header string true "JWT"
// @Success 200 {object} dto.Response "响应用户基本信息"
// @Failure 503 {object} dto.ErrResponse "mysql不可用"
// @Router /user/info [post]
func (u *User) GetInfo(ctx *gin.Context) {
	// 获取请求头中的 Authorization 值
	address := ctx.GetString("user_address")
	userInfo, err := services.UserService.GetUserInfo(address)
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, err.ToErrResponse())
	}
	ctx.JSON(http.StatusOK, userInfo)

	// 打印查询到的用户信息
	log.Printf("找到用户: %+v\n", userInfo)

}

// GetGNFTList 返回用户藏品信息
func (u *User) GetGNFTList(ctx *gin.Context) {

}
