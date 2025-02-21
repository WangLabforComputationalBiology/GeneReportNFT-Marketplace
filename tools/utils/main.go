package utils

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"strings"
)

// KMS_KEY_ID 配置 KMS 密钥
var KMS_KEY_ID = "your-kms-key-id"                                // 请替换为你的 KMS 密钥 ID
var ENCRYPTED_PRIVATE_KEY = "your-encrypted-private-key-from-kms" // 请替换为加密的私钥

// 获取 Relayer 私钥并解密
func getRelayerPrivateKeyFromKMS() (string, error) {
	// 使用 AWS SDK 创建一个 session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})
	if err != nil {
		return "", fmt.Errorf("unable to create AWS session: %v", err)
	}

	// 创建 KMS 客户端
	kmsClient := kms.New(sess)

	// 获取加密的密钥，假设密钥已通过 KMS 加密并存储
	encryptedKey := ENCRYPTED_PRIVATE_KEY

	// 解密密钥
	result, err := kmsClient.Decrypt(&kms.DecryptInput{
		CiphertextBlob: []byte(encryptedKey),
	})
	if err != nil {
		return "", fmt.Errorf("unable to decrypt private key: %v", err)
	}

	// 返回解密后的密钥
	return string(result.Plaintext), nil
}

func main() {
	// 获取 Relayer 私钥
	privateKeyStr, err := getRelayerPrivateKeyFromKMS()
	if err != nil {
		log.Fatalf("Error retrieving private key: %v", err)
	}

	// 打印私钥（仅为调试，实际应用中请勿暴露）
	fmt.Println("Relayer private key retrieved:", privateKeyStr)

	// 设置 Web3 客户端（例如 Infura 或者本地以太坊节点）
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/YOUR_INFURA_PROJECT_ID") // 请替换为你自己的 Infura URL 或者节点 URL
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// 示例：创建一个 MetaTransaction 签名并发送
	sendTransaction(client, privateKeyStr)
}

// MetaTransaction 签名并发送交易
func sendTransaction(client *ethclient.Client, privateKeyStr string) {
	// 设置交易参数
	toAddress := common.HexToAddress("0xYourContractAddress") // 请替换为目标合约地址
	amount := big.NewInt(1000000000000000000)                 // 发送的数量，单位为 wei

	// 获取合约 ABI 和方法
	contractABI, _ := abi.JSON(strings.NewReader(`[{"constant":false,"inputs":[{"name":"amount","type":"uint256"}],"name":"transfer","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]`))
	method := "transfer"
	params := []interface{}{amount}

	// 获取方法签名
	data, err := contractABI.Pack(method, params...)
	if err != nil {
		log.Fatalf("Error packing contract method: %v", err)
	}

	// 获取交易计数（nonce）
	nonce, err := client.PendingNonceAt(context.Background(), common.HexToAddress("0xYourRelayerAddress")) // 请替换为你的 Relayer 地址
	if err != nil {
		log.Fatalf("Error getting nonce: %v", err)
	}

	// 设置交易 gas 参数
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Error getting suggested gas price: %v", err)
	}
	gasLimit := uint64(21000) // 假设这是一个简单的转账交易

	// 创建交易对象
	tx := types.NewTransaction(
		nonce,
		toAddress,
		amount,
		gasLimit,
		gasPrice,
		data,
	)

	// 将私钥字符串转换为 *ecdsa.PrivateKey
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		log.Fatalf("Error converting private key string to ecdsa: %v", err)
	}

	// 使用私钥签名交易
	chainID, _ := client.NetworkID(context.Background())
	signer := types.NewEIP155Signer(chainID)
	signedTx, err := types.SignTx(tx, signer, privateKey)
	if err != nil {
		log.Fatalf("Error signing transaction: %v", err)
	}

	// 发送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatalf("Error sending transaction: %v", err)
	}

	// 打印交易哈希
	fmt.Println("Transaction hash:", signedTx.Hash().Hex())
}
