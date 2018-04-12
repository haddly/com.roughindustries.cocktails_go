// Copyright 2017 Rough Industries LLC. All rights reserved.
//connectors/db_connection.go: This is a singleton that provides a pool for
//connecting to the database
package connectors

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
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

//Database struct
type DBItem struct {
	DB       *sql.DB
	DBaddr   string
	DBpasswd string
	DBuser   string
	DBproto  string
	DBport   string
	DBname   string
	DBType   DBTypesConst
	Site     string
}

//Map of database connections
var db_list = make(map[string]DBItem)

//Set the database variables from a list
func SetDBMap(dbs map[string]DBItem) {
	for k, v := range dbs {
		db_list[k] = v
	}
	log.Infoln(db_list)
}

func GetDBType(site string) DBTypesConst {
	return db_list[site].DBType
}

//Get a connection to the database
func GetDBFromMap(site string) (*sql.DB, error) {
	//log.Infoln(site)
	//log.Infoln(db_list)
	if _, ok := db_list[site]; ok {
		log.Infoln(db_list[site])
		if db_list[site].DB == nil {
			//which db do you want to use
			var err error
			var d *sql.DB
			log.Infoln(db_list[site].DBType.String())
			if db_list[site].DBType == MySQL {
				log.Infoln("Creating a new connection: mysql", db_list[site].DBuser+":"+db_list[site].DBpasswd+"@"+db_list[site].DBproto+"("+db_list[site].DBaddr+":"+db_list[site].DBport+")/"+db_list[site].DBname+"?parseTime=true&timeout=1m")
				d, err = sql.Open("mysql", db_list[site].DBuser+":"+db_list[site].DBpasswd+"@"+db_list[site].DBproto+"("+db_list[site].DBaddr+":"+db_list[site].DBport+")/"+db_list[site].DBname+"?parseTime=true&timeout=1m")
			} else if db_list[site].DBType == SQLite {
				log.Infoln("Creating a new connection: sqllite to " + db_list[site].DBname + ".db")
				d, err = sql.Open("sqlite3", "./sql/"+db_list[site].DBname+".db")
			}
			if err != nil {
				log.Infoln("Error connecting to database")
				log.Error(err)
				return nil, err
			}
			tmp := db_list[site]
			tmp.DB = d
			db_list[site] = tmp
		} else {
			//log.Infoln(db_list[site].DB)
		}
		log.Infoln(db_list[site].DB.Stats())
		return db_list[site].DB, nil
	}
	return nil, nil
}
