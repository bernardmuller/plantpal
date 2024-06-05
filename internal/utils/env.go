package utils

import (
	"fmt"
	"github.com/joho/godotenv"
)

func InitEnv() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("failed to load environment variables: %w", err)
	}
	return nil
}
