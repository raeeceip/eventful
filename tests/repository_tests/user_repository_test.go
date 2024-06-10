package repository_tests

import (
	"eventful/models"
	"eventful/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository(t *testing.T) {
	db := SetupTestDB()
	repo := repositories.NewUserRepository(db)

	user := models.User{
		Username: "testuser",
		Password: "testpass",
		Email:    "testuser@example.com",
		Role:     "user",
	}

	// Create User
	err := repo.CreateUser(&user)
	assert.NoError(t, err)
	assert.NotZero(t, user.ID)

	// Get User by ID
	fetchedUser, err := repo.GetUserByID(user.ID)
	assert.NoError(t, err)
	assert.Equal(t, user.Username, fetchedUser.Username)
	assert.Equal(t, user.Email, fetchedUser.Email)
	assert.Equal(t, user.Role, fetchedUser.Role)

	// Get Users
	users, err := repo.GetUsers()
	assert.NoError(t, err)
	assert.NotEmpty(t, users)

	// Update User
	user.Username = "updateduser"
	err = repo.UpdateUser(&user)
	assert.NoError(t, err)

	updatedUser, err := repo.GetUserByID(user.ID)
	assert.NoError(t, err)
	assert.Equal(t, "updateduser", updatedUser.Username)

	// Delete User
	err = repo.DeleteUser(user.ID)
	assert.NoError(t, err)

	deletedUser, err := repo.GetUserByID(user.ID)
	assert.Error(t, err)
	assert.Nil(t, deletedUser)
}
