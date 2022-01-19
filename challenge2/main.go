package main

import (
	"challenge2/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)


func main() {
	mainRouter := gin.Default()

	mainRouter.Use(cors.Default())
	mainRouter.Use(static.Serve("/", static.LocalFile("./client/build", true)))

	router.InitRouter(mainRouter)
	mainRouter.Run(":8080")
}
