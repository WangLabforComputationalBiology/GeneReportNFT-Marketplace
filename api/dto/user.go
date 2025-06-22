package dto

import "time"

type LoginReq struct {
	UserAddress string `json:"user_address" binding:"required"`
	Signature   string `json:"signature" binding:"required"`
}

type UserInfoResp struct {
	Address     string    `json:"address" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Institution string    `json:"institution" binding:"required"`
	CreateAt    time.Time `json:"create_at"  binding:"required"`
	Email       string    `json:"email"    binding:"required"`
}
type GetToken struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
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

// CheckCAPTCHAReq 人机校验
type CheckCAPTCHAReq struct {
	Angle int `json:"angle"`
}

// SendEmailCodeReq 发送邮箱验证码
type SendEmailCodeReq struct {
	Institution string `json:"institution" binding:"required"`
	Email       string `json:"email" binding:"required"`
}

// VerifyEmailCodeReq 验证邮箱验证码
type VerifyEmailCodeReq struct {
	Email string `json:"email" binding:"required"`
	Code  string `json:"code" binding:"required"`
}
