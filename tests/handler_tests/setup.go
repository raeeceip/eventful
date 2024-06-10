package handler_tests

import (
	"eventful/auth"
	"eventful/config"
	"eventful/handlers"
	"eventful/repositories"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Add routes for generating tokens and login
	r.POST("/generate-token", auth.TokenGenerationHandler)
	r.POST("/login", handlers.LoginHandler)

	// Protected routes
	r.Use(auth.AuthMiddleware())

	eventRepo := repositories.NewEventRepository(config.DB)
	eventHandler := handlers.NewEventHandler(eventRepo)

	teamRepo := repositories.NewTeamRepository(config.DB)
	teamHandler := handlers.NewTeamHandler(teamRepo)

	userRepo := repositories.NewUserRepository(config.DB)
	userHandler := handlers.NewUserHandler(userRepo)

	// Event routes
	r.POST("/events", eventHandler.CreateEvent)
	r.GET("/events/:id", eventHandler.GetEventByID)
	r.PUT("/events/:id", eventHandler.UpdateEvent)
	r.DELETE("/events/:id", eventHandler.DeleteEvent)

	// Team routes
	r.POST("/teams", teamHandler.CreateTeam)
	r.GET("/teams/:id", teamHandler.GetTeamByID)
	r.PUT("/teams/:id", teamHandler.UpdateTeam)
	r.DELETE("/teams/:id", teamHandler.DeleteTeam)

	// User routes
	r.POST("/users", userHandler.CreateUser)
	r.GET("/users/:id", userHandler.GetUserByID)
	r.PUT("/users/:id", userHandler.UpdateUser)
	r.DELETE("/users/:id", userHandler.DeleteUser)

	return r
}
