package utils

import "golang.org/x/crypto/bcrypt"

func GenerateHashPassword(password string) (string, error) {
	bytePass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytePass), nil
}

func CompareHashAndPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
