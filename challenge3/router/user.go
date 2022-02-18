package router

import (
	"challenge3/db"
	"challenge3/middleware"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func generatePassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func signUp(c *gin.Context) {
	connection := db.GetDatabase()
	defer db.Closedatabase(connection)

	var user db.Account

	err := c.ShouldBind(&user)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"messsage": "Something went wrong in the server"})
		return
	}

	var dbUser db.Account
	connection.Where("email = ?", user.Email).Find(&dbUser)

	if dbUser.Email != "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "This email has been used!"})
		return
	}

	user.Password, err = generatePassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Can not generate password"})
		return
	}

	result := connection.Exec("INSERT INTO ACCOUNTS VALUES(DEFAULT,?,?,?,?)", user.Username, user.Email, user.Role_id, user.Password)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server or your role is not in the DB"})
		return
	}

	connection.Where("email = ?", user.Email).Find(&user)

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
	connection := db.GetDatabase()
	defer db.Closedatabase(connection)

	var authDetails db.Authentication

	err := c.ShouldBind(&authDetails)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"messsage": "Something went wrong in the server"})
		return
	}

	var user db.Account
	connection.Where("email = ?", authDetails.Email).Find(&user)
	if user.Email == "" {
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
	user, _ := tmp.(db.Account)

	c.JSON(http.StatusAccepted, gin.H{"user_id": user.User_id, "username": user.Username, "email": user.Email, "role_id": user.Role_id})
}

func hdDeleteUser(c *gin.Context) {
	tmp, _ := c.Get("user")
	user, _ := tmp.(db.Account)

	connection := db.GetDatabase()
	defer db.Closedatabase(connection)

	result := connection.Exec("delete from accounts where user_id = ?", user.User_id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong when deleting account"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Account deleted"})
}

func hdPutUser(c *gin.Context) {
	tmp, _ := c.Get("user")
	user, _ := tmp.(db.Account)

	var account db.Account

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

	connection := db.GetDatabase()
	defer db.Closedatabase(connection)

	result := connection.Exec("update accounts SET username=?, email=?, password=? where user_id=?", user.Username, user.Email, user.Password, user.User_id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Account updated"})
}

func hdGetUsers(c *gin.Context) {
	connection := db.GetDatabase()
	defer db.Closedatabase(connection)

	var accounts []db.Account

	result := connection.Raw("select * from accounts").Scan(&accounts)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	var usersInfo []gin.H

	for i := range accounts {
		usersInfo = append(usersInfo, gin.H{"user_id": accounts[i].User_id, "username": accounts[i].Username, "email": accounts[i].Email, "role_id": accounts[i].Role_id})
	}

	c.JSON(http.StatusAccepted, usersInfo)
}

func hdGetUsersByID(c *gin.Context){
	connection := db.GetDatabase()
	defer db.Closedatabase(connection)

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	var user db.Account
	connection.Where("user_id = ?", id).Find(&user)

	if user == (db.Account{}) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Wrong user's ID"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"user_id": user.User_id, "username": user.Username, "email": user.Email, "role_id": user.Role_id})
}

func hdDeleteUserById(c *gin.Context) {
	connection := db.GetDatabase()
	defer db.Closedatabase(connection)
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	result := connection.Exec("delete from accounts where user_id=?", id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	if result.RowsAffected < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Wrong user's ID"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Account deleted"})
}

func hdPutUserById(c *gin.Context) {
	connection := db.GetDatabase()
	defer db.Closedatabase(connection)

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}
	
	var user db.Account
	connection.Where("user_id = ?", id).Find(&user)
	
	var account db.Account

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

	result := connection.Exec("update accounts SET username=?, password=? where user_id=?", user.Username, user.Password, id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	if result.RowsAffected < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Wrong user's ID"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Info updated"})
}

func initUserRouter(router *gin.RouterGroup) {
	router.POST("/signup", middleware.ValidationMiddleware(templateRouter), signUp)
	router.POST("/signin", middleware.ValidationMiddleware(templateRouter), signIn)

	router.GET("/user", middleware.ValidationMiddleware(templateRouter),middleware.AuthMiddleware(), middleware.PermitMiddleware("get", "self"), hdGetUser)
	router.DELETE("/user", middleware.ValidationMiddleware(templateRouter),middleware.AuthMiddleware(), middleware.PermitMiddleware("delete", "self"), hdDeleteUser)
	router.PUT("/user", middleware.ValidationMiddleware(templateRouter),middleware.AuthMiddleware(), middleware.PermitMiddleware("put", "self"), hdPutUser)

	router.GET("/users", middleware.ValidationMiddleware(templateRouter),middleware.AuthMiddleware(), middleware.PermitMiddleware("get", "all"), hdGetUsers)
	router.GET("/users/:id", middleware.ValidationMiddleware(templateRouter),middleware.AuthMiddleware(), middleware.PermitMiddleware("get", "all"), hdGetUsersByID)
	router.DELETE("/users/:id", middleware.ValidationMiddleware(templateRouter),middleware.AuthMiddleware(), middleware.PermitMiddleware("delete", "all"), hdDeleteUserById)
	router.PUT("/users/:id", middleware.ValidationMiddleware(templateRouter),middleware.AuthMiddleware(), middleware.PermitMiddleware("put", "all"), hdPutUserById)
}
