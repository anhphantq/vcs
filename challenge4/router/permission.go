package router

import (
	"challenge4/middleware"
	"challenge4/models"
	"challenge4/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var permissionService services.PermissionService

func hdGetPermission(c *gin.Context) {
	var permissions []models.Permission

	permissions, err := permissionService.GetAllPermission()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	c.JSON(http.StatusAccepted, permissions)
}

func hdCreatePermission(c *gin.Context) {

	var permission models.Permission
	if err := c.ShouldBind(&permission); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	permission, err := permissionService.InsertPermission(permission)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Permission created"})
}

func hdGetPermissionByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	permission, err := permissionService.GetPermissionByID(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	c.JSON(http.StatusAccepted, permission)
}

func hdDeletePermissionByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	err = permissionService.DeletePermissionByID(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Permission deleted"})
}

func initPermissionRouter(router *gin.RouterGroup, userService services.UserService, permissionservice services.PermissionService) {
	permissionService = permissionservice
	router.GET("/permissions", middleware.ValidationMiddleware(templateRouter), middleware.RoleValidationMiddleware(userService, "admin"), hdGetPermission)
	router.POST("/permissions", middleware.ValidationMiddleware(templateRouter), middleware.RoleValidationMiddleware(userService, "admin"), hdCreatePermission)
	router.GET("/permissions/:id", middleware.ValidationMiddleware(templateRouter), middleware.RoleValidationMiddleware(userService, "admin"), hdGetPermissionByID)
	router.DELETE("/permissions/:id", middleware.ValidationMiddleware(templateRouter), middleware.RoleValidationMiddleware(userService, "admin"), hdDeletePermissionByID)
}
