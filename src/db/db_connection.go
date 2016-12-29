//model/db_connection.go
package db

import (
	"database/sql"
	//"github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var db *sql.DB = nil

func GetDB() (*sql.DB, error) {
	if db == nil {
		log.Println("Creating a new connection:")
		log.Println(os.Getenv("ccdb_password"))
		d, err := sql.Open("mysql", "root:"+"password"+"@tcp(104.196.178.43:3306)/commonwealthcocktails?timeout=1m")
		if err != nil {
			return nil, err
		}
		db = d
	}

	return db, nil
}
