package main

import (
	"GeneReport_platform/api/router"
	"GeneReport_platform/internal/setup"
	"GeneReport_platform/pkg/rocketmq"
	"github.com/gin-gonic/gin"
)

// @title			API接口文档
// @version		1.0
// @description	GeneReport平台
// @host			120.24.168.132:8080
// @BasePath		/
// @schemes		http https
func main() {

	setup.Setup()
	r := gin.Default()
	r = router.SetupRouter()

	_ = r.Run("0.0.0.0:7070")       //在配置的端口运行
	go rocketmq.Myproducer("psave") //启动生产者
}
