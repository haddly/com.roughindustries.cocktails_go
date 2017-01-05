package www

import (
	"html/template"
	"log"
	"model"
	"net/http"
	"net/url"
	"strconv"
)

type Cocktail struct {
}

//render the page based on the name of the file provided
func (cocktail *Cocktail) RenderFamilyCocktailTemplate(w http.ResponseWriter, tmpl string, fc *model.FamilyCocktail) {
	t, err := parseTempFiles(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	err = t.ExecuteTemplate(w, "base", fc)
	if err != nil {
		log.Fatal(err)
	}
}

func (cocktail *Cocktail) RenderCocktailsTemplate(w http.ResponseWriter, tmpl string, c []model.Cocktail) {
	t, err := parseTempFiles(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	err = t.ExecuteTemplate(w, "base", c)
	if err != nil {
		log.Fatal(err)
	}
}

func parseTempFiles(tmpl string) (*template.Template, error) {
	return template.ParseFiles("./view/webcontent/www/templates/"+tmpl+".html", "./view/webcontent/www/templates/head.html", "./view/webcontent/www/templates/ga.html", "./view/webcontent/www/templates/navbar.html", "./view/webcontent/www/templates/footer.html")
}

func (cocktail *Cocktail) CocktailHandler(w http.ResponseWriter, r *http.Request) {
	var fc model.FamilyCocktail
	u, err := url.Parse(r.URL.String())
	if err != nil {
		cocktail.RenderFamilyCocktailTemplate(w, "404", &fc)
	}
	//log.Println("Product: " + r.URL.Path[1:])
	m, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		cocktail.RenderFamilyCocktailTemplate(w, "404", &fc)
	}
	if len(m["ID"]) == 0 {
		cocktail.RenderFamilyCocktailTemplate(w, "404", &fc)
	} else {
		//log.Println("ID: " + m["ID"][0])

		//apply the template page info to the index page
		id, _ := strconv.Atoi(m["ID"][0])
		if len(model.Products) <= id-1 {
			cocktail.RenderFamilyCocktailTemplate(w, "404", &fc)
		} else {
			fc = model.GetCocktailByID(id)

			//log.Println(c)
			//apply the template page info to the index page
			cocktail.RenderFamilyCocktailTemplate(w, "cocktail", &fc)
		}
	}
}

//handle / requests to the server
func (cocktail *Cocktail) CocktailsHandler(w http.ResponseWriter, r *http.Request) {
	//log.Println("indexHandler: " + r.URL.Path[1:])

	var c []model.Cocktail
	c = model.GetCocktails()

	//log.Println(c)
	//apply the template page info to the index page
	cocktail.RenderCocktailsTemplate(w, "cocktails", c)
}

//handle / requests to the server
func (cocktail *Cocktail) CocktailSearchHandler(w http.ResponseWriter, r *http.Request) {
	//log.Println("indexHandler: " + r.URL.Path[1:])

	var fc model.FamilyCocktail

	//log.Println(c)
	//apply the template page info to the index page
	cocktail.RenderFamilyCocktailTemplate(w, "search", &fc)
}

//handle / requests to the server
func (cocktail *Cocktail) CocktailLandingHandler(w http.ResponseWriter, r *http.Request) {
	//log.Println("indexHandler: " + r.URL.Path[1:])

	var fc model.FamilyCocktail

	//log.Println(c)
	//apply the template page info to the index page
	cocktail.RenderFamilyCocktailTemplate(w, "index", &fc)
}

func (cocktail *Cocktail) Init() {
	//Web Service and Web Page Handlers
	log.Println("Init in www/Cocktail.go")
	http.HandleFunc("/", cocktail.CocktailLandingHandler)
	http.HandleFunc("/cocktail", cocktail.CocktailHandler)
	http.HandleFunc("/cocktails", cocktail.CocktailsHandler)
}
