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

// GNFT基础接口
type iMetadataBase interface {
	GetMetadataOverviewByOwner(owner string) (dto.GetMetadataByOwnerResp, error)
}

func RegisterMetadataService() {
	MetadataServ = &MetadataService{}
}

/*fill your method here*/

func (m *MetadataService) GetMetadataOverviewByOwner(owner string) (dto.GetMetadataByOwnerResp, error) {
	targetMetadatas, err := dao.GetMetadataDao().GetMetadataOverviewByOwner(owner)
	if err != nil {
		return dto.GetMetadataByOwnerResp{}, appErrors.New(503, "获取用户所有GNFT信息失败", err)
	}
	// 映射转换dto
	var toResp dto.GetMetadataByOwnerResp
	mapstructure.Decode(targetMetadatas, &toResp)
	return toResp, nil
}

func (m *MetadataService) GetAllMetadata() (dto.GetAllMetadataResp, error) {
	targetMetadatas, err := dao.GetMetadataDao().GetAllMetadata()
	if err != nil {
		return dto.GetAllMetadataResp{}, appErrors.New(503, "获取用户所有GNFT信息失败", err)
	}

	return dto.GetAllMetadataResp{
		MultiMetadata: targetMetadatas,
	}, nil
}
