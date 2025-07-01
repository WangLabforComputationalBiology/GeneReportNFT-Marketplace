package testCon

import (
	"context"
	"encoding/hex"
	"github.com/FISCO-BCOS/go-sdk/v3/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/v3/client"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
)

var ChainClient *client.Client

var TestContractAddressHex string

var AdminPrivateKeyHex string

func GetContractIns() *Test {

	// 初始化已部署的sharingPlatform_v3合约地址
	TestContractAddressHex = "0xE40263783e6c881bC7C3464CB03594421f662b20"

	// 实例化合约
	ContractIns, _ := NewTest(common.HexToAddress(TestContractAddressHex), ChainClient)

	return ContractIns
}

func NewAdminTransactor() *bind.TransactOpts {
	privateKey, _ := crypto.HexToECDSA(AdminPrivateKeyHex)
	admin := bind.NewKeyedTransactor(privateKey)
	// 设置其他交易参数
	admin.Value = big.NewInt(0)         // 默认无 ETH，适用于非 payable 函数
	admin.GasPrice = big.NewInt(0)      // 私链通常 Gas 价格为 0
	admin.GasLimit = big.NewInt(300000) // 默认 Gas 限制，足够覆盖 setValue
	return admin
}
func init() {
	AdminPrivateKeyHex = "9f5eb599dd2ff51f67724a793a6d702bcc273b3afe3e3bbc0e2870ed11594432"

	// 解码私钥
	privateKeyBytes, err := hex.DecodeString(AdminPrivateKeyHex)
	if err != nil {
		log.Fatalf("解码私钥失败: %v", err)
	}

	// 配置节点信息
	config := &client.Config{
		IsSMCrypto:  false,                          // 非国密
		GroupID:     "group0",                       // 确认群组ID是否为 group0
		PrivateKey:  privateKeyBytes,                // 私钥
		Host:        "10.108.10.51",                 // 节点 IP
		Port:        20200,                          // Channel 端口
		TLSCaFile:   "./assets/fisco_config/ca.crt", // Windows 路径
		TLSKeyFile:  "./assets/fisco_config/sdk.key",
		TLSCertFile: "./assets/fisco_config/sdk.crt",
		DisableSsl:  false, // 生产环境建议 false
	}

	// 初始化客户端
	ChainClient, err = client.DialContext(context.Background(), config)
	if err != nil {
		log.Fatalf("连接 FISCO BCOS v3 节点失败: %v", err)
	}
}
