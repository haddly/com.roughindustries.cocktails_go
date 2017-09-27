// Copyright 2017 Rough Industries LLC. All rights reserved.
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

//delete all the memcache entries
func MCDeleteHandler(w http.ResponseWriter, r *http.Request) {
	page := NewPage(w, r)
	if page.Authenticated {
		var buffer bytes.Buffer
		buffer.WriteString("<b>Memcache Delete</b>:<br/>")
		model.DeleteAllMemcache()

		//apply the template page info to the index page
		statStr := buffer.String()
		page.Messages["Status"] = template.HTML(statStr)
		page.RenderPageTemplate(w, r, "mcindex")
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

//load all the memcache entries from the database
func MCAddHandler(w http.ResponseWriter, r *http.Request) {
	page := NewPage(w, r)
	if page.Authenticated {
		var buffer bytes.Buffer
		buffer.WriteString("<b>Memcache Add</b>:<br/>")
		model.LoadAllMemcache()
		//apply the template page info to the index page
		statStr := buffer.String()
		page.Messages["Status"] = template.HTML(statStr)
		page.RenderPageTemplate(w, r, "mcindex")
	} else {
		http.Redirect(w, r, "/", 302)
	}
}
