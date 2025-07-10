package main

import (
	"GeneReport_platform/internal/contracts/sharingPlatformContract"
	"context"
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/v3/types"
	"log"
)

func main() {
	events := make(chan string, 1000)
	eventTopic := sharingPlatformContract.GetContractABI().Events["NewViewAccess"].ID().Hex()

	// 配置事件监听参数，忽略动态索引的 topics
	params := types.EventLogParams{
		FromBlock: 0,  // 从创世区块开始
		ToBlock:   -1, // 持续监听最新区块
		Addresses: []string{sharingPlatformContract.PlatformContractAddressHex},
		Topics:    []string{eventTopic}, // 仅设置事件签名，动态索引留空
	}

	// 定义回调函数处理事件日志
	handler := func(subscriptionID int, logs []types.Log) {
		for _, txLog := range logs {
			// 解析事件数据（示例：提取 topics 和 data）
			eventData := fmt.Sprintf("Event: subscription=%d, address=%s, topics=%v, data=%s",
				subscriptionID, txLog.Address, txLog.Topics, txLog.Data)

			select {
			case events <- eventData:
				log.Printf("Event sent to channel: %s", eventData)
			default:
				log.Printf("Event channel full, dropping event")
			}
		}
	}

	// 订阅事件
	ctx := context.Background()
	_, err := sharingPlatformContract.NewChainClient().SubscribeEventLogs(ctx, params, handler)
	if err != nil {
		log.Printf("failed to subscribe to events: %v", err)
	}

	// 保持监听运行
	go func() {
		<-ctx.Done()
		log.Println("Event listener stopped")
	}()
	select {}
}
