// Copyright 2017 Rough Industries LLC. All rights reserved.
//controller/www/product.go: Functions and handlers for dealing with products.
package www

import (
	"html/template"
	"log"
	"model"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

//Product page handler which displays the standard product page.
func ProductHandler(w http.ResponseWriter, r *http.Request) {
	page := NewPage(r)
	var p *model.BaseProductWithBDG
	//Process Form gets an ID if it was passed
	r.ParseForm()
	if len(r.Form["ID"]) == 0 {
		page.RenderPageTemplate(w, "404")
	} else {
		//apply the template page info to the index page
		id, _ := strconv.Atoi(r.Form["ID"][0])

		p = page.Product.SelectProductByIDWithBDG(id)
		page.BaseProductWithBDG = *p
		page.RenderPageTemplate(w, "product")
	}
}

//Product Modification Form page handler which displays the Product Modification
//Form page.
func ProductModFormHandler(w http.ResponseWriter, r *http.Request) {
	page := NewPage(r)
	if page.Username != "" && page.Authenticated {
		//Process Form gets an ID if it was passed
		r.ParseForm()
		var pbt model.ProductsByTypes
		pbt = page.Product.SelectProductsByTypes(true, true, false)
		var prods []model.Product
		var prod model.Product
		prods = prod.SelectProduct()
		page.Products = prods
		page.ProductsByTypes = pbt
		if len(r.Form["ID"]) == 0 {
			//apply the template page info to the index page
			page.RenderPageTemplate(w, "productmodform")
		} else {
			id, _ := strconv.Atoi(r.Form["ID"][0])
			var in model.Product
			in.ID = id
			out := in.SelectProduct()
			page.Product = out[0]
			page.BaseProductWithBDG = *page.Product.SelectBDGByProduct()
			page.RenderPageTemplate(w, "productmodform")
		}
	} else {
		page.RenderPageTemplate(w, "404")
	}
}

//Product modification form page request handler which process the product
//modification request.  This will after verifying a valid user session,
//modify the product data based on the request.
func ProductModHandler(w http.ResponseWriter, r *http.Request) {
	page := NewPage(r)
	if page.Username != "" && page.Authenticated {
		//Get the generic data that all product mod pages need
		var pbt model.ProductsByTypes
		pbt = page.Product.SelectProductsByTypes(true, true, false)
		var prods []model.Product
		var prod model.Product
		prods = prod.SelectProduct()
		page.Products = prods
		page.ProductsByTypes = pbt
		//Validate the form input and populate the product data, this also
		//parses the form in validate
		if ValidateProduct(&page.Product, &page.BaseProductWithBDG, r) {
			//did we get an add, update, or something else request
			if r.Form["button"][0] == "add" {
				ret_id := page.Product.InsertProduct()
				//handle add of bdg if derived or group
				if page.Product.ProductGroupType == model.Derived {
					var dp model.DerivedProduct
					dp.Product.ID = ret_id
					dp.BaseProduct.ID = page.BaseProductWithBDG.BaseProduct.ID
					dp.InsertDerivedProduct()
				} else if page.Product.ProductGroupType == model.Group {
					var gp model.GroupProduct
					gp.GroupProduct.ID = ret_id
					gp.Products = page.BaseProductWithBDG.GroupProducts
					gp.InsertGroupProduct()
				}
				model.LoadMCWithProductData()
				pbt = page.Product.SelectProductsByTypes(true, true, false)
				page.ProductsByTypes = pbt
				page.Product.ID = ret_id
				outProduct := page.Product.SelectProduct()
				page.Product = outProduct[0]
				page.Messages["productModifySuccess"] = "Product modified successfully and memcache updated!"
				page.RenderPageTemplate(w, "productmodform")
				return
			} else if r.Form["button"][0] == "update" {
				rows_updated := page.Product.UpdateProduct()
				//handle add of bdg if derived or group, requires the
				if page.Product.ProductGroupType == model.Derived {
					var dp model.DerivedProduct
					dp.Product.ID = page.Product.ID
					dp.BaseProduct.ID = page.BaseProductWithBDG.BaseProduct.ID
					dp.UpdateDerivedProduct()
				} else if page.Product.ProductGroupType == model.Group {
					var gp model.GroupProduct
					gp.GroupProduct.ID = page.Product.ID
					gp.Products = page.BaseProductWithBDG.GroupProducts
					gp.UpdateGroupProduct()
				}
				log.Println("Updated " + strconv.Itoa(rows_updated) + " rows")
				model.LoadMCWithProductData()
				pbt = page.Product.SelectProductsByTypes(true, true, false)
				page.ProductsByTypes = pbt
				outProduct := page.Product.SelectProduct()
				page.Product = outProduct[0]
				page.Messages["productModifySuccess"] = "Product modified successfully and memcache updated!"
				page.RenderPageTemplate(w, "productmodform")
				return
			} else {
				//we only allow add and update right now
				page.Messages["productModifyFail"] = "Product modification failed.  You tried to perform an unknown operation!"
				page.RenderPageTemplate(w, "productmodform")
				return
			}
		} else {
			//Validation failed
			log.Println("Bad product!")
			page.RenderPageTemplate(w, "/productmodform")
			return
		}
	} else {
		page.RenderPageTemplate(w, "404")
		return
	}
}

//Products page (i.e. all the products) request handler which
//displays the all the products page.
func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	page := NewPage(r)
	var p []model.Product
	p = page.Product.SelectAllProducts()
	page.Products = p
	page.RenderPageTemplate(w, "products")
}

//Parses the form and then validates the product form request
//and populates the Product struct
func ValidateProduct(product *model.Product, bpwbd *model.BaseProductWithBDG, r *http.Request) bool {
	product.Errors = make(map[string]string)
	r.ParseForm()

	if len(r.Form["productID"]) > 0 {
		product.ID, _ = strconv.Atoi(r.Form["productID"][0])
	}
	if len(r.Form["productName"]) > 0 && strings.TrimSpace(r.Form["productName"][0]) != "" {
		product.ProductName = r.Form["productName"][0]
	} else {
		product.Errors["ProductName"] = "Please enter a valid product name"
	}
	if len(r.Form["productType"]) > 0 {
		product.ProductType.ID, _ = strconv.Atoi(r.Form["productType"][0])
		log.Println(product.ProductType.ID)
		if len(r.Form["productGroupType"+r.Form["productType"][0]]) > 0 {
			pgt, _ := r.Form["productGroupType"+r.Form["productType"][0]]
			if pgt[0] == "Base" {
				product.ProductGroupType = model.Base
			} else if pgt[0] == "Derived" {
				product.ProductGroupType = model.Derived
				if len(r.Form["derived"+r.Form["productType"][0]]) > 0 {
					der_string, _ := r.Form["derived"+r.Form["productType"][0]]
					bpwbd.BaseProduct.ID, _ = strconv.Atoi(der_string[0])
				}
			} else if pgt[0] == "Group" {
				product.ProductGroupType = model.Group
				if len(r.Form["group"+r.Form["productType"][0]]) > 0 {
					grp, _ := r.Form["group"+r.Form["productType"][0]]
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
	if len(r.Form["productDescription"]) > 0 {
		product.Description = template.HTML(r.Form["productDescription"][0])
	}
	if len(r.Form["productDetails"]) > 0 {
		product.Details = template.HTML(r.Form["productDetails"][0])
	}
	if len(r.Form["productImage"]) > 0 {
		product.ImagePath, product.Image = filepath.Split(r.Form["productImage"][0])
		product.ImagePath = strings.TrimSuffix(product.ImagePath, "/")
	}
	if len(r.Form["productImageSourceName"]) > 0 {
		product.ImageSourceName = r.Form["productImageSourceName"][0]
	}
	if len(r.Form["productImageSourceLink"]) > 0 {
		product.ImageSourceLink = r.Form["productImageSourceLink"][0]
	}
	if len(r.Form["productPreText"]) > 0 {
		product.PreText = r.Form["productPreText"][0]
	}
	if len(r.Form["productPostText"]) > 0 {
		product.PostText = r.Form["productPostText"][0]
	}
	if len(r.Form["productRating"]) > 0 {
		product.Rating, _ = strconv.Atoi(r.Form["productRating"][0])
	}
	if len(r.Form["productSourceName"]) > 0 {
		product.SourceName = r.Form["productSourceName"][0]
	}
	if len(r.Form["productSourceLink"]) > 0 {
		product.SourceLink = r.Form["productSourceLink"][0]
	}
	return len(product.Errors) == 0
}
