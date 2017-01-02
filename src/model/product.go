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

type BDGCategory int

const (
	Base = 1 + iota
	Derived
	Group
)

var BDGCategoryStrings = [...]string{
	"Base",
	"Derived",
	"Group",
}

// String returns the English name of the Producttype ("Base", "Derived", "Group).
func (bdg BDGCategory) String() string { return BDGCategoryStrings[bdg-1] }

type Product struct {
	ID          int
	ProductName string
	ProductType ProductType
	Article     string
	Blurb       string
	Recipe      Recipe
	BDG         BDGCategory
	PreText     string
	PostText    string
	Image       string
}

type DerivedProduct struct {
	Product     Product
	BaseProduct Product
}

type GroupProduct struct {
	Products     []Product
	GroupProduct Product
}

type BaseProductWithBDG struct {
	Product         Product
	DerivedProducts []Product
	ProductGroups   []Product
	BaseProduct     Product
}

func GetBaseProductWithBDG(ID int) *BaseProductWithBDG {
	p := Products[ID-1]
	var bpwbdg BaseProductWithBDG
	bpwbdg.Product = p
	var dp []Product
	var bp Product
	var gp []Product
	if p.BDG == Base {
		for _, dps_element := range DerivedProducts {
			if dps_element.BaseProduct.ID == ID {
				dp = append(dp, dps_element.Product)
			}
		}
		bpwbdg.DerivedProducts = dp
	} else if p.BDG == Derived {
		for _, dps_element := range DerivedProducts {
			if dps_element.Product.ID == ID {
				bp = dps_element.BaseProduct
				break
			}
		}
		bpwbdg.BaseProduct = bp
	} else {
		for _, gps_element := range ProductGroups {
			if gps_element.GroupProduct.ID == ID {
				for _, prod := range gps_element.Products {
					gp = append(gp, prod)
				}
			}
		}
		bpwbdg.ProductGroups = gp
	}
	return &bpwbdg
}

func GetSpecificProductsFromGroup(ID int) *BaseProductWithBDG {
	p := Products[ID-1]
	var bpwbdg BaseProductWithBDG
	bpwbdg.Product = p
	var gp []Product
	if p.BDG == Group {
		for _, gps_element := range ProductGroups {
			if gps_element.GroupProduct.ID == ID {
				for _, prod := range gps_element.Products {
					gp = append(gp, prod)
				}
			}
		}
		bpwbdg.ProductGroups = gp
	}
	return &bpwbdg
}
