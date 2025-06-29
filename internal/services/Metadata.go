package services

import (
	"GeneReport_platform/api/dto"
	"GeneReport_platform/configs"
	"GeneReport_platform/internal/dao"
	"GeneReport_platform/pkg/appErrors"
	"GeneReport_platform/tools/utils"
	"bytes"
	"encoding/json"
	"github.com/mitchellh/mapstructure"
	"net/http"
	"reflect"
)

var (
	MetadataServ *MetadataService
)

type MetadataService struct {
	iMetadataBase
}

// Metadata基础接口
type iMetadataBase interface {
	GetMetadataOverviewByOwner(owner string) (dto.GetMetadataOverviewByOwnerResp, error)
}

func RegisterMetadataService() {
	MetadataServ = &MetadataService{}
}

/*fill your method here*/

func (m *MetadataService) GetMetadataOverviewByOwner(owner string) (dto.GetMetadataOverviewByOwnerResp, error) {
	targetMetadatas, err := dao.GetMetadataDao().GetMetadataOverviewByOwner(owner)
	if err != nil {
		return dto.GetMetadataOverviewByOwnerResp{}, appErrors.New(503, "获取Metadata概述信息失败", err)
	}
	// 映射转换dto
	var toResp dto.GetMetadataOverviewByOwnerResp
	_ = mapstructure.Decode(targetMetadatas, &toResp)
	return toResp, nil
}

func (m *MetadataService) GetAllMetadataOverview() (dto.GetAllMetadataOverviewResp, error) {
	results, err := dao.GetMetadataDao().GetAllMetadataOverview()
	if err != nil {
		return dto.GetAllMetadataOverviewResp{}, appErrors.New(503, "", err)
	}

	return dto.GetAllMetadataOverviewResp{
		MultiMetadata: results,
	}, nil
}

func (m *MetadataService) GetMetadataDetailByDataHash(dataHash string) (map[string]interface{}, error) {
	metadata, err := dao.GetMetadataDao().GetMetadataOverviewByDataHash(dataHash)
	if err != nil {
		return nil, appErrors.New(503, "获取Metadata详细信息失败", err)
	}

	results, err := m.GetDataImpl(metadata.ProfileID, metadata.Category)
	if err != nil {
		return nil, err
	}

	return results, nil
}

// GetDataImpl 获取Metadata的详情数据
func (m *MetadataService) GetDataImpl(profileId, category string) (map[string]interface{}, error) {
	unique := dto.UniqueProfiles{}
	//在数据库查出profileId一样记录
	if profileId != "" {
		configs.DB.Where("profile_id = ?", profileId).Find(&unique)
	}
	if unique.Status == 0 {
		return nil, appErrors.New(http.StatusOK, "数据还没处理完成，请稍等!")
	}

	// 获取结构体类型
	t, ok := dto.GetStructType(category)
	if !ok {
		return nil, appErrors.New(http.StatusBadRequest, "无效的 category", nil)
	}

	// 创建切片用于查询结果
	sliceType := reflect.SliceOf(t)
	results := reflect.New(sliceType).Interface()

	// GORM 查询
	err := configs.DB.Model(reflect.New(t).Interface()).
		Where("profile_id = ?", profileId).
		Order("report_id desc").
		Find(results).Error
	if err != nil {
		return nil, appErrors.New(http.StatusInternalServerError, "查询失败", err)
	}

	// 检查结果是否为空
	resultValue := reflect.ValueOf(results).Elem()
	if resultValue.Len() == 0 {
		return map[string]interface{}{"data": []interface{}{}}, nil // 返回空数组
	}

	// 转为 JSON 并反序列化为 map[string]interface{}
	tempData, err := json.Marshal(resultValue.Interface()) // 序列化切片值而非指针
	if err != nil {
		return nil, appErrors.New(http.StatusInternalServerError, "序列化失败", err)
	}

	var resultArray []map[string]interface{}
	if err := json.Unmarshal(tempData, &resultArray); err != nil {
		return nil, appErrors.New(http.StatusInternalServerError, "反序列化失败", err)
	}

	// 包装为 map[string]interface{}
	return map[string]interface{}{"details": resultArray}, nil
}

//	func GenerateEncryptedZIP(data []models.GenoType, w io.Writer) ([]byte, *io.PipeReader, error) {
//		// 生成 AES 密钥和 IV
//		key := make([]byte, 32) // 256 位
//		if _, err := rand.Read(key); err != nil {
//			return nil, nil, err
//		}
//		iv := make([]byte, aes.BlockSize)
//		if _, err := rand.Read(iv); err != nil {
//			return nil, nil, err
//		}
//		block, err := aes.NewCipher(key)
//		if err != nil {
//			return nil, nil, err
//		}
//		stream := cipher.NewCTR(block, iv)
//
//		// 格式化 private.key 内容
//		keyContent := fmt.Sprintf("-----BEGIN PRIVATE KEY-----\n%s\n-----END PRIVATE KEY-----", key)
//
//		// 创建管道
//		pr, pw := io.Pipe()
//
//		// 异步生成和打包 ZIP
//		go func() {
//			defer pw.Close()
//			zw := zip.NewWriter(pw)
//
//			// 添加加密的 XLSX
//			w, err := zw.Create("data.xlsx")
//			if err != nil {
//				pw.CloseWithError(err)
//				return
//			}
//			ew := cipher.StreamWriter{S: stream, W: w}
//			if err := utils.GenerateXLSX(ew, data); err != nil {
//				pw.CloseWithError(err)
//				return
//			}
//
//			// 添加 private.key
//			keyW, err := zw.Create("private.key")
//			if err != nil {
//				pw.CloseWithError(err)
//				return
//			}
//			if _, err := keyW.Write([]byte(keyContent)); err != nil {
//				pw.CloseWithError(err)
//				return
//			}
//
//			if err := zw.Close(); err != nil {
//				pw.CloseWithError(err)
//			}
//		}()
//
//		return iv, pr, nil
//	}
func (m *MetadataService) GetGenoTypeZip(dataHash, pubKey string) ([]byte, error) {
	metadata, err := dao.GetMetadataDao().GetMetadataOverviewByDataHash(dataHash)
	if err != nil {
		return nil, appErrors.New(503, "获取Metadata详细信息失败", err)
	}

	data, err := dao.GetMetadataDao().GetGenoType(metadata.ProfileID, metadata.Category)
	if err != nil {
		return nil, appErrors.New(503, "获取详细基因型数据失败", err)
	}

	// 创建XLSX缓冲区
	var buf bytes.Buffer
	err = utils.GenerateXLSX(&buf, data)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
