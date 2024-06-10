package repository_tests

import (
	"eventful/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to test database: %v", err)
	}

	// Migrate the schema
	err = db.AutoMigrate(&models.Event{}, &models.User{}, &models.Team{}, &models.Role{})
	if err != nil {
		log.Fatalf("failed to migrate test database: %v", err)
	}

	return db
}
