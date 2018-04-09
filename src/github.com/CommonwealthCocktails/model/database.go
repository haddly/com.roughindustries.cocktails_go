// Copyright 2017 Rough Industries LLC. All rights reserved.
//model/database.go:package model
package model

import (
	"github.com/CommonwealthCocktails/connectors"
	log "github.com/sirupsen/logrus"
)

//SELECTS
//Get the current name of the database being used.
func SelectCurrentDB(site string) string {
	log.Infoln("Getting CurrentDB")
	conn, _ := connectors.GetDBFromMap(site)
	log.Infoln("Getting Databases")
	var dbname string
	if connectors.GetDBType(site) == connectors.MySQL {
		rows, _ := conn.Query("SELECT DATABASE();")
		log.Infoln("Got Databases")
		if rows.Next() {
			rows.Scan(&dbname)
			log.Infoln(dbname)
		}
	} else if connectors.GetDBType(site) == connectors.SQLite {
		dbname = "./sql/commonwealthcocktails.db"
	}

	return dbname
}

func SelectTables(site string) []string {
	conn, _ := connectors.GetDBFromMap(site)
	log.Infoln("Getting Tables")
	var tables []string
	if connectors.GetDBType(site) == connectors.MySQL {
		rows, _ := conn.Query("SELECT table_name FROM information_schema.tables where table_schema='commonwealthcocktails';")
		log.Infoln("Got Tables")
		var table string
		for rows.Next() {
			rows.Scan(&table)
			tables = append(tables, table)
		}
	} else if connectors.GetDBType(site) == connectors.SQLite {
		rows, _ := conn.Query("SELECT name FROM SQLITE_MASTER WHERE type='table' ORDER BY name;")
		log.Infoln("Got Tables")
		var table string
		for rows.Next() {
			rows.Scan(&table)
			tables = append(tables, table)
		}
	}
	log.Infoln(tables)
	return tables

}
