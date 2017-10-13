// Copyright 2017 Rough Industries LLC. All rights reserved.
//controller/www/util.go: Utility functions
package book

import (
	"github.com/golang/glog"
	"net/http"
	"runtime"
)

//Helper function for producing a standard 404 page error when we through an
//panic
func Error404(w http.ResponseWriter, rec interface{}) {
	page := NewPage(nil, nil)
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	glog.Infoln("Recovered %s:%d %s\n", file, line, f.Name())
	glog.Infoln(rec)
	page.RenderBookTemplate(w, nil, "404")
}
