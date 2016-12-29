//model/database.go
package model

import (
	"db"
	"log"
)

type Database struct {
	Name string
}

func GetCurrentDB() Database {
	conn, _ := db.GetDB()
	rows, _ := conn.Query("SELECT DATABASE();")
	var dbname string
	if rows.Next() {
		rows.Scan(&dbname)
		log.Println(dbname)
	}
	db := Database{dbname}
	return db
}
