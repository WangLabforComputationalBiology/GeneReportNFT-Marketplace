package appErrors

import (
	"GeneReport_platform/api/dto"
	"errors"
	"fmt"
	"github.com/fatih/color"
	"log"
)

type IAppError interface {
	error // 嵌入标准 error 接口
	ToErrResponse() dto.ErrResponse
	ErrorWithDetail() string
}
type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`          //用户端模糊错误信息
	Detail  string `json:"detail,omitempty"` //内部错误详情
}

// New 仿errors.New
// 创建一个新的AppError实例
func New(code int, message string, errs ...error) AppError {
	if len(errs) > 0 {
		return AppError{
			Code: code, Message: message, Detail: errs[0].Error(),
		}
	} else {
		return AppError{
			Code: code, Message: message,
		}
	}
}

// 标准errors.Error接口
func (e AppError) Error() string {
	return fmt.Sprintf("code: %d ; message: %s", e.Code, e.Message)
}

// ErrorWithDetail 用于控制台输出调试附带详情
func (e AppError) ErrorWithDetail() string {
	redBold := color.New(color.FgRed, color.Bold)
	yellowBold := color.New(color.FgYellow, color.Bold)
	return redBold.Sprintf(fmt.Sprintf("↘↘↘ code: %d; message: %s;", e.Code, e.Message)) + yellowBold.Sprintf(fmt.Sprintf(" detail: %s;", e.Detail))
}

// ToErrResponse 转换为ErrResponse 进行响应
func (e AppError) ToErrResponse() dto.ErrResponse {
	log.Printf(e.ErrorWithDetail())
	return dto.ErrResponse{
		Code:    e.Code,
		Message: e.Message,
	}
}

func Is(err, target error) bool {
	//比较原生错误即可，如需比较自定义错误，请在此增加类型断言逻辑
	return errors.Is(err, target)
}
