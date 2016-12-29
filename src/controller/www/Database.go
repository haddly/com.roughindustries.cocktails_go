//www/Database.go
package www

import (
	//"database/sql"
	"html/template"
	"log"
	"model"
	"net/http"
)

type Database struct {
}

//render the page based on the name of the file provided
func (database *Database) RenderTemplate(w http.ResponseWriter, tmpl string, d *model.Database) {
	t, err := template.ParseFiles("./view/webcontent/www/templates/"+tmpl+".html", "./view/webcontent/www/templates/ga.html")
	if err != nil {
		log.Fatal(err)
	}
	err = t.ExecuteTemplate(w, "base", d)
	if err != nil {
		log.Fatal(err)
	}
}

func (database *Database) DBValidateHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("DBValidateHandler: " + r.URL.Path[1:])

	db := model.GetCurrentDB()
	//apply the template page info to the index page
	database.RenderTemplate(w, "dbvalidate", &db)
}

func (database *Database) Init() {
	log.Println("Init in db/Database.go")
	http.HandleFunc("/db_validate", database.DBValidateHandler)
}
