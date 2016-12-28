//model/product.go
package model

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
	ID          int
	ProductName string
	ProductType ProductType
	Article     string
	Blurb       string
	Recipe      Recipe
	IsBase      bool
}

type DerivedProduct struct {
	ID          int
	ProductName string
	BaseProduct Product
	ProductType ProductType
	Article     string
	Blurb       string
	Recipe      Recipe
}
