// Copyright 2017 Rough Industries LLC. All rights reserved.
//model/recipe.db.go:package model
package model

import (
	"bytes"
	"CommonwealthCocktails/connectors"
	"html"
	"html/template"
	"github.com/golang/glog"
	"strconv"
	"strings"
)

//CREATE, UPDATE, DELETE
//Insert a recipe record into the database
func InsertRecipe(recipe Recipe) int {
	recipe.ID = 0
	return processRecipe(recipe)
}

//Update a recipe record into the database.
func UpdateRecipe(recipe Recipe) int {
	return processRecipe(recipe)
}

//This will process inserts and updates. Updates delete relationships and
//underlying recipe steps, and then add them back in
func processRecipe(recipe Recipe) int {
	conn, _ := connectors.GetDB()
	var args []interface{}
	var buffer bytes.Buffer
	if recipe.ID == 0 {
		buffer.WriteString("INSERT INTO `recipe` SET ")
	} else {
		buffer.WriteString("UPDATE `recipe` SET ")
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
	glog.Infoln(query)
	glog.Infoln(args...)
	res, err := conn.Exec(query, args...)
	var recipeID int64
	if recipe.ID == 0 {
		recipeID, err = res.LastInsertId()
		if err != nil {
			glog.Error(err)
		}
	} else {
		recipeID = int64(recipe.ID)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		glog.Error(err)
	}
	glog.Infoln("Recipe ID = %d, affected = %d\n", recipeID, rowCnt)
	//We don't bother trying to figure out what we can modify or not
	//the database is small enough where we can just blow away the rows
	//and insert new ones
	clearRecipeStepsByRecipeID(recipeID)
	for _, recipestep := range recipe.RecipeSteps {
		processRecipeStep(recipestep, recipeID)
	}
	return int(recipeID)
}

//Delete the recipe steps and relationships
func clearRecipeStepsByRecipeID(recipeID int64) {
	conn, _ := connectors.GetDB()
	var buffer bytes.Buffer
	var args []interface{}
	var rsIDs []int
	//get all recipesteps from recipetorecipestep table
	buffer.WriteString("SELECT recipeToRecipeSteps.idRecipeStep FROM recipeToRecipeSteps WHERE idRecipe=?;")
	args = append(args, recipeID)
	query := buffer.String()
	glog.Infoln(query)
	glog.Infoln(args...)
	rows, err := conn.Query(query, args...)
	if err != nil {
		glog.Error(err)
	}
	defer rows.Close()
	for rows.Next() {
		var rsID int
		err := rows.Scan(&rsID)
		if err != nil {
			glog.Error(err)
		}
		glog.Infoln(rsID)
		rsIDs = append(rsIDs, rsID)
	}

	//clear all altingredients by stepid
	for _, stepID := range rsIDs {
		clearAltIngredientsByRecipeStepID(int64(stepID))
	}

	buffer.Reset()
	args = args[0:0]

	//delete all rows from recipetorecipestep table by recipe id
	buffer.WriteString("DELETE FROM `recipeToRecipeSteps` WHERE `idRecipe`=?;")
	args = append(args, recipeID)
	query = buffer.String()
	glog.Infoln(query)
	glog.Infoln(args...)
	conn.Exec(query, args...)

	//delete all steps from recipesteps table by stepid
	for _, stepID := range rsIDs {
		buffer.Reset()
		args = args[0:0]
		buffer.WriteString("DELETE FROM `recipestep` WHERE `idRecipeStep`=?;")
		args = append(args, stepID)
		query = buffer.String()
		glog.Infoln(query)
		glog.Infoln(args...)
		conn.Exec(query, args...)
	}
}

//Delete the alternate ingredient relationships
func clearAltIngredientsByRecipeStepID(stepID int64) {
	conn, _ := connectors.GetDB()
	var buffer bytes.Buffer
	var args []interface{}
	//delete all altingredients from altingredients table by stepid
	buffer.WriteString("DELETE FROM `altIngredient` WHERE `idRecipeStep`=?;")
	args = append(args, stepID)
	query := buffer.String()
	glog.Infoln(query)
	glog.Infoln(args...)
	conn.Exec(query, args...)
}

//This is a helper function, I recommend that you always clear out the
//recipesteps before you start to process anything. This function adds
//recipe steps to the database
func processRecipeStep(recipestep RecipeStep, recipeID int64) {
	conn, _ := connectors.GetDB()
	var args []interface{}
	var buffer bytes.Buffer
	buffer.WriteString("INSERT INTO `recipestep` SET ")
	if recipestep.OriginalIngredient.ID != 0 {
		buffer.WriteString("`recipestepOriginalIngredient`=?,")
		args = append(args, strconv.Itoa(recipestep.OriginalIngredient.ID))
	}
	buffer.WriteString("`recipestepRecipeCardinalFloat`=?,")
	args = append(args, strconv.FormatFloat(recipestep.RecipeCardinalFloat, 'f', -1, 32))
	if recipestep.RecipeCardinalString != "" {
		buffer.WriteString("`recipestepRecipeCardinalString`=?,")
		args = append(args, recipestep.RecipeCardinalString)
	}
	buffer.WriteString("`recipestepRecipeOrdinal`=?,")
	args = append(args, strconv.Itoa(recipestep.RecipeOrdinal))
	buffer.WriteString("`recipestepRecipeDoze`=?,")
	args = append(args, strconv.Itoa(int(recipestep.RecipeDoze.ID)))
	query := buffer.String()
	query = strings.TrimRight(query, ",")
	query = query + ";"
	glog.Infoln(query)
	res, err := conn.Exec(query, args...)
	stepID, err := res.LastInsertId()
	if err != nil {
		glog.Error(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		glog.Error(err)
	}
	glog.Infoln("Step ID = %d, affected = %d\n", stepID, rowCnt)
	if recipestep.OriginalIngredient.ID < 1 {
		rows, _ := conn.Query("SELECT idProduct, productName FROM product where productName=?;", recipestep.OriginalIngredient.ProductName)
		defer rows.Close()
		var (
			id   int
			name string
		)
		for rows.Next() {
			err := rows.Scan(&id, &name)
			if err != nil {
				glog.Error(err)
			}
			glog.Infoln(id, name)
		}
		conn.Exec("UPDATE `recipestep` SET `recipestepOriginalIngredient`=? WHERE `idRecipeStep`=?;", strconv.Itoa(id), strconv.FormatInt(stepID, 10))
	}
	for _, altingredient := range recipestep.AltIngredient {
		processAltIngredient(altingredient, stepID)
	}

	processRecipeToRecipeStep(recipeID, stepID)
}

//This is a helper function, I recommend that you always clear out the
//altingredients before you start to process anything.  This function
//adds alternate ingredients to the database.
func processAltIngredient(altingredient Product, stepID int64) {
	conn, _ := connectors.GetDB()
	var args []interface{}
	product := altingredient.SelectProduct()
	if len(product) > 0 {
		args = append(args, strconv.Itoa(product[0].ID))
		args = append(args, strconv.FormatInt(stepID, 10))
		conn.Exec("INSERT INTO `altIngredient` (`idProduct`, `idRecipeStep`) VALUES (?, ?);", args...)
	}
}

//This is a helper function, I recommend that you always clear out the
//recipesteps before you start to process anything.  This function
//adds the relationship between recipe steps and recipe records.
func processRecipeToRecipeStep(recipeID int64, stepID int64) {
	conn, _ := connectors.GetDB()
	var args []interface{}
	args = append(args, strconv.FormatInt(recipeID, 10))
	args = append(args, strconv.FormatInt(stepID, 10))
	conn.Exec("INSERT INTO `recipeToRecipeSteps` (`idRecipe`, `idRecipeStep`) VALUES (?, ?);", args...)
}

//SELECTS
//Get the recipe that is associated with a cocktail id.  The include BDG flag
//adds the group related information
func SelectRecipeByCocktail(cocktail Cocktail, includeBDG bool) Recipe {
	var ret Recipe
	conn, _ := connectors.GetDB()
	var args []interface{}
	var buffer bytes.Buffer
	buffer.WriteString("SELECT recipe.idRecipe, recipe.recipeMethod FROM recipe" +
		" JOIN cocktailToRecipe ON cocktailToRecipe.idRecipe=recipe.idRecipe" +
		" JOIN  cocktail ON cocktailToRecipe.idCocktail=cocktail.idCocktail" +
		" WHERE cocktail.idCocktail=?;")
	args = append(args, strconv.Itoa(cocktail.ID))
	query := buffer.String()
	glog.Infoln(query)
	rows, err := conn.Query(query, args...)
	if err != nil {
		glog.Error(err)
	}
	defer rows.Close()
	for rows.Next() {
		var method string
		err := rows.Scan(&ret.ID, &method)
		ret.Method = template.HTML(html.UnescapeString(method))
		ret.RecipeSteps = SelectRecipeStepsByCocktail(cocktail, includeBDG)
		if err != nil {
			glog.Error(err)
		}
		glog.Infoln(ret.ID, ret.Method)
	}
	err = rows.Err()
	if err != nil {
		glog.Error(err)
	}

	return ret
}

//Get the recipe steps that make up the cocktail.  The include BDG flag
//adds the group related information
func SelectRecipeStepsByCocktail(cocktail Cocktail, includeBDG bool) []RecipeStep {
	var ret []RecipeStep
	conn, _ := connectors.GetDB()
	var args []interface{}
	var buffer bytes.Buffer
	buffer.WriteString("SELECT recipestep.idRecipeStep, recipestep.recipestepOriginalIngredient, recipestep.recipestepRecipeCardinalFloat," +
		" COALESCE(recipestep.recipestepRecipeCardinalString, ''), recipestep.recipestepRecipeDoze" +
		" FROM recipestep" +
		" JOIN recipeToRecipeSteps ON recipeToRecipeSteps.idRecipeStep=recipestep.idRecipeStep" +
		" JOIN recipe ON  recipeToRecipeSteps.idRecipe=recipe.idRecipe" +
		" JOIN cocktailToRecipe ON cocktailToRecipe.idRecipe=recipe.idRecipe" +
		" JOIN  cocktail ON cocktailToRecipe.idCocktail=cocktail.idCocktail" +
		" WHERE cocktail.idCocktail=? ORDER BY recipestepRecipeOrdinal;")
	args = append(args, strconv.Itoa(cocktail.ID))
	query := buffer.String()
	glog.Infoln(query)
	rows, err := conn.Query(query, args...)
	if err != nil {
		glog.Error(err)
	}
	defer rows.Close()
	for rows.Next() {
		var rs RecipeStep
		var oiID int
		var doze int
		err := rows.Scan(&rs.ID, &oiID, &rs.RecipeCardinalFloat, &rs.RecipeCardinalString, &doze)
		if err != nil {
			glog.Error(err)
		}
		product := new(Product)
		rs.OriginalIngredient = product.SelectProductByID(oiID)
		rs.RecipeDoze = Doze{ID: doze}
		rs.AltIngredient = SelectAltIngredientsByRecipeStep(rs, includeBDG)
		glog.Infoln(rs.ID, rs.OriginalIngredient, rs.RecipeCardinalFloat, rs.RecipeCardinalString, int(rs.RecipeDoze.ID))
		ret = append(ret, rs)
	}
	err = rows.Err()
	if err != nil {
		glog.Error(err)
	}

	return ret

}

//Get the alternate ingredients that are associated with a recipe step.
//The include BDG flag adds the group related information
func SelectAltIngredientsByRecipeStep(rs RecipeStep, includeBDG bool) []Product {
	var ret []Product
	conn, _ := connectors.GetDB()
	var buffer bytes.Buffer
	var args []interface{}
	buffer.WriteString("SELECT altIngredient.idProduct FROM altIngredient WHERE `idRecipeStep`=?;")
	args = append(args, rs.ID)
	query := buffer.String()
	glog.Infoln(query)
	rows, err := conn.Query(query, args...)
	if err != nil {
		glog.Error(err)
	}
	defer rows.Close()
	for rows.Next() {
		var prod Product
		err := rows.Scan(&prod.ID)
		if err != nil {
			glog.Error(err)
		}
		prod = prod.SelectProductByID(prod.ID)
		ret = append(ret, prod)
	}
	err = rows.Err()
	if err != nil {
		glog.Error(err)
	}

	//TODO: Add group and derived products to the al products list
	if includeBDG {
		prod := rs.OriginalIngredient.SelectProduct()
		if len(prod) > 0 {
			bdg := prod[0].SelectBDGByProduct()
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

//Get all the doze records
func SelectDoze() []Doze {
	var ret []Doze
	conn, _ := connectors.GetDB()
	var buffer bytes.Buffer
	buffer.WriteString("SELECT doze.idDoze, doze.dozeName FROM doze;")
	query := buffer.String()
	glog.Infoln(query)
	rows, err := conn.Query(query)
	if err != nil {
		glog.Error(err)
	}
	defer rows.Close()
	for rows.Next() {
		var doze Doze
		err := rows.Scan(&doze.ID, &doze.DozeName)
		if err != nil {
			glog.Error(err)
		}
		glog.Infoln(doze.ID, doze.DozeName)
		ret = append(ret, doze)
	}
	err = rows.Err()
	if err != nil {
		glog.Error(err)
	}

	return ret
}
