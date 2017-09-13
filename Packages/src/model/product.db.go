// Copyright 2017 Rough Industries LLC. All rights reserved.
//model/product.db.go:package model
package model

import (
	"bytes"
	"connectors"
	"encoding/gob"
	"github.com/bradfitz/gomemcache/memcache"
	"html"
	"html/template"
	"log"
	"strconv"
	"strings"
)

//CREATE, UPDATE, DELETE
//Insert a product record into the database
func InsertProduct(product Product) int {
	product.ID = 0
	return processProduct(product)
}

//
func UpdateProduct(product Product) int {
	return processProduct(product)
}

//
func processProduct(product Product) int {
	conn, _ := connectors.GetDB()
	var buffer bytes.Buffer
	if product.ID == 0 {
		buffer.WriteString("INSERT INTO `commonwealthcocktails`.`product` SET ")
	} else {
		buffer.WriteString("UPDATE `commonwealthcocktails`.`product` SET ")
	}
	if product.ProductName != "" {
		sqlString := strings.Replace(string(product.ProductName), "\\", "\\\\", -1)
		sqlString = strings.Replace(sqlString, "\"", "\\\"", -1)
		buffer.WriteString("`productName`=\"" + sqlString + "\",")
	}
	buffer.WriteString(" `productType`=" + strconv.Itoa(int(product.ProductType.ID)) + ",")
	buffer.WriteString(" `productGroupType`=" + strconv.Itoa(int(product.ProductGroupType)) + ",")
	if product.Description != "" {
		buffer.WriteString("`productDescription`=\"" + html.EscapeString(string(product.Description)) + "\",")
	}
	if product.Details != "" {
		buffer.WriteString("`productDetails`=\"" + html.EscapeString(string(product.Details)) + "\",")
	}
	if product.PreText != "" {
		buffer.WriteString("`productPreText`=\"" + product.PreText + "\",")
	}
	if product.PostText != "" {
		buffer.WriteString("`productPostText`=\"" + product.PostText + "\",")
	}
	if product.Rating != 0 {
		buffer.WriteString(" `productRating`=" + strconv.Itoa(product.Rating) + ",")
	}
	if product.ImagePath != "" {
		buffer.WriteString("`productImagePath`=\"" + product.ImagePath + "\",")
	}
	if product.Image != "" {
		buffer.WriteString("`productImage`=\"" + product.Image + "\",")
	}
	if product.ImageSourceName != "" {
		buffer.WriteString("`productImageSourceName`=\"" + product.ImageSourceName + "\",")
	}
	if product.ImageSourceLink != "" {
		buffer.WriteString("`productImageSourceLink`=\"" + product.ImageSourceLink + "\",")
	}
	if product.SourceName != "" {
		buffer.WriteString("`productSourceName`=\"" + product.SourceName + "\",")
	}
	if product.SourceLink != "" {
		buffer.WriteString("`productSourceLink`=\"" + product.SourceLink + "\",")
	}
	query := buffer.String()
	query = strings.TrimRight(query, ",")
	if product.ID == 0 {
		query = query + ";"
	} else {
		query = query + " WHERE `idProduct`=" + strconv.Itoa(int(product.ID)) + ";"
	}
	log.Println(query)
	r, _ := conn.Exec(query)
	id, _ := r.LastInsertId()
	ret := int(id)
	return ret
}

//
func InsertGroupProduct(productgroup GroupProduct) {
	processGroupProduct(productgroup)
}

//
func UpdateGroupProduct(productgroup GroupProduct) {
	//clear out the old group for this id
	clearGroupProductByBaseProductID(int64(productgroup.GroupProduct.ID))
	processGroupProduct(productgroup)
}

//
func processGroupProduct(productgroup GroupProduct) {
	conn, _ := connectors.GetDB()

	//TODO: handle updates which requier deletion of old relationships
	groupproduct := SelectProduct(productgroup.GroupProduct)
	if len(groupproduct) > 0 {
		for _, productItem := range productgroup.Products {
			product := SelectProduct(productItem)
			if len(product) > 0 {
				query := "INSERT INTO `commonwealthcocktails`.`groupProduct` (`idBaseProduct`, `idProduct`) VALUES ('" + strconv.Itoa(groupproduct[0].ID) + "', '" + strconv.Itoa(product[0].ID) + "');"
				log.Println(query)
				conn.Exec(query)
			}
		}
	}
}

//
func clearGroupProductByBaseProductID(productID int64) {
	conn, _ := connectors.GetDB()

	var buffer bytes.Buffer
	var args []interface{}
	//delete all groupProduct from groupProduct table by idBaseProdcut
	buffer.WriteString("DELETE FROM `commonwealthcocktails`.`groupProduct` WHERE `idBaseProduct`=?;")
	args = append(args, productID)
	query := buffer.String()
	log.Println(query + " " + strconv.Itoa(int(productID)))
	conn.Exec(query, args...)
}

//
func InsertDerivedProduct(derivedproduct DerivedProduct) {
	processDerivedProduct(derivedproduct)
}

//
func UpdateDerivedProduct(derivedproduct DerivedProduct) {
	//clear out the old group for this id
	clearDerivedProductByProductID(int64(derivedproduct.Product.ID))
	processDerivedProduct(derivedproduct)
}

//
func processDerivedProduct(derivedproduct DerivedProduct) {
	conn, _ := connectors.GetDB()

	//TODO: handle updates which requier deletion of old relationships
	baseproduct := SelectProduct(derivedproduct.BaseProduct)
	product := SelectProduct(derivedproduct.Product)
	if len(baseproduct) > 0 && len(product) > 0 {
		query := "INSERT INTO `commonwealthcocktails`.`derivedProduct` (`idBaseProduct`, `idProduct`) VALUES ('" + strconv.Itoa(baseproduct[0].ID) + "', '" + strconv.Itoa(product[0].ID) + "');"
		log.Println(query)
		conn.Exec(query)
	}
}

//
func clearDerivedProductByProductID(productID int64) {
	conn, _ := connectors.GetDB()

	var buffer bytes.Buffer
	var args []interface{}
	//delete all altingredients from altingredients table by stepid
	buffer.WriteString("DELETE FROM `commonwealthcocktails`.`derivedProduct` WHERE `idProduct`=?;")
	args = append(args, productID)
	query := buffer.String()
	log.Println(query + " " + strconv.Itoa(int(productID)))
	conn.Exec(query, args...)
}

//SELECTS
//
func SelectProduct(product Product) []Product {
	var ret []Product
	conn, _ := connectors.GetDB()

	log.Println(product.ProductName)
	var buffer bytes.Buffer
	buffer.WriteString("SELECT `idProduct`, `productName`, `productType`, `productGroupType`, COALESCE(`productDescription`, ''), COALESCE(`productDetails`, ''), " +
		"COALESCE(`productImageSourceName`, ''), COALESCE(`productImage`, ''), COALESCE(`productImagePath`, ''), COALESCE(`productImageSourceLink`, ''), " +
		"COALESCE(`productPreText`, ''), COALESCE(`productPostText`, ''), COALESCE(`productRating`, 0), COALESCE(`productSourceName`, ''), COALESCE(`productSourceLink`, '') " +
		"FROM `commonwealthcocktails`.`product` WHERE ")
	if product.ID != 0 {
		buffer.WriteString(" `idProduct`=" + strconv.Itoa(product.ID) + " AND")
	}
	if product.ProductName != "" {
		sqlString := strings.Replace(string(product.ProductName), "\\", "\\\\", -1)
		sqlString = strings.Replace(sqlString, "\"", "\\\"", -1)
		buffer.WriteString("`productName`=\"" + sqlString + "\" AND")
	}
	if int(product.ProductType.ID) != 0 {
		buffer.WriteString(" `productType`=" + strconv.Itoa(int(product.ProductType.ID)) + " AND")
	}
	if int(product.ProductGroupType) != 0 {
		buffer.WriteString(" `productGroupType`=" + strconv.Itoa(int(product.ProductGroupType)) + " AND")
	}
	if product.Description != "" {
		buffer.WriteString("`productDescription`=\"" + html.EscapeString(string(product.Description)) + "\" AND ")
	}
	if product.Details != "" {
		buffer.WriteString("`productDescription`=\"" + html.EscapeString(string(product.Details)) + "\" AND ")
	}
	if product.PreText != "" {
		buffer.WriteString("`productPreText`=\"" + product.PreText + "\" AND")
	}
	if product.PostText != "" {
		buffer.WriteString("`productPostText`=\"" + product.PostText + "\" AND")
	}
	if product.Rating != 0 {
		buffer.WriteString(" `productRating`=" + strconv.Itoa(product.Rating) + " AND")
	}
	if product.ImagePath != "" {
		buffer.WriteString("`productImagePath`=\"" + product.ImagePath + "\" AND")
	}
	if product.Image != "" {
		buffer.WriteString("`productImage`=\"" + product.Image + "\" AND")
	}
	if product.ImageSourceName != "" {
		buffer.WriteString("`productImageSourceName`=\"" + product.ImageSourceName + "\" AND")
	}
	if product.ImageSourceLink != "" {
		buffer.WriteString("`productImageSourceLink`=\"" + product.ImageSourceLink + "\" AND")
	}
	if product.SourceName != "" {
		buffer.WriteString("`productSourceName`=\"" + product.SourceName + "\" AND")
	}
	if product.SourceLink != "" {
		buffer.WriteString("`productSourceLink`=\"" + product.SourceLink + "\" AND")
	}

	query := buffer.String()
	query = strings.TrimRight(query, " WHERE")
	query = strings.TrimRight(query, " AND")
	query = query + " ORDER BY `productType`, `productGroupType`, `productName`;"
	log.Println(query)
	rows, err := conn.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var prod Product
		var desc string
		var details string
		err := rows.Scan(&prod.ID, &prod.ProductName, &prod.ProductType.ID, &prod.ProductGroupType, &desc, &details, &prod.ImageSourceName, &prod.Image, &prod.ImagePath, &prod.ImageSourceLink, &prod.PreText, &prod.PostText, &prod.Rating, &prod.SourceName, &prod.SourceLink)
		if err != nil {
			log.Fatal(err)
		}
		prod.Description = template.HTML(html.UnescapeString(desc))
		prod.Details = template.HTML(html.UnescapeString(details))
		ret = append(ret, prod)
		log.Println(prod.ID, prod.ProductName, prod.ProductType.ID, prod.ProductGroupType, prod.Description, prod.Details, prod.ImageSourceName, prod.Image, prod.ImagePath, prod.ImageSourceLink, prod.PreText, prod.PostText, prod.Rating, prod.SourceName, prod.SourceLink)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return ret
}

//
func SelectProductsByTypes(includeIngredients bool, includeNonIngredients bool, ignoreCache bool) ProductsByTypes {
	ret := new(ProductsByTypes)
	ret = nil
	if !ignoreCache {
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
	}

	if ret == nil {
		ret = new(ProductsByTypes)
		conn, _ := connectors.GetDB()

		rows, _ := conn.Query("SELECT COUNT(*) as count FROM  `commonwealthcocktails`.`producttype`;")
		count, err := checkCount(rows)
		log.Println("Product Types Found " + strconv.Itoa(count))
		rows.Close()
		for i := 0; i < count; i++ {
			var pbt ProductsByType
			var buffer bytes.Buffer
			buffer.WriteString("SELECT `idProductType`, `productTypeName`, `productTypeIsIngredient` FROM  `commonwealthcocktails`.`producttype` WHERE idProductType='" + strconv.Itoa(i+1) + "' AND")
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
			log.Println(query)
			pbt_rows, _ := conn.Query(query)

			defer pbt_rows.Close()
			for pbt_rows.Next() {
				err = pbt_rows.Scan(&pbt.ProductType.ID, &pbt.ProductType.ProductTypeName, &pbt.ProductType.IsIngredient)
				if err != nil {
					log.Fatal(err)
				}
				if pbt.ProductType.ID != 0 {
					var inProduct Product
					inProduct.ProductType = pbt.ProductType
					//inProduct.ProductGroupType = Base
					outProduct := SelectProduct(inProduct)
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

//
func SelectProductByIDWithBDG(ID int) *BaseProductWithBDG {
	var inProduct Product
	inProduct.ID = ID
	p := SelectProduct(inProduct)
	return SelectBDGByProduct(p[0])
}

//
func SelectBDGByProduct(product Product) *BaseProductWithBDG {
	conn, _ := connectors.GetDB()

	var bpwbd BaseProductWithBDG
	var inProduct Product
	bpwbd.Product = product
	var dgp []Product
	var bp Product
	log.Println("Product With ID for BD return " + strconv.Itoa(product.ID) + "and GroupType " + strconv.Itoa(int(product.ProductGroupType)))
	if product.ProductGroupType == Base {
		bd_rows, _ := conn.Query("SELECT `idProduct` FROM  `commonwealthcocktails`.`derivedProduct` WHERE idBaseProduct='" + strconv.Itoa(product.ID) + "';")
		defer bd_rows.Close()
		for bd_rows.Next() {
			var derivedProductID int
			err := bd_rows.Scan(&derivedProductID)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("Found Derived of " + strconv.Itoa(derivedProductID))
			inProduct.ID = derivedProductID
			derivedProduct := SelectProduct(inProduct)
			dgp = append(dgp, derivedProduct[0])
		}
		bpwbd.DerivedProducts = dgp
	} else if product.ProductGroupType == Derived {
		bd_rows, _ := conn.Query("SELECT `idBaseProduct` FROM  `commonwealthcocktails`.`derivedProduct` WHERE idProduct='" + strconv.Itoa(product.ID) + "';")
		defer bd_rows.Close()
		for bd_rows.Next() {
			var baseProductID int
			err := bd_rows.Scan(&baseProductID)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("Found Base of " + strconv.Itoa(baseProductID))
			inProduct.ID = baseProductID
			baseProduct := SelectProduct(inProduct)
			bp = baseProduct[0]
		}
		bpwbd.BaseProduct = bp
	} else if product.ProductGroupType == Group {
		bd_rows, _ := conn.Query("SELECT `idProduct` FROM  `commonwealthcocktails`.`groupProduct` WHERE idBaseProduct='" + strconv.Itoa(product.ID) + "';")
		defer bd_rows.Close()
		for bd_rows.Next() {
			var groupProductID int
			err := bd_rows.Scan(&groupProductID)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("Found Group of " + strconv.Itoa(groupProductID))
			inProduct.ID = groupProductID
			groupProduct := SelectProduct(inProduct)
			dgp = append(dgp, groupProduct[0])
		}
		bpwbd.GroupProducts = dgp
	}
	return &bpwbd
}

//func SelectProductsByCocktail(cocktail Cocktail) []Product{
//SELECT * FROM commonwealthcocktails.product
//JOIN commonwealthcocktails.cocktailToProducts ON cocktailToProducts.idProduct=product.idProduct
//JOIN  commonwealthcocktails.cocktail ON cocktailToProducts.idCocktail=cocktail.idCocktail
//WHERE cocktail.idCocktail=2;
//}

//
func SelectAllProducts() []Product {
	var ret []Product
	conn, _ := connectors.GetDB()

	var buffer bytes.Buffer
	var canQuery = false
	buffer.WriteString("SELECT idProduct, productName, productType, COALESCE(productDescription, ''), COALESCE(productImagePath, '')," +
		" COALESCE(productImage, '') FROM commonwealthcocktails.product;")
	canQuery = true

	if canQuery {
		query := buffer.String()
		query = strings.TrimRight(query, " AND")
		query = query + ";"
		log.Println(query)
		rows, err := conn.Query(query)
		if err != nil {
			log.Fatal(err)
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
				log.Fatal(err)
			}
			log.Println(prod.ID, prod.ProductName, int(prod.ProductType.ID), prod.Description, prod.ImagePath, prod.Image)
			ret = append(ret, prod)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
	}
	return ret
}

//
func SelectProductByID(ID int) Product {
	var ret Product
	conn, _ := connectors.GetDB()

	var buffer bytes.Buffer
	var canQuery = false
	buffer.WriteString("SELECT idProduct, productName, productType, COALESCE(productDescription, ''), COALESCE(productImagePath, '')," +
		" COALESCE(productImage, '') FROM commonwealthcocktails.product WHERE idProduct=" + strconv.Itoa(ID) + ";")
	canQuery = true

	if canQuery {
		query := buffer.String()
		query = strings.TrimRight(query, " AND")
		query = query + ";"
		log.Println(query)
		rows, err := conn.Query(query)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			var desc string
			var pt int
			err := rows.Scan(&ret.ID, &ret.ProductName, &pt, &desc, &ret.ImagePath, &ret.Image)
			ret.Description = template.HTML(desc)
			ret.ProductType.ID = pt
			if err != nil {
				log.Fatal(err)
			}
			log.Println(ret.ID, ret.ProductName, int(ret.ProductType.ID), ret.Description, ret.ImagePath, ret.Image)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
	}
	return ret
}

//
func SelectProductsByCocktailAndProductType(ID int, pt int) []Product {
	var ret []Product
	conn, _ := connectors.GetDB()
	var args []interface{}

	var buffer bytes.Buffer
	var canQuery = false
	buffer.WriteString("SELECT product.idProduct, product.productName, product.productType, product.productGroupType," +
		" COALESCE(product.productDescription, ''), COALESCE(product.productDetails, ''), COALESCE(product.productImageSourceName, '')," +
		" COALESCE(product.productImage, ''), COALESCE(product.productImagePath, ''), COALESCE(product.productImageSourceLink, '')," +
		" COALESCE(product.productPreText, ''), COALESCE(product.productPostText, ''), COALESCE(product.productRating, 0)," +
		" COALESCE(product.productSourceName, ''), COALESCE(product.productSourceLink, '')" +
		" FROM commonwealthcocktails.product" +
		" JOIN commonwealthcocktails.cocktailToProducts ON product.idProduct=cocktailToProducts.idProduct" +
		" JOIN commonwealthcocktails.cocktail ON cocktailToProducts.idCocktail=cocktail.idCocktail" +
		" WHERE cocktail.idCocktail=? AND product.productType=?;")
	args = append(args, strconv.Itoa(ID))
	args = append(args, strconv.Itoa(pt))

	canQuery = true

	if canQuery {
		query := buffer.String()
		query = strings.TrimRight(query, " AND")
		query = query + ";"
		log.Println(query)
		rows, err := conn.Query(query, args...)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			var prod Product
			var desc string
			var details string
			err := rows.Scan(&prod.ID, &prod.ProductName, &prod.ProductType.ID, &prod.ProductGroupType, &desc, &details, &prod.ImageSourceName, &prod.Image, &prod.ImagePath, &prod.ImageSourceLink, &prod.PreText, &prod.PostText, &prod.Rating, &prod.SourceName, &prod.SourceLink)
			if err != nil {
				log.Fatal(err)
			}
			prod.Description = template.HTML(html.UnescapeString(desc))
			prod.Details = template.HTML(html.UnescapeString(details))
			log.Println(prod.ID, prod.ProductName, int(prod.ProductType.ID), prod.Description, prod.ImagePath, prod.Image)
			ret = append(ret, prod)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
	}
	return ret
}
