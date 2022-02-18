package middleware

import (
	"challenge3/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PermitMiddleware(permissionName string, scope string) gin.HandlerFunc {
	return func(c *gin.Context) {
		connection := db.GetDatabase()
		defer db.Closedatabase(connection)

		tmp, _ := c.Get("user")
		user, _ := tmp.(db.Account)

		result := connection.Exec("select * from (select permission_id as ID from rolepermissions where role_id = ?) as permitID, permissions where permitID.ID = permissions.permission_id and permissions.name = ? and permissions.scope = ?", user.Role_id, permissionName, scope)
		if result.RowsAffected < 1 || result.Error != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Permission denied"})
			return
		}
	}
}
