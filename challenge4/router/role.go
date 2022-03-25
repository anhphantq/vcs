package router

import (
	"challenge3/middleware"
	"challenge3/models"
	"challenge3/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var roleService services.RoleService

func hdGetRoles(c *gin.Context) {
	var roles []models.Role

	roles, err := roleService.GetAllRoles()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	c.JSON(http.StatusAccepted, roles)
}

func hdCreateRole(c *gin.Context) {
	var role models.Role
	if err := c.ShouldBind(&role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	role, err := roleService.SaveRole(role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Role created"})
}

func hdGetRoleByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	role, err := roleService.GetRoleByID(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server or wrong role ID"})
		return
	}

	c.JSON(http.StatusAccepted, role)
}

func hdDeleteRoleByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	err = roleService.DeleteRoleByID(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server or wrong role ID"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Role deleted"})
}

func initRoleRouter(router *gin.RouterGroup, userService services.UserService, roleService services.RoleService) {
	roleService = roleService
	router.GET("/roles", middleware.ValidationMiddleware(templateRouter), middleware.RoleValidationMiddleware(userService, "admin"), hdGetRoles)
	router.POST("/roles", middleware.ValidationMiddleware(templateRouter), middleware.RoleValidationMiddleware(userService, "admin"), hdCreateRole)
	router.GET("/roles/:id", middleware.ValidationMiddleware(templateRouter), middleware.RoleValidationMiddleware(userService, "admin"), hdGetRoleByID)
	router.DELETE("/roles/:id", middleware.ValidationMiddleware(templateRouter), middleware.RoleValidationMiddleware(userService, "admin"), hdDeleteRoleByID)
}
