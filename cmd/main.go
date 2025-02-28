package main

import (
	"GeneReport_platform/api/router"
	. "GeneReport_platform/init"
	"github.com/gin-gonic/gin"
)

// 项目入口
// @title 基因区块链社区
func main() {

	Init()
	r := gin.Default()
	//todo 确保中间件在路由前使用，否则不生效
	//r.Use(utils.JWTAuthMiddleware())
	r = router.SetupRouter()

	_ = r.Run(GlobalConfig.AppConfig.Addr) //在配置的端口运行

}
