package rocketmq

import (
	"GeneReport_platform/configs"
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

var Myp rocketmq.Producer

func Myproducer(group string) {
	host := configs.GlobalConfig.RocketMqCfg.NameServer
	// 连接 RocketMQ
	p, err := rocketmq.NewProducer(
		producer.WithNameServer([]string{host}),
		producer.WithGroupName(group), // 添加生产者组名称
	)
	if err != nil {
		fmt.Println("生成producer失败：", err)
	}
	//启动
	err = p.Start()
	if err != nil {
		fmt.Println("启动producer错误：", err)
	}
	//实例化消息
	msg := &primitive.Message{
		Topic: "test",
		Body:  []byte("this is ikun"),
	}
	//同步发送
	res, err := p.SendSync(context.Background(), msg)
	if err != nil {
		fmt.Printf("生产者发送消息失败: %s\n", err)
	} else {
		fmt.Printf("生产者发送消息成功: result=%s\n", res.String())
	}
	Myp = p
	//关闭连接
	//err = p.Shutdown()
	//if err != nil {
	//	fmt.Printf("shutdown producer error: %s", err.Error())
	//}
}

func SendMsg(topic, msg string) {
	//实例化消息
	msg1 := &primitive.Message{
		Topic: topic,
		Body:  []byte(msg),
	}
	//同步发送
	res, err := Myp.SendSync(context.Background(), msg1)
	if err != nil {
		fmt.Printf("生产者发送消息失败: %s\n", err)
	} else {
		fmt.Printf("生产者发送消息成功: result=%s\n", res.String())
	}
}
