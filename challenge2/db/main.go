package db

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetDatabase() *gorm.DB {
	connection, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres sslmode=disable dbname=challenge2 password=Phananh272")
	if err != nil {
		log.Fatalln("wrong database url", err)
	}

	sqldb := connection.DB()

	err = sqldb.Ping()
	if err != nil {
		log.Fatal("database connected")
	}

	fmt.Println("connected to database")
	return connection
}

func Closedatabase(connection *gorm.DB) {
	sqldb := connection.DB()
	sqldb.Close()
}
