package configs

import (
	"github.com/spf13/viper"
	"log"
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
type MinioConfig struct {
	Host      string `mapstructure:"host"`
	Port      int    `mapstructure:"port"`
	AccessKey string `mapstructure:"access_key"`
	SecretKey string `mapstructure:"secret_key"`
}
type EtherScanConfig struct {
	ApiKey   string `mapstructure:"apikey"`
	Endpoint string `mapstructure:"endpoint"`
	Proxy    bool   `mapstructure:"proxy"`
	ProxyUrl string `mapstructure:"proxy_url"`
}
type Config struct {
	AppConfig       AppConfig       `mapstructure:"app"`
	MysqlConfig     MysqlConfig     `mapstructure:"mysql"`
	RedisConfig     RedisConfig     `mapstructure:"redis"`
	MinioConfig     MinioConfig     `mapstructure:"minio"`
	CtxConfig       CtxConfig       `mapstructure:"ctx"`
	EtherScanConfig EtherScanConfig `mapstructure:"ether_scan"`
}

func LoadConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./configs") //在和go.mod同级的地方开始！
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("无法读取配置文件: %v", err)
	}
	//GlobalConfig 正确初始化
	GlobalConfig = &Config{}
	if err := viper.Unmarshal(GlobalConfig); err != nil {
		log.Fatalf("无法解析配置文件: %v", err)
	}
	log.Printf("已读取到配置文件: %v", GlobalConfig)
}

//func initEtherScanConfig() {
//	dao.ApiKey = GlobalConfig.EtherScanConfig.ApiKey
//	dao.EndPoint = GlobalConfig.EtherScanConfig.Endpoint
//	dao.IsProxy = GlobalConfig.EtherScanConfig.Proxy
//	dao.ProxyUrl = GlobalConfig.EtherScanConfig.ProxyUrl
//}
