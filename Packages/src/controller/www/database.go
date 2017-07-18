//www/Database.go
package www

import (
	"bytes"
	"connectors"
	"html/template"
	"model"
	"net/http"
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
		model.InitUserTables()

		conn, _ := connectors.GetDB()
		rows, _ := conn.Query("SHOW TABLES;")
		for rows.Next() {
			var temp string
			rows.Scan(&temp)
			buffer.WriteString("<br>" + temp)
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
		model.ProcessUsers()

		buffer.WriteString("<br/><b>Data Loaded!</b> ")
		//apply the template page info to the index page
		statStr := buffer.String()
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
