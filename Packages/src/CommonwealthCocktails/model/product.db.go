// Copyright 2017 Rough Industries LLC. All rights reserved.
//model/product.db.go:package model
package model

import (
	"bytes"
	"CommonwealthCocktails/connectors"
	"encoding/gob"
	"github.com/bradfitz/gomemcache/memcache"
	"html"
	"html/template"
	"github.com/golang/glog"
	"strconv"
	"strings"
)

//CREATE, UPDATE, DELETE
//Insert a product record into the database
func (product *Product) InsertProduct() int {
	product.ID = 0
	return product.processProduct()
}

//Update a product record in the database based on ID
func (product *Product) UpdateProduct() int {
	return product.processProduct()
}

//Process an insert or an update
func (product *Product) processProduct() int {
	conn, _ := connectors.GetDB()
	var args []interface{} //arguments for variables in the data struct
	var buffer bytes.Buffer
	if product.ID == 0 {
		buffer.WriteString("INSERT INTO `product` SET ")
	} else {
		buffer.WriteString("UPDATE `product` SET ")
	}
	if product.ProductName != "" {
		buffer.WriteString("`productName`=?,")
		args = append(args, html.EscapeString(string(product.ProductName)))
	}
	buffer.WriteString(" `productType`=?,")
	args = append(args, strconv.Itoa(int(product.ProductType.ID)))
	buffer.WriteString(" `productGroupType`=?,")
	args = append(args, strconv.Itoa(int(product.ProductGroupType)))
	if product.Description != "" {
		buffer.WriteString("`productDescription`=?,")
		args = append(args, html.EscapeString(string(product.Description)))
	}
	if product.Details != "" {
		buffer.WriteString("`productDetails`=?,")
		args = append(args, html.EscapeString(string(product.Details)))
	}
	if product.PreText != "" {
		buffer.WriteString("`productPreText`=?,")
		args = append(args, html.EscapeString(string(product.PreText)))
	}
	if product.PostText != "" {
		buffer.WriteString("`productPostText`=?,")
		args = append(args, html.EscapeString(string(product.PostText)))
	}
	if product.Rating != 0 {
		buffer.WriteString(" `productRating`=?,")
		args = append(args, strconv.Itoa(product.Rating))
	}
	if product.ImagePath != "" {
		buffer.WriteString("`productImagePath`=?,")
		args = append(args, product.ImagePath)
	}
	if product.Image != "" {
		buffer.WriteString("`productImage`=?,")
		args = append(args, product.Image)
	}
	if product.ImageSourceName != "" {
		buffer.WriteString("`productImageSourceName`=?,")
		args = append(args, html.EscapeString(string(product.ImageSourceName)))
	}
	if product.ImageSourceLink != "" {
		buffer.WriteString("`productImageSourceLink`=?,")
		args = append(args, product.ImageSourceLink)
	}
	if product.SourceName != "" {
		buffer.WriteString("`productSourceName`=?,")
		args = append(args, html.EscapeString(string(product.SourceName)))
	}
	if product.SourceLink != "" {
		buffer.WriteString("`productSourceLink`=?,")
		args = append(args, product.SourceLink)
	}
	query := buffer.String()
	query = strings.TrimRight(query, ",")
	if product.ID == 0 {
		query = query + ";"
	} else {
		query = query + " WHERE `idProduct`=?;"
		args = append(args, strconv.Itoa(int(product.ID)))
	}
	glog.Infoln(query)
	r, _ := conn.Exec(query, args...)
	id, _ := r.LastInsertId()
	ret := int(id)
	return ret
}

//Insert a group product record into the database
func (productgroup *GroupProduct) InsertGroupProduct() {
	productgroup.processGroupProduct()
}

//Update a group product record in the database based on ID.  Clears then adds
//the record.
func (productgroup *GroupProduct) UpdateGroupProduct() {
	//clear out the old group for this id
	productgroup.clearGroupProductByBaseProductID()
	productgroup.processGroupProduct()
}

//Process an insert or an update.  update should have cleared first.
func (productgroup *GroupProduct) processGroupProduct() {
	conn, _ := connectors.GetDB()
	var args []interface{} //arguments for variables in the data struct
	groupproduct := productgroup.GroupProduct.SelectProduct()
	if len(groupproduct) > 0 {
		for _, productItem := range productgroup.Products {
			product := productItem.SelectProduct()
			if len(product) > 0 {
				query := "INSERT INTO `groupProduct` (`idBaseProduct`, `idProduct`) VALUES (?, ?);"
				args = append(args, strconv.Itoa(groupproduct[0].ID))
				args = append(args, strconv.Itoa(product[0].ID))
				glog.Infoln(query)
				conn.Exec(query, args...)
			}
		}
	}
}

//Delete the group product based on ID.
func (productgroup *GroupProduct) clearGroupProductByBaseProductID() {
	conn, _ := connectors.GetDB()
	var buffer bytes.Buffer
	var args []interface{}
	//delete all groupProduct from groupProduct table by idBaseProdcut
	buffer.WriteString("DELETE FROM `groupProduct` WHERE `idBaseProduct`=?;")
	args = append(args, int64(productgroup.GroupProduct.ID))
	query := buffer.String()
	glog.Infoln(query + " " + strconv.Itoa(int(int64(productgroup.GroupProduct.ID))))
	conn.Exec(query, args...)
}

//Insert a derived product record into the database
func (derivedproduct *DerivedProduct) InsertDerivedProduct() {
	derivedproduct.processDerivedProduct()
}

//Update a derived product record in the database based on ID.  Clears then adds
//the record.
func (derivedproduct *DerivedProduct) UpdateDerivedProduct() {
	//clear out the old group for this id
	derivedproduct.clearDerivedProductByProductID()
	derivedproduct.processDerivedProduct()
}

//Process an insert or an update.  update should have cleared first.
func (derivedproduct *DerivedProduct) processDerivedProduct() {
	conn, _ := connectors.GetDB()
	var args []interface{}
	baseproduct := derivedproduct.BaseProduct.SelectProduct()
	product := derivedproduct.Product.SelectProduct()
	if len(baseproduct) > 0 && len(product) > 0 {
		query := "INSERT INTO `derivedProduct` (`idBaseProduct`, `idProduct`) VALUES (?, ?);"
		args = append(args, strconv.Itoa(baseproduct[0].ID))
		args = append(args, strconv.Itoa(product[0].ID))
		glog.Infoln(query)
		conn.Exec(query, args...)
	}
}

//Delete the derived product based on ID.
func (derivedproduct *DerivedProduct) clearDerivedProductByProductID() {
	conn, _ := connectors.GetDB()
	var buffer bytes.Buffer
	var args []interface{}
	//delete all altingredients from altingredients table by stepid
	buffer.WriteString("DELETE FROM `derivedProduct` WHERE `idProduct`=?;")
	args = append(args, int64(derivedproduct.Product.ID))
	query := buffer.String()
	glog.Infoln(query + " " + strconv.Itoa(int(int64(derivedproduct.Product.ID))))
	conn.Exec(query, args...)
}

//SELECTS
//Select from the product table based on the attributes set in the product object.
func (product *Product) SelectProduct() []Product {
	var ret []Product
	conn, _ := connectors.GetDB()
	var args []interface{}
	glog.Infoln(product.ProductName)
	var buffer bytes.Buffer
	buffer.WriteString("SELECT `idProduct`, `productName`, `productType`, `productGroupType`, COALESCE(`productDescription`, ''), COALESCE(`productDetails`, ''), " +
		"COALESCE(`productImageSourceName`, ''), COALESCE(`productImage`, ''), COALESCE(`productImagePath`, ''), COALESCE(`productImageSourceLink`, ''), " +
		"COALESCE(`productPreText`, ''), COALESCE(`productPostText`, ''), COALESCE(`productRating`, 0), COALESCE(`productSourceName`, ''), COALESCE(`productSourceLink`, '') " +
		"FROM `product` WHERE ")
	if product.ID != 0 {
		buffer.WriteString(" `idProduct`=? AND")
		args = append(args, strconv.Itoa(product.ID))
	}
	if product.ProductName != "" {
		buffer.WriteString("`productName`=? AND")
		args = append(args, html.EscapeString(string(product.ProductName)))
	}
	if int(product.ProductType.ID) != 0 {
		buffer.WriteString(" `productType`=? AND")
		args = append(args, strconv.Itoa(int(product.ProductType.ID)))
	}
	if int(product.ProductGroupType) != 0 {
		buffer.WriteString(" `productGroupType`=? AND")
		args = append(args, strconv.Itoa(int(product.ProductGroupType)))
	}
	if product.Description != "" {
		buffer.WriteString("`productDescription`=? AND ")
		args = append(args, html.EscapeString(string(product.Description)))
	}
	if product.Details != "" {
		buffer.WriteString("`productDescription`=? AND ")
		args = append(args, html.EscapeString(string(product.Details)))
	}
	if product.PreText != "" {
		buffer.WriteString("`productPreText`=? AND")
		args = append(args, html.EscapeString(string(product.PreText)))
	}
	if product.PostText != "" {
		buffer.WriteString("`productPostText`=? AND")
		args = append(args, html.EscapeString(string(product.PostText)))
	}
	if product.Rating != 0 {
		buffer.WriteString(" `productRating`=? AND")
		args = append(args, strconv.Itoa(product.Rating))
	}
	if product.ImagePath != "" {
		buffer.WriteString("`productImagePath`=? AND")
		args = append(args, product.ImagePath)
	}
	if product.Image != "" {
		buffer.WriteString("`productImage`=? AND")
		args = append(args, product.Image)
	}
	if product.ImageSourceName != "" {
		buffer.WriteString("`productImageSourceName`=? AND")
		args = append(args, html.EscapeString(string(product.ImageSourceName)))
	}
	if product.ImageSourceLink != "" {
		buffer.WriteString("`productImageSourceLink`=? AND")
		args = append(args, product.ImageSourceLink)
	}
	if product.SourceName != "" {
		buffer.WriteString("`productSourceName`=? AND")
		args = append(args, html.EscapeString(string(product.SourceName)))
	}
	if product.SourceLink != "" {
		buffer.WriteString("`productSourceLink`=? AND")
		args = append(args, product.SourceLink)
	}

	query := buffer.String()
	query = strings.TrimRight(query, " WHERE")
	query = strings.TrimRight(query, " AND")
	query = query + " ORDER BY `productType`, `productGroupType`, `productName`;"
	glog.Infoln(query)
	rows, err := conn.Query(query, args...)
	if err != nil {
		glog.Error(err)
	}
	defer rows.Close()
	for rows.Next() {
		var prod Product
		var desc string
		var details string
		err := rows.Scan(&prod.ID, &prod.ProductName, &prod.ProductType.ID, &prod.ProductGroupType, &desc, &details, &prod.ImageSourceName, &prod.Image, &prod.ImagePath, &prod.ImageSourceLink, &prod.PreText, &prod.PostText, &prod.Rating, &prod.SourceName, &prod.SourceLink)
		if err != nil {
			glog.Error(err)
		}
		prod.Description = template.HTML(html.UnescapeString(desc))
		prod.Details = template.HTML(html.UnescapeString(details))
		ret = append(ret, prod)
		glog.Infoln(prod.ID, prod.ProductName, prod.ProductType.ID, prod.ProductGroupType, prod.Description, prod.Details, prod.ImageSourceName, prod.Image, prod.ImagePath, prod.ImageSourceLink, prod.PreText, prod.PostText, prod.Rating, prod.SourceName, prod.SourceLink)
	}
	err = rows.Err()
	if err != nil {
		glog.Error(err)
	}

	return ret
}

////Select a set of meta records based on the flags passed in via the metatypes
//table. If ignore cache is true then the database query is run otherwise the
//cache is checked first.
func (product *Product) SelectProductsByTypes(includeIngredients bool, includeNonIngredients bool, ignoreCache bool) ProductsByTypes {
	ret := new(ProductsByTypes)
	ret = nil
	if !ignoreCache {
		ret = product.memcachedProductsByTypes(includeIngredients, includeNonIngredients)
	}

	if ret == nil {
		ret = new(ProductsByTypes)
		conn, _ := connectors.GetDB()
		var args []interface{}
		rows, _ := conn.Query("SELECT COUNT(*) as count FROM  `producttype`;")
		count, err := checkCount(rows)
		glog.Infoln("Product Types Found " + strconv.Itoa(count))
		rows.Close()
		for i := 0; i < count; i++ {
			var pbt ProductsByType
			var buffer bytes.Buffer
			args = args[0:0]
			buffer.WriteString("SELECT `idProductType`, `productTypeName`, `productTypeIsIngredient` FROM  `producttype` WHERE idProductType=? AND")
			glog.Infoln("Getting Products by Type ID " + strconv.Itoa(i+1))
			args = append(args, strconv.Itoa(i+1))
			buffer.WriteString(" (")
			if includeIngredients {
				buffer.WriteString("`productTypeIsIngredient`=1 OR ")
			}
			if includeNonIngredients {
				buffer.WriteString("`productTypeIsIngredient`=0")
			}
			query := buffer.String()
			query = strings.TrimSuffix(query, "OR ")
			query = query + ")"
			query = strings.TrimSuffix(query, " AND")
			query = query + ";"
			glog.Infoln(query)
			pbt_rows, _ := conn.Query(query, args...)

			defer pbt_rows.Close()
			for pbt_rows.Next() {
				err = pbt_rows.Scan(&pbt.ProductType.ID, &pbt.ProductType.ProductTypeName, &pbt.ProductType.IsIngredient)
				if err != nil {
					glog.Error(err)
				}
				glog.Infoln(pbt.ProductType.ID, pbt.ProductType.ProductTypeName, pbt.ProductType.IsIngredient)
				if pbt.ProductType.ID != 0 {
					var inProduct Product
					inProduct.ProductType = pbt.ProductType
					//inProduct.ProductGroupType = Base
					outProduct := inProduct.SelectProduct()
					pbt.Products = outProduct
					// for _, element := range outProduct {
					// 	GetProductByIDWithBD(element.ID)
					// }
				}
			}
			if pbt.ProductType.ID != 0 {
				ret.PBT = append(ret.PBT, pbt)
			}
		}
	}
	return *ret
}

//Memcache retrieval of products by types
func (product *Product) memcachedProductsByTypes(includeIngredients bool, includeNonIngredients bool) *ProductsByTypes {
	ret := new(ProductsByTypes)
	ret = nil
	mc, _ := connectors.GetMC()
	if mc != nil {
		item := new(memcache.Item)
		item = nil
		if includeIngredients && includeNonIngredients {
			item, _ = mc.Get("pbt_tt")
		} else if includeIngredients && !includeNonIngredients {
			item, _ = mc.Get("pbt_tf")
		} else if !includeIngredients && includeNonIngredients {
			item, _ = mc.Get("pbt_ft")
		}

		if item != nil {
			if len(item.Value) > 0 {
				read := bytes.NewReader(item.Value)
				dec := gob.NewDecoder(read)
				dec.Decode(&ret)
			}
		}
	}
	return ret
}

//Select a product by an id number and include the base, derived, or group
//information
func (product *Product) SelectProductByIDWithBDG(ID int) *BaseProductWithBDG {
	var inProduct Product
	inProduct.ID = ID
	p := inProduct.SelectProduct()
	return p[0].SelectBDGByProduct()
}

//Select the base, derived, or group information for a product
func (product *Product) SelectBDGByProduct() *BaseProductWithBDG {
	conn, _ := connectors.GetDB()
	var bpwbd BaseProductWithBDG
	var inProduct Product
	bpwbd.Product = *product
	var dgp []Product
	var bp Product
	glog.Infoln("Product With ID for BD return " + strconv.Itoa(product.ID) + "and GroupType " + strconv.Itoa(int(product.ProductGroupType)))
	if product.ProductGroupType == Base {
		bd_rows, _ := conn.Query("SELECT `idProduct` FROM  `derivedProduct` WHERE idBaseProduct=?;", strconv.Itoa(product.ID))
		defer bd_rows.Close()
		for bd_rows.Next() {
			var derivedProductID int
			err := bd_rows.Scan(&derivedProductID)
			if err != nil {
				glog.Error(err)
			}
			glog.Infoln("Found Derived of " + strconv.Itoa(derivedProductID))
			inProduct.ID = derivedProductID
			derivedProduct := inProduct.SelectProduct()
			dgp = append(dgp, derivedProduct[0])
		}
		bpwbd.DerivedProducts = dgp
	} else if product.ProductGroupType == Derived {
		bd_rows, _ := conn.Query("SELECT `idBaseProduct` FROM  `derivedProduct` WHERE idProduct=?;", strconv.Itoa(product.ID))
		defer bd_rows.Close()
		for bd_rows.Next() {
			var baseProductID int
			err := bd_rows.Scan(&baseProductID)
			if err != nil {
				glog.Error(err)
			}
			glog.Infoln("Found Base of " + strconv.Itoa(baseProductID))
			inProduct.ID = baseProductID
			baseProduct := inProduct.SelectProduct()
			bp = baseProduct[0]
		}
		bpwbd.BaseProduct = bp
	} else if product.ProductGroupType == Group {
		bd_rows, _ := conn.Query("SELECT `idProduct` FROM  `groupProduct` WHERE idBaseProduct=?;", strconv.Itoa(product.ID))
		defer bd_rows.Close()
		for bd_rows.Next() {
			var groupProductID int
			err := bd_rows.Scan(&groupProductID)
			if err != nil {
				glog.Error(err)
			}
			glog.Infoln("Found Group of " + strconv.Itoa(groupProductID))
			inProduct.ID = groupProductID
			groupProduct := inProduct.SelectProduct()
			dgp = append(dgp, groupProduct[0])
		}
		bpwbd.GroupProducts = dgp
	}
	return &bpwbd
}

//Select all products in the database
func (product *Product) SelectAllProducts() []Product {
	var ret []Product
	conn, _ := connectors.GetDB()
	var buffer bytes.Buffer
	buffer.WriteString("SELECT idProduct, productName, productType, COALESCE(productDescription, ''), COALESCE(productImagePath, '')," +
		" COALESCE(productImage, '') FROM product;")
	query := buffer.String()
	glog.Infoln(query)
	rows, err := conn.Query(query)
	if err != nil {
		glog.Error(err)
	}
	defer rows.Close()
	for rows.Next() {
		var desc string
		var pt int
		var prod Product
		err := rows.Scan(&prod.ID, &prod.ProductName, &pt, &desc, &prod.ImagePath, &prod.Image)
		prod.Description = template.HTML(desc)
		prod.ProductType.ID = pt
		if err != nil {
			glog.Error(err)
		}
		glog.Infoln(prod.ID, prod.ProductName, int(prod.ProductType.ID), prod.Description, prod.ImagePath, prod.Image)
		ret = append(ret, prod)
	}
	err = rows.Err()
	if err != nil {
		glog.Error(err)
	}

	return ret
}

//Select a product by an id number
func (product *Product) SelectProductByID(ID int) Product {
	var inProduct Product
	inProduct.ID = ID
	p := inProduct.SelectProduct()
	return p[0]
}

//Select a set of products by an associated cocktail id number and product type
//id number
func (product *Product) SelectProductsByCocktailAndProductType(ID int, pt int) []Product {
	var ret []Product
	conn, _ := connectors.GetDB()
	var args []interface{}
	var buffer bytes.Buffer
	buffer.WriteString("SELECT product.idProduct, product.productName, product.productType, product.productGroupType," +
		" COALESCE(product.productDescription, ''), COALESCE(product.productDetails, ''), COALESCE(product.productImageSourceName, '')," +
		" COALESCE(product.productImage, ''), COALESCE(product.productImagePath, ''), COALESCE(product.productImageSourceLink, '')," +
		" COALESCE(product.productPreText, ''), COALESCE(product.productPostText, ''), COALESCE(product.productRating, 0)," +
		" COALESCE(product.productSourceName, ''), COALESCE(product.productSourceLink, '')" +
		" FROM product" +
		" JOIN cocktailToProducts ON product.idProduct=cocktailToProducts.idProduct" +
		" JOIN cocktail ON cocktailToProducts.idCocktail=cocktail.idCocktail" +
		" WHERE cocktail.idCocktail=? AND product.productType=?;")
	args = append(args, strconv.Itoa(ID))
	args = append(args, strconv.Itoa(pt))
	query := buffer.String()
	glog.Infoln(query)
	rows, err := conn.Query(query, args...)
	if err != nil {
		glog.Error(err)
	}
	defer rows.Close()
	for rows.Next() {
		var prod Product
		var desc string
		var details string
		err := rows.Scan(&prod.ID, &prod.ProductName, &prod.ProductType.ID, &prod.ProductGroupType, &desc, &details, &prod.ImageSourceName, &prod.Image, &prod.ImagePath, &prod.ImageSourceLink, &prod.PreText, &prod.PostText, &prod.Rating, &prod.SourceName, &prod.SourceLink)
		if err != nil {
			glog.Error(err)
		}
		prod.Description = template.HTML(html.UnescapeString(desc))
		prod.Details = template.HTML(html.UnescapeString(details))
		glog.Infoln(prod.ID, prod.ProductName, int(prod.ProductType.ID), prod.Description, prod.ImagePath, prod.Image)
		ret = append(ret, prod)
	}
	err = rows.Err()
	if err != nil {
		glog.Error(err)
	}

	return ret
}
