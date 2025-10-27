package common

type ErrorCode struct {
	BizCode  int
	Message  string
	HttpCode int
}

var (
	ReponseSuccess = &ErrorCode{BizCode: 0, Message: "成功", HttpCode: 200}

	// 认证相关错误
	ErrUserNotFound  = &ErrorCode{BizCode: 10001, Message: "用户不存在", HttpCode: 401}
	ErrPasswordWrong = &ErrorCode{BizCode: 10002, Message: "密码错误", HttpCode: 401}
	ErrTokenInvalid  = &ErrorCode{BizCode: 10003, Message: "令牌无效", HttpCode: 401}
	ErrTokenExpired  = &ErrorCode{BizCode: 10004, Message: "令牌已过期", HttpCode: 401}
	ErrUnAuthored    = &ErrorCode{BizCode: 10005, Message: "非法操作", HttpCode: 401}

	// 参数相关错误
	ErrInvalidParams = &ErrorCode{BizCode: 20001, Message: "参数错误", HttpCode: 400}
	ErrMissingParams = &ErrorCode{BizCode: 20002, Message: "缺少必要参数", HttpCode: 400}
	ErrInvalidEmail  = &ErrorCode{BizCode: 20003, Message: "邮箱格式错误", HttpCode: 400}

	//业务错误
	ErrDatabaseBiz  = &ErrorCode{BizCode: 30001, Message: "数据库操作异常", HttpCode: 200}
	ErrPostNotFound = &ErrorCode{BizCode: 30002, Message: "文章不存在", HttpCode: 200}

	// 服务器错误
	ErrInternalServer = &ErrorCode{BizCode: 50001, Message: "服务器异常", HttpCode: 500}
	ErrDatabase       = &ErrorCode{BizCode: 50002, Message: "数据库异常", HttpCode: 500}
	ErrTokenGenerate  = &ErrorCode{BizCode: 50003, Message: "token生成错误", HttpCode: 500}
)

// func (e *ErrorCode) Error() string {
// 	return e.Message
// }
