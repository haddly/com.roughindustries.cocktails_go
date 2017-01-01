//www/product.go
package www

import (
	"html/template"
	"log"
	"model"
	"net/http"
	"net/url"
	"strconv"
)

type Product struct {
}

//render the page based on the name of the file provided
func (product *Product) RenderTemplate(w http.ResponseWriter, tmpl string, p *model.BaseProductWithBDG) {
	t, err := template.ParseFiles("./view/webcontent/www/templates/"+tmpl+".html", "./view/webcontent/www/templates/ga.html", "./view/webcontent/www/templates/navbar.html", "./view/webcontent/www/templates/footer.html")
	if err != nil {
		log.Fatal(err)
	}
	err = t.ExecuteTemplate(w, "base", p)
	if err != nil {
		log.Fatal(err)
	}
}

func (product *Product) ProductHandler(w http.ResponseWriter, r *http.Request) {
	var p *model.BaseProductWithBDG
	u, err := url.Parse(r.URL.String())
	if err != nil {
		product.RenderTemplate(w, "404", p)
	}
	//log.Println("Product: " + r.URL.Path[1:])
	m, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		product.RenderTemplate(w, "404", p)
	}
	if len(m["ID"]) == 0 {
		product.RenderTemplate(w, "404", p)
	} else {
		//log.Println("ID: " + m["ID"][0])

		//apply the template page info to the index page
		id, _ := strconv.Atoi(m["ID"][0])
		if len(model.Products) <= id-1 {
			product.RenderTemplate(w, "404", p)
		} else {
			p = model.GetBaseProductWithBDG(id)
			product.RenderTemplate(w, "product", p)
		}
	}
}

func (product *Product) Init() {
	log.Println("Init in www/product.go")
	http.HandleFunc("/product", product.ProductHandler)
	http.HandleFunc("/product/", product.ProductHandler)
}
