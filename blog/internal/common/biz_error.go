package common

import "fmt"

// BusinessError 业务错误
type BizError struct {
	Code    int
	Message string
}

func (e *BizError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

// NewBusinessError 创建业务错误
func NewBusinessError(errorCode *ErrorCode) *BizError {
	return &BizError{
		Code:    errorCode.BizCode,
		Message: errorCode.Message,
	}
}

// NewBusinessErrorWithMessage 创建带自定义消息的业务错误
func NewBusinessErrorWithMessage(code int, message string) *BizError {
	return &BizError{
		Code:    code,
		Message: message,
	}
}
