package rocketmq

import (
	"GeneReport_platform/configs"
	"GeneReport_platform/tools"
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"os"
	"strings"
)

// 定义一个处理消息的函数类型
type MessageHandler func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error)

func Consumer(handler MessageHandler, group, topic string) {
	host := configs.GlobalConfig.RocketMqCfg.NameServer
	//启动recketmq并设置负载均衡的Group
	c, _ := rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{host}),
		consumer.WithGroupName(group),
	)
	//订阅消息方式1：
	//if err := c.Subscribe(topic, consumer.MessageSelector{}, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
	//	for i := range msgs {
	//		/*todo msg 长这样
	//		[Message=[topic=test, body=牛逼, Flag=0, properties=map[CLUSTER:DefaultCluster
	//		CONSUME_START_TIME:1744811100480 MAX_OFFSET:1 MIN_OFFSET:0 MSG_REGION:DefaultRegion
	//		TRACE_ON:true UNIQ_KEY:C0A85501060C0000000051ea27600003], TransactionId=],
	//		MsgId=C0A85501060C0000000051ea27600003, OffsetMsgId=7818A88400002A9F000000000003C15A,QueueId=3,
	//		StoreSize=205, QueueOffset=0, SysFlag=0, BornTimestamp=1744811100347, BornHost=112.97.85.110:12367,
	//		StoreTimestamp=1744811100389, StoreHost=120.24.168.132:10911, CommitLogOffset=246106, BodyCRC=861123100,
	//		ReconsumeTimes=0, PreparedTransactionOffset=0]
	//		*/
	//		fmt.Printf("测试消费者处理消息: %v \n", string(msgs[i].Body))
	//
	//	}
	//	return consumer.ConsumeSuccess, nil
	//}); err != nil {
	//	fmt.Println("订阅主题失败:", err)
	//	fmt.Println(err.Error())
	//}
	// 订阅消息方式2：
	if err := c.Subscribe(topic, consumer.MessageSelector{}, handler); err != nil {
		fmt.Println("订阅主题失败:", err)
		fmt.Println(err.Error())
	}
	//启动
	err := c.Start()
	if err != nil {
		fmt.Println("启动测试消费者失败:", err)
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	fmt.Println("测试消费者已启动，正在监听消息...")

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
func HandleMsgPrint(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
	for i := range msgs {
		fmt.Printf("测试消费者处理消息: %v \n", string(msgs[i].Body))
	}
	return consumer.ConsumeSuccess, nil
}

func HandleMsgSaveData(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
	for i := range msgs {
		fmt.Printf("保存消息的消费者处理消息: %v \n", string(msgs[i].Body))
		//保存数据
		//将msgs[i].Body的数据按照":"分割 token+id
		parts := strings.Split(string(msgs[i].Body), ":")

		fmt.Printf("token:%s\nprofileId:%s", parts[0], parts[1])
		//todo 此处可以做重复新检测，表需要一个状态位来记录他的完成情况
		tools.SaveAllData(parts[0], parts[1])
		//tools.SaveDataTest(parts[0])
	}
	return consumer.ConsumeSuccess, nil
}
