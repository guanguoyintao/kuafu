package ecrypto

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword 生成密码hash
func HashPassword(passwd string) (string, error) {
	return HashPasswordWithCost(passwd, bcrypt.DefaultCost)
}

// HashPasswordWithCost 生成密码hash，cost取值范围[4,31]
func HashPasswordWithCost(passwd string, cost int) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(passwd), cost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// VerifyPassword 校验密码与密码hash是否匹配
func VerifyPassword(hashed, passwd string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(passwd)); err != nil {
		return false
	}
	return true
}

func XorEncryptDecrypt(input, key string) string {
	output := make([]byte, len(input))
	for i := 0; i < len(input); i++ {
		output[i] = input[i] ^ key[i%len(key)]
	}
	return string(output)
}
