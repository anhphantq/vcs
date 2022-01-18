package router

import (
	"challenge2/db"
	"challenge2/middleware"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusAccepted, permissions)
}

func hdCreatePermission(c *gin.Context) {
	connection := db.GetDatabase()
	defer db.Closedatabase(connection)

	var permission db.Permission
	if err := c.ShouldBind(&permission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := connection.Exec("insert into permissions values(default,?,?)", permission.Name, permission.Scope)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	c.String(http.StatusAccepted, "Permissions created")
}

func hdUpdatePermission(c *gin.Context) {
}

func hdDeletePermission(c *gin.Context) {
	connection := db.GetDatabase()
	defer db.Closedatabase(connection)

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := connection.Exec("delete from permissions where permission_id=?", id)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.String(http.StatusAccepted, "Permission deleted")
}

func InitPermissionRouter(router *gin.RouterGroup) {
	router.GET("/", middleware.AuthAdminMiddleware(), hdGetPermission)
	router.POST("/", middleware.AuthAdminMiddleware(), hdCreatePermission)
	// router.PUT("/", middleware.AuthRoleMiddleware(), hdUpdateRole)
	router.DELETE("/:id", middleware.AuthAdminMiddleware(), hdDeletePermission)
}
