package services

import (
	"GeneReport_platform/api/dto"
	"GeneReport_platform/internal/dao"
	"GeneReport_platform/pkg/appErrors"
	"github.com/mitchellh/mapstructure"
)

var (
	GNFTServ *GNFTService
)

type GNFTService struct {
	iGNFTBase
}

// GNFT基础接口
type iGNFTBase interface {
	GetGNFTInfosByOwner(owner string) (dto.GetGNFTInfosByOwnerResp, error)
}

func RegisterGNFTService() {
	GNFTServ = &GNFTService{}
}

/*fill your method here*/

func (g *GNFTService) GetGNFTInfosByOwner(owner string) (dto.GetGNFTInfosByOwnerResp, error) {
	targetGNFTs, err := dao.GetGNFTDao().GetGNFTInfosByOwner(owner)
	if err != nil {
		return dto.GetGNFTInfosByOwnerResp{}, appErrors.New(503, "获取用户所有GNFT信息失败", err)
	}
	// 映射转换dto
	var toResp dto.GetGNFTInfosByOwnerResp
	mapstructure.Decode(targetGNFTs, &toResp)
	return toResp, nil
}
