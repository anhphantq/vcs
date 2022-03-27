package router

import (
	"challenge4/middleware"
	"challenge4/models"
	"challenge4/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

var grantService services.GrantService

func hdGetGrant(c *gin.Context) {
	var Rolepermissions []models.Rolepermission

	Rolepermissions, err := grantService.GetAllGrant()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	c.JSON(http.StatusAccepted, Rolepermissions)
}

func hdCreateGrant(c *gin.Context) {
	var Rolepermission models.Rolepermission
	if err := c.ShouldBind(&Rolepermission); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something when wrong in the server"})
		return
	}

	Rolepermission, err := grantService.InsertGrant(Rolepermission)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something when wrong in the server or the role had been granted"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Granted"})
}

func hdDeleteGrant(c *gin.Context) {
	var Rolepermission models.Rolepermission
	if err := c.ShouldBind(&Rolepermission); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something when wrong in the server"})
		return
	}

	err := grantService.DeleteGrantSrv(Rolepermission)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something when wrong in the server or wrong input"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Deleted"})
}

func initGrantRouter(router *gin.RouterGroup, userService services.UserService, grantservice services.GrantService) {
	grantService = grantservice
	router.GET("/granting", middleware.ValidationMiddleware(templateRouter), middleware.RoleValidationMiddleware(userService, "admin"), hdGetGrant)
	router.POST("/granting", middleware.ValidationMiddleware(templateRouter), middleware.RoleValidationMiddleware(userService, "admin"), hdCreateGrant)
	router.DELETE("/granting", middleware.ValidationMiddleware(templateRouter), middleware.RoleValidationMiddleware(userService, "admin"), hdDeleteGrant)
}
