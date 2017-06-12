//www/product.go
package www

import (
	"log"
	"model"
	"net/http"
	"net/url"
	"strconv"
)

type Product struct {
}

//render the page based on the name of the file provided
func (product *Product) RenderPageTemplate(w http.ResponseWriter, tmpl string, page *Page) {
	t, err := parseTempFiles(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	err = t.ExecuteTemplate(w, "base", page)
	if err != nil {
		log.Fatal(err)
	}
}

func (product *Product) ProductHandler(w http.ResponseWriter, r *http.Request) {
	var p *model.BaseProductWithBD
	var page Page
	page.Username = GetUserName(r)
	u, err := url.Parse(r.URL.String())
	if err != nil {
		product.RenderPageTemplate(w, "404", &page)
	}
	//log.Println("Product: " + r.URL.Path[1:])
	m, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		product.RenderPageTemplate(w, "404", &page)
	}
	if len(m["ID"]) == 0 {
		product.RenderPageTemplate(w, "404", &page)
	} else {
		//log.Println("ID: " + m["ID"][0])

		//apply the template page info to the index page
		id, _ := strconv.Atoi(m["ID"][0])
		if len(model.Products) <= id-1 {
			product.RenderPageTemplate(w, "404", &page)
		} else {
			p = model.GetBaseProductByIDWithBD(id)
			page.BaseProductWithBD = *p
			product.RenderPageTemplate(w, "product", &page)
		}
	}
}

func (product *Product) ProductsHandler(w http.ResponseWriter, r *http.Request) {
	var p []model.Product
	var page Page
	page.Username = GetUserName(r)
	p = model.GetProducts()
	page.Products = p
	product.RenderPageTemplate(w, "products", &page)
}

func (product *Product) Init() {
	log.Println("Init in www/product.go")
	http.HandleFunc("/product", product.ProductHandler)
	http.HandleFunc("/product/", product.ProductHandler)
	http.HandleFunc("/products", product.ProductsHandler)
}
