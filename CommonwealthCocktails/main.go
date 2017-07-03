//main/main
package main

import (
	"controller/alexa"
	"controller/www"
	"db"
	"flag"
	//"io/ioutil"
	"log"
	"math/rand"
	"model"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

//The controllers
var d = www.Database{}
var c = www.Cocktail{}
var p = www.Product{}
var m = www.Meta{}
var s = www.Search{}
var l = www.Login{}
var post = www.Post{}

var a = alexa.Hello{}

func init() {

	var datasource = model.DSTtoi("DB")

	if datasource == model.DB {
		//SET THESE LINES AND ADD #gitignore to the end of the line as a comment to ignore your info
		var dbaddr string
		//dbaddr = ??
		if dbaddr == "" {
			log.Println("No DB Address set.")
			os.Exit(0)
		}
		var dbpasswd string
		//dbpasswd = ??
		if dbpasswd == "" {
			log.Println("No DB password set.")
			os.Exit(0)
		}
		var user string
		//user = ??
		if user == "" {
			log.Println("No DB user set.")
			os.Exit(0)
		}
		var proto string
		//proto = ??
		if proto == "" {
			log.Println("No DB protocol set.")
			os.Exit(0)
		}
		var port string
		//port = ??
		if port == "" {
			log.Println("No DB port set.")
			os.Exit(0)
		}
		var dbname string
		//dbname = ??
		if dbname == "" {
			log.Println("No DB name set.")
			os.Exit(0)
		}
		db.SetDBVars(dbaddr, dbpasswd, user, proto, port, dbname)
	}

	model.SetDataSourceType(datasource)
	flag.Parse()
	// wanted it to be more random so i seed it time now
	rand.Seed(time.Now().UnixNano())
	d.Init()
	c.Init()
	p.Init()
	m.Init()
	s.Init()
	a.Init()
	post.Init()
	l.Init()
	log.Println("End Init")
}

func main() {
	//log.SetOutput(ioutil.Discard)
	log.Println("Starting ... \n")
	//print out the current directory
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(dir)

	// Mandatory root-based resources and redirects for other resources
	// This is handled in the app.yaml for google cloud platform deployments
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./view/webcontent/www/images"))))
	http.Handle("/font-awesome/", http.StripPrefix("/font-awesome/", http.FileServer(http.Dir("./view/webcontent/www/font-awesome"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./view/webcontent/www/css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./view/webcontent/www/js"))))
	http.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("./view/webcontent/www/fonts"))))
	http.Handle("/slick/", http.StripPrefix("/slick/", http.FileServer(http.Dir("./view/webcontent/www/libs/slick"))))
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "view/webcontent/www/favicon.ico")
	})

	log.Println("Added Handlers ... Starting Server\n")
	//this starts up the server
	http.ListenAndServe(":8080", nil)
}
