package dto

type LoginReq struct {
	UserAddress string `json:"user_address" binding:"required"`
	Signature   string `json:"signature" binding:"required"`
}

// UpdateUser userModel for update
type UpdateUser struct {
	Address string `gorm:"column:address"`
	Name    string `gorm:"column:name"`
	Avatar  string `gorm:"column:avatar"`
}
