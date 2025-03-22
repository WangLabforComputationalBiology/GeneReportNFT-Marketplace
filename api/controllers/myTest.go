package controllers

import (
	. "GeneReport_platform/api/dto"
	"GeneReport_platform/internal/dao/global"
	"GeneReport_platform/tools/utils"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

func MyTestRoute(r *gin.Engine) {

	mytest := r.Group("/test")
	{
		mytest.POST("/encryption", encryptionTxt)
		mytest.GET("/getid", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"id": utils.GetID()})
		})
		mytest.GET("/minio", TestMinio)
	}
}

// 可以返回请求中文件的哈希值1
func encryptionTxt(c *gin.Context) {
	file, err := c.FormFile("files")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "没有文件"})
		return
	}

	//生成文件的哈希散列值
	hashString, ok := utils.FileSHA256(file)
	if !ok {
		c.JSON(http.StatusBadRequest, "文件加密失败")
	}

	c.JSON(http.StatusOK, gin.H{
		"filename": file.Filename,
		"content":  hashString,
	})

}
func TestMinio(c *gin.Context) {
	// 获取对象
	object, err := global.MinioClient.GetObject(context.Background(), "test", "1.png", minio.GetObjectOptions{})
	if err != nil {
		log.Println("minio获取对象出错！", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取对象"})
		return
	}
	defer object.Close()

	// 读取对象数据
	data, err := io.ReadAll(object)
	if err != nil {
		log.Println("读取对象数据出错！", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法读取对象数据"})
		return
	}

	// 设置响应头
	c.Header("Content-Type", "image/png")
	c.Header("Content-Disposition", "inline; filename=1.png")
	c.Data(http.StatusOK, "image/png", data)
}

// GetTransactionInfo
//
//	@Summary		摘要：根据哈希获取交易信息,参数是交易的哈希值
//	@Description	描述信息：根据哈希获取交易信息,参数是交易的哈希值
//	@Param			txHash	query		string			true	"交易的哈希值"
//	@Success		200		{object}	dto.Response	"成功响应nonce"
//	@Router			/test/txinfo [get]
func GetTransactionInfo(ctx *gin.Context) {

	// 替换为你的 Etherscan API Key
	//apiKey := "PS6539ETQH1ZKUE3FVIYH6BVFHE3KT93RB"
	apiKey := global.ApiKey
	// 替换为你要查询的交易哈希
	//txHash := "0x7815bf96f6ca2dd42fd518dc79ca6ce3d19deb4d9a8f1e6582f8811d2ebd032e"
	txHash := ctx.Query("txhash")
	// Etherscan API 请求 URL
	//testurl := fmt.Sprintf("https://api-sepolia.etherscan.io/api?module=proxy&action=eth_getTransactionByHash&txhash=%s&apikey=%s", txHash, apiKey)
	testurl := fmt.Sprintf(global.EndPoint, txHash, apiKey)
	log.Printf("=========测试网址=========>>>>>>>", testurl)
	//fixme 让请求走代理
	//配置文件是否开启代理
	if global.IsProxy {
		//proxyURL, _ := url.Parse("http://127.0.0.1:7897") // Clash HTTP 代理端口
		proxyURL, _ := url.Parse(global.ProxyUrl)
		http.DefaultTransport = &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	}
	// 设置HTTP客户端的超时时间5秒
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	// 发送 HTTP 请求
	resp, err := client.Get(testurl)
	if err != nil {
		log.Printf("HTTP请求失败: %v", err) //不能使用log.Fatal(err)，因为这个会终止程序
		ctx.JSON(http.StatusGatewayTimeout, gin.H{"error": "请求超时或网络错误"})
		return
	}
	//defer resp.Body.Close()这个也可以关闭，但是不能捕获异常
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("关闭HTTP响应的Body异常！")
		}
	}(resp.Body)

	// 读取响应数据
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	var data RpcResponse
	if err := json.Unmarshal(body, &data); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, data)

}
