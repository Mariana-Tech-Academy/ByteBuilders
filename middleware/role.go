package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoleMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "Role not found in token"})
			c.Abort()
			return
		}

		userRole, ok := role.(string)
		if !ok || userRole != requiredRole {
			c.JSON(http.StatusForbidden, gin.H{"error": "You do not have the required permissions"})
			c.Abort()
			return
		}

		c.Next()
	}
}
