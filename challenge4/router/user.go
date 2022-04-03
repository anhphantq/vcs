package router

import (
	"challenge4/middleware"
	"challenge4/models"
	"challenge4/services"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var UserService services.UserService

func generatePassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func SignUp(c *gin.Context) {
	var user models.Account

	err := c.ShouldBind(&user)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"messsage": "Something went wrong in the server"})
		return
	}

	checkEmail, err := UserService.CheckEmailUsed(user.Email)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"messsage": "Something went wrong in the server"})
		return
	}

	if checkEmail {
		c.JSON(http.StatusBadRequest, gin.H{"message": "This email has been used!"})
		return
	}

	user, err = UserService.SaveUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server or your role is not in the DB"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"user_id": user.User_id, "email": user.Email, "username": user.Username, "role_id": user.Role_id})
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
	var authDetails models.Authentication

	err := c.ShouldBind(&authDetails)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"messsage": "Something went wrong in the server"})
		return
	}

	user, err := UserService.GetUserByEmail(authDetails.Email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Incorrect Password or Email"})
		return
	}

	check := CheckPasswordHash(authDetails.Password, user.Password)
	if !check {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Incorrect Password or Email"})
		return
	}

	validToken, err := GenerateJWT(user.Email, user.Role_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Can not generate JWT"})
	}

	c.JSON(http.StatusAccepted, gin.H{"jwt": validToken})
}

func HdGetUser(c *gin.Context) {
	tmp, _ := c.Get("user")
	user, _ := tmp.(models.Account)

	c.JSON(http.StatusAccepted, gin.H{"user_id": user.User_id, "username": user.Username, "email": user.Email, "role_id": user.Role_id})
}

func HdDeleteUser(c *gin.Context) {
	tmp, _ := c.Get("user")
	user, _ := tmp.(models.Account)

	err := UserService.DeleteUserByID(user.User_id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong when deleting account"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Account deleted"})
}

func HdPutUser(c *gin.Context) {
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

	user, err = UserService.UpdateUserByID(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Account updated"})
}

func HdGetUsers(c *gin.Context) {
	var accounts []models.Account

	accounts, err := UserService.GetAllUser()

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

func HdGetUsersByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	user, err := UserService.GetUserByID(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server or wrong user's ID"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"user_id": user.User_id, "username": user.Username, "email": user.Email, "role_id": user.Role_id})
}

func HdDeleteUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	err = UserService.DeleteUserByID(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Account deleted"})
}

func HdPutUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	user, err := UserService.GetUserByID(uint(id))

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

	user, err = UserService.UpdateUserByID(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Info updated"})
}

func InitUserRouter(router *gin.RouterGroup, srv services.UserService) {
	UserService = srv
	router.POST("/signup", middleware.ValidationMiddleware(TemplateRouter), SignUp)
	router.POST("/signin", middleware.ValidationMiddleware(TemplateRouter), SignIn)

	router.GET("/user", middleware.ValidationMiddleware(TemplateRouter), middleware.AuthMiddleware(srv), middleware.PermitMiddleware(srv, "get", "self"), HdGetUser)
	router.DELETE("/user", middleware.ValidationMiddleware(TemplateRouter), middleware.AuthMiddleware(srv), middleware.PermitMiddleware(srv, "delete", "self"), HdDeleteUser)
	router.PUT("/user", middleware.ValidationMiddleware(TemplateRouter), middleware.AuthMiddleware(srv), middleware.PermitMiddleware(srv, "put", "self"), HdPutUser)

	router.GET("/users", middleware.ValidationMiddleware(TemplateRouter), middleware.AuthMiddleware(srv), middleware.PermitMiddleware(srv, "get", "all"), HdGetUsers)
	router.GET("/users/:id", middleware.ValidationMiddleware(TemplateRouter), middleware.AuthMiddleware(srv), middleware.PermitMiddleware(srv, "get", "all"), HdGetUsersByID)
	router.DELETE("/users/:id", middleware.ValidationMiddleware(TemplateRouter), middleware.AuthMiddleware(srv), middleware.PermitMiddleware(srv, "delete", "all"), HdDeleteUserById)
	router.PUT("/users/:id", middleware.ValidationMiddleware(TemplateRouter), middleware.AuthMiddleware(srv), middleware.PermitMiddleware(srv, "put", "all"), HdPutUserById)
}
