package dto

import "GeneReport_platform/internal/models"

// RpcResponse 用于调用ethscanner的api请求得到的响应数据
type RpcResponse struct {
	JSONRPC string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  Result `json:"result"`
}
type Result struct {
	BlockHash            string   `json:"blockHash"`
	BlockNumber          string   `json:"blockNumber"`
	From                 string   `json:"from"`
	Gas                  string   `json:"gas"`
	GasPrice             string   `json:"gasPrice"`
	MaxFeePerGas         string   `json:"maxFeePerGas"`
	MaxPriorityFeePerGas string   `json:"maxPriorityFeePerGas"`
	Hash                 string   `json:"hash"`
	Input                string   `json:"input"`
	Nonce                string   `json:"nonce"`
	To                   string   `json:"to"`
	TransactionIndex     string   `json:"transactionIndex"`
	Value                string   `json:"value"`
	Type                 string   `json:"type"`
	AccessList           []string `json:"accessList"`
	ChainId              string   `json:"chainId"`
	V                    string   `json:"v"`
	R                    string   `json:"r"`
	S                    string   `json:"s"`
	YParity              string   `json:"yParity"`
}

type GetGNFTInfoResp struct {
	GNFT       models.GNFT       `json:"gnft"`
	Collection models.Collection `json:"collection"`
	Quantity   int               `gorm:"column:quantity" json:"quantity"`
}

type GetGNFTInfosByOwnerReq struct {
	Owner string `json:"owner" binding:"required"`
}
type GetGNFTInfosByOwnerResp struct {
	GNFTs []GetGNFTInfoResp `json:"gnfts"`
}
