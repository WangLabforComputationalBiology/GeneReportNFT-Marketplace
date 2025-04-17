package controllers

import (
	"GeneReport_platform/api/dto"
	"GeneReport_platform/configs"
	"GeneReport_platform/internal/services"
	"GeneReport_platform/pkg/auth"
	"bytes"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type User struct{}

var UserController = User{}

func (u *User) Test(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "请求用户 controller successful！"})

}

// GetNonce
//
//	@Summary		用户获取nonce
//	@Description	检查当前redis中nonce是否过期：未过期则更新后返回，过期则重新生成
//	@Tags			用户管理
//	@Produce		json
//	@Param			user_address	path		string			true	"User address"
//	@Success		200				{object}	dto.Response	"成功响应nonce"
//	@Failure		400				{object}	dto.ErrResponse	"地址非法或无效"
//	@Failure		503				{object}	dto.ErrResponse	"redis服务不可用"
//	@Router			/user/nonce/{user_address} [get]
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
//
//	@Summary		用户登录
//	@Description	校验签名并返回生成的 JWT 令牌
//	@Tags			用户管理
//	@Accept			json
//	@Produce		json
//	@Router			/user/login [post]
//	@Param			loginRequest	body		dto.LoginReq	true	"登录请求"
//	@Success		200				{object}	dto.Response	"成功登录，返回 JWT 令牌"
//	@Failure		400				{object}	dto.ErrResponse	"请求体格式错误"
//	@Failure		400				{object}	dto.ErrResponse	"地址非法或无效"
//	@Failure		503				{object}	dto.ErrResponse	"服务不可用，确保用户存在失败"
//	@Failure		503				{object}	dto.ErrResponse	"服务不可用，获取nonce失败"
//	@Failure		401				{object}	dto.ErrResponse	"签名验证失败"
func (u *User) Login(ctx *gin.Context) {
	log.Println("进入登录接口！")

	var json dto.LoginReq
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrResponse{
			Code:    http.StatusBadRequest,
			Message: "请求体格式错误,请重新登录",
		})
		return
	}

	address := json.UserAddress
	signature := json.Signature

	//1.地址校验
	if !auth.IsValidAddress(address) {
		ctx.JSON(http.StatusBadRequest, dto.ErrResponse{
			Code:    http.StatusBadRequest,
			Message: "地址非法或无效,请重新登录",
		})
		return
	}

	//2.获取nonce,校验redis中是否有nonce
	nonce, err := services.UserService.GetNonce(address)
	if err != nil {
		log.Println("当前地址无nonce，请重新登录")
		ctx.JSON(http.StatusServiceUnavailable, err.ToErrResponse())
		return
	}

	//3.执行验签
	if isAccept, err := auth.VerifySignature(address, nonce, signature); !isAccept && err != nil {
		log.Printf("验签错误，错误为：%v", err)
		ctx.JSON(http.StatusUnauthorized, dto.ErrResponse{
			Code:    http.StatusUnauthorized,
			Message: "签名验证失败,请重新登录",
		})
		return
	}

	//4，确保用户存在，不存在执行创建
	if err := services.UserService.EnsureUserExists(address); err != nil {
		ctx.JSON(http.StatusServiceUnavailable, err.ToErrResponse())
		return
	}

	//5.生成token
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
//
//	@Summary		用户登出
//	@Description	用户退出登录，将当前 JWT 加入黑名单
//	@Tags			用户管理
//	@Produce		json
//
//	@Success		200	{object}	dto.Response	"成功登出"
//	@Failure		400	{object}	dto.ErrResponse	"服务器内部错误"
//
//	@Security		JwtAuth
//	@Router			/user/logout [post]
func (u *User) Logout(ctx *gin.Context) {
	jti, _ := ctx.Get("jti")
	//将jti加入redis黑名单
	err := configs.RedisClient.SetEX(ctx, "blacklist:"+jti.(string), "1", auth.TokenExpireDuration).Err()
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
//
//	@Summary		用户编辑用户名
//	@Description	根据用户地址更新用户名
//	@Tags			用户管理
//	@Accept			json
//	@Produce		json
//	@Security		JwtAuth
//
//	@Success		200				{object}	dto.Response	"用户名更改成功"
//	@Failure		400				{object}	dto.ErrResponse	"请求体格式错误"
//
//	@Param			Authorization	header		string			true	"JWT"
//	@Router			/user/edit/name [post]
func (u *User) EditUserName(ctx *gin.Context) {
	log.Println("进入编辑用户名接口！")
	var json dto.UpdateUser
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrResponse{
			Code:    http.StatusBadRequest,
			Message: "请求体格式错误",
		})
		return
	}

	newName := json.Name
	log.Println(" 来自post请求体的json的new_name:" + newName)
	toUpdate := dto.UpdateUser{Name: newName}
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

// UploadProfile
//
//	@Summary		上传用户头像
//	@Description	根据用户地址更新用户头像
//	@Tags			用户管理
//	@Accept			multipart/form-data
//	@Produce		image/png
//	@Security		JwtAuth
//	@Param			Authorization	header		string			true	"JWT"
//	@Param			profile			formData	file			true	"用户头像文件"
//	@Param			user_address	formData	string			true	"用户地址"
//	@Success		200				{object}	dto.Response	"头像上传成功"
//	@Failure		400				{object}	dto.ErrResponse	"请求体格式错误"
//	@Failure		503				{object}	dto.ErrResponse	"服务不可用，数据库异常"
//	@Router			/user/upload/avatar [post]
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
	err = configs.MinioClient.RemoveObject(context.Background(), "test", pictureName, minio.RemoveObjectOptions{})

	if err != nil {
		log.Println("上传失败！！\n", err)
	}

	fileSize := header.Size
	// todo 处理文件，例如保存到minio服务器,桶可以考虑用地址哈希！
	// 上传文件到 MinIO
	_, err = configs.MinioClient.PutObject(
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

	var toUpdate = dto.UpdateUser{Avatar: "/test/" + pictureName}
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
//
//	@Summary		获取用户基本信息
//	@Description	根据用户地址获取用户基本信息
//	@Tags			用户管理
//	@Produce		json
//	@Security		JwtAuth
//	@Param			Authorization	header		string			true	"JWT"
//	@Success		200				{object}	dto.Response	"响应用户基本信息"
//	@Failure		503				{object}	dto.ErrResponse	"mysql不可用"
//	@Router			/user/info [post]
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

// GetProfileOfUser 获取用户头像
//
//	@Summary		获取用户头像
//	@Description	根据用户地址获取用户头像
//	@Tags			用户管理
//	@Accept			json
//	@Produce		image/png
//	@Security		JwtAuth
//	@Param			Authorization	header		string			true	"JWT"
//	@Param			padd			query		string			false	"用户头像的路径"	default("1.png")
//	@Success		200				{file}		file			"用户头像图片文件"
//	@Failure		400				{object}	dto.ErrResponse	"请求体格式错误"
//	@Failure		503				{object}	dto.ErrResponse	"服务不可用，数据库异常"
//	@Router			/user/profile [get]
func (u *User) GetProfileOfUser(c *gin.Context) {

	// 使用Query方法获取参数，如果参数不存在则返回默认值
	//name := c.Query("name")
	p_add := c.DefaultQuery("padd", "1.png") // 默认值是20
	addresses := strings.Split(p_add, "/")   //因为参数长/xxx/sss.png,所以第一个是空的！
	// 获取对象
	object, err := configs.MinioClient.GetObject(context.Background(), addresses[1], addresses[2], minio.GetObjectOptions{})
	if err != nil {
		log.Println("minio获取对象出错！", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取对象"})
		return
	}
	defer object.Close()

	// 读取对象数据
	data, err := io.ReadAll(object)
	if err != nil {
		log.Println("读取对象数据出错！", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法读取对象数据"})
		return
	}

	// 设置响应头
	c.Header("Content-Type", "image/png")
	c.Header("Content-Disposition", "inline; filename=1.png")
	c.Data(http.StatusOK, "image/png", data)
}

// GetGNFTList
//
//	@Summary		获取用户的藏品（对他可见的数据）
//	@Description	根据用户地址更新用户名
//	@Tags			用户管理
//	@Accept			json
//	@Produce		json
//	@Security		JwtAuth
//	@Param			Authorization	header		string			true	"JWT"
//	@Success		200				{object}	dto.Response	"用户名更新成功"
//	@Failure		400				{object}	dto.ErrResponse	"请求体格式错误"
//	@Failure		503				{object}	dto.ErrResponse	"mysql异常"
//	@Router			/user/gnfts [post]
func (u *User) GetGNFTList(ctx *gin.Context) {

}

//====================================OAuth2======================================

func (u *User) Oauth2Wegene(ctx *gin.Context) {
	fmt.Println("开始重定向到wegene授权页面")
	ctx.Redirect(http.StatusMovedPermanently, "https://api.wegene.com/authorize/?redirect_uri="+
		"http://localhost:8080/user/receiveCode&response_type=code&client_id=szjsbiolab&"+
		"scope=basic rs123 athletigen skin psychology risk health ancestry haplogroups demographics web")
}

func (u *User) ReceiveCode(ctx *gin.Context) {

	//输出请求所有得东西
	fmt.Println(ctx.Request.URL.Query())

	code := ctx.Query("code")
	fmt.Println("（（（（（（（（（（（（（（（（授权码：", code, "））））））））））））））））））")

	u.GetWegeneToken(code)
	if code == "" {
		fmt.Println("授权码为空，第二次进入这个接口，无需重定向！")
		return
	}
	ctx.Redirect(301, "http://localhost:8080/swagger/index.html#/")

}

func (u *User) GetWegeneToken(code string) {

	// 设置请求的URL
	getToknUrl := "https://api.wegene.com/token/"

	// 设置请求参数
	data := url.Values{}
	fmt.Println("code:", code)
	data.Set("client_id", configs.WegeneId)
	data.Set("client_secret", configs.WegeneSecret)
	//grant_type 参数：设置为 authorization_code，这是 OAuth2 授权码模式。
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	//这里也可以和上面不一样！
	data.Set("redirect_uri", "http://localhost:8080/user/receiveCode")
	data.Set("scope", "basic rs123 athletigen skin psychology risk health ancestry haplogroups demographics web") //这里的权限范围需要和上面重定向的一样

	// 创建POST请求
	req, err := http.NewRequest("POST", getToknUrl, bytes.NewBufferString(data.Encode()))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	fmt.Println("请求token！")
	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// 打印响应
	fmt.Println("Response:", string(body))
}

/*
Response: {"access_token":"a3a401b429f273b7f2e515f4b8d51379","token_type":"bearer",
"expires_in":86400,"refresh_token":"62ca9e2733526673f7fb2d77fa82a9ef",
"scope":"basic rs123 athletigen skin psychology risk health ancestry haplogroups demographics web"}

Bearer a3a401b429f273b7f2e515f4b8d51379
*/
