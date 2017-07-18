//main/main
package main

import (
	"connectors"
	"controller/alexa"
	"controller/www"
	"flag"
	"github.com/gorilla/context"
	"log"
	"math/rand"
	"model"
	"net/http"
	//"net/url"
	"os"
	"path/filepath"
	"time"
)

//The controllers
var d = www.Database{}
var mem = www.Memcache{}
var c = www.Cocktail{}
var p = www.Product{}
var m = www.Meta{}
var s = www.Search{}
var l = www.Login{}
var post = www.Post{}
var page = www.NewPage()

var a = alexa.Hello{}

func init() {

	var datasource = model.DSTtoi("DB")

	if datasource == model.DB {
		//SET THESE LINES AND ADD #gitignore to the end of the line as a comment to ignore your info
		var dbaddr string
		var dbpasswd string
		var user string
		var proto string
		var port string
		var dbname string
		//dbaddr = ??
		//dbpasswd = ??
		//user = ??
		//proto = ??
		//port = ??
		//dbname = ??

		var mc_server string
		//mc_server = ??

		if dbaddr == "" {
			log.Println("No DB Address set.")
			os.Exit(0)
		}
		if dbpasswd == "" {
			log.Println("No DB password set.")
			os.Exit(0)
		}
		if user == "" {
			log.Println("No DB user set.")
			os.Exit(0)
		}
		if proto == "" {
			log.Println("No DB protocol set.")
			os.Exit(0)
		}
		if port == "" {
			log.Println("No DB port set.")
			os.Exit(0)
		}
		if dbname == "" {
			log.Println("No DB name set.")
			os.Exit(0)
		}
		connectors.SetDBVars(dbaddr, dbpasswd, user, proto, port, dbname)
		if mc_server != "" {
			connectors.SetMCVars(mc_server)
		} else {
			log.Println("No Memcache server set. If you want to use memcaching you will " +
				"have to set this value in main.go.")
		}
	}

	model.SetDataSourceType(datasource)
	flag.Parse()
	// wanted it to be more random so i seed it time now
	rand.Seed(time.Now().UnixNano())
	d.Init()
	mem.Init()
	c.Init()
	p.Init()
	m.Init()
	s.Init()
	a.Init()
	post.Init()
	l.Init()
	page.Init()
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
	http.HandleFunc("/memcache", func(w http.ResponseWriter, r *http.Request) {
		model.LoadMCWithProductData()
		model.LoadMCWithMetaData()
	})
	log.Println("Added Handlers ... Starting Server\n")
	//this starts up the server
	http.ListenAndServe(":8080", context.ClearHandler(http.DefaultServeMux))
}
