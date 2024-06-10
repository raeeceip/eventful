package handler_tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"eventful/auth"

	"github.com/stretchr/testify/assert"
)

func TestRoutes(t *testing.T) {
	router := SetupRouter()

	t.Run("NonExistentRoute", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/non-existent-route", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("ValidRouteNoAuth", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/events", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})
}

func TestValidRoutesWithAuth(t *testing.T) {
	router := SetupRouter()

	token, _ := auth.GenerateToken("testuser")

	t.Run("CreateEvent", func(t *testing.T) {
		w := httptest.NewRecorder()
		body := bytes.NewBufferString(`{"name":"Test Event","description":"This is a test event","start_date":"2023-01-01T00:00:00Z","end_date":"2023-01-02T00:00:00Z"}`)
		req, _ := http.NewRequest("POST", "/events", body)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("CreateTeam", func(t *testing.T) {
		w := httptest.NewRecorder()
		body := bytes.NewBufferString(`{"name":"Dev Team","description":"Development Team"}`)
		req, _ := http.NewRequest("POST", "/teams", body)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("CreateUser", func(t *testing.T) {
		w := httptest.NewRecorder()
		body := bytes.NewBufferString(`{"username":"testuser","password":"testpass","email":"testuser@example.com","role":"user"}`)
		req, _ := http.NewRequest("POST", "/users", body)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("CreateRole", func(t *testing.T) {
		w := httptest.NewRecorder()
		body := bytes.NewBufferString(`{"name":"Admin","leader":true}`)
		req, _ := http.NewRequest("POST", "/roles", body)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})
}
