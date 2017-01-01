package www

import (
	"html/template"
	"log"
	"model"
	"net/http"
	// "strings"
	//"strconv"
)

type Cocktail struct {
}

//render the page based on the name of the file provided
func (cocktail *Cocktail) RenderTemplate(w http.ResponseWriter, tmpl string, c *model.Cocktail) {
	t, err := template.ParseFiles("./view/webcontent/www/templates/"+tmpl+".html", "./view/webcontent/www/templates/ga.html", "./view/webcontent/www/templates/navbar.html", "./view/webcontent/www/templates/footer.html")
	if err != nil {
		log.Fatal(err)
	}

	//c.Description = template.HTMLEscapeString(c.Description)
	//log.Println(c.Description)
	//c.Description = strings.Replace(c.Description, "\n", "<br>", -1)

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

//handle / requests to the server
func (cocktail *Cocktail) CocktailLandingHandler(w http.ResponseWriter, r *http.Request) {
	//log.Println("indexHandler: " + r.URL.Path[1:])

	var c model.Cocktail

	//log.Println(c)
	//apply the template page info to the index page
	cocktail.RenderTemplate(w, "index", &c)
}

func (cocktail *Cocktail) Init() {
	//Web Service and Web Page Handlers
	log.Println("Init in www/Cocktail.go")
	http.HandleFunc("/", cocktail.CocktailLandingHandler)
	http.HandleFunc("/cocktails", cocktail.CocktailHandler)
	http.HandleFunc("/search", cocktail.CocktailSearchHandler)

}
