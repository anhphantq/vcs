package middleware

import (
	"challenge3/db"
	"fmt"
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
			return nil, fmt.Errorf("Invalid signing method")
		}

		return []byte("PhanDucAnh"), nil
	})

	return
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		connection := db.GetDatabase()
		defer db.Closedatabase(connection)

		authorizationHeader := c.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization header is not provided"})
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header"})
			return
		}

		if strings.ToLower(fields[0]) != authorizationBearerType {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization type"})
			return
		}

		token, err := verifyToken(fields[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "some thing went wrong when getting token claim"})
			return
		}

		exp, ok := claims["exp"].(float64)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token format (expire time)"})
			return
		}

		if (int64)(exp) < time.Now().Unix() {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token is expired"})
			return
		}

		email, ok := claims["email"]
		if email, ok = email.(string); !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "can not get email from jwt"})
			return
		}

		var user db.Account
		connection.Where("email = ?", email).Find(&user)

		if email == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "user not found"})
			return
		}

		c.Set("user", user)

		c.Next()
	}
}
