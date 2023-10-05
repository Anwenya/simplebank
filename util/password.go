package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword 对明文密码进行hash后返回
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password %w", err)
	}
	return string(hashedPassword), nil
}

// 检查明文密码与hashed密码是否匹配
func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
