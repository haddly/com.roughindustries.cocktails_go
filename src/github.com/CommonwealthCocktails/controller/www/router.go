// Copyright 2017 Rough Industries LLC. All rights reserved.
//controller/www/router.go: sets up all the routing for the webapp
package www

import (
	"github.com/CommonwealthCocktails/connectors"
	"github.com/CommonwealthCocktails/model"
	"github.com/asaskevich/govalidator"
	"github.com/golang/glog"
	"github.com/microcosm-cc/bluemonday"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"strings"
)

//The integer values for the producttype enumeration
type SystemStateConst int

const (
	Initialize = 1 + iota
	Setup
	Ready
	Down
)

//The string values for the producttype enumeration
var SystemStateStrings = [...]string{
	"Initialize",
	"Setup",
	"Ready",
	"Down",
}

// String returns the English name of the SystemStateConst ("Initialize", "Setup", ...).
func (ss SystemStateConst) String() string { return SystemStateStrings[ss-1] }

var (
	State         = Initialize
	BaseURL       = ""
	Valid_Tables  = []string{"cc_usermeta", "cc_usermetatype", "altIngredient", "altnames", "cocktail", "cocktailToAKAs", "cocktailToAltNames", "cocktailToMetas", "cocktailToPosts", "cocktailToProducts", "cocktailToRecipe", "derivedProduct", "doze", "groupProduct", "grouptype", "meta", "metatype", "oauth", "product", "producttype", "post", "recipe", "recipeToRecipeSteps", "recipestep", "userroles", "users", "usersessions"}
	Ignore_Tables = []string{"sqlite_sequence"}
)

//Init to setup the http handlers
func WWWRouterInit() {
	glog.Infoln("Init in www/router.go")
	glog.Infoln(viper.GetString("BaseURL"))
	BaseURL = viper.GetString("BaseURL")
	//Inits
	PageInit()
	//Page Routing
	http.Handle("/", RecoverHandler(MethodsHandler(PageHandler(LandingHandler), "GET")))
	//Cocktail Routing
	http.Handle("/cocktail", RecoverHandler(MethodsHandler(PageHandler(CocktailHandler), "GET")))
	http.Handle("/cocktailOfTheDay", RecoverHandler(MethodsHandler(PageHandler(CocktailOfTheDayHandler), "GET")))
	http.Handle("/cocktailTop100", RecoverHandler(MethodsHandler(PageHandler(CocktailsTop100Handler), "GET")))
	http.Handle("/cocktails", RecoverHandler(MethodsHandler(PageHandler(CocktailsHandler), "GET")))
	http.Handle("/cocktailsindex", RecoverHandler(MethodsHandler(PageHandler(CocktailsIndexHandler), "GET")))
	http.Handle("/cocktailsindex/", RecoverHandler(MethodsHandler(PageHandler(CocktailsIndexHandler), "GET")))
	http.Handle("/cocktailsByMetaID", RecoverHandler(MethodsHandler(PageHandler(CocktailsByMetaIDHandler), "GET")))
	http.Handle("/cocktailsByProductID", RecoverHandler(MethodsHandler(PageHandler(CocktailsByProductIDHandler), "GET")))
	http.Handle("/cocktailsByIngredientID", RecoverHandler(MethodsHandler(PageHandler(CocktailsByIngredientIDHandler), "GET")))
	http.Handle("/cocktailModForm", RecoverHandler(MethodsHandler(VandAPageHandler(false, false, true, ValidateCocktail, nil, CocktailModFormHandler, LandingHandler), "GET")))
	http.Handle("/cocktailMod", RecoverHandler(MethodsHandler(VandAPageHandler(false, false, false, ValidateCocktail, RequiredCocktailMod, CocktailModHandler, CocktailModFormHandler), "POST")))
	//Meta Routing
	http.Handle("/metaModForm", RecoverHandler(MethodsHandler(VandAPageHandler(false, false, true, ValidateMeta, nil, MetaModFormHandler, LandingHandler), "GET")))
	http.Handle("/metaMod", RecoverHandler(MethodsHandler(VandAPageHandler(false, false, false, ValidateMeta, RequiredMetaMod, MetaModHandler, MetaModFormHandler), "POST")))
	//Products Routing
	http.Handle("/product", RecoverHandler(MethodsHandler(PageHandler(ProductHandler), "GET")))
	http.Handle("/product/", RecoverHandler(MethodsHandler(PageHandler(ProductHandler), "GET")))
	http.Handle("/products", RecoverHandler(MethodsHandler(PageHandler(ProductsHandler), "GET")))
	http.Handle("/productModForm", RecoverHandler(MethodsHandler(VandAPageHandler(false, false, true, ValidateProduct, nil, ProductModFormHandler, LandingHandler), "GET")))
	http.Handle("/productMod", RecoverHandler(MethodsHandler(VandAPageHandler(false, false, false, ValidateProduct, RequiredProductMod, ProductModHandler, ProductModFormHandler), "POST")))
	//Post Routing
	http.Handle("/post", RecoverHandler(MethodsHandler(PageHandler(PostHandler), "GET")))
	http.Handle("/post/", RecoverHandler(MethodsHandler(PageHandler(PostHandler), "GET")))
	http.Handle("/posts", RecoverHandler(MethodsHandler(PageHandler(PostsHandler), "GET")))
	http.Handle("/postModForm", RecoverHandler(MethodsHandler(VandAPageHandler(false, false, true, ValidatePost, nil, PostModFormHandler, LandingHandler), "GET")))
	http.Handle("/postMod", RecoverHandler(MethodsHandler(VandAPageHandler(false, false, false, ValidatePost, RequiredPostMod, PostModHandler, PostModFormHandler), "POST")))
	//Static routing
	http.Handle("/images/", MethodsHandler(http.StripPrefix("/images/", http.FileServer(http.Dir("./view/webcontent/www/images"))), "GET"))
	http.Handle("/tmp/", MethodsHandler(http.StripPrefix("/tmp/", http.FileServer(http.Dir("./view/webcontent/www/tmp"))), "GET"))
	http.Handle("/font-awesome/", MethodsHandler(http.StripPrefix("/font-awesome/", http.FileServer(http.Dir("./view/webcontent/www/font-awesome"))), "GET"))
	http.Handle("/css/", MethodsHandler(http.StripPrefix("/css/", http.FileServer(http.Dir("./view/webcontent/www/css"))), "GET"))
	http.Handle("/js/", MethodsHandler(http.StripPrefix("/js/", http.FileServer(http.Dir("./view/webcontent/www/js"))), "GET"))
	http.Handle("/fonts/", MethodsHandler(http.StripPrefix("/fonts/", http.FileServer(http.Dir("./view/webcontent/www/fonts"))), "GET"))
	http.Handle("/slick/", MethodsHandler(http.StripPrefix("/slick/", http.FileServer(http.Dir("./view/webcontent/www/libs/slick"))), "GET"))
	//Memcache Routing
	http.Handle("/mc_delete", RecoverHandler(MethodsHandler(AuthenticatedPageHandler(MCDeleteHandler, false), "GET")))
	http.Handle("/mc_load", RecoverHandler(MethodsHandler(AuthenticatedPageHandler(MCAddHandler, false), "GET")))
	//Database Routing
	http.Handle("/db_tables", RecoverHandler(MethodsHandler(AuthenticatedPageHandler(DBTablesHandler, false), "GET")))
	http.Handle("/db_data", RecoverHandler(MethodsHandler(AuthenticatedPageHandler(DBDataHandler, false), "GET")))
	http.Handle("/db_test", RecoverHandler(MethodsHandler(AuthenticatedPageHandler(DBTestHandler, false), "GET")))
	//Login Routing
	http.Handle("/loginIndex", RecoverHandler(MethodsHandler(PageHandler(loginIndexHandler), "GET")))
	http.Handle("/login", RecoverHandler(MethodsHandler(VandAPageHandler(true, true, true, ValidateLogin, nil, loginHandler, loginIndexHandler), "POST")))
	http.Handle("/logout", RecoverHandler(MethodsHandler(AuthenticatedPageHandler(logoutHandler, true), "GET")))
	http.Handle("/GoogleLogin", RecoverHandler(MethodsHandler(PageHandler(handleGoogleLogin), "GET")))
	http.Handle("/GoogleCallback", RecoverHandler(MethodsHandler(PageHandler(handleGoogleCallback), "GET")))
	http.Handle("/FacebookLogin", RecoverHandler(MethodsHandler(PageHandler(handleFacebookLogin), "GET")))
	http.Handle("/FacebookCallback", RecoverHandler(MethodsHandler(PageHandler(handleFacebookCallback), "GET")))
	//Register Routing
	http.Handle("/registerForm", RecoverHandler(MethodsHandler(PageHandler(registerHandler), "GET")))
	http.Handle("/register", RecoverHandler(MethodsHandler(VandAPageHandler(true, true, true, ValidateRegister, RequiredRegister, registerFormHandler, registerHandler), "POST")))
	http.Handle("/verifyregistration", RecoverHandler(MethodsHandler(VandAPageHandler(true, true, true, ValidateEmailCode, RequiredEmailCode, verifyRegisterHandler, verifyRegisterHandler), "GET")))
	//Forgot Password
	http.Handle("/forgotPasswdForm", RecoverHandler(MethodsHandler(PageHandler(forgotPasswdHandler), "GET")))
	http.Handle("/forgotPasswd", RecoverHandler(MethodsHandler(VandAPageHandler(true, true, true, ValidateForgotPasswd, RequiredForgotPasswd, forgotPasswdFormHandler, forgotPasswdHandler), "POST")))
	http.Handle("/passwdResetForm", RecoverHandler(MethodsHandler(VandAPageHandler(true, true, true, ValidateEmailCode, RequiredEmailCode, resetPasswdHandler, resetPasswdHandler), "GET")))
	http.Handle("/passwdReset", RecoverHandler(MethodsHandler(VandAPageHandler(true, true, true, ValidateResetPasswd, RequiredResetPasswd, resetPasswdFormHandler, resetPasswdHandler), "POST")))
	//Social Routing
	http.Handle("/adminsocialpost", RecoverHandler(MethodsHandler(VandAPageHandler(false, false, true, ValidateSocialPost, nil, SocialPostHandler, nil), "GET")))
	//Image Routing
	http.Handle("/imageModForm", RecoverHandler(MethodsHandler(VandAPageHandler(false, false, true, ValidateImage, nil, ImageModFormHandler, LandingHandler), "GET")))
	http.Handle("/imageMod", RecoverHandler(MethodsHandler(VandAPageHandler(false, false, false, ValidateImage, RequiredImageMod, ImageModHandler, ImageModFormHandler), "POST")))
	http.Handle("/imageUpdate", RecoverHandler(MethodsHandler(VandAPageHandler(false, false, false, ValidateImage, RequiredImageMod, ImageUpdateHandler, ImageModFormHandler), "POST")))

}

//This only loads the page into the page datastruct, there is no authentication
//validation
func PageHandler(next func(http.ResponseWriter, *http.Request, *page)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isReadyState(w, r) {
			page := NewPage(w, r)
			page.State = State
			page.BaseURL = BaseURL
			pSP := bluemonday.StrictPolicy()
			r.ParseForm() // Required if you don't call r.FormValue()
			if len(r.Form["view"]) > 0 && strings.TrimSpace(r.Form["view"][0]) != "" {
				if govalidator.IsPrintableASCII(r.Form["view"][0]) {
					page.View = pSP.Sanitize(r.Form["view"][0])
				}
			}
			next(w, r, page)
			return
		}
	})
}

//This loads the page into the page datastruct and authenticates, there is no
//validation.  It will default to logging the user out if you fail to authenticate
//but that can be over ridden with the ignoreLogout flag
func AuthenticatedPageHandler(pass func(http.ResponseWriter, *http.Request, *page), ignoreLogout bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		page := NewPage(w, r)
		page.State = State
		page.BaseURL = BaseURL
		if page.Authenticated {
			pass(w, r, page)
			return
		} else if ignoreLogout {
			glog.Errorln("ERROR: User not authenticated and requesting restricted content, possible attack!")
			http.Redirect(w, r, "/", 302)
			return
		} else {
			glog.Errorln("ERROR: User not authenticated and requesting restricted content, possible attack!")
			http.Redirect(w, r, "/logout", 302)
			return
		}
	})
}

//This loads the page into the page datastruct, authenticates, and validates.
//You can ignore authentication by setting the ignoreAuth flag.
func VandAPageHandler(ignoreAuth bool, ignoreLogout bool, ignoreCSRF bool, validator func(http.ResponseWriter, *http.Request, *page) bool, require func(http.ResponseWriter, *http.Request, *page) bool, pass func(http.ResponseWriter, *http.Request, *page), fail func(http.ResponseWriter, *http.Request, *page)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		page := NewPage(w, r)
		page.State = State
		page.BaseURL = BaseURL
		glog.Infoln("VandAPageHandler")
		if page.Authenticated || ignoreAuth {
			//was a require fields method passed
			if require != nil {
				//check for required fields
				if !require(w, r, page) {
					//Validate the form input and populate the meta data
					if validator(w, r, page) {
						//validate the CSRF ID to make sure we don't double submit or
						//have an attack
						if !ignoreCSRF && !ValidateCSRF(r, page) {
							http.Redirect(w, r, "/logout", 302)
							return
						} else {
							pass(w, r, page)
							return
						}
					} else {
						//Validation failed
						glog.Infoln("Bad validation!")
						fail(w, r, page)
						return
					}
				} else {
					//check for required failed
					glog.Infoln("Missing required fields!")
					fail(w, r, page)
					return
				}
			} else {
				//Validate the form input and populate the meta data
				if validator(w, r, page) {
					//validate the CSRF ID to make sure we don't double submit or
					//have an attack
					if !ignoreCSRF && !ValidateCSRF(r, page) {
						http.Redirect(w, r, "/logout", 302)
						return
					} else {
						pass(w, r, page)
						return
					}
				} else {
					//Validation failed
					glog.Infoln("Bad validation!")
					fail(w, r, page)
					return
				}
			}
		} else if ignoreLogout {
			glog.Errorln("ERROR: User not authenticated and requesting restricted content, possible attack!")
			http.Redirect(w, r, "/", 302)
			return
		} else {
			glog.Errorln("ERROR: User not authenticated and requesting restricted content, possible attack!")
			http.Redirect(w, r, "/logout", 302)
			return
		}
	})
}

//This handler is designed to return a 404 error after a panic has occured
func MethodsHandler(h http.Handler, methods ...string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isValidMethod := false
		for _, v := range methods {
			if strings.ToUpper(v) == r.Method {
				isValidMethod = true
			}
		}
		if !isValidMethod {
			glog.Errorln("ERROR: Invalid Method used to access content, possible attack!")
			Error404(w, "ERROR: Invalid Method used to access content, possible attack!")
			return
		}
		h.ServeHTTP(w, r) // call next
		return
	})
}

//This handler is designed to return a 404 error after a panic has occured
func RecoverHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// catch all errors and return 404
		defer func() {
			// recover from panic if one occured. Set err to nil otherwise.
			if rec := recover(); rec != nil {
				switch State {
				case Initialize:
					glog.Errorln("We didn't get past the initialization")
					os.Exit(3)
				case Setup:
					RenderSetupTemplate(w, rec)
					return
				case Ready:
					Error404(w, rec)
					return
				case Down:
					glog.Errorln("This instance seems to have gone down.  Shutting it off.")
					os.Exit(3)
				default:
					return
				}
				return
			}
		}()
		h.ServeHTTP(w, r) // call next
		return
	})
}

//This handler is designed to return a 404 error after a panic has occured
func isReadyState(w http.ResponseWriter, r *http.Request) bool {
	switch State {
	case Initialize:
		glog.Errorln("We didn't get past the initialization")
		os.Exit(3)
	case Setup:
		glog.Infoln("In setup mode.  Waiting for all setup requirements to be completed before moving to new state.")
		page := NewSetupPage(w, r)
		page.State = State
		//Check DB
		conn, _ := connectors.GetDB()
		err := conn.Ping()
		if err != nil {
			glog.Errorln("No database connection.  Fix the connection issue and then retry setup.")
			page.RenderSetupTemplate(w, r, "/setup")
			return false
		}
		tables := model.SelectTables()
		found_all_tables := true
		if len(tables) == 0 {
			found_all_tables = false
		}
		for _, table := range tables {
			found := false
			for _, val_table := range Valid_Tables {
				if val_table == table {
					glog.Infoln("Validated table " + table)
					found = true
					break
				}
			}
			if !found {
				ignore_table := false
				for _, ign_table := range Ignore_Tables {
					if ign_table == table {
						glog.Infoln("Ignoring table " + table)
						ignore_table = true
						break
					}
				}
				if !ignore_table {
					glog.Errorln("Missing a table! " + table)
					found_all_tables = false
				}
			}
		}
		if !found_all_tables {
			//try loading the tables
			glog.Infoln("Trying to load the tables")
			DBTablesHandler(w, r, page)
		}
		State = Ready
		return true
	case Ready:
		return true
	case Down:
		glog.Errorln("This instance seems to have gone down.  Shutting it off.")
		os.Exit(3)
	default:
		return false
	}
	return false
}
