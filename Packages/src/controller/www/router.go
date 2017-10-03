// Copyright 2017 Rough Industries LLC. All rights reserved.
//controller/www/router.go: sets up all the routing for the webapp
package www

import (
	"github.com/golang/glog"
	"net/http"
	"strings"
)

//Init to setup the http handlers
func WWWRouterInit() {
	glog.Infoln("Init in www/router.go")
	//Page Routing
	http.Handle("/", RecoverHandler(MethodsHandler(PageHandler(LandingHandler), "GET")))
	//Cocktail Routing
	http.Handle("/cocktail", RecoverHandler(MethodsHandler(PageHandler(CocktailHandler), "GET")))
	http.Handle("/cocktails", RecoverHandler(MethodsHandler(PageHandler(CocktailsHandler), "GET")))
	http.Handle("/cocktailsindex", RecoverHandler(MethodsHandler(PageHandler(CocktailsIndexHandler), "GET")))
	http.Handle("/cocktailsindex/", RecoverHandler(MethodsHandler(PageHandler(CocktailsIndexHandler), "GET")))
	http.Handle("/cocktailsByMetaID", RecoverHandler(MethodsHandler(PageHandler(CocktailsByMetaIDHandler), "GET")))
	http.Handle("/cocktailsByProductID", RecoverHandler(MethodsHandler(PageHandler(CocktailsByProductIDHandler), "GET")))
	http.Handle("/cocktailModForm", RecoverHandler(MethodsHandler(AuthenticatedPageHandler(CocktailModFormHandler, false), "GET")))
	http.Handle("/cocktailMod", RecoverHandler(MethodsHandler(VandAPageHandler(false, false, false, ValidateCocktail, nil, CocktailModHandler, CocktailModFormHandler), "POST")))
	//Meta Routing
	http.Handle("/metaModForm", RecoverHandler(MethodsHandler(VandAPageHandler(false, false, true, ValidateMeta, nil, MetaModFormHandler, LandingHandler), "GET")))
	http.Handle("/metaMod", RecoverHandler(MethodsHandler(VandAPageHandler(false, false, false, ValidateMeta, RequiredMetaMod, MetaModHandler, MetaModFormHandler), "POST")))
	//Products Routing
	http.Handle("/product", RecoverHandler(MethodsHandler(PageHandler(ProductHandler), "GET")))
	http.Handle("/product/", RecoverHandler(MethodsHandler(PageHandler(ProductHandler), "GET")))
	http.Handle("/products", RecoverHandler(MethodsHandler(PageHandler(ProductsHandler), "GET")))
	http.Handle("/productModForm", RecoverHandler(MethodsHandler(AuthenticatedPageHandler(ProductModFormHandler, false), "GET")))
	http.Handle("/productMod", RecoverHandler(MethodsHandler(VandAPageHandler(false, false, false, ValidateMeta, nil, ProductModHandler, ProductModFormHandler), "POST")))
	//Static routing
	http.Handle("/images/", MethodsHandler(http.StripPrefix("/images/", http.FileServer(http.Dir("./view/webcontent/www/images"))), "GET"))
	http.Handle("/font-awesome/", MethodsHandler(http.StripPrefix("/font-awesome/", http.FileServer(http.Dir("./view/webcontent/www/font-awesome"))), "GET"))
	http.Handle("/css/", MethodsHandler(http.StripPrefix("/css/", http.FileServer(http.Dir("./view/webcontent/www/css"))), "GET"))
	http.Handle("/js/", MethodsHandler(http.StripPrefix("/js/", http.FileServer(http.Dir("./view/webcontent/www/js"))), "GET"))
	http.Handle("/fonts/", MethodsHandler(http.StripPrefix("/fonts/", http.FileServer(http.Dir("./view/webcontent/www/fonts"))), "GET"))
	http.Handle("/slick/", MethodsHandler(http.StripPrefix("/slick/", http.FileServer(http.Dir("./view/webcontent/www/libs/slick"))), "GET"))
	http.Handle("/favicon.ico", MethodsHandler(http.StripPrefix("/", http.FileServer(http.Dir("./view/webcontent/www/favicon.ico"))), "GET"))
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

}

//This only loads the page into the page datastruct, there is no authentication
//validation
func PageHandler(next func(http.ResponseWriter, *http.Request, *page)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		page := NewPage(w, r)
		next(w, r, page)
		return
	})
}

//This loads the page into the page datastruct and authenticates, there is no
//validation.  It will default to logging the user out if you fail to authenticate
//but that can be over ridden with the ignoreLogout flag
func AuthenticatedPageHandler(pass func(http.ResponseWriter, *http.Request, *page), ignoreLogout bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		page := NewPage(w, r)
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
func VandAPageHandler(ignoreAuth bool, ignoreLogout bool, ignoreCSRF bool, validator func(http.ResponseWriter, *http.Request, *page) bool, require func(*page) bool, pass func(http.ResponseWriter, *http.Request, *page), fail func(http.ResponseWriter, *http.Request, *page)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		page := NewPage(w, r)
		if page.Authenticated || ignoreAuth {
			//Validate the form input and populate the meta data
			if validator(w, r, page) {
				//validate the CSRF ID to make sure we don't double submit or
				//have an attack
				if !ignoreCSRF && !ValidateCSRF(r, page) {
					http.Redirect(w, r, "/logout", 302)
					return
				} else {
					//was a require fields method passed
					if require != nil {
						//check for required fields
						if !require(page) {
							pass(w, r, page)
						} else {
							//check for required failed
							glog.Infoln("Missing required fields!")
							fail(w, r, page)
							return
						}
					} else {
						pass(w, r, page)
					}
					return
				}
			} else {
				//Validation failed
				glog.Infoln("Bad validation!")
				fail(w, r, page)
				return
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
		glog.Infoln(r.Method)
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
				Error404(w, rec)
				return
			}
		}()
		h.ServeHTTP(w, r) // call next
		return
	})
}
