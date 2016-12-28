package www

import (
	"html/template"
	"log"
	"math/rand"
	"model"
	"net/http"
	"strconv"
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
	//parse the url and get the information after the localhost:8080/
	//stick that in the name
	//name := r.URL.Path[1:]
	//or setup a default for now
	c := &model.Cocktails[rand.Intn(len(model.Cocktails))]
	prod_ignore := []int{}

	for ad_index, ad_element := range model.Advertisements {
		for _, adcocktails_element := range ad_element.Cocktails {
			if c.ID == adcocktails_element.ID {
				c.Advertisement = model.Advertisements[ad_index]
				log.Println(strconv.Itoa(c.ID) + " " + strconv.Itoa(adcocktails_element.ID))
				for index, element := range c.Recipe.RecipeSteps {
					// element is the element from someSlice for where we are
					// is this a base product
					for _, adprod_element := range ad_element.Products {
						log.Println(strconv.Itoa(element.Ingredient.ID) + " " + strconv.Itoa(adprod_element.BaseProduct.ID))
						if element.Ingredient.ID == adprod_element.BaseProduct.ID {
							c.Recipe.RecipeSteps[index].Ingredient = &adprod_element.AdvertisedProduct.Product
							prod_ignore = append(prod_ignore, adprod_element.BaseProduct.ID)
						}
					}
				}
			}
		}
	}

	//recipe ingredient ad replacement
	for index, element := range c.Recipe.RecipeSteps {
		// element is the element from someSlice for where we are
		// is this a base product
		for _, ad_element := range model.Advertisements {
			for _, adprod_element := range ad_element.Products {
				if element.Ingredient.ID == adprod_element.BaseProduct.ID {
					c.Recipe.RecipeSteps[index].Ingredient = &adprod_element.AdvertisedProduct.Product
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
					c.Drinkware[index] = &adprod_element.AdvertisedProduct.Product
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
					c.Tool[index] = &adprod_element.AdvertisedProduct.Product
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
					c.Garnish[index] = &adprod_element.AdvertisedProduct.Product
				}
			}
		}
	}
	log.Println(c)
	//apply the template page info to the index page
	cocktail.RenderTemplate(w, "index", c)
}

func (cocktail *Cocktail) Init() {
	//Web Service and Web Page Handlers
	log.Println("Init in www/Cocktail.go")
	http.HandleFunc("/", cocktail.IndexHandler)
}
