package auth

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/okta/okta-jwt-verifier-golang"
	"github.com/okta/okta-sdk-golang/okta"
)

func GetVerifier() *okta.JwtVerifier {
	toValidate := map[string]string{"aud": "api://default", "cid": os.Getenv("OKTA_CLIENT_ID")}

	jwtVerifierSetup := okta.JwtVerifier{
		Issuer:           os.Getenv("OKTA_ISSUER"),
		ClaimsToValidate: toValidate,
	}

	return jwtVerifierSetup.New()
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		token = strings.TrimPrefix(token, "Bearer ")
		verifier := GetVerifier()
		_, err := verifier.VerifyAccessToken(context.TODO(), token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}
