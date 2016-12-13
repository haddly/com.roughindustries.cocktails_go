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

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
        
	t, err := template.ParseFiles("./webcontent/" + tmpl + ".html")
        if err != nil {
            log.Fatal(err)
        }
        t.Execute(w, p)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("indexHandler: " + r.URL.Path[1:])
	name := r.URL.Path[1:]
	p := &Page{Title: "Commonwealth Cocktails", Name: name}
	renderTemplate(w, "cocktailTemplate", p)
}

func jamaicanQuaaludeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("jamaicanQuaaludeHandler: " + r.URL.String())
	p := &Page{Title: "Commonwealth Cocktails", Name: "Jamaican Quaalude"}
	renderTemplate(w, "jamaicanQuaalude", p)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("searchHandler: " + r.URL.String())
	p := &Page{Title: "Commonwealth Cocktails", Name: "Magarita"}
	renderTemplate(w, "search", p)
}

func main() {
	log.Println("Starting ... \n")
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
	http.HandleFunc("/jamaicanQuaalude", jamaicanQuaaludeHandler)
	http.HandleFunc("/search", searchHandler)

	log.Println("Added Handlers ... Starting Server\n")
	http.ListenAndServe(":8080", nil)
}
