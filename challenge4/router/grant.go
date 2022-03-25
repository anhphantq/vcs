package router

import (
	"challenge3/middleware"
	"challenge3/models"
	"challenge3/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

var grantService services.GrantService

func hdGetGrant(c *gin.Context) {
	var rolepermissions []models.RolePermission

	rolepermissions, err := grantService.GetAllGrant()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	c.JSON(http.StatusAccepted, rolepermissions)
}

func hdCreateGrant(c *gin.Context) {
	var rolepermission models.RolePermission
	if err := c.ShouldBind(&rolepermission); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something when wrong in the server"})
		return
	}

	rolepermission, err := grantService.SaveGrant(rolepermission)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something when wrong in the server or the role had been granted"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Granted"})
}

func hdDeleteGrant(c *gin.Context) {
	var rolepermission models.RolePermission
	if err := c.ShouldBind(&rolepermission); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something when wrong in the server"})
		return
	}

	err := grantService.DeleteGrant(rolepermission)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something when wrong in the server or wrong input"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Deleted"})
}

func initGrantRouter(router *gin.RouterGroup, userService services.UserService, grantService services.GrantService) {
	grantService = grantService
	router.GET("/granting", middleware.ValidationMiddleware(templateRouter), middleware.RoleValidationMiddleware(userService, "admin"), hdGetGrant)
	router.POST("/granting", middleware.ValidationMiddleware(templateRouter), middleware.RoleValidationMiddleware(userService, "admin"), hdCreateGrant)
	router.DELETE("/granting", middleware.ValidationMiddleware(templateRouter), middleware.RoleValidationMiddleware(userService, "admin"), hdDeleteGrant)
}
