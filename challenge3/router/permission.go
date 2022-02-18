package router

import (
	"challenge3/db"
	"challenge3/middleware"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func hdGetPermission(c *gin.Context) {
	connection := db.GetDatabase()
	defer db.Closedatabase(connection)

	var permissions []db.Permission

	result := connection.Raw("select * from permissions").Scan(&permissions)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	c.JSON(http.StatusAccepted, permissions)
}

func hdCreatePermission(c *gin.Context) {
	connection := db.GetDatabase()
	defer db.Closedatabase(connection)

	var permission db.Permission
	if err := c.ShouldBind(&permission); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	result := connection.Exec("insert into permissions values(default,?,?)", permission.Name, permission.Scope)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Permission created"})
}

func hdGetPermissionByID(c *gin.Context) {
	connection := db.GetDatabase()
	defer db.Closedatabase(connection)

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	result := connection.Exec("select * from permissions where permission_id=?", id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	if result.RowsAffected < 1{
		c.JSON(http.StatusBadRequest, gin.H{"message": "Wrong permission ID"})
		return
	}

	var permission db.Permission

	result.First(&permission)

	c.JSON(http.StatusAccepted, permission)
}

func hdDeletePermissionByID(c *gin.Context) {
	connection := db.GetDatabase()
	defer db.Closedatabase(connection)

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	result := connection.Exec("delete from permissions where permission_id=?", id)
	
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	if result.RowsAffected < 1{
		c.JSON(http.StatusBadRequest, gin.H{"message": "Wrong permission ID"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message" : "Permission deleted"})
}

func initPermissionRouter(router *gin.RouterGroup) {
	router.GET("/permissions", middleware.ValidationMiddleware(templateRouter), middleware.AuthAdminMiddleware(), hdGetPermission)
	router.POST("/permissions", middleware.ValidationMiddleware(templateRouter), middleware.AuthAdminMiddleware(), hdCreatePermission)
	router.GET("/permissions/:id", middleware.ValidationMiddleware(templateRouter), middleware.AuthAdminMiddleware(), hdGetPermissionByID)
	router.DELETE("/permissions/:id", middleware.ValidationMiddleware(templateRouter), middleware.AuthAdminMiddleware(), hdDeletePermissionByID)
}
