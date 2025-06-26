package dto

import "GeneReport_platform/internal/models"

type GetMetadataResp struct {
	Metadata   models.Metadata    `json:"metadata"`
	Collection models.GeneSharing `json:"geneSharing"`
}

type GetMetadataByOwnerReq struct {
	Owner string `json:"owner" binding:"required"`
}

type GetMetadataByOwnerResp struct {
	Metadata []GetMetadataResp `json:"multi_metadata"`
}

type GetAllMetadataResp struct {
	MultiMetadata []models.Metadata `json:"multi_metadata"`
}
