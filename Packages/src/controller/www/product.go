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

func (product *Product) ProductHandler(w http.ResponseWriter, r *http.Request) {
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
	var p *model.BaseProductWithBD
	u, err := url.Parse(r.URL.String())
	if err != nil {
		page.RenderPageTemplate(w, "404")
	}
	//log.Println("Product: " + r.URL.Path[1:])
	m, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		page.RenderPageTemplate(w, "404")
	}
	if len(m["ID"]) == 0 {
		page.RenderPageTemplate(w, "404")
	} else {
		//apply the template page info to the index page
		id, _ := strconv.Atoi(m["ID"][0])
		if len(model.Products) <= id-1 {
			page.RenderPageTemplate(w, "404")
		} else {
			p = model.GetBaseProductByIDWithBD(id)
			page.BaseProductWithBD = *p
			page.RenderPageTemplate(w, "product")
		}
	}
}

func (product *Product) ProductAddFormHandler(w http.ResponseWriter, r *http.Request) {
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
		var pbt model.ProductsByTypes
		pbt = model.GetProductsByTypes(true, true, false)
		var prods []model.Product
		var prod model.Product
		prods = model.SelectProduct(prod)
		page.Products = prods
		page.ProductsByTypes = pbt
		//apply the template page info to the index page
		page.RenderPageTemplate(w, "productaddform")
	} else {
		page.RenderPageTemplate(w, "404")
	}
}

//handle / requests to the server
func (product *Product) ProductAddHandler(w http.ResponseWriter, r *http.Request) {
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
	u, err := url.Parse(r.URL.String())
	log.Println(u)
	if err != nil {
		page.RenderPageTemplate(w, "404")
	}
	m, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		page.RenderPageTemplate(w, "404")
	}
	log.Println(m)
	log.Println(r.FormValue("productType"))
	log.Println(m["productType"][0])
	log.Println(strconv.Itoa(model.ProductTypeStringToInt(m["productType"][0])))
	pgt := "productGroupType" + m["productType"][0]
	log.Println(pgt)
	log.Println(m[pgt][0])
	log.Println(strconv.Itoa(model.GroupTypeStringToInt(m[pgt][0])))
	var prod model.Product
	prod.ProductName = m["productName"][0]
	page.Errors["productName"] = "Please enter a valid name"
	prod.ProductType.ID = model.ProductTypeStringToInt(m["productType"][0])
	prod.ProductType.ProductTypeName = m["productType"][0]
	prod.ProductGroupType = model.GroupType(model.GroupTypeStringToInt(m[pgt][0]))
	prod.Description = template.HTML(m["productDescription"][0])
	prod.Details = template.HTML(m["productDetails"][0])
	log.Println(prod)
	//id := model.ProcessProduct(prod)
	//log.Println(strconv.Itoa(id))
	//prod.ID = id
	var derProd model.DerivedProduct
	derProd.Product = prod
	if m[pgt][0] == "Derived" {
		d := "derived" + m["productType"][0]
		log.Println(d)
		log.Println(strings.TrimRight(m[d][0], ","))
		derProd.BaseProduct.ID, _ = strconv.Atoi(strings.TrimRight(m[d][0], ","))
		//model.ProcessDerivedProduct(derProd)
	}
	var pbt model.ProductsByTypes
	pbt = model.GetProductsByTypes(true, true, false)
	page.ProductsByTypes = pbt
	page.Product = prod
	page.RenderPageTemplate(w, "productaddform")
}

func (product *Product) ProductsHandler(w http.ResponseWriter, r *http.Request) {
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
	var p []model.Product
	p = model.GetProducts()
	page.Products = p
	page.RenderPageTemplate(w, "products")
}

func (product *Product) Init() {
	log.Println("Init in www/product.go")
	http.HandleFunc("/product", product.ProductHandler)
	http.HandleFunc("/product/", product.ProductHandler)
	http.HandleFunc("/products", product.ProductsHandler)
	http.HandleFunc("/productAddForm", product.ProductAddFormHandler)
	http.HandleFunc("/productAdd", product.ProductAddHandler)
}
