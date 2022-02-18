package router

import (
	"challenge3/db"
	"challenge3/middleware"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func hdGetRoles(c *gin.Context) {
	connection := db.GetDatabase()
	defer db.Closedatabase(connection)

	var roles []db.Role

	result := connection.Raw("select * from roles").Scan(&roles)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	c.JSON(http.StatusAccepted, roles)
}

func hdCreateRole(c *gin.Context) {
	connection := db.GetDatabase()
	defer db.Closedatabase(connection)

	var role db.Role
	if err := c.ShouldBind(&role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	result := connection.Exec("insert into roles values(default,?)", role.Name)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Role created"})
}

func hdGetRoleByID(c *gin.Context) {
	connection := db.GetDatabase()
	defer db.Closedatabase(connection)

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	result := connection.Exec("select * from roles where role_id=?", id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	if result.RowsAffected < 1{
		c.JSON(http.StatusBadRequest, gin.H{"message": "Wrong role ID"})
		return
	}

	var role db.Role

	result.First(&role)

	c.JSON(http.StatusAccepted, role)
}

func hdDeleteRoleByID(c *gin.Context) {
	connection := db.GetDatabase()
	defer db.Closedatabase(connection)

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	result := connection.Exec("delete from roles where role_id=?", id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	if result.RowsAffected < 1{
		c.JSON(http.StatusBadRequest, gin.H{"message": "Wrong role ID"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Role deleted"})
}

func initRoleRouter(router *gin.RouterGroup) {
	router.GET("/roles", middleware.ValidationMiddleware(templateRouter), middleware.AuthAdminMiddleware(), hdGetRoles)
	router.POST("/roles", middleware.ValidationMiddleware(templateRouter), middleware.AuthAdminMiddleware(), hdCreateRole)
	router.GET("/roles/:id", middleware.ValidationMiddleware(templateRouter), middleware.AuthAdminMiddleware(), hdGetRoleByID)
	router.DELETE("/roles/:id", middleware.ValidationMiddleware(templateRouter), middleware.AuthAdminMiddleware(), hdDeleteRoleByID)
}
