package main

import (
	"log"

	"github.com/smochii/go-clean-architecture/infrastructure/database"
	"github.com/smochii/go-clean-architecture/infrastructure/database/model"
)

func main() {
	db := database.NewConnection()
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("Failed to close database: %v", err)
		}
		sqlDB.Close()
	}()

	err := db.AutoMigrate(
		&model.User{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
