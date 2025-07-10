package utils

import (
	"GeneReport_platform/pkg/appErrors"
	"archive/zip"
	"crypto/rand"
	"encoding/csv"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"io"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// GenerateShortCode 生成 8 位短链接
func GenerateShortCode() string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 8)
	if _, err := rand.Read(b); err != nil {
		panic(err) // 生产环境中应记录日志并返回错误
	}
	for i := 0; i < 8; i++ {
		b[i] = chars[b[i]%62] // 映射到字符集
	}
	return string(b)
}

// GenerateCsvZip 向zipWriter写入Csv文件
func GenerateCsvZip(writer io.Writer, category string, data any) error {
	// 使用反射检查是否为切片
	v := reflect.ValueOf(data)
	if v.Kind() != reflect.Slice {
		return fmt.Errorf("data must be a slice")
	}
	if v.Len() == 0 {
		return nil
	}

	// 检查切片元素是否为结构体
	t := v.Index(0).Type()
	if t.Kind() != reflect.Struct {
		return fmt.Errorf("slice elements must be structs")
	}

	// 创建 ZIP 写入器
	zw := zip.NewWriter(writer)
	defer zw.Close()

	// 创建 ZIP 条目
	zipEntry, err := zw.Create(category + ".csv")
	if err != nil {
		return err
	}

	// 添加 UTF-8 BOM（兼容 Excel）
	_, err = zipEntry.Write([]byte{0xEF, 0xBB, 0xBF})
	if err != nil {
		return appErrors.New(http.StatusInternalServerError, "服务繁忙，请稍后再试", errors.New("设置UTF-8 BOM失败"))
	}

	// 创建 CSV 写入器
	csvWriter := csv.NewWriter(zipEntry)
	defer csvWriter.Flush()

	// 写入表头
	headers := make([]string, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		headers[i] = t.Field(i).Name
	}
	if err := csvWriter.Write(headers); err != nil {
		return err
	}

	// 分块写入数据
	for i := 0; i < v.Len(); i++ {
		elem := v.Index(i)
		row := make([]string, elem.NumField())
		for j := 0; j < elem.NumField(); j++ {
			// 将字段值转换为字符串
			row[j] = fmt.Sprint(elem.Field(j).Interface())
		}
		if err := csvWriter.Write(row); err != nil {
			return err
		}

		// 每 10 行 flush 一次，确保分块传输
		if (i+1)%10 == 0 {
			csvWriter.Flush()
			if err := csvWriter.Error(); err != nil {
				return err
			}
		}
	}

	// 确保最后一批数据写入
	csvWriter.Flush()
	if err := csvWriter.Error(); err != nil {
		return err
	}

	// 关闭 ZIP 写入器
	return zw.Close()
}

// 生成对称密钥（32 字节）
func generateSymmetricKey() ([]byte, error) {
	key := make([]byte, 32) // AES-256 需要 32 字节密钥
	_, err := rand.Read(key)
	if err != nil {
		return nil, fmt.Errorf("生成对称密钥失败: %v", err)
	}
	return key, nil
}

// 使用 ECIES 加密对称密钥
func encryptWithECIES(pubKeyHex string, data []byte) ([]byte, error) {
	// 解析公钥
	pubKeyBytes, err := hex.DecodeString(strings.TrimPrefix(pubKeyHex, "0x"))
	if err != nil {
		return nil, fmt.Errorf("解码公钥失败: %v", err)
	}
	pubKey, err := crypto.UnmarshalPubkey(pubKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("解析公钥失败: %v", err)
	}
	eciesPubKey := ecies.ImportECDSAPublic(pubKey)

	// 执行 ECIES 加密
	encrypted, err := ecies.Encrypt(rand.Reader, eciesPubKey, data, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("ECIES 加密失败: %v", err)
	}
	return encrypted, nil
}

// 零宽度字符编码（每个字符 4 位二进制）
func encodeToZeroWidth(text string, secret string) string {
	const ZWSP = "\u200B"   // 表示 0
	const ZWNBSP = "\u200C" // 表示 1
	binary := ""
	for _, c := range secret {
		// 将字符（0-9）转换为 4 位二进制
		n, _ := strconv.Atoi(string(c))  // 假设 secret 只包含数字
		binary += fmt.Sprintf("%04b", n) // 例如 '1' -> "0001"
	}
	encoded := ""
	for _, bit := range binary {
		if bit == '0' {
			encoded += ZWSP
		} else {
			encoded += ZWNBSP
		}
	}
	return text + encoded
}

// 保存为 .key 文件
func saveToKeyFile(filename string, data string) error {
	err := os.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		return fmt.Errorf("保存 .key 文件失败: %v", err)
	}
	return nil
}

// 解析零宽度字符（按 4 位解码）
func decodeZeroWidth(text string) (string, string) {
	const ZWSP = "\u200B"
	const ZWNBSP = "\u200C"
	var binary strings.Builder
	for _, r := range text {
		if r == []rune(ZWSP)[0] {
			binary.WriteString("0")
		} else if r == []rune(ZWNBSP)[0] {
			binary.WriteString("1")
		}
	}
	binaryStr := binary.String()
	secret := ""
	for i := 0; i < len(binaryStr); i += 4 {
		if i+4 <= len(binaryStr) {
			bits := binaryStr[i : i+4]
			n, _ := strconv.ParseInt(bits, 2, 64)
			secret += fmt.Sprintf("%d", n) // 转换为数字字符
		}
	}
	originalText := strings.ReplaceAll(strings.ReplaceAll(text, ZWSP, ""), ZWNBSP, "")
	return originalText, secret
}

// EncryptDistributionV1 加密分发实现V1
// 原始加密密钥为originKeyHex
// 展示给用户的是KeyStep2
func EncryptDistributionV1(originKeyHex, pubKeyHex string, expiry int64) (string, error) {
	KeyStep1, err := encryptWithECIES(pubKeyHex, common.Hex2Bytes(originKeyHex))
	if err != nil {
		return "", err
	}
	KeyStep2 := encodeToZeroWidth(common.Bytes2Hex(KeyStep1), strconv.FormatInt(expiry, 10))
	return KeyStep2, nil
}

//// EncryptDistributionV2 加密分发实现V2
//// 原始加密密钥为KeyStep2、
//// 展示给用户是KeyStep1
//func EncryptDistributionV2(originKeyHex, pubKeyHex string, expiry int64) (string, error) {
//	encryptedExpiry, err := encryptWithECIES(pubKeyHex, common.Hex2Bytes(strconv.FormatInt(expiry, 10)))
//	if err != nil {
//		return "", err
//	}
//	//把encryptedExpiry隐写到originKeyHex中
//	KeyStep1, err := LBSWrite(originKeyHex, encryptedExpiry)
//	//KeyStep1和时间戳派生出KeyStep2，作为对文件加密的原始加密密钥
//	KeyStep2 := KeyStep1 + strconv.FormatInt(expiry, 10)
//}
