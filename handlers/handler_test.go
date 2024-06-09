package handlers

//write test for handler functions

import (
	"bytes"
	"encoding/json"
	"eventful/eventtypes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetEvents(t *testing.T) {
	db.InitDB()
	defer db.CloseDB()

	req, err := http.NewRequest("GET", "/api/events", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetEvents)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetEvent(t *testing.T) {
	db.InitDB()
	defer db.CloseDB()

	req, err := http.NewRequest("GET", "/api/events/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetEvent)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestCreateEvent(t *testing.T) {
	db.InitDB()
	defer db.CloseDB()

	event := eventtypes.Event{
		Name:     "Test Event",
		Location: "Test Location",
		Date:     "2020-01-01",
	}

	body, err := json.Marshal(event)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/api/events", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateEvent)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestUpdateEvent(t *testing.T) {
	db.InitDB()
	defer db.CloseDB()

	event := eventtypes.Event{
		ID:       1,
		Name:     "Test Event",
		Location: "Test Location",
		Date:     "2020 - 01 - 01",
	}

	body, err := json.Marshal(event)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PUT", "/api/events/1", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateEvent)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestDeleteEvent(t *testing.T) {
	db.InitDB()
	defer db.CloseDB()

	req, err := http.NewRequest("DELETE", "/api/events/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteEvent)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetEventsHandler(t *testing.T) {
	db.InitDB()
	defer db.CloseDB()

	req, err := http.NewRequest("GET", "/api/events", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/api/events", GetEvents).Methods("GET")

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetEventHandler(t *testing.T) {
	db.InitDB()
	defer db.CloseDB()

	req, err := http.NewRequest("GET", "/api/events/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/api/events/{id}", GetEvent).Methods("GET")

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestCreateEventHandler(t *testing.T) {
	db.InitDB()
	defer db.CloseDB()

	event := eventtypes.Event{
		Name:     "Test Event",
		Location: "Test Location",
		Date:     "2020-01-01",
	}

	body, err := json.Marshal(event)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/api/events", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/api/events", CreateEvent).Methods("POST")

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}
