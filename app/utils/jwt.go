package utils

import (
	"lost-and-found-backend/configs/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserId   int
	UserName string
	Role     string
	jwt.RegisteredClaims
}

func GenerateJWT(userId int, userName string, role string) (string, error) {
	var key = []byte(config.Config.GetString("key"))
	claims := Claims{
		UserId:   userId,
		UserName: userName,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    config.Config.GetString("issuer"),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenString, nil

}
