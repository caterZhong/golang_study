package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func CheckBcrypt(inputPass string, hashpassWithSalt string) bool {
	// 从哈希字符串中提取盐值并验证（自动完成）
	err := bcrypt.CompareHashAndPassword([]byte(hashpassWithSalt), []byte(inputPass))
	if err != nil {
		return false
	}

	return true
}

func GenHashedPass(inputPass string) (string, error) {
	// 生成带盐的哈希（盐值自动生成并嵌入）
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(inputPass), bcrypt.DefaultCost)

	// 将hashedPassword 转为string类型
	return string(hashedPassword), err
}
