package config

import (
	"domain-app/internal/store/postgres"
	"fmt"
	"net/http"
)

type ApiConfig struct {
	Database *postgres.Queries
	Router   *http.ServeMux
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
