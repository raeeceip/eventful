package handlers

//write test for handler functions

import (
	"bytes"
	"encoding/json"
	"eventful/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRole(t *testing.T) {
	role := models.Role{
		Name: "Admin",
	}
	roleJSON, _ := json.Marshal(role)
	req, err := http.NewRequest("POST", "/roles", bytes.NewBuffer(roleJSON))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.CreateRole)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetRoleByID(t *testing.T) {
	req, err := http.NewRequest("GET", "/roles/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetRoleByID)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetRoles(t *testing.T) {
	req, err := http.NewRequest("GET", "/roles", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetRoles)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestUpdateRole(t *testing.T) {
	role := models.Role{
		Name: "Admin",
	}
	roleJSON, _ := json.Marshal(role)
	req, err := http.NewRequest("PUT", "/roles/1", bytes.NewBuffer(roleJSON))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.UpdateRole)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestDeleteRole(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/roles/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.DeleteRole)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestCreateTeam(t *testing.T) {

	team := models.Team{
		Name: "Team 1",
	}
	teamJSON, _ := json.Marshal(team)
	req, err := http.NewRequest("POST", "/teams", bytes.NewBuffer(teamJSON))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateTeam)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetTeamByID(t *testing.T) {
	req, err := http.NewRequest("GET", "/teams/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetTeamByID)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetTeams(t *testing.T) {
	req, err := http.NewRequest("GET", "/teams", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetTeams)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestUpdateTeam(t *testing.T) {
	team := models.Team{
		Name: "Team 1",
	}
	teamJSON, _ := json.Marshal(team)
	req, err := http.NewRequest("PUT", "/teams/1", bytes.NewBuffer(teamJSON))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.UpdateTeam)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestDeleteTeam(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/teams/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.DeleteTeam)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestCreateEvent(t *testing.T) {
	event := models.Event{
		Name: "Event 1",
	}
	eventJSON, _ := json.Marshal(event)
	req, err := http.NewRequest("POST", "/events", bytes.NewBuffer(eventJSON))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.CreateEvent)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetEventByID(t *testing.T) {
	req, err := http.NewRequest("GET", "/events/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetEventByID)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetEvents(t *testing.T) {
	req, err := http.NewRequest("GET", "/events", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetEvents)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestUpdateEvent(t *testing.T) {
	event := models.Event{
		Name: "Event 1",
	}
	eventJSON, _ := json.Marshal(event)
	req, err := http.NewRequest("PUT", "/events/1", bytes.NewBuffer(eventJSON))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.UpdateEvent)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestDeleteEvent(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/events/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hanlers.DeleteEvent)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestCreateUser(t *testing.T) {
	user := models.User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "test@gmail.com",
	}
	userJSON, _ := json.Marshal(user)
	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(userJSON))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.CreateUser)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetUserByID(t *testing.T) {
	req, err := http.NewRequest("GET", "/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetUserByID)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetUsers(t *testing.T) {
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetUsers)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestUpdateUser(t *testing.T) {
	user := models.User{
		Username: "johndoe",
		Password: "password",
		Email:    "test@gmail.com",
		Role:     "Admin",
		ID:       1,
	}

	userJSON, _ := json.Marshal(user)
	req, err := http.NewRequest("PUT", "/users/1", bytes.NewBuffer(userJSON))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.UpdateUser)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestDeleteUser(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.DeleteUser)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}
