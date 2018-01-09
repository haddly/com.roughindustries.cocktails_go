// Copyright 2017 Rough Industries LLC. All rights reserved.
//model/cocktail.go:package model
package model

import (
	"html/template"
)

//DATA STRUCTURES
//Cocktail data structure
type Cocktail struct {
	ID                  int
	Title               string
	Name                template.HTML
	DisplayName         string
	AlternateName       []Name
	SpokenName          string
	Origin              template.HTML
	AKA                 []Name
	Description         template.HTML
	Comment             template.HTML
	Footnotes           template.HTML
	Keywords            string
	ImagePath           string
	Image               string
	ImageSourceName     string
	ImageSourceLink     string
	LabeledImageLink    string
	SourceName          string
	SourceLink          string
	Rating              int
	Top100Index         int
	IgnoreRecipeUpdate  bool
	Recipe              Recipe
	IgnoreProductUpdate bool
	Garnish             []Product
	Drinkware           []Product
	Tool                []Product
	IgnoreMetaUpdate    bool
	Flavor              []Meta
	Occasion            []Meta
	Formula             []Meta
	Type                []Meta
	BaseSpirit          []Meta
	Served              []Meta
	Technique           []Meta
	Strength            []Meta
	Difficulty          []Meta
	TOD                 []Meta
	Style               []Meta
	Ratio               []Meta
	Family              []Meta
	Drink               []Meta
	IsFamilyRoot        bool
	Errors              map[string]string
}

//The name struct allows for easier access to an array of strings in
//the templating system
type Name struct {
	Name string
}

//This data struct represents a combination of cocktails with the cocktail
//of interest being the Cocktail object with either a root cocktail or a
//set of children cocktails along with the related products and metadata
type CocktailSet struct {
	ChildCocktails []Cocktail
	RootCocktail   Cocktail
	Cocktail       Cocktail
	Metadata       Meta
	Product        Product
}

//All the cocktails requested in a alpha numeric map
type CocktailsByAlphaNums struct {
	CBA map[string][]Cocktail
}
