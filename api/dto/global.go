package dto

// Response 结构化响应
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// ErrResponse 该dto当前仅用于swagger文档注释
type ErrResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
