package middleware

import (
	"net/http"
	"ticket-api/utils"

	"github.com/gin-gonic/gin"
)

func RoleMiddleware(requiredRole string) gin.HandlerFunc {
		return func(c *gin.Context) {
		userRole, err := utils.ExtractRoleFromJWT(c)
		if err != nil || userRole != requiredRole {
			c.JSON(http.StatusForbidden, gin.H{"error": "forbidden: access denied"})
			c.Abort()
			return
		}
		c.Next()
	}
}