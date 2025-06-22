package models

import (
	"time"
)

type User struct {
	Address     string    `gorm:"primaryKey;type:varchar(42)" json:"address"` //用户地址
	Name        string    `gorm:"type:varchar(32)" json:"name"`               //用户名
	CreateAt    time.Time `gorm:"type:datetime" json:"create_at"`             //用户创建时间
	Institution string    `gorm:"type:varchar(32)" json:"institution"`        //用户机构
	Email       string    `gorm:"type:varchar(32)" json:"email"`              //用户机构邮箱
}
