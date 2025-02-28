package controllers

import (
	"GeneReport_platform/internal/dao/global"
	"GeneReport_platform/tools/utils"
	"github.com/gin-gonic/gin"
	"time"
)

type User struct{}

var UserController = User{}

func (u *User) Test(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "请求用户 controller successful！"})

}

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
	ctx.JSON(200, gin.H{"account": account, "jwt": jwt})
}

func (u *User) GetInfo(ctx *gin.Context) {

}

func (u *User) GetGNFTList(ctx *gin.Context) {

}
