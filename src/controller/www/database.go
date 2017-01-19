//www/Database.go
package www

import (
	//"database/sql"
	"bytes"
	"db"
	"html/template"
	"log"
	"model"
	"net/http"
)

type Database struct {
}

type Status struct {
	Status template.HTML
}

//render the page based on the name of the file provided
func (database *Database) RenderTemplate(w http.ResponseWriter, tmpl string, s *Status) {
	t, err := template.ParseFiles("./view/webcontent/www/templates/"+tmpl+".html", "./view/webcontent/www/templates/ga.html")
	if err != nil {
		log.Fatal(err)
	}
	err = t.ExecuteTemplate(w, "base", s)
	if err != nil {
		log.Fatal(err)
	}
}

func (database *Database) DBValidateHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("DBValidateHandler: " + r.URL.Path[1:])

	var buffer bytes.Buffer
	buffer.WriteString("<b>Database</b>:<br/>")
	buffer.WriteString(model.GetCurrentDB() + "<br/>")
	buffer.WriteString("<br/><b>Tables</b>: ")
	model.InitCocktailTables()
	model.InitPostTables()
	model.InitMetaTables()
	conn, _ := db.GetDB()
	rows, _ := conn.Query("SHOW TABLES;")
	for rows.Next() {
		var temp string
		rows.Scan(&temp)
		buffer.WriteString("<br>" + temp)
	}
	//apply the template page info to the index page
	statStr := buffer.String()
	status := Status{template.HTML(statStr)}
	database.RenderTemplate(w, "dbindex", &status)
}

func (database *Database) Init() {
	http.HandleFunc("/db_validate", database.DBValidateHandler)
}
