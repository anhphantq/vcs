package db

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB_HOST, DB_PORT, DB_USER, DB_NAME, DB_PWD string

func InitVar(host, port, user, name, pwd string) {
	DB_HOST, DB_PORT, DB_USER, DB_NAME, DB_PWD = host, port, user, name, pwd
}

func GetDatabase() *gorm.DB {
	connection, err := gorm.Open("postgres", "host="+DB_HOST+" port="+DB_PORT+" user="+DB_USER+" sslmode=disable dbname="+DB_NAME+" password="+DB_PWD)
	if err != nil {
		log.Fatalln("wrong database url", err, DB_HOST, "hi")
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
