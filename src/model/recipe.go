//model/recipe.go
package model

type Recipe struct {
	ID          int
	RecipeSteps []RecipeStep
	Method      string
}

type RecipeStep struct {
	Ingredient     Product
	RecipeCardinal float64
	RecipeDoze     Doze
	RecipeOrdinal  int
}

type Doze int

const (
	Shot = 1 + iota
	Ounce
	Whole
	Dash
	Slice
)

var Dozes = [...]string{
	"shot",
	"oz.",
	"whole",
	"dash",
	"slice",
}

// String returns the English name of the doze ("Shot", "Ounce", ...).
func (d Doze) String() string { return Dozes[d-1] }
