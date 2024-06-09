package repositories

import (
	"eventful/models"

	"gorm.io/gorm"
)

type EventRepository struct {
	DB *gorm.DB
}

func NewEventRepository(db *gorm.DB) *EventRepository {
	return &EventRepository{DB: db}
}

func (repo *EventRepository) GetEventByID(id uint) (*models.Event, error) {
	var event models.Event
	err := repo.DB.First(&event, id).Error
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (repo *EventRepository) GetEvents() ([]models.Event, error) {
	var events []models.Event
	err := repo.DB.Find(&events).Error
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (repo *EventRepository) CreateEvent(event *models.Event) error {
	return repo.DB.Create(event).Error
}

func (repo *EventRepository) UpdateEvent(event *models.Event) error {
	return repo.DB.Save(event).Error
}

func (repo *EventRepository) DeleteEvent(id uint) error {
	return repo.DB.Delete(&models.Event{}, id).Error
}
