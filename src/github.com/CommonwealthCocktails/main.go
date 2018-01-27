// Copyright 2017 Rough Industries LLC. All rights reserved.
//CommonwealthCocktails/main.go:
package main

import (
	"flag"
	"github.com/CommonwealthCocktails/connectors"
	"github.com/CommonwealthCocktails/controller/alexa"
	"github.com/CommonwealthCocktails/controller/www"
	"github.com/CommonwealthCocktails/utils"
	"github.com/gorilla/context"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
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
	//Add the log command line config option to print everything to stderr
	os.Args = append(os.Args, "-logtostderr=true")
	flag.Parse()

	if *passwdPtr != "" {
		// Hashing the password with the default cost of 10
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*passwdPtr), bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}
		log.Infoln(string(hashedPassword))
		os.Exit(0)
	}

	log.SetLevel(log.InfoLevel)
	Formatter := new(log.TextFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true
	utils.AddHook(true, true, true, log.AllLevels)
	log.SetFormatter(Formatter)
	log.Infoln("Info Log Level")
	log.Errorln("Error Log Level")

	log.Infoln("Reading Configuration")

	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		log.Errorf("Fatal error config file: %s \n", err)
		panic(err)
	}
	log.Infoln("Start Init")
	log.Infoln(www.State)
	log.Infoln(os.TempDir())
	//SET THESE LINES AND ADD #gitignore to the end of the line as a comment to ignore your info
	var dbaddr string
	var dbpasswd string
	var user string
	var proto string
	var port string
	var dbname string
	var dbtype connectors.DBTypesConst
	dbaddr = viper.GetString("dbaddr")
	dbpasswd = viper.GetString("dbpasswd")
	user = viper.GetString("user")
	proto = viper.GetString("proto")
	port = viper.GetString("port")
	dbname = viper.GetString("dbname")
	if viper.GetString("dbtype") == "MySQL" {
		dbtype = connectors.MySQL
	} else if viper.GetString("dbtype") == "SQLite" {
		dbtype = connectors.SQLite
	}

	var mc_server string
	mc_server = viper.GetString("mc_server")

	if dbaddr == "" || dbpasswd == "" || user == "" || proto == "" || port == "" || dbname == "" {
		log.Infoln("Not all DB parameters are set.  If your DB isn't connecting check these values.")
	}
	connectors.SetDBVars(dbaddr, dbpasswd, user, proto, port, dbname, dbtype)
	if mc_server != "" {
		connectors.SetMCVars(mc_server)
	} else {
		log.Infoln("No Memcache server set. If you want to use memcaching you will " +
			"have to set this value in main.go.")
	}
	// wanted it to be more random so i seed it time now
	rand.Seed(time.Now().UnixNano())
	// init the routing
	www.WWWRouterInit()
	alexa.AlexaRouterInit()
	www.State = www.Setup
	log.Infoln("End Init")
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

// func redirect(w http.ResponseWriter, req *http.Request) {
// 	// remove/add not default ports from req.Host
// 	target := "https://" + req.Host + req.URL.Path
// 	if len(req.URL.RawQuery) > 0 {
// 		target += "?" + req.URL.RawQuery
// 	}
// 	log.Infoln("redirect to: %s", target)
// 	http.Redirect(w, req, target, http.StatusTemporaryRedirect)
// }

func main() {
	log.Infoln(www.State)
	//print out the current directory
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Error(err)
	}
	log.Infoln(dir)
	AddLetsEncrypt()
	log.Infoln("Starting Server ... \n")
	//this starts up the server
	// redirect every http request to https
	//go http.ListenAndServe(":80", http.HandlerFunc(redirect))
	http.ListenAndServe(":8080", context.ClearHandler(http.DefaultServeMux))
}
