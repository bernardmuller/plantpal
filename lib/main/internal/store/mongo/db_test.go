package mongo

import (
	"testing"
)

func TestCMSDBConnection(t *testing.T) {
	db, err := Connect_cms_db()
	if err != nil {
		t.Fatalf(`db_connect() = %s`, err)
	}

	Disconnect_cms_db(db)
}
