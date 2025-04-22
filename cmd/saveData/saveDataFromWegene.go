package main

import (
	"GeneReport_platform/pkg/rocketmq"
	"fmt"
)

func main() {

	rocketmq.Myproducer("ptest")

	//tools.SaveDataTest("660ac4359ca7fd5d0b4b9a121c88507b")
	//return

	go rocketmq.Consumer(rocketmq.HandleMsgPrint, "one", "test")         //启动测试消费者监听
	go rocketmq.Consumer(rocketmq.HandleMsgSaveData, "save", "saveData") //监听保存数据的消费者
	//从标准输入读取数据
	for {
		var msg string
		fmt.Scanln(&msg)
		rocketmq.SendMsg("test", msg)
	}
	// 阻塞主线程，防止程序立即退出
	select {}

}
