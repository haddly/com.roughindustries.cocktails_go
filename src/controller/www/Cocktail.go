package www

import (
	"html/template"
	"log"
	"math/rand"
	"model"
	"net/http"
)

type Cocktail struct {
}

var cocktail = new(Cocktail)

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
	log.Println(c)
	//apply the template page info to the index page
	cocktail.RenderTemplate(w, "index", c)
}

func init() {
	//Web Service and Web Page Handlers
	log.Println("Init in www/Cocktail.go")
	http.HandleFunc("/", cocktail.IndexHandler)
}
