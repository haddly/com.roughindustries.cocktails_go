// Copyright 2017 Rough Industries LLC. All rights reserved.
//model/cocktail.db.go:package model
package model

import (
	"bytes"
	"encoding/gob"
	"github.com/CommonwealthCocktails/connectors"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/golang/glog"
	"html"
	"html/template"
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

//Update cocktail images
func (cocktail *Cocktail) UpdateCocktailImages() int {
	conn, _ := connectors.GetDB()
	var args []interface{}
	var buffer bytes.Buffer
	buffer.WriteString("UPDATE `cocktail` SET ")

	buffer.WriteString("`cocktailImagePath`=?,")
	args = append(args, cocktail.ImagePath)
	buffer.WriteString("`cocktailImage`=?,")
	args = append(args, cocktail.Image)
	buffer.WriteString("`cocktailLabeledImageLink`=?,")
	args = append(args, cocktail.LabeledImageLink)
	query := buffer.String()
	query = strings.TrimRight(query, ",")
	query = query + " WHERE `idCocktail`=?;"
	args = append(args, strconv.Itoa(cocktail.ID))
	glog.Infoln(query)
	res, err := conn.Exec(query, args...)
	if err != nil {
		glog.Error(err)
	}
	lastCocktailId := int64(cocktail.ID)

	rowCnt, err := res.RowsAffected()
	if err != nil {
		glog.Error(err)
	}
	glog.Infoln("Cocktail ID = %d, affected = %d\n", lastCocktailId, rowCnt)

	return int(lastCocktailId)
}

//Process an insert or an update
func (cocktail *Cocktail) processCocktail() int {
	conn, _ := connectors.GetDB()
	var args []interface{}
	var buffer bytes.Buffer
	if cocktail.ID == 0 {
		buffer.WriteString("INSERT INTO `cocktail` ( ")
	} else {
		buffer.WriteString("UPDATE `cocktail` SET ")
	}
	if cocktail.Title != "" {
		if cocktail.ID == 0 {
			buffer.WriteString("`cocktailTitle`,")
		} else {
			buffer.WriteString("`cocktailTitle`=?,")
		}
		args = append(args, cocktail.Title)
	}
	if cocktail.Name != "" {
		if cocktail.ID == 0 {
			buffer.WriteString("`cocktailName`,")
		} else {
			buffer.WriteString("`cocktailName`=?,")
		}
		args = append(args, string(cocktail.Name))
	}
	if cocktail.Description != "" {
		if cocktail.ID == 0 {
			buffer.WriteString("`cocktailDescription`,")
		} else {
			buffer.WriteString("`cocktailDescription`=?,")
		}
		args = append(args, html.EscapeString(string(cocktail.Description)))
	}
	if cocktail.Comment != "" {
		if cocktail.ID == 0 {
			buffer.WriteString("`cocktailComment`,")
		} else {
			buffer.WriteString("`cocktailComment`=?,")
		}
		args = append(args, html.EscapeString(string(cocktail.Comment)))
	}
	if cocktail.Footnotes != "" {
		if cocktail.ID == 0 {
			buffer.WriteString("`cocktailFootnotes`,")
		} else {
			buffer.WriteString("`cocktailFootnotes`=?,")
		}
		args = append(args, html.EscapeString(string(cocktail.Footnotes)))
	}
	if cocktail.Keywords != "" {
		if cocktail.ID == 0 {
			buffer.WriteString("`cocktailKeywords`,")
		} else {
			buffer.WriteString("`cocktailKeywords`=?,")
		}
		args = append(args, html.EscapeString(string(cocktail.Keywords)))
	}
	if cocktail.ImagePath != "" {
		if cocktail.ID == 0 {
			buffer.WriteString("`cocktailImagePath`,")
		} else {
			buffer.WriteString("`cocktailImagePath`=?,")
		}
		args = append(args, cocktail.ImagePath)
	}
	if cocktail.Image != "" {
		if cocktail.ID == 0 {
			buffer.WriteString("`cocktailImage`,")
		} else {
			buffer.WriteString("`cocktailImage`=?,")
		}
		args = append(args, cocktail.Image)
	}
	if cocktail.ImageSourceName != "" {
		if cocktail.ID == 0 {
			buffer.WriteString("`cocktailImageSourceName`,")
		} else {
			buffer.WriteString("`cocktailImageSourceName`=?,")
		}
		args = append(args, cocktail.ImageSourceName)
	}
	if cocktail.ImageSourceLink != "" {
		if cocktail.ID == 0 {
			buffer.WriteString("`cocktailImageSourceLink`,")
		} else {
			buffer.WriteString("`cocktailImageSourceLink`=?,")
		}
		args = append(args, cocktail.ImageSourceLink)
	}
	if cocktail.LabeledImageLink != "" {
		if cocktail.ID == 0 {
			buffer.WriteString("`cocktailLabeledImageLink`,")
		} else {
			buffer.WriteString("`cocktailLabeledImageLink`=?,")
		}
		args = append(args, cocktail.LabeledImageLink)
	}
	if cocktail.SourceName != "" {
		if cocktail.ID == 0 {
			buffer.WriteString("`cocktailSourceName`,")
		} else {
			buffer.WriteString("`cocktailSourceName`=?,")
		}
		args = append(args, cocktail.SourceName)
	}
	if cocktail.SourceLink != "" {
		if cocktail.ID == 0 {
			buffer.WriteString("`cocktailSourceLink`,")
		} else {
			buffer.WriteString("`cocktailSourceLink`=?,")
		}
		args = append(args, cocktail.SourceLink)
	}
	if cocktail.Rating != 0 {
		if cocktail.ID == 0 {
			buffer.WriteString("`cocktailRating`,")
		} else {
			buffer.WriteString("`cocktailRating`=?,")
		}
		args = append(args, strconv.Itoa(cocktail.Rating))
	}
	if cocktail.Origin != "" {
		if cocktail.ID == 0 {
			buffer.WriteString("`cocktailOrigin`,")
		} else {
			buffer.WriteString("`cocktailOrigin`=?,")
		}
		args = append(args, html.EscapeString(string(cocktail.Origin)))
	}
	if cocktail.SpokenName != "" {
		if cocktail.ID == 0 {
			buffer.WriteString("`cocktailSpokenName`,")
		} else {
			buffer.WriteString("`cocktailSpokenName`=?,")
		}
		args = append(args, cocktail.SpokenName)
	}
	if cocktail.DisplayName != "" {
		if cocktail.ID == 0 {
			buffer.WriteString("`cocktailDisplayName`,")
		} else {
			buffer.WriteString("`cocktailDisplayName`=?,")
		}
		args = append(args, cocktail.DisplayName)
	}
	query := buffer.String()
	query = strings.TrimRight(query, ",")
	if cocktail.ID == 0 {
		vals := strings.Repeat("?,", len(args))
		vals = strings.TrimRight(vals, ",")
		query = query + ") VALUES (" + vals + ");"
	} else {
		query = query + " WHERE `idCocktail`=?;"
		args = append(args, strconv.Itoa(cocktail.ID))
	}
	glog.Infoln(query)
	res, err := conn.Exec(query, args...)
	if err != nil {
		glog.Error(err)
	}
	var lastCocktailId int64
	if cocktail.ID == 0 {
		lastCocktailId, err = res.LastInsertId()
		if err != nil {
			glog.Error(err)
		}
	} else {
		lastCocktailId = int64(cocktail.ID)
	}

	rowCnt, err := res.RowsAffected()
	if err != nil {
		glog.Error(err)
	}
	glog.Infoln("Cocktail ID = %d, affected = %d\n", lastCocktailId, rowCnt)

	var recipeID int
	if cocktail.ID == 0 {
		recipeID = InsertRecipe(cocktail.Recipe)
	} else {
		if !cocktail.IgnoreRecipeUpdate {
			recipeID = UpdateRecipe(cocktail.Recipe)
		}
	}
	cocktail.clearAltNamesAndAKAsByCocktailID(lastCocktailId)
	cocktail.processAKAs(cocktail.AKA, lastCocktailId)
	cocktail.processAltNames(cocktail.AlternateName, lastCocktailId)

	if cocktail.ID == 0 || (cocktail.ID != 0 && !cocktail.IgnoreProductUpdate) {
		cocktail.clearCocktailToProductsByCocktailID(lastCocktailId)
		cocktail.processCocktailToProducts(cocktail.Garnish, lastCocktailId)
		cocktail.processCocktailToProducts(cocktail.Drinkware, lastCocktailId)
		cocktail.processCocktailToProducts(cocktail.Tool, lastCocktailId)
	}

	if cocktail.ID == 0 || (cocktail.ID != 0 && !cocktail.IgnoreMetaUpdate) {
		glog.Infoln("processing Cocktail to Metas")
		cocktail.clearCocktailToMetasByCocktailID(lastCocktailId)
		glog.Infoln("processing Cocktail to Family")
		cocktail.processCocktailToMetas(cocktail.Family, lastCocktailId, btoi(cocktail.IsFamilyRoot))
		glog.Infoln("processing Cocktail to Flavor")
		cocktail.processCocktailToMetas(cocktail.Flavor, lastCocktailId, 0)
		glog.Infoln("processing Cocktail to Type")
		cocktail.processCocktailToMetas(cocktail.Type, lastCocktailId, 0)
		glog.Infoln("processing Cocktail to BaseSpirit")
		cocktail.processCocktailToMetas(cocktail.BaseSpirit, lastCocktailId, 0)
		glog.Infoln("processing Cocktail to Served")
		cocktail.processCocktailToMetas(cocktail.Served, lastCocktailId, 0)
		glog.Infoln("processing Cocktail to Technique")
		cocktail.processCocktailToMetas(cocktail.Technique, lastCocktailId, 0)
		glog.Infoln("processing Cocktail to Strength")
		cocktail.processCocktailToMetas(cocktail.Strength, lastCocktailId, 0)
		glog.Infoln("processing Cocktail to Difficulty")
		cocktail.processCocktailToMetas(cocktail.Difficulty, lastCocktailId, 0)
		glog.Infoln("processing Cocktail to TOD")
		cocktail.processCocktailToMetas(cocktail.TOD, lastCocktailId, 0)
		glog.Infoln("processing Cocktail to Occasion")
		cocktail.processCocktailToMetas(cocktail.Occasion, lastCocktailId, 0)
		glog.Infoln("processing Cocktail to Style")
		cocktail.processCocktailToMetas(cocktail.Style, lastCocktailId, 0)
		glog.Infoln("processing Cocktail to Ratio")
		cocktail.processCocktailToMetas(cocktail.Ratio, lastCocktailId, 0)
	}

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
		glog.Error(err)
	}
	defer rows.Close()
	for rows.Next() {
		var nameID int
		err := rows.Scan(&nameID)
		if err != nil {
			glog.Error(err)
		}
		glog.Infoln(nameID)
		nameIDs = append(nameIDs, nameID)
	}
	buffer.Reset()
	args = args[0:0]
	buffer.WriteString("SELECT cocktailToAKAs.idAKAName FROM cocktailToAKAs WHERE idCocktail=?;")
	args = append(args, cocktailID)
	query = buffer.String()
	rows, err = conn.Query(query, args...)
	if err != nil {
		glog.Error(err)
	}
	defer rows.Close()
	for rows.Next() {
		var nameID int
		err := rows.Scan(&nameID)
		if err != nil {
			glog.Error(err)
		}
		glog.Infoln(nameID)
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
		buffer.WriteString("INSERT INTO `altnames` ( ")
		if name.Name != "" {
			buffer.WriteString("`altNamesString`,")
			args = append(args, name.Name)
		}
		query := buffer.String()
		query = strings.TrimRight(query, ",")
		vals := strings.Repeat("?,", len(args))
		vals = strings.TrimRight(vals, ",")
		query = query + ") VALUES (" + vals + ");"
		glog.Infoln(query)
		res, err := conn.Exec(query, args...)
		lastAltNameId, err := res.LastInsertId()
		if err != nil {
			glog.Error(err)
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
		buffer.WriteString("INSERT INTO `altnames` ( ")
		if name.Name != "" {
			buffer.WriteString("`altNamesString`,")
			args = append(args, name.Name)
		}
		query := buffer.String()
		query = strings.TrimRight(query, ",")
		vals := strings.Repeat("?,", len(args))
		vals = strings.TrimRight(vals, ",")
		query = query + ") VALUES (" + vals + ");"
		glog.Infoln(query)
		res, err := conn.Exec(query, args...)
		lastAltNameId, err := res.LastInsertId()
		if err != nil {
			glog.Error(err)
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
			glog.Infoln(query)
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
			glog.Infoln(query)
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
		glog.Infoln(query)
		cba_rows, _ := conn.Query(query)
		defer cba_rows.Close()
		for cba_rows.Next() {
			var cocktail Cocktail
			var name string
			err := cba_rows.Scan(&cocktail.ID, &cocktail.Title, &name)
			if err != nil {
				glog.Error(err)
			}
			cocktail.Name = template.HTML(html.UnescapeString(name))
			if _, ok := ret.CBA[string(cocktail.Title[0])]; ok {
				//append
				glog.Infoln("Appending to " + string(cocktail.Title[0]) + " with " + cocktail.Title)
				ret.CBA[string(cocktail.Title[0])] = append(ret.CBA[string(cocktail.Title[0])], cocktail)
			} else {
				//add
				glog.Infoln("Creating " + string(cocktail.Title[0]) + " with " + cocktail.Title)
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
	glog.Infoln(meta.MetaName)
	var buffer bytes.Buffer
	var canQuery = false
	buffer.WriteString("SELECT cocktail.idCocktail, COALESCE(cocktail.cocktailTitle, ''), COALESCE(cocktail.cocktailName, ''), cocktail.cocktailRating," +
		" COALESCE(cocktail.cocktailImagePath, ''), COALESCE(cocktail.cocktailImage, ''), COALESCE(cocktail.cocktailDescription, ''), COALESCE(cocktail.cocktailComment, '')," +
		" COALESCE(cocktail.cocktailFootnotes, ''), COALESCE(cocktail.cocktailKeywords, ''), COALESCE(cocktailTop100Index, 0)" +
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
		glog.Infoln(query)
		rows, err := conn.Query(query, args...)
		if err != nil {
			glog.Error(err)
		}
		defer rows.Close()
		for rows.Next() {
			var cocktail Cocktail
			var name string
			var desc string
			var comment string
			var footnotes string
			err := rows.Scan(&cocktail.ID, &cocktail.Title, &name, &cocktail.Rating, &cocktail.ImagePath, &cocktail.Image, &desc, &comment, &footnotes, &cocktail.Keywords, &cocktail.Top100Index)
			cocktail.Name = template.HTML(html.UnescapeString(name))
			cocktail.Description = template.HTML(html.UnescapeString(desc))
			cocktail.Comment = template.HTML(html.UnescapeString(comment))
			cocktail.Footnotes = template.HTML(html.UnescapeString(footnotes))
			if err != nil {
				glog.Error(err)
			}
			ret = append(ret, cocktail)
			glog.Infoln(cocktail.ID, cocktail.Title, cocktail.Name, cocktail.Rating, cocktail.ImagePath, cocktail.Image, cocktail.Description, cocktail.Comment, cocktail.Footnotes, cocktail.Keywords, cocktail.Top100Index)
		}
		err = rows.Err()
		if err != nil {
			glog.Error(err)
		}
	}
	return ret
}

//Select a set of cocktails that have a relationship through a product id
func (cocktail *Cocktail) SelectCocktailsByProduct(product Product) []Cocktail {
	var ret []Cocktail
	conn, _ := connectors.GetDB()
	var args []interface{}
	glog.Infoln(product.ProductName)
	var buffer bytes.Buffer
	var canQuery = false
	buffer.WriteString("SELECT cocktail.idCocktail, COALESCE(cocktail.cocktailTitle, ''), COALESCE(cocktail.cocktailName, ''), cocktail.cocktailRating," +
		" COALESCE(cocktail.cocktailImagePath, ''), COALESCE(cocktail.cocktailImage, ''), COALESCE(cocktail.cocktailDescription, ''), COALESCE(cocktail.cocktailComment, '')," +
		" COALESCE(cocktail.cocktailFootnotes, ''), COALESCE(cocktail.cocktailKeywords, ''), COALESCE(cocktailTop100Index, 0)" +
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
		glog.Infoln(query)
		rows, err := conn.Query(query, args...)
		if err != nil {
			glog.Error(err)
		}
		defer rows.Close()
		for rows.Next() {
			var cocktail Cocktail
			var name string
			var desc string
			var comment string
			var footnotes string
			err := rows.Scan(&cocktail.ID, &cocktail.Title, &name, &cocktail.Rating, &cocktail.ImagePath, &cocktail.Image, &desc, &comment, &footnotes, &cocktail.Keywords, &cocktail.Top100Index)
			cocktail.Name = template.HTML(html.UnescapeString(name))
			cocktail.Description = template.HTML(html.UnescapeString(desc))
			cocktail.Comment = template.HTML(html.UnescapeString(comment))
			cocktail.Footnotes = template.HTML(html.UnescapeString(footnotes))
			if err != nil {
				glog.Error(err)
			}
			ret = append(ret, cocktail)
			glog.Infoln(cocktail.ID, cocktail.Title, cocktail.Name, cocktail.Rating, cocktail.ImagePath, cocktail.Image, cocktail.Description, cocktail.Comment, cocktail.Footnotes, cocktail.Keywords, cocktail.Top100Index)
		}
		err = rows.Err()
		if err != nil {
			glog.Error(err)
		}
	}
	return ret
}

//Select a set of cocktails that are in the top 100
func (cocktail *Cocktail) SelectTop100Cocktails() []Cocktail {
	var ret []Cocktail
	conn, _ := connectors.GetDB()
	var args []interface{}
	var buffer bytes.Buffer
	buffer.WriteString("SELECT cocktail.idCocktail, COALESCE(cocktail.cocktailTitle, ''), COALESCE(cocktail.cocktailName, ''), cocktail.cocktailRating," +
		" COALESCE(cocktail.cocktailImagePath, ''), COALESCE(cocktail.cocktailImage, ''), COALESCE(cocktail.cocktailDescription, ''), COALESCE(cocktail.cocktailComment, '')," +
		" COALESCE(cocktail.cocktailFootnotes, ''), COALESCE(cocktail.cocktailKeywords, ''), COALESCE(cocktailTop100Index, 0)" +
		" FROM cocktail" +
		" WHERE cocktail.cocktailTop100Index > 0 ORDER BY cocktail.cocktailTop100Index")

	query := buffer.String()
	query = query + ";"
	glog.Infoln(query)
	rows, err := conn.Query(query, args...)
	if err != nil {
		glog.Error(err)
	}
	defer rows.Close()
	for rows.Next() {
		var cocktail Cocktail
		var name string
		var desc string
		var comment string
		var footnotes string
		err := rows.Scan(&cocktail.ID, &cocktail.Title, &name, &cocktail.Rating, &cocktail.ImagePath, &cocktail.Image, &desc, &comment, &footnotes, &cocktail.Keywords, &cocktail.Top100Index)
		cocktail.Name = template.HTML(html.UnescapeString(name))
		cocktail.Description = template.HTML(html.UnescapeString(desc))
		cocktail.Comment = template.HTML(html.UnescapeString(comment))
		cocktail.Footnotes = template.HTML(html.UnescapeString(footnotes))
		if err != nil {
			glog.Error(err)
		}
		ret = append(ret, cocktail)
		glog.Infoln(cocktail.ID, cocktail.Title, cocktail.Name, cocktail.Rating, cocktail.ImagePath, cocktail.Image, cocktail.Description, cocktail.Comment, cocktail.Footnotes, cocktail.Keywords, cocktail.Top100Index)
	}
	err = rows.Err()
	if err != nil {
		glog.Error(err)
	}

	return ret
}

//Select a cocktail that has this day of year set to the int being passed in
func (cocktail *Cocktail) SelectCocktailsByDayOfYear(doy int, includeBDG bool) CocktailSet {
	glog.Infoln(doy)
	var cs CocktailSet
	var args []interface{}
	conn, _ := connectors.GetDB()
	var buffer bytes.Buffer
	buffer.WriteString("SELECT cocktail.idCocktail, COALESCE(cocktail.cocktailTitle, ''), COALESCE(cocktail.cocktailName, ''), cocktail.cocktailRating," +
		" COALESCE(cocktail.cocktailImagePath, ''), COALESCE(cocktail.cocktailImage, ''), COALESCE(cocktail.cocktailDescription, ''), COALESCE(cocktail.cocktailComment, ''), COALESCE(cocktail.cocktailFootnotes, '')," +
		" COALESCE(cocktail.cocktailSourceName, ''), COALESCE(cocktail.cocktailSourceLink, ''), COALESCE(cocktail.cocktailImageSourceName, ''), COALESCE(cocktail.cocktailImageSourceLink, '')," +
		" COALESCE(cocktail.cocktailLabeledImageLink, ''), COALESCE(cocktail.cocktailOrigin, ''), COALESCE(cocktail.cocktailSpokenName, ''), COALESCE(cocktail.cocktailDisplayName, '')," +
		" COALESCE(cocktail.cocktailKeywords, ''), COALESCE(cocktailTop100Index, 0)" +
		" FROM cocktail" +
		" WHERE cocktailOfTheDay=?;")
	args = append(args, strconv.Itoa(doy))
	query := buffer.String()
	glog.Infoln(query)
	rows, err := conn.Query(query, args...)
	if err != nil {
		glog.Error(err)
	}
	defer rows.Close()
	for rows.Next() {
		var cocktail Cocktail
		var name string
		var desc string
		var comment string
		var footnotes string
		var origin string
		err := rows.Scan(&cocktail.ID, &cocktail.Title, &name, &cocktail.Rating, &cocktail.ImagePath,
			&cocktail.Image, &desc, &comment, &footnotes, &cocktail.SourceName, &cocktail.SourceLink, &cocktail.ImageSourceName, &cocktail.ImageSourceLink,
			&cocktail.LabeledImageLink, &origin, &cocktail.SpokenName, &cocktail.DisplayName, &cocktail.Keywords, &cocktail.Top100Index)
		cocktail.Name = template.HTML(html.UnescapeString(name))
		cocktail.Description = template.HTML(html.UnescapeString(desc))
		cocktail.Comment = template.HTML(html.UnescapeString(comment))
		cocktail.Footnotes = template.HTML(html.UnescapeString(footnotes))
		cocktail.Origin = template.HTML(html.UnescapeString(origin))
		if err != nil {
			glog.Error(err)
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
		cocktail.Occasion, _ = meta.SelectMetasByCocktailAndMetaType(cocktail.ID, int(Occasion))
		cocktail.Style, _ = meta.SelectMetasByCocktailAndMetaType(cocktail.ID, int(Style))
		cocktail.Ratio, _ = meta.SelectMetasByCocktailAndMetaType(cocktail.ID, int(Ratio))
		cocktail.Family, cocktail.IsFamilyRoot = meta.SelectMetasByCocktailAndMetaType(cocktail.ID, int(Family))

		//add cocktail to cocktail family
		cs.Cocktail = cocktail

		glog.Infoln(cocktail.ID, cocktail.Title, cocktail.Name, cocktail.Rating, cocktail.ImagePath,
			cocktail.Image, cocktail.Description, cocktail.Comment, cocktail.Footnotes, cocktail.SourceName, cocktail.SourceLink, cocktail.ImageSourceName, cocktail.ImageSourceLink,
			cocktail.LabeledImageLink, cocktail.SpokenName, cocktail.DisplayName, cocktail.Keywords, cocktail.Top100Index)
	}
	err = rows.Err()
	if err != nil {
		glog.Error(err)
	}

	return cs
}

//Select a set of cocktails based on a list of product IDs
func (cocktail *Cocktail) SelectCocktailsByIngredientIDs(productIDs []string) []Cocktail {
	var ret []Cocktail
	conn, _ := connectors.GetDB()
	var args []interface{}
	var buffer bytes.Buffer
	glog.Info(productIDs)
	var canQuery = false
	var cid int
	if cocktail.ID != 0 {
		cid = cocktail.ID
	}
	buffer.WriteString("SELECT cocktail.idCocktail, COALESCE(cocktail.cocktailTitle, ''), COALESCE(cocktail.cocktailName, ''), cocktail.cocktailRating," +
		" COALESCE(cocktail.cocktailImagePath, ''), COALESCE(cocktail.cocktailImage, ''), COALESCE(cocktail.cocktailDescription, ''), COALESCE(cocktail.cocktailComment, '')," +
		" COALESCE(cocktail.cocktailFootnotes, ''), COALESCE(cocktail.cocktailKeywords, ''), COALESCE(cocktailTop100Index, 0)" +
		" FROM commonwealthcocktails.cocktail JOIN commonwealthcocktails.cocktailToRecipe ON cocktailToRecipe.idCocktail=cocktail.idCocktail" +
		" JOIN commonwealthcocktails.recipeToRecipeSteps ON recipeToRecipeSteps.idRecipe=cocktailToRecipe.idRecipe" +
		" JOIN commonwealthcocktails.recipestep ON recipeToRecipeSteps.idRecipeStep=recipestep.idRecipeStep WHERE")
	if len(productIDs) > 0 {
		buffer.WriteString(" recipestep.recipestepOriginalIngredient IN (?" + strings.Repeat(",?", len(productIDs)-1) + ")")
		for i := range productIDs {
			args = append(args, productIDs[i])
		}
		canQuery = true
	}
	buffer.WriteString("UNION SELECT cocktail.idCocktail, COALESCE(cocktail.cocktailTitle, ''), COALESCE(cocktail.cocktailName, ''), cocktail.cocktailRating," +
		" COALESCE(cocktail.cocktailImagePath, ''), COALESCE(cocktail.cocktailImage, ''), COALESCE(cocktail.cocktailDescription, ''), COALESCE(cocktail.cocktailComment, '')," +
		" COALESCE(cocktail.cocktailFootnotes, ''), COALESCE(cocktail.cocktailKeywords, ''), COALESCE(cocktailTop100Index, 0)" +
		" FROM commonwealthcocktails.cocktail JOIN commonwealthcocktails.cocktailToRecipe ON cocktailToRecipe.idCocktail=cocktail.idCocktail" +
		" JOIN commonwealthcocktails.recipeToRecipeSteps ON recipeToRecipeSteps.idRecipe=cocktailToRecipe.idRecipe" +
		" JOIN commonwealthcocktails.recipestep ON recipeToRecipeSteps.idRecipeStep=recipestep.idRecipeStep " +
		" JOIN commonwealthcocktails.derivedProduct ON recipestep.recipestepOriginalIngredient=derivedProduct.idBaseProduct WHERE")
	if len(productIDs) > 0 && canQuery {
		buffer.WriteString(" derivedProduct.idProduct IN (?" + strings.Repeat(",?", len(productIDs)-1) + ")")
		for i := range productIDs {
			args = append(args, productIDs[i])
		}
	}
	glog.Infoln(args)
	if canQuery {
		query := buffer.String()
		query = strings.TrimRight(query, " AND")
		query = query + ";"
		glog.Infoln(query)
		rows, err := conn.Query(query, args...)
		if err != nil {
			glog.Error(err)
		}
		defer rows.Close()
		for rows.Next() {
			var cocktail Cocktail
			var name string
			var desc string
			var comment string
			var footnotes string
			err := rows.Scan(&cocktail.ID, &cocktail.Title, &name, &cocktail.Rating, &cocktail.ImagePath, &cocktail.Image, &desc, &comment, &footnotes, &cocktail.Keywords, &cocktail.Top100Index)
			cocktail.Name = template.HTML(html.UnescapeString(name))
			cocktail.Description = template.HTML(html.UnescapeString(desc))
			cocktail.Comment = template.HTML(html.UnescapeString(comment))
			cocktail.Footnotes = template.HTML(html.UnescapeString(footnotes))
			if err != nil {
				glog.Error(err)
			}
			//ignore the cocktail that was passed cocktail
			if cid != cocktail.ID {
				ret = append(ret, cocktail)
				glog.Infoln(cocktail.ID, cocktail.Title, cocktail.Name, cocktail.Rating, cocktail.ImagePath, cocktail.Image, cocktail.Description, cocktail.Comment, cocktail.Footnotes, cocktail.Keywords, cocktail.Top100Index)
			}
		}
		err = rows.Err()
		if err != nil {
			glog.Error(err)
		}
	}
	return ret

}

//Select a set of cocktails that have a relationship through a product id
func (cocktail *Cocktail) SelectCocktailsByIngredientID(product Product) []Cocktail {
	return cocktail.SelectCocktailsByIngredientIDs([]string{strconv.Itoa(product.ID)})
}

//Select all the cocktails
func (cocktail *Cocktail) SelectAllCocktails(includeBDG bool) []Cocktail {
	var c []Cocktail
	conn, _ := connectors.GetDB()
	var buffer bytes.Buffer
	buffer.WriteString("SELECT cocktail.idCocktail, COALESCE(cocktail.cocktailTitle, ''), COALESCE(cocktail.cocktailName, ''), cocktail.cocktailRating," +
		" COALESCE(cocktail.cocktailImagePath, ''), COALESCE(cocktail.cocktailImage, ''), COALESCE(cocktail.cocktailDescription, ''), COALESCE(cocktail.cocktailComment, ''), COALESCE(cocktail.cocktailFootnotes, '')," +
		" COALESCE(cocktail.cocktailSourceName, ''), COALESCE(cocktail.cocktailSourceLink, ''), COALESCE(cocktail.cocktailImageSourceName, ''), COALESCE(cocktail.cocktailImageSourceLink, '')," +
		" COALESCE(cocktail.cocktailLabeledImageLink, ''), COALESCE(cocktail.cocktailOrigin, ''), COALESCE(cocktail.cocktailSpokenName, ''), COALESCE(cocktail.cocktailDisplayName, '')," +
		" COALESCE(cocktail.cocktailKeywords, ''), COALESCE(cocktailTop100Index, 0)" +
		" FROM cocktail ORDER BY cocktail.cocktailName;")
	query := buffer.String()
	glog.Infoln(query)
	rows, err := conn.Query(query)
	if err != nil {
		glog.Error(err)
	}
	defer rows.Close()
	for rows.Next() {
		var cocktail Cocktail
		var name string
		var desc string
		var comment string
		var footnotes string
		var origin string
		err := rows.Scan(&cocktail.ID, &cocktail.Title, &name, &cocktail.Rating, &cocktail.ImagePath,
			&cocktail.Image, &desc, &comment, &footnotes, &cocktail.SourceName, &cocktail.SourceLink, &cocktail.ImageSourceName, &cocktail.ImageSourceLink,
			&cocktail.LabeledImageLink, &origin, &cocktail.SpokenName, &cocktail.DisplayName, &cocktail.Keywords, &cocktail.Top100Index)
		cocktail.Name = template.HTML(html.UnescapeString(name))
		cocktail.Description = template.HTML(html.UnescapeString(desc))
		cocktail.Comment = template.HTML(html.UnescapeString(comment))
		cocktail.Footnotes = template.HTML(html.UnescapeString(footnotes))
		cocktail.Origin = template.HTML(html.UnescapeString(origin))
		if err != nil {
			glog.Error(err)
		}
		//add cocktail to cocktail family
		c = append(c, cocktail)
		glog.Infoln(cocktail.ID, cocktail.Title, cocktail.Name, cocktail.Rating, cocktail.ImagePath,
			cocktail.Image, cocktail.Description, cocktail.Comment, cocktail.Footnotes, cocktail.SourceName, cocktail.SourceLink, cocktail.ImageSourceName, cocktail.ImageSourceLink,
			cocktail.LabeledImageLink, cocktail.SpokenName, cocktail.DisplayName, cocktail.Keywords, cocktail.Top100Index)
	}
	err = rows.Err()
	if err != nil {
		glog.Error(err)
	}
	return c
}

//Select a cocktail based on an id and all associated infomration. If BDG
//flag is set set related information
func (cocktail *Cocktail) SelectCocktailsByID(ID int, includeBDG bool) CocktailSet {
	glog.Infoln(ID)
	var cs CocktailSet
	ret := cocktail.memcachedCocktailsByID(ID)
	if ret == nil {
		var args []interface{}
		conn, _ := connectors.GetDB()
		var buffer bytes.Buffer
		buffer.WriteString("SELECT cocktail.idCocktail, COALESCE(cocktail.cocktailTitle, ''), COALESCE(cocktail.cocktailName, ''), cocktail.cocktailRating," +
			" COALESCE(cocktail.cocktailImagePath, ''), COALESCE(cocktail.cocktailImage, ''), COALESCE(cocktail.cocktailDescription, ''), COALESCE(cocktail.cocktailComment, ''), COALESCE(cocktail.cocktailFootnotes, '')," +
			" COALESCE(cocktail.cocktailSourceName, ''), COALESCE(cocktail.cocktailSourceLink, ''), COALESCE(cocktail.cocktailImageSourceName, ''), COALESCE(cocktail.cocktailImageSourceLink, '')," +
			" COALESCE(cocktail.cocktailLabeledImageLink, ''), COALESCE(cocktail.cocktailOrigin, ''), COALESCE(cocktail.cocktailSpokenName, ''), COALESCE(cocktail.cocktailDisplayName, '')," +
			" COALESCE(cocktail.cocktailKeywords, ''), COALESCE(cocktailTop100Index, 0)," +
			" recipe.idRecipe, COALESCE(recipe.recipeMethod, '')," +
			" recipestep.idRecipeStep, recipestep.recipestepOriginalIngredient, recipestep.recipestepRecipeOrdinal, recipestep.recipestepRecipeCardinalFloat," +
			" COALESCE(recipestep.recipestepRecipeCardinalString, ''), recipestep.recipestepRecipeDoze," +
			" product.idProduct, product.productName, product.productType, product.productGroupType," +
			" COALESCE(product.productDescription, ''), COALESCE(product.productDetails, '')," +
			" COALESCE(product.productImageSourceName, ''), COALESCE(product.productImage, '')," +
			" COALESCE(product.productImagePath, ''), COALESCE(product.productImageSourceLink, '')," +
			" COALESCE(product.productLabeledImageLink, ''), COALESCE(product.productPreText, '')," +
			" COALESCE(product.productPostText, ''), COALESCE(product.productRating, 0)," +
			" COALESCE(product.productSourceName, ''), COALESCE(product.productSourceLink, '')," +
			" COALESCE(product.productAmazonLink, '') " +
			" FROM recipestep" +
			" JOIN product ON recipestep.recipestepOriginalIngredient=product.idProduct " +
			" JOIN recipeToRecipeSteps ON recipeToRecipeSteps.idRecipeStep=recipestep.idRecipeStep" +
			" JOIN recipe ON  recipeToRecipeSteps.idRecipe=recipe.idRecipe" +
			" JOIN cocktailToRecipe ON cocktailToRecipe.idRecipe=recipe.idRecipe" +
			" JOIN  cocktail ON cocktailToRecipe.idCocktail=cocktail.idCocktail" +
			" WHERE cocktail.idCocktail=? ORDER BY recipestepRecipeOrdinal;")
		args = append(args, strconv.Itoa(ID))
		query := buffer.String()
		glog.Infoln(query)
		rows, err := conn.Query(query, args...)
		if err != nil {
			glog.Error(err)
		}
		defer rows.Close()
		var cocktail Cocktail
		var r Recipe
		for rows.Next() {
			var name string
			var desc string
			var comment string
			var footnotes string
			var origin string
			var method string
			var rs RecipeStep
			var oiID int
			var doze int
			var prod Product
			var prod_name string
			var prod_desc string
			var details string
			err := rows.Scan(&cocktail.ID, &cocktail.Title, &name, &cocktail.Rating, &cocktail.ImagePath,
				&cocktail.Image, &desc, &comment, &footnotes, &cocktail.SourceName, &cocktail.SourceLink, &cocktail.ImageSourceName, &cocktail.ImageSourceLink,
				&cocktail.LabeledImageLink, &origin, &cocktail.SpokenName, &cocktail.DisplayName, &cocktail.Keywords, &cocktail.Top100Index,
				&r.ID, &method, &rs.ID, &oiID, &rs.RecipeOrdinal, &rs.RecipeCardinalFloat, &rs.RecipeCardinalString, &doze,
				&prod.ID, &prod_name, &prod.ProductType.ID, &prod.ProductGroupType, &prod_desc, &details, &prod.ImageSourceName,
				&prod.Image, &prod.ImagePath, &prod.ImageSourceLink, &prod.LabeledImageLink, &prod.PreText, &prod.PostText,
				&prod.Rating, &prod.SourceName, &prod.SourceLink, &prod.AmazonLink)
			prod.ProductName = template.HTML(html.UnescapeString(prod_name))
			prod.Description = template.HTML(html.UnescapeString(prod_desc))
			prod.Details = template.HTML(html.UnescapeString(details))
			rs.OriginalIngredient = prod
			rs.RecipeDoze = Doze{ID: doze}
			r.RecipeSteps = append(r.RecipeSteps, rs)
			r.Method = template.HTML(html.UnescapeString(method))
			//add recipe to cocktail
			//cocktail.Recipe = SelectRecipeByCocktail(cocktail, includeBDG)
			cocktail.Recipe = r
			cocktail.Name = template.HTML(html.UnescapeString(name))
			cocktail.Description = template.HTML(html.UnescapeString(desc))
			cocktail.Comment = template.HTML(html.UnescapeString(comment))
			cocktail.Footnotes = template.HTML(html.UnescapeString(footnotes))
			cocktail.Origin = template.HTML(html.UnescapeString(origin))
			if err != nil {
				glog.Error(err)
			}
		}
		err = rows.Err()
		if err != nil {
			glog.Error(err)
		}

		glog.Infoln(cocktail.ID, cocktail.Title, cocktail.Name, cocktail.Rating, cocktail.ImagePath,
			cocktail.Image, cocktail.Description, cocktail.Comment, cocktail.Footnotes, cocktail.SourceName, cocktail.SourceLink, cocktail.ImageSourceName, cocktail.ImageSourceLink,
			cocktail.LabeledImageLink, cocktail.SpokenName, cocktail.DisplayName, cocktail.Keywords, cocktail.Top100Index)

		meta := new(Meta)
		product := new(Product)

		products := product.SelectProductsByCocktail(cocktail.ID)
		for _, e := range products {
			if e.ProductType.ID == Drinkware {
				cocktail.Drinkware = append(cocktail.Drinkware, e)
			} else if e.ProductType.ID == Garnish {
				cocktail.Garnish = append(cocktail.Garnish, e)
			} else if e.ProductType.ID == Tool {
				cocktail.Tool = append(cocktail.Tool, e)
			}
		}

		metas := meta.SelectMetasByCocktail(cocktail.ID)
		for _, e := range metas {
			if e.MetaType.ID == Flavor {
				cocktail.Flavor = append(cocktail.Flavor, e)
			} else if e.MetaType.ID == BaseSpirit {
				cocktail.BaseSpirit = append(cocktail.BaseSpirit, e)
			} else if e.MetaType.ID == Type {
				cocktail.Type = append(cocktail.Type, e)
			} else if e.MetaType.ID == Served {
				cocktail.Served = append(cocktail.Served, e)
			} else if e.MetaType.ID == Technique {
				cocktail.Technique = append(cocktail.Technique, e)
			} else if e.MetaType.ID == Strength {
				cocktail.Strength = append(cocktail.Strength, e)
			} else if e.MetaType.ID == Difficulty {
				cocktail.Difficulty = append(cocktail.Difficulty, e)
			} else if e.MetaType.ID == TOD {
				cocktail.TOD = append(cocktail.TOD, e)
			} else if e.MetaType.ID == Occasion {
				cocktail.Occasion = append(cocktail.Occasion, e)
			} else if e.MetaType.ID == Style {
				cocktail.Style = append(cocktail.Style, e)
			} else if e.MetaType.ID == Ratio {
				cocktail.Ratio = append(cocktail.Ratio, e)
			} else if e.MetaType.ID == Family {
				cocktail.Family = append(cocktail.Family, e)
				for _, f := range cocktail.Family {
					if f.IsRoot {
						cocktail.IsFamilyRoot = true
					}
				}
			}
		}
		cs.Cocktail = cocktail

	} else {
		cs = *ret
	}
	return cs
}

//Memcache retrieval of cocktail by id
func (cocktail *Cocktail) memcachedCocktailsByID(ID int) *CocktailSet {
	var cs CocktailSet
	var cocktails Cocktail
	mc, _ := connectors.GetMC()
	if mc != nil {
		item := new(memcache.Item)
		item, _ = mc.Get("cocktail_" + strconv.Itoa(ID))
		if item != nil {
			if len(item.Value) > 0 {
				read := bytes.NewReader(item.Value)
				dec := gob.NewDecoder(read)
				dec.Decode(&cocktails)
			}
			cs.Cocktail = cocktails
		} else {
			return nil
		}
		// var found bool
		// found = false
		// for _, ct := range cocktails.List {
		// 	if ct.ID == ID {
		// 		//add cocktail to cocktail family
		// 		cs.Cocktail = ct
		// 		found = true
		// 		break
		// 	}
		// }
		// if !found {
		// 	return nil
		// }
	}
	return &cs
}
