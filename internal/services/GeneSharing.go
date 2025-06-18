package services

import (
	"GeneReport_platform/api/dto"
	"GeneReport_platform/internal/dao"
	"GeneReport_platform/pkg/appErrors"
	"github.com/mitchellh/mapstructure"
)

var (
	GeneSharingServ *GeneSharingService
)

type GeneSharingService struct {
	iGeneSharingBase
}

// collection基础接口
type iGeneSharingBase interface {
}

func RegisterGeneSharingService() {
	GeneSharingServ = &GeneSharingService{}
}

// GetGeneSharingDetailsByContractAddress 将[]dao.CollectionWithGNFT 整理为 dto.GetCollectionInfoResp
func (c *GeneSharingService) GetGeneSharingDetailsByContractAddress(geneSharingID string) (dto.GetGeneSharingDetailsByContractAddressResp, error) {
	var toResp dto.GetGeneSharingDetailsByContractAddressResp

	results, err := dao.GetGeneSharingDao().GetGeneSharingDetailsByAddress(geneSharingID)
	if err != nil {
		return dto.GetGeneSharingDetailsByContractAddressResp{}, appErrors.New(503, "获取集合信息失败", err)
	}

	_ = mapstructure.Decode(results, toResp)
	return toResp, nil
}

func (c *GeneSharingService) GetGeneSharingOverviewByCreator(creatorAddress string) (dto.GetGeneSharingOverviewByCreatorResp, error) {

	results, err := dao.GetGeneSharingDao().GetGeneSharingOverviewByCreator(creatorAddress)
	if err != nil {
		return dto.GetGeneSharingOverviewByCreatorResp{}, appErrors.New(503, "failed", err)
	}

	return dto.GetGeneSharingOverviewByCreatorResp{GeneSharings: results}, nil
}
