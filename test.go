package main

import (
	"GeneReport_platform/internal/contracts/sharingPlatformContract"
	"fmt"
)

func main() {

	//txHash := ""
	//
	////获取交易回执
	//receipt, err := sharingPlatformContract.NewChainClient().GetTransactionReceipt(context.Background(), common.HexToHash(txHash), true)
	//if err != nil {
	//	return
	//}
	//
	//// 定义事件签名和哈希
	//eventSignatures := map[string]string{
	//	"NewViewAccess":     "0x" + common.Bytes2Hex(crypto.Keccak256([]byte("NewViewAccess(address,bytes32,uint256)"))),
	//	"RenewalViewAccess": "0x" + common.Bytes2Hex(crypto.Keccak256([]byte("RenewalViewAccess(address,bytes32,uint256)"))),
	//}
	eventRawName := sharingPlatformContract.GetContractABI().Events["NewViewAccess"].RawName
	eventName := sharingPlatformContract.GetContractABI().Events["NewViewAccess"].Name

	fmt.Println(eventRawName)
	fmt.Println(eventName)
}

//// EventListener 事件监听器
//type EventListener struct {
//	client *BlockchainClient
//	events chan string
//}
//
//func (e *EventListener) ListenEvents(ctx context.Context, contractAddress string, eventSignature string) error {
//	// 计算事件签名哈希
//	eventTopic := sharingPlatformContract.GetContractABI().Events["NewViewAccess"].RawName
//
//	// 配置事件监听参数，忽略动态索引的 topics
//	params := types.EventLogParams{
//		FromBlock: 0,  // 从创世区块开始
//		ToBlock:   -1, // 持续监听最新区块
//		Addresses: []string{contractAddress},
//		Topics:    []string{eventTopic, "", ""}, // 仅设置事件签名，动态索引留空
//	}
//
//	// 定义回调函数处理事件日志
//	handler := func(logs []types.Log) {
//		for _, log := range logs {
//			// 解析事件数据（示例：提取 topics 和 data）
//			eventData := fmt.Sprintf("Event: subscription=%d, address=%s, topics=%v, data=%s",
//				subscriptionID, log.Address, log.Topics, log.Data)
//			select {
//			case e.events <- eventData:
//				log.Printf("Event sent to channel: %s", eventData)
//			default:
//				log.Println("Event channel full, dropping event")
//			}
//		}
//	}
//
//	// 订阅事件
//	_, err := e.client.conn.SubscribeEventLogs(params, handler)
//	if err != nil {
//		return fmt.Errorf("failed to subscribe to events: %v", err)
//	}
//
//	// 保持监听运行
//	go func() {
//		<-ctx.Done()
//		log.Println("Event listener stopped")
//	}()
//
//	return nil
//}
