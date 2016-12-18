package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type Page struct {
	Title       string
	Name        string
	Description string
}

var cockatils = []Page{
	Page{
		Title:       "Jamaican Quaalude",
		Name:        "Jamaican Quaalude",
		Description: "I'll assume that this delicious cocktail's name is derived from its tropical flavors (Jamaican), and its mind numbing effects (Quaalude). With five spirits, and a bit of cream to blend it all together, this rich drink is a great dessert cocktail that will definitely keep the evening going. We hope you'll try our featured cocktail, the Jamaican Quaalude!",
	},
}

//render the page based on the name of the file provided
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, err := template.ParseFiles("./webcontent/" + tmpl + ".html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, p)
}

//handle / requests to the server
func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("indexHandler: " + r.URL.Path[1:])
	//parse the url and get the information after the localhost:8080/
	//stick that in the name
	name := r.URL.Path[1:]
	p := &Page{}
	for _, element := range cockatils {
		// index is the index where we are
		// element is the element from someSlice for where we are
		if element.Name == name {
			p = &element
			break
		}
	}
	//apply the template page info to the index page
	renderTemplate(w, "index", p)
}

func init() {
	//Web Service and Web Page Handlers
	//http.HandleFunc("/", indexHandler)
}

//where it all starts
func main() {
	log.Println("Starting ... \n")
	//print out the current directory
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(dir)

	// Mandatory root-based resources and redirects for other resources
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./webcontent/images"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./webcontent/css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./webcontent/js"))))
	http.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("./webcontent/fonts"))))
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "webcontent/favicon.ico")
	})

	//Web Service and Web Page Handlers
	http.HandleFunc("/", indexHandler)

	log.Println("Added Handlers ... Starting Server\n")
	//this starts up the server
	http.ListenAndServe(":8080", nil)
}
