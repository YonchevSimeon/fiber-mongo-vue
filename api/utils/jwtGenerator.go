package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJwtToken() (string, error) {
	secret := os.Getenv("JWT_SECRET")

	claims := jwt.MapClaims{}

	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tkn, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return tkn, nil
}
