//model/meta.db.go
package model

import (
	"database/sql"
	"db"
	"log"
)

func InitMetaTables() {
	conn, _ := db.GetDB()
	var temp string
	if err := conn.QueryRow("SHOW TABLES LIKE 'metatype';").Scan(&temp); err == nil {
		log.Println("metatype Table Exists")
	} else if err == sql.ErrNoRows {
		log.Println("Creating metatype Table")
		conn.Query("CREATE TABLE `commonwealthcocktails`.`metatype` (`idMetaType` INT NOT NULL AUTO_INCREMENT,  PRIMARY KEY (`idMetaType`));")
	}

	if err := conn.QueryRow("SHOW TABLES LIKE 'meta';").Scan(&temp); err == nil {
		log.Println("Meta Table Exists")
	} else if err == sql.ErrNoRows {
		log.Println("Creating Meta Table")
		conn.Query("CREATE TABLE `commonwealthcocktails`.`meta` (`idMeta` INT NOT NULL AUTO_INCREMENT,PRIMARY KEY (`idMeta`));")
	}
}
