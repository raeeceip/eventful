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

func TestCreateUser(t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()
	body := bytes.NewBufferString(`{"username":"testuser","password":"testpass","email":"testuser@example.com","role":"user"}`)
	req, _ := http.NewRequest("POST", "/users", body)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "testuser")
}

func TestGetUserByID(t *testing.T) {
	router := SetupRouter()

	// Create a user first
	w := httptest.NewRecorder()
	body := bytes.NewBufferString(`{"username":"testuser","password":"testpass","email":"testuser@example.com","role":"user"}`)
	req, _ := http.NewRequest("POST", "/users", body)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// Extract the created user ID
	var response map[string]interface{}
	_ = json.Unmarshal(w.Body.Bytes(), &response)
	userID := int(response["id"].(float64))

	// Get the user by ID
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/users/"+strconv.Itoa(userID), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "testuser")
}
