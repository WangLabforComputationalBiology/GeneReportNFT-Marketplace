package models

import (
	"time"
)

type GeneSharing struct {
	ContractAddress string    `gorm:"primaryKey;type:varchar(42)" json:"contract_address"`                   //sharing集合地址
	Name            string    `gorm:"type:varchar(32)" json:"name"`                                          //sharing集合名称
	Description     string    `gorm:"type:text" json:"description"`                                          //sharing集合描述
	CreatorAddress  string    `gorm:"type:varchar(42);index:idx_creatorgeneSharings" json:"creator_address"` //sharing集合创建者
	CreatedAt       time.Time `gorm:"type:datetime" json:"created_at"`                                       //sharing集合创建时间
	ExplorerLink    string    `gorm:"type:varchar(255)" json:"explorer_link"`                                //链上交易链接
	ItemAmount      int       `gorm:"type:int;" json:"item_amount"`                                          //MetaData数量
	IsOfficial      bool      `gorm:"type:tinyint(1);" json:"is_official"`                                   //是否第三方官方授权构建
	Tags            string    `gorm:"type:varchar(255)" json:"tags"`                                         //标签，以分号分隔，,ex:third party:wegene;...
}
