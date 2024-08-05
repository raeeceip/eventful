package main

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Location    string    `json:"location"`
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"` // Password won't be exposed in JSON
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type Role struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsAdmin     bool   `json:"is_admin"`
}

func NewEvent() *Event {
	return &Event{ID: uuid.New().String()}
}

func NewUser() *User {
	return &User{ID: uuid.New().String()}
}

func NewRole() *Role {
	return &Role{ID: uuid.New().String()}
}
