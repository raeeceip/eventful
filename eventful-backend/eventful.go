package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func main() {
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		redisURL = "redis://localhost:6379"
	}

	opts, err := redis.ParseURL(redisURL)
	if err != nil {
		log.Fatal("Failed to parse Redis URL: ", err)
	}

	rdb := redis.NewClient(opts)

	ctx := context.Background()

	// Ping Redis to check connection
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Failed to connect to Redis: ", err)
	}

	repo := NewRepository(rdb)
	handler := NewHandler(repo)

	r := gin.Default()
	// add cors
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"} // Allow requests from Electron app
	r.Use(cors.New(config))
	// Serve static files
	r.Static("/static", "./static")

	// Load HTML templates
	r.LoadHTMLGlob("templates/*")

	// Home page
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Eventful",
		})
	})

	// API routes
	api := r.Group("/api")
	{
		api.GET("/events", handler.GetAllEvents)
		api.POST("/events", handler.CreateEvent)
		api.GET("/events/:id", handler.GetEvent)
		api.PUT("/events/:id", handler.UpdateEvent)
		api.DELETE("/events/:id", handler.DeleteEvent)

		api.GET("/users", handler.GetAllUsers)
		api.POST("/users", handler.CreateUser)
		api.GET("/users/:id", handler.GetUser)
		api.PUT("/users/:id", handler.UpdateUser)
		api.DELETE("/users/:id", handler.DeleteUser)

		api.POST("/login", handler.Login)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on port %s", port)
	r.Run(":" + port)
}
