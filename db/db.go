package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Connect() *sql.DB {
	log.Println("Connecting to mysql server...")

	dbConnect, err := sql.Open("mysql", "root:anmap1234@tcp(127.0.0.1:3306)/demoAPI")
	if err != nil {
		panic(err.Error())
	}
	db = dbConnect

	return db
}

func Execute(query string) (*sql.Rows, error) {
	result, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}

	return result, err
}

func Close() {
	log.Println("Closing connect mysql server...")
	defer db.Close()
}
