package setup

import (
	"GeneReport_platform/configs"
	"GeneReport_platform/internal/dao"
	"GeneReport_platform/internal/services"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"strconv"
	"sync"
	"time"
)

func registerServices() {
	services.RegisterUserService()
	services.RegisterGNFTService()
	services.RegisterOrderService()
}
func registerDAO() {
	dao.RegisterUserDao()
	dao.RegisterGNFTDao()
	dao.RegisterOrderDao()
}
func setupMysql() {
	//初始化数据库连接
	dsn := configs.GlobalConfig.MysqlConfig.Dsn
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get sql db: %v", err)
	}
	sqlDB.SetMaxIdleConns(configs.GlobalConfig.MysqlConfig.MaxIdleConns)
	sqlDB.SetMaxOpenConns(configs.GlobalConfig.MysqlConfig.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour)
	dao.DB = db
}

func setupMinio() {
	url := configs.GlobalConfig.MinioConfig.Host + ":" + strconv.Itoa(configs.GlobalConfig.MinioConfig.Port)
	accessKey := configs.GlobalConfig.MinioConfig.AccessKey
	secretKey := configs.GlobalConfig.MinioConfig.SecretKey
	fmt.Print("main的viper配置在这个包前生效！url:", url)
	fmt.Printf("accessKey:%s,secretKey:%s", accessKey, secretKey)
	var err error
	dao.MinioClient, err = minio.New(url, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: false, // 如果MinIO服务器使用HTTPS，请设置为true
	})
	if err != nil {
		log.Fatalf("Failed to create MinIO client: %v", err)
	}
}
func setupRedis() {
	dao.RedisClient = redis.NewClient(&redis.Options{
		Addr:        configs.GlobalConfig.RedisConfig.Addr,                                 // Redis 服务器地址
		Password:    configs.GlobalConfig.RedisConfig.Password,                             // Redis 密码
		DB:          configs.GlobalConfig.RedisConfig.Db,                                   // Redis 数据库
		DialTimeout: time.Duration(configs.GlobalConfig.RedisConfig.Timeout) * time.Second, // 连接超时时间
	})

	// 测试 Redis 连接
	_, err := dao.RedisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Connected to Redis")

}

func Setup() {

	once := sync.Once{}
	once.Do(func() {
		configs.LoadConfig()
		setupMinio()
		setupMysql()
		setupMysql()
		setupRedis()
		registerServices()
		registerDAO()
	})

}
