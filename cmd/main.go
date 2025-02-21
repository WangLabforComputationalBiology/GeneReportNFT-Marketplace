package main

import (
	"GeneReport_platform/api/router"
	"GeneReport_platform/init"
)

// 项目入口
func main() {

	init.Init()
	r := router.SetupRouter()
	_ = r.Run(init.GlobalConfig.AppConfig.Port)

}
