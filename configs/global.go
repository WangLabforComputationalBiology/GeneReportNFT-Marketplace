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

	WegeneId           string
	WegeneSecret       string
	WegeneRedirectHost string
)
