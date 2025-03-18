package errors

//
//import (
//	"fmt"
//	"github.com/fatih/color"
//)
//
//const (
//	ErrCodeInvalidInput int32 = 1001
//	ErrCodeNotFound     int32 = 1002
//	ErrCodeInternal     int32 = 1003
//	ErrCodeDefault      int32 = 1004
//)
//
//var errMessages = map[int32]string{
//	ErrCodeInvalidInput: "Invalid input",
//	ErrCodeNotFound:     "Resource not found",
//	ErrCodeInternal:     "Internal server error",
//	ErrCodeDefault:      "Unknown error",
//}
//
//type CustomError struct {
//	Code    int32  `json:"code"`
//	Message string `json:"message"`
//	Detail  string `json:"detail"`
//}
//
//// New 标准errors.New接口
//func New(code int32, err ...interface{}) *CustomError {
//	//类型断言
//	if len(err) > 0 && err[0] != nil {
//		if _, ok := err[0].(error); ok {
//			return &CustomError{
//				Code: code, Message: errMessages[code],
//				Detail: fmt.Sprintf("%v", err[0].Error()),
//			}
//		}
//	}
//	return &CustomError{
//		Code: code, Message: errMessages[code],
//	}
//
//}
//
//// 标准errors.Error接口
//func (e *CustomError) Error() string {
//	return fmt.Sprintf("code: %d ; message: %s", e.Code, e.Message)
//}
//
//// ErrorWithColor 用于控制台输出调试
//func (e *CustomError) ErrorWithColor() string {
//	redBold := color.New(color.FgRed, color.Bold)
//	return redBold.Sprintf(fmt.Sprintf("↘↘↘ code: %d; message: %s;", e.Code, e.Message))
//}
//
//// ErrorWithDetail 用于控制台输出调试附带详情
//func (e *CustomError) ErrorWithDetail() string {
//	yellowBold := color.New(color.FgYellow, color.Bold)
//	return e.ErrorWithColor() + yellowBold.Sprintf(fmt.Sprintf(" detail: %s;", e.Detail))
//}
