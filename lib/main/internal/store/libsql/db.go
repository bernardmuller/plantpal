package libsql

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func ConnectDb() (*sql.DB, error) {

	databaseName := os.Getenv("TEST_DATABASE_URL")
	authToken := os.Getenv("TURSO_AUTH_TOKEN")

	if databaseName == "" || authToken == "" {
		return nil, fmt.Errorf("TEST_DATABASE_URL and TURSO_AUTH_TOKEN environment variables must be set")
	}

	url := fmt.Sprintf("[%s]?authToken=%s", databaseName, authToken)

	db, err := sql.Open("libsql", url)

	if err != nil {
		return nil, fmt.Errorf("failed to open db %s: %s", url, err)
	}
	return db, nil
}
