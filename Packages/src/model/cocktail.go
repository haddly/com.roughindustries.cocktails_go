//model/cocktail.go
package model

import (
	"bytes"
	"encoding/gob"
	"html/template"
	"log"
	"math/rand"
)

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
	Description     template.HTML
	Comment         template.HTML
	Recipe          Recipe
	Garnish         []Product
	ImagePath       string
	Image           string
	ImageSourceName string
	ImageSourceLink string
	Drinkware       []Product
	Tool            []Product
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
	Ratio           []Meta
	Family          []Meta
	IsFamilyRoot    bool
	About           Post
	Articles        []Post

	//Advertiser Info
	Advertisement Advertisement
}

type Name struct {
	Name string
}

type CocktailSet struct {
	ChildCocktails []Cocktail
	RootCocktail   Cocktail
	Cocktail       Cocktail
	Metadata       Meta
	Product        Product
}

type CocktailSearch struct {
	Products []Product
	Metadata []Meta
}

func GetCocktailSearch() CocktailSearch {
	var cs CocktailSearch
	for _, element := range Products {
		if element.ProductGroupType == Base {
			cs.Products = append(cs.Products, element)
		}
	}
	cs.Metadata = Metadata
	return cs
}

func copyCocktail(ID int) Cocktail {
	var c Cocktail
	var network bytes.Buffer        // Stand-in for a network connection
	enc := gob.NewEncoder(&network) // Will write to network.
	dec := gob.NewDecoder(&network) // Will read from network.
	var err error
	// Encode (send) the value.
	if ID <= 0 {
		err = enc.Encode(Cocktails[rand.Intn(len(Cocktails))])
	} else {
		err = enc.Encode(Cocktails[ID-1])
	}

	if err != nil {
		log.Fatal("encode error:", err)
	}
	// Decode (receive) the value.
	err = dec.Decode(&c)
	if err != nil {
		log.Fatal("decode error:", err)
	}

	return c
}

func GetCocktailByID(ID int) CocktailSet {
	var c Cocktail
	dst := GetDataSourceType()
	switch dst {
	case Internal:
		c = copyCocktail(ID)
		return processCocktailRequest(c)
	default:
		c.ID = ID
		return processCocktailRequest(c)

	}
}

func GetCocktail() CocktailSet {
	var c Cocktail
	c = copyCocktail(-1)
	return processCocktailRequest(c)
}

func GetCocktails() []Cocktail {
	var c []Cocktail
	c = Cocktails
	return c
}

func processCocktailRequest(c Cocktail) CocktailSet {
	dst := GetDataSourceType()
	switch dst {
	case Internal:
		return processInternalCocktailRequest(c)
	default:
		var cs CocktailSet
		cs = SelectCocktailsByID(c.ID)
		return cs

	}
}

func processInternalCocktailRequest(c Cocktail) CocktailSet {

	var cs CocktailSet

	prod_ignore := []int{}

	//This will have to be gotten through derived products I think TCH
	//for index, element := range c.Recipe.RecipeSteps {
	//	c.Recipe.RecipeSteps[index].AltIngredient = *GetSpecificProductsFromGroup(element.OriginalIngredient.ID)
	//}

	for ad_index, ad_element := range Advertisements {
		for _, adcocktails_element := range ad_element.Cocktails {
			if c.ID == adcocktails_element.ID {
				c.Advertisement = Advertisements[ad_index]
				for index, element := range c.Recipe.RecipeSteps {
					// element is the element from someSlice for where we are
					// is this a base product
					for _, adprod_element := range ad_element.Products {
						if element.OriginalIngredient.ID == adprod_element.BaseProduct.ID {
							c.Recipe.RecipeSteps[index].AdIngredient = adprod_element.AdvertisedProduct
							prod_ignore = append(prod_ignore, element.OriginalIngredient.ID)
						}
					}
				}
			}
		}
	}

	//recipe OriginalIngredient ad replacement
	for index, element := range c.Recipe.RecipeSteps {
		// element is the element from someSlice for where we are
		// is this a base product
		for _, ad_element := range Advertisements {
			for _, adprod_element := range ad_element.Products {
				if element.OriginalIngredient.ID == adprod_element.BaseProduct.ID {
					if !intInSlice(element.OriginalIngredient.ID, prod_ignore) {
						c.Recipe.RecipeSteps[index].AdIngredient = adprod_element.AdvertisedProduct
					}
				}
			}
		}
	}
	//drinkware ad replacement
	for index, element := range c.Drinkware {
		// element is the element from someSlice for where we are
		// is this a base product
		for _, ad_element := range Advertisements {
			for _, adprod_element := range ad_element.Products {
				if element.ID == adprod_element.BaseProduct.ID {
					c.Drinkware[index] = adprod_element.AdvertisedProduct
				}
			}
		}
	}
	//tool ad replacement
	for index, element := range c.Tool {
		// element is the element from someSlice for where we are
		// is this a base product
		for _, ad_element := range Advertisements {
			for _, adprod_element := range ad_element.Products {
				if element.ID == adprod_element.BaseProduct.ID {
					c.Tool[index] = adprod_element.AdvertisedProduct
				}
			}
		}
	}
	//garnish ad replacement
	for index, element := range c.Garnish {
		// element is the element from someSlice for where we are
		// is this a base product
		for _, ad_element := range Advertisements {
			for _, adprod_element := range ad_element.Products {
				if element.ID == adprod_element.BaseProduct.ID {
					c.Garnish[index] = adprod_element.AdvertisedProduct
				}
			}
		}
	}

	//put the the cocktails in a set struct
	if c.IsFamilyRoot {
		for _, element := range CS {
			if element.RootCocktail.ID == c.ID {
				cs.ChildCocktails = element.ChildCocktails
			}
		}
		cs.Cocktail = c
	} else {
		for _, cocktail := range CS {
			for _, element := range cocktail.ChildCocktails {
				if element.ID == c.ID {
					cs.RootCocktail = cocktail.RootCocktail
				}
			}
		}
		cs.Cocktail = c
	}
	return cs
}

func intInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
