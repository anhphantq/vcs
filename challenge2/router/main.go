package router

import "github.com/gin-gonic/gin"

func InitRouter(router *gin.Engine) {

	userRouter := router.Group("/user-management/user")
	initUserRouter(userRouter)

	roleRouter := router.Group("/user-management/role")
	initRoleRouter(roleRouter)

	grantRouter := router.Group("/user-management/grant")
	initGrantRouter(grantRouter)

	permissionRouter := router.Group("/user-management/permission")
	initPermissionRouter(permissionRouter)
}
