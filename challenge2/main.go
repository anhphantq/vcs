package main

import (
	"challenge2/router"

	"github.com/gin-gonic/gin"
)

func main() {
	mainRouter := gin.Default()

	router.InitRouter(mainRouter)

	mainRouter.Run(":8080")
}
