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
	PreText     string
	PostText    string
	Image       string
}

type DerivedProduct struct {
	Product     Product
	BaseProduct Product
}

type BaseProductWithDerived struct {
	Product         Product
	DerivedProducts []Product
	BaseProduct     Product
}

func GetBaseProductWithDerived(ID int) *BaseProductWithDerived {
	p := Products[ID-1]
	var bpwd BaseProductWithDerived
	bpwd.Product = p
	var dp []Product
	var bp Product
	if p.IsBase {
		for _, dps_element := range DerivedProducts {
			if dps_element.BaseProduct.ID == ID {
				dp = append(dp, dps_element.Product)
			}
		}
		bpwd.DerivedProducts = dp
	} else {
		for _, dps_element := range DerivedProducts {
			if dps_element.Product.ID == ID {
				bp = dps_element.BaseProduct
				break
			}
		}
		bpwd.BaseProduct = bp
	}
	return &bpwd
}
