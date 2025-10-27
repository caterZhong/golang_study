package service

import (
	"golang_study/blog/internal/common"
	"golang_study/blog/internal/dto/user_dto"
	"golang_study/blog/internal/repo"
	"golang_study/blog/internal/utils"
)

type UserService struct {
	userRepository *repo.UserRepository
	jwtService     *utils.JWTService
}

func NewUserService(userRepository *repo.UserRepository, jwtService *utils.JWTService) *UserService {
	return &UserService{
		userRepository: userRepository,
		jwtService:     jwtService,
	}

}

// 用户登录服务
func (userService *UserService) Login(email string, password string) (*user_dto.LoginResponse, error) {
	// 1. 从数据库获取email 对应的用户信息， 校验密码
	user, err := userService.userRepository.GetByEmail(email)

	// 用户不存在， 抛出异常
	if err != nil {
		return nil, err

	}

	// 密码错误，抛出异常
	// password 为 加盐的哈希
	// 先获取 salt, 然后将前端传入的pass + salt 做哈希 和数据库中的password中的哈希比对， 一致为通过
	if !utils.CheckBcrypt(password, user.Password) {
		return nil, common.NewBusinessError(common.ErrPasswordWrong)
	}

	// 2. 如果校验通过， 生成token 返回
	// 生成新的访问令牌
	accessToken, err := userService.jwtService.GenerateToken(user.ID, user.UserName, user.Email)
	if err != nil {
		return nil, common.NewBusinessError(common.ErrTokenGenerate)
	}

	// 解析token获取过期时间
	claims, err := userService.jwtService.ParseToken(accessToken)
	if err != nil {
		return nil, common.NewBusinessError(common.ErrTokenGenerate)
	}

	return &user_dto.LoginResponse{
		UserID:      user.ID,
		Username:    user.UserName,
		Email:       user.Email,
		AccessToken: accessToken,
		ExpiresAt:   claims.ExpiresAt.Unix(),
		TokenType:   "Bearer",
	}, nil
}
