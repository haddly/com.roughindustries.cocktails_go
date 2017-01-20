//model/recipe.go
package model

import (
	"html/template"
)

type Recipe struct {
	ID          int
	RecipeSteps []RecipeStep
	Method      template.HTML
}

type RecipeStep struct {
	ID                   int
	OriginalIngredient   Product
	AltIngredient        []Product
	AdIngredient         Product
	RecipeCardinalFloat  float64
	RecipeCardinalString string
	RecipeDoze           Doze
	RecipeOrdinal        int
	BDG                  BaseProductWithBDG
}

type Doze int

const (
	Shot = 1 + iota
	Ounce
	Whole
	Dash
	Slice
	TopOffWith
	Fresh
)

var DozeStrings = [...]string{
	"shot",
	"oz.",
	"whole",
	"dash",
	"slice",
	"top off with",
	"fresh",
}

// String returns the English name of the doze ("Shot", "Ounce", ...).
func (d Doze) String() string { return DozeStrings[d-1] }
