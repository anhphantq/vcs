package middleware

import (
	"challenge3/services"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func RoleValidationMiddleware(srv services.UserService, role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Authorization header is not provided"})
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid authorization header"})
			return
		}

		if strings.ToLower(fields[0]) != authorizationBearerType {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid authorization type"})
			return
		}

		token, err := verifyToken(fields[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid JWT"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong when getting token claim"})
			return
		}

		exp, ok := claims["exp"].(float64)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token format (expire time)"})
			return
		}

		if (int64)(exp) < time.Now().Unix() {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Token is expired"})
			return
		}

		email, ok := claims["email"].(string)

		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token format (email)"})
			return
		}

		user, err := srv.GetUserByEmail(email)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "User not found"})
			return
		}

		check, err := srv.CheckRole(user.User_id, role)

		if !check || err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Permission denied"})
			return
		}

		c.Set("user", user)

		c.Next()
	}
}
