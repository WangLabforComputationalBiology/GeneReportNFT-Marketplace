package controllers

import (
	"GeneReport_platform/api/dto/user"
	"GeneReport_platform/internal/dao/global"
	"GeneReport_platform/tools/utils"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"io"
	"log"
	"net/http"
	"time"
)

type User struct{}

var UserController = User{}

func (u *User) Test(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "请求用户 controller successful！"})

}

// 请求nonce
func (u *User) GetNonce(ctx *gin.Context) {
	account := ctx.Param("user_address")

	//fixme 在这里要验证用户是否为新用户，是的话需要写进MySQL！

	jwt, _ := utils.GenToken(account)
	// 设置 key 并指定过期时间为 3分钟
	err := global.RedisClient.Set(ctx, account, jwt, time.Minute*3).Err()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "redis服务出现问题！"})
		return
	}
	//todo 测试1
	token, err := utils.ParseToken(jwt)
	log.Println(" 用户地址是： ", token.User_address)
	ctx.JSON(200, gin.H{"account": account, "jwt": jwt})
}

//发送验签

//登出操作

// 编辑用户名称
func (u *User) EditUserName(ctx *gin.Context) {
	log.Println("进入编辑用户名接口！")
	var json user.EditUserName
	if err := ctx.ShouldBindJSON(&json); err != nil {
		// 处理错误，比如返回错误信息
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	name := json.Name
	log.Println(" 来自post请求体的json的new_name:" + name)

	var update user.UpdateUserName = user.UpdateUserName{Name: name, Address: json.Address}

	// 修改数据库的内容
	if err := global.DB.Model(&user.UpdateUserName{}).Where("address = ?", json.Address).Updates(update).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "更新用户名失败"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "用户名更新成功"})

}

// 上传用户头像
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

	fileSize := header.Size
	// todo 处理文件，例如保存到minio服务器
	// 上传文件到 MinIO
	_, err = global.MinioClient.PutObject(
		context.Background(),
		"test", // 替换为你的桶名称
		address+".png",
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

	//todo 下面是把对象返回，测试用的
	// 读取对象数据
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
func (u *User) GetInfo(ctx *gin.Context) {

}

func (u *User) GetGNFTList(ctx *gin.Context) {

}
