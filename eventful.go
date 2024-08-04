package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func main() {
	// Connect to Redis
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	ctx := context.Background()

	// Ping Redis to check connection
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Failed to connect to Redis: ", err)
	}

	repo := NewRepository(rdb)
	handler := NewHandler(repo)

	r := gin.Default()

	// Load HTML templates
	r.LoadHTMLGlob("templates/*")

	// Serve static files
	r.Static("/static", "./static")

	// Home page
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Eventful",
		})
	})

	// API routes
	api := r.Group("/api")
	{
		// Event routes
		api.POST("/events", handler.CreateEvent)
		api.GET("/events/:id", handler.GetEvent)
		api.PUT("/events/:id", handler.UpdateEvent)
		api.DELETE("/events/:id", handler.DeleteEvent)
		api.GET("/events", handler.GetAllEvents)

		// User routes
		api.POST("/users", handler.CreateUser)
		api.GET("/users/:id", handler.GetUser)
		api.PUT("/users/:id", handler.UpdateUser)
		api.DELETE("/users/:id", handler.DeleteUser)
		api.GET("/users", handler.GetAllUsers)

		// Admin routes
		admin := api.Group("/admin")
		{
			admin.POST("/roles", handler.CreateRole)
			admin.PUT("/users/role", handler.AssignUserRole)
		}
	}

	r.Run(":8080")
}
