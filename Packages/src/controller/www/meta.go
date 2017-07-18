//www/meta.go
package www

import (
	"html/template"
	"log"
	"model"
	"net/http"
	"net/url"
	"strconv"
)

type Meta struct {
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
	if page.Username != "" {
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
			var inMeta model.Meta
			inMeta.ID = id
			outMeta := model.SelectMeta(inMeta)
			page.Meta = outMeta[0]
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
	if page.Username != "" {
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
		log.Println(m["metaID"][0])
		log.Println(m["metaName"][0])
		log.Println(m["metaType"][0])
		log.Println(m["metaBlurb"][0])

		var inMeta model.Meta
		inMeta.ID, _ = strconv.Atoi(m["metaID"][0])
		inMeta.MetaName = m["metaName"][0]
		inMeta.MetaType.ID, _ = strconv.Atoi(m["metaType"][0])
		inMeta.Blurb = template.HTML(m["metaBlurb"][0])
		if len(m["add"]) > 0 {
			ret_id := model.ProcessMeta(inMeta)
			inMeta.ID = ret_id
			outMeta := model.SelectMeta(inMeta)
			page.Meta = outMeta[0]
			var mbt model.MetasByTypes
			model.LoadMCWithMetaData()
			mbt = model.GetMetaByTypes(false, true, false)
			page.MetasByTypes = mbt
			page.Messages["metaModifySuccess"] = "Metadata modified successfully and memcache updated!"
			page.RenderPageTemplate(w, "metamodform")
			return
		} else {
			outMeta := model.SelectMeta(inMeta)
			page.Meta = outMeta[0]
			page.Messages["metaModifyFail"] = "Metadata modification failed.  You tried to perform an unknown operation!"
			page.RenderPageTemplate(w, "metamodform")
			return
		}
	} else {
		page.RenderPageTemplate(w, "404")
		return
	}
}

func (meta *Meta) Init() {
	log.Println("Init in www/meta.go")
	http.HandleFunc("/metaModForm", meta.MetaModFormHandler)
	http.HandleFunc("/metaMod", meta.MetaModHandler)
}
