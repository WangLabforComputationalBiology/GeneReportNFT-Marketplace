package dto

type LoginReq struct {
	UserAddress string `json:"user_address" binding:"required"`
	Signature   string `json:"signature" binding:"required"`
}

// UpdateUser userModel for update
type UpdateUser struct {
	Address string `json:"address"`
	Name    string `json:"name"`
	Avatar  string `json:"avatar"`
}

type GetToken struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

type UserInfoRes struct {
	Address string `json:"address" binding:"required"`
	Name    string `json:"name" binding:"required"`
	Avatar  string `json:"avatar" binding:"required"`
	Country string `json:"country" binding:"required"`
}

type GetReportId struct {
	Id    int    `json:"id"`
	Email string `json:"email"`

	Profiles []Profile `json:"profiles"`
}

type Profile struct {
	Id     string `json:"id"`
	Format string `json:"format"`
	Name   string `json:"name"`
	Sex    int    `json:"sex"`
}

type ToSave struct {
	Code      string `json:"code"`
	ProfileId string `json:"profileId"`
}
