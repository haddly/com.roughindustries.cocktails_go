package www

import (
	"bytes"
	"encoding/gob"
	"html/template"
	"log"
	"math/rand"
	"model"
	"net/http"
	//"strconv"
)

type Cocktail struct {
}

//render the page based on the name of the file provided
func (cocktail *Cocktail) RenderTemplate(w http.ResponseWriter, tmpl string, c *model.Cocktail) {
	t, err := template.ParseFiles("./view/webcontent/www/templates/"+tmpl+".html", "./view/webcontent/www/templates/ga.html")
	if err != nil {
		log.Fatal(err)
	}
	err = t.ExecuteTemplate(w, "base", c)
	if err != nil {
		log.Fatal(err)
	}
}

//handle / requests to the server
func (cocktail *Cocktail) IndexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("indexHandler: " + r.URL.Path[1:])

	var c model.Cocktail
	var network bytes.Buffer        // Stand-in for a network connection
	enc := gob.NewEncoder(&network) // Will write to network.
	dec := gob.NewDecoder(&network) // Will read from network.
	// Encode (send) the value.
	err := enc.Encode(model.Cocktails[rand.Intn(len(model.Cocktails))])
	if err != nil {
		log.Fatal("encode error:", err)
	}
	// Decode (receive) the value.
	err = dec.Decode(&c)
	if err != nil {
		log.Fatal("decode error:", err)
	}
	//c := &model.Cocktails[rand.Intn(len(model.Cocktails))]
	prod_ignore := []int{}

	for ad_index, ad_element := range model.Advertisements {
		for _, adcocktails_element := range ad_element.Cocktails {
			if c.ID == adcocktails_element.ID {
				c.Advertisement = model.Advertisements[ad_index]
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
		for _, ad_element := range model.Advertisements {
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
		for _, ad_element := range model.Advertisements {
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
		for _, ad_element := range model.Advertisements {
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
		for _, ad_element := range model.Advertisements {
			for _, adprod_element := range ad_element.Products {
				if element.ID == adprod_element.BaseProduct.ID {
					c.Garnish[index] = adprod_element.AdvertisedProduct.Product
				}
			}
		}
	}
	log.Println(c)
	//apply the template page info to the index page
	cocktail.RenderTemplate(w, "index", &c)
}

func intInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func (cocktail *Cocktail) Init() {
	//Web Service and Web Page Handlers
	log.Println("Init in www/Cocktail.go")
	http.HandleFunc("/", cocktail.IndexHandler)
}
