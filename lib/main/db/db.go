package db

import (
	"database/sql"
	"fmt"
  "os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func connect_db() (*sql.DB, error) {
  database_name := os.Getenv("TEST_DATABASE_URL")
  auth_token := os.Getenv("TURSO_AUTH_TOKEN")

  if database_name == "" || auth_token == "" {
    return nil, fmt.Errorf("DATABASE_NAME and AUTH_TOKEN environment variables must be set")
  }

  url := fmt.Sprintf("[%s]?authToken=%s", database_name, auth_token)

  db, err := sql.Open("libsql", url)
  if err != nil {
    return nil, fmt.Errorf("failed to open db %s: %s", url, err)
  }
  return db, nil
}

