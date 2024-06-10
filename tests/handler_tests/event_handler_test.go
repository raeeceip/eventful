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

func TestCreateEvent(t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()
	body := bytes.NewBufferString(`{"name":"Test Event","description":"This is a test event","start_date":"2023-01-01T00:00:00Z","end_date":"2023-01-02T00:00:00Z"}`)
	req, _ := http.NewRequest("POST", "/events", body)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Test Event")
}

func TestGetEventByID(t *testing.T) {
	router := SetupRouter()

	// Create an event first
	w := httptest.NewRecorder()
	body := bytes.NewBufferString(`{"name":"Test Event","description":"This is a test event","start_date":"2023-01-01T00:00:00Z","end_date":"2023-01-02T00:00:00Z"}`)
	req, _ := http.NewRequest("POST", "/events", body)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// Extract the created event ID
	var response map[string]interface{}
	_ = json.Unmarshal(w.Body.Bytes(), &response)
	eventID := int(response["id"].(float64))

	// Get the event by ID
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/events/"+strconv.Itoa(eventID), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Test Event")
}
