// Copyright 2017 Rough Industries LLC. All rights reserved.
//controller/www/page.go: Functions and handlers for dealing with pages.  This
//is the container for the pages that we serve over http.  The page struct
//represents all that can be sent to the templates.  I put in a gereric
//load page here that just does a processing wheel until the page is loaded.
//
package www

import (
	"github.com/golang/glog"
	"html/template"
	"model"
	"net/http"
	"strings"
	"time"
)

//Is the website public or private, i.e. do you want to offer admin
//accessability to the website
var AllowAdmin = true

//Use Google Analytics in the page.  This is to stop using google analytics
//code in the test environment.
var UseGA = false

//the page struct is all the things a template could display or use when it
//generates a page
type page struct {
	Username             string
	Redirect             string
	Authenticated        bool
	AllowAdmin           bool
	UseGA                bool
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
	Errors               map[string]string
	Messages             map[string]template.HTML
}

var counter = 0

//the page template renderer.  This should be the basic method for displaying
//all pages.
func (page *page) RenderPageTemplate(w http.ResponseWriter, r *http.Request, tmpl string) {
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
		//set the ip for this session
		page.UserSession.LastRemoteAddr = r.RemoteAddr
		page.UserSession.LastXForwardedFor = r.Header.Get("X-Forwarded-For")
		page.UserSession.LastSeenTime = time.Now()
		page.UserSession.CSRF = encrypt([]byte(page.UserSession.CSRFKey), page.UserSession.CSRFBase)
		SetSession(w, r, &page.UserSession, false)
	}
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

//The parse template files includes not only the page that is being requested
//but also the header, footer, google analytics, and navigation for
//provide a complete page
func parseTempFiles(tmpl string) (*template.Template, error) {
	t, e := template.ParseFiles("./view/webcontent/www/templates/"+tmpl+".html", "./view/webcontent/www/templates/head.html", "./view/webcontent/www/templates/ga.html", "./view/webcontent/www/templates/navbar.html", "./view/webcontent/www/templates/footer.html")
	return t, e
}

// The main index page handler
func LandingHandler(w http.ResponseWriter, r *http.Request, page *page) {
	//apply the template page info to the index page
	page.RenderPageTemplate(w, r, "index")
}

//The load page does a processing wheel until the actual page is loaded. It
//works as a redirect basically with the wheel showing till the redirected page
//is ready
func Load(w http.ResponseWriter, r *http.Request) {
	page := NewPage(w, r)
	glog.Infoln("Loader")
	glog.Infoln(r.URL.EscapedPath() + "?" + r.URL.RawQuery)
	glog.Infoln(strings.Replace(r.URL.EscapedPath(), "/load/", "", 1))
	redirect := strings.Replace(r.URL.EscapedPath(), "/load/", "", 1)
	if redirect == "" {
		page.Redirect = "index"
	} else {
		page.Redirect = redirect + "?" + r.URL.RawQuery
	}
	page.RenderPageTemplate(w, nil, "loader")
}

//An initialization function that provides an initialized page object
func NewPage(w http.ResponseWriter, r *http.Request) *page {
	p := new(page)
	p.Messages = make(map[string]template.HTML)
	p.Errors = make(map[string]string)
	p.AllowAdmin = AllowAdmin
	p.UseGA = UseGA
	if r != nil {
		p.UserSession, p.Authenticated = GetSession(w, r)
	}
	return p
}
