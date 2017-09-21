// Copyright 2017 Rough Industries LLC. All rights reserved.
//model/cocktail.db.go:package model
package model

import (
	"bytes"
	"connectors"
	"encoding/gob"
	"github.com/bradfitz/gomemcache/memcache"
	"html"
	"html/template"
	"log"
	"strconv"
	"strings"
)

//CREATE, UPDATE, DELETE
//Insert a cocktail record into the database
func (cocktail *Cocktail) InsertCocktail() int {
	cocktail.ID = 0
	return cocktail.processCocktail()
}

//Update a cocktail record in the database based on ID
func (cocktail *Cocktail) UpdateCocktail() int {
	return cocktail.processCocktail()
}

//Process an insert or an update
func (cocktail *Cocktail) processCocktail() int {
	conn, _ := connectors.GetDB()
	var args []interface{}
	var buffer bytes.Buffer
	if cocktail.ID == 0 {
		buffer.WriteString("INSERT INTO `cocktail` SET ")
	} else {
		buffer.WriteString("UPDATE `cocktail` SET ")
	}
	if cocktail.Title != "" {
		buffer.WriteString("`cocktailTitle`=?,")
		args = append(args, cocktail.Title)
	}
	if cocktail.Name != "" {
		buffer.WriteString("`cocktailName`=?,")
		args = append(args, cocktail.Name)
	}
	if cocktail.Description != "" {
		buffer.WriteString("`cocktailDescription`=?,")
		args = append(args, html.EscapeString(string(cocktail.Description)))
	}
	if cocktail.Comment != "" {
		buffer.WriteString("`cocktailComment`=?,")
		args = append(args, html.EscapeString(string(cocktail.Comment)))
	}
	if cocktail.ImagePath != "" {
		buffer.WriteString("`cocktailImagePath`=?,")
		args = append(args, cocktail.ImagePath)
	}
	if cocktail.Image != "" {
		buffer.WriteString("`cocktailImage`=?,")
		args = append(args, cocktail.Image)
	}
	if cocktail.ImageSourceName != "" {
		buffer.WriteString("`cocktailImageSourceName`=?,")
		args = append(args, cocktail.ImageSourceName)
	}
	if cocktail.ImageSourceLink != "" {
		buffer.WriteString("`cocktailImageSourceLink`=?,")
		args = append(args, cocktail.ImageSourceLink)
	}
	if cocktail.SourceName != "" {
		buffer.WriteString("`cocktailSourceName`=?,")
		args = append(args, cocktail.SourceName)
	}
	if cocktail.SourceLink != "" {
		buffer.WriteString("`cocktailSourceLink`=?,")
		args = append(args, cocktail.SourceLink)
	}
	if cocktail.Rating != 0 {
		buffer.WriteString(" `cocktailRating`=?,")
		args = append(args, strconv.Itoa(cocktail.Rating))
	}
	if cocktail.Origin != "" {
		buffer.WriteString("`cocktailOrigin`=?,")
		args = append(args, html.EscapeString(string(cocktail.Origin)))
	}
	if cocktail.SpokenName != "" {
		buffer.WriteString("`cocktailSpokenName`=?,")
		args = append(args, cocktail.SpokenName)
	}
	if cocktail.DisplayName != "" {
		buffer.WriteString("`cocktailDisplayName`=?,")
		args = append(args, cocktail.DisplayName)
	}
	query := buffer.String()
	query = strings.TrimRight(query, ",")
	if cocktail.ID == 0 {
		query = query + ";"
	} else {
		query = query + " WHERE `idCocktail`=?;"
		args = append(args, strconv.Itoa(cocktail.ID))
	}
	log.Println(query)
	res, err := conn.Exec(query, args...)
	if err != nil {
		log.Fatal(err)
	}
	var lastCocktailId int64
	if cocktail.ID == 0 {
		lastCocktailId, err = res.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		lastCocktailId = int64(cocktail.ID)
	}

	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Cocktail ID = %d, affected = %d\n", lastCocktailId, rowCnt)

	var recipeID int
	if cocktail.ID == 0 {
		recipeID = InsertRecipe(cocktail.Recipe)
	} else {
		recipeID = UpdateRecipe(cocktail.Recipe)
	}
	cocktail.clearAltNamesAndAKAsByCocktailID(lastCocktailId)
	cocktail.processAKAs(cocktail.AKA, lastCocktailId)
	cocktail.processAltNames(cocktail.AlternateName, lastCocktailId)

	cocktail.clearCocktailToProductsByCocktailID(lastCocktailId)
	cocktail.processCocktailToProducts(cocktail.Garnish, lastCocktailId)
	cocktail.processCocktailToProducts(cocktail.Drinkware, lastCocktailId)
	cocktail.processCocktailToProducts(cocktail.Tool, lastCocktailId)

	log.Println("processing Cocktail to Metas")
	cocktail.clearCocktailToMetasByCocktailID(lastCocktailId)
	log.Println("processing Cocktail to Family")
	cocktail.processCocktailToMetas(cocktail.Family, lastCocktailId, btoi(cocktail.IsFamilyRoot))
	log.Println("processing Cocktail to Flavor")
	cocktail.processCocktailToMetas(cocktail.Flavor, lastCocktailId, 0)
	log.Println("processing Cocktail to Type")
	cocktail.processCocktailToMetas(cocktail.Type, lastCocktailId, 0)
	log.Println("processing Cocktail to BaseSpirit")
	cocktail.processCocktailToMetas(cocktail.BaseSpirit, lastCocktailId, 0)
	log.Println("processing Cocktail to Served")
	cocktail.processCocktailToMetas(cocktail.Served, lastCocktailId, 0)
	log.Println("processing Cocktail to Technique")
	cocktail.processCocktailToMetas(cocktail.Technique, lastCocktailId, 0)
	log.Println("processing Cocktail to Strength")
	cocktail.processCocktailToMetas(cocktail.Strength, lastCocktailId, 0)
	log.Println("processing Cocktail to Difficulty")
	cocktail.processCocktailToMetas(cocktail.Difficulty, lastCocktailId, 0)
	log.Println("processing Cocktail to TOD")
	cocktail.processCocktailToMetas(cocktail.TOD, lastCocktailId, 0)
	log.Println("processing Cocktail to Ratio")
	cocktail.processCocktailToMetas(cocktail.Ratio, lastCocktailId, 0)

	if cocktail.ID == 0 {
		args = args[0:0]
		args = append(args, strconv.FormatInt(lastCocktailId, 10))
		args = append(args, strconv.Itoa(recipeID))
		conn.Exec("INSERT INTO `cocktailToRecipe` (`idCocktail`, `idRecipe`) VALUES (?, ?);", args...)
	}
	return int(lastCocktailId)
}

//Delete the alt names and aka names based on ID.
func (cocktail *Cocktail) clearAltNamesAndAKAsByCocktailID(cocktailID int64) {
	conn, _ := connectors.GetDB()
	var buffer bytes.Buffer
	var args []interface{}
	var nameIDs []int
	buffer.WriteString("SELECT cocktailToAltNames.idAltName FROM cocktailToAltNames WHERE idCocktail=?;")
	args = append(args, cocktailID)
	query := buffer.String()
	rows, err := conn.Query(query, args...)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var nameID int
		err := rows.Scan(&nameID)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(nameID)
		nameIDs = append(nameIDs, nameID)
	}
	buffer.Reset()
	args = args[0:0]
	buffer.WriteString("SELECT cocktailToAKAs.idAKAName FROM cocktailToAKAs WHERE idCocktail=?;")
	args = append(args, cocktailID)
	query = buffer.String()
	rows, err = conn.Query(query, args...)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var nameID int
		err := rows.Scan(&nameID)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(nameID)
		nameIDs = append(nameIDs, nameID)
	}
	//clear all altingredients by stepid
	for _, nameID := range nameIDs {
		buffer.Reset()
		args = args[0:0]
		buffer.WriteString("DELETE FROM `altnames` WHERE `idAltNames`=?;")
		args = append(args, nameID)
		query = buffer.String()
		conn.Exec(query, args...)
	}
	buffer.Reset()
	args = args[0:0]
	buffer.WriteString("DELETE FROM `cocktailToAltNames` WHERE `idCocktail`=?;")
	args = append(args, cocktailID)
	query = buffer.String()
	conn.Exec(query, args...)
	buffer.Reset()
	args = args[0:0]
	buffer.WriteString("DELETE FROM `cocktailToAKAs` WHERE `idCocktail`=?;")
	args = append(args, cocktailID)
	query = buffer.String()
	conn.Exec(query, args...)
}

//Process an insert into aka names and the relationship between aka and cocktail
func (cocktail *Cocktail) processAKAs(names []Name, cocktailID int64) {
	conn, _ := connectors.GetDB()
	var args []interface{}
	var buffer bytes.Buffer
	for _, name := range names {
		buffer.WriteString("INSERT INTO `altnames` SET ")
		if name.Name != "" {
			buffer.WriteString("`altNamesString`=?,")
			args = append(args, name.Name)
		}
		query := buffer.String()
		query = strings.TrimRight(query, ",")
		query = query + ";"
		log.Println(query)
		res, err := conn.Exec(query, args...)
		lastAltNameId, err := res.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		args = args[0:0]
		args = append(args, strconv.FormatInt(cocktailID, 10))
		args = append(args, strconv.FormatInt(lastAltNameId, 10))
		conn.Exec("INSERT INTO `cocktailToAKAs` (`idCocktail`, `idAKAName`) VALUES (?, ?);")
	}
}

//Process an insert into alt names and the relationship between alt names
//and cocktail
func (cocktail *Cocktail) processAltNames(names []Name, cocktailID int64) {
	conn, _ := connectors.GetDB()
	var args []interface{}
	var buffer bytes.Buffer
	for _, name := range names {
		buffer.WriteString("INSERT INTO `altnames` SET ")
		if name.Name != "" {
			buffer.WriteString("`altNamesString`=?,")
			args = append(args, name.Name)
		}
		query := buffer.String()
		query = strings.TrimRight(query, ",")
		query = query + ";"
		log.Println(query)
		res, err := conn.Exec(query, args...)
		lastAltNameId, err := res.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		args = args[0:0]
		args = append(args, strconv.FormatInt(cocktailID, 10))
		args = append(args, strconv.FormatInt(lastAltNameId, 10))
		conn.Exec("INSERT INTO `cocktailToAltNames` (`idCocktail`, `idAltName`) VALUES (?, ?);", args...)
	}
}

//Deletes rows from cocktailToProducts table by cocktail id
func (cocktail *Cocktail) clearCocktailToProductsByCocktailID(cocktailID int64) {
	conn, _ := connectors.GetDB()
	var buffer bytes.Buffer
	var args []interface{}
	//delete all rows from cocktailToProducts table by cocktail id
	buffer.WriteString("DELETE FROM `cocktailToProducts` WHERE `idCocktail`=?;")
	args = append(args, cocktailID)
	query := buffer.String()
	conn.Exec(query, args...)
}

//Insert rows into cocktailToProducts table by cocktail id and a set of products
func (cocktail *Cocktail) processCocktailToProducts(products []Product, cocktailID int64) {
	conn, _ := connectors.GetDB()
	var args []interface{}
	for _, product := range products {
		prodTo := product.SelectProduct()
		if len(prodTo) > 0 {
			query := "INSERT INTO `cocktailToProducts` (`idCocktail`, `idProduct`, `idProductType`) VALUES (?, ?, ?);"
			args = args[0:0]
			args = append(args, strconv.FormatInt(cocktailID, 10))
			args = append(args, strconv.Itoa(prodTo[0].ID))
			args = append(args, strconv.Itoa(int(prodTo[0].ProductType.ID)))
			log.Println(query)
			conn.Exec(query, args...)
		}
	}
}

//Deletes rows from cocktailToMEtas table by cocktail id
func (cocktail *Cocktail) clearCocktailToMetasByCocktailID(cocktailID int64) {
	conn, _ := connectors.GetDB()
	var buffer bytes.Buffer
	var args []interface{}
	//delete all rows from cocktailToMetas table by cocktail id
	buffer.WriteString("DELETE FROM `cocktailToMetas` WHERE `idCocktail`=?;")
	args = append(args, cocktailID)
	query := buffer.String()
	conn.Exec(query, args...)
}

//Insert rows into cocktailToMetas table by cocktail id and a set of metas and
//set the root cocktail flag for the set of metas
func (cocktail *Cocktail) processCocktailToMetas(metas []Meta, cocktailID int64, isRootCocktailForMeta int) {
	conn, _ := connectors.GetDB()
	var args []interface{}
	for _, meta := range metas {
		metaTo := meta.SelectMeta()
		if len(metaTo) > 0 {
			query := "INSERT INTO `cocktailToMetas` (`idCocktail`, `idMeta`, `idMetaType`, `isRootCocktailForMeta`) VALUES (?, ?, ?, ?);"
			args = args[0:0]
			args = append(args, strconv.FormatInt(cocktailID, 10))
			args = append(args, strconv.Itoa(metaTo[0].ID))
			args = append(args, strconv.Itoa(int(metaTo[0].MetaType.ID)))
			args = append(args, strconv.Itoa(int(isRootCocktailForMeta)))
			log.Println(query)
			conn.Exec(query, args...)
		}
	}
}

//SELECTS
//Select all cocktails in alpha num order in alpha num map
func (cocktail *Cocktail) SelectCocktailsByAlphaNums(ignoreCache bool) CocktailsByAlphaNums {
	ret := new(CocktailsByAlphaNums)
	ret = nil
	if !ignoreCache {
		ret = cocktail.memcachedCocktailsByAlphaNums()
	}
	if ret == nil {
		ret = new(CocktailsByAlphaNums)
		ret.CBA = make(map[string][]Cocktail)
		conn, _ := connectors.GetDB()
		var buffer bytes.Buffer
		buffer.WriteString("SELECT cocktail.idCocktail, cocktail.cocktailTitle, cocktail.cocktailName" +
			" FROM cocktail ORDER BY cocktail.cocktailTitle;")
		query := buffer.String()
		log.Println(query)
		cba_rows, _ := conn.Query(query)
		defer cba_rows.Close()
		for cba_rows.Next() {
			var cocktail Cocktail
			err := cba_rows.Scan(&cocktail.ID, &cocktail.Title, &cocktail.Name)
			if err != nil {
				log.Fatal(err)
			}
			if _, ok := ret.CBA[string(cocktail.Title[0])]; ok {
				//append
				log.Println("Appending to " + string(cocktail.Title[0]) + " with " + cocktail.Title)
				ret.CBA[string(cocktail.Title[0])] = append(ret.CBA[string(cocktail.Title[0])], cocktail)
			} else {
				//add
				log.Println("Creating " + string(cocktail.Title[0]) + " with " + cocktail.Title)
				ret.CBA[string(cocktail.Title[0])] = []Cocktail{cocktail}
			}
		}
	}
	return *ret

}

//Memcache retrieval of cocktails in alpha num order
func (cocktail *Cocktail) memcachedCocktailsByAlphaNums() *CocktailsByAlphaNums {
	ret := new(CocktailsByAlphaNums)
	ret = nil
	mc, _ := connectors.GetMC()
	if mc != nil {
		item := new(memcache.Item)
		item, _ = mc.Get("cba")
		if item != nil {
			if len(item.Value) > 0 {
				read := bytes.NewReader(item.Value)
				dec := gob.NewDecoder(read)
				dec.Decode(&ret)
			}
		}
	}
	return ret
}

//Select a set of cocktails that have a relationship through a meta id
func (cocktail *Cocktail) SelectCocktailsByMeta(meta Meta) []Cocktail {
	var ret []Cocktail
	conn, _ := connectors.GetDB()
	var args []interface{}
	log.Println(meta.MetaName)
	var buffer bytes.Buffer
	var canQuery = false
	buffer.WriteString("SELECT cocktail.idCocktail, cocktail.cocktailTitle, cocktail.cocktailName, cocktail.cocktailRating," +
		" cocktail.cocktailImagePath, cocktail.cocktailImage, COALESCE(cocktail.cocktailDescription, ''), COALESCE(cocktail.cocktailComment, '')" +
		" FROM cocktail" +
		" JOIN cocktailToMetas ON cocktailToMetas.idCocktail=cocktail.idCocktail WHERE")
	if meta.ID != 0 {
		buffer.WriteString(" `cocktailToMetas`.`idMeta`=? AND")
		args = append(args, strconv.Itoa(meta.ID))
		canQuery = true
	}
	if canQuery {
		query := buffer.String()
		query = strings.TrimRight(query, " AND")
		query = query + ";"
		log.Println(query)
		rows, err := conn.Query(query, args...)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			var cocktail Cocktail
			var desc string
			var comment string
			err := rows.Scan(&cocktail.ID, &cocktail.Title, &cocktail.Name, &cocktail.Rating, &cocktail.ImagePath, &cocktail.Image, &desc, &comment)
			cocktail.Description = template.HTML(html.UnescapeString(desc))
			cocktail.Comment = template.HTML(html.UnescapeString(comment))
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

//Select a set of cocktails that have a relationship through a product id
func (cocktail *Cocktail) SelectCocktailsByProduct(product Product) []Cocktail {
	var ret []Cocktail
	conn, _ := connectors.GetDB()
	var args []interface{}
	log.Println(product.ProductName)
	var buffer bytes.Buffer
	var canQuery = false
	buffer.WriteString("SELECT cocktail.idCocktail, cocktail.cocktailTitle, cocktail.cocktailName, cocktail.cocktailRating," +
		" cocktail.cocktailImagePath, cocktail.cocktailImage, COALESCE(cocktail.cocktailDescription, ''), COALESCE(cocktail.cocktailComment, '')" +
		" FROM cocktail" +
		" JOIN cocktailToProducts ON cocktailToProducts.idCocktail=cocktail.idCocktail WHERE")
	if product.ID != 0 {
		buffer.WriteString(" `cocktailToProducts`.`idProduct`=? AND")
		args = append(args, strconv.Itoa(product.ID))
		canQuery = true
	}
	if int(product.ProductType.ID) != 0 {
		buffer.WriteString(" `cocktailToProducts`.`idProductType`=? AND")
		args = append(args, strconv.Itoa(int(product.ProductType.ID)))
		canQuery = true
	}
	if canQuery {
		query := buffer.String()
		query = strings.TrimRight(query, " AND")
		query = query + ";"
		log.Println(query)
		rows, err := conn.Query(query, args...)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			var cocktail Cocktail
			var desc string
			var comment string
			err := rows.Scan(&cocktail.ID, &cocktail.Title, &cocktail.Name, &cocktail.Rating, &cocktail.ImagePath, &cocktail.Image, &desc, &comment)
			cocktail.Description = template.HTML(html.UnescapeString(desc))
			cocktail.Comment = template.HTML(html.UnescapeString(comment))
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

//Select all the cocktails
func (cocktail *Cocktail) SelectAllCocktails(includeBDG bool) []Cocktail {
	var c []Cocktail
	conn, _ := connectors.GetDB()
	var buffer bytes.Buffer
	buffer.WriteString("SELECT cocktail.idCocktail, cocktail.cocktailTitle, cocktail.cocktailName, cocktail.cocktailRating," +
		" cocktail.cocktailImagePath, cocktail.cocktailImage, COALESCE(cocktail.cocktailDescription, ''), COALESCE(cocktail.cocktailComment, '')," +
		" cocktail.cocktailSourceName, cocktail.cocktailSourceLink, cocktail.cocktailImageSourceName, cocktail.cocktailImageSourceLink," +
		" COALESCE(cocktail.cocktailOrigin, ''), COALESCE(cocktail.cocktailSpokenName, ''), COALESCE(cocktail.cocktailDisplayName, '')" +
		" FROM cocktail;")
	query := buffer.String()
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
		var origin string
		err := rows.Scan(&cocktail.ID, &cocktail.Title, &cocktail.Name, &cocktail.Rating, &cocktail.ImagePath,
			&cocktail.Image, &desc, &comment, &cocktail.SourceName, &cocktail.SourceLink, &cocktail.ImageSourceName, &cocktail.ImageSourceLink,
			&origin, &cocktail.SpokenName, &cocktail.DisplayName)
		cocktail.Description = template.HTML(html.UnescapeString(desc))
		cocktail.Comment = template.HTML(html.UnescapeString(comment))
		cocktail.Origin = template.HTML(html.UnescapeString(origin))
		if err != nil {
			log.Fatal(err)
		}
		//add cocktail to cocktail family
		c = append(c, cocktail)
		log.Println(cocktail.ID, cocktail.Title, cocktail.Name, cocktail.Rating, cocktail.ImagePath,
			cocktail.Image, cocktail.Description, cocktail.Comment, cocktail.SourceName, cocktail.SourceLink, cocktail.ImageSourceName, cocktail.ImageSourceLink,
			cocktail.SpokenName, cocktail.DisplayName)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return c
}

//Select a cocktail based on an id and all associated infomration. If BDG
//flag is set set related information
func (cocktail *Cocktail) SelectCocktailsByID(ID int, includeBDG bool) CocktailSet {
	var cs CocktailSet
	var args []interface{}
	conn, _ := connectors.GetDB()
	var buffer bytes.Buffer
	buffer.WriteString("SELECT cocktail.idCocktail, cocktail.cocktailTitle, cocktail.cocktailName, cocktail.cocktailRating," +
		" cocktail.cocktailImagePath, cocktail.cocktailImage, COALESCE(cocktail.cocktailDescription, ''), COALESCE(cocktail.cocktailComment, '')," +
		" cocktail.cocktailSourceName, cocktail.cocktailSourceLink, cocktail.cocktailImageSourceName, cocktail.cocktailImageSourceLink," +
		" COALESCE(cocktail.cocktailOrigin, ''), COALESCE(cocktail.cocktailSpokenName, ''), COALESCE(cocktail.cocktailDisplayName, '')" +
		" FROM cocktail" +
		" WHERE idCocktail=?;")
	args = append(args, strconv.Itoa(ID))
	query := buffer.String()
	log.Println(query)
	rows, err := conn.Query(query, args...)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var cocktail Cocktail
		var desc string
		var comment string
		var origin string
		err := rows.Scan(&cocktail.ID, &cocktail.Title, &cocktail.Name, &cocktail.Rating, &cocktail.ImagePath,
			&cocktail.Image, &desc, &comment, &cocktail.SourceName, &cocktail.SourceLink, &cocktail.ImageSourceName, &cocktail.ImageSourceLink,
			&origin, &cocktail.SpokenName, &cocktail.DisplayName)
		cocktail.Description = template.HTML(html.UnescapeString(desc))
		cocktail.Comment = template.HTML(html.UnescapeString(comment))
		cocktail.Origin = template.HTML(html.UnescapeString(origin))
		if err != nil {
			log.Fatal(err)
		}
		meta := new(Meta)
		product := new(Product)
		//add recipe to cocktail
		cocktail.Recipe = SelectRecipeByCocktail(cocktail, includeBDG)
		cocktail.Drinkware = product.SelectProductsByCocktailAndProductType(cocktail.ID, int(Drinkware))
		cocktail.Garnish = product.SelectProductsByCocktailAndProductType(cocktail.ID, int(Garnish))
		cocktail.Tool = product.SelectProductsByCocktailAndProductType(cocktail.ID, int(Tool))
		cocktail.Flavor, _ = meta.SelectMetasByCocktailAndMetaType(cocktail.ID, int(Flavor))
		cocktail.BaseSpirit, _ = meta.SelectMetasByCocktailAndMetaType(cocktail.ID, int(BaseSpirit))
		cocktail.Type, _ = meta.SelectMetasByCocktailAndMetaType(cocktail.ID, int(Type))

		cocktail.Served, _ = meta.SelectMetasByCocktailAndMetaType(cocktail.ID, int(Served))
		cocktail.Technique, _ = meta.SelectMetasByCocktailAndMetaType(cocktail.ID, int(Technique))
		cocktail.Strength, _ = meta.SelectMetasByCocktailAndMetaType(cocktail.ID, int(Strength))
		cocktail.Difficulty, _ = meta.SelectMetasByCocktailAndMetaType(cocktail.ID, int(Difficulty))
		cocktail.TOD, _ = meta.SelectMetasByCocktailAndMetaType(cocktail.ID, int(TOD))
		cocktail.Ratio, _ = meta.SelectMetasByCocktailAndMetaType(cocktail.ID, int(Ratio))
		cocktail.Family, cocktail.IsFamilyRoot = meta.SelectMetasByCocktailAndMetaType(cocktail.ID, int(Family))

		//add cocktail to cocktail family
		cs.Cocktail = cocktail

		log.Println(cocktail.ID, cocktail.Title, cocktail.Name, cocktail.Rating, cocktail.ImagePath,
			cocktail.Image, cocktail.Description, cocktail.Comment, cocktail.SourceName, cocktail.SourceLink, cocktail.ImageSourceName, cocktail.ImageSourceLink,
			cocktail.SpokenName, cocktail.DisplayName)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return cs
}
