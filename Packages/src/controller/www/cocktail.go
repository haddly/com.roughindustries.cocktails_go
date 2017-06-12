package www

import (
	"html/template"
	"log"
	"model"
	"net/http"
	"net/url"
	"strconv"
)

type Cocktail struct {
}

func (cocktail *Cocktail) RenderPageTemplate(w http.ResponseWriter, tmpl string, page *Page) {
	t, err := parseTempFiles(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	err = t.ExecuteTemplate(w, "base", page)
	if err != nil {
		log.Fatal(err)
	}
}

func parseTempFiles(tmpl string) (*template.Template, error) {
	t, e := template.ParseFiles("./view/webcontent/www/templates/"+tmpl+".html", "./view/webcontent/www/templates/head.html", "./view/webcontent/www/templates/ga.html", "./view/webcontent/www/templates/navbar.html", "./view/webcontent/www/templates/footer.html")
	return t, e
}

func (cocktail *Cocktail) CocktailHandler(w http.ResponseWriter, r *http.Request) {
	var cs model.CocktailSet
	var page Page
	page.Username = GetUserName(r)
	u, err := url.Parse(r.URL.String())
	if err != nil {
		cocktail.RenderPageTemplate(w, "404", &page)
	}
	//log.Println("Product: " + r.URL.Path[1:])
	m, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		cocktail.RenderPageTemplate(w, "404", &page)
	}
	if len(m["ID"]) == 0 {
		cocktail.RenderPageTemplate(w, "404", &page)
	} else {
		//log.Println("ID: " + m["ID"][0])

		//apply the template page info to the index page
		id, _ := strconv.Atoi(m["ID"][0])
		if len(model.Products) <= id-1 {
			cocktail.RenderPageTemplate(w, "404", &page)
		} else {
			cs = model.GetCocktailByID(id)
			page.CocktailSet = cs
			//log.Println(c)
			//apply the template page info to the index page
			cocktail.RenderPageTemplate(w, "cocktail", &page)
		}
	}
}

//handle / requests to the server
func (cocktail *Cocktail) CocktailsHandler(w http.ResponseWriter, r *http.Request) {
	//log.Println("indexHandler: " + r.URL.Path[1:])
	var cs model.CocktailSet
	var c []model.Cocktail
	c = model.GetCocktails()
	cs.ChildCocktails = c
	var page Page
	page.CocktailSet = cs
	page.Username = GetUserName(r)
	//log.Println(c)
	//apply the template page info to the index page
	cocktail.RenderPageTemplate(w, "cocktails", &page)
}

//handle / requests to the server
func (cocktail *Cocktail) CocktailsIndexHandler(w http.ResponseWriter, r *http.Request) {
	//log.Println("indexHandler: " + r.URL.Path[1:])

	var m model.MetasByTypes
	m = model.GetMetaByTypes(true, true)

	var page Page
	page.MetasByTypes = m
	page.Username = GetUserName(r)
	//log.Println(c)
	//apply the template page info to the index page
	cocktail.RenderPageTemplate(w, "cocktailsindex", &page)
}

//handle / requests to the server
func (cocktail *Cocktail) CocktailsByMetaIDHandler(w http.ResponseWriter, r *http.Request) {
	var cs model.CocktailSet
	var page Page
	page.Username = GetUserName(r)
	u, err := url.Parse(r.URL.String())
	if err != nil {
		cocktail.RenderPageTemplate(w, "404", &page)
	}
	m, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		cocktail.RenderPageTemplate(w, "404", &page)
	}
	if len(m["ID"]) == 0 {
		cocktail.RenderPageTemplate(w, "404", &page)
	} else {
		id, _ := strconv.Atoi(m["ID"][0])
		log.Println("Meta ID: " + m["ID"][0])
		var inMeta model.Meta
		inMeta.ID = id
		var c []model.Cocktail
		c = model.SelectCocktailsByMeta(inMeta)
		cs.ChildCocktails = c
		meta := model.SelectMeta(inMeta)
		cs.Metadata = meta[0]
		page.CocktailSet = cs
		cocktail.RenderPageTemplate(w, "cocktails", &page)
	}
}

func (cocktail *Cocktail) CocktailsByProductIDHandler(w http.ResponseWriter, r *http.Request) {
	var cs model.CocktailSet
	var page Page
	page.Username = GetUserName(r)
	u, err := url.Parse(r.URL.String())
	if err != nil {
		cocktail.RenderPageTemplate(w, "404", &page)
	}
	m, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		cocktail.RenderPageTemplate(w, "404", &page)
	}
	if len(m["ID"]) == 0 {
		cocktail.RenderPageTemplate(w, "404", &page)
	} else {
		id, _ := strconv.Atoi(m["ID"][0])
		log.Println("Product ID: " + m["ID"][0])
		if len(model.Products) <= id-1 {
			cocktail.RenderPageTemplate(w, "404", &page)
		} else {
			var inProduct model.Product
			inProduct.ID = id
			var c []model.Cocktail
			c = model.SelectCocktailsByProduct(inProduct)
			cs.ChildCocktails = c
			page.CocktailSet = cs
			cocktail.RenderPageTemplate(w, "cocktails", &page)
		}
	}
}

//handle / requests to the server
func (cocktail *Cocktail) CocktailSearchHandler(w http.ResponseWriter, r *http.Request) {
	var page Page
	page.Username = GetUserName(r)
	//log.Println(c)
	//apply the template page info to the index page
	cocktail.RenderPageTemplate(w, "search", &page)
}

//handle / requests to the server
func (cocktail *Cocktail) CocktailLandingHandler(w http.ResponseWriter, r *http.Request) {
	//log.Println("indexHandler: " + r.URL.Path[1:])

	var cs model.CocktailSet

	var page Page
	page.CocktailSet = cs
	page.Username = GetUserName(r)
	//log.Println(c)
	//apply the template page info to the index page
	cocktail.RenderPageTemplate(w, "index", &page)
}

func (cocktail *Cocktail) Init() {
	//Web Service and Web Page Handlers
	log.Println("Init in www/Cocktail.go")
	http.HandleFunc("/", cocktail.CocktailLandingHandler)
	http.HandleFunc("/cocktail", cocktail.CocktailHandler)
	http.HandleFunc("/cocktails", cocktail.CocktailsHandler)
	http.HandleFunc("/cocktailsindex", cocktail.CocktailsIndexHandler)
	http.HandleFunc("/cocktailsByMetaID", cocktail.CocktailsByMetaIDHandler)
	http.HandleFunc("/cocktailsByProductID", cocktail.CocktailsByProductIDHandler)
}
