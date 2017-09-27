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
	rows, _ := conn.Query("SELECT DATABASE();")
	glog.Infoln("Got Databases")
	var dbname string
	if rows.Next() {
		rows.Scan(&dbname)
		glog.Infoln(dbname)
	}
	return dbname
}
