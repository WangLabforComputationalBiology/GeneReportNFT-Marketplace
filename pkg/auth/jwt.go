package auth

import (
	"GeneReport_platform/configs"
	"GeneReport_platform/pkg/appContext"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"time"
)

type CustomClaims struct {
	UserAddress string `json:"user_address"`
	PubKey      string `json:"pub_key"`
	jwt.RegisteredClaims
}

// SecretKey 加密密钥
var SecretKey = []byte("SZTU")

// TokenExpireDuration 默认过期时间
var TokenExpireDuration = time.Minute * 3000

// GenerateJTI 生成JTI
func GenerateJTI() string {
	// 调用生成随机ID并添加时间戳的函数
	return uuid.New().String()
}

// GenerateToken 生成JWT
func GenerateToken(userAddress, pubKeyHex string) (string, error) {

	// 实例化声明
	c := CustomClaims{
		userAddress, // 用户地址
		pubKeyHex,   //用户钱包公钥
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)), // 过期时间
			Issuer:    "GeneReport_platform",                                   // 签发人
			ID:        GenerateJTI(),                                           //jti
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	fmt.Println(token.SignedString(SecretKey))
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(SecretKey)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*CustomClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func CheckBlacklist(jti string) (bool, error) {
	if err := configs.RedisClient.Get(appContext.NewTimeoutContext(), "blacklist:"+jti).Err(); err != nil {
		// 不在黑名单中
		if errors.Is(err, redis.Nil) {
			return false, nil
		}
		//redis服务异常
		return false, err
	}
	// 在黑名单中
	return true, nil
}
