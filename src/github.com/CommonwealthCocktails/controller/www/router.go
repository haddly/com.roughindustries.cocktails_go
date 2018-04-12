// Copyright 2017 Rough Industries LLC. All rights reserved.
//controller/www/router.go: sets up all the routing for the webapp
package www

import (
	"github.com/CommonwealthCocktails/connectors"
	"github.com/CommonwealthCocktails/model"
	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"strconv"
	"strings"
)

//The integer values for the enumeration
type SystemStateConst int

const (
	Initialize = 1 + iota
	Setup
	Ready
	Down
)

//The string values for the enumeration
var SystemStateStrings = [...]string{
	"Initialize",
	"Setup",
	"Ready",
	"Down",
}

// String returns the English name of the SystemStateConst ("Initialize", "Setup", ...).
func (ss SystemStateConst) String() string { return SystemStateStrings[ss-1] }

//The integer values for the enumeration
type SystemModeConst int

const (
	Dev = 1 + iota
	Production
)

//The string values for the enumeration
var SystemModeStrings = [...]string{
	"Dev",
	"Production",
}

// String returns the English name of the SystemModeConst ("Dev", "Production", ...).
func (sm SystemModeConst) String() string { return SystemModeStrings[sm-1] }

var (
	Mode          = Dev
	State         = Initialize
	BaseURL       = ""
	Valid_Tables  = []string{"cc_usermeta", "cc_usermetatype", "altIngredient", "altnames", "cocktail", "cocktailToAKAs", "cocktailToAltNames", "cocktailToMetas", "cocktailToPosts", "cocktailToProducts", "cocktailToRecipe", "derivedProduct", "doze", "groupProduct", "grouptype", "meta", "metatype", "oauth", "product", "producttype", "post", "recipe", "recipeToRecipeSteps", "recipestep", "userroles", "users", "usersessions"}
	Ignore_Tables = []string{"sqlite_sequence"}
)

//Init to setup the http handlers
func WWWRouterInit() {
	rtr := mux.NewRouter()

	log.Infoln("Init in www/router.go")
	log.Infoln(viper.GetString("BaseURL"))
	BaseURL = viper.GetString("BaseURL")
	//Inits
	PageInit()
	//Page Routing
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}", PageHandler(LandingHandler))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/", PageHandler(LandingHandler))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/{locale:(?:|[a-z][a-z]-[A-Z][A-Z])}/", PageHandler(LandingHandler))
	//Cocktail Routing
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/cocktail/{cocktailID}", RecoverHandler(MethodsHandler(PageHandler(CocktailHandler), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/cocktail/{cocktailID}/", RecoverHandler(MethodsHandler(PageHandler(CocktailHandler), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/{locale:(?:|[a-z][a-z]-[A-Z][A-Z])}/cocktail/{cocktailID}", RecoverHandler(MethodsHandler(PageHandler(CocktailHandler), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/{locale:(?:|[a-z][a-z]-[A-Z][A-Z])}/cocktail/{cocktailID}/", RecoverHandler(MethodsHandler(PageHandler(CocktailHandler), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/cocktail/{cocktailID}/{keywords}", RecoverHandler(MethodsHandler(PageHandler(CocktailHandler), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/cocktail/{cocktailID}/{keywords}/", RecoverHandler(MethodsHandler(PageHandler(CocktailHandler), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/{locale:(?:|[a-z][a-z]-[A-Z][A-Z])}/cocktail/{cocktailID}/{keywords}", RecoverHandler(MethodsHandler(PageHandler(CocktailHandler), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/{locale:(?:|[a-z][a-z]-[A-Z][A-Z])}/cocktail/{cocktailID}/{keywords}/", RecoverHandler(MethodsHandler(PageHandler(CocktailHandler), "GET")))
	//http.Handle("/cocktail", RecoverHandler(MethodsHandler(PageHandler(CocktailHandler), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/cocktailOfTheDay", RecoverHandler(MethodsHandler(PageHandler(CocktailOfTheDayHandler), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/{locale:(?:|[a-z][a-z]-[A-Z][A-Z])}/cocktailOfTheDay", RecoverHandler(MethodsHandler(PageHandler(CocktailOfTheDayHandler), "GET")))
	cocktailTop100SubRoute := rtr.PathPrefix("/{view:(?:|[a-zA-Z0-9]+)}/cocktailTop100").Subrouter()
	cocktailTop100SubRoute.Handle("", RecoverHandler(MethodsHandler(PageHandler(CocktailsTop100Handler), "GET")))
	cocktailTop100SubRoute.Handle("/", RecoverHandler(MethodsHandler(PageHandler(CocktailsTop100Handler), "GET")))
	cocktailTop100SubRoute.Handle("/{page:(?:|[0-9]+)}", RecoverHandler(MethodsHandler(PageHandler(CocktailsTop100Handler), "GET")))
	cocktailTop100SubRoute.Handle("/{page:(?:|[0-9]+)}/", RecoverHandler(MethodsHandler(PageHandler(CocktailsTop100Handler), "GET")))
	cocktailTop100LocaleSubRoute := rtr.PathPrefix("/{view:(?:|[a-zA-Z0-9]+)}/{locale:(?:|[a-z][a-z]-[A-Z][A-Z])}/cocktailTop100").Subrouter()
	cocktailTop100LocaleSubRoute.Handle("", RecoverHandler(MethodsHandler(PageHandler(CocktailsTop100Handler), "GET")))
	cocktailTop100LocaleSubRoute.Handle("/", RecoverHandler(MethodsHandler(PageHandler(CocktailsTop100Handler), "GET")))
	cocktailTop100LocaleSubRoute.Handle("/{page:(?:|[0-9]+)}", RecoverHandler(MethodsHandler(PageHandler(CocktailsTop100Handler), "GET")))
	cocktailTop100LocaleSubRoute.Handle("/{page:(?:|[0-9]+)}/", RecoverHandler(MethodsHandler(PageHandler(CocktailsTop100Handler), "GET")))
	cocktailsSubRoute := rtr.PathPrefix("/{view:(?:|[a-zA-Z0-9]+)}/cocktails").Subrouter()
	cocktailsSubRoute.Handle("", RecoverHandler(MethodsHandler(PageHandler(CocktailsHandler), "GET")))
	cocktailsSubRoute.Handle("/", RecoverHandler(MethodsHandler(PageHandler(CocktailsHandler), "GET")))
	cocktailsSubRoute.Handle("/{page:(?:|[0-9]+)}", RecoverHandler(MethodsHandler(PageHandler(CocktailsHandler), "GET")))
	cocktailsSubRoute.Handle("/{page:(?:|[0-9]+)}/", RecoverHandler(MethodsHandler(PageHandler(CocktailsHandler), "GET")))
	cocktailsLocaleSubRoute := rtr.PathPrefix("/{view:(?:|[a-zA-Z0-9]+)}/{locale:(?:|[a-z][a-z]-[A-Z][A-Z])}/cocktails").Subrouter()
	cocktailsLocaleSubRoute.Handle("", RecoverHandler(MethodsHandler(PageHandler(CocktailsHandler), "GET")))
	cocktailsLocaleSubRoute.Handle("/", RecoverHandler(MethodsHandler(PageHandler(CocktailsHandler), "GET")))
	cocktailsLocaleSubRoute.Handle("/{page:(?:|[0-9]+)}", RecoverHandler(MethodsHandler(PageHandler(CocktailsHandler), "GET")))
	cocktailsLocaleSubRoute.Handle("/{page:(?:|[0-9]+)}/", RecoverHandler(MethodsHandler(PageHandler(CocktailsHandler), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/cocktailsindex", RecoverHandler(MethodsHandler(PageHandler(CocktailsIndexHandler), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/cocktailsindex/", RecoverHandler(MethodsHandler(PageHandler(CocktailsIndexHandler), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/{locale:(?:|[a-z][a-z]-[A-Z][A-Z])}/cocktailsindex", RecoverHandler(MethodsHandler(PageHandler(CocktailsIndexHandler), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/{locale:(?:|[a-z][a-z]-[A-Z][A-Z])}/cocktailsindex/", RecoverHandler(MethodsHandler(PageHandler(CocktailsIndexHandler), "GET")))

	cocktailsByMetaSubRoute := rtr.PathPrefix("/{view:(?:|[a-zA-Z0-9]+)}/cocktailsByMetaID/{metaID}").Subrouter()
	cocktailsByMetaSubRoute.Handle("", RecoverHandler(MethodsHandler(PageHandler(CocktailsByMetaIDHandler), "GET")))
	cocktailsByMetaSubRoute.Handle("/", RecoverHandler(MethodsHandler(PageHandler(CocktailsByMetaIDHandler), "GET")))
	cocktailsByMetaSubRoute.Handle("/{page:(?:|[0-9]+)}", RecoverHandler(MethodsHandler(PageHandler(CocktailsByMetaIDHandler), "GET")))
	cocktailsByMetaSubRoute.Handle("/{page:(?:|[0-9]+)}/", RecoverHandler(MethodsHandler(PageHandler(CocktailsByMetaIDHandler), "GET")))

	cocktailsByMetaLocaleSubRoute := rtr.PathPrefix("/{view:(?:|[a-zA-Z0-9]+)}/{locale:(?:|[a-z][a-z]-[A-Z][A-Z])}/cocktailsByMetaID/{metaID}").Subrouter()
	cocktailsByMetaLocaleSubRoute.Handle("", RecoverHandler(MethodsHandler(PageHandler(CocktailsByMetaIDHandler), "GET")))
	cocktailsByMetaLocaleSubRoute.Handle("/", RecoverHandler(MethodsHandler(PageHandler(CocktailsByMetaIDHandler), "GET")))
	cocktailsByMetaLocaleSubRoute.Handle("/{page:(?:|[0-9]+)}", RecoverHandler(MethodsHandler(PageHandler(CocktailsByMetaIDHandler), "GET")))
	cocktailsByMetaLocaleSubRoute.Handle("/{page:(?:|[0-9]+)}/", RecoverHandler(MethodsHandler(PageHandler(CocktailsByMetaIDHandler), "GET")))

	cocktailsByProductSubRoute := rtr.PathPrefix("/{view:(?:|[a-zA-Z0-9]+)}/cocktailsByProductID/{productID}").Subrouter()
	cocktailsByProductSubRoute.Handle("", RecoverHandler(MethodsHandler(PageHandler(CocktailsByProductIDHandler), "GET")))
	cocktailsByProductSubRoute.Handle("/", RecoverHandler(MethodsHandler(PageHandler(CocktailsByProductIDHandler), "GET")))
	cocktailsByProductSubRoute.Handle("/{page:(?:|[0-9]+)}", RecoverHandler(MethodsHandler(PageHandler(CocktailsByProductIDHandler), "GET")))
	cocktailsByProductSubRoute.Handle("/{page:(?:|[0-9]+)}/", RecoverHandler(MethodsHandler(PageHandler(CocktailsByProductIDHandler), "GET")))

	cocktailsByProductLocaleSubRoute := rtr.PathPrefix("/{view:(?:|[a-zA-Z0-9]+)}/{locale:(?:|[a-z][a-z]-[A-Z][A-Z])}/cocktailsByProductID/{productID}").Subrouter()
	cocktailsByProductLocaleSubRoute.Handle("", RecoverHandler(MethodsHandler(PageHandler(CocktailsByProductIDHandler), "GET")))
	cocktailsByProductLocaleSubRoute.Handle("/", RecoverHandler(MethodsHandler(PageHandler(CocktailsByProductIDHandler), "GET")))
	cocktailsByProductLocaleSubRoute.Handle("/{page:(?:|[0-9]+)}", RecoverHandler(MethodsHandler(PageHandler(CocktailsByProductIDHandler), "GET")))
	cocktailsByProductLocaleSubRoute.Handle("/{page:(?:|[0-9]+)}/", RecoverHandler(MethodsHandler(PageHandler(CocktailsByProductIDHandler), "GET")))

	cocktailsByIngredientSubRoute := rtr.PathPrefix("/{view:(?:|[a-zA-Z0-9]+)}/cocktailsByIngredientID/{productID}").Subrouter()
	cocktailsByIngredientSubRoute.Handle("", RecoverHandler(MethodsHandler(PageHandler(CocktailsByIngredientIDHandler), "GET")))
	cocktailsByIngredientSubRoute.Handle("/", RecoverHandler(MethodsHandler(PageHandler(CocktailsByIngredientIDHandler), "GET")))
	cocktailsByIngredientSubRoute.Handle("/{page:(?:|[0-9]+)}", RecoverHandler(MethodsHandler(PageHandler(CocktailsByIngredientIDHandler), "GET")))
	cocktailsByIngredientSubRoute.Handle("/{page:(?:|[0-9]+)}/", RecoverHandler(MethodsHandler(PageHandler(CocktailsByIngredientIDHandler), "GET")))

	cocktailsByIngredientLocaleSubRoute := rtr.PathPrefix("/{view:(?:|[a-zA-Z0-9]+)}/{locale:(?:|[a-z][a-z]-[A-Z][A-Z])}/cocktailsByIngredientID/{productID}").Subrouter()
	cocktailsByIngredientLocaleSubRoute.Handle("", RecoverHandler(MethodsHandler(PageHandler(CocktailsByIngredientIDHandler), "GET")))
	cocktailsByIngredientLocaleSubRoute.Handle("/", RecoverHandler(MethodsHandler(PageHandler(CocktailsByIngredientIDHandler), "GET")))
	cocktailsByIngredientLocaleSubRoute.Handle("/{page:(?:|[0-9]+)}", RecoverHandler(MethodsHandler(PageHandler(CocktailsByIngredientIDHandler), "GET")))
	cocktailsByIngredientLocaleSubRoute.Handle("/{page:(?:|[0-9]+)}/", RecoverHandler(MethodsHandler(PageHandler(CocktailsByIngredientIDHandler), "GET")))

	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/cocktailModForm", RecoverHandler(MethodsHandler(VandAPageHandler(false, false, true, ValidateCocktail, nil, CocktailModFormHandler, LandingHandler), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/cocktailModForm/", RecoverHandler(MethodsHandler(VandAPageHandler(false, false, true, ValidateCocktail, nil, CocktailModFormHandler, LandingHandler), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/cocktailModForm/{cocktailID:(?:|[0-9]+)}", RecoverHandler(MethodsHandler(VandAPageHandler(false, false, true, ValidateCocktail, nil, CocktailModFormHandler, LandingHandler), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/cocktailMod", RecoverHandler(MethodsHandler(VandAPageHandler(false, false, false, ValidateCocktail, RequiredCocktailMod, CocktailModHandler, CocktailModFormHandler), "POST")))
	//Meta Routing
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/metaModForm", RecoverHandler(MethodsHandler(VandAPageHandler(false, false, true, ValidateMeta, nil, MetaModFormHandler, LandingHandler), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/metaModForm/", RecoverHandler(MethodsHandler(VandAPageHandler(false, false, true, ValidateMeta, nil, MetaModFormHandler, LandingHandler), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/metaModForm/{metaID:(?:|[0-9]+)}", RecoverHandler(MethodsHandler(VandAPageHandler(false, false, true, ValidateMeta, nil, MetaModFormHandler, LandingHandler), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/metaMod", RecoverHandler(MethodsHandler(VandAPageHandler(false, false, false, ValidateMeta, RequiredMetaMod, MetaModHandler, MetaModFormHandler), "POST")))
	//Products Routing
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/product/{productID}", RecoverHandler(MethodsHandler(PageHandler(ProductHandler), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/product/{productID}/", RecoverHandler(MethodsHandler(PageHandler(ProductHandler), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/{locale:(?:|[a-z][a-z]-[A-Z][A-Z])}/product/{productID}", RecoverHandler(MethodsHandler(PageHandler(ProductHandler), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/{locale:(?:|[a-z][a-z]-[A-Z][A-Z])}/product/{productID}/", RecoverHandler(MethodsHandler(PageHandler(ProductHandler), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/product/{productID}/{keywords}", RecoverHandler(MethodsHandler(PageHandler(ProductHandler), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/product/{productID}/{keywords}/", RecoverHandler(MethodsHandler(PageHandler(ProductHandler), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/{locale:(?:|[a-z][a-z]-[A-Z][A-Z])}/product/{productID}/{keywords}", RecoverHandler(MethodsHandler(PageHandler(ProductHandler), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/{locale:(?:|[a-z][a-z]-[A-Z][A-Z])}/product/{productID}/{keywords}/", RecoverHandler(MethodsHandler(PageHandler(ProductHandler), "GET")))

	productsSubRoute := rtr.PathPrefix("/{view:(?:|[a-zA-Z0-9]+)}/products").Subrouter()
	productsSubRoute.Handle("", RecoverHandler(MethodsHandler(PageHandler(ProductsHandler), "GET")))
	productsSubRoute.Handle("/", RecoverHandler(MethodsHandler(PageHandler(ProductsHandler), "GET")))
	productsSubRoute.Handle("/{page:(?:|[0-9]+)}", RecoverHandler(MethodsHandler(PageHandler(ProductsHandler), "GET")))
	productsSubRoute.Handle("/{page:(?:|[0-9]+)}/", RecoverHandler(MethodsHandler(PageHandler(ProductsHandler), "GET")))
	productsLocaleSubRoute := rtr.PathPrefix("/{view:(?:|[a-zA-Z0-9]+)}/{locale:(?:|[a-z][a-z]-[A-Z][A-Z])}/products").Subrouter()
	productsLocaleSubRoute.Handle("", RecoverHandler(MethodsHandler(PageHandler(ProductsHandler), "GET")))
	productsLocaleSubRoute.Handle("/", RecoverHandler(MethodsHandler(PageHandler(ProductsHandler), "GET")))
	productsLocaleSubRoute.Handle("/{page:(?:|[0-9]+)}", RecoverHandler(MethodsHandler(PageHandler(ProductsHandler), "GET")))
	productsLocaleSubRoute.Handle("/{page:(?:|[0-9]+)}/", RecoverHandler(MethodsHandler(PageHandler(ProductsHandler), "GET")))

	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/productModForm", RecoverHandler(MethodsHandler(VandAPageHandler(false, false, true, ValidateProduct, nil, ProductModFormHandler, LandingHandler), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/productModForm/", RecoverHandler(MethodsHandler(VandAPageHandler(false, false, true, ValidateProduct, nil, ProductModFormHandler, LandingHandler), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/productModForm/{productID:(?:|[0-9]+)}", RecoverHandler(MethodsHandler(VandAPageHandler(false, false, true, ValidateProduct, nil, ProductModFormHandler, LandingHandler), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/productMod", RecoverHandler(MethodsHandler(VandAPageHandler(false, false, false, ValidateProduct, RequiredProductMod, ProductModHandler, ProductModFormHandler), "POST")))
	//Post Routing
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/post/{postID}", RecoverHandler(MethodsHandler(PageHandler(PostHandler), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/post/{postID}/", RecoverHandler(MethodsHandler(PageHandler(PostHandler), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/{locale:(?:|[a-z][a-z]-[A-Z][A-Z])}/post/{postID}", RecoverHandler(MethodsHandler(PageHandler(PostHandler), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/{locale:(?:|[a-z][a-z]-[A-Z][A-Z])}/post/{postID}/", RecoverHandler(MethodsHandler(PageHandler(PostHandler), "GET")))

	postsSubRoute := rtr.PathPrefix("/{view:(?:|[a-zA-Z0-9]+)}/posts").Subrouter()
	postsSubRoute.Handle("", RecoverHandler(MethodsHandler(PageHandler(PostsHandler), "GET")))
	postsSubRoute.Handle("/", RecoverHandler(MethodsHandler(PageHandler(PostsHandler), "GET")))
	postsSubRoute.Handle("/{page:(?:|[0-9]+)}", RecoverHandler(MethodsHandler(PageHandler(PostsHandler), "GET")))
	postsSubRoute.Handle("/{page:(?:|[0-9]+)}/", RecoverHandler(MethodsHandler(PageHandler(PostsHandler), "GET")))
	postsLocaleSubRoute := rtr.PathPrefix("/{view:(?:|[a-zA-Z0-9]+)}/{locale:(?:|[a-z][a-z]-[A-Z][A-Z])}/posts").Subrouter()
	postsLocaleSubRoute.Handle("", RecoverHandler(MethodsHandler(PageHandler(PostsHandler), "GET")))
	postsLocaleSubRoute.Handle("/", RecoverHandler(MethodsHandler(PageHandler(PostsHandler), "GET")))
	postsLocaleSubRoute.Handle("/{page:(?:|[0-9]+)}", RecoverHandler(MethodsHandler(PageHandler(PostsHandler), "GET")))
	postsLocaleSubRoute.Handle("/{page:(?:|[0-9]+)}/", RecoverHandler(MethodsHandler(PageHandler(PostsHandler), "GET")))

	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/postModForm", RecoverHandler(MethodsHandler(VandAPageHandler(false, false, true, ValidatePost, nil, PostModFormHandler, LandingHandler), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/postModForm/", RecoverHandler(MethodsHandler(VandAPageHandler(false, false, true, ValidatePost, nil, PostModFormHandler, LandingHandler), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/postModForm/{postID:(?:|[0-9]+)}", RecoverHandler(MethodsHandler(VandAPageHandler(false, false, true, ValidatePost, nil, PostModFormHandler, LandingHandler), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/postMod", RecoverHandler(MethodsHandler(VandAPageHandler(false, false, false, ValidatePost, RequiredPostMod, PostModHandler, PostModFormHandler), "POST")))
	//Search Routing
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/search", RecoverHandler(MethodsHandler(PageHandler(SearchHandler), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/{locale:(?:|[a-z][a-z]-[A-Z][A-Z])}/search", RecoverHandler(MethodsHandler(PageHandler(SearchHandler), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/searchForm", RecoverHandler(MethodsHandler(VandAPageHandler(true, false, true, ValidateSearch, RequiredSearch, SearchFormHandler, SearchHandler), "POST")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/{locale:(?:|[a-z][a-z]-[A-Z][A-Z])}/searchForm", RecoverHandler(MethodsHandler(VandAPageHandler(true, false, true, ValidateSearch, RequiredSearch, SearchHandler, SearchFormHandler), "POST")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/quickSearchForm", RecoverHandler(MethodsHandler(VandAPageHandler(true, false, true, ValidateQuickSearch, RequiredQuickSearch, QuickSearchFormHandler, LandingHandler), "POST")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/{locale:(?:|[a-z][a-z]-[A-Z][A-Z])}/quickSearchForm", RecoverHandler(MethodsHandler(VandAPageHandler(true, false, true, ValidateQuickSearch, RequiredQuickSearch, QuickSearchFormHandler, LandingHandler), "POST")))
	//Memcache Routing
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/mc_delete", RecoverHandler(MethodsHandler(AuthenticatedPageHandler(MCDeleteHandler, false), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/mc_load", RecoverHandler(MethodsHandler(AuthenticatedPageHandler(MCAddHandler, false), "GET")))
	//Database Routing
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/db_tables", RecoverHandler(MethodsHandler(AuthenticatedPageHandler(DBTablesHandler, false), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/db_data", RecoverHandler(MethodsHandler(AuthenticatedPageHandler(DBDataHandler, false), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/db_test", RecoverHandler(MethodsHandler(AuthenticatedPageHandler(DBTestHandler, false), "GET")))
	//Login Routing
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/loginIndex", RecoverHandler(MethodsHandler(PageHandler(loginIndexHandler), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/{locale:(?:|[a-z][a-z]-[A-Z][A-Z])}/loginIndex", RecoverHandler(MethodsHandler(PageHandler(loginIndexHandler), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/login", RecoverHandler(MethodsHandler(VandAPageHandler(true, true, true, ValidateLogin, nil, loginHandler, loginIndexHandler), "POST")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/{locale:(?:|[a-z][a-z]-[A-Z][A-Z])}/login", RecoverHandler(MethodsHandler(VandAPageHandler(true, true, true, ValidateLogin, nil, loginHandler, loginIndexHandler), "POST")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/logout", RecoverHandler(MethodsHandler(AuthenticatedPageHandler(logoutHandler, true), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/{locale:(?:|[a-z][a-z]-[A-Z][A-Z])}/logout", RecoverHandler(MethodsHandler(AuthenticatedPageHandler(logoutHandler, true), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/GoogleLogin", RecoverHandler(MethodsHandler(PageHandler(handleGoogleLogin), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/GoogleCallback", RecoverHandler(MethodsHandler(PageHandler(handleGoogleCallback), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/FacebookLogin", RecoverHandler(MethodsHandler(PageHandler(handleFacebookLogin), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/FacebookCallback", RecoverHandler(MethodsHandler(PageHandler(handleFacebookCallback), "GET")))
	//Register Routing
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/registerForm", RecoverHandler(MethodsHandler(PageHandler(registerHandler), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/{locale:(?:|[a-z][a-z]-[A-Z][A-Z])}/registerForm", RecoverHandler(MethodsHandler(PageHandler(registerHandler), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/register", RecoverHandler(MethodsHandler(VandAPageHandler(true, true, true, ValidateRegister, RequiredRegister, registerFormHandler, registerHandler), "POST")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/{locale:(?:|[a-z][a-z]-[A-Z][A-Z])}/register", RecoverHandler(MethodsHandler(VandAPageHandler(true, true, true, ValidateRegister, RequiredRegister, registerFormHandler, registerHandler), "POST")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/verifyregistration", RecoverHandler(MethodsHandler(VandAPageHandler(true, true, true, ValidateEmailCode, RequiredEmailCode, verifyRegisterHandler, verifyRegisterHandler), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/{locale:(?:|[a-z][a-z]-[A-Z][A-Z])}/verifyregistration", RecoverHandler(MethodsHandler(VandAPageHandler(true, true, true, ValidateEmailCode, RequiredEmailCode, verifyRegisterHandler, verifyRegisterHandler), "GET")))
	//Forgot Password
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/forgotPasswdForm", RecoverHandler(MethodsHandler(PageHandler(forgotPasswdHandler), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/{locale:(?:|[a-z][a-z]-[A-Z][A-Z])}/forgotPasswdForm", RecoverHandler(MethodsHandler(PageHandler(forgotPasswdHandler), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/forgotPasswd", RecoverHandler(MethodsHandler(VandAPageHandler(true, true, true, ValidateForgotPasswd, RequiredForgotPasswd, forgotPasswdFormHandler, forgotPasswdHandler), "POST")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/passwdResetForm", RecoverHandler(MethodsHandler(VandAPageHandler(true, true, true, ValidateEmailCode, RequiredEmailCode, resetPasswdHandler, resetPasswdHandler), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/passwdReset", RecoverHandler(MethodsHandler(VandAPageHandler(true, true, true, ValidateResetPasswd, RequiredResetPasswd, resetPasswdFormHandler, resetPasswdHandler), "POST")))
	//Social Routing
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/adminsocialpost", RecoverHandler(MethodsHandler(VandAPageHandler(false, false, true, ValidateSocialPost, nil, SocialPostHandler, nil), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/{locale:(?:|[a-z][a-z]-[A-Z][A-Z])}/adminsocialpost", RecoverHandler(MethodsHandler(VandAPageHandler(false, false, true, ValidateSocialPost, nil, SocialPostHandler, nil), "GET")))
	//Image Routing
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/imageModForm", RecoverHandler(MethodsHandler(VandAPageHandler(false, false, true, ValidateImage, nil, ImageModFormHandler, LandingHandler), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/imageMod", RecoverHandler(MethodsHandler(VandAPageHandler(false, false, false, ValidateImage, RequiredImageMod, ImageModHandler, ImageModFormHandler), "POST")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/imageUpdate", RecoverHandler(MethodsHandler(VandAPageHandler(false, false, false, ValidateImage, RequiredImageMod, ImageUpdateHandler, ImageModFormHandler), "POST")))
	//Static routing
	rtr.PathPrefix("/{view:(?:|[a-zA-Z0-9]+)}/images/").Handler(MethodsHandler(StaticFileHandler("/images"), "GET"))
	rtr.Handle("/tmp/", MethodsHandler(http.StripPrefix("/tmp/", http.FileServer(http.Dir("./view/webcontent/www/tmp"))), "GET"))
	rtr.PathPrefix("/{view:(?:|[a-zA-Z0-9]+)}/font-awesome/").Handler(MethodsHandler(StaticFileHandler("/font-awesome"), "GET"))
	rtr.PathPrefix("/{view:(?:|[a-zA-Z0-9]+)}/css/").Handler(MethodsHandler(StaticFileHandler("/css"), "GET"))
	rtr.PathPrefix("/{view:(?:|[a-zA-Z0-9]+)}/js/").Handler(MethodsHandler(StaticFileHandler("/js"), "GET"))
	rtr.PathPrefix("/{view:(?:|[a-zA-Z0-9]+)}/fonts/").Handler(MethodsHandler(StaticFileHandler("/fonts"), "GET"))
	rtr.PathPrefix("/{view:(?:|[a-zA-Z0-9]+)}/slick/").Handler(MethodsHandler(StaticFileHandler("/slick"), "GET"))

	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/static/{file:(?:|[a-zA-Z0-9]+)}", RecoverHandler(MethodsHandler(PageHandler(PassthroughFileHandler), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/static/{file:(?:|[a-zA-Z0-9]+)}/", RecoverHandler(MethodsHandler(PageHandler(PassthroughFileHandler), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/{locale:(?:|[a-z][a-z]-[A-Z][A-Z])}/static/{file:(?:|[a-zA-Z0-9]+)}", RecoverHandler(MethodsHandler(PageHandler(PassthroughFileHandler), "GET")))
	rtr.Handle("/{view:(?:|[a-zA-Z0-9]+)}/{locale:(?:|[a-z][a-z]-[A-Z][A-Z])}/static/{file:(?:|[a-zA-Z0-9]+)}/", RecoverHandler(MethodsHandler(PageHandler(PassthroughFileHandler), "GET")))

	//Set the 404 page not found handler
	rtr.NotFoundHandler = http.HandlerFunc(notFound)
	//Apply the router
	http.Handle("/", rtr)
}

//This handler is designed to return a 404 error
func notFound(w http.ResponseWriter, r *http.Request) {
	log.Infoln("notFound Req: " + r.Host + " " + r.URL.Path)
	if r.URL.Path == "/robot.txt" {
		Error404(w, "You crazy robot!", "www")
	} else {
		Error404(w, "ERROR: Page not found!", "www")
	}
}

func StaticFileHandler(subfolder string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		view := ""
		locale := ""
		params := mux.Vars(r)
		if govalidator.IsASCII(params["view"]) {
			view, _ = params["view"]
		}
		if len(view) == 0 {
			view = "www"
		}
		if govalidator.IsASCII(params["locale"]) {
			locale, _ = params["locale"]
		}
		if len(locale) == 0 {
			http.StripPrefix("/"+view+subfolder, http.FileServer(http.Dir("./view/webcontent/"+view+subfolder))).ServeHTTP(w, r)
		} else {
			http.StripPrefix("/"+view+"/"+locale+subfolder, http.FileServer(http.Dir("./view/webcontent/"+view+subfolder))).ServeHTTP(w, r)
		}
		return
	})
}

func PassthroughFileHandler(w http.ResponseWriter, r *http.Request, page *page) {
	view := ""
	locale := ""
	file := ""
	params := mux.Vars(r)
	if govalidator.IsASCII(params["view"]) {
		view, _ = params["view"]
	}
	if len(view) == 0 {
		view = "www"
	}
	if govalidator.IsASCII(params["file"]) {
		file, _ = params["file"]
	}
	if len(file) == 0 {
		file = "about"
	}
	if govalidator.IsASCII(params["locale"]) {
		locale, _ = params["locale"]
	}
	page.View = view
	page.Locale = locale
	page.RenderPageTemplate(w, r, file)
}

//This only loads the page into the page datastruct, there is no authentication
//validation
func PageHandler(next func(http.ResponseWriter, *http.Request, *page)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Infoln("PageHandler")
		r.ParseForm() // Required if you don't call r.FormValue()
		params := mux.Vars(r)
		view := "www"
		if govalidator.IsASCII(params["view"]) && len(params["view"]) != 0 {
			view, _ = params["view"]
		}
		currentPage := 1
		if govalidator.IsInt(params["page"]) && len(params["page"]) != 0 {
			currentPage, _ = strconv.Atoi(params["page"])
		}
		if isReadyState(w, r, view) {
			page := NewPage(w, r, view)
			page.Pagination.CurrentPage = currentPage
			page.State = State
			page.BaseURL = BaseURL
			if govalidator.IsASCII(params["locale"]) {
				page.Locale, _ = params["locale"]
			}
			page.View = view
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
		r.ParseForm() // Required if you don't call r.FormValue()
		params := mux.Vars(r)
		view := "www"
		if govalidator.IsASCII(params["view"]) && len(params["view"]) != 0 {
			view, _ = params["view"]
		}
		page := NewPage(w, r, view)
		page.State = State
		page.BaseURL = BaseURL
		page.View = view
		if govalidator.IsASCII(params["locale"]) {
			page.Locale, _ = params["locale"]
		}
		if page.Authenticated {
			pass(w, r, page)
			return
		} else if ignoreLogout {
			log.Errorln("ERROR: User not authenticated and requesting restricted content, possible attack!")
			http.Redirect(w, r, "/", 302)
			return
		} else {
			log.Errorln("ERROR: User not authenticated and requesting restricted content, possible attack!")
			http.Redirect(w, r, "/logout", 302)
			return
		}
	})
}

//This loads the page into the page datastruct, authenticates, and validates.
//You can ignore authentication by setting the ignoreAuth flag.
func VandAPageHandler(ignoreAuth bool, ignoreLogout bool, ignoreCSRF bool, validator func(http.ResponseWriter, *http.Request, *page) bool, require func(http.ResponseWriter, *http.Request, *page) bool, pass func(http.ResponseWriter, *http.Request, *page), fail func(http.ResponseWriter, *http.Request, *page)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm() // Required if you don't call r.FormValue()
		params := mux.Vars(r)
		view := "www"
		if govalidator.IsASCII(params["view"]) && len(params["view"]) != 0 {
			view, _ = params["view"]
		}
		page := NewPage(w, r, view)
		page.State = State
		page.BaseURL = BaseURL
		page.View = view
		if govalidator.IsASCII(params["locale"]) {
			page.Locale, _ = params["locale"]
		}
		log.Infoln("VandAPageHandler")
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
						log.Infoln("Bad validation!")
						fail(w, r, page)
						return
					}
				} else {
					//check for required failed
					log.Infoln("Missing required fields!")
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
					log.Infoln("Bad validation!")
					fail(w, r, page)
					return
				}
			}
		} else if ignoreLogout {
			log.Errorln("ERROR: User not authenticated and requesting restricted content, possible attack!")
			http.Redirect(w, r, "/", 302)
			return
		} else {
			log.Errorln("ERROR: User not authenticated and requesting restricted content, possible attack!")
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
			log.Errorln("ERROR: Invalid Method used to access content, possible attack!")
			Error404(w, "ERROR: Invalid Method used to access content, possible attack!", "www")
			return
		}
		h.ServeHTTP(w, r) // call next
		return
	})
}

//This handler is designed to return a 404 error after a panic has occured
func RecoverHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Infoln("RecoverHandler")
		// catch all errors and return 404
		defer func() {
			// recover from panic if one occured. Set err to nil otherwise.
			if rec := recover(); rec != nil {
				switch State {
				case Initialize:
					log.Errorln("We didn't get past the initialization")
					os.Exit(3)
				case Setup:
					RenderSetupTemplate(w, rec)
					return
				case Ready:
					Error404(w, rec, "www")
					return
				case Down:
					log.Errorln("This instance seems to have gone down.  Shutting it off.")
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
func isReadyState(w http.ResponseWriter, r *http.Request, site string) bool {
	switch State {
	case Initialize:
		log.Errorln("We didn't get past the initialization")
		os.Exit(3)
	case Setup:
		log.Infoln("In setup mode.  Waiting for all setup requirements to be completed before moving to new state.")
		page := NewSetupPage(w, r)
		page.State = State
		page.View = site
		//Check DB
		conn, err := connectors.GetDBFromMap(site)
		if err != nil {
			log.Errorln("No database connection.  Fix the connection issue and then retry setup.")
			page.RenderSetupTemplate(w, r, "/setup")
			return false
		}
		log.Infoln(conn)
		err = conn.Ping()
		if err != nil {
			log.Errorln("No database connection.  Fix the connection issue and then retry setup.")
			page.RenderSetupTemplate(w, r, "/setup")
			return false
		}
		tables := model.SelectTables(site)
		found_all_tables := true
		if len(tables) == 0 {
			found_all_tables = false
		}
		for _, table := range tables {
			found := false
			for _, val_table := range Valid_Tables {
				if val_table == table {
					log.Infoln("Validated table " + table)
					found = true
					break
				}
			}
			if !found {
				ignore_table := false
				for _, ign_table := range Ignore_Tables {
					if ign_table == table {
						log.Infoln("Ignoring table " + table)
						ignore_table = true
						break
					}
				}
				if !ignore_table {
					log.Errorln("Missing a table! " + table)
					found_all_tables = false
				}
			}
		}
		if !found_all_tables {
			//try loading the tables
			log.Infoln("Trying to load the tables")
			DBTablesHandler(w, r, page)
		}
		State = Ready
		return true
	case Ready:
		return true
	case Down:
		log.Errorln("This instance seems to have gone down.  Shutting it off.")
		os.Exit(3)
	default:
		return false
	}
	return false
}
