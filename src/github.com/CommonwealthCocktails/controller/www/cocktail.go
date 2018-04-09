// Copyright 2017 Rough Industries LLC. All rights reserved.
//controller/www/cocktail.go: Functions and handlers for dealing with cocktails.
//TODO: migrate cocktail by meta or product id to a single function that
//is passed a meta or a product id parameter
package www

import (
	"github.com/CommonwealthCocktails/model"
	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	"github.com/microcosm-cc/bluemonday"
	log "github.com/sirupsen/logrus"
	"html"
	"html/template"
	"net/http"
	"net/url"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//Cocktail page handler which displays the standard cocktail page.
func CocktailHandler(w http.ResponseWriter, r *http.Request, page *page) {
	log.Infoln("CocktailHandler")
	var cs model.CocktailSet
	if !ValidateCocktailPath(w, r, page) {
		page.RenderPageTemplate(w, r, "404")
	} else {
		cs = page.Cocktail.SelectCocktailsByID(page.Cocktail.ID, true, page.View)
		page.CocktailSet = cs
		if len(page.CocktailSet.Cocktail.RelatedCocktails) <= 0 {
			var temp []string
			for _, e := range cs.Cocktail.Recipe.RecipeSteps {
				temp = append(temp, strconv.Itoa(e.OriginalIngredient.ID))
			}
			page.Cocktails = append(page.Cocktails, page.CocktailSet.Cocktail.SelectCocktailsByIngredientIDs(temp, page.View)...)
		} else {
			page.Cocktails = page.CocktailSet.Cocktail.RelatedCocktails
		}
		//apply the template page info to the index page
		page.RenderPageTemplate(w, r, "cocktail")
	}
}

//Cocktail page handler which displays the standard cocktail page.
func CocktailOfTheDayHandler(w http.ResponseWriter, r *http.Request, page *page) {
	var cs model.CocktailSet
	cs = page.Cocktail.SelectCocktailsByDayOfYear(time.Now().YearDay(), true, page.View)
	page.CocktailSet = cs
	if cs.Cocktail.ID == 0 {
		page.RenderPageTemplate(w, r, "404")
		return
	}
	var temp []string
	for _, e := range cs.Cocktail.Recipe.RecipeSteps {
		temp = append(temp, strconv.Itoa(e.OriginalIngredient.ID))
	}
	page.Cocktails = append(page.Cocktails, page.CocktailSet.Cocktail.SelectCocktailsByIngredientIDs(temp, page.View)...)
	//apply the template page info to the index page
	page.RenderPageTemplate(w, r, "cocktail")
}

//Cocktails page (i.e. all the cocktails) request handler which
//displays the all the cocktails page.
func CocktailsHandler(w http.ResponseWriter, r *http.Request, page *page) {
	var cs model.CocktailSet
	var c []model.Cocktail
	c = page.Cocktail.SelectAllCocktails(page.View)
	totalC := len(c)
	diff := len(c) - ((page.Pagination.CurrentPage - 1) * 25)
	if diff > 25 {
		c = c[(page.Pagination.CurrentPage-1)*25 : ((page.Pagination.CurrentPage-1)*25)+25]
	} else {
		c = c[(page.Pagination.CurrentPage-1)*25 : ((page.Pagination.CurrentPage-1)*25)+diff]
	}
	cs.ChildCocktails = c
	page.CocktailSet = cs
	PaginationCalculate(page, page.Pagination.CurrentPage, 25, totalC, 3)
	page.SubrouteURL = "cocktails"
	//apply the template page info to the index page
	page.RenderPageTemplate(w, r, "cocktails")
}

//Cocktails Top 100 pagerequest handler which
//displays the top 100 the cocktails page.
func CocktailsTop100Handler(w http.ResponseWriter, r *http.Request, page *page) {
	var cs model.CocktailSet
	var c []model.Cocktail
	log.Infoln("Top100 Request")
	c = page.Cocktail.SelectTop100Cocktails(page.View)
	totalC := len(c)
	diff := len(c) - ((page.Pagination.CurrentPage - 1) * 25)
	if diff > 6 {
		c = c[(page.Pagination.CurrentPage-1)*25 : ((page.Pagination.CurrentPage-1)*25)+25]
	} else {
		c = c[(page.Pagination.CurrentPage-1)*25 : ((page.Pagination.CurrentPage-1)*25)+diff]
	}
	cs.ChildCocktails = c
	cs.Metadata.MetaName = "Top 100"
	page.CocktailSet = cs
	PaginationCalculate(page, page.Pagination.CurrentPage, 25, totalC, 2)
	page.SubrouteURL = "cocktailTop100"
	//apply the template page info to the index page
	page.RenderPageTemplate(w, r, "cocktails")
}

//Cocktail Modification Form page handler which displays the Cocktail Modification
//Form page.
func CocktailModFormHandler(w http.ResponseWriter, r *http.Request, page *page) {
	u, _ := url.Parse(r.URL.String())
	log.Infoln(u)
	r.ParseForm()
	var cba model.CocktailsByAlphaNums
	cba = page.Cocktail.SelectCocktailsByAlphaNums(false, page.View)
	page.CocktailsByAlphaNums = cba
	page.Doze = model.SelectDoze(page.View)
	var mbt model.MetasByTypes
	mbt = page.Meta.SelectMetaByTypes(false, true, false, page.View)
	page.MetasByTypes = mbt
	var ingredients model.ProductsByTypes
	ingredients = page.Product.SelectProductsByTypes(true, false, false, page.View)
	page.Ingredients = ingredients
	var nonIngredients model.ProductsByTypes
	nonIngredients = page.Product.SelectProductsByTypes(false, true, false, page.View)
	page.NonIngredients = nonIngredients
	log.Infoln(r.Form["ID"])
	page.IsForm = true
	if len(r.Form["ID"]) == 0 {
		//apply the template page info to the index page
		page.RenderPageTemplate(w, r, "cocktailmodform")
	} else {
		id, _ := strconv.Atoi(r.Form["ID"][0])
		var in model.Cocktail
		in.ID = id
		out := page.Cocktail.SelectCocktailsByID(id, false, page.View)
		page.Cocktail = out.Cocktail
		page.RenderPageTemplate(w, r, "cocktailmodform")
	}
}

//Cocktail modification form page request handler which process the cocktail
//modification request.  This will after verifying a valid user session,
//modify the cocktail data based on the request.
func CocktailModHandler(w http.ResponseWriter, r *http.Request, page *page) {
	r.ParseForm()
	log.Infoln(r.Form)
	var cba model.CocktailsByAlphaNums
	cba = page.Cocktail.SelectCocktailsByAlphaNums(false, page.View)
	page.CocktailsByAlphaNums = cba
	page.Doze = model.SelectDoze(page.View)
	var mbt model.MetasByTypes
	mbt = page.Meta.SelectMetaByTypes(false, true, false, page.View)
	page.MetasByTypes = mbt
	var ingredients model.ProductsByTypes
	ingredients = page.Product.SelectProductsByTypes(true, false, false, page.View)
	page.Ingredients = ingredients
	var nonIngredients model.ProductsByTypes
	nonIngredients = page.Product.SelectProductsByTypes(false, true, false, page.View)
	page.NonIngredients = nonIngredients
	page.IsForm = true
	if r.Form["button"][0] == "add" {
		ret_id := page.Cocktail.InsertCocktail(page.View)
		model.LoadMCWithCocktailByAlphaNumsData(page.View)
		outCocktail := page.Cocktail.SelectCocktailsByID(ret_id, false, page.View)
		page.Cocktail = outCocktail.Cocktail
		page.Messages["cocktailModifySuccess"] = "Cocktail modified successfully and memcache updated!"
		page.RenderPageTemplate(w, r, "cocktailmodform")
		return
	} else if r.Form["button"][0] == "update" {
		ret_id := page.Cocktail.UpdateCocktail(page.View)
		model.LoadMCWithCocktailByAlphaNumsData(page.View)
		outCocktail := page.Cocktail.SelectCocktailsByID(ret_id, false, page.View)
		page.Cocktail = outCocktail.Cocktail
		page.Messages["cocktailModifySuccess"] = "Cocktail modified successfully and memcache updated!"
		page.RenderPageTemplate(w, r, "cocktailmodform")
		return
	} else {
		page.Messages["cocktailModifyFail"] = "Cocktail modification failed.  You tried to perform an unknown operation!"
		page.RenderPageTemplate(w, r, "cocktailmodform")
		return
	}
}

//Cocktails Index page i.e. page that gets you to cocktails via header links,
//metas, etc
func CocktailsIndexHandler(w http.ResponseWriter, r *http.Request, page *page) {
	var m model.MetasByTypes
	m = page.Meta.SelectMetaByTypes(true, true, false, page.View)
	page.MetasByTypes = m
	//apply the template page info to the index page
	page.RenderPageTemplate(w, r, "cocktailsindex")
}

//Cocktails by meta id page handler that shows all the cocktails that are
//related to the meta id provided
func CocktailsByMetaIDHandler(w http.ResponseWriter, r *http.Request, page *page) {
	r.ParseForm() // Required if you don't call r.FormValue()
	var cs model.CocktailSet
	if !ValidateCocktailPath(w, r, page) {
		page.RenderPageTemplate(w, r, "404")
	} else {
		meta := page.Meta.SelectMeta(page.View)
		cs.Metadata = meta[0]
		//hijack some meta datas for special views
		if cs.Metadata.MetaName == "Top 100" {
			CocktailsTop100Handler(w, r, page)
			return
		}
		var c []model.Cocktail
		c = page.Cocktail.SelectCocktailsByMeta(page.Meta, page.View)
		totalC := len(c)
		diff := len(c) - ((page.Pagination.CurrentPage - 1) * 25)
		if diff > 25 {
			c = c[(page.Pagination.CurrentPage-1)*25 : ((page.Pagination.CurrentPage-1)*25)+25]
		} else {
			c = c[(page.Pagination.CurrentPage-1)*25 : ((page.Pagination.CurrentPage-1)*25)+diff]
		}
		cs.ChildCocktails = c
		page.CocktailSet = cs
		PaginationCalculate(page, page.Pagination.CurrentPage, 25, totalC, 3)
		page.SubrouteURL = "cocktailsByMetaID/" + strconv.Itoa(page.Meta.ID)
		page.RenderPageTemplate(w, r, "cocktails")
	}
}

//Cocktails by product id page handler that shows all the cocktails that are
//related to the product id provided
func CocktailsByProductIDHandler(w http.ResponseWriter, r *http.Request, page *page) {
	var cs model.CocktailSet
	if !ValidateCocktailPath(w, r, page) {
		page.RenderPageTemplate(w, r, "404")
	} else {
		var c []model.Cocktail
		c = page.Cocktail.SelectCocktailsByProduct(page.Product, page.View)
		totalC := len(c)
		diff := len(c) - ((page.Pagination.CurrentPage - 1) * 25)
		if diff > 25 {
			c = c[(page.Pagination.CurrentPage-1)*25 : ((page.Pagination.CurrentPage-1)*25)+25]
		} else {
			c = c[(page.Pagination.CurrentPage-1)*25 : ((page.Pagination.CurrentPage-1)*25)+diff]
		}
		cs.ChildCocktails = c
		page.CocktailSet = cs
		PaginationCalculate(page, page.Pagination.CurrentPage, 25, totalC, 3)
		page.SubrouteURL = "cocktailsByProductID/" + strconv.Itoa(page.Product.ID)
		prod := page.Product.SelectProduct(page.View)
		cs.Product = prod[0]
		page.RenderPageTemplate(w, r, "cocktails")

	}
}

//Cocktails by Ingredient id page handler that shows all the cocktails that are
//related to the Ingredient id provided
func CocktailsByIngredientIDHandler(w http.ResponseWriter, r *http.Request, page *page) {
	var cs model.CocktailSet
	if !ValidateCocktailPath(w, r, page) {
		page.RenderPageTemplate(w, r, "404")
	} else {
		var c []model.Cocktail
		c = page.Cocktail.SelectCocktailsByIngredientID(page.Product, page.View)
		totalC := len(c)
		diff := len(c) - ((page.Pagination.CurrentPage - 1) * 25)
		if diff > 25 {
			c = c[(page.Pagination.CurrentPage-1)*25 : ((page.Pagination.CurrentPage-1)*25)+25]
		} else {
			c = c[(page.Pagination.CurrentPage-1)*25 : ((page.Pagination.CurrentPage-1)*25)+diff]
		}
		cs.ChildCocktails = c
		page.CocktailSet = cs
		PaginationCalculate(page, page.Pagination.CurrentPage, 25, totalC, 3)
		page.SubrouteURL = "cocktailsByIngredientID/" + strconv.Itoa(page.Product.ID)
		prod := page.Product.SelectProduct(page.View)
		cs.Product = prod[0]
		page.RenderPageTemplate(w, r, "cocktails")

	}
}

//Validates the cocktail form request and populates the Cocktail
//struct
func ValidateCocktailPath(w http.ResponseWriter, r *http.Request, page *page) bool {
	page.Cocktail.Errors = make(map[string]string)
	params := mux.Vars(r)
	pUGCP := bluemonday.UGCPolicy()
	pUGCP.AllowElements("img")
	log.Infoln("Cocktail Validate")
	log.Infoln(params["cocktailID"])
	if len(params["cocktailID"]) > 0 {
		if govalidator.IsInt(params["cocktailID"]) {
			page.Cocktail.ID, _ = strconv.Atoi(params["cocktailID"])
		} else {
			page.Cocktail.Errors["CocktailID"] = "Please enter a valid cocktail id. "
		}
	}

	if len(params["metaID"]) > 0 {
		if govalidator.IsInt(params["metaID"]) {
			page.Meta.ID, _ = strconv.Atoi(params["metaID"])
		} else {
			page.Cocktail.Errors["MetaID"] = "Please enter a valid meta id. "
		}
	}

	if len(params["productID"]) > 0 {
		if govalidator.IsInt(params["productID"]) {
			page.Product.ID, _ = strconv.Atoi(params["productID"])
		} else {
			page.Cocktail.Errors["ProductID"] = "Please enter a valid product id. "
		}
	}

	if len(page.Cocktail.Errors) > 0 {
		page.Errors["cocktailErrors"] = "You have errors in your input"
	}
	return len(page.Cocktail.Errors) == 0
}

//Validates the cocktail form request and populates the Cocktail
//struct
func ValidateCocktail(w http.ResponseWriter, r *http.Request, page *page) bool {
	page.Cocktail.Errors = make(map[string]string)
	r.ParseForm() // Required if you don't call r.FormValue()
	pUGCP := bluemonday.UGCPolicy()
	pUGCP.AllowElements("img")
	pSP := bluemonday.StrictPolicy()
	log.Infoln("Cocktail Validate")
	if len(r.Form["cocktailID"]) > 0 && strings.TrimSpace(r.Form["cocktailID"][0]) != "" {
		if govalidator.IsInt(r.Form["cocktailID"][0]) {
			page.Cocktail.ID, _ = strconv.Atoi(r.Form["cocktailID"][0])
		} else {
			page.Cocktail.Errors["CocktailID"] = "Please enter a valid cocktail id. "
		}
	}
	if len(r.Form["cocktailTitle"]) > 0 {
		page.Cocktail.Title = r.Form["cocktailTitle"][0]
	}
	if len(r.Form["cocktailName"]) > 0 && strings.TrimSpace(r.Form["cocktailName"][0]) != "" {
		if govalidator.IsPrintableASCII(r.Form["cocktailName"][0]) {
			page.Cocktail.Name = template.HTML(pSP.Sanitize(html.EscapeString(r.Form["cocktailName"][0])))
		} else {
			page.Cocktail.Errors["CocktailName"] = "Please enter a valid product name. "
		}
	}
	if len(r.Form["cocktailDisplayName"]) > 0 {
		page.Cocktail.DisplayName = r.Form["cocktailDisplayName"][0]
	}
	if len(r.Form["cocktailAlternateNames"]) > 0 {
		//page.Cocktail.AlternateNames = r.Form["cocktailAlternateNames"][0]
	}
	if len(r.Form["cocktailSpokenName"]) > 0 {
		page.Cocktail.SpokenName = r.Form["cocktailSpokenName"][0]
	}
	if len(r.Form["cocktailOrigin"]) > 0 {
		re := regexp.MustCompile(`\r?\n`)
		page.Cocktail.Origin = template.HTML(re.ReplaceAllString(r.Form["cocktailOrigin"][0], " "))
	}
	if len(r.Form["cocktailAKA"]) > 0 {
		//page.Cocktail.AKA = r.Form["cocktailAKA"][0]
	}
	if len(r.Form["cocktailDescription"]) > 0 {
		re := regexp.MustCompile(`\r?\n`)
		page.Cocktail.Description = template.HTML(re.ReplaceAllString(r.Form["cocktailDescription"][0], " "))
	}
	if len(r.Form["cocktailComment"]) > 0 {
		re := regexp.MustCompile(`\r?\n`)
		page.Cocktail.Comment = template.HTML(re.ReplaceAllString(r.Form["cocktailComment"][0], " "))
	}
	if len(r.Form["cocktailFootnotes"]) > 0 {
		re := regexp.MustCompile(`\r?\n`)
		page.Cocktail.Footnotes = template.HTML(re.ReplaceAllString(r.Form["cocktailFootnotes"][0], " "))
	}
	if len(r.Form["cocktailKeywords"]) > 0 {
		page.Cocktail.Keywords = r.Form["cocktailKeywords"][0]
	}
	if len(r.Form["cocktailImage"]) > 0 {
		page.Cocktail.ImagePath, page.Cocktail.Image = filepath.Split(r.Form["cocktailImage"][0])
		page.Cocktail.ImagePath = strings.TrimSuffix(page.Cocktail.ImagePath, "/")
	}
	if len(r.Form["cocktailLabeledImageLink"]) > 0 {
		page.Cocktail.LabeledImageLink = r.Form["cocktailLabeledImageLink"][0]
	}
	if len(r.Form["cocktailImageSourceName"]) > 0 {
		page.Cocktail.ImageSourceName = r.Form["cocktailImageSourceName"][0]
	}
	if len(r.Form["cocktailImageSourceLink"]) > 0 {
		page.Cocktail.ImageSourceLink = r.Form["cocktailImageSourceLink"][0]
	}
	if len(r.Form["cocktailRating"]) > 0 {
		page.Cocktail.Rating, _ = strconv.Atoi(r.Form["cocktailRating"][0])
	}
	if len(r.Form["cocktailSourceName"]) > 0 {
		page.Cocktail.SourceName = r.Form["cocktailSourceName"][0]
	}
	if len(r.Form["cocktailSourceLink"]) > 0 {
		page.Cocktail.SourceLink = r.Form["cocktailSourceLink"][0]
	}
	if len(r.Form["recipeMethod"]) > 0 {
		re := regexp.MustCompile(`\r?\n`)
		page.Cocktail.Recipe.Method = template.HTML(re.ReplaceAllString(r.Form["recipeMethod"][0], " "))
	}
	if len(r.Form["ignoreRecipe"]) > 0 {
		page.Cocktail.IgnoreRecipeUpdate, _ = strconv.ParseBool(r.Form["ignoreRecipe"][0])
	}
	if len(r.Form["recipeID"]) > 0 {
		page.Cocktail.Recipe.ID, _ = strconv.Atoi(r.Form["recipeID"][0])
	}
	if len(r.Form["recipeSteps"]) > 0 {
		stepIndexer := r.Form["recipeSteps"][0]
		// Split on comma.
		result := strings.Split(stepIndexer, ",")

		// Display all elements.
		ord := 0
		for i := range result {
			index, _ := strconv.Atoi(result[i])
			if index != -1 {
				var rs model.RecipeStep
				var prod model.Product
				var doze model.Doze
				if len(r.Form["recipestep["+result[i]+"].Ingredient"]) > 0 {
					prod.ID, _ = strconv.Atoi(r.Form["recipestep["+result[i]+"].Ingredient"][0])
					rs.OriginalIngredient = prod
				}
				rs.OriginalIngredient = prod
				if len(r.Form["recipestep["+result[i]+"].Doze"]) > 0 {
					doze.ID, _ = strconv.Atoi(r.Form["recipestep["+result[i]+"].Doze"][0])
					rs.RecipeDoze = doze
				}
				if len(r.Form["recipestep["+result[i]+"].Quantity"]) > 0 {
					rs.RecipeCardinalFloat, _ = strconv.ParseFloat(r.Form["recipestep["+result[i]+"].Quantity"][0], 64)
					quant, _ := strconv.ParseFloat(r.Form["recipestep["+result[i]+"].Quantity"][0], 64)
					rs.RecipeCardinalString = FloatToVulgar(quant)
				}
				rs.RecipeOrdinal = ord
				if len(r.Form["recipestep["+result[i]+"].AltIngredients"]) > 0 {
					for j := range r.Form["recipestep["+result[i]+"].AltIngredients"] {
						var prod model.Product
						prod.ID, _ = strconv.Atoi(r.Form["recipestep["+result[i]+"].AltIngredients"][j])
						rs.AltIngredient = append(rs.AltIngredient, prod)
					}
				}
				page.Cocktail.Recipe.RecipeSteps = append(page.Cocktail.Recipe.RecipeSteps, rs)
				ord++
			}
		}
	}
	if len(r.Form["ignoreProducts"]) > 0 {
		page.Cocktail.IgnoreProductUpdate, _ = strconv.ParseBool(r.Form["ignoreProducts"][0])
	}
	if len(r.Form["Garnish"]) > 0 {
		for _, id := range r.Form["Garnish"] {
			if id != "" {
				var p model.Product
				p.ID, _ = strconv.Atoi(id)
				page.Cocktail.Garnish = append(page.Cocktail.Garnish, p)
			}
		}
	} else {
		page.Cocktail.Garnish = nil
	}
	if len(r.Form["Drinkware"]) > 0 {
		for _, id := range r.Form["Drinkware"] {
			if id != "" {
				var p model.Product
				p.ID, _ = strconv.Atoi(id)
				page.Cocktail.Drinkware = append(page.Cocktail.Drinkware, p)
			}
		}
	}
	if len(r.Form["Tool"]) > 0 {
		for _, id := range r.Form["Tool"] {
			if id != "" {
				var p model.Product
				p.ID, _ = strconv.Atoi(id)
				page.Cocktail.Tool = append(page.Cocktail.Tool, p)
			}
		}
	}

	if len(r.Form["ignoreMetas"]) > 0 {
		page.Cocktail.IgnoreMetaUpdate, _ = strconv.ParseBool(r.Form["ignoreMetas"][0])
	}
	if len(r.Form["Flavor"]) > 0 {
		for _, id := range r.Form["Flavor"] {
			if id != "" {
				var m model.Meta
				m.ID, _ = strconv.Atoi(id)
				page.Cocktail.Flavor = append(page.Cocktail.Flavor, m)
			}
		}
	}
	if len(r.Form["BaseSpirit"]) > 0 {
		for _, id := range r.Form["BaseSpirit"] {
			if id != "" {
				var m model.Meta
				m.ID, _ = strconv.Atoi(id)
				page.Cocktail.BaseSpirit = append(page.Cocktail.BaseSpirit, m)
			}
		}
	}
	if len(r.Form["Type"]) > 0 {
		for _, id := range r.Form["Type"] {
			if id != "" {
				var m model.Meta
				m.ID, _ = strconv.Atoi(id)
				page.Cocktail.Type = append(page.Cocktail.Type, m)
			}
		}
	}
	if len(r.Form["Family"]) > 0 {
		for _, id := range r.Form["Family"] {
			if id != "" {
				var m model.Meta
				m.ID, _ = strconv.Atoi(id)
				log.Infoln("Family id = " + id)
				page.Cocktail.Family = append(page.Cocktail.Family, m)
			}
		}
	}
	if r.Form["IsFamilyRoot"] != nil {
		page.Cocktail.IsFamilyRoot = true
	}
	if len(r.Form["Served"]) > 0 {
		for _, id := range r.Form["Served"] {
			if id != "" {
				var m model.Meta
				m.ID, _ = strconv.Atoi(id)
				page.Cocktail.Served = append(page.Cocktail.Served, m)
			}
		}
	}
	if len(r.Form["Technique"]) > 0 {
		for _, id := range r.Form["Technique"] {
			if id != "" {
				var m model.Meta
				m.ID, _ = strconv.Atoi(id)
				page.Cocktail.Technique = append(page.Cocktail.Technique, m)
			}
		}
	}
	if len(r.Form["Strength"]) > 0 {
		for _, id := range r.Form["Strength"] {
			if id != "" {
				var m model.Meta
				m.ID, _ = strconv.Atoi(id)
				page.Cocktail.Strength = append(page.Cocktail.Strength, m)
			}
		}
	}
	if len(r.Form["Difficulty"]) > 0 {
		for _, id := range r.Form["Difficulty"] {
			if id != "" {
				var m model.Meta
				m.ID, _ = strconv.Atoi(id)
				page.Cocktail.Difficulty = append(page.Cocktail.Difficulty, m)
			}
		}
	}
	if len(r.Form["TimeofDay"]) > 0 {
		for _, id := range r.Form["TimeofDay"] {
			if id != "" {
				var m model.Meta
				m.ID, _ = strconv.Atoi(id)
				page.Cocktail.TOD = append(page.Cocktail.TOD, m)
			}
		}
	}
	if len(r.Form["Occasion"]) > 0 {
		for _, id := range r.Form["Occasion"] {
			if id != "" {
				var m model.Meta
				m.ID, _ = strconv.Atoi(id)
				page.Cocktail.Occasion = append(page.Cocktail.Occasion, m)
			}
		}
	}
	if len(r.Form["Style"]) > 0 {
		for _, id := range r.Form["Style"] {
			if id != "" {
				var m model.Meta
				m.ID, _ = strconv.Atoi(id)
				page.Cocktail.Style = append(page.Cocktail.Style, m)
			}
		}
	}
	if len(r.Form["Ratio"]) > 0 {
		for _, id := range r.Form["Ratio"] {
			if id != "" {
				var m model.Meta
				m.ID, _ = strconv.Atoi(id)
				page.Cocktail.Ratio = append(page.Cocktail.Ratio, m)
			}
		}
	}
	if len(r.Form["Drink"]) > 0 {
		for _, id := range r.Form["Drink"] {
			if id != "" {
				var m model.Meta
				m.ID, _ = strconv.Atoi(id)
				page.Cocktail.Drink = append(page.Cocktail.Drink, m)
			}
		}
	}
	if len(page.Cocktail.Errors) > 0 {
		page.Errors["cocktailErrors"] = "You have errors in your input"
	}
	return len(page.Cocktail.Errors) == 0
}

//Checks the page meta struct that required fields are filled in.
func RequiredCocktailMod(w http.ResponseWriter, r *http.Request, page *page) bool {
	page.Cocktail.Errors = make(map[string]string)
	r.ParseForm() // Required if you don't call r.FormValue()
	missingRequired := false
	if r.Form["cocktailName"] == nil || len(r.Form["cocktailName"]) == 0 || strings.TrimSpace(r.Form["cocktailName"][0]) == "" {
		page.Cocktail.Errors["CocktailName"] = "Cocktail name is required."
		missingRequired = true
	}
	return missingRequired
}
