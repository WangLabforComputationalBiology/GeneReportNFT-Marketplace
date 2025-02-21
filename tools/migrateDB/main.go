package main

import (
	"GeneReport_platform/init"
	"GeneReport_platform/internal/dao"
	"GeneReport_platform/internal/models/entity"
	"log"
)

var modelList = []interface{}{
	&entity.GNFT{},  //文件
	&entity.User{},  //用户
	&entity.Order{}, //权限
}

func main() {
	init.Init()
	// 自动迁移所有模型
	for _, model := range modelList {
		if err := dao.AutoMigrate(model); err != nil {
			log.Fatalf("Failed to migrate model %T: %v", model, err)
		}
	}
	log.Println("Migration complete!")
}
