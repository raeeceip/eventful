package handler_tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"eventful/auth"

	"github.com/stretchr/testify/assert"
)

func TestCreateRole(t *testing.T) {
	router := SetupRouter()

	token, _ := auth.GenerateToken("testuser")

	w := httptest.NewRecorder()
	body := bytes.NewBufferString(`{"name":"Admin","leader":true}`)
	req, _ := http.NewRequest("POST", "/roles", body)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Admin")
}

func TestGetRoleByID(t *testing.T) {
	router := SetupRouter()

	token, _ := auth.GenerateToken("testuser")

	// Create a role first
	w := httptest.NewRecorder()
	body := bytes.NewBufferString(`{"name":"Admin","leader":true}`)
	req, _ := http.NewRequest("POST", "/roles", body)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// Extract the created role ID
	var response map[string]interface{}
	_ = json.Unmarshal(w.Body.Bytes(), &response)
	roleID := int(response["id"].(float64))

	// Get the role by ID
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/roles/"+strconv.Itoa(roleID), nil)
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Admin")
}
