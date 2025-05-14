package models

import (
	"time"
)

type Collection struct {
	Id              string    `gorm:"primaryKey;type:varchar(32)" json:"Id"`             //以profile_id作为collection_id
	Name            string    `gorm:"type:varchar(32)" json:"name"`                      //藏品集合名称
	Address         string    `gorm:"type:varchar(42)" json:"address"`                   //藏品集合地址
	Description     string    `gorm:"type:text" json:"description"`                      //藏品集合描述
	Creator         string    `gorm:"type:varchar(42);index:idx_creator" json:"creator"` //藏品集合创建者,并构建索引idx_creator(creator)
	CreatorEarning  string    `gorm:"type:varchar(32)" json:"creator_earning"`           //藏品集合创建者收益
	AvailableRegion string    `gorm:"type:varchar(32)" json:"available_region"`          //藏品集合可操作区域
	CreatedAt       time.Time `gorm:"type:datetime" json:"created_at"`                   //藏品集合创建时间
	ExplorerLink    string    `gorm:"type:varchar(255)" json:"explorer_link"`            //区块浏览器链接
	ItemAmount      int       `gorm:"type:int;default:8" json:"item_amount"`             //GNFT数量
}
