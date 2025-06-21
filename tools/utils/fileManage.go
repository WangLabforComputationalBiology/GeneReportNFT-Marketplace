package utils

import (
	"crypto/rand"
	"fmt"
	"github.com/xuri/excelize/v2"
	"io"
	"reflect"
)

// GenerateXLSX 根据切片生成xlsx文件
func GenerateXLSX(w io.Writer, data any) error {
	f := excelize.NewFile()
	defer f.Close()

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

	// 设置表头
	headers := make([]interface{}, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		headers[i] = t.Field(i).Name
	}
	f.SetSheetRow("Sheet1", "A1", &headers)

	// 写入数据行
	for i := 0; i < v.Len(); i++ {
		elem := v.Index(i)
		row := make([]interface{}, elem.NumField())
		for j := 0; j < elem.NumField(); j++ {
			row[j] = elem.Field(j).Interface()
		}
		f.SetSheetRow("Sheet1", fmt.Sprintf("A%d", i+2), &row)
	}
	_, err := f.WriteTo(w)
	return err
}

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
