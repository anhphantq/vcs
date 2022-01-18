package router

import (
	"challenge2/db"
	"challenge2/middleware"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func hdGetGrant(c *gin.Context) {
	connection := db.GetDatabase()
	defer db.Closedatabase(connection)

	var roles []db.Role

	result := connection.Raw("select * from accounts").Scan(&roles)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusAccepted, roles)
}

func hdCreateGrant(c *gin.Context) {
	connection := db.GetDatabase()
	defer db.Closedatabase(connection)

	var role db.Role
	if err := c.ShouldBind(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := connection.Exec("insert into roles values(default,?)", role.Name)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	c.String(http.StatusAccepted, "Role created")
}

func hdUpdateGrant(c *gin.Context) {
}

func hdDeleteGrant(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.Exec("delete from roles where role_id=?", id)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.String(http.StatusAccepted, "Role deleted")
}

func InitGrantRouter(router *gin.RouterGroup) {
	router.GET("/", middleware.AuthRoleMiddleware(), hdGetRole)
	router.POST("/", middleware.AuthRoleMiddleware(), hdCreateRole)
	// router.PUT("/", middleware.AuthRoleMiddleware(), hdUpdateRole)
	router.DELETE("/:id", middleware.AuthRoleMiddleware(), hdDeleteRole)
}
