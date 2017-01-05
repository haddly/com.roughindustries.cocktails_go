//model/database.go
package model

import (
	"db"
	"log"
)

func GetCurrentDB() string {
	conn, _ := db.GetDB()
	rows, _ := conn.Query("SELECT DATABASE();")
	var dbname string
	if rows.Next() {
		rows.Scan(&dbname)
		log.Println(dbname)
	}
	return dbname
}
