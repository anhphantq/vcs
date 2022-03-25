package router

import (
	"challenge3/middleware"
	"challenge3/models"
	"challenge3/services"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var userService services.UserService

func generatePassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func signUp(c *gin.Context) {
	var user models.Account

	err := c.ShouldBind(&user)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"messsage": "Something went wrong in the server"})
		return
	}

	checkEmail, err := userService.CheckEmailUsed(user.Email)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"messsage": "Something went wrong in the server"})
		return
	}

	if checkEmail {
		c.JSON(http.StatusBadRequest, gin.H{"message": "This email has been used!"})
		return
	}

	user.Password, err = generatePassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Can not generate password"})
		return
	}

	user, err = userService.SaveUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server or your role is not in the DB"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"user_id": user.User_id, "email": user.Email, "username": user.Username, "role_id": user.Role_id})
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func generateJWT(email string, role uint) (string, error) {
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

func signIn(c *gin.Context) {
	var authDetails models.Authentication

	err := c.ShouldBind(&authDetails)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"messsage": "Something went wrong in the server"})
		return
	}

	user, err := userService.GetUserByEmail(authDetails.Email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Incorrect Password or Email"})
		return
	}

	check := checkPasswordHash(authDetails.Password, user.Password)
	if !check {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Incorrect Password or Email"})
		return
	}

	validToken, err := generateJWT(user.Email, user.Role_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Can not generate JWT"})
	}

	c.JSON(http.StatusAccepted, gin.H{"jwt": validToken})
}

func hdGetUser(c *gin.Context) {
	tmp, _ := c.Get("user")
	user, _ := tmp.(models.Account)

	c.JSON(http.StatusAccepted, gin.H{"user_id": user.User_id, "username": user.Username, "email": user.Email, "role_id": user.Role_id})
}

func hdDeleteUser(c *gin.Context) {
	tmp, _ := c.Get("user")
	user, _ := tmp.(models.Account)

	err := userService.DeleteUserByID(user.User_id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong when deleting account"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Account deleted"})
}

func hdPutUser(c *gin.Context) {
	tmp, _ := c.Get("user")
	user, _ := tmp.(models.Account)

	var account models.Account

	err := c.ShouldBind(&account)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	if account.Password != "" {
		user.Password = account.Password
	}

	if account.Username != "" {
		user.Username = account.Username
	}

	if account.Password != "" {
		user.Password, err = generatePassword(user.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
			return
		}
	}

	user, err = userService.UpdateUser(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Account updated"})
}

func hdGetUsers(c *gin.Context) {
	var accounts []models.Account

	accounts, err := userService.GetAllUser()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong in the server"})
		return
	}

	var usersInfo []gin.H

	for i := range accounts {
		usersInfo = append(usersInfo, gin.H{"user_id": accounts[i].User_id, "username": accounts[i].Username, "email": accounts[i].Email, "role_id": accounts[i].Role_id})
	}

	c.JSON(http.StatusAccepted, usersInfo)
}

func hdGetUsersByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	user, err := userService.GetUserByID(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server or wrong user's ID"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"user_id": user.User_id, "username": user.Username, "email": user.Email, "role_id": user.Role_id})
}

func hdDeleteUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	err = userService.DeleteUserByID(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server or wrong user's ID"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Account deleted"})
}

func hdPutUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	user, err := userService.GetUserByID(uint(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Wrong user's ID"})
		return
	}

	var account models.Account

	if err := c.ShouldBind(&account); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	if account.Password != "" {
		user.Password = account.Password
	}

	if account.Username != "" {
		user.Username = account.Username
	}

	if account.Password != "" {
		user.Password, err = generatePassword(user.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
			return
		}
	}

	user, err = userService.UpdateUser(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Info updated"})
}

func initUserRouter(router *gin.RouterGroup, srv services.UserService) {
	userService = srv
	router.POST("/signup", middleware.ValidationMiddleware(templateRouter), signUp)
	router.POST("/signin", middleware.ValidationMiddleware(templateRouter), signIn)

	router.GET("/user", middleware.ValidationMiddleware(templateRouter), middleware.AuthMiddleware(srv), middleware.PermitMiddleware(srv, "get", "self"), hdGetUser)
	router.DELETE("/user", middleware.ValidationMiddleware(templateRouter), middleware.AuthMiddleware(srv), middleware.PermitMiddleware(srv, "delete", "self"), hdDeleteUser)
	router.PUT("/user", middleware.ValidationMiddleware(templateRouter), middleware.AuthMiddleware(srv), middleware.PermitMiddleware(srv, "put", "self"), hdPutUser)

	router.GET("/users", middleware.ValidationMiddleware(templateRouter), middleware.AuthMiddleware(srv), middleware.PermitMiddleware(srv, "get", "all"), hdGetUsers)
	router.GET("/users/:id", middleware.ValidationMiddleware(templateRouter), middleware.AuthMiddleware(srv), middleware.PermitMiddleware(srv, "get", "all"), hdGetUsersByID)
	router.DELETE("/users/:id", middleware.ValidationMiddleware(templateRouter), middleware.AuthMiddleware(srv), middleware.PermitMiddleware(srv, "delete", "all"), hdDeleteUserById)
	router.PUT("/users/:id", middleware.ValidationMiddleware(templateRouter), middleware.AuthMiddleware(srv), middleware.PermitMiddleware(srv, "put", "all"), hdPutUserById)
}
