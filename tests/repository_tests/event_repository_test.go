package repository_tests

import (
	"eventful/models"
	"eventful/repositories"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEventRepository(t *testing.T) {
	db := SetupTestDB()
	repo := repositories.NewEventRepository(db)

	event := models.Event{
		Title:       "Test Event",
		Description: "This is a test event",
		StartDate:   time.Now(),
		EndDate:     time.Now().Add(24 * time.Hour),
	}

	// Create Event
	err := repo.CreateEvent(&event)
	assert.NoError(t, err)
	assert.NotZero(t, event.ID)

	// Get Event by ID
	fetchedEvent, err := repo.GetEventByID(event.ID)
	assert.NoError(t, err)
	assert.Equal(t, event.Title, fetchedEvent.Title)
	assert.Equal(t, event.Description, fetchedEvent.Description)

	// Get Events
	events, err := repo.GetEvents()
	assert.NoError(t, err)
	assert.NotEmpty(t, events)

	// Update Event
	event.Title = "Updated Test Event"
	err = repo.UpdateEvent(&event)
	assert.NoError(t, err)

	updatedEvent, err := repo.GetEventByID(event.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Test Event", updatedEvent.Title)

	// Delete Event
	err = repo.DeleteEvent(event.ID)
	assert.NoError(t, err)

	deletedEvent, err := repo.GetEventByID(event.ID)
	assert.Error(t, err)
	assert.Nil(t, deletedEvent)
}
