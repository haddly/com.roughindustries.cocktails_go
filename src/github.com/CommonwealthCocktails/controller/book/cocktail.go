// Copyright 2017 Rough Industries LLC. All rights reserved.
//controller/www/cocktail.go: Functions and handlers for dealing with cocktails.
//TODO: migrate cocktail by meta or product id to a single function that
//is passed a meta or a product id parameter
package book

import (
	"github.com/CommonwealthCocktails/model"
	"github.com/golang/glog"
	"html/template"
	"net/http"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

//Cocktail page handler which displays the standard cocktail page.
func CocktailHandler(w http.ResponseWriter, r *http.Request, page *page) {
	var cs model.CocktailSet
	if !ValidateCocktail(w, r, page) {
		page.RenderBookTemplate(w, r, "404")
	} else {
		cs = page.Cocktail.SelectCocktailsByID(page.Cocktail.ID, true)
		page.CocktailSet = cs
		//apply the template page info to the index page
		page.RenderBookTemplate(w, r, "cocktail")
	}
}

//Validates the cocktail form request and populates the Cocktail
//struct
func ValidateCocktail(w http.ResponseWriter, r *http.Request, page *page) bool {
	page.Cocktail.Errors = make(map[string]string)
	r.ParseForm() // Required if you don't call r.FormValue()

	if len(r.Form["cocktailID"]) > 0 {
		if _, err := strconv.Atoi(r.Form["cocktailID"][0]); err == nil {
			page.Cocktail.ID, _ = strconv.Atoi(r.Form["cocktailID"][0])
		} else {
			glog.Errorln("Invalid CocktailID: " + r.Form["cocktailID"][0])
			page.Cocktail.Errors["CocktailID"] = "Invalid CocktailID"
		}
	}
	if len(r.Form["cocktailTitle"]) > 0 {
		page.Cocktail.Title = r.Form["cocktailTitle"][0]
	}
	if len(r.Form["cocktailName"]) > 0 && strings.TrimSpace(r.Form["cocktailName"][0]) != "" {
		page.Cocktail.Name = r.Form["cocktailName"][0]
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
	if len(r.Form["cocktailImage"]) > 0 {
		page.Cocktail.ImagePath, page.Cocktail.Image = filepath.Split(r.Form["cocktailImage"][0])
		page.Cocktail.ImagePath = strings.TrimSuffix(page.Cocktail.ImagePath, "/")
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
					//quant, _ := strconv.ParseFloat(r.Form["recipestep["+result[i]+"].Quantity"][0], 64)
					//rs.RecipeCardinalString = FloatToVulgar(quant)
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
				glog.Infoln("Family id = " + id)
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
	return len(page.Cocktail.Errors) == 0
}
