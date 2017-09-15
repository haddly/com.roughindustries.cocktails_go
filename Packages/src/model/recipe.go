// Copyright 2017 Rough Industries LLC. All rights reserved.
//model/recipe.go:package model
package model

import (
	"html/template"
)

//DATA STRUCTURES
//Recipe data structure. A collection of steps with the method to bring those
//steps together
type Recipe struct {
	ID          int
	RecipeSteps []RecipeStep
	Method      template.HTML
}

//Recipe Step data structure. An individual component of a recipe.  Consists
//of an ingredient, order it fits into the recipe, is unit of measure, and
//quantity at the unit of measure.
type RecipeStep struct {
	ID                   int
	OriginalIngredient   Product
	AltIngredient        []Product
	AdIngredient         Product
	RecipeCardinalFloat  float64
	RecipeCardinalString string
	RecipeDoze           Doze
	RecipeOrdinal        int
}

//ENUMERATIONS - These must match the database one for one in both ID and order
//
type Doze struct {
	ID       int
	DozeName string
}

//
const (
	Shot = 1 + iota
	Ounce
	Whole
	Dash
	Slice
	TopOffWith
	Fresh
	Splash
)

//
var DozeStrings = [...]string{
	"shot",
	"oz.",
	"whole",
	"dash",
	"slice",
	"top off with",
	"fresh",
	"splash",
}

// String returns the English name of the doze ("Shot", "Ounce", ...).
func (d Doze) String() string { return DozeStrings[d.ID-1] }
