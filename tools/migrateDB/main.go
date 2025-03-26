package main

import (
	"GeneReport_platform/init"
	"GeneReport_platform/internal/dao/global"
	"GeneReport_platform/internal/models"
	"log"
)

var modelList = []interface{}{
	&models.GNFT{},  //文件
	&models.User{},  //用户
	&models.Order{}, //权限
}

func main() {
	init.Init()
	// 自动迁移所有模型
	for _, model := range modelList {
		if err := global.DB.AutoMigrate(model); err != nil {
			log.Fatalf("Failed to migrate model %T: %v", model, err)
		}
	}
	log.Println("Migration complete!")
}
