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

type Cocktail struct {
	Title           string
	Name            string
	Description     string
	Recipe          Recipe
	CocktailMethod  string
	Garnish         []Component
	Image           string
	ImageSourceName string
	ImageSourceLink string
	Drinkware       []Component
	Tool            []Component
	SourceName      string
	SourceLink      string
	Rating          float32
	Flavor          []Meta
	Type            []Meta
	BaseSpirit      []Meta
	Served          []Meta
	Technique       []Meta
	Strength        []Meta
	Difficulty      []Meta
	TOD             []Meta
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
	"shot",
	"oz.",
	"whole",
	"dash",
}

// String returns the English name of the doze ("Shot", "Ounce", ...).
func (d Doze) String() string { return Dozes[d-1] }

type ComponentType int

const (
	Spirit = 1 + iota
	Liqueur
	Wine
	Mixer
	Beer
	Garnish
	Drinkware
	Tool
)

var ComponentTypeStrings = [...]string{
	"Spirit",
	"Liqueur",
	"Wine",
	"Mixer",
	"Beer",
	"Garnish",
	"Drinkware",
	"Tool",
}

// String returns the English name of the componenttype ("Spirit", "Liqueur", ...).
func (ct ComponentType) String() string { return ComponentTypeStrings[ct-1] }

type Component struct {
	ComponentName string
	ComponentType ComponentType
}

type MetaType int

const (
	Flavor = 1 + iota
	BaseSpirit
	Type
	Served
	Technique
	Strength
	Difficulty
	TOD
)

var MetaTypeStrings = [...]string{
	"Flavor",
	"Base Spirit",
	"Type",
	"Served",
	"Technique",
	"Strength",
	"Difficulty",
	"Time of Day",
}

// String returns the English name of the metatype ("Flavor", "Base Spirit", ...).
func (mt MetaType) String() string { return MetaTypeStrings[mt-1] }

type Meta struct {
	MetaName string
	MetaType MetaType
}

var cockatils = []Cocktail{
	Cocktail{
		Title:           "Jamaican Quaalude",
		Name:            "Jamaican Quaalude",
		Description:     "I'll assume that this delicious cocktail's name is derived from its tropical flavors (Jamaican), and its mind numbing effects (Quaalude). With five spirits, and a bit of cream to blend it all together, this rich drink is a great dessert cocktail that will definitely keep the evening going. We hope you'll try our featured cocktail, the Jamaican Quaalude!",
		CocktailMethod:  "Combine all of the ingredients in an ice filled cocktail shaker.  Cover, shake well, and pour into a Rocks glass.  Add a couple of sipping straws, garnish accordingly.",
		Image:           "jamaican_quaalude_750_750.png",
		ImageSourceName: "Unknown",
		ImageSourceLink: "",
		SourceName:      "Hampton Roads Happy Hour",
		SourceLink:      "http://hamptonroadshappyhour.com/jamaican-quaalude",
		Rating:          5.0,
		Tool: []Component{
			Component{
				ComponentName: "Shaker",
				ComponentType: Tool,
			},
		},
		Strength: []Meta{
			Meta{
				MetaName: "Medium",
				MetaType: Flavor,
			},
		},
		Difficulty: []Meta{
			Meta{
				MetaName: "Easy",
				MetaType: Difficulty,
			},
		},
		TOD: []Meta{
			Meta{
				MetaName: "Evening",
				MetaType: TOD,
			},
		},
		Flavor: []Meta{
			Meta{
				MetaName: "Creamy",
				MetaType: Flavor,
			},
		},
		Type: []Meta{
			Meta{
				MetaName: "Creamy",
				MetaType: Type,
			},
		},
		Served: []Meta{
			Meta{
				MetaName: "On the Rocks",
				MetaType: Served,
			},
		},
		Technique: []Meta{
			Meta{
				MetaName: "Shaken",
				MetaType: Technique,
			},
		},
		BaseSpirit: []Meta{},
		Garnish: []Component{
			Component{
				ComponentName: "Cherry",
				ComponentType: Garnish,
			},
			Component{
				ComponentName: "Slice of Starfruit",
				ComponentType: Garnish,
			},
		},
		Drinkware: []Component{
			Component{
				ComponentName: "Old Fashioned",
				ComponentType: Drinkware,
			},
		},
		Recipe: Recipe{
			RecipeSteps: []RecipeStep{
				//1 oz. Kahlua
				RecipeStep{
					Ingredient: Component{
						ComponentName: "Kahlua",
						ComponentType: Liqueur,
					},
					RecipeCardinal: 1.0,
					RecipeDoze:     Ounce,
					RecipeOrdinal:  0,
				},
				//1 oz. Coconut Rum
				RecipeStep{
					Ingredient: Component{
						ComponentName: "Coconut Rum",
						ComponentType: Spirit,
					},
					RecipeCardinal: 1.0,
					RecipeDoze:     Ounce,
					RecipeOrdinal:  1,
				},
				//1 oz. Baileys Irish Cream
				RecipeStep{
					Ingredient: Component{
						ComponentName: "Irish Cream",
						ComponentType: Liqueur,
					},
					RecipeCardinal: 1.0,
					RecipeDoze:     Ounce,
					RecipeOrdinal:  2,
				},
				//.5 oz Amaretto
				RecipeStep{
					Ingredient: Component{
						ComponentName: "Amaretto",
						ComponentType: Liqueur,
					},
					RecipeCardinal: 0.5,
					RecipeDoze:     Ounce,
					RecipeOrdinal:  3,
				},
				//.5 oz Frangelico
				RecipeStep{
					Ingredient: Component{
						ComponentName: "Frangelico",
						ComponentType: Liqueur,
					},
					RecipeCardinal: 0.5,
					RecipeDoze:     Ounce,
					RecipeOrdinal:  4,
				},
				//1 oz Cream
				RecipeStep{
					Ingredient: Component{
						ComponentName: "Cream",
						ComponentType: Mixer,
					},
					RecipeCardinal: 1.0,
					RecipeDoze:     Ounce,
					RecipeOrdinal:  5,
				},
			},
		},
	},
}

//render the page based on the name of the file provided
func renderTemplate(w http.ResponseWriter, tmpl string, c *Cocktail) {
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
	c := &Cocktail{}
	for _, element := range cockatils {
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
