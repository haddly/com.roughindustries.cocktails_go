// Copyright 2017 Rough Industries LLC. All rights reserved.
//controller/www/router.go: sets up all the routing for the webapp
package www

import (
	"log"
	"model"
	"net/http"
)

//Init to setup the http handlers
func WWWRouterInit() {
	log.Println("Init in www/router.go")
	//Page Routing
	http.HandleFunc("/", LandingHandler)
	http.HandleFunc("/load/", Load)
	//Cocktail Routing
	http.HandleFunc("/cocktail", CocktailHandler)
	http.HandleFunc("/cocktails", CocktailsHandler)
	http.HandleFunc("/cocktailsindex", CocktailsIndexHandler)
	http.HandleFunc("/cocktailsindex/", CocktailsIndexHandler)
	http.HandleFunc("/cocktailsByMetaID", CocktailsByMetaIDHandler)
	http.HandleFunc("/cocktailsByProductID", CocktailsByProductIDHandler)
	http.HandleFunc("/cocktailModForm", CocktailModFormHandler)
	http.HandleFunc("/cocktailMod", CocktailModHandler)
	//Meta Routing
	http.HandleFunc("/metaModForm", MetaModFormHandler)
	http.HandleFunc("/metaMod", MetaModHandler)
	//Products Routing
	http.HandleFunc("/product", ProductHandler)
	http.HandleFunc("/product/", ProductHandler)
	http.HandleFunc("/products", ProductsHandler)
	http.HandleFunc("/productModForm", ProductModFormHandler)
	http.HandleFunc("/productMod", ProductModHandler)
	//Search Routing
	http.HandleFunc("/search", CocktailSearchHandler)
	//Static routing
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./view/webcontent/www/images"))))
	http.Handle("/font-awesome/", http.StripPrefix("/font-awesome/", http.FileServer(http.Dir("./view/webcontent/www/font-awesome"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./view/webcontent/www/css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./view/webcontent/www/js"))))
	http.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("./view/webcontent/www/fonts"))))
	http.Handle("/slick/", http.StripPrefix("/slick/", http.FileServer(http.Dir("./view/webcontent/www/libs/slick"))))
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "view/webcontent/www/favicon.ico")
	})
	//Memcache Routing
	http.HandleFunc("/mc_delete", MCDeleteHandler)
	http.HandleFunc("/mc_load", MCAddHandler)
	http.HandleFunc("/memcache", func(w http.ResponseWriter, r *http.Request) {
		model.LoadMCWithProductData()
		model.LoadMCWithMetaData()
	})
	//Database Routing
	http.HandleFunc("/db_tables", DBTablesHandler)
	http.HandleFunc("/db_data", DBDataHandler)
	http.HandleFunc("/db_test", DBTestHandler)
	//Login Routing
	http.HandleFunc("/loginIndex", loginIndexHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/GoogleLogin", handleGoogleLogin)
	http.HandleFunc("/GoogleCallback", handleGoogleCallback)
	http.HandleFunc("/FacebookLogin", handleFacebookLogin)
	http.HandleFunc("/FacebookCallback", handleFacebookCallback)

}
