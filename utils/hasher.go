package utils

import (
	"crypto/sha256"
	"io"
	"math/big"
	"os"
)

// Txt2Hash 把.txt文件内容使用sha256算法转换成hash
func Txt2Hash(file *os.File) (uint256 *big.Int, err error) {

	hash := sha256.New()

	// 将文件内容拷贝到哈希对象中
	_, err = io.Copy(hash, file)
	if err != nil {
		return nil, err
	}

	// 获取哈希值的字节切片（32字节 = 256位）
	hashBytes := hash.Sum(nil)

	// 将哈希值字节切片转换为 *big.Int（可以表示任意大小的整数）
	uint256 = new(big.Int).SetBytes(hashBytes)

	return uint256, nil
}
