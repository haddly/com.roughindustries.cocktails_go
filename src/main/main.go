package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type Page struct {
	Title string
	Name  string
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
	//load of the name and title for doing tempating
	p := &Page{Title: "Commonwealth Cocktails", Name: name}
	//apply the template page info to the index page
	renderTemplate(w, "index", p)
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
