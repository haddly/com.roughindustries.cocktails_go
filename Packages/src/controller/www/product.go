//www/product.go
package www

import (
	"html/template"
	"log"
	"model"
	"net/http"
	"net/url"
	"strconv"
	"strings"
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

func (product *Product) ProductAddFormHandler(w http.ResponseWriter, r *http.Request) {
	userName := GetUserName(r)
	if userName != "" {
		var page Page
		page.Username = GetUserName(r)
		var pbt model.ProductsByTypes
		pbt = model.GetProductsByTypes(true, true)
		page.ProductsByTypes = pbt
		//apply the template page info to the index page
		product.RenderPageTemplate(w, "productaddform", &page)
	} else {
		product.RenderPageTemplate(w, "404", nil)
	}
}

//handle / requests to the server
func (product *Product) ProductAddHandler(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.String())
	log.Println(u)
	if err != nil {
		product.RenderPageTemplate(w, "404", nil)
	}
	m, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		product.RenderPageTemplate(w, "404", nil)
	}
	log.Println(m)
	log.Println(m["productType"][0])
	log.Println(strconv.Itoa(model.ProductTypeStringToInt(m["productType"][0])))
	pgt := "productGroupType" + m["productType"][0]
	log.Println(pgt)
	log.Println(m[pgt][0])
	log.Println(strconv.Itoa(model.GroupTypeStringToInt(m[pgt][0])))
	var prod model.Product
	prod.ProductName = m["productName"][0]
	prod.ProductType.ID = model.ProductTypeStringToInt(m["productType"][0])
	prod.ProductGroupType = model.GroupType(model.GroupTypeStringToInt(m[pgt][0]))
	prod.Description = template.HTML(m["productDescription"][0])
	prod.Details = template.HTML(m["productDetails"][0])
	log.Println(prod)
	id := model.ProcessProduct(prod)
	log.Println(strconv.Itoa(id))
	prod.ID = id
	var derProd model.DerivedProduct
	derProd.Product = prod
	if m[pgt][0] == "Derived" {
		d := "derived" + m["productType"][0]
		log.Println(d)
		log.Println(strings.TrimRight(m[d][0], ","))
		derProd.BaseProduct.ID, _ = strconv.Atoi(strings.TrimRight(m[d][0], ","))
		model.ProcessDerivedProduct(derProd)
	}
	product.RenderPageTemplate(w, "404", nil)
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
	http.HandleFunc("/productAddForm", product.ProductAddFormHandler)
	http.HandleFunc("/productAdd", product.ProductAddHandler)
}
