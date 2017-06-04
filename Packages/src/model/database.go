//model/database.go
package model

import (
	"db"
	"log"
)

func GetCurrentDB() string {
	log.Println("Getting CurrentDB")
	conn, _ := db.GetDB()
	log.Println("Getting Databases")
	rows, _ := conn.Query("SELECT DATABASE();")
	log.Println("Got Databases")
	var dbname string
	if rows.Next() {
		rows.Scan(&dbname)
		log.Println(dbname)
	}
	return dbname
}
