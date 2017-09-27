// Copyright 2017 Rough Industries LLC. All rights reserved.
//controller/www/util.go: Utility functions
package www

import (
	"github.com/golang/glog"
	"math"
	"math/rand"
	"net/http"
	"runtime"
	"strconv"
	"time"
)

//all alpha numeric ascii characters upper and lower case
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

//Generate a random sequence of length n characters from the alpha numeric
//ascii characters upper and lower case
func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

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
	page.RenderPageTemplate(w, nil, "404")
}

//Validates the CSRF ID from the last page request.  True means the CSRf is good,
//and false means that the CSRF is not the one past to the previous page.
//This could indicate a CSRF attack.
func ValidateCSRF(r *http.Request, page *page) bool {
	if len(r.Form["CSRF"]) > 0 {
		if (r.Form["CSRF"][0] != page.UserSession.CSRF) || (time.Since(page.UserSession.CSRFGenTime).Minutes() > 10) {
			page.Messages["metaModifyFail"] = "Metadata modification failed.  You took more then 10 minutes to enter data or you tried to navigate backwards and resubmit!"
			if r.Form["CSRF"][0] != page.UserSession.CSRF {
				glog.Errorln("ERROR: Incorrect CSRF, possible CSRF attack!")
			}
			return false
		}
	} else {
		panic("ERROR: No CSRF ID provided, possible CSRF attack!")
	}
	return true
}

//Converts a float value to a vulgar fractional string i.e. .5 to ½
func FloatToVulgar(val float64) string {
	realPart := val
	integerPart := math.Floor(realPart)
	decimalPart := realPart - integerPart
	var intStringPart string
	if int(integerPart) == 0 {
		intStringPart = ""
	} else {
		intStringPart = strconv.Itoa(int(integerPart))
	}
	if decimalPart == 0.0 {
		return intStringPart
	} else if decimalPart <= 0.125 {
		return intStringPart + "⅛"
	} else if decimalPart <= 0.25 {
		return intStringPart + "¼"
	} else if decimalPart <= 0.375 {
		return intStringPart + "⅜"
	} else if decimalPart <= .5 {
		return intStringPart + "½"
	} else if decimalPart <= .625 {
		return intStringPart + "⅝"
	} else if decimalPart <= .75 {
		return intStringPart + "¾"
	} else if decimalPart <= .875 {
		return intStringPart + "⅞"
	}
	return strconv.Itoa(int(math.Ceil(realPart)))
}
