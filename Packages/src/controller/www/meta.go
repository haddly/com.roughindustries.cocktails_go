// Copyright 2017 Rough Industries LLC. All rights reserved.
//controller/www/meta.go:
package www

import (
	"github.com/golang/glog"
	"github.com/microcosm-cc/bluemonday"
	"html/template"
	"model"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

//Meta Modification Form page handler which displays the Meta Modification
//Form page.
func MetaModFormHandler(w http.ResponseWriter, r *http.Request) {
	page := NewPage(w, r)
	if page.Authenticated {
		u, err := url.Parse(r.URL.String())
		glog.Infoln(u)
		if err != nil {
			page.RenderPageTemplate(w, r, "404")
		}
		m, err := url.ParseQuery(u.RawQuery)
		if err != nil {
			page.RenderPageTemplate(w, r, "404")
		}
		var mbt model.MetasByTypes
		mbt = page.Meta.SelectMetaByTypes(false, true, false)
		page.MetasByTypes = mbt
		if len(m["ID"]) == 0 {
			//apply the template page info to the index page
			page.RenderPageTemplate(w, r, "metamodform")
		} else {
			id, _ := strconv.Atoi(m["ID"][0])
			var in model.Meta
			in.ID = id
			out := in.SelectMeta()
			page.Meta = out[0]
			page.RenderPageTemplate(w, r, "metamodform")
		}
	} else {
		glog.Errorln("ERROR: User not authenticated and requesting meta data modification form, possible attack!")
		http.Redirect(w, r, "/", 302)
	}
}

//Meta modification form page request handler which process the meta
//modification request.  This will after verifying a valid user session,
//modify the meta data based on the request.
func MetaModHandler(w http.ResponseWriter, r *http.Request) {
	page := NewPage(w, r)
	if page.Authenticated {
		//Get the generic data that all meta mod pages need
		var mbt model.MetasByTypes
		mbt = page.Meta.SelectMetaByTypes(false, true, false)
		page.MetasByTypes = mbt
		//Validate the form input and populate the meta data
		if ValidateMeta(&page.Meta, r) {
			//validate the CSRF ID to make sure we don't double submit or
			//have an attack
			if !ValidateCSRF(r, page) {
				page.RenderPageTemplate(w, r, "metamodform")
				return
			} else {
				//did we get an add, update, or something else request
				if r.Form["button"][0] == "add" {
					ret_id := page.Meta.InsertMeta()
					model.LoadMCWithMetaData()
					page.Meta.ID = ret_id
					outMeta := page.Meta.SelectMeta()
					page.Meta = outMeta[0]
					page.Messages["metaModifySuccess"] = "Metadata modified successfully and memcache updated!"
					page.RenderPageTemplate(w, r, "metamodform")
					return
				} else if r.Form["button"][0] == "update" {
					rows_updated := page.Meta.UpdateMeta()
					model.LoadMCWithMetaData()
					glog.Infoln("Updated " + strconv.Itoa(rows_updated) + " rows")
					outMeta := page.Meta.SelectMeta()
					page.Meta = outMeta[0]
					page.Messages["metaModifySuccess"] = "Metadata modified successfully and memcache updated!"
					page.RenderPageTemplate(w, r, "metamodform")
					return
				} else if len(r.Form["ID"]) != 0 {
					id, _ := strconv.Atoi(r.Form["ID"][0])
					var in model.Meta
					in.ID = id
					out := in.SelectMeta()
					page.Meta = out[0]
					page.RenderPageTemplate(w, r, "metamodform")
				} else {
					//we only allow add and update right now
					page.Messages["metaModifyFail"] = "Metadata modification failed.  You tried to perform an unknown operation!"
					page.RenderPageTemplate(w, r, "metamodform")
					return
				}
			}
		} else {
			//Validation failed
			glog.Infoln("Bad meta!")
			if len(r.Form["ID"]) == 0 {
				//apply the template page info to the index page
				page.RenderPageTemplate(w, r, "metamodform")
			} else {
				id, _ := strconv.Atoi(r.Form["ID"][0])
				var in model.Meta
				in.ID = id
				out := in.SelectMeta()
				page.Meta = out[0]
				page.RenderPageTemplate(w, r, "metamodform")
			}
			return
		}
	} else {
		glog.Errorln("ERROR: User not authenticated and requesting update on meta data, possible attack!")
		http.Redirect(w, r, "/", 302)
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
	//Required
	if len(r.Form["metaName"]) > 0 && strings.TrimSpace(r.Form["metaName"][0]) != "" {
		meta.MetaName = r.Form["metaName"][0]
	} else {
		meta.Errors["MetaName"] = "Please enter a valid meta name"
	}
	//Required
	if len(r.Form["metaType"]) > 0 {
		meta.MetaType.ID, _ = strconv.Atoi(r.Form["metaType"][0])
	} else {
		meta.Errors["MetaType"] = "Please select a valid meta type"
	}
	//Can be blank
	if len(r.Form["metaBlurb"]) > 0 {
		//sanitize the input, we don't want XSS
		p := bluemonday.UGCPolicy()
		p.AllowElements("img")
		html := p.Sanitize(r.Form["metaBlurb"][0])
		meta.Blurb = template.HTML(html)
	} else {
		meta.Blurb = ""
	}
	return len(meta.Errors) == 0
}
