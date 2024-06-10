package repositories

import (
	"eventful/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (repo *UserRepository) CreateUser(user *models.User) error {
	return repo.DB.Create(user).Error
}

func (repo *UserRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := repo.DB.First(&user, id).Error
	return &user, err
}

func (repo *UserRepository) GetUsers() ([]models.User, error) {
	var users []models.User
	err := repo.DB.Find(&users).Error
	return users, err
}

func (repo *UserRepository) UpdateUser(user *models.User) error {
	return repo.DB.Save(user).Error
}

func (repo *UserRepository) DeleteUser(id uint) error {
	return repo.DB.Delete(&models.User{}, id).Error
}
