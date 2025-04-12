package rocketmq

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"os"
)

func Myconsumer() {
	//启动recketmq并设置负载均衡的Group
	c, _ := rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{"120.24.168.132:9876"}),
		consumer.WithGroupName("one"),
	)
	//订阅消息
	if err := c.Subscribe("test", consumer.MessageSelector{}, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for i := range msgs {
			fmt.Printf("subscribe callback: %v \n", msgs[i])
		}
		return consumer.ConsumeSuccess, nil
	}); err != nil {
		fmt.Println("订阅主题失败:", err)
		fmt.Println(err.Error())
	}
	//启动
	err := c.Start()
	if err != nil {
		fmt.Println("启动消费者失败:", err)
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	fmt.Println("消费者已启动，正在监听消息...")

	// 持续运行，直到手动停止
	select {} //select {} 是 Go 语言中的一种特殊语法，表示一个没有分支的 select 语句。它的作用是让当前 Goroutine 进入永久阻塞状态，直到程序被手动终止（例如通过信号或外部干预）。
	//阻塞就意味着他不能往下执行，直到被手动终止。
	//阻塞主线程
	//time.Sleep(time.Hour)
	//关闭连接
	//err = c.Shutdown()
	//if err != nil {
	//	fmt.Printf("shutdown Consumer error: %s", err.Error())
	//}
}
