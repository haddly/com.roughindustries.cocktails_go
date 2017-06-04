//www/Database.go
package www

import (
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

func (database *Database) DBTablesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("DBTablesHandler: " + r.URL.Path[1:])

	var buffer bytes.Buffer
	buffer.WriteString("<b>Database</b>:<br/>")
	buffer.WriteString(model.GetCurrentDB() + "<br/>")
	buffer.WriteString("<br/><b>Tables</b>: ")
	model.InitProductTables()
	model.InitCocktailTables()
	model.InitPostTables()
	model.InitMetaTables()
	model.InitRecipeTables()
	model.InitAdvertisementTables()
	model.InitCocktailReferences()
	model.InitAdvertisementReferences()
	model.InitMetaReferences()
	model.InitProductReferences()
	model.InitRecipeReferences()

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

func (database *Database) DBDataHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("DBValidateHandler: " + r.URL.Path[1:])

	var buffer bytes.Buffer
	buffer.WriteString("<b>Database</b>:<br/>")
	buffer.WriteString(model.GetCurrentDB() + "<br/>")

	model.ProcessPosts()
	model.ProcessMetaTypes()
	model.ProcessProducts()
	model.ProcessMetas()
	model.ProcessCocktails()
	model.ProcessRecipes()
	model.ProcessDerivedProducts()
	model.ProcessProductGroups()

	buffer.WriteString("<br/><b>Data Loaded!</b> ")

	//apply the template page info to the index page
	statStr := buffer.String()
	status := Status{template.HTML(statStr)}
	database.RenderTemplate(w, "dbindex", &status)
}

func (database *Database) DBTestHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("DBValidateHandler: " + r.URL.Path[1:])

	var buffer bytes.Buffer
	buffer.WriteString("<b>Database</b>:<br/>")
	buffer.WriteString(model.GetCurrentDB() + "<br/>")

	model.GetMetaByTypes(false, false)
	model.GetBaseProductByTypes()

	//apply the template page info to the index page
	statStr := buffer.String()
	status := Status{template.HTML(statStr)}
	database.RenderTemplate(w, "dbindex", &status)
}

func (database *Database) Init() {
	http.HandleFunc("/db_tables", database.DBTablesHandler)
	http.HandleFunc("/db_data", database.DBDataHandler)
	http.HandleFunc("/db_test", database.DBTestHandler)
}
