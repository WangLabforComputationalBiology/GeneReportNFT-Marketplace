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

type GetToken struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

type GetReportId struct {
	Id       string    `json:"id"`
	Profiles []Profile `json:"profiles"`
}

type Profile struct {
	Id     string `json:"id"`
	Format string `json:"format"`
	Sex    int    `json:"sex"`
}
