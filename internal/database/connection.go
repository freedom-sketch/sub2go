package database

import (
	"github.com/freedom-sketch/sub2go/config"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func Connect(cfg *config.DataBase) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(cfg.Name+".db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := AutoMigrate(db); err != nil {
		return nil, err
	}

	return db, nil
}

func ConnectInMemory() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := AutoMigrate(db); err != nil {
		return nil, err
	}
	return db, nil
}
