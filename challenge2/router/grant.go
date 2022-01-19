package router

import (
	"challenge2/db"
	"challenge2/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func hdGetGrant(c *gin.Context) {
	connection := db.GetDatabase()
	defer db.Closedatabase(connection)

	var rolepermissions []db.RolePermission

	result := connection.Raw("select * from rolepermissions").Scan(&rolepermissions)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusAccepted, rolepermissions)
}

func hdCreateGrant(c *gin.Context) {
	connection := db.GetDatabase()
	defer db.Closedatabase(connection)

	var rolepermission db.RolePermission
	if err := c.ShouldBind(&rolepermission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := connection.Exec("insert into rolepermissions values(?,?)", rolepermission.Role_id, rolepermission.Permission_id)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	c.String(http.StatusAccepted, "Granted")
}

func hdUpdateGrant(c *gin.Context) {
}

func hdDeleteGrant(c *gin.Context) {
	connection := db.GetDatabase()
	defer db.Closedatabase(connection)

	var rolepermission db.RolePermission
	if err := c.ShouldBind(&rolepermission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := connection.Exec("delete from rolepermissions where role_id=? and permission_id=?", rolepermission.Role_id, rolepermission.Permission_id)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	c.String(http.StatusAccepted, "Grant deleted")
}

func initGrantRouter(router *gin.RouterGroup) {
	router.GET("", middleware.AuthAdminMiddleware(), hdGetGrant)
	router.POST("", middleware.AuthAdminMiddleware(), hdCreateGrant)
	router.DELETE("", middleware.AuthAdminMiddleware(), hdDeleteGrant)
}
