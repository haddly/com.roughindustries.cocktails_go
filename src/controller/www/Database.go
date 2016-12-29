//www/Database.go
package www

import (
	//"database/sql"
	"bytes"
	"html/template"
	"log"
	"model"
	"net/http"
)

type Database struct {
}

type Status struct {
	Status string
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
	buffer.WriteString(model.GetCurrentDB() + "/<br/>")
	buffer.WriteString(model.InitCocktailTable() + "/<br/>")
	//apply the template page info to the index page
	status := Status{buffer.String()}
	database.RenderTemplate(w, "dbindex", &status)
}

func (database *Database) Init() {
	log.Println("Init in db/Database.go")
	http.HandleFunc("/db_validate", database.DBValidateHandler)
}
