package libsql

import (
	"testing"
)

func TestDBConnection(t *testing.T) {
	_, err := ConnectDb()
	if err != nil {
		t.Fatalf(`db_connect() = %s`, err)
	}
}
