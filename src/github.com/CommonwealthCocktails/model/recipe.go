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
//A doze represents a unit of measure.
type Doze struct {
	ID       int
	DozeName string
}

//Doze constants represent the unites of measure you can have in a recipe step
const (
	Shot = 1 + iota
	Ounce
	Whole
	Dash
	Slice
	TopOffWith
	Fresh
	Splash
	Leaf
	Drop
	Segment
	Grind
	Pinch
	Wedge
	Teaspoon
)

//String values for the int consts
var DozeStrings = [...]string{
	"shot(s)",
	"ounce(s)",
	"whole",
	"dash(es)",
	"slice(s)",
	"top off with",
	"fresh",
	"splash(es)",
	"leaf(ves)",
	"drop(s)",
	"segment(s)",
	"grind(s)",
	"pinch(es)",
	"wedge(s)",
	"teaspoon(s)",
}

// String returns the English name of the doze ("Shot", "Ounce", ...).
func (d Doze) String() string { return DozeStrings[d.ID-1] }
