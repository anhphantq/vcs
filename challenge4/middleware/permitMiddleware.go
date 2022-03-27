package middleware

import (
	"challenge4/models"
	"challenge4/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PermitMiddleware(srv services.UserService, permissionName string, scope string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tmp, _ := c.Get("user")
		user, _ := tmp.(models.Account)

		result, err := srv.CheckPermission(user.Role_id, permissionName, scope)
		if !result || err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Permission denied"})
			return
		}
	}
}
