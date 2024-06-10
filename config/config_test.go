package config

import (
	"eventful/models"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
}

func TestInitDB(t *testing.T) {
	// Create a temporary .env file
	f, err := os.CreateTemp("", ".env")
	assert.NoError(t, err)
	defer os.Remove(f.Name())

	_, err = f.WriteString("DB_USER=test\nDB_PASSWORD=test\nDB_HOST=localhost\nDB_PORT=3306\nDB_NAME=testdb")
	assert.NoError(t, err)

	err = godotenv.Load(f.Name())
	assert.NoError(t, err)

	InitDB()
	defer CloseDB()

	assert.NotNil(t, DB)

	// Check if tables were created
	err = DB.AutoMigrate(&models.Event{}, &models.User{}, &models.Team{}, &models.Role{})
	assert.NoError(t, err)
}

func TestCloseDB(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	DB = db

	CloseDB()

	sqlDB, err := DB.DB()
	assert.NoError(t, err)

	err = sqlDB.Ping()
	assert.Error(t, err, "sql: database is closed")
}
