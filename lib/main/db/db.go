package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func connect_db() (*sql.DB, error) {
  // database_name := os.Getenv("TEST_DATABASE_URL")
  // auth_token := os.Getenv("TURSO_AUTH_TOKEN")
  //
  // if database_name == "" || auth_token == "" {
  //   return nil, fmt.Errorf("DATABASE_NAME and AUTH_TOKEN environment variables must be set")
  // }

  url := "libsql://domain-db-test-bernardmuller.turso.io?authToken=eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE3MDg5MzA3OTMsImlkIjoiNTQxOTZhYjgtZDQ3My0xMWVlLWIzMmEtZGU0YWQ3NTljZGZlIn0.HEKTJ5RbybF8ckh4KfRBgImE7Al-yRw7YU-ONUX8Ig7mwnYpUihlQrX3FCEDZVQZQkM4TDdugDzFHPqqxT3UBw"
  // url := fmt.Sprintf("[%s]?authToken=%s", database_name, auth_token)

  db, err := sql.Open("libsql", url)
  if err != nil {
    return nil, fmt.Errorf("failed to open db %s: %s", url, err)
  }
  return db, nil
}

