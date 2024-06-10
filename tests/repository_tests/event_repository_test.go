package repository_tests

import (
	"eventful/models"
	"eventful/repositories"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrate the schema
	db.AutoMigrate(&models.Event{})

	return db, nil
}

func TestEventRepository(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Failed to setup test database: %v", err)
	}

	repo := repositories.NewEventRepository(db)

	event := models.Event{
		Title:       "Test Event",
		Description: "This is a test event",
		StartDate:   time.Now(),
		EndDate:     time.Now().Add(time.Hour * 24),
	}

	err = repo.CreateEvent(&event)
	if err != nil {
		t.Fatalf("Failed to create event: %v", err)
	}

	assert.NotEqual(t, 0, event.ID)

	fetchedEvent, err := repo.GetEventByID(event.ID)
	if err != nil {
		t.Fatalf("Failed to fetch event: %v", err)
	}

	assert.Equal(t, event.Title, fetchedEvent.Title)
	assert.Equal(t, event.Description, fetchedEvent.Description)
	assert.Equal(t, event.StartDate.Unix(), fetchedEvent.StartDate.Unix())
	assert.Equal(t, event.EndDate.Unix(), fetchedEvent.EndDate.Unix())

	events, err := repo.GetEvents()
	if err != nil {
		t.Fatalf("Failed to fetch events: %v", err)
	}

	assert.NotEmpty(t, events)

	event.Title = "Updated Test Event"
	err = repo.UpdateEvent(&event)
	if err != nil {
		t.Fatalf("Failed to update event: %v", err)
	}

	updatedEvent, err := repo.GetEventByID(event.ID)
	if err != nil {
		t.Fatalf("Failed to fetch updated event: %v", err)
	}

	assert.Equal(t, event.Title, updatedEvent.Title)

	err = repo.DeleteEvent(event.ID)
	if err != nil {
		t.Fatalf("Failed to delete event: %v", err)
	}

	deletedEvent, err := repo.GetEventByID(event.ID)
	assert.Nil(t, deletedEvent)
	assert.NotNil(t, err)
}

func TestEventRepository_GetEventByID(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Failed to setup test database: %v", err)
	}

	repo := repositories.NewEventRepository(db)

	event := models.Event{
		ID:          1,
		Title:       "Test Event",
		Description: "This is a test event",

		StartDate: time.Now(),
		EndDate:   time.Now().Add(time.Hour * 24),
		Location:  "Lagos",
	}

	err = repo.CreateEvent(&event)
	if err != nil {
		t.Fatalf("Failed to create event: %v", err)
	}

	fetchedEvent, err := repo.GetEventByID(event.ID)
	if err != nil {
		t.Fatalf("Failed to fetch event: %v", err)
	}

	assert.Equal(t, event.ID, fetchedEvent.ID)
	assert.Equal(t, event.Description, fetchedEvent.Description)
	assert.Equal(t, event.StartDate.Unix(), fetchedEvent.StartDate.Unix())
	assert.Equal(t, event.EndDate.Unix(), fetchedEvent.EndDate.Unix())
}

func TestEventRepository_GetEvents(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Failed to setup test database: %v", err)
	}

	repo := repositories.NewEventRepository(db)

	event1 := models.Event{
		Title:       "Test Event 1",
		Description: "This is test event 1",
		StartDate:   time.Now(),
		EndDate:     time.Now().Add(time.Hour * 24),
	}

	event2 := models.Event{
		Title:       "Test Event 2",
		Description: "This is test event 2",
		StartDate:   time.Now(),
		EndDate:     time.Now().Add(time.Hour * 24),
	}

	err = repo.CreateEvent(&event1)
	assert.NoError(t, err)

	err = repo.CreateEvent(&event2)
	assert.NoError(t, err)

	events, err := repo.GetEvents()
	assert.NoError(t, err)
	assert.Len(t, events, 2)
}

func TestEventRepository_CreateEvent(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Failed to setup test database: %v", err)
	}

	repo := repositories.NewEventRepository(db)

	event := models.Event{
		Title:       "Test Event",
		Description: "This is a test event",
		StartDate:   time.Now(),
		EndDate:     time.Now().Add(time.Hour * 24),
	}

	err = repo.CreateEvent(&event)
	assert.NoError(t, err)
	assert.NotZero(t, event.ID)
}

func TestEventRepository_UpdateEvent(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Failed to setup test database: %v", err)
	}

	repo := repositories.NewEventRepository(db)

	event := models.Event{
		Title:       "Test Event",
		Description: "This is a test event",
		StartDate:   time.Now(),
		EndDate:     time.Now().Add(time.Hour * 24),
	}

	err = repo.CreateEvent(&event)
	assert.NoError(t, err)

	event.Title = "Updated Test Event"
	err = repo.UpdateEvent(&event)
	assert.NoError(t, err)

	updatedEvent, err := repo.GetEventByID(event.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Test Event", updatedEvent.Title)
}

func TestEventRepository_DeleteEvent(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Failed to setup test database: %v", err)
	}

	repo := repositories.NewEventRepository(db)

	event := models.Event{
		Title:       "Test Event",
		Description: "This is a test event",
		StartDate:   time.Now(),
		EndDate:     time.Now().Add(time.Hour * 24),
	}

	err = repo.CreateEvent(&event)
	assert.NoError(t, err)

	err = repo.DeleteEvent(event.ID)
	assert.NoError(t, err)

	deletedEvent, err := repo.GetEventByID(event.ID)
	assert.Nil(t, deletedEvent)
	assert.NotNil(t, err)
}
