package models

import (
	"time"
)

type User struct {
	Address  string    `gorm:"primaryKey;type:varchar(42)" json:"address"`      //用户地址
	Name     string    `gorm:"type:varchar(32)" json:"name"`                    //用户名
	Avatar   string    `gorm:"type:varchar(32)" json:"avatar"`                  //用户头像文件路径
	CreateAt time.Time `gorm:"type:datetime" json:"create_at"`                  //用户创建时间
	Country  string    `gorm:"type:varchar(32);default:UNKNOWN" json:"country"` //用户所在国家
}
