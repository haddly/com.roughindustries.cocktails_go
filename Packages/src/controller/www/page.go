package www

import (
	"html/template"
	"log"
	"model"
	"net/http"
	"strings"
)

func (page *page) RenderPageTemplate(w http.ResponseWriter, tmpl string) {
	// CATCH ONLY HEADER START
	defer func() {
		// recover from panic if one occured. Set err to nil otherwise.
		if rec := recover(); rec != nil {
			Error404(w, rec)
		}
	}()
	// CATCH ONLY HEADER START
	t, err := parseTempFiles(tmpl)
	if err != nil {
		Error404(w, err)
		page.RenderPageTemplate(w, "404")
		return
	}
	err = t.ExecuteTemplate(w, "base", page)
	if err != nil {
		Error404(w, err)
		return
	}
}

func parseTempFiles(tmpl string) (*template.Template, error) {
	t, e := template.ParseFiles("./view/webcontent/www/templates/"+tmpl+".html", "./view/webcontent/www/templates/head.html", "./view/webcontent/www/templates/ga.html", "./view/webcontent/www/templates/navbar.html", "./view/webcontent/www/templates/footer.html")
	return t, e
}

func Load(w http.ResponseWriter, r *http.Request) {
	// STANDARD HANDLER HEADER START
	// catch all errors and return 404
	defer func() {
		// recover from panic if one occured. Set err to nil otherwise.
		if rec := recover(); rec != nil {
			Error404(w, rec)
		}
	}()
	page := NewPage()
	//page.Username, page.Authenticated = GetSession(r)
	// STANDARD HANLDER HEADER END
	log.Println("Loader")
	log.Println(r.URL.EscapedPath() + "?" + r.URL.RawQuery)
	log.Println(strings.Replace(r.URL.EscapedPath(), "/load/", "", 1))
	redirect := strings.Replace(r.URL.EscapedPath(), "/load/", "", 1)
	if redirect == "" {
		page.Redirect = "index"
	} else {
		page.Redirect = redirect + "?" + r.URL.RawQuery
	}
	page.RenderPageTemplate(w, "loader")
}

type page struct {
	Username             string
	Redirect             string
	Authenticated        bool
	CocktailSearch       model.CocktailSearch
	CocktailSet          model.CocktailSet
	MetasByTypes         model.MetasByTypes
	Ingredients          model.ProductsByTypes
	NonIngredients       model.ProductsByTypes
	Cocktail             model.Cocktail
	Product              model.Product
	BaseProductWithBDG   model.BaseProductWithBDG
	Meta                 model.Meta
	Products             []model.Product
	Cocktails            []model.Cocktail
	CocktailsByAlphaNums model.CocktailsByAlphaNums
	Metas                []model.Meta
	ProductsByTypes      model.ProductsByTypes
	Doze                 []model.Doze
	DerivedProduct       model.DerivedProduct
	GroupProduct         model.GroupProduct
	User                 model.User
	Errors               map[string]string
	Messages             map[string]template.HTML
}

func NewPage() page {
	var p page
	p.Messages = make(map[string]template.HTML)
	p.Errors = make(map[string]string)
	return p
}

// Init
func (page *page) Init() {
	//Web Service and Web Page Handlers
	log.Println("Init in www/page.go")
	http.HandleFunc("/load/", Load)
}
