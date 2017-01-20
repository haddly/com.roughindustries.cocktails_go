//model/advertisement.db.go
package model

import (
	"database/sql"
	"db"
	"log"
)

func InitAdvertisementTables() {
	conn, _ := db.GetDB()
	var temp string
	if err := conn.QueryRow("SHOW TABLES LIKE 'adtype';").Scan(&temp); err == nil {
		log.Println("adtype Table Exists")
	} else if err == sql.ErrNoRows {
		log.Println("Creating adtype Table")
		conn.Query("CREATE TABLE `commonwealthcocktails`.`adtype` (`idAdType` INT NOT NULL AUTO_INCREMENT,  PRIMARY KEY (`idAdType`));")
	}

	if err := conn.QueryRow("SHOW TABLES LIKE 'advertisement';").Scan(&temp); err == nil {
		log.Println("advertisement Table Exists")
	} else if err == sql.ErrNoRows {
		log.Println("Creating advertisement Table")
		conn.Query("CREATE TABLE `commonwealthcocktails`.`advertisement` (`idAdvertisement` INT NOT NULL AUTO_INCREMENT,PRIMARY KEY (`idAdvertisement`));")
	}
}
