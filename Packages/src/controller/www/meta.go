//www/meta.go
package www

import (
	"log"
	"model"
	"net/http"
	"net/url"
)

type Meta struct {
}

//render the page based on the name of the file provided
func (meta *Meta) RenderPageTemplate(w http.ResponseWriter, tmpl string, page *Page) {
	t, err := parseTempFiles(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	err = t.ExecuteTemplate(w, "base", page)
	if err != nil {
		log.Fatal(err)
	}
}

func (meta *Meta) MetaAddFormHandler(w http.ResponseWriter, r *http.Request) {
	userName := GetUserName(r)
	if userName != "" {
		var page Page
		page.Username = GetUserName(r)
		var mbt model.MetasByTypes
		mbt = model.GetMetaByTypes(false, true)
		page.MetasByTypes = mbt
		//apply the template page info to the index page
		meta.RenderPageTemplate(w, "metaaddform", &page)
	} else {
		meta.RenderPageTemplate(w, "404", nil)
	}
}

//handle / requests to the server
func (meta *Meta) MetaAddHandler(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.String())
	log.Println(u)
	if err != nil {
		meta.RenderPageTemplate(w, "404", nil)
	}
	m, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		meta.RenderPageTemplate(w, "404", nil)
	}
	log.Println(m)
	meta.RenderPageTemplate(w, "404", nil)
}

func (meta *Meta) Init() {
	log.Println("Init in www/meta.go")
	http.HandleFunc("/metaAddForm", meta.MetaAddFormHandler)
	http.HandleFunc("/metaAdd", meta.MetaAddHandler)
}
