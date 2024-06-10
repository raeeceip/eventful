package repository_tests

import (
	"eventful/models"
	"eventful/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoleRepository(t *testing.T) {
	db := SetupTestDB()
	repo := repositories.NewRoleRepository(db)

	role := models.Role{
		Name:   "Admin",
		Leader: true,
	}

	// Create Role
	err := repo.CreateRole(&role)
	assert.NoError(t, err)
	assert.NotZero(t, role.ID)

	// Get Role by ID
	fetchedRole, err := repo.GetRoleByID(role.ID)
	assert.NoError(t, err)
	assert.Equal(t, role.Name, fetchedRole.Name)
	assert.Equal(t, role.Leader, fetchedRole.Leader)

	// Get Roles
	roles, err := repo.GetRoles()
	assert.NoError(t, err)
	assert.NotEmpty(t, roles)

	// Update Role
	role.Name = "Updated Admin"
	err = repo.UpdateRole(&role)
	assert.NoError(t, err)

	updatedRole, err := repo.GetRoleByID(role.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Admin", updatedRole.Name)

	// Delete Role
	err = repo.DeleteRole(role.ID)
	assert.NoError(t, err)

	deletedRole, err := repo.GetRoleByID(role.ID)
	assert.Error(t, err)
	assert.Nil(t, deletedRole)
}
