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
		AllowHeaders:     []string{"*"},                                       // 允许的请求头
		ExposeHeaders:    []string{"*"},                                       // 客户端可以访问的响应头
		AllowCredentials: false,                                               // 是否允许携带 Cookie
		MaxAge:           12 * time.Hour,                                      // 预检请求的缓存时间
	})
}

//return func(c *gin.Context) {
//	method := c.Request.Method
//	origin := c.Request.Header.Get("Origin")
//	if origin != "" {
//		c.Header("Access-Control-Allow-Origin", "*")  // 可将将 * 替换为指定的域名
//		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
//		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
//		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
//		c.Header("Access-Control-Allow-Credentials", "true")
//	}
//	if method == "OPTIONS" {
//		c.AbortWithStatus(http.StatusNoContent)
//	}
//	c.Next()
//}
