// Copyright 2017 Rough Industries LLC. All rights reserved.
//model/database.go:package model
package model

import (
	"github.com/CommonwealthCocktails/connectors"
	"github.com/golang/glog"
)

//SELECTS
//Get the current name of the database being used.
func SelectCurrentDB() string {
	glog.Infoln("Getting CurrentDB")
	conn, _ := connectors.GetDB()
	glog.Infoln("Getting Databases")
	var dbname string
	if connectors.DBType == connectors.MySQL {
		rows, _ := conn.Query("SELECT DATABASE();")
		glog.Infoln("Got Databases")
		if rows.Next() {
			rows.Scan(&dbname)
			glog.Infoln(dbname)
		}
	} else if connectors.DBType == connectors.SQLite {
		dbname = "./sql/commonwealthcocktails.db"
	}

	return dbname
}

func SelectTables() []string {
	conn, _ := connectors.GetDB()
	glog.Infoln("Getting Tables")
	var tables []string
	if connectors.DBType == connectors.MySQL {
		rows, _ := conn.Query("SELECT table_name FROM information_schema.tables where table_schema='commonwealthcocktails';")
		glog.Infoln("Got Tables")
		var table string
		for rows.Next() {
			rows.Scan(&table)
			tables = append(tables, table)
		}
	} else if connectors.DBType == connectors.SQLite {
		rows, _ := conn.Query("SELECT name FROM SQLITE_MASTER WHERE type='table' ORDER BY name;")
		glog.Infoln("Got Tables")
		var table string
		for rows.Next() {
			rows.Scan(&table)
			tables = append(tables, table)
		}
	}
	glog.Infoln(tables)
	return tables

}
