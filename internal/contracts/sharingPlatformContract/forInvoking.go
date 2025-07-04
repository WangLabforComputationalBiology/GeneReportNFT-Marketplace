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

var FiscoConfig *client.Config

var PlatformContractAddressHex string

var AdminPrivateKeyHex string

var MetaDataContractAddress string

func GetContractIns() *SharingPlatformContract {
	// 初始化客户端
	ChainClient, err := client.DialContext(context.Background(), FiscoConfig)

	if err != nil {
		log.Fatalf("连接 FISCO BCOS v3 节点失败: %v", err)
	}

	// 实例化合约
	ContractIns, _ := NewSharingPlatformContract(common.HexToAddress(PlatformContractAddressHex), ChainClient)

	return ContractIns
}

func NewAdminTransactor() *bind.TransactOpts {
	privateKey, _ := crypto.HexToECDSA(AdminPrivateKeyHex)
	admin := bind.NewKeyedTransactor(privateKey)
	// 设置其他交易参数
	admin.Value = big.NewInt(0)             // 默认无 ETH，适用于非 payable 函数
	admin.GasPrice = big.NewInt(0)          // 私链通常 Gas 价格为 0
	admin.GasLimit = big.NewInt(3000000000) // 默认 Gas 限制，足够覆盖 setValue
	return admin
}

func init() {
	AdminPrivateKeyHex = "9f5eb599dd2ff51f67724a793a6d702bcc273b3afe3e3bbc0e2870ed11594432"

	//初始化已部署的Metadata合约地址
	MetaDataContractAddress = "0xE4386cf966aC0Ca0757020d6F0C07AB3dCE13a2D"

	// 初始化已部署的sharingPlatform_v3合约地址
	PlatformContractAddressHex = "0x9FF72c7E4a7c0d3126632C05e07028e1aa3004d3"

	// 解码私钥
	privateKeyBytes, err := hex.DecodeString(AdminPrivateKeyHex)
	if err != nil {
		log.Fatalf("解码私钥失败: %v", err)
	}

	// 配置节点信息
	FiscoConfig = &client.Config{
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

}
