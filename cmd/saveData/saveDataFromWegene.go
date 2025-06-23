package main

import (
	"GeneReport_platform/api/dto"
	"GeneReport_platform/configs"
	"GeneReport_platform/internal/setup"
	"GeneReport_platform/pkg/rocketmq"
	"fmt"
)

func main() {

	setup.Setup()
	fmt.Println("数据库迁移开始")
	// 自动迁移
	configs.DB.AutoMigrate(
		&dto.Psychology{},
		&dto.Skin{},
		&dto.Athletigen{},
		//&dto.HealthyDrug{},
		&dto.HealthyTraits{},
		&dto.HealthyCarrier{},
		&dto.HealthyMetabolism{},
		&dto.HealthResult{},
		&dto.Risk{},
		&dto.Ancestry{},
		&dto.Haplogroups{},
		&dto.Demographics{},
		&dto.Genotype{},
		&dto.UniqueProfiles{},
		&dto.Metadatas{}, //fixme 这里的Metadatas自动迁移有问题
		&dto.DataVisitRecord{},
	)

	rocketmq.Myproducer("ptest")

	//tools.SaveDataTest("660ac4359ca7fd5d0b4b9a121c88507b")
	//return

	go rocketmq.Consumer(rocketmq.HandleMsgPrint, "one", "testCon")      //启动测试消费者监听
	go rocketmq.Consumer(rocketmq.HandleMsgSaveData, "save", "saveData") //监听保存数据的消费者
	//从标准输入读取数据
	for {
		var msg string
		fmt.Scanln(&msg)
		rocketmq.SendMsg("testCon", msg)
	}
	// 阻塞主线程，防止程序立即退出
	select {}

}
