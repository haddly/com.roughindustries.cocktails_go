// Copyright 2017 Rough Industries LLC. All rights reserved.
//model/database.go:package model
package model

import (
	"connectors"
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
