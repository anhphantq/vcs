package router

import (
	"challenge2/db"
	"challenge2/middleware"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GeneratePassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func SignUp(c *gin.Context) {
	connection := db.GetDatabase()
	defer db.Closedatabase(connection)

	var user db.Account

	err := c.ShouldBind(&user)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	var dbUser db.Account
	connection.Where("email = ?", user.Email).Find(&dbUser)

	if dbUser.Email != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "This email has been used!"})
		return
	}

	user.Password, err = GeneratePassword(user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Internal Error"})
		return
	}

	result := connection.Exec("INSERT INTO ACCOUNTS VALUES(DEFAULT,?,?,?,?)", user.Username, user.Email, user.Role_id, user.Password)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.String(http.StatusAccepted, "account created")
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateJWT(email string, role uint) (string, error) {
	var mySigningKey = []byte("PhanDucAnh")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Printf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func SignIn(c *gin.Context) {
	connection := db.GetDatabase()
	defer db.Closedatabase(connection)

	var authDetails db.Authentication

	err := c.ShouldBind(&authDetails)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var user db.Account
	connection.Where("email = ?", authDetails.Email).Find(&user)
	if user.Email == "" {
		c.JSON(http.StatusForbidden, gin.H{"msg": "Incorrect Password or Email"})
		return
	}

	check := CheckPasswordHash(authDetails.Password, user.Password)
	if !check {
		c.JSON(http.StatusForbidden, gin.H{"msg": "Incorrect Password or Email"})
		return
	}

	validToken, err := GenerateJWT(user.Email, user.Role_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Can not generate JWT"})
	}

	c.JSON(http.StatusAccepted, gin.H{"msg": "authenticated", "jwt": validToken})
}

func hdGetUser(c *gin.Context) {

}

func hdDeleteUser(c *gin.Context) {

}

func hdPutUser(c *gin.Context) {

}

func hdGetUsers(c *gin.Context) {

}

func hdDeleteUserById(c *gin.Context) {

}

func hdPutUserById(c *gin.Context) {

}

func InitUserRouter(router *gin.RouterGroup) {
	router.POST("/signup", SignUp)
	router.POST("/signin", SignIn)

	router.GET("/user", middleware.AuthMiddleware(), middleware.PermitMiddleware("get", "self"), hdGetUser)
	router.DELETE("/user", middleware.AuthMiddleware(), middleware.PermitMiddleware("delete", "self"), hdDeleteUser)
	router.PUT("/user", middleware.AuthMiddleware(), middleware.PermitMiddleware("put", "self"), hdPutUser)

	router.GET("/users", middleware.AuthMiddleware(), middleware.PermitMiddleware("get", "all"), hdGetUsers)
	router.DELETE("/users/:id", middleware.AuthMiddleware(), middleware.PermitMiddleware("delete", "all"), hdDeleteUserById)
	router.PUT("/user/:id", middleware.AuthMiddleware(), middleware.PermitMiddleware("put", "all"), hdPutUserById)
}
