package repository_tests

import (
	"eventful/models"
	"eventful/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTeamRepository(t *testing.T) {
	db := SetupTestDB()
	repo := repositories.NewTeamRepository(db)

	team := models.Team{
		Name:        "Dev Team",
		Description: "Development Team",
	}

	// Create Team
	err := repo.CreateTeam(&team)
	assert.NoError(t, err)
	assert.NotZero(t, team.ID)

	// Get Team by ID
	fetchedTeam, err := repo.GetTeamByID(team.ID)
	assert.NoError(t, err)
	assert.Equal(t, team.Name, fetchedTeam.Name)
	assert.Equal(t, team.Description, fetchedTeam.Description)

	// Get Teams
	teams, err := repo.GetTeams()
	assert.NoError(t, err)
	assert.NotEmpty(t, teams)

	// Update Team
	team.Name = "Updated Dev Team"
	err = repo.UpdateTeam(&team)
	assert.NoError(t, err)

	updatedTeam, err := repo.GetTeamByID(team.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Dev Team", updatedTeam.Name)

	// Delete Team
	err = repo.DeleteTeam(team.ID)
	assert.NoError(t, err)

	deletedTeam, err := repo.GetTeamByID(team.ID)
	assert.Error(t, err)
	assert.Nil(t, deletedTeam)
}
