// Copyright 2017 Rough Industries LLC. All rights reserved.
//controller/alexa/router.go: sets up all the routing for the alexa webapp
package book

import (
	"github.com/golang/glog"
	"net/http"
	"strings"
)

//Init to setup the http handlers
func BookRouterInit() {
	glog.Infoln("Init in book/router.go")
	http.Handle("/book", RecoverHandler(MethodsHandler(PageHandler(LandingHandler), "GET")))
	http.Handle("/book/", RecoverHandler(MethodsHandler(PageHandler(LandingHandler), "GET")))
}

//This only loads the page into the page datastruct, there is no authentication
//validation
func PageHandler(next func(http.ResponseWriter, *http.Request, *page)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		page := NewPage(w, r)
		next(w, r, page)
		return
	})
}

//This handler is designed to return a 404 error after a panic has occured
func MethodsHandler(h http.Handler, methods ...string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isValidMethod := false
		for _, v := range methods {
			if strings.ToUpper(v) == r.Method {
				isValidMethod = true
			}
		}
		if !isValidMethod {
			glog.Errorln("ERROR: Invalid Method used to access content, possible attack!")
			Error404(w, "ERROR: Invalid Method used to access content, possible attack!")
			return
		}
		h.ServeHTTP(w, r) // call next
		return
	})
}

//This handler is designed to return a 404 error after a panic has occured
func RecoverHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// catch all errors and return 404
		defer func() {
			// recover from panic if one occured. Set err to nil otherwise.
			if rec := recover(); rec != nil {
				Error404(w, rec)
			}
			return
		}()
		h.ServeHTTP(w, r) // call next
		return
	})
}
