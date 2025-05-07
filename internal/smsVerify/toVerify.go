package smsVerify

import (
	"crypto/rand"
	"fmt"
	unisms "github.com/apistd/uni-go-sdk/sms"
	"math/big"
)

// UniSMSClient sdk中已默认配置了请求端点和hmac模式
var UniSMSClient *unisms.UniSMSClient

func init() {
	// 初始化 smsVerify 客户端
	//UniSMSClient = unisms.NewClient("Nkkp1zkmw2ZvdqjN8EQmiB9MmGSHsX3HPL4LDzuaiqRKt422Y", "Dv6QEBu7xwv83U3dk7Wqt4ZGGSKsg2")//HMAC模式
	UniSMSClient = unisms.NewClient("Nkkp1zkmw2ZvdqjN8EQmiB9MmGSHsX3HPL4LDzuaiqRKt422Y") //简易模式
	fmt.Println("smsVerify client initialized")
}

// GenerateSMSCode 生成 0 到 999999 之间的随机数格式化为 6 位字符串，不足 6 位补 0
func GenerateSMSCode() string {
	n, _ := rand.Int(rand.Reader, big.NewInt(1000000))
	return fmt.Sprintf("%06d", n)
}
