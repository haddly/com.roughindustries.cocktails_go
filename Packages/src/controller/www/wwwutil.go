// Copyright 2017 Rough Industries LLC. All rights reserved.
//controller/www/util.go: Utility functions
package www

import (
	"log"
	"math"
	"math/rand"
	"net/http"
	"runtime"
	"strconv"
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
	page := NewPage(nil)
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	log.Printf("Recovered %s:%d %s\n", file, line, f.Name())
	log.Println(rec)
	page.RenderPageTemplate(w, "404")
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
