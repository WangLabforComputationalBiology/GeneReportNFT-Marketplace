package dto

type CheckCAPTCHAReq struct {
	Angle int `json:"angle"`
}
type SendSMSCodeReq struct {
	Phone string `json:"phone" binding:"required"`
}
type VerifySMSCodeReq struct {
	Phone string `json:"phone" binding:"required"`
	Code  string `json:"code" binding:"required"`
}
type CreateCollectionReq struct {
	ProfileID      string `json:"profile_id"` //profile_id将作为collection_id
	CollectionName string `json:"collection_name"`
	Symbol         string `json:"symbol"`
	Salt           string `json:"salt"`
	Description    string `json:"description"`
}
