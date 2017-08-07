// Package cocktail
package www

import (
	"html/template"
	"log"
	"math"
	"model"
	"net/http"
	"net/url"
	"path/filepath"
	"strconv"
	"strings"
)

// Cocktail
type Cocktail struct {
}

// CocktailHandler
func (cocktail *Cocktail) CocktailHandler(w http.ResponseWriter, r *http.Request) {
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
		if len(model.Products) <= id-1 {
			page.RenderPageTemplate(w, "404")
		} else {
			cs = model.GetCocktailByID(id)
			page.CocktailSet = cs
			//apply the template page info to the index page
			page.RenderPageTemplate(w, "cocktail")
		}
	}
}

// CocktailsHandler
func (cocktail *Cocktail) CocktailsHandler(w http.ResponseWriter, r *http.Request) {
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
	c = model.GetCocktails()
	cs.ChildCocktails = c
	page.CocktailSet = cs
	//apply the template page info to the index page
	page.RenderPageTemplate(w, "cocktails")
}

// CocktailAddFormHandler
func (cocktail *Cocktail) CocktailModFormHandler(w http.ResponseWriter, r *http.Request) {
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
	if page.Username != "" {
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
			out := model.SelectCocktailsByID(id)
			page.Cocktail = out.Cocktail
			page.RenderPageTemplate(w, "cocktailmodform")
		}
	} else {
		page.RenderPageTemplate(w, "404")
	}
}

// CocktailAddHandler
func (cocktail *Cocktail) CocktailModHandler(w http.ResponseWriter, r *http.Request) {
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

	if page.Username != "" {
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
				page.Messages["cocktailModifySuccess"] = "Cocktail modified successfully and memcache updated!"
				page.RenderPageTemplate(w, "cocktailmodform")
				return
			} else if m["button"][0] == "update" {
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

// CocktailsIndexHandler
func (cocktail *Cocktail) CocktailsIndexHandler(w http.ResponseWriter, r *http.Request) {
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

// CocktailsByMetaIDHandler
func (cocktail *Cocktail) CocktailsByMetaIDHandler(w http.ResponseWriter, r *http.Request) {
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

// CocktailsByProductIDHandler
func (cocktail *Cocktail) CocktailsByProductIDHandler(w http.ResponseWriter, r *http.Request) {
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
		log.Println("Product ID: " + m["ID"][0])
		if len(model.Products) <= id-1 {
			page.RenderPageTemplate(w, "404")
		} else {
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
}

// CocktailSearchHandler
func (cocktail *Cocktail) CocktailSearchHandler(w http.ResponseWriter, r *http.Request) {
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
	//apply the template page info to the index page
	page.RenderPageTemplate(w, "search")
}

// CocktailLandingHandler
func (cocktail *Cocktail) CocktailLandingHandler(w http.ResponseWriter, r *http.Request) {
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
	page.CocktailSet = cs
	//apply the template page info to the index page
	page.RenderPageTemplate(w, "index")
}

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
		cocktail.Origin = template.HTML(m["cocktailOrigin"][0])
	}
	if len(m["cocktailAKA"]) > 0 {
		//cocktail.AKA = m["cocktailAKA"][0]
	}
	if len(m["cocktailDescription"]) > 0 {
		cocktail.Description = template.HTML(m["cocktailDescription"][0])
	}
	if len(m["cocktailComment"]) > 0 {
		cocktail.Comment = template.HTML(m["cocktailComment"][0])
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
		cocktail.Recipe.Method = template.HTML(m["recipeMethod"][0])
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
			var p model.Product
			p.ID, _ = strconv.Atoi(id)
			cocktail.Garnish = append(cocktail.Garnish, p)
		}
	}
	if len(m["Drinkware"]) > 0 {
		for _, id := range m["Drinkware"] {
			var p model.Product
			p.ID, _ = strconv.Atoi(id)
			cocktail.Drinkware = append(cocktail.Drinkware, p)
		}
	}
	if len(m["Tool"]) > 0 {
		for _, id := range m["Tool"] {
			var p model.Product
			p.ID, _ = strconv.Atoi(id)
			cocktail.Tool = append(cocktail.Tool, p)
		}
	}
	if len(m["Flavor"]) > 0 {
		for _, id := range m["Flavor"] {
			var m model.Meta
			m.ID, _ = strconv.Atoi(id)
			cocktail.Flavor = append(cocktail.Flavor, m)
		}
	}
	if len(m["BaseSpirit"]) > 0 {
		for _, id := range m["BaseSpirit"] {
			var m model.Meta
			m.ID, _ = strconv.Atoi(id)
			cocktail.BaseSpirit = append(cocktail.BaseSpirit, m)
		}
	}
	if len(m["Type"]) > 0 {
		for _, id := range m["Type"] {
			var m model.Meta
			m.ID, _ = strconv.Atoi(id)
			cocktail.Type = append(cocktail.Type, m)
		}
	}
	if len(m["Family"]) > 0 {
		for _, id := range m["Family"] {
			var m model.Meta
			m.ID, _ = strconv.Atoi(id)
			cocktail.Family = append(cocktail.Family, m)
		}
	}
	if m["IsFamilyRoot"] != nil {
		cocktail.IsFamilyRoot = true
	}
	if len(m["Served"]) > 0 {
		for _, id := range m["Served"] {
			var m model.Meta
			m.ID, _ = strconv.Atoi(id)
			cocktail.Served = append(cocktail.Served, m)
		}
	}
	if len(m["Technique"]) > 0 {
		for _, id := range m["Technique"] {
			var m model.Meta
			m.ID, _ = strconv.Atoi(id)
			cocktail.Technique = append(cocktail.Technique, m)
		}
	}
	if len(m["Strength"]) > 0 {
		for _, id := range m["Strength"] {
			var m model.Meta
			m.ID, _ = strconv.Atoi(id)
			cocktail.Strength = append(cocktail.Strength, m)
		}
	}
	if len(m["Difficulty"]) > 0 {
		for _, id := range m["Difficulty"] {
			var m model.Meta
			m.ID, _ = strconv.Atoi(id)
			cocktail.Difficulty = append(cocktail.Difficulty, m)
		}
	}
	if len(m["TimeofDay"]) > 0 {
		for _, id := range m["TimeofDay"] {
			var m model.Meta
			m.ID, _ = strconv.Atoi(id)
			cocktail.TOD = append(cocktail.TOD, m)
		}
	}
	if len(m["Ratio"]) > 0 {
		for _, id := range m["Ratio"] {
			var m model.Meta
			m.ID, _ = strconv.Atoi(id)
			cocktail.Ratio = append(cocktail.Ratio, m)
		}
	}
	if len(m["Drink"]) > 0 {
		for _, id := range m["Drink"] {
			var m model.Meta
			m.ID, _ = strconv.Atoi(id)
			cocktail.Drink = append(cocktail.Drink, m)
		}
	}
	return len(cocktail.Errors) == 0
}

func FloatToVulgar(val float64) string {
	realPart := val
	integerPart := math.Floor(realPart)
	decimalPart := realPart - integerPart
	var intStringPart string
	if int(integerPart) == 0 {
		intStringPart = ""
	} else {
		intStringPart = strconv.Itoa(int(integerPart))
	}
	if decimalPart == 0.0 {
		return intStringPart
	} else if decimalPart <= 0.125 {
		return intStringPart + "⅛"
	} else if decimalPart <= 0.25 {
		return intStringPart + "¼"
	} else if decimalPart <= 0.375 {
		return intStringPart + "⅜"
	} else if decimalPart <= .5 {
		return intStringPart + "½"
	} else if decimalPart <= .625 {
		return intStringPart + "⅝"
	} else if decimalPart <= .75 {
		return intStringPart + "¾"
	} else if decimalPart <= .875 {
		return intStringPart + "⅞"
	}
	return strconv.Itoa(int(math.Ceil(realPart)))
}

// Init
func (cocktail *Cocktail) Init() {
	//Web Service and Web Page Handlers
	log.Println("Init in www/Cocktail.go")
	http.HandleFunc("/", cocktail.CocktailLandingHandler)
	http.HandleFunc("/cocktail", cocktail.CocktailHandler)
	http.HandleFunc("/cocktails", cocktail.CocktailsHandler)
	http.HandleFunc("/cocktailsindex", cocktail.CocktailsIndexHandler)
	http.HandleFunc("/cocktailsByMetaID", cocktail.CocktailsByMetaIDHandler)
	http.HandleFunc("/cocktailsByProductID", cocktail.CocktailsByProductIDHandler)
	http.HandleFunc("/cocktailModForm", cocktail.CocktailModFormHandler)
	http.HandleFunc("/cocktailMod", cocktail.CocktailModHandler)
}
