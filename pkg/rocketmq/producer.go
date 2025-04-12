package rocketmq

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

var Myp rocketmq.Producer

func init() {
	Myproducer()
}
func Myproducer() {
	//连接recketmq
	p, err := rocketmq.NewProducer(producer.WithNameServer([]string{"120.24.168.132:9876"}))
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
		fmt.Printf("send message error: %s\n", err)
	} else {
		fmt.Printf("send message success: result=%s\n", res.String())
	}
	Myp = p
	//关闭连接
	//err = p.Shutdown()
	//if err != nil {
	//	fmt.Printf("shutdown producer error: %s", err.Error())
	//}
}

func SendMsg(msg string) {
	//实例化消息
	msg1 := &primitive.Message{
		Topic: "test",
		Body:  []byte(msg),
	}
	//同步发送
	res, err := Myp.SendSync(context.Background(), msg1)
	if err != nil {
		fmt.Printf("send message error: %s\n", err)
	} else {
		fmt.Printf("send message success: result=%s\n", res.String())
	}
}
