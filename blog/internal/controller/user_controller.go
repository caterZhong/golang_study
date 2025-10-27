package controller

import (
	"golang_study/blog/internal/common"
	"golang_study/blog/internal/dto/user_dto"
	"golang_study/blog/internal/service"

	"github.com/gin-gonic/gin"
)

// 业务控制器结构体
type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (u *UserController) Login(c *gin.Context) {
	// 处理登录请求
	// 获取post 通过json传输的参数email和password， 验证用户信息， 生成token 返回
	var loginDTO user_dto.LoginDTO
	if err := c.ShouldBind(&loginDTO); err != nil {
		common.BadRequest(c, common.ErrInvalidParams.Message)
	}

	jwt, err := u.userService.Login(loginDTO.Email, loginDTO.Password)

	if err != nil {
		// 检查是否为自定义业务错误
		if bizErr, ok := err.(*common.BizError); ok {
			common.ErrorWithBiz(c, bizErr)
			return
		}

		// 其他类型错误
		common.InternalServerError(c, "服务器异常")
		return
	}

	// 成功返回jwt
	common.Success(c, jwt)
	return
}

func (u *UserController) GetUserInfo(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "获取用户信息成功",
	})
}
