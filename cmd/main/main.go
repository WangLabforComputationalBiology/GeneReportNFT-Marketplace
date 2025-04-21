package main

import (
	"GeneReport_platform/api/router"
	"GeneReport_platform/configs"
	_ "GeneReport_platform/docs" //不写的话访问/swagger/index.html#/会报错！！
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

	go rocketmq.Myproducer("psave")                //启动生产者
	_ = r.Run(configs.GlobalConfig.AppConfig.Addr) //在配置的端口运行

}
