package www

import (
	"model"
)

type Page struct {
	Username          string
	CocktailSet       model.CocktailSet
	MetasByTypes      model.MetasByTypes
	Cocktail          model.Cocktail
	BaseProductWithBD model.BaseProductWithBD
	Products          []model.Product
}
