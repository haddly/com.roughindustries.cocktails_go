//main/main
package main

import (
	"html/template"
	"log"
	"model"
	"net/http"
	"os"
	"path/filepath"
)

//render the page based on the name of the file provided
func renderTemplate(w http.ResponseWriter, tmpl string, c *model.Cocktail) {
	t, err := template.ParseFiles("./webcontent/" + tmpl + ".html")
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(w, c)
	if err != nil {
		log.Fatal(err)
	}
}

//handle / requests to the server
func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("indexHandler: " + r.URL.Path[1:])
	//parse the url and get the information after the localhost:8080/
	//stick that in the name
	//name := r.URL.Path[1:]
	//or setup a default for now
	name := "Jamaican Quaalude"
	c := &model.Cocktail{}
	for _, element := range model.Cocktails {
		// index is the index where we are
		// element is the element from someSlice for where we are
		if element.Name == name {
			c = &element
			break
		}
	}
	log.Println(c)
	//apply the template page info to the index page
	renderTemplate(w, "index", c)
}

func init() {
	//Web Service and Web Page Handlers
	http.HandleFunc("/", indexHandler)
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
	// This is handled in the app.yaml for google cloud platform deployments
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./webcontent/images"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./webcontent/css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./webcontent/js"))))
	http.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("./webcontent/fonts"))))
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "webcontent/favicon.ico")
	})

	log.Println("Added Handlers ... Starting Server\n")
	//this starts up the server
	http.ListenAndServe(":8080", nil)
}
