package global

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
)

var (
	DB          *gorm.DB
	Ctx         context.Context
	CtxCancel   context.CancelFunc
	RedisClient *redis.Client
	MinioClient *minio.Client

	// 下面是etherscan的配置：
	IsProxy  bool
	ProxyUrl string
	ApiKey   string
	EndPoint string
)
