//model/recipe.go
package model

type Recipe struct {
	ID          int
	RecipeSteps []RecipeStep
	Method      string
}

type RecipeStep struct {
	Ingredient           *Product
	RecipeCardinalFloat  float64
	RecipeCardinalString string
	RecipeDoze           Doze
	RecipeOrdinal        int
}

type Doze int

const (
	Shot = 1 + iota
	Ounce
	Whole
	Dash
	Slice
	TopOffWith
)

var Dozes = [...]string{
	"shot",
	"oz.",
	"whole",
	"dash",
	"slice",
	"top off with",
}

// String returns the English name of the doze ("Shot", "Ounce", ...).
func (d Doze) String() string { return Dozes[d-1] }
