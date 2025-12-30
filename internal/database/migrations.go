package database

import "github.com/freedom-sketch/project-noob/internal/database/models"

func AutoMigrate() error {
	return DB.AutoMigrate(
		&models.User{},
		&models.Admin{},
		&models.Subscription{},
		&models.Country{},
		&models.Server{},
	)
}
