package dto

import (
	"GeneReport_platform/internal/models"
)

type GetCollectionWithGNFTsByIDReq struct {
	CollectionID string `json:"collection_id"` //集合ID
}

type GetCollectionWithGNFTsByIDResp struct {
	Collection models.Collection
	GNFTs      []models.GNFT `json:"GNFTs"` //该集合下的所有GNFT
}

type GetCollectionsWithGNFTsByCreatorReq struct {
	Creator string `json:"creator"` //创建者地址
}

type GetCollectionsWithGNFTsByCreatorResp struct {
	Collections []GetCollectionWithGNFTsByIDResp `json:"collections"`
}
