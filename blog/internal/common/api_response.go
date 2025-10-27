package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 统一返回结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	TraceID string      `json:"traceId,omitempty"`
}

// PageResult 分页返回结构
type PageResult struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}

// 成功返回
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}

// 带分页的成功返回
func SuccessWithPage(c *gin.Context, list interface{}, total int64, page, pageSize int) {
	Success(c, PageResult{
		List:     list,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	})
}

// 失败返回
func Error(c *gin.Context, code int, message string) {
	c.AbortWithStatusJSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}

// 失败返回
func ErrorWithBiz(c *gin.Context, err *BizError) {
	c.AbortWithStatusJSON(http.StatusOK, Response{
		Code:    err.Code,
		Message: err.Message,
		Data:    nil,
	})
}

// 参数错误返回
func BadRequest(c *gin.Context, message string) {
	Error(c, 400, message)
}

// 未授权返回
func Unauthorized(c *gin.Context, message string) {
	Error(c, 401, message)
}

// 禁止访问返回
func Forbidden(c *gin.Context, message string) {
	Error(c, 403, message)
}

// 服务器错误返回
func InternalServerError(c *gin.Context, message string) {
	Error(c, 500, message)
}
