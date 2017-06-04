package www

import (
	"html/template"
	"log"
	"model"
	"net/http"
)

type Search struct {
}

//render the page based on the name of the file provided
func (cocktail *Search) RenderTemplate(w http.ResponseWriter, tmpl string, cs *model.CocktailSearch) {
	t, err := template.ParseFiles("./view/webcontent/www/templates/"+tmpl+".html", "./view/webcontent/www/templates/head.html", "./view/webcontent/www/templates/ga.html", "./view/webcontent/www/templates/navbar.html", "./view/webcontent/www/templates/footer.html")
	if err != nil {
		log.Fatal(err)
	}

	//c.Description = template.HTMLEscapeString(c.Description)
	//log.Println(c.Description)
	//c.Description = strings.Replace(c.Description, "\n", "<br>", -1)

	err = t.ExecuteTemplate(w, "base", cs)
	if err != nil {
		log.Fatal(err)
	}
}

//handle / requests to the server
func (cocktail *Search) CocktailSearchHandler(w http.ResponseWriter, r *http.Request) {
	//log.Println("indexHandler: " + r.URL.Path[1:])

	var cs model.CocktailSearch
	cs = model.GetCocktailSearch()

	//log.Println(c)
	//apply the template page info to the index page
	cocktail.RenderTemplate(w, "search", &cs)
}

func (cocktail *Search) Init() {
	log.Println("Init in www/search.go")
	http.HandleFunc("/search", cocktail.CocktailSearchHandler)

}
