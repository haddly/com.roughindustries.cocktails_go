// Copyright 2017 Rough Industries LLC. All rights reserved.
//CommonwealthCocktails/main.go:
package main

import (
	"connectors"
	"controller/alexa"
	"controller/www"
	"flag"
	"github.com/gorilla/context"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

//Initialize the database connection, memcache connection, and all the
//controllers
func init() {

	//SET THESE LINES AND ADD #gitignore to the end of the line as a comment to ignore your info
	var dbaddr string
	var dbpasswd string
	var user string
	var proto string
	var port string
	var dbname string
	var dbtype connectors.DBTypesConst
	//dbaddr = ??
	//dbpasswd = ??
	//user = ??
	//proto = ??
	//port = ??
	//dbname = ??
	//dbtype = www.DBTypesConst.MySQL

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
	connectors.SetDBVars(dbaddr, dbpasswd, user, proto, port, dbname, dbtype)
	if mc_server != "" {
		connectors.SetMCVars(mc_server)
	} else {
		log.Println("No Memcache server set. If you want to use memcaching you will " +
			"have to set this value in main.go.")
	}

	flag.Parse()
	// wanted it to be more random so i seed it time now
	rand.Seed(time.Now().UnixNano())
	// init the routing
	www.WWWRouterInit()
	alexa.AlexaRouterInit()
	log.Println("End Init")
}

//Let's Encrypt HTTPS setup, Challenges are based on the number of domains
func AddLetsEncrypt() {
	//Challenge 1
	//http.HandleFunc("/.well-known/acme-challenge/{Your Let's Encrypt Page Code}", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "{Your Let's Encrypt Conent Code}")
	//})
	//Challenge N...
	//http.HandleFunc("/.well-known/acme-challenge/{Your Let's Encrypt Page Code}", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "{Your Let's Encrypt Conent Code}")
	//})
}

func main() {
	log.Println("Initializing ... \n")
	//print out the current directory
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(dir)
	AddLetsEncrypt()
	log.Println("Starting Server ... \n")
	//this starts up the server
	http.ListenAndServe(":8080", context.ClearHandler(http.DefaultServeMux))
}
