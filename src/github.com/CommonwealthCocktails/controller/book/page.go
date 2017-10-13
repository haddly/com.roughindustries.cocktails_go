// Copyright 2017 Rough Industries LLC. All rights reserved.
//controller/book/page.go: Functions and handlers for dealing with pages.  This
//is the container for the pages that we serve over http.  The page struct
//represents all that can be sent to the templates.  I put in a gereric
//load page here that just does a processing wheel until the page is loaded.
//
package book

import (
	"github.com/CommonwealthCocktails/model"
	"html/template"
	"net/http"
)

//the page struct is all the things a template could display or use when it
//generates a page
type page struct {
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
}

//the page template renderer.  This should be the basic method for displaying
//all pages.
func (page *page) RenderBookTemplate(w http.ResponseWriter, r *http.Request, tmpl string) {
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
		page.RenderBookTemplate(w, r, "404")
		return
	}
	err = t.ExecuteTemplate(w, "base", page)
	if err != nil {
		Error404(w, err)
		return
	}
}

//The parse template files includes not only the page that is being requested
//but also the header, footer, google analytics, and navigation for
//provide a complete page
func parseTempFiles(tmpl string) (*template.Template, error) {
	t, e := template.ParseFiles("./view/webcontent/book/templates/"+tmpl+".html", "./view/webcontent/book/templates/head.html", "./view/webcontent/book/templates/footer.html")
	return t, e
}

// The main index page handler
func LandingHandler(w http.ResponseWriter, r *http.Request, page *page) {
	//apply the template page info to the index page
	page.RenderBookTemplate(w, r, "index")
}

//An initialization function that provides an initialized page object
func NewPage(w http.ResponseWriter, r *http.Request) *page {
	p := new(page)
	return p
}
