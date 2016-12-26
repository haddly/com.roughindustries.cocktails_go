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
	t, err := template.ParseFiles("./view/webcontent/www/" + tmpl + ".html")
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(w, c)
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

	for _, ad_element := range model.Advertisements {
		for ad_index, adcocktails_element := range ad_element.Cocktails {
			if c.ID == adcocktails_element {
				c.Advertisement = model.Advertisements[ad_index]
				log.Println(strconv.Itoa(c.ID) + " " + strconv.Itoa(adcocktails_element))
				for index, element := range c.Recipe.RecipeSteps {
					// element is the element from someSlice for where we are
					// is this a base product
					for _, adprod_element := range ad_element.Products {
						log.Println(strconv.Itoa(element.Ingredient.ID) + " " + strconv.Itoa(adprod_element.BaseProductID))
						if element.Ingredient.ID == adprod_element.BaseProductID {
							c.Recipe.RecipeSteps[index].Ingredient = model.Products[adprod_element.AdvertisedProductID-1]
							prod_ignore = append(prod_ignore, adprod_element.BaseProductID)
						}
					}
				}
			}
		}
	}

	for index, element := range c.Recipe.RecipeSteps {
		// element is the element from someSlice for where we are
		// is this a base product
		for _, ad_element := range model.Advertisements {
			for _, adprod_element := range ad_element.Products {
				if element.Ingredient.ID == adprod_element.BaseProductID {
					c.Recipe.RecipeSteps[index].Ingredient = model.Products[adprod_element.AdvertisedProductID-1]
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
