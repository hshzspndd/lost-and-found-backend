package utils

import (
	"lost-and-found-backend/configs/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserId   int
	Username string
	Role     string
	jwt.RegisteredClaims
}

// 生成token
func GenerateJWT(userId int, username string, role string) (string, time.Time, error) {
	var key = []byte(config.Config.GetString("jwt.key"))
	var expiredAt time.Time = time.Now().Add(2 * time.Hour)
	claims := Claims{
		UserId:   userId,
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    config.Config.GetString("jwt.issuer"),
			ExpiresAt: jwt.NewNumericDate(expiredAt),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", time.Now(), err
	}

	return tokenString, expiredAt, nil

}
