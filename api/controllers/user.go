package controllers

import (
	"GeneReport_platform/tools/utils"
	"github.com/gin-gonic/gin"
)

type User struct{}

var UserController = User{}

func (u *User) Test(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "请求用户 controller successful！"})

}

func (u *User) GetNonce(ctx *gin.Context) {
	account := ctx.Param("account")
	jwt, _ := utils.GenToken(account)
	ctx.JSON(200, gin.H{"account": account, "jwt": jwt})
}

func (u *User) GetInfo(ctx *gin.Context) {

}

func (u *User) GetGNFTList(ctx *gin.Context) {

}
