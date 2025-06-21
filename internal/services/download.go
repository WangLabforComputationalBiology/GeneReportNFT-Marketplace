package services

import (
	"GeneReport_platform/configs"
	"GeneReport_platform/pkg/appErrors"
	"GeneReport_platform/tools/utils"
	"context"
	"net/http"
	"time"
)

var (
	DownloadServ *DownloadService
)

type DownloadService struct {
	iDownloadBase
}

// collection基础接口
type iDownloadBase interface {
}

func RegisterDownloadService() {
	DownloadServ = &DownloadService{}
}

func (d *DownloadService) GenerateDownloadLink(dataHash, userAddress, pubKey string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var shortCode string
	//生成下载短码,并存储到redis
	for i := 0; i < 3; i++ { //最多尝试3次
		shortCode = utils.GenerateShortCode()
		getAllCmd := configs.RedisClient.HGetAll(context.Background(), shortCode)

		if getAllCmd.Err() != nil { //内部错误
			return "", appErrors.New(http.StatusInternalServerError, "redis error", getAllCmd.Err())
		} else if getAllCmd.Val() == nil { //短码不存在或已过期，该短码可用
			break
		} else if i == 2 { //尝试3次后仍然失败
			return "", appErrors.New(http.StatusInternalServerError, "generate short code failed")
		}
	}
	//找到可用短码后，存储到redis
	pipe := configs.RedisClient.Pipeline()

	// 添加 HSet 命令
	setCmd := pipe.HSet(ctx, shortCode, map[string]interface{}{
		"data_hash":    dataHash,
		"user_address": userAddress,
		"pub_key":      pubKey,
	})
	pipe.Expire(ctx, shortCode, 5*time.Minute)
	// 执行 Pipeline
	_, err := pipe.Exec(ctx)
	if err != nil || setCmd.Err() != nil {
		return "", appErrors.New(http.StatusInternalServerError, "服务繁忙，请稍后再试", err)
	}

	return "dl/" + shortCode, nil
}
