package controllers

import (
	"GeneReport_platform/api/dto"
	"GeneReport_platform/internal/dao/global"
	"GeneReport_platform/tools/utils"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
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
	address := ctx.Param("user_address")

	//fixme 在这里要验证用户是否为新用户，是的话需要写进MySQL！

	jwt, _ := utils.GenToken(address)
	// 设置 key 并指定过期时间为 3分钟
	err := global.RedisClient.Set(ctx, address, jwt, time.Minute*3).Err()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "redis服务出现问题！"})
		return
	}
	//todo 测试1
	token, err := utils.ParseToken(jwt)
	log.Println(" 用户地址是： ", token.User_address)
	ctx.JSON(200, gin.H{"account": address, "jwt": jwt})
}

// 发送验签-login接口
func (u *User) LogIng(ctx *gin.Context) {
	log.Println("进入登录接口！")
	var json dto.EditUserName
	if err := ctx.ShouldBindJSON(&json); err != nil {
		// 处理错误，比如返回错误信息
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	address := json.Address
	signature := json.Signature

	log.Println("请求里的地址是："+address, "请求里的签名是："+signature)
	jwt, _ := utils.GenToken(address, 3)

	// 设置 key 并指定过期时间为 3分钟
	err := global.RedisClient.Set(ctx, address, jwt, time.Minute*3).Err()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "redis服务出现问题！"})
		return
	}

	//todo
	ctx.JSON(200, gin.H{"reflash_token": "这东西是什么，我不知道啊，文档没说清楚", "access_token": jwt})
}

// 登出操作
func (u *User) LogOut(ctx *gin.Context) {
	var json dto.EditUserName
	if err := ctx.ShouldBindJSON(&json); err != nil {
		// 处理错误，比如返回错误信息
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	address := json.Address

	//将redis里的jwt删除！
	err := global.RedisClient.Del(ctx, address).Err()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "redis服务出现问题！"})
		return
	}
	ctx.JSON(200, gin.H{"msg:": "成功退出登录！"})
}

// 编辑用户名称
func (u *User) EditUserName(ctx *gin.Context) {
	log.Println("进入编辑用户名接口！")
	var json dto.EditUserName
	if err := ctx.ShouldBindJSON(&json); err != nil {
		// 处理错误，比如返回错误信息
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	name := json.Name
	log.Println(" 来自post请求体的json的new_name:" + name)

	var update dto.UpdateUser = dto.UpdateUser{Name: name, Address: json.Address}

	// 修改数据库的内容
	if err := global.DB.Model(&dto.Users{}).Where("address = ?", json.Address).Updates(update).Error; err != nil {
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

	var update dto.UpdateUser = dto.UpdateUser{Picture: "/test/" + pictureName}
	// 修改数据库的内容
	if err := global.DB.Model(&dto.Users{}).Where("address = ?", address).Updates(update).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "更新用户名失败"})
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

// 返回用户信息
func (u *User) GetInfo(ctx *gin.Context) {
	// 获取请求头中的 Authorization 值
	address := ctx.GetHeader("Authorization")

	// 创建一个 Users 实例用于接收查询结果
	var user dto.Users

	// 根据地址查询用户
	result := global.DB.Where("address = ?", address).First(&user)

	// 检查错误
	if result.Error != nil {
		// 如果没有找到记录，result.Error 将会是 gorm.ErrRecordNotFound
		if result.Error == gorm.ErrRecordNotFound {
			fmt.Println("没有找到地址为", address, "的用户")
		} else {
			fmt.Println("查询出错:", result.Error)
		}
		return
	}

	// 打印查询到的用户信息
	fmt.Printf("找到用户: %+v\n", user)
	ctx.JSON(200, gin.H{"user": user})
}

// 返回用户藏品信息
func (u *User) GetGNFTList(ctx *gin.Context) {

}
