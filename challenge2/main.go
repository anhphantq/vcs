package main

import (
	"challenge2/db"
	"challenge2/router"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.InitVar(os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_PWD"))
	mainRouter := gin.Default()

	mainRouter.Use(cors.Default())
	mainRouter.Use(static.Serve("/", static.LocalFile("./client/build", true)))

	router.InitRouter(mainRouter)
	mainRouter.Run(":8080")
}
