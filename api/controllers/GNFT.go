package controllers

import (
	. "GeneReport_platform/api/dto"
	"GeneReport_platform/internal/dao/global"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type GNFTController struct{}

var (
	GNFTCtrller = &GNFTController{}
)

// 根据哈希获取交易信息
func (cGNFT *GNFTController) GetTransaction(txHash string) RpcResponse {

	testurl := fmt.Sprintf(global.EndPoint, txHash, global.ApiKey)
	proxyURL, _ := url.Parse(global.ProxyUrl)
	http.DefaultTransport = &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	resp, err := http.Get(testurl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var data RpcResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return data
	}
	// 输出交易信息
	fmt.Println(data)
	return data
}

func (cGNFT *GNFTController) GetInfo(ctx *gin.Context) {

}

func (cGNFT *GNFTController) GetList(ctx *gin.Context) {

}

func (cGNFT *GNFTController) Mint(ctx *gin.Context) {

}

// Burn 销毁GNFT
func (cGNFT *GNFTController) Burn(ctx *gin.Context) {

}

// PutOnSale 上架GNFT
func (cGNFT *GNFTController) PutOnSale() string {
	// 添加逻辑代码
	return "PutOnSale"
}

// TakeOffSale 下架GNFT
func (cGNFT *GNFTController) TakeOffSale() string {
	// 添加逻辑代码
	return "TakeOffSale"
}
