package repositories

import (
	"eventful/models"

	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (repo *UserRepo) GetUserByID(id int) (*models.User, error) {
	var user models.User
	err := repo.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepo) GetUsers() ([]models.User, error) {
	var users []models.User
	err := repo.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepo) CreateUser(user *models.User) error {
	return repo.DB.Create(user).Error
}

func (repo *UserRepo) UpdateUser(user *models.User) error {
	return repo.DB.Save(user).Error
}

func (repo *UserRepo) DeleteUser(id int) error {
	return repo.DB.Delete(&models.User{}, id).Error
}
