package dto

import (
	"GeneReport_platform/internal/models"
	"time"
)

type GeneSharingOverview struct {
	ContractAddress string    `gorm:"primaryKey;type:varchar(42)" json:"contract_address"` //sharing集合地址
	Name            string    `gorm:"type:varchar(32)" json:"name"`                        //sharing集合名称
	Description     string    `gorm:"type:text" json:"description"`                        //sharing集合描述
	Creator         string    `gorm:"type:varchar(42);index:idx_creator" json:"creator"`   //sharing集合创建者
	CreatedAt       time.Time `gorm:"type:datetime" json:"created_at"`                     //sharing集合创建时间
	ItemAmount      int       `gorm:"type:int;" json:"item_amount"`                        //MetaData数量
	IsOfficial      bool      `gorm:"type:tinyint(1);" json:"is_official"`                 //是否第三方官方授权构建
	Tags            string    `gorm:"type:varchar(255)" json:"tags"`                       //标签，以分号分隔，,ex:third party:wegene;...
}

type GetGeneSharingDetailsByContractAddressReq struct {
	GeneSharingContractAddress string `json:"geneSharing_contract_address"` //集合ID
}

type GetGeneSharingDetailsByContractAddressResp struct {
	GeneSharing       models.GeneSharing
	MetadataOverviews []MetadataOverview `json:"metadata_overviews"` //该集合下的所有GNFT
}

type GetGeneSharingOverviewByCreatorReq struct {
	CreatorAddress string `json:"creator_address"` //创建者地址
}

type GetGeneSharingOverviewByCreatorResp struct {
	GeneSharings []GeneSharingOverview `json:"geneSharings"`
}

type GetAllGeneSharingOverviewResp struct {
	MultiGeneSharing []GeneSharingOverview `json:"multi_geneSharing"`
	PageNum          int                   `json:"page_num"`
}
