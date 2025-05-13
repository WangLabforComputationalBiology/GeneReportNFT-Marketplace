package main

import (
	"GeneReport_platform/configs"
	"GeneReport_platform/internal/models"
	"GeneReport_platform/internal/setup"
	"log"
)

var modelList = []interface{}{
	&models.GNFT{},    //文件
	&models.User{},    //用户
	&models.Listing{}, //权限
}

func main() {
	setup.Setup()
	// 自动迁移所有模型
	for _, model := range modelList {
		if err := configs.DB.AutoMigrate(model); err != nil {
			log.Fatalf("Failed to migrate model %T: %v", model, err)
		}
	}
	log.Println("Migration complete!")
}
