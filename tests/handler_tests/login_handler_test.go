package handler_tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoginHandler(t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()
	body := bytes.NewBufferString(`{"username":"testuser", "password":"testpass"}`)
	req, _ := http.NewRequest("POST", "/login", body)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "token")
}
