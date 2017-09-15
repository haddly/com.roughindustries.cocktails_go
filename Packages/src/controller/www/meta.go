// Copyright 2017 Rough Industries LLC. All rights reserved.
//controller/www/meta.go:
package www

import (
	"html/template"
	"log"
	"model"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

//Meta Modification Form page handler which displays the Meta Modification
//Form page.
func MetaModFormHandler(w http.ResponseWriter, r *http.Request) {
	page := NewPage(r)
	if page.Username != "" && page.Authenticated {
		u, err := url.Parse(r.URL.String())
		log.Println(u)
		if err != nil {
			page.RenderPageTemplate(w, "404")
		}
		m, err := url.ParseQuery(u.RawQuery)
		if err != nil {
			page.RenderPageTemplate(w, "404")
		}
		var mbt model.MetasByTypes
		mbt = page.Meta.SelectMetaByTypes(false, true, false)
		page.MetasByTypes = mbt
		if len(m["ID"]) == 0 {
			//apply the template page info to the index page
			page.RenderPageTemplate(w, "metamodform")
		} else {
			id, _ := strconv.Atoi(m["ID"][0])
			var in model.Meta
			in.ID = id
			out := in.SelectMeta()
			page.Meta = out[0]
			page.RenderPageTemplate(w, "metamodform")
		}
	} else {
		page.RenderPageTemplate(w, "404")
	}
}

//Meta modification form page request handler which process the meta
//modification request.  This will after verifying a valid user session,
//modify the meta data based on the request.
func MetaModHandler(w http.ResponseWriter, r *http.Request) {
	page := NewPage(r)
	if page.Username != "" && page.Authenticated {
		//Get the generic data that all meta mod pages need
		var mbt model.MetasByTypes
		mbt = page.Meta.SelectMetaByTypes(false, true, false)
		page.MetasByTypes = mbt
		//Validate the form input and populate the meta data
		if ValidateMeta(&page.Meta, r) {
			//did we get an add, update, or something else request
			if r.Form["button"][0] == "add" {
				ret_id := page.Meta.InsertMeta()
				model.LoadMCWithMetaData()
				page.Meta.ID = ret_id
				outMeta := page.Meta.SelectMeta()
				page.Meta = outMeta[0]
				page.Messages["metaModifySuccess"] = "Metadata modified successfully and memcache updated!"
				page.RenderPageTemplate(w, "metamodform")
				return
			} else if r.Form["button"][0] == "update" {
				rows_updated := page.Meta.UpdateMeta()
				model.LoadMCWithMetaData()
				log.Println("Updated " + strconv.Itoa(rows_updated) + " rows")
				outMeta := page.Meta.SelectMeta()
				page.Meta = outMeta[0]
				page.Messages["metaModifySuccess"] = "Metadata modified successfully and memcache updated!"
				page.RenderPageTemplate(w, "metamodform")
				return
			} else if len(r.Form["ID"]) != 0 {
				id, _ := strconv.Atoi(r.Form["ID"][0])
				var in model.Meta
				in.ID = id
				out := in.SelectMeta()
				page.Meta = out[0]
				page.RenderPageTemplate(w, "metamodform")
			} else {
				//we only allow add and update right now
				page.Messages["metaModifyFail"] = "Metadata modification failed.  You tried to perform an unknown operation!"
				page.RenderPageTemplate(w, "metamodform")
				return
			}
		} else {
			//Validation failed
			log.Println("Bad meta!")
			if len(r.Form["ID"]) == 0 {
				//apply the template page info to the index page
				page.RenderPageTemplate(w, "metamodform")
			} else {
				id, _ := strconv.Atoi(r.Form["ID"][0])
				var in model.Meta
				in.ID = id
				out := in.SelectMeta()
				page.Meta = out[0]
				page.RenderPageTemplate(w, "metamodform")
			}
			return
		}
	} else {
		page.RenderPageTemplate(w, "404")
		return
	}
}

//Parses the form and then validates the meta form request and
//populates the Meta struct
func ValidateMeta(meta *model.Meta, r *http.Request) bool {
	meta.Errors = make(map[string]string)
	r.ParseForm() // Required if you don't call r.FormValue()

	if len(r.Form["metaID"]) > 0 {
		meta.ID, _ = strconv.Atoi(r.Form["metaID"][0])
	}
	if len(r.Form["metaName"]) > 0 && strings.TrimSpace(r.Form["metaName"][0]) != "" {
		meta.MetaName = r.Form["metaName"][0]
	} else {
		meta.Errors["MetaName"] = "Please enter a valid meta name"
	}
	if len(r.Form["metaType"]) > 0 {
		meta.MetaType.ID, _ = strconv.Atoi(r.Form["metaType"][0])
	} else {
		meta.Errors["MetaType"] = "Please select a valid meta type"
	}
	if len(r.Form["metaBlurb"]) > 0 {
		meta.Blurb = template.HTML(r.Form["metaBlurb"][0])
	}
	return len(meta.Errors) == 0
}
