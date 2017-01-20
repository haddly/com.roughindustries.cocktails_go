//model/product.db.go
package model

import (
	"database/sql"
	"db"
	"log"
)

func InitProductTables() {
	conn, _ := db.GetDB()
	var temp string
	if err := conn.QueryRow("SHOW TABLES LIKE 'producttype';").Scan(&temp); err == nil {
		log.Println("producttype Table Exists")
	} else if err == sql.ErrNoRows {
		log.Println("Creating producttype Table")
		conn.Query("CREATE TABLE `commonwealthcocktails`.`producttype` (`idProductType` INT NOT NULL AUTO_INCREMENT,PRIMARY KEY (`idProductType`));")
	}

	if err := conn.QueryRow("SHOW TABLES LIKE 'bdgtype';").Scan(&temp); err == nil {
		log.Println("bdgtype Table Exists")
	} else if err == sql.ErrNoRows {
		log.Println("Creating bdgtype Table")
		conn.Query("CREATE TABLE `commonwealthcocktails`.`bdgtype` (`idBDGType` INT NOT NULL AUTO_INCREMENT,PRIMARY KEY (`idBDGType`));")
	}

	if err := conn.QueryRow("SHOW TABLES LIKE 'product';").Scan(&temp); err == nil {
		log.Println("product Table Exists")
	} else if err == sql.ErrNoRows {
		log.Println("Creating product Table")
		conn.Query("CREATE TABLE `commonwealthcocktails`.`product` (`idProduct` INT NOT NULL AUTO_INCREMENT,PRIMARY KEY (`idProduct`));")
	}
}
