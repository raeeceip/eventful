package handler_tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTeam(t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()
	body := bytes.NewBufferString(`{"name":"Dev Team","description":"Development Team"}`)
	req, _ := http.NewRequest("POST", "/teams", body)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Dev Team")
}

func TestGetTeamByID(t *testing.T) {
	router := SetupRouter()

	// Create a team first
	w := httptest.NewRecorder()
	body := bytes.NewBufferString(`{"name":"Dev Team","description":"Development Team"}`)
	req, _ := http.NewRequest("POST", "/teams", body)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// Extract the created team ID
	var response map[string]interface{}
	_ = json.Unmarshal(w.Body.Bytes(), &response)
	teamID := int(response["id"].(float64))

	// Get the team by ID
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/teams/"+strconv.Itoa(teamID), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Dev Team")
}
