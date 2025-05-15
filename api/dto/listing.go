package dto

import "GeneReport_platform/internal/models"

type GetListingsReq struct {
	CollectionId string `json:"collection_id"`
	Identifier   string `json:"identifier"`
}
type GetListingsResp struct {
	Listings []models.Listing `json:"listings"`
}
