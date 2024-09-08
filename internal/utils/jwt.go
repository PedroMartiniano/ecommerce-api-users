package utils

import (
	"errors"
	"github.com/PedroMartiniano/ecommerce-api-users/internal/configs"
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
)

func GenerateJWT(id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	secretJWT := configs.GetEnv("SECRET_JWT")

	signedToken, err := token.SignedString([]byte(secretJWT))

	return signedToken, err
}

func VerifyJWT(token string) (string, error) {
	tokenArr := strings.Split(token, " ")

	if len(tokenArr) != 2 {
		return "", errors.New("invalid token")
	}

	secretJWT := configs.GetEnv("SECRET_JWT")

	parsedToken, err := jwt.Parse(tokenArr[1], func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(secretJWT), nil
	})

	if err != nil {
		return "", err
	}

	if !parsedToken.Valid {
		return "", errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid token")
	}

	id, _ := claims["id"].(string)

	return id, nil
}
