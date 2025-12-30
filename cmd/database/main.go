package main

import (
	"log"

	"github.com/freedom-sketch/project-noob/internal/database"
)

func main() {
	err := database.Connect()
	if err != nil {
		log.Fatal("Connect failed:", err)
	}
	log.Println("✅ Database file created successfully")
}
