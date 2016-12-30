//model/cocktail.go
package model

import (
	"bytes"
	"encoding/gob"
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
	Description     string
	Comment         string
	Recipe          Recipe
	Garnish         []Product
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

	//Advertiser Info
	Advertisement Advertisement
}

type Name struct {
	Name string
}

func InitCocktailTable() string {
	log.Println("CocktailTable Init")
	return "CocktailTable Init"
}

func GetCocktail() Cocktail {
	var c Cocktail
	var network bytes.Buffer        // Stand-in for a network connection
	enc := gob.NewEncoder(&network) // Will write to network.
	dec := gob.NewDecoder(&network) // Will read from network.
	// Encode (send) the value.
	err := enc.Encode(Cocktails[rand.Intn(len(Cocktails))])
	if err != nil {
		log.Fatal("encode error:", err)
	}
	// Decode (receive) the value.
	err = dec.Decode(&c)
	if err != nil {
		log.Fatal("decode error:", err)
	}
	//c := &Cocktails[rand.Intn(len(Cocktails))]
	prod_ignore := []int{}

	for ad_index, ad_element := range Advertisements {
		for _, adcocktails_element := range ad_element.Cocktails {
			if c.ID == adcocktails_element.ID {
				c.Advertisement = Advertisements[ad_index]
				for index, element := range c.Recipe.RecipeSteps {
					// element is the element from someSlice for where we are
					// is this a base product
					for _, adprod_element := range ad_element.Products {
						if element.OriginalIngredient.ID == adprod_element.BaseProduct.ID {
							c.Recipe.RecipeSteps[index].AdIngredient = adprod_element.AdvertisedProduct.Product
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
						c.Recipe.RecipeSteps[index].AdIngredient = adprod_element.AdvertisedProduct.Product
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
					c.Drinkware[index] = adprod_element.AdvertisedProduct.Product
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
					c.Tool[index] = adprod_element.AdvertisedProduct.Product
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
					c.Garnish[index] = adprod_element.AdvertisedProduct.Product
				}
			}
		}
	}
	return c
}

func intInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
