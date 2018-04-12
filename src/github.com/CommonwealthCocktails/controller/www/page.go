// Copyright 2017 Rough Industries LLC. All rights reserved.
//controller/www/page.go: Functions and handlers for dealing with pages.  This
//is the container for the pages that we serve over http.  The page struct
//represents all that can be sent to the templates.  I put in a gereric
//load page here that just does a processing wheel until the page is loaded.
//
package www

import (
	"github.com/CommonwealthCocktails/model"
	log "github.com/sirupsen/logrus"
	"html/template"
	"net/http"
	//"strings"
	"github.com/microcosm-cc/bluemonday"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/google"
	"html"
	"math/rand"
	//"os"
	//"path/filepath"
	"errors"
	"math"
	"reflect"
	"regexp"
	"strings"
	"time"
)

//the page struct is all the things a template could display or use when it
//generates a page
type page struct {
	State                int
	BaseURL              string
	SubrouteURL          string
	Username             string
	Authenticated        bool
	AllowAdmin           bool
	UseGA                bool
	ShowAffiliateLinks   bool
	ReCAPTCHASiteKey     string
	ReCAPTCHASiteKeyInv  string
	TinymceAPIKey        string
	SocialSource         string
	IsForm               bool
	Locale               string
	Template             string
	Pagination           pagination
	CocktailSet          model.CocktailSet
	MetasByTypes         model.MetasByTypes
	Ingredients          model.ProductsByTypes
	NonIngredients       model.ProductsByTypes
	Cocktail             model.Cocktail
	Post                 model.Post
	Product              model.Product
	BaseProductWithBDG   model.BaseProductWithBDG
	Meta                 model.Meta
	Products             []model.Product
	Cocktails            []model.Cocktail
	Image                model.Image
	Posts                []model.Post
	CocktailsByAlphaNums model.CocktailsByAlphaNums
	Metas                []model.Meta
	ProductsByTypes      model.ProductsByTypes
	Doze                 []model.Doze
	DerivedProduct       model.DerivedProduct
	GroupProduct         model.GroupProduct
	User                 model.User
	UserSession          model.UserSession
	Search               model.Search
	SubmitButtonString   string
	View                 string
	Errors               map[string]string
	Messages             map[string]template.HTML
}

type pagination struct {
	NumPages           int
	HasPrev, HasNext   bool
	PrevPage, NextPage int
	ItemsPerPage       int
	CurrentPage        int
	NumItems           int
	PageOffset         int
	PageNums           []int
}

type translation struct {
	Translation string `json:"translation"`
}

var counter = 0

//Variables for use throughout the package
var (
	//OAuth
	googleOauthConfig   *oauth2.Config
	facebookOauthConfig *oauth2.Config

	// Some random string, random for each request
	// this way could create a memory leak sense I don't clear out the map ever, just a heads up
	oauthStateString = make(map[string]bool)

	//Default user is the user you can get into the system with at all times
	allowDefault    bool
	defaultUser     string
	defaultPassword string

	//reCAPTCHA
	sitekey    string
	re         reCAPTCHA
	sitekeyInv string
	reInv      reCAPTCHA

	//registration email
	registerFromEmailAddress string
	registerEmailPasswd      string

	//Is the website public or private, i.e. do you want to offer admin
	//accessability to the website
	AllowAdmin bool

	//Use Google Analytics in the page.  This is to stop using google analytics
	//code in the test environment.
	UseGA bool

	//show or don't show affiliate links
	ShowAffiliateLinks bool

	//the sessions are stored here in cookies
	store_key string

	TinymceAPIKey string
)

//Init variables from config
func PageInit() {
	log.Infoln("Login Init")
	//default user
	allowDefault = viper.GetBool("allowDefault")
	defaultUser = viper.GetString("defaultUser")
	//hash is = password
	defaultPassword = viper.GetString("defaultPassword")

	AllowAdmin = viper.GetBool("AllowAdmin")
	UseGA = viper.GetBool("UseGA")

	ShowAffiliateLinks = viper.GetBool("ShowAffiliateLinks")

	store_key = viper.GetString("cookieStoreKey")

	TinymceAPIKey = viper.GetString("TinymceAPIKey")

	//reCAPTCHA
	sitekey = viper.GetString("reCAPTCHASiteKey")
	re = reCAPTCHA{
		Secret: viper.GetString("reCAPTCHASecret"),
	}
	sitekeyInv = viper.GetString("reCAPTCHASiteKeyInv")
	reInv = reCAPTCHA{
		Secret: viper.GetString("reCAPTCHASecretInv"),
	}

	googleOauthConfig = &oauth2.Config{
		ClientID:     viper.GetString("googleOauthConfigClientID"),
		ClientSecret: viper.GetString("googleOauthConfigClientSecret"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	facebookOauthConfig = &oauth2.Config{
		ClientID:     viper.GetString("facebookOauthConfigClientID"),
		ClientSecret: viper.GetString("facebookOauthConfigClientSecret"),
		Scopes:       []string{"public_profile", "email", "pages_show_list", "manage_pages", "publish_pages"},
		Endpoint:     facebook.Endpoint,
	}

	//I am using gmail smtp.  If you have 2 step authentication get an app
	//password that corresponds to the from email account you use.
	registerFromEmailAddress = viper.GetString("registerFromEmailAddress")
	registerEmailPasswd = viper.GetString("registerEmailPasswd")

}

//the page template renderer.  This should be the basic method for displaying
//all pages.
func (page *page) RenderSetupTemplate(w http.ResponseWriter, r *http.Request, tmpl string) {
	page.Template = tmpl
	// CATCH ONLY HEADER START
	defer func() {
		// recover from panic if one occured. Set err to nil otherwise.
		if rec := recover(); rec != nil {
			Error404(w, rec, page.View)
		}
	}()
	// CATCH ONLY HEADER START
	// No CSRF check because we are in Setup state
	t, err := parseTempFiles(*page, tmpl)
	if err != nil {
		Error404(w, err, page.View)
		return
	}
	err = t.ExecuteTemplate(w, "base", page)
	if err != nil {
		Error404(w, err, page.View)
		return
	}
}

//the page template renderer.  This should be the basic method for displaying
//all pages.
func (this *page) RenderPageTemplate(w http.ResponseWriter, r *http.Request, tmpl string) {
	this.Template = tmpl
	// CATCH ONLY HEADER START
	defer func() {
		// recover from panic if one occured. Set err to nil otherwise.
		if rec := recover(); rec != nil {
			Error404(w, rec, this.View)
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
			log.Infoln(this.UserSession.CSRF)
			log.Infoln(this.UserSession.CSRFBase)
		}
		SetSession(w, r, &this.UserSession, false, this.View)
	}

	if this.View == "" {
		this.View = "www"
	}

	t, err := parseTempFiles(*this, tmpl)
	if err != nil {
		Error404(w, err, "www")
		return
	}

	err = t.ExecuteTemplate(w, "base", this)
	if err != nil {
		Error404(w, err, "www")
		return
	}
}

//The parse template files includes not only the page that is being requested
//but also the header, footer, google analytics, and navigation for
//provide a complete page
func parseTempFiles(page page, tmpl string) (*template.Template, error) {
	funcMap := template.FuncMap{
		"now": func() string { return time.Now().UTC().Format(time.RFC3339) },
		"attr": func(s string) template.HTMLAttr {
			log.Infoln("attr" + s)
			log.Infoln(template.HTMLAttr(s))
			return template.HTMLAttr(s)
		},
		"safe": func(s string) template.HTML {
			log.Infoln("safe" + s)
			log.Infoln(template.HTML(s))
			return template.HTML(s)
		},
		"sanitize": func(temp template.HTML) string {
			pSP := bluemonday.StrictPolicy()
			unescaped := html.UnescapeString(pSP.Sanitize(string(temp)))
			// Make a Regex to say we only want
			reg, err := regexp.Compile("[^a-zA-Z0-9 ().-]+")
			if err != nil {
				log.Fatal(err)
			}
			processedString := reg.ReplaceAllString(string(unescaped), "")
			return processedString
		},
		"replaceSpaceWithDash": func(temp template.HTML) string {
			pSP := bluemonday.StrictPolicy()
			unescaped := html.UnescapeString(pSP.Sanitize(string(temp)))
			// Make a Regex to say we only want
			reg, err := regexp.Compile("[^a-zA-Z0-9 -]+")
			if err != nil {
				log.Fatal(err)
			}
			processedString := reg.ReplaceAllString(string(unescaped), "")
			processedString = strings.Replace(processedString, " ", "-", -1)
			return processedString
		},
		"mksliceCocktail": func(c []model.Cocktail, size int) []model.Cocktail {
			rv := reflect.ValueOf(c)
			swap := reflect.Swapper(c)
			length := rv.Len()
			for i := length - 1; i > 0; i-- {
				j := rand.Intn(i + 1)
				swap(i, j)
			}
			if len(c) > size {
				return c[0:size]
			} else {
				return c
			}
		},
		"dayOfYear": func() string {
			return time.Now().Format("January-2")
		},
		"showAffiliateLinks": func() bool {
			return viper.GetBool("ShowAffiliateLinks")
		},
		"I18n": func(view string, locale string, key string) string {
			trans := ""
			if locale == "" {
				locale = "en-US"
			}
			t := GetI18nMap()
			trans = t[view].Locales[locale].Locale[key].Translation
			return trans
		},
		"dict": func(values ...interface{}) (map[string]interface{}, error) {
			if len(values)%2 != 0 {
				return nil, errors.New("invalid dict call")
			}
			dict := make(map[string]interface{}, len(values)/2)
			for i := 0; i < len(values); i += 2 {
				key, ok := values[i].(string)
				if !ok {
					return nil, errors.New("dict keys must be strings")
				}
				dict[key] = values[i+1]
			}
			return dict, nil
		},
	}
	t, e := template.New("page").Funcs(funcMap).ParseFiles("./view/webcontent/"+page.View+"/templates/"+tmpl+".html", "./view/webcontent/"+page.View+"/templates/head.html", "./view/webcontent/"+page.View+"/templates/loader.html", "./view/webcontent/"+page.View+"/templates/ga.html", "./view/webcontent/"+page.View+"/templates/navbar.html", "./view/webcontent/"+page.View+"/templates/footer.html")
	return t, e
}

// The main index page handler
func LandingHandler(w http.ResponseWriter, r *http.Request, page *page) {
	log.Infoln("Landing")
	//apply the template page info to the index page
	page.RenderPageTemplate(w, r, "index")
}

// The main index page handler
func TestHandler(w http.ResponseWriter, r *http.Request, page *page) {
	log.Infoln("Test")
	//apply the template page info to the index page
	page.RenderPageTemplate(w, r, "index")
}

//An initialization function that provides an initialized page object
func NewPage(w http.ResponseWriter, r *http.Request, site string) *page {
	p := new(page)
	p.Messages = make(map[string]template.HTML)
	p.Errors = make(map[string]string)
	p.AllowAdmin = AllowAdmin
	p.UseGA = UseGA
	p.ReCAPTCHASiteKey = sitekey
	p.ReCAPTCHASiteKeyInv = sitekeyInv
	p.TinymceAPIKey = TinymceAPIKey
	if r != nil {
		p.UserSession, p.Authenticated = GetSession(w, r, site)
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
	p.TinymceAPIKey = TinymceAPIKey
	p.UserSession = *new(model.UserSession)
	p.Authenticated = false
	return p
}

// pages start at 1 - not 0
func PaginationCalculate(page *page, currentPage, itemsPerPage, numItems, pageOffset int) {

	page.Pagination.CurrentPage = currentPage
	page.Pagination.ItemsPerPage = itemsPerPage
	page.Pagination.NumItems = numItems
	page.Pagination.PageOffset = pageOffset
	page.Pagination.PageNums = []int{}

	// calculate number of pages
	d := float64(page.Pagination.NumItems) / float64(page.Pagination.ItemsPerPage)
	page.Pagination.NumPages = int(math.Ceil(d))

	// HasPrev, HasNext?
	page.Pagination.HasPrev = page.Pagination.CurrentPage-page.Pagination.PageOffset > 1
	page.Pagination.HasNext = page.Pagination.CurrentPage+page.Pagination.PageOffset < page.Pagination.NumPages

	// calculate them
	if page.Pagination.HasPrev {
		page.Pagination.PrevPage = page.Pagination.CurrentPage - page.Pagination.PageOffset - 1
	}
	if page.Pagination.HasNext {
		page.Pagination.NextPage = page.Pagination.CurrentPage + page.Pagination.PageOffset + 1
	}

	upperLimit := page.Pagination.NumPages
	if page.Pagination.NumPages > page.Pagination.CurrentPage+page.Pagination.PageOffset {
		upperLimit = page.Pagination.CurrentPage + page.Pagination.PageOffset
	}

	for i := page.Pagination.CurrentPage - page.Pagination.PageOffset; i <= upperLimit; i++ {
		if i >= 1 {
			page.Pagination.PageNums = append(page.Pagination.PageNums, i)
		}
	}

	log.Infoln(page.Pagination)

	return
}
