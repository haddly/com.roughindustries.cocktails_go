// Copyright 2017 Rough Industries LLC. All rights reserved.
//controller/www/database.go: Functions and handlers for interacting with the
//database at the level above the individual tables.  This includes loading
//the tables and data from sql mysqldump files.
package www

import (
	"bufio"
	"bytes"
	"github.com/CommonwealthCocktails/connectors"
	"github.com/CommonwealthCocktails/model"
	"github.com/golang/glog"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

//Handler for loading the sql mysqldump file for db tables
func DBTablesHandler(w http.ResponseWriter, r *http.Request, page *page) {
	// Check for a valid user and that authentication
	var buffer bytes.Buffer
	var dat []byte
	if connectors.DBType == connectors.MySQL {
		dat, _ = ioutil.ReadFile("sql/ccschemadump.sql")
	} else if connectors.DBType == connectors.SQLite {
		dat, _ = ioutil.ReadFile("sql/ccsqlite3schema.sql")
	}
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
			glog.Infoln(string(request))
			if len(string(request)) > 0 {
				_, err := conn.Exec(string(request))
				if err != nil {
					glog.Infoln(err)
				}
			}
		}
	}
	//apply the template page info to the index page
	statStr := buffer.String()
	page.Messages["Status"] = template.HTML(statStr)
	page.RenderPageTemplate(w, r, "dbindex")

}

//Handler for loading the sql mysqldump file for db data
func DBDataHandler(w http.ResponseWriter, r *http.Request, page *page) {
	var buffer bytes.Buffer
	buffer.WriteString("<br/><b>Data Loaded!</b> ")
	dir, _ := os.Getwd()
	var dat *os.File
	if connectors.DBType == connectors.MySQL {
		dat, _ = os.Open("sql/ccdatadump.sql")
	} else if connectors.DBType == connectors.SQLite {
		dat, _ = os.Open("sql/ccsqlite3data.sql")
	}
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
		glog.Infoln(string(request))
		if len(string(request)) > 0 {
			_, err := conn.Exec(string(request))
			if err != nil {
				glog.Infoln(err)
			}
		}
	}
	//apply the template page info to the index page
	statStr := buffer.String()
	glog.Infoln(statStr)
	page.Messages["Status"] = template.HTML(statStr)
	page.RenderPageTemplate(w, r, "dbindex")
}

//Handler for running db sanity checks
func DBTestHandler(w http.ResponseWriter, r *http.Request, page *page) {
	if page.Authenticated {
		var buffer bytes.Buffer
		buffer.WriteString("<b>Database</b>:<br/>")
		buffer.WriteString(model.SelectCurrentDB() + "<br/>")

		page.Meta.SelectMetaByTypes(false, false, true)
		page.Product.SelectProductsByTypes(true, true, true)
		//apply the template page info to the index page
		statStr := buffer.String()
		page.Messages["Status"] = template.HTML(statStr)
		page.RenderPageTemplate(w, r, "dbindex")
	} else {
		http.Redirect(w, r, "/", 302)
	}
}
