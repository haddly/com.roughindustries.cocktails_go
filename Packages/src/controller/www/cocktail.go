// Package cocktail
package www

import (
	"log"
	"model"
	"net/http"
	"net/url"
	"strconv"
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
func (cocktail *Cocktail) CocktailAddFormHandler(w http.ResponseWriter, r *http.Request) {
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
		//apply the template page info to the index page
		page.RenderPageTemplate(w, "cocktailaddform")
	} else {
		page.RenderPageTemplate(w, "404")
	}
}

// CocktailAddHandler
func (cocktail *Cocktail) CocktailAddHandler(w http.ResponseWriter, r *http.Request) {
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
	log.Println("In Add Cocktail handler")
	u, err := url.Parse(r.URL.String())
	log.Println(u)
	if err != nil {
		page.RenderPageTemplate(w, "404")
	}
	m, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		page.RenderPageTemplate(w, "404")
	}
	log.Println(m)
	page.RenderPageTemplate(w, "404")
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
	http.HandleFunc("/cocktailAddForm", cocktail.CocktailAddFormHandler)
	http.HandleFunc("/cocktailAdd", cocktail.CocktailAddHandler)
}
