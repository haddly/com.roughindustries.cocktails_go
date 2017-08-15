//model/cocktail.connectors.go
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

func InsertCocktail(cocktail Cocktail) int {
	cocktail.ID = 0
	return ProcessCocktail(cocktail)
}

func UpdateCocktail(cocktail Cocktail) int {
	return ProcessCocktail(cocktail)
}

func ProcessCocktail(cocktail Cocktail) int {
	conn, _ := connectors.GetDB()
	var args []interface{}

	var buffer bytes.Buffer
	if cocktail.ID == 0 {
		buffer.WriteString("INSERT INTO `commonwealthcocktails`.`cocktail` SET ")
	} else {
		buffer.WriteString("UPDATE `commonwealthcocktails`.`cocktail` SET ")
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
	ClearAltNamesAndAKAsByCocktailID(lastCocktailId)
	ProcessAKAs(cocktail.AKA, lastCocktailId)
	ProcessAltNames(cocktail.AlternateName, lastCocktailId)

	ClearCocktailToProductsByCocktailID(lastCocktailId)
	ProcessCocktailToProducts(cocktail.Garnish, lastCocktailId)
	ProcessCocktailToProducts(cocktail.Drinkware, lastCocktailId)
	ProcessCocktailToProducts(cocktail.Tool, lastCocktailId)

	log.Println("Processing Cocktail to Metas")
	ClearCocktailToMetasByCocktailID(lastCocktailId)
	log.Println("Processing Cocktail to Family")
	ProcessCocktailToMetas(cocktail.Family, lastCocktailId, btoi(cocktail.IsFamilyRoot))
	log.Println("Processing Cocktail to Flavor")
	ProcessCocktailToMetas(cocktail.Flavor, lastCocktailId, 0)
	log.Println("Processing Cocktail to Type")
	ProcessCocktailToMetas(cocktail.Type, lastCocktailId, 0)
	log.Println("Processing Cocktail to BaseSpirit")
	ProcessCocktailToMetas(cocktail.BaseSpirit, lastCocktailId, 0)
	log.Println("Processing Cocktail to Served")
	ProcessCocktailToMetas(cocktail.Served, lastCocktailId, 0)
	log.Println("Processing Cocktail to Technique")
	ProcessCocktailToMetas(cocktail.Technique, lastCocktailId, 0)
	log.Println("Processing Cocktail to Strength")
	ProcessCocktailToMetas(cocktail.Strength, lastCocktailId, 0)
	log.Println("Processing Cocktail to Difficulty")
	ProcessCocktailToMetas(cocktail.Difficulty, lastCocktailId, 0)
	log.Println("Processing Cocktail to TOD")
	ProcessCocktailToMetas(cocktail.TOD, lastCocktailId, 0)
	log.Println("Processing Cocktail to Ratio")
	ProcessCocktailToMetas(cocktail.Ratio, lastCocktailId, 0)

	if cocktail.ID == 0 {
		args = args[0:0]
		args = append(args, strconv.FormatInt(lastCocktailId, 10))
		args = append(args, strconv.Itoa(recipeID))
		conn.Exec("INSERT INTO `commonwealthcocktails`.`cocktailToRecipe` (`idCocktail`, `idRecipe`) VALUES (?, ?);", args...)
	}
	return int(lastCocktailId)
}

func ClearAltNamesAndAKAsByCocktailID(cocktailID int64) {
	conn, _ := connectors.GetDB()

	var buffer bytes.Buffer
	var args []interface{}
	var nameIDs []int
	buffer.WriteString("SELECT cocktailToAltNames.idAltName FROM commonwealthcocktails.cocktailToAltNames WHERE idCocktail=?;")
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

	buffer.WriteString("SELECT cocktailToAKAs.idAKAName FROM commonwealthcocktails.cocktailToAKAs WHERE idCocktail=?;")
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
		buffer.WriteString("DELETE FROM `commonwealthcocktails`.`altnames` WHERE `idAltNames`=?;")
		args = append(args, nameID)
		query = buffer.String()
		conn.Exec(query, args...)
	}

	buffer.Reset()
	args = args[0:0]

	buffer.WriteString("DELETE FROM `commonwealthcocktails`.`cocktailToAltNames` WHERE `idCocktail`=?;")
	args = append(args, cocktailID)
	query = buffer.String()
	conn.Exec(query, args...)

	buffer.Reset()
	args = args[0:0]

	buffer.WriteString("DELETE FROM `commonwealthcocktails`.`cocktailToAKAs` WHERE `idCocktail`=?;")
	args = append(args, cocktailID)
	query = buffer.String()
	conn.Exec(query, args...)
}

func ProcessAKAs(names []Name, cocktailID int64) {
	conn, _ := connectors.GetDB()
	var args []interface{}
	var buffer bytes.Buffer
	for _, name := range names {
		buffer.WriteString("INSERT INTO `commonwealthcocktails`.`altnames` SET ")

		if name.Name != "" {
			buffer.WriteString("`altNamesString`=\"?\",")
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
		conn.Exec("INSERT INTO `commonwealthcocktails`.`cocktailToAKAs` (`idCocktail`, `idAKAName`) VALUES (?, ?);")
	}
}

func ProcessAltNames(names []Name, cocktailID int64) {
	conn, _ := connectors.GetDB()
	var args []interface{}

	var buffer bytes.Buffer
	for _, name := range names {
		buffer.WriteString("INSERT INTO `commonwealthcocktails`.`altnames` SET ")

		if name.Name != "" {
			buffer.WriteString("`altNamesString`=\"?\",")
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
		conn.Exec("INSERT INTO `commonwealthcocktails`.`cocktailToAltNames` (`idCocktail`, `idAltName`) VALUES (?, ?);", args...)
	}
}

func ClearCocktailToProductsByCocktailID(cocktailID int64) {
	conn, _ := connectors.GetDB()
	var buffer bytes.Buffer
	var args []interface{}

	//delete all rows from cocktailToProducts table by cocktail id
	buffer.WriteString("DELETE FROM `commonwealthcocktails`.`cocktailToProducts` WHERE `idCocktail`=?;")
	args = append(args, cocktailID)
	query := buffer.String()
	conn.Exec(query, args...)
}

func ProcessCocktailToProducts(products []Product, cocktailID int64) {
	conn, _ := connectors.GetDB()
	var args []interface{}

	for _, product := range products {
		prodTo := SelectProduct(product)
		if len(prodTo) > 0 {
			query := "INSERT INTO `commonwealthcocktails`.`cocktailToProducts` (`idCocktail`, `idProduct`, `idProductType`) VALUES (?, ?, ?);"
			args = args[0:0]
			args = append(args, strconv.FormatInt(cocktailID, 10))
			args = append(args, strconv.Itoa(prodTo[0].ID))
			args = append(args, strconv.Itoa(int(prodTo[0].ProductType.ID)))
			log.Println(query)
			conn.Exec(query, args...)
		}
	}
}

func ClearCocktailToMetasByCocktailID(cocktailID int64) {
	conn, _ := connectors.GetDB()
	var buffer bytes.Buffer
	var args []interface{}

	//delete all rows from cocktailToMetas table by cocktail id
	buffer.WriteString("DELETE FROM `commonwealthcocktails`.`cocktailToMetas` WHERE `idCocktail`=?;")
	args = append(args, cocktailID)
	query := buffer.String()
	conn.Exec(query, args...)
}

func ProcessCocktailToMetas(metas []Meta, cocktailID int64, isRootCocktailForMeta int) {
	conn, _ := connectors.GetDB()
	var args []interface{}

	for _, meta := range metas {
		metaTo := SelectMeta(meta)
		if len(metaTo) > 0 {
			query := "INSERT INTO `commonwealthcocktails`.`cocktailToMetas` (`idCocktail`, `idMeta`, `idMetaType`, `isRootCocktailForMeta`) VALUES (?, ?, ?, ?);"
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

func GetCocktailsByAlphaNums(ignoreCache bool) CocktailsByAlphaNums {
	ret := new(CocktailsByAlphaNums)
	ret = nil
	if !ignoreCache {
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
	}

	if ret == nil {
		ret = new(CocktailsByAlphaNums)
		ret.CBA = make(map[string][]Cocktail)
		conn, _ := connectors.GetDB()
		var buffer bytes.Buffer
		buffer.WriteString("SELECT cocktail.idCocktail, cocktail.cocktailTitle, cocktail.cocktailName" +
			" FROM commonwealthcocktails.cocktail ORDER BY cocktail.cocktailTitle;")
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

func SelectCocktailsByMeta(meta Meta) []Cocktail {
	var ret []Cocktail
	conn, _ := connectors.GetDB()

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

func SelectCocktailsByProduct(product Product) []Cocktail {
	var ret []Cocktail
	conn, _ := connectors.GetDB()

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
	if int(product.ProductType.ID) != 0 {
		buffer.WriteString(" `cocktailToProducts`.`idProductType`=" + strconv.Itoa(int(product.ProductType.ID)) + " AND")
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

func SelectAllCocktails(includeBDG bool) []Cocktail {
	var c []Cocktail
	conn, _ := connectors.GetDB()

	var buffer bytes.Buffer
	var canQuery = false
	buffer.WriteString("SELECT cocktail.idCocktail, cocktail.cocktailTitle, cocktail.cocktailName, cocktail.cocktailRating," +
		" cocktail.cocktailImagePath, cocktail.cocktailImage, COALESCE(cocktail.cocktailDescription, ''), COALESCE(cocktail.cocktailComment, '')," +
		" cocktail.cocktailSourceName, cocktail.cocktailSourceLink, cocktail.cocktailImageSourceName, cocktail.cocktailImageSourceLink," +
		" COALESCE(cocktail.cocktailOrigin, ''), COALESCE(cocktail.cocktailSpokenName, ''), COALESCE(cocktail.cocktailDisplayName, '')" +
		" FROM commonwealthcocktails.cocktail;")
	canQuery = true
	if canQuery {
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
	}
	return c
}

func SelectCocktailsByID(ID int, includeBDG bool) CocktailSet {
	var cs CocktailSet
	conn, _ := connectors.GetDB()

	var buffer bytes.Buffer
	var canQuery = false
	buffer.WriteString("SELECT cocktail.idCocktail, cocktail.cocktailTitle, cocktail.cocktailName, cocktail.cocktailRating," +
		" cocktail.cocktailImagePath, cocktail.cocktailImage, COALESCE(cocktail.cocktailDescription, ''), COALESCE(cocktail.cocktailComment, '')," +
		" cocktail.cocktailSourceName, cocktail.cocktailSourceLink, cocktail.cocktailImageSourceName, cocktail.cocktailImageSourceLink," +
		" COALESCE(cocktail.cocktailOrigin, ''), COALESCE(cocktail.cocktailSpokenName, ''), COALESCE(cocktail.cocktailDisplayName, '')" +
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
			//add recipe to cocktail
			cocktail.Recipe = SelectRecipeByCocktail(cocktail, includeBDG)
			cocktail.Drinkware = SelectProductsByCocktailAndProductType(cocktail.ID, int(Drinkware))
			cocktail.Garnish = SelectProductsByCocktailAndProductType(cocktail.ID, int(Garnish))
			cocktail.Tool = SelectProductsByCocktailAndProductType(cocktail.ID, int(Tool))
			cocktail.Flavor, _ = SelectMetasByCocktailAndMetaType(cocktail.ID, int(Flavor))
			cocktail.BaseSpirit, _ = SelectMetasByCocktailAndMetaType(cocktail.ID, int(BaseSpirit))
			cocktail.Type, _ = SelectMetasByCocktailAndMetaType(cocktail.ID, int(Type))

			cocktail.Served, _ = SelectMetasByCocktailAndMetaType(cocktail.ID, int(Served))
			cocktail.Technique, _ = SelectMetasByCocktailAndMetaType(cocktail.ID, int(Technique))
			cocktail.Strength, _ = SelectMetasByCocktailAndMetaType(cocktail.ID, int(Strength))
			cocktail.Difficulty, _ = SelectMetasByCocktailAndMetaType(cocktail.ID, int(Difficulty))
			cocktail.TOD, _ = SelectMetasByCocktailAndMetaType(cocktail.ID, int(TOD))
			cocktail.Ratio, _ = SelectMetasByCocktailAndMetaType(cocktail.ID, int(Ratio))
			cocktail.Family, cocktail.IsFamilyRoot = SelectMetasByCocktailAndMetaType(cocktail.ID, int(Family))

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
	}
	return cs
}

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}
