//controller/www/util.go: Utility functions
package www

import (
	"log"
	"math/rand"
	"net/http"
	"runtime"
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
	page := NewPage()
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	log.Printf("Recovered %s:%d %s\n", file, line, f.Name())
	log.Println(rec)
	page.RenderPageTemplate(w, "404")
}
