// Copyright 2017 Rough Industries LLC. All rights reserved.
//controller/www/search.go: Functions and handlers for dealing with searching.
package www

import (
	"model"
	"net/http"
)

//Search page handler which displays the standard search page.
func CocktailSearchHandler(w http.ResponseWriter, r *http.Request) {
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
	page.CocktailSearch = model.GetCocktailSearch()

	//apply the template page info to the index page
	page.RenderPageTemplate(w, "search")
}
