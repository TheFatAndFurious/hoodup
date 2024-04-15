package utils

import "golang.org/x/crypto/bcrypt"

func Hashpassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 20)
	return string(bytes), err
}