package models

type TraceData struct {
	MetadataHash string `json:"metadata_hash"` // 元数据哈希
	Operator     string `json:"operator"`      // 操作者地址

	//追溯事件状态:
	//0：被operator创建,
	//1：operator通过SharingAddress新增查看许可,
	//2：operator续约查看许可,
	//3：在工坊被operator添加到SharingAddress,
	//4：operator修改共享状态
	OperationType int `json:"operation_type"`

	Timestamp    int64  `json:"timestamp"`     // 时间戳
	Remark       string `json:"remark"`        // 研究去向备注
	ExplorerLink string `json:"explorer_link"` // 链上交易链接
}
