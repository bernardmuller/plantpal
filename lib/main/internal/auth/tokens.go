package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"os"
	"time"
)

func getSecretKey() (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", fmt.Errorf("failed to load environment variables: %w", err)
	}
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		return "", fmt.Errorf("JWT_SECRET environment variables must be set")
	}
	return secretKey, nil
}

func CreateToken(payload interface{}) (string, error) {
	secretKey, err := getSecretKey()
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"payload": payload,
			"exp":     time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	secretKey, err := getSecretKey()
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
