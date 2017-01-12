//model/cocktail.db.go
package model

import (
	"db"
	"log"
)

func InitCocktailTable() string {
	conn, _ := db.GetDB()
	rows, _ := conn.Query("SHOW TABLES;")
	var tables string
	for rows.Next() {
		var temp string
		rows.Scan(&temp)
		tables = tables + "<br>" + temp
		log.Println(tables)
	}
	return tables
}
