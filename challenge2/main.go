package main

import (
	"challenge2/router"

	"github.com/gin-gonic/gin"
)

func main() {
	mainRouter := gin.Default()

	userRouter := mainRouter.Group("/user-management/user")
	router.InitUserRouter(userRouter)

	mainRouter.Run(":8080")
}
