package handlers

import (
	"eventful/auth"
	"eventful/config"
	"eventful/repositories"

	"github.com/gin-gonic/gin"
)

// NewRouter sets up the Gin router and routes
func NewRouter() *gin.Engine {
	r := gin.Default()

	// Route for generating tokens
	r.POST("/generate-token", auth.TokenGenerationHandler)

	// Route for login
	r.POST("/login", LoginHandler)

	// Protected routes
	r.Use(auth.AuthMiddleware())

	eventRepo := repositories.NewEventRepository(config.DB)
	eventHandler := NewEventHandler(eventRepo)

	roleRepo := repositories.NewRoleRepository(config.DB)
	roleHandler := NewRoleHandler(roleRepo)

	// Event routes
	r.POST("/events", eventHandler.CreateEvent)
	r.GET("/events/:id", eventHandler.GetEventByID)
	r.PUT("/events/:id", eventHandler.UpdateEvent)
	r.DELETE("/events/:id", eventHandler.DeleteEvent)

	// Role routes
	r.POST("/roles", roleHandler.CreateRole)
	r.GET("/roles/:id", roleHandler.GetRoleByID)
	r.GET("/roles", roleHandler.GetRoles)
	r.PUT("/roles/:id", roleHandler.UpdateRole)
	r.DELETE("/roles/:id", roleHandler.DeleteRole)

	// Add more routes for users, teams, etc.

	return r
}
