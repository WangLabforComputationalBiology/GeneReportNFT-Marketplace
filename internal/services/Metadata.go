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

func (g *MetadataService) GetMetadataOverviewByOwner(owner string) (dto.GetMetadataByOwnerResp, error) {
	targetGNFTs, err := dao.GetMetadataDao().GetMetadataOverviewByOwner(owner)
	if err != nil {
		return dto.GetMetadataByOwnerResp{}, appErrors.New(503, "获取用户所有GNFT信息失败", err)
	}
	// 映射转换dto
	var toResp dto.GetMetadataByOwnerResp
	mapstructure.Decode(targetGNFTs, &toResp)
	return toResp, nil
}
