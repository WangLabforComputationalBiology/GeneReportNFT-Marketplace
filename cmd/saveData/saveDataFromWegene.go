package main

import (
	"GeneReport_platform/pkg/rocketmq"
	"GeneReport_platform/tools"
	"fmt"
)

/*
Response: {"access_token":"165ef5fc2c8a8a667831fba135b11427","token_type":"bearer",
"expires_in":86400,"refresh_token":"44426f83fbe18af2dd6bf74675c9d279",
"scope":"basic rs123 athletigen skin psychology risk health ancestry haplogroups demographics web"}
*/
func main() {

	//rocketmq.Myproducer()//已经在init那里使用了

	tools.SaveData("165ef5fc2c8a8a667831fba135b11427")

	return
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
