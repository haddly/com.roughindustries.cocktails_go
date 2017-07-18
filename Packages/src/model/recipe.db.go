//model/recipe.connectors.go
package model

import (
	"bytes"
	"connectors"
	"database/sql"
	"html/template"
	"log"
	"strconv"
	"strings"
)

func InitRecipeTables() {
	conn, _ := connectors.GetDB()

	var temp string

	if err := conn.QueryRow("SHOW TABLES LIKE 'doze';").Scan(&temp); err == nil {
		log.Println("doze Table Exists")
	} else if err == sql.ErrNoRows {
		log.Println("Creating doze Table")
		conn.Query("CREATE TABLE `commonwealthcocktails`.`doze` (`idDoze` INT NOT NULL AUTO_INCREMENT,PRIMARY KEY (`idDoze`));")
		conn.Query("ALTER TABLE `commonwealthcocktails`.`doze` ADD COLUMN `dozeName` VARCHAR(150) NOT NULL AFTER `idDoze`;")
		conn.Exec("INSERT INTO `commonwealthcocktails`.`doze` (`idDoze`, `dozeName`) VALUES ('1', 'Shot');")
		conn.Exec("INSERT INTO `commonwealthcocktails`.`doze` (`idDoze`, `dozeName`) VALUES ('2', 'Ounce');")
		conn.Exec("INSERT INTO `commonwealthcocktails`.`doze` (`idDoze`, `dozeName`) VALUES ('3', 'Whole');")
		conn.Exec("INSERT INTO `commonwealthcocktails`.`doze` (`idDoze`, `dozeName`) VALUES ('4', 'Dash');")
		conn.Exec("INSERT INTO `commonwealthcocktails`.`doze` (`idDoze`, `dozeName`) VALUES ('5', 'Slice');")
		conn.Exec("INSERT INTO `commonwealthcocktails`.`doze` (`idDoze`, `dozeName`) VALUES ('6', 'Top Off With');")
		conn.Exec("INSERT INTO `commonwealthcocktails`.`doze` (`idDoze`, `dozeName`) VALUES ('7', 'Fresh');")
	}

	if err := conn.QueryRow("SHOW TABLES LIKE 'recipestep';").Scan(&temp); err == nil {
		log.Println("recipestep Table Exists")
	} else if err == sql.ErrNoRows {
		log.Println("Creating recipestep Table")
		conn.Query("CREATE TABLE `commonwealthcocktails`.`recipestep` (`idRecipeStep` INT NOT NULL AUTO_INCREMENT,PRIMARY KEY (`idRecipeStep`));") //ID
		conn.Query("ALTER TABLE `commonwealthcocktails`.`recipestep` " +
			"ADD COLUMN `recipestepOriginalIngredient` INT AFTER `idRecipeStep`," + //OriginalIngredient
			"ADD COLUMN `recipestepRecipeCardinalFloat` FLOAT NOT NULL AFTER `recipestepOriginalIngredient`," + //RecipeCardinalFloat
			"ADD COLUMN `recipestepRecipeCardinalString` VARCHAR(15) NOT NULL AFTER `recipestepRecipeCardinalFloat`," + //RecipeCardinalString
			"ADD COLUMN `recipestepRecipeOrdinal` INT NOT NULL AFTER `recipestepRecipeCardinalString`," + //RecipeOrdinal
			"ADD COLUMN `recipestepRecipeDoze` INT NOT NULL AFTER `recipestepRecipeOrdinal`," + //RecipeDoze
			"ADD COLUMN `recipestepAdIngredient` INT AFTER `recipestepRecipeDoze`;") //AdIngredient
	}

	if err := conn.QueryRow("SHOW TABLES LIKE 'recipe';").Scan(&temp); err == nil {
		log.Println("recipe Table Exists")
	} else if err == sql.ErrNoRows {
		log.Println("Creating recipe Table")
		conn.Query("CREATE TABLE `commonwealthcocktails`.`recipe` (`idRecipe` INT NOT NULL AUTO_INCREMENT,PRIMARY KEY (`idRecipe`));") //ID
		conn.Query("ALTER TABLE `commonwealthcocktails`.`recipe` " +
			"ADD COLUMN `recipeMethod` VARCHAR(1500) AFTER `idRecipe`;") //Method

	}
}

func InitRecipeReferences() {
	conn, _ := connectors.GetDB()

	conn.Query("ALTER TABLE `commonwealthcocktails`.`recipestep`" +
		"ADD CONSTRAINT recipestep_recipesteporiginalingredient_id_fk FOREIGN KEY(recipestepOriginalIngredient) REFERENCES product(idProduct)," +
		"ADD CONSTRAINT recipestep_recipestepadingredient_id_fk FOREIGN KEY(recipestepAdIngredient) REFERENCES product(idProduct)," +
		"ADD CONSTRAINT recipestep_recipesteprecipedoze_id_fk FOREIGN KEY(recipestepRecipeDoze) REFERENCES doze(idDoze);")
	addRecipeToRecipeStepTables()
	addRecipeStepToAltIngredientTables()
}

func addRecipeToRecipeStepTables() {
	conn, _ := connectors.GetDB()

	var temp string
	if err := conn.QueryRow("SHOW TABLES LIKE 'recipeToRecipeSteps';").Scan(&temp); err == nil {
		log.Println("recipeToRecipeSteps Table Exists")
	} else if err == sql.ErrNoRows {
		log.Println("Creating recipeToRecipeSteps Table")
		//Drink
		conn.Exec("CREATE TABLE `commonwealthcocktails`.`recipeToRecipeSteps` (`idRecipeToRecipeSteps` INT NOT NULL AUTO_INCREMENT," +
			" `idRecipe` INT NOT NULL," +
			" `idRecipeStep` INT NOT NULL," +
			" PRIMARY KEY (`idRecipeToRecipeSteps`)," +
			" CONSTRAINT recipeToRecipeSteps_idRecipe_id_fk FOREIGN KEY(idRecipe) REFERENCES recipe(idRecipe)," +
			" CONSTRAINT recipeToRecipeSteps_idRecipeStep_id_fk FOREIGN KEY(idRecipeStep) REFERENCES recipestep(idRecipeStep));")
	}
}

func addRecipeStepToAltIngredientTables() {
	conn, _ := connectors.GetDB()

	var temp string
	if err := conn.QueryRow("SHOW TABLES LIKE 'altIngredient';").Scan(&temp); err == nil {
		log.Println("altIngredient Table Exists")
	} else if err == sql.ErrNoRows {
		log.Println("Creating altIngredient Table")
		//Drink
		conn.Exec("CREATE TABLE `commonwealthcocktails`.`altIngredient` (`idAltIngredient` INT NOT NULL AUTO_INCREMENT," +
			" `idProduct` INT NOT NULL," +
			" `idRecipeStep` INT NOT NULL," +
			" PRIMARY KEY (`idAltIngredient`)," +
			" CONSTRAINT altIngredient_idProduct_id_fk FOREIGN KEY(idProduct) REFERENCES product(idProduct)," +
			" CONSTRAINT altIngredient_idRecipeStep_id_fk FOREIGN KEY(idRecipeStep) REFERENCES recipestep(idRecipeStep));")
	}
}

func ProcessRecipes() {
	for _, recipe := range Recipes {
		ProcessRecipe(recipe)
	}
}

func ProcessRecipe(recipe Recipe) int {
	conn, _ := connectors.GetDB()

	var buffer bytes.Buffer
	buffer.WriteString("INSERT INTO `commonwealthcocktails`.`recipe` SET ")
	if recipe.Method != "" {
		sqlString := strings.Replace(string(recipe.Method), "\\", "\\\\", -1)
		sqlString = strings.Replace(sqlString, "\"", "\\\"", -1)
		buffer.WriteString("`recipeMethod`=\"" + sqlString + "\",")
	}
	query := buffer.String()
	query = strings.TrimRight(query, ",")
	query = query + ";"
	log.Println(query)
	res, err := conn.Exec(query)
	lastRecipeId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Recipe ID = %d, affected = %d\n", lastRecipeId, rowCnt)

	for _, recipestep := range recipe.RecipeSteps {
		buffer.Reset()
		buffer.WriteString("INSERT INTO `commonwealthcocktails`.`recipestep` SET ")
		buffer.WriteString("`recipestepRecipeCardinalFloat`=" + strconv.FormatFloat(recipestep.RecipeCardinalFloat, 'f', -1, 32) + ",")
		if recipestep.RecipeCardinalString != "" {
			buffer.WriteString("`recipestepRecipeCardinalString`=\"" + recipestep.RecipeCardinalString + "\",")
		}
		buffer.WriteString("`recipestepRecipeOrdinal`=" + strconv.Itoa(recipestep.RecipeOrdinal) + ",")
		buffer.WriteString("`recipestepRecipeDoze`=" + strconv.Itoa(int(recipestep.RecipeDoze.ID)) + ",")
		query := buffer.String()
		query = strings.TrimRight(query, ",")
		query = query + ";"
		log.Println(query)
		res, err := conn.Exec(query)
		lastStepId, err := res.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		rowCnt, err := res.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Step ID = %d, affected = %d\n", lastStepId, rowCnt)
		rows, _ := conn.Query("SELECT idProduct, productName FROM commonwealthcocktails.product where productName = '" + recipestep.OriginalIngredient.ProductName + "';")
		defer rows.Close()
		var (
			id   int
			name string
		)
		for rows.Next() {
			err := rows.Scan(&id, &name)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(id, name)
		}
		//TODO: Add alt ingredients
		for _, altingredient := range recipestep.AltIngredient {
			product := SelectProduct(altingredient)
			if len(product) > 0 {
				conn.Exec("INSERT INTO `commonwealthcocktails`.`altIngredient` (`idProduct`, `idRecipeStep`) VALUES ('" + strconv.Itoa(product[0].ID) + "', '" + strconv.FormatInt(lastStepId, 10) + "');")
			}
		}
		conn.Exec("UPDATE `commonwealthcocktails`.`recipestep` SET `recipestepOriginalIngredient`='" + strconv.Itoa(id) + "' WHERE `idRecipeStep`='" + strconv.FormatInt(lastStepId, 10) + "';")
		conn.Exec("INSERT INTO `commonwealthcocktails`.`recipeToRecipeSteps` (`idRecipe`, `idRecipeStep`) VALUES ('" + strconv.FormatInt(lastRecipeId, 10) + "', '" + strconv.FormatInt(lastStepId, 10) + "');")
	}
	return int(lastRecipeId)
}

func SelectRecipeByCocktail(cocktail Cocktail) Recipe {

	var ret Recipe
	conn, _ := connectors.GetDB()

	var buffer bytes.Buffer
	var canQuery = false
	buffer.WriteString("SELECT recipe.recipeMethod FROM commonwealthcocktails.recipe" +
		" JOIN commonwealthcocktails.cocktailToRecipe ON cocktailToRecipe.idRecipe=recipe.idRecipe" +
		" JOIN  commonwealthcocktails.cocktail ON cocktailToRecipe.idCocktail=cocktail.idCocktail" +
		" WHERE cocktail.idCocktail=" + strconv.Itoa(cocktail.ID) + ";")
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
			var method string
			err := rows.Scan(&method)
			ret.Method = template.HTML(method)
			ret.RecipeSteps = SelectRecipeStepsByCocktail(cocktail)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(ret.Method)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
	}
	return ret
}

func SelectRecipeStepsByCocktail(cocktail Cocktail) []RecipeStep {

	var ret []RecipeStep
	conn, _ := connectors.GetDB()

	var buffer bytes.Buffer
	var canQuery = false
	buffer.WriteString("SELECT recipestep.recipestepOriginalIngredient, recipestep.recipestepRecipeCardinalFloat," +
		" COALESCE(recipestep.recipestepRecipeCardinalString, ''), recipestep.recipestepRecipeDoze" +
		" FROM commonwealthcocktails.recipestep" +
		" JOIN commonwealthcocktails.recipeToRecipeSteps ON recipeToRecipeSteps.idRecipeStep=recipestep.idRecipeStep" +
		" JOIN commonwealthcocktails.recipe ON  recipeToRecipeSteps.idRecipe=recipe.idRecipe" +
		" JOIN commonwealthcocktails.cocktailToRecipe ON cocktailToRecipe.idRecipe=recipe.idRecipe" +
		" JOIN  commonwealthcocktails.cocktail ON cocktailToRecipe.idCocktail=cocktail.idCocktail" +
		" WHERE cocktail.idCocktail=" + strconv.Itoa(cocktail.ID) + " ORDER BY recipestepRecipeOrdinal;")
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
			var rs RecipeStep
			var oiID int
			var doze int
			err := rows.Scan(&oiID, &rs.RecipeCardinalFloat, &rs.RecipeCardinalString, &doze)
			if err != nil {
				log.Fatal(err)
			}
			rs.OriginalIngredient = SelectProductByID(oiID)
			rs.RecipeDoze = Doze{ID: doze}
			log.Println(rs.OriginalIngredient, rs.RecipeCardinalFloat, rs.RecipeCardinalString, int(rs.RecipeDoze.ID))
			ret = append(ret, rs)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
	}
	return ret

}

func SelectDoze() []Doze {
	var ret []Doze
	conn, _ := connectors.GetDB()

	var buffer bytes.Buffer
	buffer.WriteString("SELECT doze.idDoze, doze.dozeName FROM commonwealthcocktails.doze;")

	query := buffer.String()
	log.Println(query)
	rows, err := conn.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var doze Doze
		err := rows.Scan(&doze.ID, &doze.DozeName)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(doze.ID, doze.DozeName)
		ret = append(ret, doze)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return ret
}
