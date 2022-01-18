package main

import (
	"challenge2/router"

	"github.com/gin-gonic/gin"
)

func main() {
	mainRouter := gin.Default()

	userRouter := mainRouter.Group("/user-management/user")
	router.InitUserRouter(userRouter)

	roleRouter := mainRouter.Group("/user-management/role")
	router.InitRoleRouter(roleRouter)

	grant

	mainRouter.Run(":8080")
}
