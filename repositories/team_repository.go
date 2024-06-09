package repositories

import (
	"eventful/models"

	"gorm.io/gorm"
)

type TeamRepository struct {
	DB *gorm.DB
}

func NewTeamRepository(db *gorm.DB) *TeamRepository {
	return &TeamRepository{DB: db}
}

func (repo *TeamRepository) GetTeamByID(id uint) (*models.Team, error) {
	var team models.Team
	err := repo.DB.First(&team, id).Error
	if err != nil {
		return nil, err
	}
	return &team, nil
}

func (repo *TeamRepository) GetTeams() ([]models.Team, error) {
	var teams []models.Team
	err := repo.DB.Find(&teams).Error
	if err != nil {
		return nil, err
	}
	return teams, nil
}

func (repo *TeamRepository) CreateTeam(team *models.Team) error {
	return repo.DB.Create(team).Error
}

func (repo *TeamRepository) UpdateTeam(team *models.Team) error {
	return repo.DB.Save(team).Error
}

func (repo *TeamRepository) DeleteTeam(id uint) error {
	return repo.DB.Delete(&models.Team{}, id).Error
}

// GetTeamsByUserID retrieves all teams by user ID
func (repo *TeamRepository) GetTeamsByUserID(userID uint) ([]models.Team, error) {
	var teams []models.Team
	err := repo.DB.Where("user_id = ?", userID).Find(&teams).Error
	if err != nil {
		return nil, err
	}
	return teams, nil
}

// GetTeamsByEventID retrieves all teams by event ID
func (repo *TeamRepository) GetTeamsByEventID(eventID uint) ([]models.Team, error) {

	var teams []models.Team
	err := repo.DB.Where("event_id = ?", eventID).Find(&teams).Error
	if err != nil {
		return nil, err
	}
	return teams, nil
}
