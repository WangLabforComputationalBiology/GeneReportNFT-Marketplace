package global

import (
	"context"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	DB          *gorm.DB
	Ctx         context.Context
	RedisClient *redis.Client
)
