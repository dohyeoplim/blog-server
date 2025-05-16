package services

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID string) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	expMin := os.Getenv("JWT_EXP_MINUTES")
	duration := time.Minute * 60
	if expMin != "" {
		if d, err := time.ParseDuration(expMin + "m"); err == nil {
			duration = d
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(duration).Unix(),
	})

	return token.SignedString([]byte(secret))
}

func ParseJWT(tokenString string) (*jwt.Token, error) {
	secret := os.Getenv("JWT_SECRET")
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
}
