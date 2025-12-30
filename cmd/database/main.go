// This file is temporary and will be deleted along with the cmd/database folder
package main

import (
	"github.com/freedom-sketch/project-noob/internal/database"
	"github.com/freedom-sketch/project-noob/internal/logger"
)

func main() {
	log, _ := logger.New("database.log")

	if err := database.Connect(); err != nil {
		log.Fatal("Connect failed:", err)
	}

	log.Info("✅ Successful connection to the database")
}
