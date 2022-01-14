package middleware

import (
	"challenge2/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PermitMiddleware(permissionName string, scope string) gin.HandlerFunc {
	return func(c *gin.Context) {
		connection := db.GetDatabase()
		defer db.Closedatabase(connection)

		email, ok := c.Get("email")
		if email, ok = email.(string); !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "can not get email from jwt"})
			return
		}

		var user db.Account
		connection.Where("email = ?", email).Find(&user)

		connection.Exec("select from (select permission_id from rolespermissions where role_id == ?) as permitID, ")

	}
}
