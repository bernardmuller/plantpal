package config

import (
	"domain-app/internal/store/postgres"
	"fmt"
	"github.com/labstack/echo/v4"
)

type ApiConfig struct {
	Database *postgres.Queries
	Router   *echo.Echo
	PORT     string
}

func CreateAPIConfig() *ApiConfig {
	router := CreateRouter()
	database, err := postgres.ConnectDB()
	if err != nil {
		fmt.Println("Error connecting to DB: ", err)
	}
	config := ApiConfig{
		Database: database,
		Router:   router,
		PORT:     ":8080",
	}
	return &config
}
