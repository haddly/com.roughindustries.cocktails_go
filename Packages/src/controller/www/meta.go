//www/meta.go
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

type Meta struct {
}

func (meta *Meta) Init() {
	log.Println("Init in www/meta.go")
	http.HandleFunc("/metaModForm", meta.MetaModFormHandler)
	http.HandleFunc("/metaMod", meta.MetaModHandler)
}

func (meta *Meta) MetaModFormHandler(w http.ResponseWriter, r *http.Request) {
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
		mbt = model.GetMetaByTypes(false, true, false)
		page.MetasByTypes = mbt
		if len(m["ID"]) == 0 {
			//apply the template page info to the index page
			page.RenderPageTemplate(w, "metamodform")
		} else {
			id, _ := strconv.Atoi(m["ID"][0])
			var in model.Meta
			in.ID = id
			out := model.SelectMeta(in)
			page.Meta = out[0]
			page.RenderPageTemplate(w, "metamodform")
		}
	} else {
		page.RenderPageTemplate(w, "404")
	}
}

//handle / requests to the server
func (meta *Meta) MetaModHandler(w http.ResponseWriter, r *http.Request) {
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
		u, err := url.Parse(r.URL.String())
		log.Println(u)
		if err != nil {
			page.RenderPageTemplate(w, "404")
			return
		}
		m, err := url.ParseQuery(u.RawQuery)
		if err != nil {
			page.RenderPageTemplate(w, "404")
			return
		}
		log.Println(m)

		if ValidateMeta(&page.Meta, m) {
			if m["button"][0] == "add" {
				ret_id := model.InsertMeta(page.Meta)
				var mbt model.MetasByTypes
				model.LoadMCWithMetaData()
				mbt = model.GetMetaByTypes(false, true, false)
				page.MetasByTypes = mbt
				page.Meta.ID = ret_id
				outMeta := model.SelectMeta(page.Meta)
				page.Meta = outMeta[0]
				page.Messages["metaModifySuccess"] = "Metadata modified successfully and memcache updated!"
				page.RenderPageTemplate(w, "metamodform")
				return
			} else if m["button"][0] == "update" {
				rows_updated := model.UpdateMeta(page.Meta)
				log.Println("Updated " + strconv.Itoa(rows_updated) + " rows")
				var mbt model.MetasByTypes
				model.LoadMCWithMetaData()
				mbt = model.GetMetaByTypes(false, true, false)
				page.MetasByTypes = mbt
				outMeta := model.SelectMeta(page.Meta)
				page.Meta = outMeta[0]
				page.Messages["metaModifySuccess"] = "Metadata modified successfully and memcache updated!"
				page.RenderPageTemplate(w, "metamodform")
				return
			} else {
				var mbt model.MetasByTypes
				mbt = model.GetMetaByTypes(false, true, false)
				page.MetasByTypes = mbt
				outMeta := model.SelectMeta(page.Meta)
				page.Meta = outMeta[0]
				page.Messages["metaModifyFail"] = "Metadata modification failed.  You tried to perform an unknown operation!"
				page.RenderPageTemplate(w, "metamodform")
				return
			}
		} else {
			var mbt model.MetasByTypes
			mbt = model.GetMetaByTypes(false, true, false)
			page.MetasByTypes = mbt
			log.Println("Bad meta!")
			page.RenderPageTemplate(w, "/metamodform")
			return
		}
	} else {
		page.RenderPageTemplate(w, "404")
		return
	}
}

func ValidateMeta(meta *model.Meta, m map[string][]string) bool {
	meta.Errors = make(map[string]string)
	if len(m["metaID"]) > 0 {
		meta.ID, _ = strconv.Atoi(m["metaID"][0])
	}
	if len(m["metaName"]) > 0 && strings.TrimSpace(m["metaName"][0]) != "" {
		meta.MetaName = m["metaName"][0]
	} else {
		meta.Errors["MetaName"] = "Please enter a valid meta name"
	}
	if len(m["metaType"]) > 0 {
		meta.MetaType.ID, _ = strconv.Atoi(m["metaType"][0])
	} else {
		meta.Errors["MetaType"] = "Please select a valid meta type"
	}
	if len(m["metaBlurb"]) > 0 {
		meta.Blurb = template.HTML(m["metaBlurb"][0])
	}
	return len(meta.Errors) == 0
}
