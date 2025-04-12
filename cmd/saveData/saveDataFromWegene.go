package main

import (
	"GeneReport_platform/pkg/rocketmq"
	"fmt"
)

func main() {

	//rocketmq.Myproducer()//已经在init那里使用了
	go rocketmq.Myconsumer() //启动消费者监听

	//从标准输入读取数据
	for {
		var msg string
		fmt.Scanln(&msg)
		rocketmq.SendMsg(msg)
	}
	// 阻塞主线程，防止程序立即退出
	select {}

}
