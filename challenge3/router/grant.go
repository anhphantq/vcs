package router

import (
	"challenge3/db"
	"challenge3/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func hdGetGrant(c *gin.Context) {
	connection := db.GetDatabase()
	defer db.Closedatabase(connection)

	var rolepermissions []db.RolePermission

	result := connection.Raw("select * from rolepermissions").Scan(&rolepermissions)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	c.JSON(http.StatusAccepted, rolepermissions)
}

func hdCreateGrant(c *gin.Context) {
	connection := db.GetDatabase()
	defer db.Closedatabase(connection)

	var rolepermission db.RolePermission
	if err := c.ShouldBind(&rolepermission); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something when wrong in the server"})
		return
	}

	result := connection.Exec("insert into rolepermissions values(?,?)", rolepermission.Role_id, rolepermission.Permission_id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something when wrong in the server or the role had been granted"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Granted"})
}

func hdDeleteGrant(c *gin.Context) {
	connection := db.GetDatabase()
	defer db.Closedatabase(connection)

	var rolepermission db.RolePermission
	if err := c.ShouldBind(&rolepermission); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something when wrong in the server"})
		return
	}

	result := connection.Exec("delete from rolepermissions where role_id=? and permission_id=?", rolepermission.Role_id, rolepermission.Permission_id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something when wrong in the server"})
		return
	}

	if result.RowsAffected < 1{
		c.JSON(http.StatusBadRequest, gin.H{"message": "Wrong role's ID or permission's ID"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Deleted"})
}

func initGrantRouter(router *gin.RouterGroup) {
	router.GET("/granting", middleware.ValidationMiddleware(templateRouter), middleware.AuthAdminMiddleware(), hdGetGrant)
	router.POST("/granting", middleware.ValidationMiddleware(templateRouter), middleware.AuthAdminMiddleware(), hdCreateGrant)
	router.DELETE("/granting", middleware.ValidationMiddleware(templateRouter), middleware.AuthAdminMiddleware(), hdDeleteGrant)
}
