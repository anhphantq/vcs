package router

import (
	"challenge3/db"
	"challenge3/middleware"
	"challenge3/models"
	"challenge3/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var permissionService services.PermissionService

func hdGetPermission(c *gin.Context) {
	var permissions []models.Permission

	permissions, err := permissionService.GetAllPermisson()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	c.JSON(http.StatusAccepted, permissions)
}

func hdCreatePermission(c *gin.Context) {
	connection := db.GetDatabase()
	defer db.Closedatabase(connection)

	var permission models.Permission
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

	if result.RowsAffected < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Wrong permission ID"})
		return
	}

	var permission models.Permission

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

	if result.RowsAffected < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Wrong permission ID"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Permission deleted"})
}

func initPermissionRouter(router *gin.RouterGroup, userService services.UserService, permissionService services.PermissionService) {
	permissionService = permissionService
	router.GET("/permissions", middleware.ValidationMiddleware(templateRouter), middleware.RoleValidationMiddleware(userService, "admin"), hdGetPermission)
	router.POST("/permissions", middleware.ValidationMiddleware(templateRouter), middleware.RoleValidationMiddleware(userService, "admin"), hdCreatePermission)
	router.GET("/permissions/:id", middleware.ValidationMiddleware(templateRouter), middleware.RoleValidationMiddleware(userService, "admin"), hdGetPermissionByID)
	router.DELETE("/permissions/:id", middleware.ValidationMiddleware(templateRouter), middleware.RoleValidationMiddleware(userService, "admin"), hdDeletePermissionByID)
}
