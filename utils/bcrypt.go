package utils

import (
	"golang.org/x/crypto/bcrypt"
)

var GBcrypt Bcrypt

type Bcrypt struct {
}

// HashPwd 将密码变成带盐值的哈希值
func (Bcrypt) HashPwd(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) // 第二个参数是cost值，越高安全性越强
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPwd 验证密码是否正确
func (Bcrypt) CheckPwd(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
