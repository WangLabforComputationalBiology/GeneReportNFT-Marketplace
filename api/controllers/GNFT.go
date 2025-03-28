package controllers

import (
	. "GeneReport_platform/api/dto"
	"GeneReport_platform/internal/dao"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

type GNFT struct{}

var (
	GNFTController = &GNFT{}
)

func (cGNFT *GNFT) GetTransaction(txHash string) RpcResponse {
	testurl := fmt.Sprintf(dao.EndPoint, txHash, dao.ApiKey)
	if dao.IsProxy {
		proxyURL, _ := url.Parse(dao.ProxyUrl)
		http.DefaultTransport = &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	}
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(testurl)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	var data RpcResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return data
	}
	// 输出交易信息
	fmt.Println(data)
	return data
}

func (cGNFT *GNFT) GetInfo(ctx *gin.Context) {

}

func (cGNFT *GNFT) GetList(ctx *gin.Context) {

}

func (cGNFT *GNFT) Mint(ctx *gin.Context) {

}

// Burn 销毁GNFT
func (cGNFT *GNFT) Burn(ctx *gin.Context) {

}

// PutOnSale 上架GNFT
func (cGNFT *GNFT) PutOnSale() string {
	// 添加逻辑代码
	return "PutOnSale"
}

// TakeOffSale 下架GNFT
func (cGNFT *GNFT) TakeOffSale() string {
	// 添加逻辑代码
	return "TakeOffSale"
}
