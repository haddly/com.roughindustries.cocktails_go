package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// recipe:
//   - !recipeStep
//       ingredient:
//          componentName: Pineapple
//          componentType: 3
//       recipeCardinal: 1.0
//       recipeDoze: whole
//       recipeOrdinal: 1

type Page struct {
	Title       string
	Name        string
	Description string
	Recipe      Recipe
}

type Recipe struct {
	RecipeSteps []RecipeStep
}

type RecipeStep struct {
	Ingredient     Component
	RecipeCardinal float64
	RecipeDoze     Doze
	RecipeOrdinal  int
}

type Doze int

const (
	Shot = 1 + iota
	Ounce
	Whole
	Dash
)

var Dozes = [...]string{
	"Shot",
	"Ounce",
	"Whole",
	"Dash",
}

// String returns the English name of the doze ("Shot", "Ounce", ...).
func (d Doze) String() string { return Dozes[d-1] }

// Spirit (0)
// Liqueur (1)
// Wine (2)
// Mixer (3)
// Beer (4)

type ComponentType int

const (
	Spirit = 1 + iota
	Liqueur
	Wine
	Mixer
	Beer
)

var ComponentTypes = [...]string{
	"Spirit",
	"Liqueur",
	"Wine",
	"Mixer",
	"Beer",
}

type Component struct {
	ComponentName       string
	ComponentType       ComponentType
	ComponentTypeString string
}

// String returns the English name of the doze ("Shot", "Ounce", ...).
func (ct ComponentType) String() string { return ComponentTypes[ct-1] }

var cockatils = []Page{
	Page{
		Title:       "Jamaican Quaalude",
		Name:        "Jamaican Quaalude",
		Description: "I'll assume that this delicious cocktail's name is derived from its tropical flavors (Jamaican), and its mind numbing effects (Quaalude). With five spirits, and a bit of cream to blend it all together, this rich drink is a great dessert cocktail that will definitely keep the evening going. We hope you'll try our featured cocktail, the Jamaican Quaalude!",
		Recipe: Recipe{
			RecipeSteps: []RecipeStep{
				RecipeStep{
					Ingredient: Component{
						ComponentName:       "Pineapple",
						ComponentType:       Spirit,
						ComponentTypeString: ComponentTypes[Spirit-1],
					},
					RecipeCardinal: 1.0,
					RecipeDoze:     1,
					RecipeOrdinal:  0,
				},
			},
		},
	},
}

//render the page based on the name of the file provided
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, err := template.ParseFiles("./webcontent/" + tmpl + ".html")
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(w, p)
	if err != nil {
		log.Fatal(err)
	}
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
	log.Println(p)
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
