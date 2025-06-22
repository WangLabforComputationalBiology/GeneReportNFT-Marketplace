package dto

import (
	"GeneReport_platform/internal/dao"
	"GeneReport_platform/internal/models"
)

type GetGeneSharingDetailsByContractAddressReq struct {
	GeneSharingContractAddress string `json:"geneSharing_contract_address"` //集合ID
}

type GetGeneSharingDetailsByContractAddressResp struct {
	GeneSharing       models.GeneSharing
	MetadataOverviews []dao.MetadataOverview `json:"metadata_overviews"` //该集合下的所有GNFT
}

type GetGeneSharingOverviewByCreatorReq struct {
	CreatorAddress string `json:"creator_address"` //创建者地址
}

type GetGeneSharingOverviewByCreatorResp struct {
	GeneSharings []dao.GeneSharingOverview `json:"geneSharings"`
}
