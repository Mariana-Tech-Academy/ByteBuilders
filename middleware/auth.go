package middleware

import (
	"digital-library/config"
	"digital-library/models"
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenWithBearer := c.GetHeader("Authorization")
		if tokenWithBearer == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			c.Abort()
			return
		}

		// Split token from Bearer
		parts := strings.Split(tokenWithBearer, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		tokenString := parts[1]

		// Check if the token is blacklisted
		var blacklistedToken models.BlacklistedToken
		result := config.DB.Where("token = ?", tokenString).First(&blacklistedToken)
		if result.RowsAffected > 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is blacklisted"})
			c.Abort()
			return
		}

		jwtKey := os.Getenv("JWT_SECRET")

		// Parse the token and claims
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Ensure Role is present
		if claims.Role == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Role not found in token"})
			c.Abort()
			return
		}

		// Store username and role in context
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)

		c.Next()
	}
}

func GenerateJWT(username, role string) (string, error) {
	jwtKey := os.Getenv("JWT_SECRET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"username": username, "role": role, "exp": time.Now().Add(24 * time.Hour).Unix()})
	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", errors.New("Could not create token")
	}

	return tokenString, nil
}
