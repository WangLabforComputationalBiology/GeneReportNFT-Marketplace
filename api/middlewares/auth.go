package middlewares

import (
	"GeneReport_platform/pkg/auth"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"time"
)

// AuthMiddleware
//
//	@securityDefinitions.apikey	JwtAuth
//	@in							header
//	@name						Authorization
//	@description				JWT token in format "Bearer {token}"
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取请求中的 Authorization header
		tokenString := ctx.GetHeader("Authorization")
		parts := strings.Split(tokenString, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			log.Printf("接收到jwt为%v", tokenString)
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "未授权或非法的token格式"})
			ctx.Abort()
			return
		}

		// 解析 JWT
		claim, err := auth.ParseToken(parts[0])
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "未授权或非法的token格式"})
			ctx.Abort()
			return
		}

		//检查token是否过期
		if claim.ExpiresAt.Before(time.Now()) {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token已过期"})
			ctx.Abort()
			return
		}
		//检查黑名单
		isInBlackList, err := auth.CheckBlacklist(claim.ID)
		if isInBlackList && err == nil {
			//在黑名单中
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token已失效"})
			ctx.Abort()
			return
		} else if err != nil {
			ctx.JSON(http.StatusServiceUnavailable, gin.H{"error": "redis服务异常"})
			ctx.Abort()
			return
		}

		// 通过校验，将user_address、JTI 存入context
		ctx.Set("user_address", claim.UserAddress)
		ctx.Set("JTI", claim.ID)

		ctx.Next()
	}
}
