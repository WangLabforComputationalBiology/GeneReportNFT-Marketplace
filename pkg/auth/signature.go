package auth

import (
	"GeneReport_platform/configs"
	"GeneReport_platform/pkg/appContext"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"log"
)

// GenerateNonce 生成验签码
func GenerateNonce() string {
	return uuid.New().String()
}

// structureMessage 构建签名信息
func structureMessage(address, nonce string) string {
	template := `Welcome to GeneReport_platform!

Click to sign in and accept the OpenSeaTerms of Service and Privacy Policy.

This request will not trigger a blockchain transaction or cost any gas fees.

Wallet address:
%v

Nonce:
%v`

	return fmt.Sprintf(template, address, nonce)

}

// VerifySignature 执行验签
func VerifySignature(address, nonce, signatureStr string) (bool, error) {
	log.Printf("VerifySignature: address：%v ;;nonce: %v ;;signatureStr: %v", address, nonce, signatureStr)
	//构造签名信息（带默认前缀）
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(structureMessage(address, nonce)), structureMessage(address, nonce))
	msgHash := crypto.Keccak256Hash([]byte(msg))
	signature := common.Hex2Bytes(signatureStr[2:]) // 十六进制字符串转字节
	log.Printf("签名字节数为：%v", len(signature))
	if len(signature) != 65 {
		return false, errors.New("签名长度无效")
	}
	// 调整 v 值
	signature[64] -= 27
	if signature[64] > 1 {
		return false, errors.New("无效的恢复参数 v")
	}

	// 恢复公钥
	pubKey, err := crypto.SigToPub(msgHash.Bytes(), signature)
	if err != nil {
		return false, errors.New("无法恢复公钥")
	}

	// 验证地址
	derivedAddress := crypto.PubkeyToAddress(*pubKey).Hex()
	expectedAddress := common.HexToAddress(address).Hex()
	if derivedAddress != expectedAddress {
		return false, errors.New("地址不匹配")
	}
	// 验证签名
	valid := crypto.VerifySignature(crypto.FromECDSAPub(pubKey), msgHash.Bytes(), signature[:64])
	if !valid {
		return false, errors.New("签名无效")
	}

	// 验证通过后删除 nonce
	if err = configs.RedisClient.Del(appContext.NewTimeoutContext(), address).Err(); err != nil {
		return false, errors.New("redis服务异常")
	}

	return true, nil
}
