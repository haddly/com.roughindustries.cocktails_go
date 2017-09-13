// Copyright 2017 Rough Industries LLC. All rights reserved.
//controller/www/database.go: Functions and handlers for interacting with the
//database at the level above the individual tables.  This includes loading
//the tables and data from sql mysqldump files.
package www

import (
	"bufio"
	"bytes"
	"connectors"
	"html/template"
	"io/ioutil"
	"log"
	"model"
	"net/http"
	"os"
	"regexp"
	"strings"
)

//Handler for loading the sql mysqldump file for db tables
func DBTablesHandler(w http.ResponseWriter, r *http.Request) {
	page := NewPage(r)
	// Check for a valid user and that authentication
	if page.Username != "" && page.Authenticated {
		var buffer bytes.Buffer
		dat, _ := ioutil.ReadFile("sql/ccschemadump.sql")
		//dat, _ := ioutil.ReadFile("sql/ccsqlite3schema.sql")
		requests := strings.Split(string(dat), ";")
		conn, _ := connectors.GetDB()
		//disable foreign key contraint sense I don't know the order we add
		//the tables
		conn.Exec("SET FOREIGN_KEY_CHECKS=0;")
		//parse the file and run only the slq commands
		for _, request := range requests {
			r, _ := regexp.Compile("(.*/*!.*)")
			if !r.MatchString(string(request)) {
				buffer.WriteString(string(request) + ";<br><br>")
				log.Println(string(request))
				if len(string(request)) > 0 {
					_, err := conn.Exec(string(request))
					if err != nil {
						log.Println(err)
					}
				}
			}
		}
		//apply the template page info to the index page
		statStr := buffer.String()
		page.Messages["Status"] = template.HTML(statStr)
		page.RenderPageTemplate(w, "dbindex")
	} else {
		page.RenderPageTemplate(w, "404")
	}
}

//Handler for loading the sql mysqldump file for db data
func DBDataHandler(w http.ResponseWriter, r *http.Request) {
	page := NewPage(r)
	if page.Username != "" && page.Authenticated {
		var buffer bytes.Buffer
		buffer.WriteString("<br/><b>Data Loaded!</b> ")
		dir, _ := os.Getwd()
		dat, _ := os.Open("sql/ccdatadump.sql")
		//dat, _ := os.Open("sql/ccsqlite3data.sql")
		defer dat.Close()
		scanner := bufio.NewScanner(dat)
		scanner.Split(bufio.ScanLines)
		buffer.WriteString(dir + "<br><br>")
		conn, _ := connectors.GetDB()
		//disable foreign key contraint sense I don't know the order we add
		//the data
		conn.Exec("SET FOREIGN_KEY_CHECKS=0;")
		//parse the file and run only the slq commands
		for scanner.Scan() {
			request := scanner.Text()
			buffer.WriteString(string(request) + ";<br><br>")
			log.Println(string(request))
			if len(string(request)) > 0 {
				_, err := conn.Exec(string(request))
				if err != nil {
					log.Println(err)
				}
			}
		}
		//apply the template page info to the index page
		statStr := buffer.String()
		log.Println(statStr)
		page.Messages["Status"] = template.HTML(statStr)
		page.RenderPageTemplate(w, "dbindex")
	} else {
		page.RenderPageTemplate(w, "404")
	}
}

//Handler for running db sanity checks
func DBTestHandler(w http.ResponseWriter, r *http.Request) {
	page := NewPage(r)
	if page.Username != "" && page.Authenticated {
		var buffer bytes.Buffer
		buffer.WriteString("<b>Database</b>:<br/>")
		buffer.WriteString(model.GetCurrentDB() + "<br/>")

		model.SelectMetaByTypes(false, false, true)
		model.SelectProductsByTypes(true, true, true)
		//apply the template page info to the index page
		statStr := buffer.String()
		page.Messages["Status"] = template.HTML(statStr)
		page.RenderPageTemplate(w, "dbindex")
	} else {
		page.RenderPageTemplate(w, "404")
	}
}
