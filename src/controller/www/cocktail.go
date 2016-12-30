package www

import (
	"html/template"
	"log"
	"model"
	"net/http"
	//"strconv"
)

type Cocktail struct {
}

//render the page based on the name of the file provided
func (cocktail *Cocktail) RenderTemplate(w http.ResponseWriter, tmpl string, c *model.Cocktail) {
	t, err := template.ParseFiles("./view/webcontent/www/templates/"+tmpl+".html", "./view/webcontent/www/templates/ga.html", "./view/webcontent/www/templates/header.html", "./view/webcontent/www/templates/footer.html")
	if err != nil {
		log.Fatal(err)
	}
	err = t.ExecuteTemplate(w, "base", c)
	if err != nil {
		log.Fatal(err)
	}
}

//handle / requests to the server
func (cocktail *Cocktail) CocktailHandler(w http.ResponseWriter, r *http.Request) {
	//log.Println("indexHandler: " + r.URL.Path[1:])

	var c model.Cocktail
	c = model.GetCocktail()

	//log.Println(c)
	//apply the template page info to the index page
	cocktail.RenderTemplate(w, "cocktail", &c)
}

//handle / requests to the server
func (cocktail *Cocktail) CocktailSearchHandler(w http.ResponseWriter, r *http.Request) {
	//log.Println("indexHandler: " + r.URL.Path[1:])

	var c model.Cocktail

	//log.Println(c)
	//apply the template page info to the index page
	cocktail.RenderTemplate(w, "search", &c)
}

func (cocktail *Cocktail) Init() {
	//Web Service and Web Page Handlers
	log.Println("Init in www/Cocktail.go")
	http.HandleFunc("/", cocktail.CocktailHandler)
	http.HandleFunc("/cocktail", cocktail.CocktailHandler)
	http.HandleFunc("/search", cocktail.CocktailSearchHandler)

}
