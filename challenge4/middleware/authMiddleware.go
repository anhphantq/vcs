package middleware

import (
	"challenge4/services"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationBearerType = "bearer"
)

func verifyToken(tokenString string) (token *jwt.Token, err error) {
	token, err = jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte("PhanDucAnh"), nil
	})

	return
}

func AuthMiddleware(srv services.UserService) gin.HandlerFunc {
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
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong when geeting claim from JWT"})
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

		c.Set("user", user)

		c.Next()
	}
}
