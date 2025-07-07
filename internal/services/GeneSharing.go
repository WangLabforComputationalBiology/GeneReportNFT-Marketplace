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

// GetGeneSharingDetailsByContractAddress 根据合约地址获取geneSharing合集详情
func (g *GeneSharingService) GetGeneSharingDetailsByContractAddress(geneSharingID string) (dto.GetGeneSharingDetailsByContractAddressResp, error) {
	var toResp dto.GetGeneSharingDetailsByContractAddressResp

	results, err := dao.GetGeneSharingDao().GetGeneSharingDetailsByAddress(geneSharingID)
	if err != nil {
		return dto.GetGeneSharingDetailsByContractAddressResp{}, appErrors.New(503, "获取geneSharing合集信息失败", err)
	}

	_ = mapstructure.Decode(results, toResp)
	return toResp, nil
}

func (g *GeneSharingService) GetGeneSharingOverviewByCreator(creatorAddress string) (dto.GetGeneSharingOverviewByCreatorResp, error) {

	results, err := dao.GetGeneSharingDao().GetGeneSharingOverviewByCreator(creatorAddress)
	if err != nil {
		return dto.GetGeneSharingOverviewByCreatorResp{}, appErrors.New(503, "failed", err)
	}

	return dto.GetGeneSharingOverviewByCreatorResp{GeneSharings: results}, nil
}

func (g *GeneSharingService) GetAllGeneSharingOverview(page int) (dto.GetAllGeneSharingOverviewResp, error) {
	results, pageNum, err := dao.GetGeneSharingDao().GetAllGeneSharingOverview(page)
	if err != nil {
		return dto.GetAllGeneSharingOverviewResp{}, appErrors.New(503, "获取GeneSharing合集概览失败", err)
	}
	return dto.GetAllGeneSharingOverviewResp{
		MultiGeneSharing: results,
		PageNum:          pageNum,
	}, nil
}
