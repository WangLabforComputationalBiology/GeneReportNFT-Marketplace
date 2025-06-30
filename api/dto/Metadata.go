package dto

import (
	"GeneReport_platform/internal/models"
	"time"
)

type MetadataOverview struct {
	DataHash        string    `gorm:"primaryKey;type:varchar(36)" json:"data_hash"`
	Name            string    `gorm:"index:idx_name;type:varchar(32)" json:"name"`
	Sex             bool      `gorm:"type:varchar(32)" json:"sex"`
	Category        string    `gorm:"type:varchar(32)" json:"category"`
	Format          string    `gorm:"type:varchar(32)" json:"format"`
	ContractAddress string    `gorm:"type:varchar(32)" json:"contract_address"`
	IsSharable      bool      `gorm:"type:tinyint(1)" json:"is_sharable"`
	CreatedAt       time.Time `gorm:"type:datetime" json:"created_at"`
	Tags            string    `gorm:"type:varchar(32)" json:"tags"`
}

type UpdateMetadata struct {
	GeneSharingContractAddress string `json:"gene_sharing_contract_address"`
	ContractAddress            string `json:"contract_address"`
	Description                string `json:"description"`
	Tags                       string `json:"tags"`
	IsSharable                 bool   `json:"is_sharable"`
}

type GetMetadataResp struct {
	Metadata    models.Metadata    `json:"metadata"`
	GeneSharing models.GeneSharing `json:"geneSharing"`
}

type GetMetadataOverviewByOwnerReq struct {
	Owner string `json:"owner" binding:"required"`
}

type GetMetadataOverviewByOwnerResp struct {
	Metadata []MetadataOverview `json:"multi_metadata"`
}

type GetAllMetadataOverviewResp struct {
	MultiMetadata []MetadataOverview `json:"multi_metadata"`
}

type NewViewAccessReq struct {
	DataHash string `json:"data_hash" binding:"required"`
	Remark   string `json:"remark" binding:"required"`
}
