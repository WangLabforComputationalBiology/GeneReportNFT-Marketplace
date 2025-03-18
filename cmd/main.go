package main

import (
	"GeneReport_platform/api/router"
	_ "GeneReport_platform/docs"
	. "GeneReport_platform/init"
	"github.com/gin-gonic/gin"
)

// @title           API接口文档
// @version         1.0
// @description     GeneReport平台
// @host            localhost:8080
// @BasePath        /docs
func main() {

	Init()
	r := gin.Default()
	//todo 确保中间件在路由前使用，否则不生效
	//r.Use(utils.JWTAuthMiddleware())
	r = router.SetupRouter()

	_ = r.Run(GlobalConfig.AppConfig.Addr) //在配置的端口运行

}
