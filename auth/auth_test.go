package auth

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func init() {
	os.Setenv("JWT_SECRET", "testsecret")
	jwtSecret = []byte(os.Getenv("JWT_SECRET"))
}

func TestGenerateToken(t *testing.T) {
	token, err := GenerateToken("testuser")
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// Parse token to verify claims
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return jwtSecret, nil
	})
	assert.NoError(t, err)
	assert.NotNil(t, parsedToken)
	assert.True(t, parsedToken.Valid)

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	assert.True(t, ok)
	assert.Equal(t, "testuser", claims["username"])
	assert.Equal(t, true, claims["authorized"])
	assert.WithinDuration(t, time.Now().Add(24*time.Hour), time.Unix(int64(claims["exp"].(float64)), 0), time.Minute)
}

func TestAuthMiddleware(t *testing.T) {
	token, _ := GenerateToken("testuser")

	tests := []struct {
		name       string
		token      string
		wantStatus int
	}{
		{
			name:       "No Token",
			token:      "",
			wantStatus: http.StatusUnauthorized,
		},
		{
			name:       "Invalid Token",
			token:      "invalidtoken",
			wantStatus: http.StatusUnauthorized,
		},
		{
			name:       "Valid Token",
			token:      token,
			wantStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			r := gin.New()

			r.Use(AuthMiddleware())
			r.GET("/protected", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"status": "success"})
			})

			req, _ := http.NewRequest(http.MethodGet, "/protected", nil)
			if tt.token != "" {
				req.Header.Set("Authorization", "Bearer "+tt.token)
			}
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
		})
	}
}

func TestAuthMiddleware_InvalidSigningMethod(t *testing.T) {
	// Generate token with different signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS384, jwt.MapClaims{
		"authorized": true,
		"username":   "testuser",
		"exp":        time.Now().Add(24 * time.Hour).Unix(),
	})
	tokenString, _ := token.SignedString(jwtSecret)

	gin.SetMode(gin.TestMode)
	r := gin.New()

	r.Use(AuthMiddleware())
	r.GET("/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "success"})
	})

	req, _ := http.NewRequest(http.MethodGet, "/protected", nil)
	req.Header.Set("Authorization", "Bearer "+tokenString)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
