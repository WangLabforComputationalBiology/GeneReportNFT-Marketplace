package dto

import (
	"GeneReport_platform/internal/models"
)

type GetCollectionInfoByIDReq struct {
	CollectionID string `json:"collection_id"` //集合ID
}

type GetCollectionInfoResp struct {
	Collection models.Collection
	GNFTs      []models.GNFT `json:"GNFTs"` //该集合下的所有GNFT
}

type GetCollectionsInfosByCreatorReq struct {
	Creator string `json:"creator"` //创建者地址
}

type GetCollectionInfosByCreatorResp struct {
	Collections []GetCollectionInfoResp `json:"collections"`
}
