//model/cocktail.go
package model

// recipe:
//   - !recipeStep
//       ingredient:
//          ProductName: Pineapple
//          ProductType: 3
//       recipeCardinal: 1.0
//       recipeDoze: whole
//       recipeOrdinal: 1

type Cocktail struct {
	Title           string
	Name            string
	DisplayName     string
	AlternateName   []Name
	SpokenName      string
	Origin          string
	AKA             []Name
	Description     string
	Comment         string
	Recipe          Recipe
	Method          string
	Garnish         []Product
	Image           string
	ImageSourceName string
	ImageSourceLink string
	Drinkware       []Product
	Tool            []Product
	SourceName      string
	SourceLink      string
	Rating          float32
	Flavor          []Meta
	Type            []Meta
	BaseSpirit      []Meta
	Served          []Meta
	Technique       []Meta
	Strength        []Meta
	Difficulty      []Meta
	TOD             []Meta
}

type Name struct {
	Name string
}

type Recipe struct {
	RecipeSteps []RecipeStep
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
)

var Dozes = [...]string{
	"shot",
	"oz.",
	"whole",
	"dash",
}

// String returns the English name of the doze ("Shot", "Ounce", ...).
func (d Doze) String() string { return Dozes[d-1] }

type ProductType int

const (
	Spirit = 1 + iota
	Liqueur
	Wine
	Mixer
	Beer
	Garnish
	Drinkware
	Tool
)

var ProductTypeStrings = [...]string{
	"Spirit",
	"Liqueur",
	"Wine",
	"Mixer",
	"Beer",
	"Garnish",
	"Drinkware",
	"Tool",
}

// String returns the English name of the Producttype ("Spirit", "Liqueur", ...).
func (ct ProductType) String() string { return ProductTypeStrings[ct-1] }

type Product struct {
	ProductName string
	ProductType ProductType
	Article     string
	Blurb       string
}

type MetaType int

const (
	Flavor = 1 + iota
	BaseSpirit
	Type
	Occasion
	Family
	Formula
	Served
	Technique
	Strength
	Difficulty
	TOD
)

var MetaTypeStrings = [...]string{
	"Flavor",      //Fruity, Bitter, Creamy, ...
	"Base Spirit", //Vodka, Gin, Bourbon, ...
	"Type",        //Tiki,
	"Occasion",    //Christmas, 4th of July, Halloween, ...
	"Family",      //Margarita, Martini, ...
	"Formula",
	"Served",      //Highball, Martini, Old Fashioned, ...
	"Technique",   //Shaking, Stirring, Straining
	"Strength",    //Weak, Medium, Strong
	"Difficulty",  //Easy, Medium, Hard
	"Time of Day", //Evening, Dessert, Brunch, ...
}

// String returns the English name of the metatype ("Flavor", "Base Spirit", ...).
func (mt MetaType) String() string { return MetaTypeStrings[mt-1] }

type Meta struct {
	MetaName string
	MetaType MetaType
	Article  string
	Blurb    string
}
