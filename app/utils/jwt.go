package utils

import (
	"lost-and-found-backend/configs/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID   int
	Username string
	Role     string
	jwt.RegisteredClaims
}

// 生成token
func GenerateJWT(userID int, username string, role string) (string, time.Time, error) {
	var key = []byte(config.Config.GetString("jwt.key"))
	expiresAt := time.Now().Add(2 * time.Hour)
	claims := Claims{
		UserID:   userID,
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    config.Config.GetString("jwt.issuer"),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", time.Time{}, err
	}

	return tokenString, expiresAt, nil

}
