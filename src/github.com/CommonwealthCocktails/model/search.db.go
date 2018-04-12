// Copyright 2017 Rough Industries LLC. All rights reserved.
//model/search.db.go:package model
package model

import (
	"bytes"
	"fmt"
	"github.com/CommonwealthCocktails/connectors"
	log "github.com/sirupsen/logrus"
	"html"
	"html/template"
	"strconv"
	"strings"
	"unicode"
)

//SELECTS
func (search *Search) SearchForCoctails(site string) []Cocktail {
	c := make([]Cocktail, 0)
	if len(search.Exclude_Ingredients) > 0 {
		ingEx := make([]Cocktail, 0)
		ingEx = append(ingEx, getIngredientExcludeSQL(search, site)...)
		if len(c) != 0 {
			tmp := make([]Cocktail, len(c))
			copy(tmp, c)
			c = c[0:0]
			for _, ingExItem := range ingEx {
				found := false
				for _, cItem := range tmp {
					if cItem.ID == ingExItem.ID {
						found = true
						break
					}
				}
				if found {
					c = append(c, ingExItem)
				}
			}
		} else {
			c = make([]Cocktail, len(ingEx))
			copy(c, ingEx)
		}
	}
	if len(search.Exclude_NonIngredients) > 0 {
		nonIngEx := make([]Cocktail, 0)
		nonIngEx = append(nonIngEx, getNonIngredientExcludeSQL(search, site)...)
		if len(c) != 0 {
			tmp := make([]Cocktail, len(c))
			copy(tmp, c)
			c = c[0:0]
			for _, nonIngExItem := range nonIngEx {
				found := false
				for _, cItem := range tmp {
					if cItem.ID == nonIngExItem.ID {
						found = true
						break
					}
				}
				if found {
					c = append(c, nonIngExItem)
				}
			}
		} else {
			c = make([]Cocktail, len(nonIngEx))
			copy(c, nonIngEx)
		}
	}
	if len(search.Exclude_Metas) > 0 {
		metaEx := make([]Cocktail, 0)
		metaEx = append(metaEx, getMetaExcludeSQL(search, site)...)
		if len(c) != 0 {
			tmp := make([]Cocktail, len(c))
			copy(tmp, c)
			c = c[0:0]
			for _, metaExItem := range metaEx {
				found := false
				for _, cItem := range tmp {
					if cItem.ID == metaExItem.ID {
						found = true
						break
					}
				}
				if found {
					c = append(c, metaExItem)
				}
			}
		} else {
			c = make([]Cocktail, len(metaEx))
			copy(c, metaEx)
		}
	}
	if len(search.Include_Ingredients) > 0 {
		ingEx := make([]Cocktail, 0)
		ingEx = append(ingEx, getIngredientIncludeSQL(search, site)...)
		if len(c) != 0 {
			tmp := make([]Cocktail, len(c))
			copy(tmp, c)
			c = c[0:0]
			for _, ingExItem := range ingEx {
				found := false
				for _, cItem := range tmp {
					if cItem.ID == ingExItem.ID {
						found = true
						break
					}
				}
				if found {
					c = append(c, ingExItem)
				}
			}
		} else {
			c = make([]Cocktail, len(ingEx))
			copy(c, ingEx)
		}
	}
	if len(search.Include_Metas) > 0 {
		metaIn := make([]Cocktail, 0)
		metaIn = append(metaIn, getMetaIncludeSQL(search, site)...)
		if len(c) != 0 {
			tmp := make([]Cocktail, len(c))
			copy(tmp, c)
			c = c[0:0]
			for _, metaInItem := range metaIn {
				found := false
				for _, cItem := range tmp {
					if cItem.ID == metaInItem.ID {
						found = true
						break
					}
				}
				if found {
					c = append(c, metaInItem)
				}
			}
		} else {
			c = make([]Cocktail, len(metaIn))
			copy(c, metaIn)
		}
	}
	if len(search.Include_NonIngredients) > 0 {
		nonIngEx := make([]Cocktail, 0)
		nonIngEx = append(nonIngEx, getNonIngredientIncludeSQL(search, site)...)
		if len(c) != 0 {
			tmp := make([]Cocktail, len(c))
			copy(tmp, c)
			c = c[0:0]
			for _, nonIngExItem := range nonIngEx {
				found := false
				for _, cItem := range tmp {
					if cItem.ID == nonIngExItem.ID {
						found = true
						break
					}
				}
				if found {
					c = append(c, nonIngExItem)
				}
			}
		} else {
			c = make([]Cocktail, len(nonIngEx))
			copy(c, nonIngEx)
		}
	}
	if len(search.Keywords) > 0 {
		keywords := make([]Cocktail, 0)
		keywords = append(keywords, getKeywordSQL(search, site)...)
		if len(c) != 0 {
			tmp := make([]Cocktail, len(c))
			copy(tmp, c)
			c = c[0:0]
			for _, nonKeywordsItem := range keywords {
				found := false
				for _, cItem := range tmp {
					if cItem.ID == nonKeywordsItem.ID {
						found = true
						break
					}
				}
				if found {
					c = append(c, nonKeywordsItem)
				}
			}
		} else {
			c = make([]Cocktail, len(keywords))
			copy(c, keywords)
		}
	}
	return c
}

func getKeywordSQL(search *Search, site string) []Cocktail {
	var buffer bytes.Buffer
	var args []interface{}
	var c []Cocktail
	conn, _ := connectors.GetDBFromMap(site)
	str := strings.ToLower(search.Keywords)
	buffer.WriteString("SELECT cocktail.idCocktail, COALESCE(cocktail.cocktailTitle, ''), COALESCE(cocktail.cocktailName, ''), cocktail.cocktailRating," +
		" COALESCE(cocktail.cocktailImagePath, ''), COALESCE(cocktail.cocktailImage, ''), COALESCE(cocktail.cocktailDescription, ''), COALESCE(cocktail.cocktailComment, ''), COALESCE(cocktail.cocktailFootnotes, '')," +
		" COALESCE(cocktail.cocktailSourceName, ''), COALESCE(cocktail.cocktailSourceLink, ''), COALESCE(cocktail.cocktailImageSourceName, ''), COALESCE(cocktail.cocktailImageSourceLink, '')," +
		" COALESCE(cocktail.cocktailLabeledImageLink, ''), COALESCE(cocktail.cocktailOrigin, ''), COALESCE(cocktail.cocktailSpokenName, ''), COALESCE(cocktail.cocktailDisplayName, '')," +
		" COALESCE(cocktail.cocktailKeywords, ''), COALESCE(cocktail.cocktailTop100Index, 0)" +
		" FROM cocktail WHERE ")

	parts := NGrams(str, 4)
	parts = RemoveBelowNLettersWords(parts, 3)
	parts = append(parts, RemoveBelowNLettersWords(SplitOnNonLetters(str), 3)...)
	fmt.Printf("%+v\n", parts)
	parts = RemoveDuplicates(parts)
	cocktailNameSQL := ""
	cocktailDescSQL := ""

	for i := 0; i < len(parts); i++ {
		cocktailNameSQL = cocktailNameSQL + " cocktailName LIKE ? OR "
		cocktailDescSQL = cocktailDescSQL + " cocktailDescription LIKE ? OR "
	}

	var total_parts []string
	for i := 0; i < 2; i++ {
		total_parts = append(total_parts, parts...)
	}

	if str != "" {
		buffer.WriteString("(" + cocktailNameSQL)
		buffer.WriteString(strings.TrimSuffix(cocktailDescSQL, "OR "))
		buffer.WriteString(") ")
	}

	ng := Sqlize(total_parts)
	if len(ng) == 0 {
		return c
	}

	for _, v := range ng {
		args = append(args, v)
	}

	buffer.WriteString(" AND (cocktail.cocktailRating>=? AND cocktail.cocktailRating<=?)")
	args = append(args, strconv.Itoa(search.RatingMin))
	args = append(args, strconv.Itoa(search.RatingMax))
	buffer.WriteString(";")
	query := buffer.String()
	log.Infoln(query)
	log.Infoln(args)
	rows, err := conn.Query(query, args...)
	if err != nil {
		log.Error(err)
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
			log.Error(err)
		}
		//add cocktail to cocktail family
		c = append(c, cocktail)
		log.Infoln(cocktail.ID, cocktail.Title, cocktail.Name, cocktail.Rating, cocktail.ImagePath,
			cocktail.Image, cocktail.Description, cocktail.Comment, cocktail.Footnotes, cocktail.SourceName, cocktail.SourceLink, cocktail.ImageSourceName, cocktail.ImageSourceLink,
			cocktail.LabeledImageLink, cocktail.SpokenName, cocktail.DisplayName, cocktail.Keywords, cocktail.Top100Index)
	}
	err = rows.Err()
	if err != nil {
		log.Error(err)
	}
	return c
}

func getIngredientIncludeSQL(search *Search, site string) []Cocktail {
	var buffer bytes.Buffer
	var args []interface{}
	var c []Cocktail
	conn, _ := connectors.GetDBFromMap(site)
	buffer.WriteString("SELECT DISTINCT cocktail.idCocktail, cocktail.cocktailTitle, cocktail.cocktailName," +
		" cocktail.cocktailRating, cocktail.cocktailImagePath, cocktail.cocktailImage, cocktail.cocktailDescription, cocktail.cocktailComment," +
		" cocktail.cocktailFootnotes, cocktail.cocktailSourceName, cocktail.cocktailSourceLink, cocktail.cocktailImageSourceName," +
		" cocktail.cocktailImageSourceLink, cocktail.cocktailLabeledImageLink, cocktail.cocktailOrigin, cocktail.cocktailSpokenName," +
		" cocktail.cocktailDisplayName, cocktail.cocktailKeywords, cocktailTop100Index FROM commonwealthcocktails.cocktail" +
		" JOIN commonwealthcocktails.cocktailToRecipe ON cocktailToRecipe.idCocktail=cocktail.idCocktail JOIN commonwealthcocktails.recipeToRecipeSteps" +
		" ON recipeToRecipeSteps.idRecipe=cocktailToRecipe.idRecipe JOIN commonwealthcocktails.recipestep ON" +
		" recipeToRecipeSteps.idRecipeStep=recipestep.idRecipeStep WHERE ")
	buffer.WriteString(" recipestep.recipestepOriginalIngredient IN (?" + strings.Repeat(",?", len(search.Include_Ingredients)-1) + ")")
	for i := range search.Include_Ingredients {
		args = append(args, search.Include_Ingredients[i])
	}
	buffer.WriteString(" AND (cocktail.cocktailRating>=? AND cocktail.cocktailRating<=?)")
	args = append(args, strconv.Itoa(search.RatingMin))
	args = append(args, strconv.Itoa(search.RatingMax))
	buffer.WriteString(";")
	query := buffer.String()
	log.Infoln(query)
	log.Infoln(args)
	rows, err := conn.Query(query, args...)
	if err != nil {
		log.Error(err)
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
			log.Error(err)
		}
		//add cocktail to cocktail family
		c = append(c, cocktail)
		log.Infoln(cocktail.ID, cocktail.Title, cocktail.Name, cocktail.Rating, cocktail.ImagePath,
			cocktail.Image, cocktail.Description, cocktail.Comment, cocktail.Footnotes, cocktail.SourceName, cocktail.SourceLink, cocktail.ImageSourceName, cocktail.ImageSourceLink,
			cocktail.LabeledImageLink, cocktail.SpokenName, cocktail.DisplayName, cocktail.Keywords, cocktail.Top100Index)
	}
	err = rows.Err()
	if err != nil {
		log.Error(err)
	}
	return c
}

func getMetaIncludeSQL(search *Search, site string) []Cocktail {
	var buffer bytes.Buffer
	var args []interface{}
	var c []Cocktail
	conn, _ := connectors.GetDBFromMap(site)
	buffer.WriteString("SELECT DISTINCT cocktail.idCocktail, cocktail.cocktailTitle, cocktail.cocktailName," +
		" cocktail.cocktailRating, cocktail.cocktailImagePath, cocktail.cocktailImage, cocktail.cocktailDescription, cocktail.cocktailComment," +
		" cocktail.cocktailFootnotes, cocktail.cocktailSourceName, cocktail.cocktailSourceLink, cocktail.cocktailImageSourceName," +
		" cocktail.cocktailImageSourceLink, cocktail.cocktailLabeledImageLink, cocktail.cocktailOrigin, cocktail.cocktailSpokenName," +
		" cocktail.cocktailDisplayName, cocktail.cocktailKeywords, cocktailTop100Index FROM commonwealthcocktails.cocktail" +
		" JOIN commonwealthcocktails.cocktailToMetas ON cocktailToMetas.idCocktail=cocktail.idCocktail WHERE ")
	buffer.WriteString(" cocktailToMetas.idMeta IN (?" + strings.Repeat(",?", len(search.Include_Metas)-1) + ")")
	for i := range search.Include_Metas {
		args = append(args, search.Include_Metas[i])
	}
	buffer.WriteString(" AND (cocktail.cocktailRating>=? AND cocktail.cocktailRating<=?)")
	args = append(args, strconv.Itoa(search.RatingMin))
	args = append(args, strconv.Itoa(search.RatingMax))
	buffer.WriteString(";")
	query := buffer.String()
	log.Infoln(query)
	log.Infoln(args)
	rows, err := conn.Query(query, args...)
	if err != nil {
		log.Error(err)
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
			log.Error(err)
		}
		//add cocktail to cocktail family
		c = append(c, cocktail)
		log.Infoln(cocktail.ID, cocktail.Title, cocktail.Name, cocktail.Rating, cocktail.ImagePath,
			cocktail.Image, cocktail.Description, cocktail.Comment, cocktail.Footnotes, cocktail.SourceName, cocktail.SourceLink, cocktail.ImageSourceName, cocktail.ImageSourceLink,
			cocktail.LabeledImageLink, cocktail.SpokenName, cocktail.DisplayName, cocktail.Keywords, cocktail.Top100Index)
	}
	err = rows.Err()
	if err != nil {
		log.Error(err)
	}
	return c
}

func getIngredientExcludeSQL(search *Search, site string) []Cocktail {
	var buffer bytes.Buffer
	var args []interface{}
	var c []Cocktail
	conn, _ := connectors.GetDBFromMap(site)
	buffer.WriteString("SELECT DISTINCT cocktail.idCocktail, COALESCE(cocktail.cocktailTitle, ''), COALESCE(cocktail.cocktailName, ''), cocktail.cocktailRating," +
		" COALESCE(cocktail.cocktailImagePath, ''), COALESCE(cocktail.cocktailImage, ''), COALESCE(cocktail.cocktailDescription, ''), COALESCE(cocktail.cocktailComment, ''), COALESCE(cocktail.cocktailFootnotes, '')," +
		" COALESCE(cocktail.cocktailSourceName, ''), COALESCE(cocktail.cocktailSourceLink, ''), COALESCE(cocktail.cocktailImageSourceName, ''), COALESCE(cocktail.cocktailImageSourceLink, '')," +
		" COALESCE(cocktail.cocktailLabeledImageLink, ''), COALESCE(cocktail.cocktailOrigin, ''), COALESCE(cocktail.cocktailSpokenName, ''), COALESCE(cocktail.cocktailDisplayName, '')," +
		" COALESCE(cocktail.cocktailKeywords, ''), COALESCE(cocktailTop100Index, 0) FROM commonwealthcocktails.cocktail JOIN commonwealthcocktails.cocktailToRecipe" +
		" ON cocktailToRecipe.idCocktail=cocktail.idCocktail JOIN commonwealthcocktails.recipeToRecipeSteps ON" +
		" recipeToRecipeSteps.idRecipe=cocktailToRecipe.idRecipe WHERE recipeToRecipeSteps.idRecipe NOT IN (SELECT recipeToRecipeSteps.idRecipe" +
		" FROM commonwealthcocktails.recipeToRecipeSteps JOIN commonwealthcocktails.recipestep ON recipeToRecipeSteps.idRecipeStep=recipestep.idRecipeStep" +
		" WHERE ")
	buffer.WriteString(" recipestep.recipestepOriginalIngredient IN (?" + strings.Repeat(",?", len(search.Exclude_Ingredients)-1) + "))")
	for j := range search.Exclude_Ingredients {
		args = append(args, search.Exclude_Ingredients[j])
	}
	buffer.WriteString(" AND (cocktail.cocktailRating>=? AND cocktail.cocktailRating<=?)")
	args = append(args, strconv.Itoa(search.RatingMin))
	args = append(args, strconv.Itoa(search.RatingMax))
	buffer.WriteString(";")
	query := buffer.String()
	log.Infoln(query)
	log.Infoln(args)
	rows, err := conn.Query(query, args...)
	if err != nil {
		log.Error(err)
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
			log.Error(err)
		}
		//add cocktail to cocktail family
		c = append(c, cocktail)
		log.Infoln(cocktail.ID, cocktail.Title, cocktail.Name, cocktail.Rating, cocktail.ImagePath,
			cocktail.Image, cocktail.Description, cocktail.Comment, cocktail.Footnotes, cocktail.SourceName, cocktail.SourceLink, cocktail.ImageSourceName, cocktail.ImageSourceLink,
			cocktail.LabeledImageLink, cocktail.SpokenName, cocktail.DisplayName, cocktail.Keywords, cocktail.Top100Index)
	}
	err = rows.Err()
	if err != nil {
		log.Error(err)
	}
	return c
}

func getNonIngredientIncludeSQL(search *Search, site string) []Cocktail {
	var buffer bytes.Buffer
	var args []interface{}
	var c []Cocktail
	conn, _ := connectors.GetDBFromMap(site)
	buffer.WriteString("SELECT DISTINCT cocktail.idCocktail, cocktail.cocktailTitle, cocktail.cocktailName," +
		" cocktail.cocktailRating, cocktail.cocktailImagePath, cocktail.cocktailImage, cocktail.cocktailDescription, cocktail.cocktailComment," +
		" cocktail.cocktailFootnotes, cocktail.cocktailSourceName, cocktail.cocktailSourceLink, cocktail.cocktailImageSourceName," +
		" cocktail.cocktailImageSourceLink, cocktail.cocktailLabeledImageLink, cocktail.cocktailOrigin, cocktail.cocktailSpokenName," +
		" cocktail.cocktailDisplayName, cocktail.cocktailKeywords, cocktailTop100Index FROM commonwealthcocktails.cocktail" +
		" JOIN commonwealthcocktails.cocktailToProducts ON cocktailToProducts.idCocktail=cocktail.idCocktail WHERE ")
	buffer.WriteString(" cocktailToProducts.idProduct IN (?" + strings.Repeat(",?", len(search.Include_NonIngredients)-1) + ")")
	for i := range search.Include_NonIngredients {
		args = append(args, search.Include_NonIngredients[i])
	}
	buffer.WriteString(" AND (cocktail.cocktailRating>=? AND cocktail.cocktailRating<=?)")
	args = append(args, strconv.Itoa(search.RatingMin))
	args = append(args, strconv.Itoa(search.RatingMax))
	buffer.WriteString(";")
	query := buffer.String()
	log.Infoln(query)
	log.Infoln(args)
	rows, err := conn.Query(query, args...)
	if err != nil {
		log.Error(err)
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
			log.Error(err)
		}
		//add cocktail to cocktail family
		c = append(c, cocktail)
		log.Infoln(cocktail.ID, cocktail.Title, cocktail.Name, cocktail.Rating, cocktail.ImagePath,
			cocktail.Image, cocktail.Description, cocktail.Comment, cocktail.Footnotes, cocktail.SourceName, cocktail.SourceLink, cocktail.ImageSourceName, cocktail.ImageSourceLink,
			cocktail.LabeledImageLink, cocktail.SpokenName, cocktail.DisplayName, cocktail.Keywords, cocktail.Top100Index)
	}
	err = rows.Err()
	if err != nil {
		log.Error(err)
	}
	return c
}

func getMetaExcludeSQL(search *Search, site string) []Cocktail {
	var buffer bytes.Buffer
	var args []interface{}
	var c []Cocktail
	conn, _ := connectors.GetDBFromMap(site)
	buffer.WriteString("SELECT DISTINCT cocktail.idCocktail, COALESCE(cocktail.cocktailTitle, ''), COALESCE(cocktail.cocktailName, ''), cocktail.cocktailRating," +
		" COALESCE(cocktail.cocktailImagePath, ''), COALESCE(cocktail.cocktailImage, ''), COALESCE(cocktail.cocktailDescription, ''), COALESCE(cocktail.cocktailComment, ''), COALESCE(cocktail.cocktailFootnotes, '')," +
		" COALESCE(cocktail.cocktailSourceName, ''), COALESCE(cocktail.cocktailSourceLink, ''), COALESCE(cocktail.cocktailImageSourceName, ''), COALESCE(cocktail.cocktailImageSourceLink, '')," +
		" COALESCE(cocktail.cocktailLabeledImageLink, ''), COALESCE(cocktail.cocktailOrigin, ''), COALESCE(cocktail.cocktailSpokenName, ''), COALESCE(cocktail.cocktailDisplayName, '')," +
		" COALESCE(cocktail.cocktailKeywords, ''), COALESCE(cocktailTop100Index, 0) FROM cocktail WHERE cocktail.idCocktail NOT IN (SELECT DISTINCT idCocktail FROM cocktailToMetas WHERE ")
	buffer.WriteString(" cocktailToMetas.idMeta IN (?" + strings.Repeat(",?", len(search.Exclude_Metas)-1) + "))")
	for j := range search.Exclude_Metas {
		args = append(args, search.Exclude_Metas[j])
	}
	buffer.WriteString(" AND (cocktail.cocktailRating>=? AND cocktail.cocktailRating<=?)")
	args = append(args, strconv.Itoa(search.RatingMin))
	args = append(args, strconv.Itoa(search.RatingMax))
	buffer.WriteString(";")
	query := buffer.String()
	log.Infoln(query)
	log.Infoln(args)
	rows, err := conn.Query(query, args...)
	if err != nil {
		log.Error(err)
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
			log.Error(err)
		}
		//add cocktail to cocktail family
		c = append(c, cocktail)
		log.Infoln(cocktail.ID, cocktail.Title, cocktail.Name, cocktail.Rating, cocktail.ImagePath,
			cocktail.Image, cocktail.Description, cocktail.Comment, cocktail.Footnotes, cocktail.SourceName, cocktail.SourceLink, cocktail.ImageSourceName, cocktail.ImageSourceLink,
			cocktail.LabeledImageLink, cocktail.SpokenName, cocktail.DisplayName, cocktail.Keywords, cocktail.Top100Index)
	}
	err = rows.Err()
	if err != nil {
		log.Error(err)
	}
	return c
}

func getNonIngredientExcludeSQL(search *Search, site string) []Cocktail {
	var buffer bytes.Buffer
	var args []interface{}
	var c []Cocktail
	conn, _ := connectors.GetDBFromMap(site)
	buffer.WriteString("SELECT DISTINCT cocktail.idCocktail, COALESCE(cocktail.cocktailTitle, ''), COALESCE(cocktail.cocktailName, ''), cocktail.cocktailRating," +
		" COALESCE(cocktail.cocktailImagePath, ''), COALESCE(cocktail.cocktailImage, ''), COALESCE(cocktail.cocktailDescription, ''), COALESCE(cocktail.cocktailComment, ''), COALESCE(cocktail.cocktailFootnotes, '')," +
		" COALESCE(cocktail.cocktailSourceName, ''), COALESCE(cocktail.cocktailSourceLink, ''), COALESCE(cocktail.cocktailImageSourceName, ''), COALESCE(cocktail.cocktailImageSourceLink, '')," +
		" COALESCE(cocktail.cocktailLabeledImageLink, ''), COALESCE(cocktail.cocktailOrigin, ''), COALESCE(cocktail.cocktailSpokenName, ''), COALESCE(cocktail.cocktailDisplayName, '')," +
		" COALESCE(cocktail.cocktailKeywords, ''), COALESCE(cocktailTop100Index, 0) FROM cocktail WHERE cocktail.idCocktail NOT IN (SELECT DISTINCT idCocktail FROM cocktailToProducts WHERE ")
	buffer.WriteString(" cocktailToProducts.idProduct IN (?" + strings.Repeat(",?", len(search.Exclude_NonIngredients)-1) + "))")
	for j := range search.Exclude_NonIngredients {
		args = append(args, search.Exclude_NonIngredients[j])
	}
	buffer.WriteString(" AND (cocktail.cocktailRating>=? AND cocktail.cocktailRating<=?)")
	args = append(args, strconv.Itoa(search.RatingMin))
	args = append(args, strconv.Itoa(search.RatingMax))
	buffer.WriteString(";")
	query := buffer.String()
	log.Infoln(query)
	log.Infoln(args)
	rows, err := conn.Query(query, args...)
	if err != nil {
		log.Error(err)
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
			log.Error(err)
		}
		//add cocktail to cocktail family
		c = append(c, cocktail)
		log.Infoln(cocktail.ID, cocktail.Title, cocktail.Name, cocktail.Rating, cocktail.ImagePath,
			cocktail.Image, cocktail.Description, cocktail.Comment, cocktail.Footnotes, cocktail.SourceName, cocktail.SourceLink, cocktail.ImageSourceName, cocktail.ImageSourceLink,
			cocktail.LabeledImageLink, cocktail.SpokenName, cocktail.DisplayName, cocktail.Keywords, cocktail.Top100Index)
	}
	err = rows.Err()
	if err != nil {
		log.Error(err)
	}
	return c
}

func (search *Search) QuickSearch(site string) ([]Cocktail, []Post, []Product) {
	var cocktails []Cocktail
	var posts []Post
	var products []Product

	conn, _ := connectors.GetDBFromMap(site)
	str := strings.ToLower(search.Keywords)
	parts := NGrams(str, 4)
	parts = RemoveBelowNLettersWords(parts, 3)
	parts = append(parts, RemoveBelowNLettersWords(SplitOnNonLetters(str), 3)...)
	fmt.Printf("%+v\n", parts)
	parts = RemoveDuplicates(parts)
	cocktailNameSQL := ""
	cocktailDescSQL := ""
	productNameSQL := ""
	productDescSQL := ""
	postTitleSQL := ""
	postContentSQL := ""
	for i := 0; i < len(parts); i++ {
		cocktailNameSQL = cocktailNameSQL + " cocktailName LIKE ? OR "
		cocktailDescSQL = cocktailDescSQL + " cocktailDescription LIKE ? OR "

		productNameSQL = productNameSQL + " productName LIKE ? OR "
		productDescSQL = productDescSQL + " productDescription LIKE ? OR "

		postTitleSQL = postTitleSQL + " postTitle LIKE ? OR "
		postContentSQL = postContentSQL + " postContent LIKE ? OR "
	}

	//we are doing this the number of times we have SQL statements from above
	var total_parts []string
	for i := 0; i < 6; i++ {
		total_parts = append(total_parts, parts...)
	}

	var buffer bytes.Buffer
	buffer.WriteString("SELECT idCocktail, cocktailName, '' as idPost, '' as postTitle, '' as idProduct, '' as productName FROM cocktail WHERE ")
	buffer.WriteString(cocktailNameSQL)
	buffer.WriteString(strings.TrimSuffix(cocktailDescSQL, "OR "))
	buffer.WriteString(" UNION ")
	buffer.WriteString("SELECT '' as idCocktail, '' as cocktailName, idPost, postTitle, '' as idProduct, '' as productName FROM post WHERE ")
	buffer.WriteString(postTitleSQL)
	buffer.WriteString(strings.TrimSuffix(postContentSQL, "OR "))
	buffer.WriteString(" UNION ")
	buffer.WriteString("SELECT '' as idCocktail, '' as cocktailName, '' as idPost, '' as postTitle, idProduct, productName FROM product WHERE ")
	buffer.WriteString(productNameSQL)
	buffer.WriteString(strings.TrimSuffix(productDescSQL, "OR "))
	buffer.WriteString(";")
	query := buffer.String()
	log.Infoln(query)
	//   SELECT idCocktail, cocktailName, NULL as idPost, NULL as postTitle, NULL as idProduct, NULL as productName FROM cocktail WHERE cocktailName LIKE "%gla%" OR cocktailName LIKE "%mar%" OR cocktailDescription LIKE "%gla%" OR cocktailDescription LIKE "%mar%"
	//   UNION
	//   SELECT NULL as idCocktail, NULL as cocktailName, idPost, postTitle, NULL as idProduct, NULL as productName FROM post WHERE postTitle LIKE "%mar%" OR postTitle LIKE "%gla%" OR postContent LIKE "%mar%" OR postContent LIKE "%gla%"
	//   UNION
	//   SELECT NULL as idCocktail, NULL as cocktailName, NULL as idPost, NULL as postTitle, idProduct, productName FROM product WHERE productName LIKE "%mar%" OR productName LIKE "%gla%" OR productDescription LIKE "%mar%" OR productDescription LIKE "%gla%"
	//   ;

	fmt.Println(total_parts)
	ng := Sqlize(total_parts)
	fmt.Println(ng)
	if len(ng) == 0 {
		return cocktails, posts, products
	}

	args := make([]interface{}, len(ng))
	for i, v := range ng {
		args[i] = v
	}
	rows, err := conn.Query(query, args...)
	if err != nil {
		log.Error(err)
	}
	defer rows.Close()
	for rows.Next() {
		cocktailID := ""
		cocktailName := ""
		productID := ""
		productName := ""
		postID := ""
		postTitle := ""
		err := rows.Scan(&cocktailID, &cocktailName, &postID, &postTitle, &productID, &productName)
		if err != nil {
			log.Error(err)
		}
		if cocktailID != "" {
			log.Infoln(cocktailID, cocktailName)
			id, _ := strconv.Atoi(cocktailID)
			cocktail := Cocktail{
				ID:   id,
				Name: template.HTML(html.UnescapeString(cocktailName)),
			}
			cocktails = append(cocktails, cocktail)
		} else if postID != "" {
			log.Infoln(postID, postTitle)
			id, _ := strconv.Atoi(postID)
			post := Post{
				ID:        id,
				PostTitle: template.HTML(html.UnescapeString(postTitle)),
			}
			posts = append(posts, post)
		} else if productID != "" {
			log.Infoln(productID, productName)
			id, _ := strconv.Atoi(productID)
			product := Product{
				ID:          id,
				ProductName: template.HTML(html.UnescapeString(productName)),
			}
			products = append(products, product)
		}
	}
	return cocktails, posts, products
}

// SplitOnNonLetters splits a string on non-letter runes
func SplitOnNonLetters(s string) []string {
	notALetter := func(char rune) bool { return !unicode.IsLetter(char) }
	return strings.FieldsFunc(s, notALetter)
}

func NGrams(s string, n int) []string {
	sub := ""
	subs := []string{}

	runes := bytes.Runes([]byte(s))
	l := len(runes)
	for i, r := range runes {
		sub = sub + string(r)
		if (i+1)%n == 0 {
			subs = append(subs, sub)
			sub = ""
		} else if (i + 1) == l {
			subs = append(subs, sub)
		}
	}

	return subs
}

func Sqlize(words []string) []string {
	var ret []string
	for i, _ := range words {
		ret = append(ret, "%"+words[i]+"%")
	}
	return ret
}

func RemoveBelowNLettersWords(elements []string, n int) []string {
	result := []string{}

	for v := range elements {
		if len(elements[v]) >= n {
			// Append to result slice.
			result = append(result, elements[v])
		} else if _, err := strconv.Atoi(elements[v]); err == nil {
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}

func RemoveDuplicates(elements []string) []string {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	result := []string{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}
