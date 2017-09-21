// Copyright 2017 Rough Industries LLC. All rights reserved.
//model/database.go:package model
package model

import (
	"connectors"
	"log"
)

//SELECTS
//Get the current name of the database being used.
func SelectCurrentDB() string {
	log.Println("Getting CurrentDB")
	conn, _ := connectors.GetDB()
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
