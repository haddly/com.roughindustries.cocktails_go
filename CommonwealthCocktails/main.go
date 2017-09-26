// Copyright 2017 Rough Industries LLC. All rights reserved.
//CommonwealthCocktails/main.go:
package main

import (
	"connectors"
	"controller/alexa"
	"controller/www"
	"flag"
	"github.com/gorilla/context"
	"golang.org/x/crypto/bcrypt"
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
	//Utility for getting a hashed password.  I use this to setup the initial
	//admin account password sense I use bcrypt to check password hashes.
	//Set it up and then start the server if you want ot use a default admin
	//account.
	passwdPtr := flag.String("password", "", "Gives a password hash, but doesn't start the server. Don't forget special characters need to be escaped on the command-line, i.e. ! ? $ % $ # & * ( ) blank tab | ' ; \" < > \\ ~ ` [ ] { }")
	flag.Parse()
	if *passwdPtr != "" {
		// Hashing the password with the default cost of 10
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*passwdPtr), bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}
		log.Println(string(hashedPassword))
		os.Exit(0)
	}
	log.Println("Start Init")
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

	if dbaddr == "" || dbpasswd == "" || user == "" || proto == "" || port == "" || dbname == "" {
		log.Println("Not all DB parameters are set.  If your DB isn't connecting check these values.")
	}
	connectors.SetDBVars(dbaddr, dbpasswd, user, proto, port, dbname, dbtype)
	if mc_server != "" {
		connectors.SetMCVars(mc_server)
	} else {
		log.Println("No Memcache server set. If you want to use memcaching you will " +
			"have to set this value in main.go.")
	}
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
