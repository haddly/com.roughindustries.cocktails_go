// Copyright 2017 Rough Industries LLC. All rights reserved.
//connectors/db_connection.go: This is a singleton that provides a pool for
//connecting to the database
package connectors

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	//_ "github.com/mattn/go-sqlite3"
)

//Database types constant enumeration
type DBTypesConst int

//Types of databases supported in the enum
const (
	MySQL = 1 + iota
	SQLite
)

//String values for the database enumeration
var DBTypeStrings = [...]string{
	"MySQL",
	"SQLite",
}

// String returns the English name of the database types ("MySSQL", "SQLite", ...).
func (dt DBTypesConst) String() string { return DBTypeStrings[dt-1] }

//Database variables
var db *sql.DB = nil
var dbaddr string
var dbpasswd string
var user string
var proto string
var port string
var dbname string
var DBType DBTypesConst

//Set the database variables for connecting to the database
func SetDBVars(in_dbaddr string, in_dbpasswd string, in_user string, in_proto string, in_port string, in_dbname string, in_dbtype DBTypesConst) {
	dbaddr = in_dbaddr
	dbpasswd = in_dbpasswd
	user = in_user
	proto = in_proto
	port = in_port
	dbname = in_dbname
	DBType = in_dbtype
	glog.Infoln(DBType)
	if DBType == MySQL {
		glog.Infoln(user + ":" + dbpasswd + "@" + proto + "(" + dbaddr + ":" + port + ")/" + dbname + "?parseTime=true&timeout=1m")
	} else if DBType == SQLite {
		glog.Infoln("SQLite connecting to ./sql/commonwealthcocktails.db")
	}
}

//Get a connection to the database
func GetDB() (*sql.DB, error) {
	if db == nil {
		//which db do you want to use
		var err error
		var d *sql.DB
		if DBType == MySQL {
			glog.Infoln("Creating a new connection: mysql", user+":"+dbpasswd+"@"+proto+"("+dbaddr+":"+port+")/"+dbname+"?parseTime=true&timeout=1m")
			d, err = sql.Open("mysql", user+":"+dbpasswd+"@"+proto+"("+dbaddr+":"+port+")/"+dbname+"?parseTime=true&timeout=1m")
		} else if DBType == SQLite {
			glog.Infoln("Creating a new connection: sqllite to commonwealthcocktails.db")
			d, err = sql.Open("sqlite3", "./sql/commonwealthcocktails.db")
		}
		if err != nil {
			glog.Infoln("Error connecting to database")
			glog.Error(err)
			return nil, err
		}
		db = d
	}
	return db, nil
}
