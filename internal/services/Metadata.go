package services

import (
	"GeneReport_platform/api/dto"
	"GeneReport_platform/internal/dao"
	"GeneReport_platform/pkg/appErrors"
	"github.com/mitchellh/mapstructure"
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

	results, err := GetDataImpl(metadata.ProfileID, metadata.Category)
	if err != nil {
		return nil, err
	}

	return results, nil
}
