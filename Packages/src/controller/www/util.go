package www

import (
	"log"
	"math/rand"
	"net/http"
	"runtime"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

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

func Loader(w http.ResponseWriter) {
	page := NewPage()
	page.RenderPageTemplate(w, "loader")
}
