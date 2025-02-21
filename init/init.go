package init

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

var GlobalConfig *Config

type AppConfig struct {
	Name string `mapstructure:"name"`
	Port string `mapstructure:"port"`
}
type MysqlConfig struct {
	Dsn          string `mapstructure:"dsn"`
	MaxIdleConns int    `mapstructure:"maxidleconns"`
	MaxOpenConns int    `mapstructure:"maxopenconns"`
}
type RedisConfig struct {
}
type Config struct {
	AppConfig   AppConfig   `mapstructure:"app"`
	MysqlConfig MysqlConfig `mapstructure:"mysql"`
}

func initConfig() {
	// 初始化配置文件
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local" // 默认环境是本地开发环境
	}
	viper.SetConfigName(fmt.Sprintf("config_%s", env))
	viper.AddConfigPath("./init")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("无法读取配置文件: %v", err)
	}
	//初始化GlobalConfig
	GlobalConfig = &Config{}
	if err := viper.Unmarshal(GlobalConfig); err != nil {
		log.Fatalf("无法解析配置文件: %v", err)
	}
}
func initDB() {
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
func initServices() {

}
func initDAO() {

}
func Init() {
	initConfig()
	initDB()
	initServices()
	initDAO()
}
