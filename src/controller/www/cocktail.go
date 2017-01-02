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
func (cocktail *Cocktail) RenderTemplate(w http.ResponseWriter, tmpl string, fc *model.FamilyCocktail) {
	t, err := template.ParseFiles("./view/webcontent/www/templates/"+tmpl+".html", "./view/webcontent/www/templates/ga.html", "./view/webcontent/www/templates/navbar.html", "./view/webcontent/www/templates/footer.html")
	if err != nil {
		log.Fatal(err)
	}

	//c.Description = template.HTMLEscapeString(c.Description)
	//log.Println(c.Description)
	//c.Description = strings.Replace(c.Description, "\n", "<br>", -1)

	err = t.ExecuteTemplate(w, "base", fc)
	if err != nil {
		log.Fatal(err)
	}
}

func (cocktail *Cocktail) CocktailHandler(w http.ResponseWriter, r *http.Request) {
	var fc model.FamilyCocktail
	u, err := url.Parse(r.URL.String())
	if err != nil {
		cocktail.RenderTemplate(w, "404", &fc)
	}
	//log.Println("Product: " + r.URL.Path[1:])
	m, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		cocktail.RenderTemplate(w, "404", &fc)
	}
	if len(m["ID"]) == 0 {
		cocktail.RenderTemplate(w, "404", &fc)
	} else {
		//log.Println("ID: " + m["ID"][0])

		//apply the template page info to the index page
		id, _ := strconv.Atoi(m["ID"][0])
		if len(model.Products) <= id-1 {
			cocktail.RenderTemplate(w, "404", &fc)
		} else {
			fc = model.GetCocktailByID(id)

			//log.Println(c)
			//apply the template page info to the index page
			cocktail.RenderTemplate(w, "cocktail", &fc)
		}
	}
}

//handle / requests to the server
func (cocktail *Cocktail) CocktailsHandler(w http.ResponseWriter, r *http.Request) {
	//log.Println("indexHandler: " + r.URL.Path[1:])

	var fc model.FamilyCocktail
	fc = model.GetCocktail()

	//log.Println(c)
	//apply the template page info to the index page
	cocktail.RenderTemplate(w, "cocktail", &fc)
}

//handle / requests to the server
func (cocktail *Cocktail) CocktailSearchHandler(w http.ResponseWriter, r *http.Request) {
	//log.Println("indexHandler: " + r.URL.Path[1:])

	var fc model.FamilyCocktail

	//log.Println(c)
	//apply the template page info to the index page
	cocktail.RenderTemplate(w, "search", &fc)
}

//handle / requests to the server
func (cocktail *Cocktail) CocktailLandingHandler(w http.ResponseWriter, r *http.Request) {
	//log.Println("indexHandler: " + r.URL.Path[1:])

	var fc model.FamilyCocktail

	//log.Println(c)
	//apply the template page info to the index page
	cocktail.RenderTemplate(w, "index", &fc)
}

func (cocktail *Cocktail) Init() {
	//Web Service and Web Page Handlers
	log.Println("Init in www/Cocktail.go")
	http.HandleFunc("/", cocktail.CocktailLandingHandler)
	http.HandleFunc("/cocktail", cocktail.CocktailHandler)
	http.HandleFunc("/cocktails", cocktail.CocktailsHandler)
	http.HandleFunc("/search", cocktail.CocktailSearchHandler)

}
