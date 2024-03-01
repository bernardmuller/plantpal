package cms_db

import (
	"testing"
)

func TestCMSDBConnection(t *testing.T) {
	db, err := connect_cms_db()
	if err != nil {
		t.Fatalf(`db_connect() = %s`, err)
	}

  disconnect_cms_db(db)
}
