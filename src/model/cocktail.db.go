//model/cocktail.db.go
package model

import (
	"bytes"
	"database/sql"
	"db"
	"log"
	"strconv"
	"strings"
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
		query := "ALTER TABLE `commonwealthcocktails`.`cocktail`" +                                                                          //
			" ADD COLUMN `cocktailTitle` VARCHAR(150) NOT NULL AFTER `idCocktail`," + //Title
			" ADD COLUMN `cocktailName` VARCHAR(150) NOT NULL AFTER `cocktailTitle`," + //Name
			" ADD COLUMN `cocktailDisplayName` VARCHAR(150) NULL AFTER `cocktailName`," + //DisplayName
			" ADD COLUMN `cocktailSpokenName` VARCHAR(150) NULL AFTER `cocktailDisplayName`," + //SpokenName
			" ADD COLUMN `cocktailOrigin` VARCHAR(150) NULL AFTER `cocktailSpokenName`," + //Origin
			" ADD COLUMN `cocktailDescription` VARCHAR(1500) NULL AFTER `cocktailOrigin`," + //Description
			" ADD COLUMN `cocktailComment` VARCHAR(1500) NULL AFTER `cocktailDescription`," + //Comment
			" ADD COLUMN `cocktailImagePath` VARCHAR(1000) NULL AFTER `cocktailComment`," + //ImagePath
			" ADD COLUMN `cocktailImage` VARCHAR(250) NULL AFTER `cocktailImagePath`," + //Image
			" ADD COLUMN `cocktailImageSourceName` VARCHAR(250) NULL AFTER `cocktailImage`," + //ImageSourceName
			" ADD COLUMN `cocktailImageSourceLink` VARCHAR(1000) NULL AFTER `cocktailImageSourceName`," + //ImageSourceName
			" ADD COLUMN `cocktailRating` INT(1) NULL AFTER `cocktailImageSourceLink`," + //Rating
			" ADD COLUMN `cocktailFamily` INT NULL AFTER `cocktailRating`," +
			" ADD COLUMN `cocktailIsFamilyRoot` BOOLEAN NULL AFTER `cocktailFamily`," + //IsFamilyRoot
			" ADD COLUMN `cocktailSourceName` VARCHAR(150) NULL AFTER `cocktailIsFamilyRoot`," + //SourceName
			" ADD COLUMN `cocktailSourceLink` VARCHAR(150) AFTER `cocktailSourceName`;" //SourceLink
		log.Println(query)
		conn.Query(query)
	}
}

func InitCocktailReferences() {
	conn, _ := db.GetDB()
	query := "ALTER TABLE `commonwealthcocktails`.`cocktail`" +
		" ADD CONSTRAINT cocktail_cocktailFamily_id_fk FOREIGN KEY(cocktailFamily) REFERENCES meta(idMeta);"
	log.Println(query)
	conn.Query(query)
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
	//AKA
	if err := conn.QueryRow("SHOW TABLES LIKE 'cocktailToAKAs';").Scan(&temp); err == nil {
		log.Println("cocktailToAKAs Table Exists")
	} else if err == sql.ErrNoRows {
		log.Println("Creating cocktailToAKAs Table")
		conn.Exec("CREATE TABLE `commonwealthcocktails`.`cocktailToAKAs` (`idCocktailToAKANames` INT NOT NULL AUTO_INCREMENT," +
			" `idCocktail` INT NOT NULL," +
			" `idAKAName` INT NOT NULL," +
			" PRIMARY KEY (`idCocktailToAKANames`)," +
			" CONSTRAINT cocktailToAKAs_idCocktail_id_fk FOREIGN KEY(idCocktail) REFERENCES cocktail(idCocktail)," +
			" CONSTRAINT cocktailToAKAs_idAKAName_id_fk FOREIGN KEY(idAKAName) REFERENCES altnames(idAltNames));")
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
			" `idMetaType` INT NOT NULL," +
			" PRIMARY KEY (`idCocktailToMetas`)," +
			" CONSTRAINT cocktailToMetas_idCocktail_id_fk FOREIGN KEY(idCocktail) REFERENCES cocktail(idCocktail)," +
			" CONSTRAINT cocktailToMetas_idMeta_id_fk FOREIGN KEY(idMeta) REFERENCES meta(idMeta)," +
			" CONSTRAINT cocktailToMetas_idMetaType_id_fk FOREIGN KEY(idMetaType) REFERENCES metatype(idMetaType));")
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
			" `idProductType` INT NOT NULL," +
			" PRIMARY KEY (`idCocktailToProducts`)," +
			" CONSTRAINT cocktailToProducts_idCocktail_id_fk FOREIGN KEY(idCocktail) REFERENCES cocktail(idCocktail)," +
			" CONSTRAINT cocktailToProducts_idProduct_id_fk FOREIGN KEY(idProduct) REFERENCES product(idProduct)," +
			" CONSTRAINT cocktailToProducts_idProductType_id_fk FOREIGN KEY(idProductType) REFERENCES producttype(idProductType));")
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

func ProcessCocktails() {
	for _, cocktail := range Cocktails {
		ProcessCocktail(cocktail)
	}
}

func ProcessCocktail(cocktail Cocktail) int {
	conn, _ := db.GetDB()
	var buffer bytes.Buffer
	buffer.WriteString("INSERT INTO `commonwealthcocktails`.`cocktail` SET ")
	if cocktail.Title != "" {
		buffer.WriteString("`cocktailTitle`=\"" + cocktail.Title + "\",")
	}
	if cocktail.Name != "" {
		buffer.WriteString("`cocktailName`=\"" + cocktail.Name + "\",")
	}
	if cocktail.Description != "" {
		sqlString := strings.Replace(string(cocktail.Description), "\\", "\\\\", -1)
		sqlString = strings.Replace(sqlString, "\"", "\\\"", -1)
		buffer.WriteString("`cocktailDescription`=\"" + sqlString + "\",")
	}
	if cocktail.ImagePath != "" {
		buffer.WriteString("`cocktailImagePath`=\"" + cocktail.ImagePath + "\",")
	}
	if cocktail.Image != "" {
		buffer.WriteString("`cocktailImage`=\"" + cocktail.Image + "\",")
	}
	if cocktail.ImageSourceName != "" {
		buffer.WriteString("`cocktailImageSourceName`=\"" + cocktail.ImageSourceName + "\",")
	}
	if cocktail.ImageSourceLink != "" {
		buffer.WriteString("`cocktailImageSourceLink`=\"" + cocktail.ImageSourceLink + "\",")
	}
	if cocktail.SourceName != "" {
		buffer.WriteString("`cocktailSourceName`=\"" + cocktail.SourceName + "\",")
	}
	if cocktail.SourceLink != "" {
		buffer.WriteString("`cocktailSourceLink`=\"" + cocktail.SourceLink + "\",")
	}
	if cocktail.Rating != 0 {
		buffer.WriteString(" `cocktailRating`=" + strconv.Itoa(cocktail.Rating) + ",")
	}
	if cocktail.Origin != "" {
		sqlString := strings.Replace(string(cocktail.Origin), "\\", "\\\\", -1)
		sqlString = strings.Replace(sqlString, "\"", "\\\"", -1)
		buffer.WriteString("`cocktailOrigin`=\"" + sqlString + "\",")
	}
	if cocktail.SpokenName != "" {
		buffer.WriteString("`cocktailSpokenName`=\"" + cocktail.SpokenName + "\",")
	}
	if cocktail.DisplayName != "" {
		buffer.WriteString("`cocktailDisplayName`=\"" + cocktail.DisplayName + "\",")
	}
	if cocktail.IsFamilyRoot {
		buffer.WriteString(" `cocktailIsFamilyRoot`='1',")
	} else {
		buffer.WriteString(" `cocktailIsFamilyRoot`='0',")
	}

	query := buffer.String()
	query = strings.TrimRight(query, ",")
	query = query + ";"
	log.Println(query)
	res, err := conn.Exec(query)
	lastCocktailId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Cocktail ID = %d, affected = %d\n", lastCocktailId, rowCnt)

	recipeID := ProcessRecipe(cocktail.Recipe)

	ProcessCocktailToMeta(cocktail.Family, lastCocktailId)

	ProcessAKAs(cocktail.AKA, lastCocktailId)
	ProcessAltNames(cocktail.AlternateName, lastCocktailId)
	ProcessCocktailToProducts(cocktail.Garnish, lastCocktailId)
	ProcessCocktailToProducts(cocktail.Drinkware, lastCocktailId)
	ProcessCocktailToProducts(cocktail.Tool, lastCocktailId)

	ProcessCocktailToMetas(cocktail.Flavor, lastCocktailId)
	ProcessCocktailToMetas(cocktail.Type, lastCocktailId)
	ProcessCocktailToMetas(cocktail.BaseSpirit, lastCocktailId)
	ProcessCocktailToMetas(cocktail.Served, lastCocktailId)
	ProcessCocktailToMetas(cocktail.Technique, lastCocktailId)
	ProcessCocktailToMetas(cocktail.Strength, lastCocktailId)
	ProcessCocktailToMetas(cocktail.Difficulty, lastCocktailId)
	ProcessCocktailToMetas(cocktail.TOD, lastCocktailId)
	ProcessCocktailToMetas(cocktail.Ratio, lastCocktailId)

	conn.Exec("INSERT INTO `commonwealthcocktails`.`cocktailToRecipe` (`idCocktail`, `idRecipe`) VALUES ('" + strconv.FormatInt(lastCocktailId, 10) + "', '" + strconv.Itoa(recipeID) + "');")

	return int(lastCocktailId)
}

func ProcessAKAs(names []Name, cocktailID int64) {
	conn, _ := db.GetDB()
	var buffer bytes.Buffer
	for _, name := range names {
		buffer.WriteString("INSERT INTO `commonwealthcocktails`.`altnames` SET ")

		if name.Name != "" {
			buffer.WriteString("`altNamesString`=\"" + name.Name + "\",")
		}

		query := buffer.String()
		query = strings.TrimRight(query, ",")
		query = query + ";"
		log.Println(query)
		res, err := conn.Exec(query)
		lastAltNameId, err := res.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}

		conn.Exec("INSERT INTO `commonwealthcocktails`.`cocktailToAKAs` (`idCocktail`, `idAKAName`) VALUES ('" + strconv.FormatInt(cocktailID, 10) + "', '" + strconv.FormatInt(lastAltNameId, 10) + "');")
	}
}

func ProcessAltNames(names []Name, cocktailID int64) {
	conn, _ := db.GetDB()
	var buffer bytes.Buffer
	for _, name := range names {
		buffer.WriteString("INSERT INTO `commonwealthcocktails`.`altnames` SET ")

		if name.Name != "" {
			buffer.WriteString("`altNamesString`=\"" + name.Name + "\",")
		}

		query := buffer.String()
		query = strings.TrimRight(query, ",")
		query = query + ";"
		log.Println(query)
		res, err := conn.Exec(query)
		lastAltNameId, err := res.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}

		conn.Exec("INSERT INTO `commonwealthcocktails`.`cocktailToAltNames` (`idCocktail`, `idAltName`) VALUES ('" + strconv.FormatInt(cocktailID, 10) + "', '" + strconv.FormatInt(lastAltNameId, 10) + "');")
	}
}

func ProcessCocktailToMeta(meta Meta, cocktailID int64) {
	conn, _ := db.GetDB()
	metaTo := SelectMeta(meta)
	query := "UPDATE `commonwealthcocktails`.`cocktail` SET `cocktailFamily`='" + strconv.Itoa(metaTo.ID) + "' WHERE `idCocktail`='" + strconv.FormatInt(cocktailID, 10) + "';"
	log.Println(query)
	conn.Exec(query)
}

func ProcessCocktailToProducts(products []Product, cocktailID int64) {
	conn, _ := db.GetDB()
	for _, product := range products {
		prodTo := SelectProduct(product)
		query := "INSERT INTO `commonwealthcocktails`.`cocktailToProducts` (`idCocktail`, `idProduct`, `idProductType`) VALUES ('" + strconv.FormatInt(cocktailID, 10) + "', '" + strconv.Itoa(prodTo.ID) + "', '" + strconv.Itoa(int(prodTo.ProductType)) + "');"
		log.Println(query)
		conn.Exec(query)
	}

}

func ProcessCocktailToMetas(metas []Meta, cocktailID int64) {
	conn, _ := db.GetDB()
	for _, meta := range metas {
		metaTo := SelectMeta(meta)
		query := "INSERT INTO `commonwealthcocktails`.`cocktailToMetas` (`idCocktail`, `idMeta`, `idMetaType`) VALUES ('" + strconv.FormatInt(cocktailID, 10) + "', '" + strconv.Itoa(metaTo.ID) + "', '" + strconv.Itoa(int(metaTo.MetaType)) + "');"
		log.Println(query)
		conn.Exec(query)
	}

}
