//controller/www/memcache.go: Functions and handlers for interacting with the
//memcache at a high level.  This includes delteing and loading large sets
//of data from the database.
package www

import (
	"bytes"
	"html/template"
	"model"
	"net/http"
)

//Memcache struct for defining methods to
type Memcache struct {
}

//Init to setup the http handlers
func (memcache *Memcache) Init() {
	http.HandleFunc("/mc_delete", memcache.MCDeleteHandler)
	http.HandleFunc("/mc_load", memcache.MCAddHandler)
}

//delete all the memcache entries
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
	if page.Username != "" && page.Authenticated {
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

//load all the memcache entries from the database
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
	if page.Username != "" && page.Authenticated {
		var buffer bytes.Buffer
		buffer.WriteString("<b>Memcache Add</b>:<br/>")
		model.LoadAllMemcache()
		//apply the template page info to the index page
		statStr := buffer.String()
		page.Messages["Status"] = template.HTML(statStr)
		page.RenderPageTemplate(w, "mcindex")
	}
}
