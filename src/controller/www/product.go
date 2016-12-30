//www/product.go
package www

import (
	"html/template"
	"log"
	"model"
	"net/http"
)

type Product struct {
}

//render the page based on the name of the file provided
func (product *Product) RenderTemplate(w http.ResponseWriter, tmpl string, p *model.Product) {
	t, err := template.ParseFiles("./view/webcontent/www/templates/"+tmpl+".html", "./view/webcontent/www/templates/ga.html", "./view/webcontent/www/templates/header.html", "./view/webcontent/www/templates/footer.html")
	if err != nil {
		log.Fatal(err)
	}
	err = t.ExecuteTemplate(w, "base", p)
	if err != nil {
		log.Fatal(err)
	}
}

func (product *Product) ProductHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Product: " + r.URL.Path[1:])

	//apply the template page info to the index page
	p := &model.Products[0]
	product.RenderTemplate(w, "product", p)
}

func (product *Product) Init() {
	log.Println("Init in www/product.go")
	http.HandleFunc("/product", product.ProductHandler)
}
