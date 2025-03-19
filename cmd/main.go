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
// @BasePath        /
// @schemes         http
func main() {

	Init()
	r := gin.Default()
	r = router.SetupRouter()

	_ = r.Run(GlobalConfig.AppConfig.Addr) //在配置的端口运行

}
