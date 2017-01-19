//model/cocktail.db.go
package model

import (
	"database/sql"
	"db"
	"log"
)

func InitCocktailTables() {
	conn, _ := db.GetDB()
	var temp string
	if err := conn.QueryRow("SHOW TABLES LIKE 'cocktail';").Scan(&temp); err == nil {
		log.Println("cocktail Table Exists")
	} else if err == sql.ErrNoRows {
		log.Println("Creating cocktail Table")
		conn.Query("CREATE TABLE `commonwealthcocktails`.`cocktail` (`idcocktail` INT NOT NULL AUTO_INCREMENT,PRIMARY KEY (`idcocktail`));")
	}
}
