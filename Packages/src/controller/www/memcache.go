//www/memcache.go
package www

import (
	"bytes"
	"html/template"
	"model"
	"net/http"
)

type Memcache struct {
}

func (memcache *Memcache) MCDeleteHandler(w http.ResponseWriter, r *http.Request) {
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
		var buffer bytes.Buffer
		buffer.WriteString("<b>Memcache Delete</b>:<br/>")
		model.DeleteAllMemcache()
		page.Username, page.Authenticated = GetSession(r)
		//apply the template page info to the index page
		statStr := buffer.String()
		page.Messages["Status"] = template.HTML(statStr)
		page.RenderPageTemplate(w, "mcindex")
	}
}

func (memcache *Memcache) MCAddHandler(w http.ResponseWriter, r *http.Request) {
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
		var buffer bytes.Buffer
		buffer.WriteString("<b>Memcache Add</b>:<br/>")
		model.LoadAllMemcache()
		//apply the template page info to the index page
		statStr := buffer.String()
		page.Messages["Status"] = template.HTML(statStr)
		page.RenderPageTemplate(w, "mcindex")
	}
}

func (memcache *Memcache) Init() {
	http.HandleFunc("/mc_delete", memcache.MCDeleteHandler)
	http.HandleFunc("/mc_load", memcache.MCAddHandler)
}
