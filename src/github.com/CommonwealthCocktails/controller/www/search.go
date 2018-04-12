// Copyright 2017 Rough Industries LLC. All rights reserved.
package www

import (
	"github.com/CommonwealthCocktails/model"
	"github.com/asaskevich/govalidator"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"
)

//Search page handler which displays the standard post page.
func SearchHandler(w http.ResponseWriter, r *http.Request, page *page) {
	var mbt model.MetasByTypes
	mbt = page.Meta.SelectMetaByTypes(false, true, false, page.View)
	page.MetasByTypes = mbt
	//var ingredients model.ProductsByTypes
	//ingredients = page.Product.SelectProductsByTypes(true, false, false)
	//page.Ingredients = ingredients
	//var nonIngredients model.ProductsByTypes
	//nonIngredients = page.Product.SelectProductsByTypes(false, true, false)
	//page.NonIngredients = nonIngredients
	page.Ingredients, page.NonIngredients = page.Product.SelectProductsForSearch(page.View)
	page.RenderPageTemplate(w, r, "search")
}

func SearchFormHandler(w http.ResponseWriter, r *http.Request, page *page) {
	var cs model.CocktailSet
	var c []model.Cocktail
	c = page.Search.SearchForCoctails(page.View)
	cs.ChildCocktails = c
	page.CocktailSet = cs
	//apply the template page info to the index page
	page.RenderPageTemplate(w, r, "cocktails")
}

func QuickSearchFormHandler(w http.ResponseWriter, r *http.Request, page *page) {
	page.Cocktails, page.Posts, page.Products = page.Search.QuickSearch(page.View)
	page.RenderPageTemplate(w, r, "quickSearch")
}

func ValidateSearch(w http.ResponseWriter, r *http.Request, page *page) bool {
	page.Search.Errors = make(map[string]string)
	r.ParseForm() // Required if you don't call r.FormValue()
	log.Infoln(r.Form)

	if len(r.Form["rating_from_input"]) > 0 && strings.TrimSpace(r.Form["rating_from_input"][0]) != "" {
		if govalidator.IsInt(r.Form["rating_from_input"][0]) {
			log.Println("Min Rating " + r.Form["rating_from_input"][0])
			page.Search.RatingMin, _ = strconv.Atoi(r.Form["rating_from_input"][0])
		} else {
			page.Search.Errors["rating_from_input"] = "Please enter a valid min rating."
		}
	}
	if len(r.Form["rating_to_input"]) > 0 && strings.TrimSpace(r.Form["rating_to_input"][0]) != "" {
		if govalidator.IsInt(r.Form["rating_to_input"][0]) {
			log.Println("Max Rating " + r.Form["rating_to_input"][0])
			page.Search.RatingMax, _ = strconv.Atoi(r.Form["rating_to_input"][0])
		} else {
			page.Search.Errors["rating_to_input"] = "Please enter a valid max rating."
		}
	}
	if len(r.Form["keywords"]) > 0 && strings.TrimSpace(r.Form["keywords"][0]) != "" {
		if govalidator.IsPrintableASCII(r.Form["keywords"][0]) {
			log.Println("Keywords " + r.Form["keywords"][0])
			page.Search.Keywords = r.Form["keywords"][0]
		} else {
			page.Search.Errors["keywords"] = "Please enter valid keywords."
		}
	}
	for y, x := range r.Form {
		if x[0] == "exclude" {
			if strings.Contains(y, "_Ingredients") {
				log.Println("exclude ingredient Num " + strings.Split(y, "_")[0])
				exc, _ := strconv.Atoi(strings.Split(y, "_")[0])
				page.Search.Exclude_Ingredients = append(page.Search.Exclude_Ingredients, exc)
			} else if strings.Contains(y, "_NonIngredients") {
				log.Println("exclude non ingredient Num " + strings.Split(y, "_")[0])
				exc, _ := strconv.Atoi(strings.Split(y, "_")[0])
				page.Search.Exclude_NonIngredients = append(page.Search.Exclude_NonIngredients, exc)
			} else if strings.Contains(y, "_Meta") {
				log.Println("exclude meta Num " + strings.Split(y, "_")[0])
				exc, _ := strconv.Atoi(strings.Split(y, "_")[0])
				page.Search.Exclude_Metas = append(page.Search.Exclude_Metas, exc)
			}
		} else if x[0] == "include" {
			if strings.Contains(y, "_Ingredients") {
				log.Println("include ingredient Num " + strings.Split(y, "_")[0])
				inc, _ := strconv.Atoi(strings.Split(y, "_")[0])
				page.Search.Include_Ingredients = append(page.Search.Include_Ingredients, inc)
			} else if strings.Contains(y, "_NonIngredients") {
				log.Println("include non ingredient Num " + strings.Split(y, "_")[0])
				inc, _ := strconv.Atoi(strings.Split(y, "_")[0])
				page.Search.Include_NonIngredients = append(page.Search.Include_NonIngredients, inc)
			} else if strings.Contains(y, "_Meta") {
				log.Println("include meta Num " + strings.Split(y, "_")[0])
				inc, _ := strconv.Atoi(strings.Split(y, "_")[0])
				page.Search.Include_Metas = append(page.Search.Include_Metas, inc)
			}
		}
	}

	if len(page.Search.Errors) > 0 {
		page.Errors["searchErrors"] = "You have errors in your input. "
	}
	log.Infoln(page.Search)
	return len(page.Search.Errors) == 0
}

func RequiredSearch(w http.ResponseWriter, r *http.Request, page *page) bool {
	page.Search.Errors = make(map[string]string)
	r.ParseForm() // Required if you don't call r.FormValue()
	missingRequired := false
	log.Errorln(page.Search)
	return missingRequired
}

func ValidateQuickSearch(w http.ResponseWriter, r *http.Request, page *page) bool {
	page.Search.Errors = make(map[string]string)
	r.ParseForm() // Required if you don't call r.FormValue()
	log.Infoln(r.Form)

	if len(r.Form["keywords"]) > 0 && strings.TrimSpace(r.Form["keywords"][0]) != "" {
		if govalidator.IsPrintableASCII(r.Form["keywords"][0]) {
			log.Println("Keywords " + r.Form["keywords"][0])
			page.Search.Keywords = r.Form["keywords"][0]
		} else {
			page.Search.Errors["keywords"] = "Please enter valid keywords."
		}
	}

	if len(page.Search.Errors) > 0 {
		page.Errors["searchErrors"] = "You have errors in your input. "
	}
	log.Infoln(page.Search)
	return len(page.Search.Errors) == 0
}

func RequiredQuickSearch(w http.ResponseWriter, r *http.Request, page *page) bool {
	page.Search.Errors = make(map[string]string)
	r.ParseForm() // Required if you don't call r.FormValue()
	missingRequired := false
	log.Errorln(page.Search)
	return missingRequired
}
