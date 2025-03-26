package models

import (
	"time"
)

type User struct {
	Address  string    `gorm:"primaryKey;type:varchar(64)" json:"address"` //用户地址
	Name     string    `gorm:"type:varchar(30)" json:"name"`               //用户名
	Avatar   string    `gorm:"type:varchar(30)" json:"avatar"`             //用户头像文件路径
	CreateAt time.Time `gorm:"type:datetime" json:"create_at"`             //用户创建时间
	GNFTs    []GNFT    `gorm:"many2many:user_gnfts;" json:"gnfts"`         //用户持有的所有gnft
}
