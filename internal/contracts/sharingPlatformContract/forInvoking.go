package sharingPlatformContract

import (
	"encoding/hex"
	"github.com/FISCO-BCOS/go-sdk/v3/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/v3/client"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/net/context"
	"log"
	"math/big"
)

var ChainClient *client.Client

var ContractAddressHex string

var AdminPrivateKeyHex string

var ContractIns *SharingPlatformContract

func GetContractIns() *SharingPlatformContract {
	return ContractIns
}

func GetClient() *client.Client {
	//todo 补充重连机制
	return ChainClient
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
	// 解码私钥
	privateKey, err := hex.DecodeString("145e247e170ba3afd6ae97e88f00dbc976c2345d511b0f6713355d19d8b80b58")
	if err != nil {
		log.Fatalf("解码私钥失败: %v", err)
	}

	// 配置节点信息
	config := &client.Config{
		IsSMCrypto:  false,           // 非国密
		GroupID:     "group0",        // 确认群组ID是否为 group0
		PrivateKey:  privateKey,      // 私钥
		Host:        "127.0.0.1",     // 节点 IP
		Port:        20200,           // Channel 端口
		TLSCaFile:   "./conf/ca.crt", // Windows 路径
		TLSKeyFile:  "./conf/sdk.key",
		TLSCertFile: "./conf/sdk.crt",
		DisableSsl:  false, // 生产环境建议 false
	}

	// 初始化客户端
	ChainClient, err := client.DialContext(context.Background(), config)
	if err != nil {
		log.Fatalf("连接 FISCO BCOS v3 节点失败: %v", err)
	}

	// 初始化已部署的合约地址
	ContractAddressHex = ""

	// 实例化合约
	ContractIns, err = NewSharingPlatformContract(common.HexToAddress(ContractAddressHex), ChainClient)
	if err != nil {
		log.Fatalf("Failed to instantiate contract: %v", err)
	}
}

func NewTransactorSession() SharingPlatformContractTransactorSession {
	return SharingPlatformContractTransactorSession{
		Contract:     &ContractIns.SharingPlatformContractTransactor,
		TransactOpts: *NewAdminTransactor(),
	}
}
