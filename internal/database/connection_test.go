package database

import (
	"testing"
)

func TestConnect(t *testing.T) {
	db, err := ConnectInMemory()
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	if db == nil {
		t.Fatal("db should not be nil")
	}

	var count int
	err = db.Raw("SELECT COUNT(*) FROM sqlite_master WHERE type='table' AND name='users'").Scan(&count).Error
	if err != nil {
		t.Fatalf("SQLite query error: %v", err)
	}

	t.Log("Migrations were successful, tables were created")
}
