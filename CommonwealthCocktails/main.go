//main/main
package main

import (
	"controller/alexa"
	"controller/www"
	"db"
	"flag"
	"log"
	"math/rand"
	"model"
	"net/http"
	"os"
	//"path/filepath"
	"time"
)

//The controllers
var d = www.Database{}
var c = www.Cocktail{}
var p = www.Product{}
var s = www.Search{}
var post = www.Post{}
var a = alexa.Hello{}

func init() {

	var datasource = model.DSTtoi(*flag.String("DataSource", "Internal", "DataSource Type"))
	if os.Getenv("DataSource") != "" {
		datasource = model.DSTtoi(os.Getenv("DataSource"))
	}

	if datasource == model.DB {
		log.Println("In DB setup")
		var dbaddr = *flag.String("DBADDR", "0.0.0.0", "Database IP Address")
		if os.Getenv("DBADDR") != "" {
			dbaddr = os.Getenv("DBADDR")
		}
		var dbpasswd = *flag.String("DBPASSWD", "password", "Database password")
		if os.Getenv("DBPASSWD") != "" {
			dbpasswd = os.Getenv("DBPASSWD")
		}
		var user = *flag.String("DBUSERNAME", "user", "Database username to login as")
		if os.Getenv("DBUSERNAME") != "" {
			user = os.Getenv("DBUSERNAME")
		}
		var proto = *flag.String("DBPROTOCOL", "tcp", "Database protocol to use to connect over")
		if os.Getenv("DBPROTOCOL") != "" {
			proto = os.Getenv("DBPROTOCOL")
		}
		var port = *flag.String("DBPORT", "3306", "Database port to connect over")
		if os.Getenv("DBPORT") != "" {
			port = os.Getenv("DBPORT")
		}
		var dbname = *flag.String("DBNAME", "db", "Database name")
		if os.Getenv("DBNAME") != "" {
			dbname = os.Getenv("DBNAME")
		}
		db.SetDBVars(dbaddr, dbpasswd, user, proto, port, dbname)
	} else {
		log.Println(datasource)
	}

	model.SetDataSourceType(datasource)
	flag.Parse()
	// wanted it to be more random so i seed it time now
	rand.Seed(time.Now().UnixNano())
	d.Init()
	c.Init()
	p.Init()
	s.Init()
	a.Init()
	post.Init()
	log.Println("End Init")
}

func main() {
	// log.Println("Starting ... \n")
	// //print out the current directory
	// dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(dir)

	// // Mandatory root-based resources and redirects for other resources
	// // This is handled in the app.yaml for google cloud platform deployments
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
