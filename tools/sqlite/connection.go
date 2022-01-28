package sqlite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var connection *sql.DB

func setConnection() {
	var err error
	connection, err = sql.Open("sqlite3", "./foo.db")
	if err != nil {
		log.Fatal(err)
	}
}
func GetConnection() *sql.DB {
	if connection == nil {
		setConnection()
	}
	return connection
}
