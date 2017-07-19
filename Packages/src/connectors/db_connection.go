//connectors/db_connection.go
package connectors

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB = nil
var dbaddr string
var dbpasswd string
var user string
var proto string
var port string
var dbname string

func SetDBVars(in_dbaddr string, in_dbpasswd string, in_user string, in_proto string, in_port string, in_dbname string) {
	dbaddr = in_dbaddr
	dbpasswd = in_dbpasswd
	user = in_user
	proto = in_proto
	port = in_port
	dbname = in_dbname
	log.Println(user + ":" + dbpasswd + "@" + proto + "(" + dbaddr + ":" + port + ")/" + dbname + "?timeout=1m")
}

func GetDB() (*sql.DB, error) {
	if db == nil {
		log.Println("Creating a new connection: mysql", user+":"+dbpasswd+"@"+proto+"("+dbaddr+":"+port+")/"+dbname+"?timeout=1m")
		d, err := sql.Open("mysql", user+":"+dbpasswd+"@"+proto+"("+dbaddr+":"+port+")/"+dbname+"?timeout=1m")
		if err != nil {
			log.Println("Error connecting to database")
			log.Fatal(err)
			return nil, err
		}
		db = d
	}
	return db, nil
}