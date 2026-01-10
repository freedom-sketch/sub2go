package database

import "testing"

func TestConnect(t *testing.T) {
	db, err := ConnectInMemory()
	if err != nil {
		t.Fatalf("failed to connect: %v", err)
	}

	requiredTables := []string{"users", "subscriptions", "servers", "inbounds", "admins"}

	for _, table := range requiredTables {
		var count int
		err := db.Raw("SELECT COUNT(*) FROM sqlite_master WHERE type='table' AND name=?", table).Scan(&count).Error
		if err != nil {
			t.Fatalf("query error for table %s: %v", table, err)
		}
		if count == 0 {
			t.Errorf("table %s was not created", table)
		}
	}

	t.Log("Migrations were successful, tables were created")
}
