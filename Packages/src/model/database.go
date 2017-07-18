//model/database.go
package model

import (
	"connectors"
	"log"
	"strings"
)

func GetCurrentDB() string {
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

func MysqlRealEscapeString(value string) string {
	replace := map[string]string{"\\": "\\\\", "'": `\'`, "\\0": "\\\\0", "\n": "\\n", "\r": "\\r", `"`: `\"`, "\x1a": "\\Z"}

	for b, a := range replace {
		value = strings.Replace(value, b, a, -1)
	}

	return value
}
