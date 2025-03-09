package controllers

import (
	"GeneReport_platform/internal/dao/global"
	"GeneReport_platform/tools/utils"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"io"
	"log"
	"net/http"
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
