package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	Repo *Repository
}

func NewHandler(repo *Repository) *Handler {
	return &Handler{Repo: repo}
}

func (h *Handler) CreateEvent(c *gin.Context) {
	var event Event
	if err := c.ShouldBind(&event); err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"error": err.Error()})
		return
	}
	event.ID = NewEvent().ID
	if err := h.Repo.CreateEvent(c, &event); err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "event.html", event)
}

func (h *Handler) GetEvent(c *gin.Context) {
	id := c.Param("id")
	var event Event
	if err := h.Repo.Get(c, "event:"+id, &event); err != nil {
		c.HTML(http.StatusNotFound, "error.html", gin.H{"error": "Event not found"})
		return
	}
	c.HTML(http.StatusOK, "event.html", event)
}

func (h *Handler) UpdateEvent(c *gin.Context) {
	id := c.Param("id")
	var event Event
	if err := c.ShouldBind(&event); err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"error": err.Error()})
		return
	}
	event.ID = id
	if err := h.Repo.Update(c, "event:"+id, &event); err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "event.html", event)
}

func (h *Handler) DeleteEvent(c *gin.Context) {
	id := c.Param("id")
	if err := h.Repo.Delete(c, "event:"+id); err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "")
}

func (h *Handler) GetAllEvents(c *gin.Context) {
	events, err := h.Repo.GetAll(c, "event:")
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": err.Error()})
		return
	}
	var eventsList []Event
	for _, eventJSON := range events {
		var event Event
		if err := json.Unmarshal([]byte(eventJSON), &event); err != nil {
			c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": "Failed to parse event data"})
			return
		}
		eventsList = append(eventsList, event)
	}
	c.HTML(http.StatusOK, "events.html", gin.H{"events": eventsList})
}

func (h *Handler) CreateUser(c *gin.Context) {
	var user User
	if err := c.ShouldBind(&user); err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"error": err.Error()})
		return
	}
	user.ID = NewUser().ID
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)
	if err := h.Repo.CreateUser(c, &user); err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "user.html", user)
}

func (h *Handler) GetUser(c *gin.Context) {
	id := c.Param("id")
	var user User
	if err := h.Repo.Get(c, "user:"+id, &user); err != nil {
		c.HTML(http.StatusNotFound, "error.html", gin.H{"error": "User not found"})
		return
	}
	c.HTML(http.StatusOK, "user.html", user)
}

func (h *Handler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user User
	if err := c.ShouldBind(&user); err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"error": err.Error()})
		return
	}
	user.ID = id
	if err := h.Repo.Update(c, "user:"+id, &user); err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "user.html", user)
}

func (h *Handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := h.Repo.Delete(c, "user:"+id); err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "")
}

func (h *Handler) GetAllUsers(c *gin.Context) {
	users, err := h.Repo.GetAll(c, "user:")
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": err.Error()})
		return
	}
	var usersList []User
	for _, userJSON := range users {
		var user User
		if err := json.Unmarshal([]byte(userJSON), &user); err != nil {
			c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": "Failed to parse user data"})
			return
		}
		usersList = append(usersList, user)
	}
	c.HTML(http.StatusOK, "users.html", gin.H{"users": usersList})
}

func (h *Handler) CreateRole(c *gin.Context) {
	var role Role
	if err := c.ShouldBind(&role); err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"error": err.Error()})
		return
	}
	role.ID = NewRole().ID
	if err := h.Repo.CreateRole(c, &role); err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "role.html", role)
}

func (h *Handler) AssignUserRole(c *gin.Context) {
	var request struct {
		UserID string `form:"user_id"`
		RoleID string `form:"role_id"`
	}
	if err := c.ShouldBind(&request); err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"error": err.Error()})
		return
	}

	var user User
	if err := h.Repo.Get(c, "user:"+request.UserID, &user); err != nil {
		c.HTML(http.StatusNotFound, "error.html", gin.H{"error": "User not found"})
		return
	}

	var role Role
	if err := h.Repo.Get(c, "role:"+request.RoleID, &role); err != nil {
		c.HTML(http.StatusNotFound, "error.html", gin.H{"error": "Role not found"})
		return
	}

	user.Role = role.Name
	if err := h.Repo.Update(c, "user:"+user.ID, &user); err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "user.html", user)
}

// Add this to your handler.go file

func (h *Handler) Login(c *gin.Context) {
	var loginData struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Here you would typically check the username and password against your database
	// For this example, we'll just check if the username and password are not empty
	if loginData.Username != "" && loginData.Password != "" {
		// In a real application, you would generate a proper JWT token here
		token := "sample_token"
		c.JSON(http.StatusOK, gin.H{"token": token})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}
}
