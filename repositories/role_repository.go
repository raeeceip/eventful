package repositories

import (
	"eventful/models"

	"gorm.io/gorm"
)

type RoleRepository struct {
	DB *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{DB: db}
}

func (repo *RoleRepository) GetRoleByID(id uint) (*models.Role, error) {
	var role models.Role
	err := repo.DB.First(&role, id).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (repo *RoleRepository) GetRoles() ([]models.Role, error) {
	var roles []models.Role
	err := repo.DB.Find(&roles).Error
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (repo *RoleRepository) CreateRole(role *models.Role) error {
	return repo.DB.Create(role).Error
}

func (repo *RoleRepository) UpdateRole(role *models.Role) error {
	return repo.DB.Save(role).Error
}

func (repo *RoleRepository) DeleteRole(id uint) error {
	return repo.DB.Delete(&models.Role{}, id).Error
}
