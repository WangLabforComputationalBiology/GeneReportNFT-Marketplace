package main

import (
	"GeneReport_platform/api/router"
	. "GeneReport_platform/init"
	"github.com/gin-gonic/gin"
)

// 项目入口
func main() {

	Init()
	r := gin.Default()
	r = router.SetupRouter()
	_ = r.Run(GlobalConfig.AppConfig.Addr)

}
