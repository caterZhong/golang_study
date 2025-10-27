package utils

import (
	"errors"
	"fmt"
	"golang_study/blog/config"
	"time"

	"sync"

	"github.com/golang-jwt/jwt/v4"
)

// 自定义Claims
type CustomClaims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

var (
	JWT  *JWTService
	once sync.Once
)

// JWT服务
type JWTService struct {
	config *config.JWTConfig
}

// 新建JWT服务
func NewJWTService(config *config.JWTConfig) error {
	once.Do(func() {
		JWT = &JWTService{
			config: config,
		}
	})
	return nil

}

// 生成Token
func (j *JWTService) GenerateToken(userID uint, username string, email string) (string, error) {
	claims := CustomClaims{
		UserID:   userID,
		Username: username,
		Email:    email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.config.ExpiresTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			// Issuer:    j.config.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.config.Secret))
}

// 生成刷新Token
func (j *JWTService) GenerateRefreshToken(userID int64) (string, error) {
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.config.RefreshExpires)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		// Issuer:    j.config.Issuer,
		Subject: fmt.Sprintf("%d", userID),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.config.Secret))
}

// 解析Token
func (j *JWTService) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.config.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// 解析刷新Token
func (j *JWTService) ParseRefreshToken(tokenString string) (int64, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.config.Secret), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		var userID int64
		fmt.Sscanf(claims.Subject, "%d", &userID)
		return userID, nil
	}

	return 0, errors.New("invalid refresh token")
}

// 刷新Token
// func (j *JWTService) RefreshToken(refreshToken string, user *UserInfo) (string, error) {
//     userID, err := j.ParseRefreshToken(refreshToken)
//     if err != nil {
//         return "", err
//     }

//     if userID != user.UserID {
//         return "", errors.New("user id mismatch")
//     }

//     return j.GenerateToken(user.UserID, user.Username, user.Email)
// }

// 用户信息
type UserInfo struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// 从Token获取用户信息
// func (j *JWTService) GetUserInfoFromToken(tokenString string) (*model.User, error) {
//     claims, err := j.ParseToken(tokenString)
//     if err != nil {
//         return nil, err
//     }
// 	user := &model.User{
// 		ID: claims.ID,
//         UserName: claims.Username,
//         Email:    claims.Email,
// 	}

// 	fmt.Println(user)

// 	return &model.User{
// 		ID: claims.ID,
//         UserName: claims.Username,
//         Email:    claims.Email,
// 	}, nil
// }
