// Copyright 2017 Rough Industries LLC. All rights reserved.
//controller/www/cocktail.go: Functions and handlers for dealing with cocktails.
//TODO: migrate cocktail by meta or product id to a single function that
//is passed a meta or a product id parameter
package www

import (
	"html/template"
	"log"
	"model"
	"net/http"
	"net/url"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

//Cocktail page handler which displays the standard cocktail page.
func CocktailHandler(w http.ResponseWriter, r *http.Request) {
	// STANDARD HANDLER HEADER START
	// catch all errors and return 404
	defer func() {
		// recover from panic if one occured. Set err to nil otherwise.
		if rec := recover(); rec != nil {
			Error404(w, rec)
		}
	}()
	page := NewPage()
	page.Username, page.Authenticated = GetSession(r)
	// STANDARD HANLDER HEADER END
	var cs model.CocktailSet
	u, err := url.Parse(r.URL.String())
	if err != nil {
		page.RenderPageTemplate(w, "404")
	}
	m, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		page.RenderPageTemplate(w, "404")
	}
	if len(m["ID"]) == 0 {
		page.RenderPageTemplate(w, "404")
	} else {
		//apply the template page info to the index page
		id, _ := strconv.Atoi(m["ID"][0])
		cs = model.GetCocktailByID(id, true)
		page.CocktailSet = cs
		//apply the template page info to the index page
		page.RenderPageTemplate(w, "cocktail")

	}
}

//Cocktails page (i.e. all the cocktails) request handler which
//displays the all the cocktails page.
func CocktailsHandler(w http.ResponseWriter, r *http.Request) {
	// STANDARD HANDLER HEADER START
	// catch all errors and return 404
	defer func() {
		// recover from panic if one occured. Set err to nil otherwise.
		if rec := recover(); rec != nil {
			Error404(w, rec)
		}
	}()
	page := NewPage()
	page.Username, page.Authenticated = GetSession(r)
	// STANDARD HANLDER HEADER END
	var cs model.CocktailSet
	var c []model.Cocktail
	c = model.SelectAllCocktails(false)
	cs.ChildCocktails = c
	page.CocktailSet = cs
	//apply the template page info to the index page
	page.RenderPageTemplate(w, "cocktails")
}

//Cocktail Modification Form page handler which displays the Cocktail Modification
//Form page.
func CocktailModFormHandler(w http.ResponseWriter, r *http.Request) {
	// STANDARD HANDLER HEADER START
	// catch all errors and return 404
	defer func() {
		// recover from panic if one occured. Set err to nil otherwise.
		if rec := recover(); rec != nil {
			Error404(w, rec)
		}
	}()
	page := NewPage()
	page.Username, page.Authenticated = GetSession(r)
	// STANDARD HANLDER HEADER END
	log.Println("In Add Cocktail Form handler")
	if page.Username != "" && page.Authenticated {
		u, err := url.Parse(r.URL.String())
		log.Println(u)
		if err != nil {
			page.RenderPageTemplate(w, "404")
		}
		m, err := url.ParseQuery(u.RawQuery)
		if err != nil {
			page.RenderPageTemplate(w, "404")
		}
		var cba model.CocktailsByAlphaNums
		cba = model.GetCocktailsByAlphaNums(false)
		page.CocktailsByAlphaNums = cba
		page.Doze = model.SelectDoze()
		var mbt model.MetasByTypes
		mbt = model.GetMetaByTypes(false, true, false)
		page.MetasByTypes = mbt
		var ingredients model.ProductsByTypes
		ingredients = model.GetProductsByTypes(true, false, false)
		page.Ingredients = ingredients
		var nonIngredients model.ProductsByTypes
		nonIngredients = model.GetProductsByTypes(false, true, false)
		page.NonIngredients = nonIngredients
		if len(m["ID"]) == 0 {
			//apply the template page info to the index page
			page.RenderPageTemplate(w, "cocktailmodform")
		} else {
			id, _ := strconv.Atoi(m["ID"][0])
			var in model.Cocktail
			in.ID = id
			out := model.SelectCocktailsByID(id, false)
			page.Cocktail = out.Cocktail
			page.RenderPageTemplate(w, "cocktailmodform")
		}
	} else {
		page.RenderPageTemplate(w, "404")
	}
}

//Cocktail modification form page request handler which process the cocktail
//modification request.  This will after verifying a valid user session,
//modify the cocktail data based on the request.
func CocktailModHandler(w http.ResponseWriter, r *http.Request) {
	// STANDARD HANDLER HEADER START
	// catch all errors and return 404
	defer func() {
		// recover from panic if one occured. Set err to nil otherwise.
		if rec := recover(); rec != nil {
			Error404(w, rec)
		}
	}()
	page := NewPage()
	page.Username, page.Authenticated = GetSession(r)
	// STANDARD HANLDER HEADER END

	if page.Username != "" && page.Authenticated {
		u, err := url.Parse(r.URL.String())
		log.Println(u)
		if err != nil {
			page.RenderPageTemplate(w, "404")
			return
		}
		m, err := url.ParseQuery(u.RawQuery)
		if err != nil {
			page.RenderPageTemplate(w, "404")
			return
		}
		log.Println(m)
		var cba model.CocktailsByAlphaNums
		cba = model.GetCocktailsByAlphaNums(false)
		page.CocktailsByAlphaNums = cba
		page.Doze = model.SelectDoze()
		var mbt model.MetasByTypes
		mbt = model.GetMetaByTypes(false, true, false)
		page.MetasByTypes = mbt
		var ingredients model.ProductsByTypes
		ingredients = model.GetProductsByTypes(true, false, false)
		page.Ingredients = ingredients
		var nonIngredients model.ProductsByTypes
		nonIngredients = model.GetProductsByTypes(false, true, false)
		page.NonIngredients = nonIngredients
		if ValidateCocktail(&page.Cocktail, m) {
			if m["button"][0] == "add" {
				ret_id := model.InsertCocktail(page.Cocktail)
				model.LoadMCWithCocktailByAlphaNumsData()
				outCocktail := model.SelectCocktailsByID(ret_id, false)
				page.Cocktail = outCocktail.Cocktail
				page.Messages["cocktailModifySuccess"] = "Cocktail modified successfully and memcache updated!"
				page.RenderPageTemplate(w, "cocktailmodform")
				return
			} else if m["button"][0] == "update" {
				ret_id := model.UpdateCocktail(page.Cocktail)
				model.LoadMCWithCocktailByAlphaNumsData()
				outCocktail := model.SelectCocktailsByID(ret_id, false)
				page.Cocktail = outCocktail.Cocktail
				page.Messages["cocktailModifySuccess"] = "Cocktail modified successfully and memcache updated!"
				page.RenderPageTemplate(w, "cocktailmodform")
				return
			} else {
				page.Messages["cocktailModifyFail"] = "Cocktail modification failed.  You tried to perform an unknown operation!"
				page.RenderPageTemplate(w, "cocktailmodform")
				return
			}
		} else {
			log.Println("Bad cocktail!")
			page.RenderPageTemplate(w, "/cocktailmodform")
			return
		}
	} else {
		page.RenderPageTemplate(w, "404")
		return
	}
}

//Cocktails Index page i.e. page that gets you to cocktails via header links,
//metas, etc
func CocktailsIndexHandler(w http.ResponseWriter, r *http.Request) {
	// STANDARD HANDLER HEADER START
	// catch all errors and return 404
	defer func() {
		// recover from panic if one occured. Set err to nil otherwise.
		if rec := recover(); rec != nil {
			Error404(w, rec)
		}
	}()
	page := NewPage()
	page.Username, page.Authenticated = GetSession(r)
	// STANDARD HANLDER HEADER END
	var m model.MetasByTypes
	m = model.GetMetaByTypes(true, true, false)
	page.MetasByTypes = m
	//apply the template page info to the index page
	page.RenderPageTemplate(w, "cocktailsindex")
}

//Cocktails by meta id page handler that shows all the cocktails that are
//related to the meta id provided
func CocktailsByMetaIDHandler(w http.ResponseWriter, r *http.Request) {
	// STANDARD HANDLER HEADER START
	// catch all errors and return 404
	defer func() {
		// recover from panic if one occured. Set err to nil otherwise.
		if rec := recover(); rec != nil {
			Error404(w, rec)
		}
	}()
	page := NewPage()
	page.Username, page.Authenticated = GetSession(r)
	// STANDARD HANLDER HEADER END
	var cs model.CocktailSet
	u, err := url.Parse(r.URL.String())
	if err != nil {
		page.RenderPageTemplate(w, "404")
	}
	m, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		page.RenderPageTemplate(w, "404")
	}
	if len(m["ID"]) == 0 {
		page.RenderPageTemplate(w, "404")
	} else {
		id, _ := strconv.Atoi(m["ID"][0])
		log.Println("Meta ID: " + m["ID"][0])
		var inMeta model.Meta
		inMeta.ID = id
		var c []model.Cocktail
		c = model.SelectCocktailsByMeta(inMeta)
		cs.ChildCocktails = c
		meta := model.SelectMeta(inMeta)
		cs.Metadata = meta[0]
		page.CocktailSet = cs
		page.RenderPageTemplate(w, "cocktails")
	}
}

//Cocktails by product id page handler that shows all the cocktails that are
//related to the product id provided
func CocktailsByProductIDHandler(w http.ResponseWriter, r *http.Request) {
	// STANDARD HANDLER HEADER START
	// catch all errors and return 404
	defer func() {
		// recover from panic if one occured. Set err to nil otherwise.
		if rec := recover(); rec != nil {
			Error404(w, rec)
		}
	}()
	page := NewPage()
	page.Username, page.Authenticated = GetSession(r)
	// STANDARD HANLDER HEADER END
	var cs model.CocktailSet
	u, err := url.Parse(r.URL.String())
	if err != nil {
		page.RenderPageTemplate(w, "404")
	}
	m, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		page.RenderPageTemplate(w, "404")
	}
	if len(m["ID"]) == 0 {
		page.RenderPageTemplate(w, "404")
	} else {
		id, _ := strconv.Atoi(m["ID"][0])

		var inProduct model.Product
		inProduct.ID = id
		var c []model.Cocktail
		c = model.SelectCocktailsByProduct(inProduct)
		cs.ChildCocktails = c
		prod := model.SelectProduct(inProduct)
		cs.Product = prod[0]
		page.CocktailSet = cs
		page.RenderPageTemplate(w, "cocktails")

	}
}

//Validates the cocktail form request and populates the Cocktail
//struct
func ValidateCocktail(cocktail *model.Cocktail, m map[string][]string) bool {
	cocktail.Errors = make(map[string]string)

	if len(m["cocktailID"]) > 0 {
		cocktail.ID, _ = strconv.Atoi(m["cocktailID"][0])
	}
	if len(m["cocktailTitle"]) > 0 {
		cocktail.Title = m["cocktailTitle"][0]
	}
	if len(m["cocktailName"]) > 0 && strings.TrimSpace(m["cocktailName"][0]) != "" {
		cocktail.Name = m["cocktailName"][0]
	} else {
		cocktail.Errors["CocktailName"] = "Please enter a valid cocktail name"
	}
	if len(m["cocktailDisplayName"]) > 0 {
		cocktail.DisplayName = m["cocktailDisplayName"][0]
	}
	if len(m["cocktailAlternateNames"]) > 0 {
		//cocktail.AlternateNames = m["cocktailAlternateNames"][0]
	}
	if len(m["cocktailSpokenName"]) > 0 {
		cocktail.SpokenName = m["cocktailSpokenName"][0]
	}
	if len(m["cocktailOrigin"]) > 0 {
		re := regexp.MustCompile(`\r?\n`)
		cocktail.Origin = template.HTML(re.ReplaceAllString(m["cocktailOrigin"][0], " "))
	}
	if len(m["cocktailAKA"]) > 0 {
		//cocktail.AKA = m["cocktailAKA"][0]
	}
	if len(m["cocktailDescription"]) > 0 {
		re := regexp.MustCompile(`\r?\n`)
		cocktail.Description = template.HTML(re.ReplaceAllString(m["cocktailDescription"][0], " "))
	}
	if len(m["cocktailComment"]) > 0 {
		re := regexp.MustCompile(`\r?\n`)
		cocktail.Comment = template.HTML(re.ReplaceAllString(m["cocktailComment"][0], " "))
	}
	if len(m["cocktailImage"]) > 0 {
		cocktail.ImagePath, cocktail.Image = filepath.Split(m["cocktailImage"][0])
		cocktail.ImagePath = strings.TrimSuffix(cocktail.ImagePath, "/")
	}
	if len(m["cocktailImageSourceName"]) > 0 {
		cocktail.ImageSourceName = m["cocktailImageSourceName"][0]
	}
	if len(m["cocktailImageSourceLink"]) > 0 {
		cocktail.ImageSourceLink = m["cocktailImageSourceLink"][0]
	}
	if len(m["cocktailRating"]) > 0 {
		cocktail.Rating, _ = strconv.Atoi(m["cocktailRating"][0])
	}
	if len(m["cocktailSourceName"]) > 0 {
		cocktail.SourceName = m["cocktailSourceName"][0]
	}
	if len(m["cocktailSourceLink"]) > 0 {
		cocktail.SourceLink = m["cocktailSourceLink"][0]
	}
	if len(m["recipeMethod"]) > 0 {
		re := regexp.MustCompile(`\r?\n`)
		cocktail.Recipe.Method = template.HTML(re.ReplaceAllString(m["recipeMethod"][0], " "))
	}
	if len(m["recipeID"]) > 0 {
		cocktail.Recipe.ID, _ = strconv.Atoi(m["recipeID"][0])
	}
	if len(m["recipeSteps"]) > 0 {
		stepIndexer := m["recipeSteps"][0]
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
				if len(m["recipestep["+result[i]+"].Ingredient"]) > 0 {
					prod.ID, _ = strconv.Atoi(m["recipestep["+result[i]+"].Ingredient"][0])
					rs.OriginalIngredient = prod
				}
				rs.OriginalIngredient = prod
				if len(m["recipestep["+result[i]+"].Doze"]) > 0 {
					doze.ID, _ = strconv.Atoi(m["recipestep["+result[i]+"].Doze"][0])
					rs.RecipeDoze = doze
				}
				if len(m["recipestep["+result[i]+"].Quantity"]) > 0 {
					rs.RecipeCardinalFloat, _ = strconv.ParseFloat(m["recipestep["+result[i]+"].Quantity"][0], 64)
					quant, _ := strconv.ParseFloat(m["recipestep["+result[i]+"].Quantity"][0], 64)
					rs.RecipeCardinalString = FloatToVulgar(quant)
				}
				rs.RecipeOrdinal = ord
				if len(m["recipestep["+result[i]+"].AltIngredients"]) > 0 {
					for j := range m["recipestep["+result[i]+"].AltIngredients"] {
						var prod model.Product
						prod.ID, _ = strconv.Atoi(m["recipestep["+result[i]+"].AltIngredients"][j])
						rs.AltIngredient = append(rs.AltIngredient, prod)
					}
				}
				cocktail.Recipe.RecipeSteps = append(cocktail.Recipe.RecipeSteps, rs)
				ord++
			}
		}
	}
	if len(m["Garnish"]) > 0 {
		for _, id := range m["Garnish"] {
			if id != "" {
				var p model.Product
				p.ID, _ = strconv.Atoi(id)
				cocktail.Garnish = append(cocktail.Garnish, p)
			}
		}
	} else {
		cocktail.Garnish = nil
	}
	if len(m["Drinkware"]) > 0 {
		for _, id := range m["Drinkware"] {
			if id != "" {
				var p model.Product
				p.ID, _ = strconv.Atoi(id)
				cocktail.Drinkware = append(cocktail.Drinkware, p)
			}
		}
	}
	if len(m["Tool"]) > 0 {
		for _, id := range m["Tool"] {
			if id != "" {
				var p model.Product
				p.ID, _ = strconv.Atoi(id)
				cocktail.Tool = append(cocktail.Tool, p)
			}
		}
	}
	if len(m["Flavor"]) > 0 {
		for _, id := range m["Flavor"] {
			if id != "" {
				var m model.Meta
				m.ID, _ = strconv.Atoi(id)
				cocktail.Flavor = append(cocktail.Flavor, m)
			}
		}
	}
	if len(m["BaseSpirit"]) > 0 {
		for _, id := range m["BaseSpirit"] {
			if id != "" {
				var m model.Meta
				m.ID, _ = strconv.Atoi(id)
				cocktail.BaseSpirit = append(cocktail.BaseSpirit, m)
			}
		}
	}
	if len(m["Type"]) > 0 {
		for _, id := range m["Type"] {
			if id != "" {
				var m model.Meta
				m.ID, _ = strconv.Atoi(id)
				cocktail.Type = append(cocktail.Type, m)
			}
		}
	}
	if len(m["Family"]) > 0 {
		for _, id := range m["Family"] {
			if id != "" {
				var m model.Meta
				m.ID, _ = strconv.Atoi(id)
				log.Println("Family id = " + id)
				cocktail.Family = append(cocktail.Family, m)
			}
		}
	}
	if m["IsFamilyRoot"] != nil {
		cocktail.IsFamilyRoot = true
	}
	if len(m["Served"]) > 0 {
		for _, id := range m["Served"] {
			if id != "" {
				var m model.Meta
				m.ID, _ = strconv.Atoi(id)
				cocktail.Served = append(cocktail.Served, m)
			}
		}
	}
	if len(m["Technique"]) > 0 {
		for _, id := range m["Technique"] {
			if id != "" {
				var m model.Meta
				m.ID, _ = strconv.Atoi(id)
				cocktail.Technique = append(cocktail.Technique, m)
			}
		}
	}
	if len(m["Strength"]) > 0 {
		for _, id := range m["Strength"] {
			if id != "" {
				var m model.Meta
				m.ID, _ = strconv.Atoi(id)
				cocktail.Strength = append(cocktail.Strength, m)
			}
		}
	}
	if len(m["Difficulty"]) > 0 {
		for _, id := range m["Difficulty"] {
			if id != "" {
				var m model.Meta
				m.ID, _ = strconv.Atoi(id)
				cocktail.Difficulty = append(cocktail.Difficulty, m)
			}
		}
	}
	if len(m["TimeofDay"]) > 0 {
		for _, id := range m["TimeofDay"] {
			if id != "" {
				var m model.Meta
				m.ID, _ = strconv.Atoi(id)
				cocktail.TOD = append(cocktail.TOD, m)
			}
		}
	}
	if len(m["Ratio"]) > 0 {
		for _, id := range m["Ratio"] {
			if id != "" {
				var m model.Meta
				m.ID, _ = strconv.Atoi(id)
				cocktail.Ratio = append(cocktail.Ratio, m)
			}
		}
	}
	if len(m["Drink"]) > 0 {
		for _, id := range m["Drink"] {
			if id != "" {
				var m model.Meta
				m.ID, _ = strconv.Atoi(id)
				cocktail.Drink = append(cocktail.Drink, m)
			}
		}
	}
	return len(cocktail.Errors) == 0
}
