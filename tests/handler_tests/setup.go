package handler_tests

import (
	"eventful/auth"
	"eventful/handlers"
	"eventful/models"
	"eventful/repositories"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to test database: " + err.Error())
	}

	err = db.AutoMigrate(&models.Event{}, &models.User{}, &models.Team{}, &models.Role{})
	if err != nil {
		panic("failed to migrate test database: " + err.Error())
	}

	return db
}

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	// Use custom recovery middleware to handle panics
	r.Use(gin.Recovery())

	// Add routes for generating tokens and login
	r.POST("/generate-token", auth.TokenGenerationHandler)
	r.POST("/login", handlers.LoginHandler)

	// Protected routes
	r.Use(auth.AuthMiddleware())

	// Setup database
	db := SetupTestDB()

	// Initialize repositories
	eventRepo := repositories.NewEventRepository(db)
	teamRepo := repositories.NewTeamRepository(db)
	userRepo := repositories.NewUserRepository(db)
	roleRepo := repositories.NewRoleRepository(db)

	// Initialize handlers
	eventHandler := handlers.NewEventHandler(eventRepo)
	teamHandler := handlers.NewTeamHandler(teamRepo)
	userHandler := handlers.NewUserHandler(userRepo)
	roleHandler := handlers.NewRoleHandler(roleRepo)

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

	// Role routes
	r.POST("/roles", roleHandler.CreateRole)
	r.GET("/roles/:id", roleHandler.GetRoleByID)
	r.GET("/roles", roleHandler.GetRoles)
	r.PUT("/roles/:id", roleHandler.UpdateRole)
	r.DELETE("/roles/:id", roleHandler.DeleteRole)

	return r
}

func TestMain(m *testing.M) {
	// Set the JWT secret for testing
	os.Setenv("JWT_SECRET", "testsecret")

	// Run the tests
	code := m.Run()

	// Exit with the test result code
	os.Exit(code)
}
