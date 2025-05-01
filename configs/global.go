package configs

import (
	"github.com/go-redis/redis/v8"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB

	RedisClient *redis.Client
	MinioClient *minio.Client

	// 下面是etherscan的配置：
	IsProxy  bool
	ProxyUrl string
	ApiKey   string
	EndPoint string

	WegeneId           string
	WegeneSecret       string
	WegeneRedirectHost string
)
