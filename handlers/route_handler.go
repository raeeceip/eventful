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

	userRepo := repositories.NewUserRepository(config.DB)
	userHandler := NewUserHandler(userRepo)

	teamRepo := repositories.NewTeamRepository(config.DB)
	teamHandler := NewTeamHandler(teamRepo)

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

	// user routes
	r.POST("/users", userHandler.CreateUser)
	r.GET("/users/:id", userHandler.GetUserByID)
	r.GET("/users", userHandler.GetUsers)
	r.PUT("/users/:id", userHandler.UpdateUser)
	r.DELETE("/users/:id", userHandler.DeleteUser)

	// team routes
	r.POST("/teams", teamHandler.CreateTeam)
	r.GET("/teams/:id", teamHandler.GetTeamByID)
	r.GET("/teams", teamHandler.GetTeams)
	r.PUT("/teams/:id", teamHandler.UpdateTeam)
	r.DELETE("/teams/:id", teamHandler.DeleteTeam)

	return r

}
