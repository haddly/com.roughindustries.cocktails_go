// Copyright 2017 Rough Industries LLC. All rights reserved.
//controller/www/cocktail.go: Functions and handlers for dealing with cocktails.
//TODO: migrate cocktail by meta or product id to a single function that
//is passed a meta or a product id parameter
package www

import (
	"github.com/golang/glog"
	"html/template"
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
	page := NewPage(w, r)
	var cs model.CocktailSet
	if !ValidateCocktail(&page.Cocktail, r) {
		page.RenderPageTemplate(w, r, "404")
	} else {
		cs = page.Cocktail.SelectCocktailsByID(page.Cocktail.ID, true)
		page.CocktailSet = cs
		//apply the template page info to the index page
		page.RenderPageTemplate(w, r, "cocktail")
	}
}

//Cocktails page (i.e. all the cocktails) request handler which
//displays the all the cocktails page.
func CocktailsHandler(w http.ResponseWriter, r *http.Request) {
	page := NewPage(w, r)
	var cs model.CocktailSet
	var c []model.Cocktail
	c = page.Cocktail.SelectAllCocktails(false)
	cs.ChildCocktails = c
	page.CocktailSet = cs
	//apply the template page info to the index page
	page.RenderPageTemplate(w, r, "cocktails")
}

//Cocktail Modification Form page handler which displays the Cocktail Modification
//Form page.
func CocktailModFormHandler(w http.ResponseWriter, r *http.Request) {
	page := NewPage(w, r)
	glog.Infoln("In Add Cocktail Form handler")
	if page.Authenticated {
		u, _ := url.Parse(r.URL.String())
		glog.Infoln(u)
		var cba model.CocktailsByAlphaNums
		cba = page.Cocktail.SelectCocktailsByAlphaNums(false)
		page.CocktailsByAlphaNums = cba
		page.Doze = model.SelectDoze()
		var mbt model.MetasByTypes
		mbt = page.Meta.SelectMetaByTypes(false, true, false)
		page.MetasByTypes = mbt
		var ingredients model.ProductsByTypes
		ingredients = page.Product.SelectProductsByTypes(true, false, false)
		page.Ingredients = ingredients
		var nonIngredients model.ProductsByTypes
		nonIngredients = page.Product.SelectProductsByTypes(false, true, false)
		page.NonIngredients = nonIngredients
		if len(r.Form["ID"]) == 0 {
			//apply the template page info to the index page
			page.RenderPageTemplate(w, r, "cocktailmodform")
		} else {
			id, _ := strconv.Atoi(r.Form["ID"][0])
			var in model.Cocktail
			in.ID = id
			out := page.Cocktail.SelectCocktailsByID(id, false)
			page.Cocktail = out.Cocktail
			page.RenderPageTemplate(w, r, "cocktailmodform")
		}
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

//Cocktail modification form page request handler which process the cocktail
//modification request.  This will after verifying a valid user session,
//modify the cocktail data based on the request.
func CocktailModHandler(w http.ResponseWriter, r *http.Request) {
	page := NewPage(w, r)
	if page.Authenticated {
		glog.Infoln(r.Form)
		var cba model.CocktailsByAlphaNums
		cba = page.Cocktail.SelectCocktailsByAlphaNums(false)
		page.CocktailsByAlphaNums = cba
		page.Doze = model.SelectDoze()
		var mbt model.MetasByTypes
		mbt = page.Meta.SelectMetaByTypes(false, true, false)
		page.MetasByTypes = mbt
		var ingredients model.ProductsByTypes
		ingredients = page.Product.SelectProductsByTypes(true, false, false)
		page.Ingredients = ingredients
		var nonIngredients model.ProductsByTypes
		nonIngredients = page.Product.SelectProductsByTypes(false, true, false)
		page.NonIngredients = nonIngredients
		if ValidateCocktail(&page.Cocktail, r) {
			if r.Form["button"][0] == "add" {
				ret_id := page.Cocktail.InsertCocktail()
				model.LoadMCWithCocktailByAlphaNumsData()
				outCocktail := page.Cocktail.SelectCocktailsByID(ret_id, false)
				page.Cocktail = outCocktail.Cocktail
				page.Messages["cocktailModifySuccess"] = "Cocktail modified successfully and memcache updated!"
				page.RenderPageTemplate(w, r, "cocktailmodform")
				return
			} else if r.Form["button"][0] == "update" {
				ret_id := page.Cocktail.UpdateCocktail()
				model.LoadMCWithCocktailByAlphaNumsData()
				outCocktail := page.Cocktail.SelectCocktailsByID(ret_id, false)
				page.Cocktail = outCocktail.Cocktail
				page.Messages["cocktailModifySuccess"] = "Cocktail modified successfully and memcache updated!"
				page.RenderPageTemplate(w, r, "cocktailmodform")
				return
			} else {
				page.Messages["cocktailModifyFail"] = "Cocktail modification failed.  You tried to perform an unknown operation!"
				page.RenderPageTemplate(w, r, "cocktailmodform")
				return
			}
		} else {
			glog.Infoln("Bad cocktail!")
			page.RenderPageTemplate(w, r, "/cocktailmodform")
			return
		}
	} else {
		http.Redirect(w, r, "/", 302)
		return
	}
}

//Cocktails Index page i.e. page that gets you to cocktails via header links,
//metas, etc
func CocktailsIndexHandler(w http.ResponseWriter, r *http.Request) {
	page := NewPage(w, r)
	var m model.MetasByTypes
	m = page.Meta.SelectMetaByTypes(true, true, false)
	page.MetasByTypes = m
	//apply the template page info to the index page
	page.RenderPageTemplate(w, r, "cocktailsindex")
}

//Cocktails by meta id page handler that shows all the cocktails that are
//related to the meta id provided
func CocktailsByMetaIDHandler(w http.ResponseWriter, r *http.Request) {
	page := NewPage(w, r)
	r.ParseForm() // Required if you don't call r.FormValue()
	var cs model.CocktailSet
	if len(r.Form["ID"]) == 0 {
		page.RenderPageTemplate(w, r, "404")
	} else {
		id, _ := strconv.Atoi(r.Form["ID"][0])
		glog.Infoln("Meta ID: " + r.Form["ID"][0])
		var inMeta model.Meta
		inMeta.ID = id
		var c []model.Cocktail
		c = page.Cocktail.SelectCocktailsByMeta(inMeta)
		cs.ChildCocktails = c
		meta := inMeta.SelectMeta()
		cs.Metadata = meta[0]
		page.CocktailSet = cs
		page.RenderPageTemplate(w, r, "cocktails")
	}
}

//Cocktails by product id page handler that shows all the cocktails that are
//related to the product id provided
func CocktailsByProductIDHandler(w http.ResponseWriter, r *http.Request) {
	page := NewPage(w, r)
	var cs model.CocktailSet
	if len(r.Form["ID"]) == 0 {
		page.RenderPageTemplate(w, r, "404")
	} else {
		id, _ := strconv.Atoi(r.Form["ID"][0])

		var inProduct model.Product
		inProduct.ID = id
		var c []model.Cocktail
		c = page.Cocktail.SelectCocktailsByProduct(inProduct)
		cs.ChildCocktails = c
		prod := inProduct.SelectProduct()
		cs.Product = prod[0]
		page.CocktailSet = cs
		page.RenderPageTemplate(w, r, "cocktails")

	}
}

//Validates the cocktail form request and populates the Cocktail
//struct
func ValidateCocktail(cocktail *model.Cocktail, r *http.Request) bool {
	cocktail.Errors = make(map[string]string)
	r.ParseForm() // Required if you don't call r.FormValue()

	if len(r.Form["cocktailID"]) > 0 {
		if _, err := strconv.Atoi(r.Form["cocktailID"][0]); err == nil {
			cocktail.ID, _ = strconv.Atoi(r.Form["cocktailID"][0])
		} else {
			glog.Errorln("Invalid CocktailID: " + r.Form["cocktailID"][0])
			cocktail.Errors["CocktailID"] = "Invalid CocktailID"
		}
	}
	if len(r.Form["cocktailTitle"]) > 0 {
		cocktail.Title = r.Form["cocktailTitle"][0]
	}
	if len(r.Form["cocktailName"]) > 0 && strings.TrimSpace(r.Form["cocktailName"][0]) != "" {
		cocktail.Name = r.Form["cocktailName"][0]
	}
	if len(r.Form["cocktailDisplayName"]) > 0 {
		cocktail.DisplayName = r.Form["cocktailDisplayName"][0]
	}
	if len(r.Form["cocktailAlternateNames"]) > 0 {
		//cocktail.AlternateNames = r.Form["cocktailAlternateNames"][0]
	}
	if len(r.Form["cocktailSpokenName"]) > 0 {
		cocktail.SpokenName = r.Form["cocktailSpokenName"][0]
	}
	if len(r.Form["cocktailOrigin"]) > 0 {
		re := regexp.MustCompile(`\r?\n`)
		cocktail.Origin = template.HTML(re.ReplaceAllString(r.Form["cocktailOrigin"][0], " "))
	}
	if len(r.Form["cocktailAKA"]) > 0 {
		//cocktail.AKA = r.Form["cocktailAKA"][0]
	}
	if len(r.Form["cocktailDescription"]) > 0 {
		re := regexp.MustCompile(`\r?\n`)
		cocktail.Description = template.HTML(re.ReplaceAllString(r.Form["cocktailDescription"][0], " "))
	}
	if len(r.Form["cocktailComment"]) > 0 {
		re := regexp.MustCompile(`\r?\n`)
		cocktail.Comment = template.HTML(re.ReplaceAllString(r.Form["cocktailComment"][0], " "))
	}
	if len(r.Form["cocktailImage"]) > 0 {
		cocktail.ImagePath, cocktail.Image = filepath.Split(r.Form["cocktailImage"][0])
		cocktail.ImagePath = strings.TrimSuffix(cocktail.ImagePath, "/")
	}
	if len(r.Form["cocktailImageSourceName"]) > 0 {
		cocktail.ImageSourceName = r.Form["cocktailImageSourceName"][0]
	}
	if len(r.Form["cocktailImageSourceLink"]) > 0 {
		cocktail.ImageSourceLink = r.Form["cocktailImageSourceLink"][0]
	}
	if len(r.Form["cocktailRating"]) > 0 {
		cocktail.Rating, _ = strconv.Atoi(r.Form["cocktailRating"][0])
	}
	if len(r.Form["cocktailSourceName"]) > 0 {
		cocktail.SourceName = r.Form["cocktailSourceName"][0]
	}
	if len(r.Form["cocktailSourceLink"]) > 0 {
		cocktail.SourceLink = r.Form["cocktailSourceLink"][0]
	}
	if len(r.Form["recipeMethod"]) > 0 {
		re := regexp.MustCompile(`\r?\n`)
		cocktail.Recipe.Method = template.HTML(re.ReplaceAllString(r.Form["recipeMethod"][0], " "))
	}
	if len(r.Form["recipeID"]) > 0 {
		cocktail.Recipe.ID, _ = strconv.Atoi(r.Form["recipeID"][0])
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
				cocktail.Recipe.RecipeSteps = append(cocktail.Recipe.RecipeSteps, rs)
				ord++
			}
		}
	}
	if len(r.Form["Garnish"]) > 0 {
		for _, id := range r.Form["Garnish"] {
			if id != "" {
				var p model.Product
				p.ID, _ = strconv.Atoi(id)
				cocktail.Garnish = append(cocktail.Garnish, p)
			}
		}
	} else {
		cocktail.Garnish = nil
	}
	if len(r.Form["Drinkware"]) > 0 {
		for _, id := range r.Form["Drinkware"] {
			if id != "" {
				var p model.Product
				p.ID, _ = strconv.Atoi(id)
				cocktail.Drinkware = append(cocktail.Drinkware, p)
			}
		}
	}
	if len(r.Form["Tool"]) > 0 {
		for _, id := range r.Form["Tool"] {
			if id != "" {
				var p model.Product
				p.ID, _ = strconv.Atoi(id)
				cocktail.Tool = append(cocktail.Tool, p)
			}
		}
	}
	if len(r.Form["Flavor"]) > 0 {
		for _, id := range r.Form["Flavor"] {
			if id != "" {
				var m model.Meta
				m.ID, _ = strconv.Atoi(id)
				cocktail.Flavor = append(cocktail.Flavor, m)
			}
		}
	}
	if len(r.Form["BaseSpirit"]) > 0 {
		for _, id := range r.Form["BaseSpirit"] {
			if id != "" {
				var m model.Meta
				m.ID, _ = strconv.Atoi(id)
				cocktail.BaseSpirit = append(cocktail.BaseSpirit, m)
			}
		}
	}
	if len(r.Form["Type"]) > 0 {
		for _, id := range r.Form["Type"] {
			if id != "" {
				var m model.Meta
				m.ID, _ = strconv.Atoi(id)
				cocktail.Type = append(cocktail.Type, m)
			}
		}
	}
	if len(r.Form["Family"]) > 0 {
		for _, id := range r.Form["Family"] {
			if id != "" {
				var m model.Meta
				m.ID, _ = strconv.Atoi(id)
				glog.Infoln("Family id = " + id)
				cocktail.Family = append(cocktail.Family, m)
			}
		}
	}
	if r.Form["IsFamilyRoot"] != nil {
		cocktail.IsFamilyRoot = true
	}
	if len(r.Form["Served"]) > 0 {
		for _, id := range r.Form["Served"] {
			if id != "" {
				var m model.Meta
				m.ID, _ = strconv.Atoi(id)
				cocktail.Served = append(cocktail.Served, m)
			}
		}
	}
	if len(r.Form["Technique"]) > 0 {
		for _, id := range r.Form["Technique"] {
			if id != "" {
				var m model.Meta
				m.ID, _ = strconv.Atoi(id)
				cocktail.Technique = append(cocktail.Technique, m)
			}
		}
	}
	if len(r.Form["Strength"]) > 0 {
		for _, id := range r.Form["Strength"] {
			if id != "" {
				var m model.Meta
				m.ID, _ = strconv.Atoi(id)
				cocktail.Strength = append(cocktail.Strength, m)
			}
		}
	}
	if len(r.Form["Difficulty"]) > 0 {
		for _, id := range r.Form["Difficulty"] {
			if id != "" {
				var m model.Meta
				m.ID, _ = strconv.Atoi(id)
				cocktail.Difficulty = append(cocktail.Difficulty, m)
			}
		}
	}
	if len(r.Form["TimeofDay"]) > 0 {
		for _, id := range r.Form["TimeofDay"] {
			if id != "" {
				var m model.Meta
				m.ID, _ = strconv.Atoi(id)
				cocktail.TOD = append(cocktail.TOD, m)
			}
		}
	}
	if len(r.Form["Ratio"]) > 0 {
		for _, id := range r.Form["Ratio"] {
			if id != "" {
				var m model.Meta
				m.ID, _ = strconv.Atoi(id)
				cocktail.Ratio = append(cocktail.Ratio, m)
			}
		}
	}
	if len(r.Form["Drink"]) > 0 {
		for _, id := range r.Form["Drink"] {
			if id != "" {
				var m model.Meta
				m.ID, _ = strconv.Atoi(id)
				cocktail.Drink = append(cocktail.Drink, m)
			}
		}
	}
	return len(cocktail.Errors) == 0
}
