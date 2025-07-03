package utils

import (
	"archive/zip"
	"crypto/rand"
	"encoding/csv"
	"fmt"
	"io"
	"reflect"
)

//// GenerateXLSX 根据切片生成xlsx文件
//func GenerateXLSX(w io.Writer, data any) error {
//	f := excelize.NewFile()
//	defer f.Close()
//
//	// 使用反射检查是否为切片
//	v := reflect.ValueOf(data)
//	if v.Kind() != reflect.Slice {
//		return fmt.Errorf("data must be a slice")
//	}
//	if v.Len() == 0 {
//		return nil
//	}
//
//	// 检查切片元素是否为结构体
//	t := v.Index(0).Type()
//	if t.Kind() != reflect.Struct {
//		return fmt.Errorf("slice elements must be structs")
//	}
//
//	// 设置表头
//	headers := make([]interface{}, t.NumField())
//	for i := 0; i < t.NumField(); i++ {
//		headers[i] = t.Field(i).Name
//	}
//	f.SetSheetRow("Sheet1", "A1", &headers)
//
//	// 写入数据行
//	for i := 0; i < v.Len(); i++ {
//		elem := v.Index(i)
//		row := make([]interface{}, elem.NumField())
//		for j := 0; j < elem.NumField(); j++ {
//			row[j] = elem.Field(j).Interface()
//		}
//		f.SetSheetRow("Sheet1", fmt.Sprintf("A%d", i+2), &row)
//	}
//	_, err := f.WriteTo(w)
//	return err
//}

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
func GenerateCsvZip(writer io.Writer, data any) error {
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
	zipEntry, err := zw.Create("example.csv")
	if err != nil {
		return err
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
