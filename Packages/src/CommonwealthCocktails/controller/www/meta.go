// Copyright 2017 Rough Industries LLC. All rights reserved.
//controller/www/meta.go:
package www

import (
	"github.com/asaskevich/govalidator"
	"github.com/golang/glog"
	"github.com/microcosm-cc/bluemonday"
	"html/template"
	"CommonwealthCocktails/model"
	"net/http"
	"strconv"
	"strings"
)

//Meta Modification Form page handler which displays the Meta Modification
//Form page.
func MetaModFormHandler(w http.ResponseWriter, r *http.Request, page *page) {
	var mbt model.MetasByTypes
	mbt = page.Meta.SelectMetaByTypes(false, true, false)
	page.MetasByTypes = mbt
	if page.Meta.ID == 0 {
		//apply the template page info to the index page
		page.RenderPageTemplate(w, r, "metamodform")
	} else {
		out := page.Meta.SelectMeta()
		page.Meta = out[0]
		page.RenderPageTemplate(w, r, "metamodform")
	}
}

//Meta modification form page request handler which process the meta
//modification request.  This will after verifying a valid user session,
//modify the meta data based on the request.
func MetaModHandler(w http.ResponseWriter, r *http.Request, page *page) {
	//Get the generic data that all meta mod pages need
	var mbt model.MetasByTypes
	mbt = page.Meta.SelectMetaByTypes(false, true, false)
	page.MetasByTypes = mbt
	//did we get an add, update, or something else request
	if page.SubmitButtonString == "add" {
		ret_id := page.Meta.InsertMeta()
		model.LoadMCWithMetaData()
		page.Meta.ID = ret_id
		outMeta := page.Meta.SelectMeta()
		page.Meta = outMeta[0]
		page.Messages["metaModifySuccess"] = "Metadata modified successfully and memcache updated!"
		page.RenderPageTemplate(w, r, "metamodform")
		return
	} else if page.SubmitButtonString == "update" {
		rows_updated := page.Meta.UpdateMeta()
		model.LoadMCWithMetaData()
		glog.Infoln("Updated " + strconv.Itoa(rows_updated) + " rows")
		outMeta := page.Meta.SelectMeta()
		page.Meta = outMeta[0]
		page.Messages["metaModifySuccess"] = "Metadata modified successfully and memcache updated!"
		page.RenderPageTemplate(w, r, "metamodform")
		return
	} else {
		//we only allow add and update right now
		page.Messages["metaModifyFail"] = "Metadata modification failed.  You tried to perform an unknown operation!"
		page.RenderPageTemplate(w, r, "metamodform")
		return
	}
}

//Parses the form and then validates the meta form request and
//populates the Meta struct
func ValidateMeta(w http.ResponseWriter, r *http.Request, page *page) bool {
	page.Meta.Errors = make(map[string]string)
	r.ParseForm() // Required if you don't call r.FormValue()
	pUGCP := bluemonday.UGCPolicy()
	pUGCP.AllowElements("img")
	pSP := bluemonday.StrictPolicy()
	if len(r.Form["metaID"]) > 0 && strings.TrimSpace(r.Form["metaID"][0]) != "" {
		if govalidator.IsInt(r.Form["metaID"][0]) {
			page.Meta.ID, _ = strconv.Atoi(r.Form["metaID"][0])
		} else {
			page.Meta.Errors["MetaID"] = "Please enter a valid meta id. "
		}
	}
	if len(r.Form["button"]) > 0 && strings.TrimSpace(r.Form["button"][0]) != "" {
		if govalidator.IsAlpha(r.Form["button"][0]) {
			page.SubmitButtonString = pSP.Sanitize(r.Form["button"][0])
		} else {
			page.Meta.Errors["button"] = "Please click a valid button. "
		}
	}
	if len(r.Form["metaName"]) > 0 && strings.TrimSpace(r.Form["metaName"][0]) != "" {
		if govalidator.IsPrintableASCII(r.Form["metaName"][0]) {
			page.Meta.MetaName = pSP.Sanitize(r.Form["metaName"][0])
		} else {
			page.Meta.Errors["MetaName"] = "Please enter a valid meta name. "
		}
	}
	//Required
	if len(r.Form["metaType"]) > 0 && strings.TrimSpace(r.Form["metaType"][0]) != "" {
		mtID, _ := strconv.Atoi(r.Form["metaType"][0])
		if govalidator.IsInt(r.Form["metaType"][0]) && mtID <= len(model.MetaTypeStrings) {
			page.Meta.MetaType.ID = mtID
		} else {
			page.Meta.Errors["MetaType"] = "Please select a valid meta type. "
		}
	}
	//Can be blank
	if len(r.Form["metaBlurb"]) > 0 && strings.TrimSpace(r.Form["metaBlurb"][0]) != "" {
		if govalidator.IsPrintableASCII(r.Form["metaBlurb"][0]) {
			//sanitize the input, we don't want XSS
			html := pUGCP.Sanitize(r.Form["metaBlurb"][0])
			page.Meta.Blurb = template.HTML(html)
		} else {
			page.Meta.Errors["MetaBlurb"] = "Please enter a valid meta blurb. "
		}
	} else {
		page.Meta.Blurb = ""
	}
	if len(page.Meta.Errors) > 0 {
		page.Errors["metaErrors"] = "You have errors in your input. "
	}
	return len(page.Meta.Errors) == 0
}

//Checks the page meta struct that required fields are filled in.
func RequiredMetaMod(page *page) bool {
	missingRequired := false
	if strings.TrimSpace(page.Meta.MetaName) == "" {
		page.Meta.Errors["MetaName"] = "Meta name is required."
		missingRequired = true
	}
	if page.Meta.MetaType.ID == 0 {
		page.Meta.Errors["MetaType"] = "Meta type is required."
		missingRequired = true
	}
	return missingRequired
}
