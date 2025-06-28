package services

import (
	"GeneReport_platform/api/dto"
	"GeneReport_platform/configs"
	"GeneReport_platform/internal/dao"
	"GeneReport_platform/pkg/appErrors"
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
	mapstructure.Decode(targetMetadatas, &toResp)
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

// GetDataImpl 获取GeneType数据
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
	results := reflect.MakeSlice(sliceType, 0, 0)

	// GORM 查询
	err := configs.DB.Where("profile_id = ?", profileId).Find(&results).Error
	if err != nil {
		return nil, appErrors.New(http.StatusInternalServerError, "查询失败", err)
	}

	// 转为 map[string]interface{}
	data, err := json.Marshal(results.Interface())
	if err != nil {
		return nil, appErrors.New(http.StatusInternalServerError, "序列化失败", err)
	}
	var resultMap map[string]interface{}
	if err := json.Unmarshal(data, &resultMap); err != nil {
		return nil, appErrors.New(http.StatusInternalServerError, "反序列化失败", err)
	}

	return resultMap, nil
}
