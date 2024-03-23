package postgres

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDB() (*Queries, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load environment variables: %w", err)
	}
	uri := os.Getenv("POSTGRES_URI")
	if uri == "" {
		return nil, fmt.Errorf("POSTGRES_URI environment variables must be set")
	}
	conn, err2 := sql.Open("postgres", uri)
	if err2 != nil {
		return nil, fmt.Errorf("failed to open db %s: %w", uri, err)
	}

	return New(conn), nil
}
