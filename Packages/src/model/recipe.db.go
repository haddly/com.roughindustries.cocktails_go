//model/recipe.connectors.go
package model

import (
	"bytes"
	"connectors"
	"html"
	"html/template"
	"log"
	"strconv"
	"strings"
)

func InsertRecipe(recipe Recipe) int {
	recipe.ID = 0
	return ProcessRecipe(recipe)
}

func UpdateRecipe(recipe Recipe) int {
	return ProcessRecipe(recipe)
}

func ProcessRecipe(recipe Recipe) int {
	conn, _ := connectors.GetDB()
	var args []interface{}

	var buffer bytes.Buffer
	if recipe.ID == 0 {
		buffer.WriteString("INSERT INTO `commonwealthcocktails`.`recipe` SET ")
	} else {
		buffer.WriteString("UPDATE `commonwealthcocktails`.`recipe` SET ")
	}
	if recipe.Method != "" {
		buffer.WriteString("`recipeMethod`=?,")
		args = append(args, html.EscapeString(string(recipe.Method)))
	}
	query := buffer.String()
	query = strings.TrimRight(query, ",")
	if recipe.ID == 0 {
		query = query + ";"
	} else {
		query = query + " WHERE `idRecipe`=?;"
		args = append(args, strconv.Itoa(recipe.ID))
	}
	log.Println(query)
	log.Println(args...)
	res, err := conn.Exec(query, args...)
	var recipeID int64
	if recipe.ID == 0 {
		recipeID, err = res.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		recipeID = int64(recipe.ID)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Recipe ID = %d, affected = %d\n", recipeID, rowCnt)
	//We don't bother trying to figure out what we can modify or not
	//the database is small enough where we can just blow away the rows
	//and insert new ones
	ClearRecipeStepsByRecipeID(recipeID)
	for _, recipestep := range recipe.RecipeSteps {
		ProcessRecipeStep(recipestep, recipeID)
	}
	return int(recipeID)
}

func ClearRecipeStepsByRecipeID(recipeID int64) {
	conn, _ := connectors.GetDB()

	var buffer bytes.Buffer
	var args []interface{}
	var rsIDs []int
	//get all recipesteps from recipetorecipestep table
	buffer.WriteString("SELECT recipeToRecipeSteps.idRecipeStep FROM commonwealthcocktails.recipeToRecipeSteps WHERE idRecipe=?;")
	args = append(args, recipeID)
	query := buffer.String()
	log.Println(query)
	log.Println(args...)
	rows, err := conn.Query(query, args...)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var rsID int
		err := rows.Scan(&rsID)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(rsID)
		rsIDs = append(rsIDs, rsID)
	}

	//clear all altingredients by stepid
	for _, stepID := range rsIDs {
		ClearAltIngredientsByRecipeStepID(int64(stepID))
	}

	buffer.Reset()
	args = args[0:0]

	//delete all rows from recipetorecipestep table by recipe id
	buffer.WriteString("DELETE FROM `commonwealthcocktails`.`recipeToRecipeSteps` WHERE `idRecipe`=?;")
	args = append(args, recipeID)
	query = buffer.String()
	log.Println(query)
	log.Println(args...)
	conn.Exec(query, args...)

	//delete all steps from recipesteps table by stepid
	for _, stepID := range rsIDs {
		buffer.Reset()
		args = args[0:0]
		buffer.WriteString("DELETE FROM `commonwealthcocktails`.`recipestep` WHERE `idRecipeStep`=?;")
		args = append(args, stepID)
		query = buffer.String()
		log.Println(query)
		log.Println(args...)
		conn.Exec(query, args...)
	}
}

func ClearAltIngredientsByRecipeStepID(stepID int64) {
	conn, _ := connectors.GetDB()

	var buffer bytes.Buffer
	var args []interface{}
	//delete all altingredients from altingredients table by stepid
	buffer.WriteString("DELETE FROM `commonwealthcocktails`.`altIngredient` WHERE `idRecipeStep`=?;")
	args = append(args, stepID)
	query := buffer.String()
	log.Println(query)
	log.Println(args...)
	conn.Exec(query, args...)
}

//This is a helper function, I recommend that you always clear out the
//recipesteps before you start to process anything
func ProcessRecipeStep(recipestep RecipeStep, recipeID int64) {
	conn, _ := connectors.GetDB()

	var buffer bytes.Buffer
	buffer.WriteString("INSERT INTO `commonwealthcocktails`.`recipestep` SET ")
	if recipestep.OriginalIngredient.ID != 0 {
		buffer.WriteString("`recipestepOriginalIngredient`=" + strconv.Itoa(recipestep.OriginalIngredient.ID) + ",")
	}
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
	stepID, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Step ID = %d, affected = %d\n", stepID, rowCnt)
	if recipestep.OriginalIngredient.ID < 1 {
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
		conn.Exec("UPDATE `commonwealthcocktails`.`recipestep` SET `recipestepOriginalIngredient`='" + strconv.Itoa(id) + "' WHERE `idRecipeStep`='" + strconv.FormatInt(stepID, 10) + "';")
	}
	for _, altingredient := range recipestep.AltIngredient {
		ProcessAltIngredient(altingredient, stepID)
	}

	ProcessRecipeToRecipeStep(recipeID, stepID)
}

//This is a helper function, I recommend that you always clear out the
//altingredients before you start to process anything
func ProcessAltIngredient(altingredient Product, stepID int64) {
	conn, _ := connectors.GetDB()

	product := SelectProduct(altingredient)
	if len(product) > 0 {
		conn.Exec("INSERT INTO `commonwealthcocktails`.`altIngredient` (`idProduct`, `idRecipeStep`) VALUES ('" + strconv.Itoa(product[0].ID) + "', '" + strconv.FormatInt(stepID, 10) + "');")
	}
}

//This is a helper function, I recommend that you always clear out the
//recipesteps before you start to process anything
func ProcessRecipeToRecipeStep(recipeID int64, stepID int64) {
	conn, _ := connectors.GetDB()
	var args []interface{}

	args = append(args, strconv.FormatInt(recipeID, 10))
	args = append(args, strconv.FormatInt(stepID, 10))
	conn.Exec("INSERT INTO `commonwealthcocktails`.`recipeToRecipeSteps` (`idRecipe`, `idRecipeStep`) VALUES (?, ?);", args...)
}

func SelectRecipeByCocktail(cocktail Cocktail, includeBDG bool) Recipe {

	var ret Recipe
	conn, _ := connectors.GetDB()

	var buffer bytes.Buffer
	var canQuery = false
	buffer.WriteString("SELECT recipe.idRecipe, recipe.recipeMethod FROM commonwealthcocktails.recipe" +
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
			err := rows.Scan(&ret.ID, &method)
			ret.Method = template.HTML(html.UnescapeString(method))
			ret.RecipeSteps = SelectRecipeStepsByCocktail(cocktail, includeBDG)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(ret.ID, ret.Method)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
	}
	return ret
}

func SelectRecipeStepsByCocktail(cocktail Cocktail, includeBDG bool) []RecipeStep {

	var ret []RecipeStep
	conn, _ := connectors.GetDB()

	var buffer bytes.Buffer
	var canQuery = false
	buffer.WriteString("SELECT recipestep.idRecipeStep, recipestep.recipestepOriginalIngredient, recipestep.recipestepRecipeCardinalFloat," +
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
			err := rows.Scan(&rs.ID, &oiID, &rs.RecipeCardinalFloat, &rs.RecipeCardinalString, &doze)
			if err != nil {
				log.Fatal(err)
			}
			rs.OriginalIngredient = SelectProductByID(oiID)
			rs.RecipeDoze = Doze{ID: doze}
			rs.AltIngredient = SelectAltIngredientsByRecipeStep(rs, includeBDG)
			log.Println(rs.ID, rs.OriginalIngredient, rs.RecipeCardinalFloat, rs.RecipeCardinalString, int(rs.RecipeDoze.ID))
			ret = append(ret, rs)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
	}
	return ret

}

func SelectAltIngredientsByRecipeStep(rs RecipeStep, includeBDG bool) []Product {
	var ret []Product
	conn, _ := connectors.GetDB()

	var buffer bytes.Buffer
	var args []interface{}
	buffer.WriteString("SELECT altIngredient.idProduct FROM commonwealthcocktails.altIngredient WHERE `idRecipeStep`=?;")
	args = append(args, rs.ID)
	query := buffer.String()
	log.Println(query)
	rows, err := conn.Query(query, args...)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var prod Product
		err := rows.Scan(&prod.ID)
		if err != nil {
			log.Fatal(err)
		}
		prod = SelectProductByID(prod.ID)
		ret = append(ret, prod)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	//TODO: Add group and derived products to the al products list
	if includeBDG {
		prod := SelectProduct(rs.OriginalIngredient)
		if len(prod) > 0 {
			bdg := GetBDGByProduct(prod[0])
			if bdg.BaseProduct.ID != 0 {
				ret = append(ret, bdg.BaseProduct)
			}
			if len(bdg.GroupProducts) > 0 {
				for _, gp := range bdg.GroupProducts {
					ret = append(ret, gp)
				}
			}
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
