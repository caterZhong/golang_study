package middleware

import (
	"strings"

	"golang_study/blog/internal/common"
	"golang_study/blog/internal/utils"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 获取并验证token字符串
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			// c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization header required"})
			common.Unauthorized(c, common.ErrTokenInvalid.Message)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			// c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "bearer token required"})
			common.Unauthorized(c, common.ErrTokenInvalid.Message)
			return
		}

		// 2. 解析token（处理错误！）
		// token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 	// 验证签名方法
		// 	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		// 		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		// 	}

		// 	// 获取密钥
		// 	secret := os.Getenv("JWT_SECRET")
		// 	if secret == "" {
		// 		return nil, fmt.Errorf("JWT_SECRET not configured")
		// 	}
		// 	return []byte(secret), nil
		// })

		// // 3. 处理解析错误
		// if err != nil {
		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token: " + err.Error()})
		// 	return
		// }

		// // 4. 验证token有效性
		// if token == nil || !token.Valid {
		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		// 	return
		// }

		claims, ok := utils.JWT.ParseToken(tokenString)

		// 5. 提取claims
		// claims, ok := token.Claims.(jwt.MapClaims)
		// if !ok {
		if ok != nil {
			common.Unauthorized(c, ok.Error())
			// c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token claims"})
			return
		}

		// 6. 验证必要字段存在

		if claims.UserID == 0 {
			common.Unauthorized(c, "userID not found in token")
			// c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "userID not found in token"})
			return
		}

		// 7. 设置到context（建议添加类型转换）
		c.Set("userID", claims.UserID)
		// if claims["roles"] != nil {
		// 	c.Set("roles", claims["roles"])
		// } else {
		// 	c.Set("roles", []string{}) // 默认空数组
		// }

		c.Next()
	}
}
