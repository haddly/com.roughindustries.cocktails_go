//model/cocktail.go
package model

import (
	"html/template"
)

// recipe:
//   - !recipeStep
//       ingredient:
//          ProductName: Pineapple
//          ProductType: 3
//       recipeCardinal: 1.0
//       recipeDoze: whole
//       recipeOrdinal: 1

type Cocktail struct {
	ID              int
	Title           string
	Name            string
	DisplayName     string
	AlternateName   []Name
	SpokenName      string
	Origin          template.HTML
	AKA             []Name
	Description     template.HTML
	Comment         template.HTML
	ImagePath       string
	Image           string
	ImageSourceName string
	ImageSourceLink string
	SourceName      string
	SourceLink      string
	Rating          int
	Recipe          Recipe
	Garnish         []Product
	Drinkware       []Product
	Tool            []Product
	Flavor          []Meta
	Occasion        []Meta
	Formula         []Meta
	Type            []Meta
	BaseSpirit      []Meta
	Served          []Meta
	Technique       []Meta
	Strength        []Meta
	Difficulty      []Meta
	TOD             []Meta
	Ratio           []Meta
	Family          []Meta
	Drink           []Meta
	IsFamilyRoot    bool
	About           Post
	Articles        []Post
	Errors          map[string]string

	//Advertiser Info
	//Advertisement Advertisement
}

type Name struct {
	Name string
}

type CocktailSet struct {
	ChildCocktails []Cocktail
	RootCocktail   Cocktail
	Cocktail       Cocktail
	Metadata       Meta
	Product        Product
}

type CocktailSearch struct {
	Products []Product
	Metadata []Meta
}

type CocktailsByAlphaNums struct {
	CBA map[string][]Cocktail
}

func GetCocktailSearch() CocktailSearch {
	var cs CocktailSearch
	// for _, element := range Products {
	// 	if element.ProductGroupType == Base {
	// 		cs.Products = append(cs.Products, element)
	// 	}
	// }
	//cs.Metadata = Metadata
	return cs
}

func GetCocktailByID(ID int, includeBDG bool) CocktailSet {
	var c Cocktail
	c.ID = ID
	return processCocktailRequest(c, includeBDG)
}

func processCocktailRequest(c Cocktail, includeBDG bool) CocktailSet {
	var cs CocktailSet
	cs = SelectCocktailsByID(c.ID, includeBDG)
	return cs
}

func intInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
