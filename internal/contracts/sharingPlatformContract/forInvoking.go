package sharingPlatformContract

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/v3/abi"
	"github.com/FISCO-BCOS/go-sdk/v3/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/v3/client"
	"github.com/FISCO-BCOS/go-sdk/v3/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/net/context"
	"log"
	"math/big"
	"os"
	"strings"
)

var FiscoConfig *client.Config

var PlatformContractAddressHex string

var AdminPrivateKeyHex string

var AdminAddress string

var MetaDataContractAddress string

var ContractABI abi.ABI

// NewContractIns 获取新的合约实例
func NewContractIns() *SharingPlatformContract {
	// 初始化客户端
	ChainClient, err := client.DialContext(context.Background(), FiscoConfig)

	if err != nil {
		log.Fatalf("连接 FISCO BCOS v3 节点失败: %v", err)
	}

	// 实例化合约
	ContractIns, _ := NewSharingPlatformContract(common.HexToAddress(PlatformContractAddressHex), ChainClient)

	return ContractIns
}

// NewChainClient 新建fisco客户端
func NewChainClient() *client.Client {
	ChainClient, _ := client.DialContext(context.Background(), FiscoConfig)
	return ChainClient
}

func GetContractABI() abi.ABI {
	return ContractABI
}

// NewAdminTransactor 获取新的管理员交易者
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
	AdminAddress = "0x86b4851Bd2F1578Cab7D7f6ff4ffC5589FA41Bd7"
	AdminPrivateKeyHex = "9f5eb599dd2ff51f67724a793a6d702bcc273b3afe3e3bbc0e2870ed11594432"

	// 初始化已部署的sharingPlatform_v3合约地址,转大小写不敏感
	PlatformContractAddressHex = strings.ToLower("0x9F054F2dFE9BeCe24609c3F4a9dCC4624d23ee75")

	//初始化已部署的Metadata合约地址
	MetaDataContractAddress = strings.ToLower("0xf1c1EbE953f500B97f85053A7740bD06A646Df51")

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

	//设置合约ABI
	data, err := os.ReadFile("./internal/contracts/.sol/build/sharingPlatform_v3.abi")
	if err != nil {
		log.Fatalf("读取 ABI 文件失败: %v", err)
	}

	// 解析 ABI JSON
	if err := json.Unmarshal(data, &ContractABI); err != nil {
		log.Fatalf("解析 ABI JSON 失败: %v", err)
	}

}

func DecodeInputData(receipt *types.Receipt) (string, map[string]interface{}, error) {
	//获取input
	input := common.Hex2Bytes(receipt.Input[2:])
	fmt.Println(len(input))
	if len(input) < 4 {
		return "", nil, fmt.Errorf("input 数据太短")
	}

	//解析函数
	methodID := input[:4]
	method, err := ContractABI.MethodByID(methodID)
	if err != nil {
		return "", nil, fmt.Errorf("无法找到匹配的函数: %v", err)
	}

	//解析参数
	params := make(map[string]interface{})
	err = method.Inputs.UnpackIntoMap(params, input[4:])
	if err != nil {
		return "", nil, fmt.Errorf("参数解码失败: %v", err)
	}
	return method.Name, params, nil
}
