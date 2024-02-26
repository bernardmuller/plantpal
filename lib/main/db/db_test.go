package db

import (
	"testing"
)

func TestDBConnection(t *testing.T) {
  _, err := connect_db()
  if err != nil {
    t.Fatalf(`db_connect() = %s`, err)
  }
}
