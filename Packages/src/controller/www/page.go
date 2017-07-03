package www

import (
	"model"
)

type Page struct {
	Username          string
	CocktailSet       model.CocktailSet
	MetasByTypes      model.MetasByTypes
	Ingredients       model.ProductsByTypes
	NonIngredients    model.ProductsByTypes
	Cocktail          model.Cocktail
	BaseProductWithBD model.BaseProductWithBD
	Products          []model.Product
	ProductsByTypes   model.ProductsByTypes
	Doze              []model.Doze
}
