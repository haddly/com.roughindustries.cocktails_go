package www

import (
	"log"
	"model"
	"net/http"
)

type Search struct {
}

//handle / requests to the server
func (cocktail *Search) CocktailSearchHandler(w http.ResponseWriter, r *http.Request) {
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

func (cocktail *Search) Init() {
	log.Println("Init in www/search.go")
	http.HandleFunc("/search", cocktail.CocktailSearchHandler)

}
