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
func (product *Product) RenderBaseProductWithBDTemplate(w http.ResponseWriter, tmpl string, p *model.BaseProductWithBD) {
	t, err := parseTempFiles(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	err = t.ExecuteTemplate(w, "base", p)
	if err != nil {
		log.Fatal(err)
	}
}

//render the page based on the name of the file provided
func (product *Product) RenderProductsTemplate(w http.ResponseWriter, tmpl string, p []model.Product) {
	t, err := parseTempFiles(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	err = t.ExecuteTemplate(w, "base", p)
	if err != nil {
		log.Fatal(err)
	}
}

func (product *Product) ProductHandler(w http.ResponseWriter, r *http.Request) {
	var p *model.BaseProductWithBD
	u, err := url.Parse(r.URL.String())
	if err != nil {
		product.RenderBaseProductWithBDTemplate(w, "404", p)
	}
	//log.Println("Product: " + r.URL.Path[1:])
	m, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		product.RenderBaseProductWithBDTemplate(w, "404", p)
	}
	if len(m["ID"]) == 0 {
		product.RenderBaseProductWithBDTemplate(w, "404", p)
	} else {
		//log.Println("ID: " + m["ID"][0])

		//apply the template page info to the index page
		id, _ := strconv.Atoi(m["ID"][0])
		if len(model.Products) <= id-1 {
			product.RenderBaseProductWithBDTemplate(w, "404", p)
		} else {
			p = model.GetBaseProductByIDWithBD(id)
			product.RenderBaseProductWithBDTemplate(w, "product", p)
		}
	}
}

func (product *Product) ProductsHandler(w http.ResponseWriter, r *http.Request) {
	var p []model.Product

	p = model.GetProducts()
	product.RenderProductsTemplate(w, "products", p)
}

func (product *Product) Init() {
	log.Println("Init in www/product.go")
	http.HandleFunc("/product", product.ProductHandler)
	http.HandleFunc("/product/", product.ProductHandler)
	http.HandleFunc("/products", product.ProductsHandler)
}
