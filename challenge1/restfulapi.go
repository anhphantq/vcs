package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type account struct {
	User_id  uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"name" binding:"required"`
	Birthday string `json:"birthday" binding:"required"`
	Sex      *bool  `json:"sex,string"`
	Email    string `json:"email" binding:"required"`
}

func main() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=accounts sslmode=disable password=Phananh272")
	router := gin.Default()

	if err != nil {
		panic(err)
	}

	fmt.Print("connect success")

	router.GET("/user-management/users/:name", func(c *gin.Context) {
		name := c.Param("name")
		var accounts []account

		db.Where("username = ?", name).Find(&accounts)
		c.JSON(http.StatusAccepted, accounts)
	})

	router.GET("/user-management/users", func(c *gin.Context) {
		var accounts []account

		db.Raw("select * from accounts").Scan(&accounts)
		c.JSON(http.StatusAccepted, accounts)
	})

	router.POST("/user-management/users", func(c *gin.Context) {
		var account account
		if err := c.ShouldBind(&account); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result := db.Exec("insert into accounts values(default,?,?,?,?)", account.Username, account.Email, account.Birthday, account.Sex)

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.String(http.StatusAccepted, "User added")
	})

	router.PUT("/user-management/users/:id", func(c *gin.Context) {
		var account account

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := c.ShouldBind(&account); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result := db.Exec("update accounts SET username=?, email=?, birthday=?, sex=? where user_id=?", account.Username, account.Email, account.Birthday, account.Sex, id)

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.String(http.StatusAccepted, "Info updated")
	})

	router.DELETE("/user-management/users/:id", func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result := db.Exec("delete from accounts where user_id=?", id)

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.String(http.StatusAccepted, "User deleted")
	})

	router.Run("localhost:8080")

	defer db.Close()
}
