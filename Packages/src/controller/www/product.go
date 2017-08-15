//www/product.go
package www

import (
	"html/template"
	"log"
	"model"
	"net/http"
	"net/url"
	"path/filepath"
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
	var p *model.BaseProductWithBDG
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

		p = model.GetProductByIDWithBDG(id)
		page.BaseProductWithBDG = *p
		page.RenderPageTemplate(w, "product")
	}
}

func (product *Product) ProductModFormHandler(w http.ResponseWriter, r *http.Request) {
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
		u, err := url.Parse(r.URL.String())
		log.Println(u)
		if err != nil {
			page.RenderPageTemplate(w, "404")
		}
		m, err := url.ParseQuery(u.RawQuery)
		if err != nil {
			page.RenderPageTemplate(w, "404")
		}
		var pbt model.ProductsByTypes
		pbt = model.GetProductsByTypes(true, true, false)
		var prods []model.Product
		var prod model.Product
		prods = model.SelectProduct(prod)
		page.Products = prods
		page.ProductsByTypes = pbt
		if len(m["ID"]) == 0 {
			//apply the template page info to the index page
			page.RenderPageTemplate(w, "productmodform")
		} else {
			id, _ := strconv.Atoi(m["ID"][0])
			var in model.Product
			in.ID = id
			out := model.SelectProduct(in)
			page.Product = out[0]
			page.BaseProductWithBDG = *model.GetBDGByProduct(out[0])
			page.RenderPageTemplate(w, "productmodform")
		}
	} else {
		page.RenderPageTemplate(w, "404")
	}
}

//handle / requests to the server
func (product *Product) ProductModHandler(w http.ResponseWriter, r *http.Request) {
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
		u, err := url.Parse(r.URL.String())
		log.Println(u)
		if err != nil {
			page.RenderPageTemplate(w, "404")
			return
		}
		m, err := url.ParseQuery(u.RawQuery)
		if err != nil {
			page.RenderPageTemplate(w, "404")
			return
		}
		log.Println(m)
		var pbt model.ProductsByTypes
		pbt = model.GetProductsByTypes(true, true, false)
		var prods []model.Product
		var prod model.Product
		prods = model.SelectProduct(prod)
		page.Products = prods
		page.ProductsByTypes = pbt
		if ValidateProduct(&page.Product, &page.BaseProductWithBDG, m) {
			if m["button"][0] == "add" {
				ret_id := model.InsertProduct(page.Product)
				//handle add of bdg if derived or group
				if page.Product.ProductGroupType == model.Derived {
					var dp model.DerivedProduct
					dp.Product.ID = ret_id
					dp.BaseProduct.ID = page.BaseProductWithBDG.BaseProduct.ID
					model.InsertDerivedProduct(dp)
				} else if page.Product.ProductGroupType == model.Group {
					var gp model.GroupProduct
					gp.GroupProduct.ID = ret_id
					gp.Products = page.BaseProductWithBDG.GroupProducts
					model.InsertGroupProduct(gp)
				}
				model.LoadMCWithProductData()
				pbt = model.GetProductsByTypes(true, true, false)
				page.ProductsByTypes = pbt
				page.Product.ID = ret_id
				outProduct := model.SelectProduct(page.Product)
				page.Product = outProduct[0]
				page.Messages["productModifySuccess"] = "Product modified successfully and memcache updated!"
				page.RenderPageTemplate(w, "productmodform")
				return
			} else if m["button"][0] == "update" {
				rows_updated := model.UpdateProduct(page.Product)
				//handle add of bdg if derived or group, requires the
				if page.Product.ProductGroupType == model.Derived {
					var dp model.DerivedProduct
					dp.Product.ID = page.Product.ID
					dp.BaseProduct.ID = page.BaseProductWithBDG.BaseProduct.ID
					model.UpdateDerivedProduct(dp)
				} else if page.Product.ProductGroupType == model.Group {
					var gp model.GroupProduct
					gp.GroupProduct.ID = page.Product.ID
					gp.Products = page.BaseProductWithBDG.GroupProducts
					model.UpdateGroupProduct(gp)
				}
				log.Println("Updated " + strconv.Itoa(rows_updated) + " rows")
				model.LoadMCWithProductData()
				pbt = model.GetProductsByTypes(true, true, false)
				page.ProductsByTypes = pbt
				outProduct := model.SelectProduct(page.Product)
				page.Product = outProduct[0]
				page.Messages["productModifySuccess"] = "Product modified successfully and memcache updated!"
				page.RenderPageTemplate(w, "productmodform")
				return
			} else {
				page.Messages["productModifyFail"] = "Product modification failed.  You tried to perform an unknown operation!"
				page.RenderPageTemplate(w, "productmodform")
				return
			}
		} else {
			log.Println("Bad product!")
			page.RenderPageTemplate(w, "/productmodform")
			return
		}
	} else {
		page.RenderPageTemplate(w, "404")
		return
	}
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
	p = model.SelectAllProducts()
	page.Products = p
	page.RenderPageTemplate(w, "products")
}

func ValidateProduct(product *model.Product, bpwbd *model.BaseProductWithBDG, m map[string][]string) bool {
	product.Errors = make(map[string]string)

	if len(m["productID"]) > 0 {
		product.ID, _ = strconv.Atoi(m["productID"][0])
	}
	if len(m["productName"]) > 0 && strings.TrimSpace(m["productName"][0]) != "" {
		product.ProductName = m["productName"][0]
	} else {
		product.Errors["ProductName"] = "Please enter a valid product name"
	}
	if len(m["productType"]) > 0 {
		product.ProductType.ID, _ = strconv.Atoi(m["productType"][0])
		log.Println(product.ProductType.ID)
		if len(m["productGroupType"+m["productType"][0]]) > 0 {
			pgt, _ := m["productGroupType"+m["productType"][0]]
			if pgt[0] == "Base" {
				product.ProductGroupType = model.Base
			} else if pgt[0] == "Derived" {
				product.ProductGroupType = model.Derived
				if len(m["derived"+m["productType"][0]]) > 0 {
					der_string, _ := m["derived"+m["productType"][0]]
					bpwbd.BaseProduct.ID, _ = strconv.Atoi(der_string[0])
				}
			} else if pgt[0] == "Group" {
				product.ProductGroupType = model.Group
				if len(m["group"+m["productType"][0]]) > 0 {
					grp, _ := m["group"+m["productType"][0]]
					var dgp []model.Product
					for i := range grp {
						var inProduct model.Product
						inProduct.ID, _ = strconv.Atoi(grp[i])
						dgp = append(dgp, inProduct)
					}
					bpwbd.GroupProducts = dgp
				}
			}
		}
	} else {
		product.Errors["ProductType"] = "Please select a valid product type"
	}
	if len(m["productDescription"]) > 0 {
		product.Description = template.HTML(m["productDescription"][0])
	}
	if len(m["productDetails"]) > 0 {
		product.Details = template.HTML(m["productDetails"][0])
	}
	if len(m["productImage"]) > 0 {
		product.ImagePath, product.Image = filepath.Split(m["productImage"][0])
		product.ImagePath = strings.TrimSuffix(product.ImagePath, "/")
	}
	if len(m["productImageSourceName"]) > 0 {
		product.ImageSourceName = m["productImageSourceName"][0]
	}
	if len(m["productImageSourceLink"]) > 0 {
		product.ImageSourceLink = m["productImageSourceLink"][0]
	}
	if len(m["productPreText"]) > 0 {
		product.PreText = m["productPreText"][0]
	}
	if len(m["productPostText"]) > 0 {
		product.PostText = m["productPostText"][0]
	}
	if len(m["productRating"]) > 0 {
		product.Rating, _ = strconv.Atoi(m["productRating"][0])
	}
	if len(m["productSourceName"]) > 0 {
		product.SourceName = m["productSourceName"][0]
	}
	if len(m["productSourceLink"]) > 0 {
		product.SourceLink = m["productSourceLink"][0]
	}
	return len(product.Errors) == 0
}

func (product *Product) Init() {
	log.Println("Init in www/product.go")
	http.HandleFunc("/product", product.ProductHandler)
	http.HandleFunc("/product/", product.ProductHandler)
	http.HandleFunc("/products", product.ProductsHandler)
	http.HandleFunc("/productModForm", product.ProductModFormHandler)
	http.HandleFunc("/productMod", product.ProductModHandler)
}
