// Copyright 2017 Rough Industries LLC. All rights reserved.
//model/cocktail.go:package model
package model

import (
	"html/template"
)

//DATA STRUCTURES
//Cocktail data structure
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
	Errors          map[string]string
}

//
type Name struct {
	Name string
}

//
type CocktailSet struct {
	ChildCocktails []Cocktail
	RootCocktail   Cocktail
	Cocktail       Cocktail
	Metadata       Meta
	Product        Product
}

//
type CocktailSearch struct {
	Products []Product
	Metadata []Meta
}

//
type CocktailsByAlphaNums struct {
	CBA map[string][]Cocktail
}
