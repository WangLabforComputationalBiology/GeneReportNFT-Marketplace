package init

import (
	"GeneReport_platform/internal/dao"
	"GeneReport_platform/internal/dao/global"
	"GeneReport_platform/internal/services"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var GlobalConfig *Config

type AppConfig struct {
	Name string `mapstructure:"name"`
	Addr string `mapstructure:"addr"`
}
type MysqlConfig struct {
	Dsn          string `mapstructure:"dsn"`
	MaxIdleConns int    `mapstructure:"maxidleconns"`
	MaxOpenConns int    `mapstructure:"maxopenconns"`
}
type RedisConfig struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	Timeout  int    `mapstructure:"timeout"`
	Db       int    `mapstructure:"db"`
}
type CtxConfig struct {
	Timeout int `mapstructure:"timeout"`
}
type Etherscan struct {
	ApiKey   string `mapstructure:apikey`
	Endpoint string `mapstructure:endpoint`
	Proxy    bool   `mapstructure:proxy`
	ProxyUrl string `mapstructure:proxyurl`
}
type Config struct {
	AppConfig   AppConfig   `mapstructure:"app"`
	MysqlConfig MysqlConfig `mapstructure:"mysql"`
	RedisConfig RedisConfig `mapstructure:"redis"`
	CtxConfig   CtxConfig   `mapstructure:"ctx"`
	EthConfig   Etherscan   `mapstructure:"etherscan"`
}

func initConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config") //在和go.mod同级的地方开始！
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("无法读取配置文件: %v", err)
	}
	//GlobalConfig 正确初始化
	GlobalConfig = &Config{}
	if err := viper.Unmarshal(GlobalConfig); err != nil {
		log.Fatalf("无法解析配置文件: %v", err)
	}

}
func initMysql() {
	//初始化数据库连接
	dsn := GlobalConfig.MysqlConfig.Dsn
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get sql db: %v", err)
	}
	sqlDB.SetMaxIdleConns(GlobalConfig.MysqlConfig.MaxIdleConns)
	sqlDB.SetMaxOpenConns(GlobalConfig.MysqlConfig.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour)
	global.DB = db
}

func initMinio() {
	//todo 没建结构体，直接读！
	url := viper.GetString("minio.host") + ":" + viper.GetString("minio.port")
	accessKey := viper.GetString("minio.access_key")
	secretKey := viper.GetString("minio.secret_key")
	//fmt.Println("main的viper配置在这个包前生效！url:", url)
	fmt.Print("main的viper配置在这个包前生效！url:", url)
	fmt.Printf("accessKey:%s,secretKey:%s", accessKey, secretKey)
	var err error
	//不用：=
	global.MinioClient, err = minio.New(url, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: false, // 如果MinIO服务器使用HTTPS，请设置为true
	})
	if err != nil {
		log.Fatalf("Failed to create MinIO client: %v", err)
	}
}
func initRedis() {
	global.RedisClient = redis.NewClient(&redis.Options{
		Addr:        GlobalConfig.RedisConfig.Addr,                                 // Redis 服务器地址
		Password:    GlobalConfig.RedisConfig.Password,                             // Redis 密码
		DB:          GlobalConfig.RedisConfig.Db,                                   // Redis 数据库
		DialTimeout: time.Duration(GlobalConfig.RedisConfig.Timeout) * time.Second, // 连接超时时间
	})

	// 测试 Redis 连接
	_, err := global.RedisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Connected to Redis")

}

func initContext() {
	global.Ctx, global.CtxCancel = context.WithTimeout(context.Background(), time.Duration(GlobalConfig.CtxConfig.Timeout)*time.Second)
}

func initEtherscan() {
	global.ApiKey = GlobalConfig.EthConfig.ApiKey
	global.EndPoint = GlobalConfig.EthConfig.Endpoint
	global.IsProxy = GlobalConfig.EthConfig.Proxy
	global.ProxyUrl = GlobalConfig.EthConfig.ProxyUrl
}
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

func Init() {
	initConfig()
	initMysql()
	initMinio()
	initRedis()
	initContext()
	registerServices()
	registerDAO()
	initEtherscan()

}
