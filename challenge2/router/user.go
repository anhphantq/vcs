package router

import (
	"challenge2/db"
	"challenge2/middleware"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

	validate := validator.New()

	if err = validate.Struct(user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "validation error"})
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
	tmp, _ := c.Get("user")
	user, _ := tmp.(db.Account)

	c.JSON(http.StatusAccepted, gin.H{"user_id": user.User_id, "username": user.Username, "email": user.Email, "role_id": user.User_id})
}

func hdDeleteUser(c *gin.Context) {
	tmp, _ := c.Get("user")
	user, _ := tmp.(db.Account)

	connection := db.GetDatabase()
	defer db.Closedatabase(connection)

	result := connection.Exec("delete from accounts where user_id = ?", user.User_id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "something went wrong when deleting account"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"msg": "account deleted"})
}

func hdPutUser(c *gin.Context) {
	tmp, _ := c.Get("user")
	user, _ := tmp.(db.Account)

	var account db.Account

	if err := c.ShouldBind(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if account.Password != "" {
		user.Password = account.Password
	}

	if account.Username != "" {
		user.Username = account.Username
	}

	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "validation error"})
		return
	}

	if account.Password != "" {
		user.Password, err = GeneratePassword(user.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	connection := db.GetDatabase()
	defer db.Closedatabase(connection)

	result := connection.Exec("update accounts SET username=?, email=?, password=? where user_id=?", user.Username, user.Email, user.Password, user.User_id)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"msg": "account updated"})
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

	c.JSON(http.StatusAccepted, accounts)
}

func hdDeleteUserById(c *gin.Context) {
	connection := db.GetDatabase()
	defer db.Closedatabase(connection)
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := connection.Exec("delete from accounts where user_id=?", id)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.String(http.StatusAccepted, "User deleted")
}

func hdPutUserById(c *gin.Context) {
	connection := db.GetDatabase()
	defer db.Closedatabase(connection)
	var account db.Account

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBind(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	err = validate.Struct(account)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "validation error"})
		return
	}

	account.Password, err = GeneratePassword(account.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := connection.Exec("update accounts SET username=?, role_id=?, password=? where user_id=?", account.Username, account.Role_id, account.Password, id)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.String(http.StatusAccepted, "Info updated")
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
