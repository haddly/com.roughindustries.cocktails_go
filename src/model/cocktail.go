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
	ID              int
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
	Garnish         []*Product
	Image           string
	ImageSourceName string
	ImageSourceLink string
	Drinkware       []*Product
	Tool            []*Product
	SourceName      string
	SourceLink      string
	Rating          int
	Flavor          []Meta
	Type            []Meta
	BaseSpirit      []Meta
	Served          []Meta
	Technique       []Meta
	Strength        []Meta
	Difficulty      []Meta
	TOD             []Meta

	//Advertiser Info
	Advertisement Advertisement
}

type Name struct {
	Name string
}
