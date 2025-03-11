package dto

import "time"

// user的全量数据
type Users struct {
	ID         uint
	Address    string    `gorm:"column:address"`
	Name       string    `gorm:"column:name"`
	Picture    string    `gorm:"column:picture"`
	CreateTime time.Time `gorm:"column:create_time"`
}

// 绑定post的
type EditUserName struct {
	Name      string `json:"new_name"`
	Address   string `json:"user_address"`
	Signature string `json:"signature"`
}

// 数据库操作的
type UpdateUser struct {
	ID        uint
	Address   string `gorm:"column:address"`
	Name      string `gorm:"column:name"`
	Picture   string `gorm:"column:picture"`
	Signature string `gorm:"column:signature"`
}
