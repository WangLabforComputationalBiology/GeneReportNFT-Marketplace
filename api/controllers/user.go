package controllers

import (
	"GeneReport_platform/api/dto"
	"GeneReport_platform/configs"
	"GeneReport_platform/internal/services"
	"GeneReport_platform/pkg/auth"
	"GeneReport_platform/pkg/rocketmq"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
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

	var req dto.LoginReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrResponse{
			Code:    http.StatusBadRequest,
			Message: "请求体格式错误,请重新登录",
		})
		return
	}

	address := req.UserAddress
	signature := req.Signature

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
	var req dto.UpdateUser
	if err := ctx.ShouldBindJSON(&req); err != nil || req.Name == "" {
		ctx.JSON(http.StatusBadRequest, dto.ErrResponse{
			Code:    http.StatusBadRequest,
			Message: "请求体格式错误",
		})
		return
	}

	newName := req.Name

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

// UploadAvatar
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
func (u *User) UploadAvatar(ctx *gin.Context) {
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
		ctx.Error(err)
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

// GetGNFTList 获取用户已获取到数据的GNFT列表
func (u *User) GetGNFTList(ctx *gin.Context) {

}

//====================================OAuth2======================================

func (u *User) Oauth2Wegene(ctx *gin.Context) {
	fmt.Println("开始重定向到wegene授权页面")
	ctx.Redirect(http.StatusMovedPermanently, "https://api.wegene.com/authorize/?redirect_uri="+
		"http://localhost:8080/user/receiveCode&response_type=code&client_id=szjsbiolab&"+
		"scope=basic rs123 athletigen skin psychology risk health ancestry haplogroups demographics web"+
		" names email")
}

func (u *User) ReceiveCode(ctx *gin.Context) {

	//输出请求所有得东西
	fmt.Println(ctx.Request.URL.Query())

	code := ctx.Query("code")
	fmt.Println("（（（（（（（（（（（（（（（（授权码：", code, "））））））））））））））））））")

	token := u.GetWegeneToken(code)

	//生成一个uuid
	uuid := uuid.New().String()
	//将uuid和toen的映射存到redis，5分钟后过期
	configs.RedisClient.Set(ctx, uuid, token, 5*time.Minute)
	fmt.Println("存到redis的code：token= ", code, "------", token)
	if code == "" {
		fmt.Println("授权码为空，第二次进入这个接口，无需重定向！")
		return
	}
	ctx.Redirect(http.StatusMovedPermanently, "http://localhost:5173/create/selectProfile/"+uuid)

}

// GetWegeneToken 根据授权码获取token
func (u *User) GetWegeneToken(code string) (tkn string) {

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
	//data.Set("redirect_uri", "http://localhost:8080/user/receiveCode") //可有可无
	//这里的权限范围是上面重定向的子集，只能少不能多，可以用户自己选
	//fixme 这里可以让用户选择授权的范围
	data.Set("scope", "basic rs123 athletigen skin psychology risk health ancestry haplogroups demographics web names email")
	//上面的范围如果有names和emali那么/user的响应会多出这两个值

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

	var token dto.GetToken
	if err := json.Unmarshal(body, &token); err != nil {
		fmt.Println("Error unmarshaling response:", err)
		return
	}

	//返回token
	tkn = token.AccessToken
	return

}

// 拿着token取请求profile,基因报告
func getReportId(token string) (usersProfile dto.GetReportId) {

	url := "https://api.wegene.com/user/"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Bearer "+token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))

	err = json.Unmarshal(body, &usersProfile)
	if err != nil {
		fmt.Println("解析获取reportId的响应数据出错！" + err.Error())
	}

	return
}

// SendSMSCode 通过手机号码获取验证码并存入redis
func (u *User) SendSMSCode(ctx *gin.Context) {
	var req dto.SendSMSCodeReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrResponse{
			Code:    http.StatusBadRequest,
			Message: "请求体格式错误,请检查",
		})
		return
	}
	if err := services.UserService.SendSMSCode(req.Phone); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.ToErrResponse())
	}
}

// VerifySMSCode 验证手机验证码
func (u *User) VerifySMSCode(ctx *gin.Context) {
	var req dto.VerifySMSCodeReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrResponse{
			Code:    http.StatusBadRequest,
			Message: "请求体格式错误,请检查",
		})
		return
	}
	if isPass, err := services.UserService.VerifySMSCode(req.Phone, req.Code); err == nil {
		if isPass {
			ctx.JSON(http.StatusOK, dto.Response{Code: 200, Message: "验证成功"})
		} else {
			ctx.JSON(http.StatusBadRequest, dto.ErrResponse{
				Code:    http.StatusBadRequest,
				Message: "验证码错误",
			})
		}
	} else {
		ctx.JSON(http.StatusInternalServerError, err.ToErrResponse())
	}
}

// GetUsersProfileByCode 重定向将token的kv映射传给前端，前端那这个key请求基因报告数据供用户选择
func (u *User) GetUsersProfileByCode(ctx *gin.Context) {
	//在get请求路径里面获取code
	code := ctx.Query("code")
	fmt.Println("code:", code)
	//根据code在redis获取token
	token := configs.RedisClient.Get(context.Background(), code).Val()
	//在go里面，如果在redis拿不到值，将会是""!!
	if token == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "授权过期或未授权！"})
		return
	}
	usersProfile := getReportId(token)
	//创建一个数组将usersProfile.Profiles[x].id中的每个元素都添加到数组中
	var profiles []string
	for _, v := range usersProfile.Profiles {
		profiles = append(profiles, v.Id)
	}
	//由于数据有限，所以在里面添加几个假数据
	profiles = append(profiles, "1false", "2false", "3false", "4false", "5false", "6false")
	ctx.JSON(http.StatusOK, gin.H{
		"profiles": profiles,
	})
}

func (u *User) SaveProfileInfo(ctx *gin.Context) {
	//获取post请求json里面的数据
	var toSave dto.ToSave
	if err := ctx.ShouldBindJSON(&toSave); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	code := toSave.Code
	//根据code在redis获取token
	token := configs.RedisClient.Get(context.Background(), code).Val()
	//拿到后删除redis的记录
	configs.RedisClient.Del(context.Background(), code)
	sendMsg := token + ":" + toSave.ProfileId
	//异步保存数据
	rocketmq.SendMsg("saveData", sendMsg)

	fmt.Println("成功！异步保存数据：", sendMsg)
	ctx.JSON(http.StatusOK, gin.H{"msg": "successful!"})

}
