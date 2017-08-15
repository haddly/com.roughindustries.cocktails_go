//www/Database.go
package www

import (
	"bufio"
	"bytes"
	"connectors"
	"html/template"
	//"io"
	"io/ioutil"
	"log"
	"model"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type Database struct {
}

type Status struct {
	Status template.HTML
}

func (database *Database) DBTablesHandler(w http.ResponseWriter, r *http.Request) {
	// STANDARD HANDLER HEADER START
	// catch all errors and return 404
	defer func() {
		// recover from panic if one occured. Set err to nil otherwise.
		if rec := recover(); rec != nil {
			Error404(w, rec)
		}
	}()
	page := NewPage()
	page.Username, page.Authenticated = GetSession(r)
	// STANDARD HANLDER HEADER END
	if page.Username != "" {
		var buffer bytes.Buffer
		dat, _ := ioutil.ReadFile("sql/ccschemadump.sql")

		requests := strings.Split(string(dat), ";")
		conn, _ := connectors.GetDB()
		conn.Exec("SET FOREIGN_KEY_CHECKS=0;")
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

func (database *Database) DBDataHandler(w http.ResponseWriter, r *http.Request) {
	// STANDARD HANDLER HEADER START
	// catch all errors and return 404
	defer func() {
		// recover from panic if one occured. Set err to nil otherwise.
		if rec := recover(); rec != nil {
			Error404(w, rec)
		}
	}()
	page := NewPage()
	page.Username, page.Authenticated = GetSession(r)
	// STANDARD HANLDER HEADER END
	if page.Username != "" {
		var buffer bytes.Buffer
		buffer.WriteString("<br/><b>Data Loaded!</b> ")
		dir, _ := os.Getwd()
		dat, _ := os.Open("sql/ccdatadump.sql")
		defer dat.Close()
		scanner := bufio.NewScanner(dat)
		scanner.Split(bufio.ScanLines)

		buffer.WriteString(dir + "<br><br>")
		conn, _ := connectors.GetDB()
		conn.Exec("SET FOREIGN_KEY_CHECKS=0;")
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

func (database *Database) DBTestHandler(w http.ResponseWriter, r *http.Request) {
	// STANDARD HANDLER HEADER START
	// catch all errors and return 404
	defer func() {
		// recover from panic if one occured. Set err to nil otherwise.
		if rec := recover(); rec != nil {
			Error404(w, rec)
		}
	}()
	page := NewPage()
	page.Username, page.Authenticated = GetSession(r)
	// STANDARD HANLDER HEADER END
	if page.Username != "" {
		var buffer bytes.Buffer
		buffer.WriteString("<b>Database</b>:<br/>")
		buffer.WriteString(model.GetCurrentDB() + "<br/>")

		model.GetMetaByTypes(false, false, true)
		model.GetProductsByTypes(true, true, true)
		//apply the template page info to the index page
		statStr := buffer.String()
		page.Messages["Status"] = template.HTML(statStr)
		page.RenderPageTemplate(w, "dbindex")
	} else {
		page.RenderPageTemplate(w, "404")
	}
}

func (database *Database) Init() {
	http.HandleFunc("/db_tables", database.DBTablesHandler)
	http.HandleFunc("/db_data", database.DBDataHandler)
	http.HandleFunc("/db_test", database.DBTestHandler)
}
