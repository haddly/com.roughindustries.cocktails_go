// Copyright 2017 Rough Industries LLC. All rights reserved.
//controller/www/page.go: Functions and handlers for dealing with pages.  This
//is the container for the pages that we serve over http.  The page struct
//represents all that can be sent to the templates.  I put in a gereric
//load page here that just does a processing wheel until the page is loaded.
//
package www

import (
	"github.com/CommonwealthCocktails/model"
	"github.com/golang/glog"
	"html/template"
	"net/http"
	//"strings"
	"time"
)

var (
	//Is the website public or private, i.e. do you want to offer admin
	//accessability to the website
	AllowAdmin = true

	//Use Google Analytics in the page.  This is to stop using google analytics
	//code in the test environment.
	UseGA = false
)

//the page struct is all the things a template could display or use when it
//generates a page
type page struct {
	State                int
	BaseURL              string
	Username             string
	Authenticated        bool
	AllowAdmin           bool
	UseGA                bool
	ReCAPTCHASiteKey     string
	ReCAPTCHASiteKeyInv  string
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
	UserSession          model.UserSession
	SubmitButtonString   string
	Errors               map[string]string
	Messages             map[string]template.HTML
}

var counter = 0

//the page template renderer.  This should be the basic method for displaying
//all pages.
func (page *page) RenderSetupTemplate(w http.ResponseWriter, r *http.Request, tmpl string) {
	// CATCH ONLY HEADER START
	defer func() {
		// recover from panic if one occured. Set err to nil otherwise.
		if rec := recover(); rec != nil {
			Error404(w, rec)
		}
	}()
	// CATCH ONLY HEADER START
	// No CSRF check because we are in Setup state
	t, err := parseTempFiles(tmpl)
	if err != nil {
		Error404(w, err)
		page.RenderPageTemplate(w, r, "404")
		return
	}
	err = t.ExecuteTemplate(w, "base", page)
	if err != nil {
		Error404(w, err)
		return
	}
}

//the page template renderer.  This should be the basic method for displaying
//all pages.
func (this *page) RenderPageTemplate(w http.ResponseWriter, r *http.Request, tmpl string) {
	// CATCH ONLY HEADER START
	defer func() {
		// recover from panic if one occured. Set err to nil otherwise.
		if rec := recover(); rec != nil {
			Error404(w, rec)
		}
	}()
	// CATCH ONLY HEADER START
	// setup the CSRF id for this page
	if r != nil {
		this.UserSession.LastSeenTime = time.Now()
		if len(this.UserSession.CSRFKey) == 0 || this.UserSession.CSRFBase == "" {
			this.UserSession.CSRF = ""
		} else {
			this.UserSession.CSRF = encrypt([]byte(this.UserSession.CSRFKey), this.UserSession.CSRFBase)
			glog.Infoln(this.UserSession.CSRF)
			glog.Infoln(this.UserSession.CSRFBase)
		}
		SetSession(w, r, &this.UserSession, false)
	}
	t, err := parseTempFiles(tmpl)
	if err != nil {
		Error404(w, err)
		this.RenderPageTemplate(w, r, "404")
		return
	}
	err = t.ExecuteTemplate(w, "base", this)
	if err != nil {
		Error404(w, err)
		return
	}
}

//The parse template files includes not only the page that is being requested
//but also the header, footer, google analytics, and navigation for
//provide a complete page
func parseTempFiles(tmpl string) (*template.Template, error) {
	t, e := template.ParseFiles("./view/webcontent/www/templates/"+tmpl+".html", "./view/webcontent/www/templates/head.html", "./view/webcontent/www/templates/loader.html", "./view/webcontent/www/templates/ga.html", "./view/webcontent/www/templates/navbar.html", "./view/webcontent/www/templates/footer.html")
	return t, e
}

// The main index page handler
func LandingHandler(w http.ResponseWriter, r *http.Request, page *page) {
	//apply the template page info to the index page
	page.RenderPageTemplate(w, r, "index")
}

//An initialization function that provides an initialized page object
func NewPage(w http.ResponseWriter, r *http.Request) *page {
	p := new(page)
	p.Messages = make(map[string]template.HTML)
	p.Errors = make(map[string]string)
	p.AllowAdmin = AllowAdmin
	p.UseGA = UseGA
	p.ReCAPTCHASiteKey = sitekey
	p.ReCAPTCHASiteKeyInv = sitekeyInv
	if r != nil {
		p.UserSession, p.Authenticated = GetSession(w, r)
	}
	return p
}

//Specific initialization for setup pages, i.e. no session data
func NewSetupPage(w http.ResponseWriter, r *http.Request) *page {
	p := new(page)
	p.Messages = make(map[string]template.HTML)
	p.Errors = make(map[string]string)
	p.AllowAdmin = AllowAdmin
	p.UseGA = UseGA
	p.ReCAPTCHASiteKey = sitekey
	p.ReCAPTCHASiteKeyInv = sitekeyInv
	p.UserSession = *new(model.UserSession)
	p.Authenticated = false
	return p
}
