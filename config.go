package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	// We'll use postgres, we need its dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var dbHost = os.Getenv("DB_HOST")
var dbPort = os.Getenv("DB_PORT")
var dbUser = os.Getenv("DB_USER")
var dbPass = os.Getenv("DB_PASS")
var dbName = os.Getenv("DB_NAME")

func connectDB() *gorm.DB {
	connStr := "host=" + dbHost + " port=" + dbPort + " user=" + dbUser +
		" password=" + dbPass + " dbname=" + dbName
	fmt.Println(connStr)
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return db
}

//DB is a connection to Postgres service
var DB = connectDB()
