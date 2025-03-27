package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

// CORS 返回一个 CORS 中间件配置
func CORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},                                       // 设置所有域名
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // 允许的请求方法
		AllowHeaders:     []string{"Content-Type", "Authorization"},           // 允许的请求头
		ExposeHeaders:    []string{"Content-Length"},                          // 客户端可以访问的响应头
		AllowCredentials: true,                                                // 是否允许携带 Cookie
		MaxAge:           12 * time.Hour,                                      // 预检请求的缓存时间
	})
}
