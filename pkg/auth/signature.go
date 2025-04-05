package auth

import (
	"GeneReport_platform/configs"
	"GeneReport_platform/pkg/appContext"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
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
	message := structureMessage(address, nonce)
	hash := crypto.Keccak256Hash([]byte(message))
	signature := common.Hex2Bytes(signatureStr) // 十六进制字符串转字节
	if len(signature) != 65 {
		return false, errors.New("签名长度无效")
	}
	// 调整 v 值
	signature[64] -= 27
	if signature[64] > 1 {
		return false, errors.New("无效的恢复参数 v")
	}

	// 恢复公钥
	pubKey, err := crypto.SigToPub(hash.Bytes(), signature)
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
	valid := crypto.VerifySignature(crypto.FromECDSAPub(pubKey), hash.Bytes(), signature[:64])
	if !valid {
		return false, errors.New("签名无效")
	}

	// 验证通过后删除 nonce
	if err = configs.RedisClient.Del(appContext.NewTimeoutContext(), address).Err(); err != nil {
		return false, errors.New("redis服务异常")
	}

	return true, nil
}
