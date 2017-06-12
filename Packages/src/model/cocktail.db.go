//model/cocktail.db.go
package model

import (
	"bytes"
	"database/sql"
	"db"
	"html/template"
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
			" ADD COLUMN `cocktailSourceName` VARCHAR(150) NULL AFTER `cocktailRating`," + //SourceName
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
			" `isRootCocktailForMeta` BOOLEAN NOT NULL," +
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
	if cocktail.Comment != "" {
		sqlString := strings.Replace(string(cocktail.Comment), "\\", "\\\\", -1)
		sqlString = strings.Replace(sqlString, "\"", "\\\"", -1)
		buffer.WriteString("`cocktailComment`=\"" + sqlString + "\",")
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
	ProcessAKAs(cocktail.AKA, lastCocktailId)
	ProcessAltNames(cocktail.AlternateName, lastCocktailId)
	ProcessCocktailToProducts(cocktail.Garnish, lastCocktailId)
	ProcessCocktailToProducts(cocktail.Drinkware, lastCocktailId)
	ProcessCocktailToProducts(cocktail.Tool, lastCocktailId)

	log.Println("Processing Cocktail to Metas")
	ProcessCocktailToMetas(cocktail.Family, lastCocktailId, btoi(cocktail.IsFamilyRoot))
	ProcessCocktailToMetas(cocktail.Flavor, lastCocktailId, 0)
	ProcessCocktailToMetas(cocktail.Type, lastCocktailId, 0)
	ProcessCocktailToMetas(cocktail.BaseSpirit, lastCocktailId, 0)
	ProcessCocktailToMetas(cocktail.Served, lastCocktailId, 0)
	ProcessCocktailToMetas(cocktail.Technique, lastCocktailId, 0)
	ProcessCocktailToMetas(cocktail.Strength, lastCocktailId, 0)
	ProcessCocktailToMetas(cocktail.Difficulty, lastCocktailId, 0)
	ProcessCocktailToMetas(cocktail.TOD, lastCocktailId, 0)
	ProcessCocktailToMetas(cocktail.Ratio, lastCocktailId, 0)

	conn.Exec("INSERT INTO `commonwealthcocktails`.`cocktailToRecipe` (`idCocktail`, `idRecipe`) VALUES ('" + strconv.FormatInt(lastCocktailId, 10) + "', '" + strconv.Itoa(recipeID) + "');")

	return int(lastCocktailId)
}

func ProcessAKAs(names []Name, cocktailID int64) {
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
		conn, _ := db.GetDB()
		res, err := conn.Exec(query)
		lastAltNameId, err := res.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		conn, _ = db.GetDB()
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

func ProcessCocktailToProducts(products []Product, cocktailID int64) {
	conn, _ := db.GetDB()

	for _, product := range products {
		prodTo := SelectProduct(product)
		if len(prodTo) > 0 {
			query := "INSERT INTO `commonwealthcocktails`.`cocktailToProducts` (`idCocktail`, `idProduct`, `idProductType`) VALUES ('" + strconv.FormatInt(cocktailID, 10) + "', '" + strconv.Itoa(prodTo[0].ID) + "', '" + strconv.Itoa(int(prodTo[0].ProductType)) + "');"
			log.Println(query)
			conn.Exec(query)
		}
	}
}

func ProcessCocktailToMetas(metas []Meta, cocktailID int64, isRootCocktailForMeta int) {
	conn, _ := db.GetDB()

	for _, meta := range metas {
		metaTo := SelectMeta(meta)
		if len(metaTo) > 0 {
			query := "INSERT INTO `commonwealthcocktails`.`cocktailToMetas` (`idCocktail`, `idMeta`, `idMetaType`, `isRootCocktailForMeta`) VALUES ('" + strconv.FormatInt(cocktailID, 10) + "', '" + strconv.Itoa(metaTo[0].ID) + "', '" + strconv.Itoa(int(metaTo[0].MetaType.ID)) + "', '" + strconv.Itoa(int(isRootCocktailForMeta)) + "');"
			log.Println(query)
			conn.Exec(query)
		}
	}
}

//func SelectCocktailRecipeSteps(cocktail Cocktail) []RecipeStep{
//SELECT * FROM commonwealthcocktails.recipestep
//JOIN commonwealthcocktails.recipeToRecipeSteps ON recipeToRecipeSteps.idRecipeStep=recipestep.idRecipeStep
//JOIN commonwealthcocktails.recipe ON  recipeToRecipeSteps.idRecipe=recipe.idRecipe
//JOIN commonwealthcocktails.cocktailToRecipe ON cocktailToRecipe.idRecipe=recipe.idRecipe
//JOIN  commonwealthcocktails.cocktail ON cocktailToRecipe.idCocktail=cocktail.idCocktail
//WHERE cocktail.idCocktail=1 ORDER BY recipestepRecipeOrdinal;
//}

func SelectCocktailsByMeta(meta Meta) []Cocktail {
	var ret []Cocktail
	conn, _ := db.GetDB()

	log.Println(meta.MetaName)
	var buffer bytes.Buffer
	var canQuery = false
	buffer.WriteString("SELECT cocktail.idCocktail, cocktail.cocktailTitle, cocktail.cocktailName, cocktail.cocktailRating," +
		" cocktail.cocktailImagePath, cocktail.cocktailImage, COALESCE(cocktail.cocktailDescription, ''), COALESCE(cocktail.cocktailComment, '')" +
		" FROM commonwealthcocktails.cocktail" +
		" JOIN commonwealthcocktails.cocktailToMetas ON cocktailToMetas.idCocktail=cocktail.idCocktail WHERE")
	//cocktailToMetas.idMeta=23;
	//
	if meta.ID != 0 {
		buffer.WriteString(" `cocktailToMetas`.`idMeta`=" + strconv.Itoa(meta.ID) + " AND")
		canQuery = true
	}
	// if meta.MetaType.ID != 0 {
	// 	buffer.WriteString(" `cocktailToMetas`.`idMetaType`=" + strconv.Itoa(meta.MetaType.ID) + " AND")
	// 	canQuery = true
	// }

	if canQuery {
		query := buffer.String()
		query = strings.TrimRight(query, " AND")
		query = query + ";"
		log.Println(query)
		rows, err := conn.Query(query)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			var cocktail Cocktail
			var desc string
			var comment string
			err := rows.Scan(&cocktail.ID, &cocktail.Title, &cocktail.Name, &cocktail.Rating, &cocktail.ImagePath, &cocktail.Image, &desc, &comment)
			cocktail.Description = template.HTML(desc)
			cocktail.Comment = template.HTML(comment)
			if err != nil {
				log.Fatal(err)
			}
			ret = append(ret, cocktail)
			log.Println(cocktail.ID, cocktail.Title, cocktail.Name, cocktail.Rating, cocktail.ImagePath, cocktail.Image, cocktail.Description, cocktail.Comment)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
	}
	return ret
}

func SelectCocktailsByProduct(product Product) []Cocktail {
	var ret []Cocktail
	conn, _ := db.GetDB()

	log.Println(product.ProductName)
	var buffer bytes.Buffer
	var canQuery = false
	buffer.WriteString("SELECT cocktail.idCocktail, cocktail.cocktailTitle, cocktail.cocktailName, cocktail.cocktailRating," +
		" cocktail.cocktailImagePath, cocktail.cocktailImage, COALESCE(cocktail.cocktailDescription, ''), COALESCE(cocktail.cocktailComment, '')" +
		" FROM commonwealthcocktails.cocktail" +
		" JOIN commonwealthcocktails.cocktailToProducts ON cocktailToProducts.idCocktail=cocktail.idCocktail WHERE")
	//cocktailToMetas.idProduct=23;
	//
	if product.ID != 0 {
		buffer.WriteString(" `cocktailToProducts`.`idProduct`=" + strconv.Itoa(product.ID) + " AND")
		canQuery = true
	}
	if int(product.ProductType) != 0 {
		buffer.WriteString(" `cocktailToProducts`.`idProductType`=" + strconv.Itoa(int(product.ProductType)) + " AND")
		canQuery = true
	}

	if canQuery {
		query := buffer.String()
		query = strings.TrimRight(query, " AND")
		query = query + ";"
		log.Println(query)
		rows, err := conn.Query(query)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			var cocktail Cocktail
			var desc string
			var comment string
			err := rows.Scan(&cocktail.ID, &cocktail.Title, &cocktail.Name, &cocktail.Rating, &cocktail.ImagePath, &cocktail.Image, &desc, &comment)
			cocktail.Description = template.HTML(desc)
			cocktail.Comment = template.HTML(comment)
			if err != nil {
				log.Fatal(err)
			}
			ret = append(ret, cocktail)
			log.Println(cocktail.ID, cocktail.Title, cocktail.Name, cocktail.Rating, cocktail.ImagePath, cocktail.Image, cocktail.Description, cocktail.Comment)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
	}
	return ret
}

func SelectCocktailsByID(ID int) CocktailSet {
	var cs CocktailSet
	conn, _ := db.GetDB()

	var buffer bytes.Buffer
	var canQuery = false
	buffer.WriteString("SELECT cocktail.idCocktail, cocktail.cocktailTitle, cocktail.cocktailName, cocktail.cocktailRating," +
		" cocktail.cocktailImagePath, cocktail.cocktailImage, COALESCE(cocktail.cocktailDescription, ''), COALESCE(cocktail.cocktailComment, '')," +
		" cocktail.cocktailSourceName, cocktail.cocktailSourceLink, cocktail.cocktailImageSourceName, cocktail.cocktailImageSourceLink" +
		" FROM commonwealthcocktails.cocktail" +
		" WHERE idCocktail=" + strconv.Itoa(ID) + ";")
	canQuery = true
	if canQuery {
		query := buffer.String()
		query = strings.TrimRight(query, " AND")
		query = query + ";"
		log.Println(query)
		rows, err := conn.Query(query)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			var cocktail Cocktail
			var desc string
			var comment string
			err := rows.Scan(&cocktail.ID, &cocktail.Title, &cocktail.Name, &cocktail.Rating, &cocktail.ImagePath,
				&cocktail.Image, &desc, &comment, &cocktail.SourceName, &cocktail.SourceLink, &cocktail.ImageSourceName, &cocktail.ImageSourceLink)
			cocktail.Description = template.HTML(desc)
			cocktail.Comment = template.HTML(comment)
			if err != nil {
				log.Fatal(err)
			}
			//add recipe to cocktail
			cocktail.Recipe = SelectRecipeByCocktail(cocktail)
			cocktail.Drinkware = SelectProductsByCocktailAndProductType(cocktail.ID, int(Drinkware))
			cocktail.Garnish = SelectProductsByCocktailAndProductType(cocktail.ID, int(Garnish))
			cocktail.Tool = SelectProductsByCocktailAndProductType(cocktail.ID, int(Tool))
			cocktail.Flavor = SelectMetasByCocktailAndMetaType(cocktail.ID, int(Flavor))
			cocktail.BaseSpirit = SelectMetasByCocktailAndMetaType(cocktail.ID, int(BaseSpirit))
			cocktail.Type = SelectMetasByCocktailAndMetaType(cocktail.ID, int(Type))

			cocktail.Served = SelectMetasByCocktailAndMetaType(cocktail.ID, int(Served))
			cocktail.Technique = SelectMetasByCocktailAndMetaType(cocktail.ID, int(Technique))
			cocktail.Strength = SelectMetasByCocktailAndMetaType(cocktail.ID, int(Strength))
			cocktail.Difficulty = SelectMetasByCocktailAndMetaType(cocktail.ID, int(Difficulty))
			cocktail.TOD = SelectMetasByCocktailAndMetaType(cocktail.ID, int(TOD))
			cocktail.Ratio = SelectMetasByCocktailAndMetaType(cocktail.ID, int(Ratio))
			cocktail.Family = SelectMetasByCocktailAndMetaType(cocktail.ID, int(Family))

			//add cocktail to cocktail family
			cs.Cocktail = cocktail

			log.Println(cocktail.ID, cocktail.Title, cocktail.Name, cocktail.Rating, cocktail.ImagePath,
				cocktail.Image, cocktail.Description, cocktail.Comment, cocktail.SourceName, cocktail.SourceLink, cocktail.ImageSourceName, cocktail.ImageSourceLink)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
	}
	return cs
}

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}
