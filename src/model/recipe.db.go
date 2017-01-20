//model/recipe.db.go
package model

import (
	"database/sql"
	"db"
	"log"
)

func InitRecipeTables() {
	conn, _ := db.GetDB()
	var temp string
	if err := conn.QueryRow("SHOW TABLES LIKE 'recipe';").Scan(&temp); err == nil {
		log.Println("recipe Table Exists")
	} else if err == sql.ErrNoRows {
		log.Println("Creating recipe Table")
		conn.Query("CREATE TABLE `commonwealthcocktails`.`recipe` (`idRecipe` INT NOT NULL AUTO_INCREMENT,PRIMARY KEY (`idRecipe`));")
	}

	if err := conn.QueryRow("SHOW TABLES LIKE 'recipestep';").Scan(&temp); err == nil {
		log.Println("recipestep Table Exists")
	} else if err == sql.ErrNoRows {
		log.Println("Creating recipestep Table")
		conn.Query("CREATE TABLE `commonwealthcocktails`.`recipestep` (`idRecipeStep` INT NOT NULL AUTO_INCREMENT,PRIMARY KEY (`idRecipeStep`));")
	}

	if err := conn.QueryRow("SHOW TABLES LIKE 'doze';").Scan(&temp); err == nil {
		log.Println("doze Table Exists")
	} else if err == sql.ErrNoRows {
		log.Println("Creating doze Table")
		conn.Query("CREATE TABLE `commonwealthcocktails`.`doze` (`idDoze` INT NOT NULL AUTO_INCREMENT,PRIMARY KEY (`idDoze`));")
	}
}
