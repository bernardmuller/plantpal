package config

import (
	"fmt"
	"github.com/bernardmuller/plantpal/store/postgres"
)

type ModuleConfig struct {
	Database *postgres.Queries
	PORT     string
}

func CreateConfig() *ModuleConfig {
	database, err := postgres.ConnectDB()
	if err != nil {
		fmt.Println("Error connecting to DB: ", err)
	}
	config := ModuleConfig{
		Database: database,
		PORT:     ":8000",
	}
	return &config
}
