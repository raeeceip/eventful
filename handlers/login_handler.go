package handlers

import (
	"eventful/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Dummy user authentication
var dummyUser = map[string]string{
	"username": "password",
}

// LoginHandler handles user login and generates JWT token
func LoginHandler(c *gin.Context) {
	var credentials map[string]string
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	password, ok := dummyUser[credentials["username"]]
	if !ok || password != credentials["password"] {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	token, err := auth.GenerateToken(credentials["username"])
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
