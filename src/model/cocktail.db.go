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

	if err := conn.QueryRow("SHOW TABLES LIKE 'altnames';").Scan(&temp); err == nil {
		log.Println("altnames Table Exists")
	} else if err == sql.ErrNoRows {
		log.Println("Creating altnames Table")
		conn.Query("CREATE TABLE `commonwealthcocktails`.`altnames` (`idAltNames` INT NOT NULL AUTO_INCREMENT,  PRIMARY KEY (`idAltNames`));")
		conn.Query("ALTER TABLE `commonwealthcocktails`.`altnames` ADD COLUMN `altNamesString` VARCHAR(250) NOT NULL AFTER `idAltNames`;")
	}

	if err := conn.QueryRow("SHOW TABLES LIKE 'cocktail';").Scan(&temp); err == nil {
		log.Println("cocktail Table Exists")
	} else if err == sql.ErrNoRows {
		log.Println("Creating cocktail Table")
		conn.Query("CREATE TABLE `commonwealthcocktails`.`cocktail` (`idCocktail` INT NOT NULL AUTO_INCREMENT,PRIMARY KEY (`idCocktail`));") //ID
		conn.Query("ALTER TABLE `commonwealthcocktails`.`cocktail`" +                                                                        //
			" ADD COLUMN `cocktailTitle` VARCHAR(150) NOT NULL AFTER `idCocktail`," + //Title
			" ADD COLUMN `cocktailName` VARCHAR(150) NOT NULL AFTER `cocktailTitle`," + //Name
			" ADD COLUMN `cocktailDisplayName` VARCHAR(150) NOT NULL AFTER `cocktailName`," + //DisplayName
			" ADD COLUMN `cocktailSpokenName` VARCHAR(150) NOT NULL AFTER `cocktailDisplayName`," + //SpokenName
			" ADD COLUMN `cocktailOrigin` VARCHAR(150) NOT NULL AFTER `cocktailSpokenName`," + //Origin
			" ADD COLUMN `cocktailDescription` VARCHAR(1500) NOT NULL AFTER `cocktailOrigin`," + //Description
			" ADD COLUMN `cocktailComment` VARCHAR(1500) NOT NULL AFTER `cocktailDescription`," + //Comment
			" ADD COLUMN `cocktailImagePath` VARCHAR(1000) NOT NULL AFTER `cocktailComment`," + //ImagePath
			" ADD COLUMN `cocktailImage` VARCHAR(250) NOT NULL AFTER `cocktailImagePath`," + //Image
			" ADD COLUMN `cocktailImageSourceName` VARCHAR(250) NOT NULL AFTER `cocktailImage`," + //ImageSourceName
			" ADD COLUMN `cocktailImageSourceLink` VARCHAR(1000) NOT NULL AFTER `cocktailImageSourceName`," + //ImageSourceName
			" ADD COLUMN `cocktailImageRating` INT(1) NOT NULL AFTER `cocktailImageSourceLink`," + //Rating
			" ADD COLUMN `cocktailIsFamilyRoot` BOOLEAN NOT NULL AFTER `cocktailImageRating`," + //IsFamilyRoot
			" ADD COLUMN `cocktailSourceName` VARCHAR(150) NOT NULL AFTER `cocktailIsFamilyRoot`," + //SourceName
			" ADD COLUMN `cocktailSourceLink` VARCHAR(150) NULL AFTER `cocktailSourceName`;") //SourceLink
	}
}

func InitCocktailReferences() {
	addCocktailToRecipeReference()
	addCocktailToAltNamesTable()
	addCocktailToMetasTable()
	addCocktailToProductsTable()
	addCocktailToPostsTable()
}

func addCocktailToRecipeReference() {
	//Recipe
	conn, _ := db.GetDB()
	var temp string
	if err := conn.QueryRow("SHOW TABLES LIKE 'cocktailToRecipe';").Scan(&temp); err == nil {
		log.Println("cocktailToRecipe Table Exists")
	} else if err == sql.ErrNoRows {
		log.Println("Creating cocktailToRecipe Table")
		conn.Exec("CREATE TABLE `commonwealthcocktails`.`cocktailToRecipe` (`idCocktailToRecipe` INT NOT NULL AUTO_INCREMENT," +
			" `idCocktail` INT NOT NULL," +
			" `idRecipe` INT NOT NULL," +
			" PRIMARY KEY (`idCocktailToRecipe`)," +
			" CONSTRAINT cocktailToRecipe_idCocktail_id_fk FOREIGN KEY(idCocktail) REFERENCES cocktail(idCocktail)," +
			" CONSTRAINT cocktailToRecipe_idRecipe_id_fk FOREIGN KEY(idRecipe) REFERENCES recipe(idRecipe));")
	}
}

func addCocktailToAltNamesTable() {
	//AlternateName
	//AKA
	conn, _ := db.GetDB()
	var temp string
	if err := conn.QueryRow("SHOW TABLES LIKE 'cocktailToAltNames';").Scan(&temp); err == nil {
		log.Println("CocktailToAltNames Table Exists")
	} else if err == sql.ErrNoRows {
		log.Println("Creating CocktailToAltNames Table")
		conn.Exec("CREATE TABLE `commonwealthcocktails`.`cocktailToAltNames` (`idCocktailToAltNames` INT NOT NULL AUTO_INCREMENT," +
			" `idCocktail` INT NOT NULL," +
			" `idAltName` INT NOT NULL," +
			" PRIMARY KEY (`idCocktailToAltNames`)," +
			" CONSTRAINT cocktailToAltNames_idCocktail_id_fk FOREIGN KEY(idCocktail) REFERENCES cocktail(idCocktail)," +
			" CONSTRAINT cocktailToAltNames_idAltName_id_fk FOREIGN KEY(idAltName) REFERENCES altnames(idAltNames));")
	}
}

func addCocktailToMetasTable() {
	//Flavor
	//Type
	//BaseSpirit
	//Served
	//Technique
	//Strength
	//Difficulty
	//TOD
	//Ratio
	//Family

	conn, _ := db.GetDB()
	var temp string
	if err := conn.QueryRow("SHOW TABLES LIKE 'cocktailToMetas';").Scan(&temp); err == nil {
		log.Println("cocktailToMetas Table Exists")
	} else if err == sql.ErrNoRows {
		log.Println("Creating cocktailToMetas Table")
		conn.Exec("CREATE TABLE `commonwealthcocktails`.`cocktailToMetas` (`idCocktailToMetas` INT NOT NULL AUTO_INCREMENT," +
			" `idCocktail` INT NOT NULL," +
			" `idMeta` INT NOT NULL," +
			" PRIMARY KEY (`idCocktailToMetas`)," +
			" CONSTRAINT cocktailToMetas_idCocktail_id_fk FOREIGN KEY(idCocktail) REFERENCES cocktail(idCocktail)," +
			" CONSTRAINT cocktailToMetas_idMeta_id_fk FOREIGN KEY(idMeta) REFERENCES meta(idMeta));")
	}
}

func addCocktailToProductsTable() {
	//Garnish
	//Drinkware
	//Tool
	conn, _ := db.GetDB()
	var temp string
	if err := conn.QueryRow("SHOW TABLES LIKE 'cocktailToProducts';").Scan(&temp); err == nil {
		log.Println("cocktailToProducts Table Exists")
	} else if err == sql.ErrNoRows {
		log.Println("Creating cocktailToProducts Table")
		conn.Exec("CREATE TABLE `commonwealthcocktails`.`cocktailToProducts` (`idCocktailToProducts` INT NOT NULL AUTO_INCREMENT," +
			" `idCocktail` INT NOT NULL," +
			" `idProduct` INT NOT NULL," +
			" PRIMARY KEY (`idCocktailToProducts`)," +
			" CONSTRAINT cocktailToProducts_idCocktail_id_fk FOREIGN KEY(idCocktail) REFERENCES cocktail(idCocktail)," +
			" CONSTRAINT cocktailToProducts_idProduct_id_fk FOREIGN KEY(idProduct) REFERENCES product(idProduct));")
	}
}

func addCocktailToPostsTable() {
	//About
	//Articles
	conn, _ := db.GetDB()
	var temp string
	if err := conn.QueryRow("SHOW TABLES LIKE 'cocktailToPosts';").Scan(&temp); err == nil {
		log.Println("cocktailToPosts Table Exists")
	} else if err == sql.ErrNoRows {
		log.Println("Creating cocktailToPosts Table")
		conn.Exec("CREATE TABLE `commonwealthcocktails`.`cocktailToPosts` (`idCocktailToPosts` INT NOT NULL AUTO_INCREMENT," +
			" `idCocktail` INT NOT NULL," +
			" `idPost` INT NOT NULL," +
			" PRIMARY KEY (`idCocktailToPosts`)," +
			" CONSTRAINT cocktailToPosts_idCocktail_id_fk FOREIGN KEY(idCocktail) REFERENCES cocktail(idCocktail)," +
			" CONSTRAINT cocktailToPosts_idPost_id_fk FOREIGN KEY(idPost) REFERENCES post(idPost));")
	}
}
