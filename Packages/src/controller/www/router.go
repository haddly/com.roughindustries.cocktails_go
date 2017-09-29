// Copyright 2017 Rough Industries LLC. All rights reserved.
//controller/www/router.go: sets up all the routing for the webapp
package www

import (
	"github.com/golang/glog"
	"model"
	"net/http"
)

//Init to setup the http handlers
func WWWRouterInit() {
	glog.Infoln("Init in www/router.go")
	//Page Routing
	http.Handle("/", RecoverHandler(PageHandler(LandingHandler)))
	http.Handle("/load/", RecoverHandler(http.HandlerFunc(Load)))
	//Cocktail Routing
	http.Handle("/cocktail", RecoverHandler(http.HandlerFunc(CocktailHandler)))
	http.Handle("/cocktails", RecoverHandler(http.HandlerFunc(CocktailsHandler)))
	http.Handle("/cocktailsindex", RecoverHandler(http.HandlerFunc(CocktailsIndexHandler)))
	http.Handle("/cocktailsindex/", RecoverHandler(http.HandlerFunc(CocktailsIndexHandler)))
	http.Handle("/cocktailsByMetaID", RecoverHandler(http.HandlerFunc(CocktailsByMetaIDHandler)))
	http.Handle("/cocktailsByProductID", RecoverHandler(http.HandlerFunc(CocktailsByProductIDHandler)))
	http.Handle("/cocktailModForm", RecoverHandler(http.HandlerFunc(CocktailModFormHandler)))
	http.Handle("/cocktailMod", RecoverHandler(http.HandlerFunc(CocktailModHandler)))
	//Meta Routing
	http.Handle("/metaModForm", RecoverHandler(AuthenticatedPageHandler(MetaModFormHandler, LandingHandler)))
	http.Handle("/metaMod", RecoverHandler(ValidatedPageHandler(ValidateMeta, MetaModHandler, MetaModFormHandler)))
	//Products Routing
	http.Handle("/product", RecoverHandler(http.HandlerFunc(ProductHandler)))
	http.Handle("/product/", RecoverHandler(http.HandlerFunc(ProductHandler)))
	http.Handle("/products", RecoverHandler(http.HandlerFunc(ProductsHandler)))
	http.Handle("/productModForm", RecoverHandler(http.HandlerFunc(ProductModFormHandler)))
	http.Handle("/productMod", RecoverHandler(http.HandlerFunc(ProductModHandler)))
	//Static routing
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./view/webcontent/www/images"))))
	http.Handle("/font-awesome/", http.StripPrefix("/font-awesome/", http.FileServer(http.Dir("./view/webcontent/www/font-awesome"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./view/webcontent/www/css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./view/webcontent/www/js"))))
	http.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("./view/webcontent/www/fonts"))))
	http.Handle("/slick/", http.StripPrefix("/slick/", http.FileServer(http.Dir("./view/webcontent/www/libs/slick"))))
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./view/webcontent/www/favicon.ico")
	})
	//Memcache Routing
	http.Handle("/mc_delete", RecoverHandler(http.HandlerFunc(MCDeleteHandler)))
	http.Handle("/mc_load", RecoverHandler(http.HandlerFunc(MCAddHandler)))
	http.HandleFunc("/memcache", func(w http.ResponseWriter, r *http.Request) {
		model.LoadMCWithProductData()
		model.LoadMCWithMetaData()
	})
	//Database Routing
	http.Handle("/db_tables", RecoverHandler(http.HandlerFunc(DBTablesHandler)))
	http.Handle("/db_data", RecoverHandler(http.HandlerFunc(DBDataHandler)))
	http.Handle("/db_test", RecoverHandler(http.HandlerFunc(DBTestHandler)))
	//Login Routing
	http.Handle("/loginIndex", RecoverHandler(http.HandlerFunc(loginIndexHandler)))
	http.Handle("/login", RecoverHandler(http.HandlerFunc(loginHandler)))
	http.Handle("/logout", RecoverHandler(http.HandlerFunc(logoutHandler)))
	http.Handle("/GoogleLogin", RecoverHandler(http.HandlerFunc(handleGoogleLogin)))
	http.Handle("/GoogleCallback", RecoverHandler(http.HandlerFunc(handleGoogleCallback)))
	http.Handle("/FacebookLogin", RecoverHandler(http.HandlerFunc(handleFacebookLogin)))
	http.Handle("/FacebookCallback", RecoverHandler(http.HandlerFunc(handleFacebookCallback)))

}

func PageHandler(next func(http.ResponseWriter, *http.Request, *page)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		page := NewPage(w, r)
		next(w, r, page)
		return
	})
}

func AuthenticatedPageHandler(pass func(http.ResponseWriter, *http.Request, *page), fail func(http.ResponseWriter, *http.Request, *page)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		page := NewPage(w, r)
		if page.Authenticated {
			pass(w, r, page)
			return
		} else {
			glog.Errorln("ERROR: User not authenticated and requesting update on meta data, possible attack!")
			http.Redirect(w, r, "/", 302)
			return
		}
	})
}

func ValidatedPageHandler(validator func(http.ResponseWriter, *http.Request, *page) bool, pass func(http.ResponseWriter, *http.Request, *page), fail func(http.ResponseWriter, *http.Request, *page)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		page := NewPage(w, r)
		if page.Authenticated {
			//Validate the form input and populate the meta data
			if validator(w, r, page) {
				//validate the CSRF ID to make sure we don't double submit or
				//have an attack
				if !ValidateCSRF(r, page) {
					fail(w, r, page)
					return
				}
			} else {
				//Validation failed
				glog.Infoln("Bad validation!")
				fail(w, r, page)
				return
			}
		} else {
			glog.Errorln("ERROR: User not authenticated and requesting update on meta data, possible attack!")
			http.Redirect(w, r, "/", 302)
			return
		}
		pass(w, r, page)
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
