package user

// 绑定post的
type EditUserName struct {
	Name    string `json:"new_name"`
	Address string `json:"user_address"`
}

// 数据库操作的
type UpdateUserName struct {
	ID      uint
	Address string `gorm:"column:address"`
	Name    string `gorm:"column:name"`
}
