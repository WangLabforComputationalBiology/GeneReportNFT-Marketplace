package auth

//
//import (
//	"github.com/golang-jwt/jwt/v4"
//	"github.com/google/uuid"
//	"time"
//)
//
//var (
//	// AccessSecret access Token密钥
//	AccessSecret = []byte("your_access_secret")
//	// RefreshSecret refresh Token密钥
//	RefreshSecret = []byte("your_refresh_secret")
//	// AccessExpire access Token过期时间
//	AccessExpire = 15 * time.Minute
//	// RefreshExpire refresh Token过期时间
//	RefreshExpire = 7 * 24 * time.Hour
//)
//
//type CustomClaims struct {
//	UserID uint `json:"user_id"`
//	jwt.RegisteredClaims
//}
//
//// GenerateTokens 生成Access Token和Refresh Token
//func GenerateTokens(userID uint) (accessToken, refreshToken string, err error) {
//	// Access Token
//	accessClaims := CustomClaims{
//		UserID: userID,
//		RegisteredClaims: jwt.RegisteredClaims{
//			ExpiresAt: jwt.NewNumericDate(time.Now().Add(AccessExpire)),
//			ID:        generateJTI(), // 唯一标识用于黑名单
//		},
//	}
//	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString(AccessSecret)
//
//	// Refresh Token（无敏感信息）
//	refreshClaims := jwt.RegisteredClaims{
//		ExpiresAt: jwt.NewNumericDate(time.Now().Add(RefreshExpire)),
//		ID:        generateJTI(),
//	}
//	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(RefreshSecret)
//
//	return
//}
//
//// ParseAccessToken 解析Access Token
//func ParseAccessToken(tokenStr string) (*CustomClaims, error) {
//	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
//		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
//			return nil, errors.New("invalid signing method")
//		}
//		return AccessSecret, nil
//	})
//	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
//		return claims, nil
//	}
//	return nil, err
//}
//
//// 生成唯一JTI（用于黑名单）
//func generateJTI() string {
//	return uuid.New().String()
//}
