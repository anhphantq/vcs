package db

import (
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var lock = &sync.Mutex{}

var singleDB *gorm.DB

func GetDB() *gorm.DB {
	if singleDB == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleDB == nil {
			singleDB = getDatabase()
		} else {
			return singleDB
		}
	} else {
		return singleDB
	}

	return singleDB
}

func getDatabase() *gorm.DB {
	dsn := "host=localhost user=postgres password=Phananh272 dbname=challenge3 port=5432 sslmode=disable"
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//connection, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres sslmode=disable dbname=challenge3 password=Phananh272")
	if err != nil {
		log.Fatalln("wrong database url", err)
	}

	sqldb, _ := connection.DB()

	err = sqldb.Ping()
	if err != nil {
		log.Fatal("database connected")
	}

	fmt.Println("connected to database")
	return connection
}
