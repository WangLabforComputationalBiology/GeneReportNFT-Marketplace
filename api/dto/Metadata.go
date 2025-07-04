package dto

import (
	"GeneReport_platform/internal/models"
	"time"
)

type MetadataOverview struct {
	DataHash        string    `gorm:"primaryKey;type:varchar(36)" json:"data_hash"`
	Name            string    `gorm:"index:idx_name;type:varchar(32)" json:"name"`
	Description     string    `gorm:"type:varchar(255)" json:"description"`
	Sex             bool      `gorm:"type:varchar(32)" json:"sex"`
	Category        string    `gorm:"type:varchar(32)" json:"category"`
	Format          string    `gorm:"type:varchar(32)" json:"format"`
	ContractAddress string    `gorm:"type:varchar(32)" json:"contract_address"`
	IsSharable      bool      `gorm:"type:tinyint(1)" json:"is_sharable"`
	CreatedAt       time.Time `gorm:"type:datetime" json:"created_at"`
	Tags            string    `gorm:"type:varchar(32)" json:"tags"`
}

type UpdateMetadata struct {
	Name                       string `json:"name"`
	GeneSharingContractAddress string `gorm:"column:geneSharing_contract_address" json:"geneSharing_contract_address"`
	ContractAddress            string `json:"contract_address"`
	Description                string `json:"description"`
	Tags                       string `json:"tags"`
	IsSharable                 bool   `json:"is_sharable"`
	IsHidden                   bool   `json:"is_hidden"`
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
	PageNum       int                `json:"page_num"`
}

type GetGenoTypeZipResp struct {
	DownloadURL  string `json:"download_url"`
	AccessStatus bool   `json:"access_status"`
}

type NewViewAccessReq struct {
	DataHash string `json:"data_hash" binding:"required"`
	Label    string `json:"label" binding:"required"`
}

type NewViewAccessResp struct {
	DownloadURL string `json:"download_url"`
}
