//model/product.db.go
package model

import (
	"bytes"
	"database/sql"
	"db"
	"html/template"
	"log"
	"strconv"
	"strings"
)

func InitProductTables() {
	conn, _ := db.GetDB()

	var temp string
	if err := conn.QueryRow("SHOW TABLES LIKE 'producttype';").Scan(&temp); err == nil {
		log.Println("producttype Table Exists")
	} else if err == sql.ErrNoRows {
		log.Println("Creating producttype Table")
		conn.Query("CREATE TABLE `commonwealthcocktails`.`producttype` (`idProductType` INT NOT NULL AUTO_INCREMENT,PRIMARY KEY (`idProductType`));")
		conn.Query("ALTER TABLE `commonwealthcocktails`.`producttype` ADD COLUMN `productTypeName` VARCHAR(150) NOT NULL AFTER `idProductType`;")
		conn.Exec("INSERT INTO `commonwealthcocktails`.`producttype` (`idProductType`, `productTypeName`) VALUES ('1', 'Spirit');")
		conn.Exec("INSERT INTO `commonwealthcocktails`.`producttype` (`idProductType`, `productTypeName`) VALUES ('2', 'Liqueur');")
		conn.Exec("INSERT INTO `commonwealthcocktails`.`producttype` (`idProductType`, `productTypeName`) VALUES ('3', 'Wine');")
		conn.Exec("INSERT INTO `commonwealthcocktails`.`producttype` (`idProductType`, `productTypeName`) VALUES ('4', 'Mixer');")
		conn.Exec("INSERT INTO `commonwealthcocktails`.`producttype` (`idProductType`, `productTypeName`) VALUES ('5', 'Beer');")
		conn.Exec("INSERT INTO `commonwealthcocktails`.`producttype` (`idProductType`, `productTypeName`) VALUES ('6', 'Garnish');")
		conn.Exec("INSERT INTO `commonwealthcocktails`.`producttype` (`idProductType`, `productTypeName`) VALUES ('7', 'Drinkware');")
		conn.Exec("INSERT INTO `commonwealthcocktails`.`producttype` (`idProductType`, `productTypeName`) VALUES ('8', 'Tool');")
	}

	if err := conn.QueryRow("SHOW TABLES LIKE 'product';").Scan(&temp); err == nil {
		log.Println("product Table Exists")
	} else if err == sql.ErrNoRows {
		log.Println("Creating product Table")
		conn.Query("CREATE TABLE `commonwealthcocktails`.`product` (`idProduct` INT NOT NULL AUTO_INCREMENT,PRIMARY KEY (`idProduct`));") //ID
		conn.Query("ALTER TABLE `commonwealthcocktails`.`product`" +
			"ADD COLUMN `productName` VARCHAR(150) NOT NULL AFTER `idProduct`," + //ProductName
			"ADD COLUMN `productType`  INT NOT NULL AFTER `productName`," + //ProductType
			"ADD COLUMN `productDescription` VARCHAR(1500) AFTER `productType`," + //Description
			"ADD COLUMN `productDetails` VARCHAR(1500) AFTER `productDescription`," + //Details
			"ADD COLUMN `productImagePath` VARCHAR(750) AFTER `productDetails`," + //ImagePath
			"ADD COLUMN `productImage` VARCHAR(500) AFTER `productImagePath`," + //Image
			"ADD COLUMN `productImageSourceName` VARCHAR(500) AFTER `productImage`," + //ImageSourceName
			"ADD COLUMN `productImageSourceLink` VARCHAR(750) AFTER `productImageSourceName`," + //ImageSourceLink
			"ADD COLUMN `productArticle` INT AFTER `productImageSourceLink`," + //Article
			"ADD COLUMN `productRecipe` INT AFTER `productArticle`," + //Recipe
			"ADD COLUMN `productGroupType` INT AFTER `productRecipe`," + //ProductGroupType
			"ADD COLUMN `productPreText` VARCHAR(250) AFTER `productGroupType`," + //PreText
			"ADD COLUMN `productPostText` VARCHAR(250) AFTER `productPreText`," + //PostText
			"ADD COLUMN `productRating` INT(1) AFTER `productPostText`," + //Rating
			"ADD COLUMN `productSourceName` VARCHAR(1500) AFTER `productRating`," + //SourceName
			"ADD COLUMN `productSourceLink` VARCHAR(1500) AFTER `productSourceName`," + //SourceLink
			"ADD COLUMN `productAbout` INT AFTER `productSourceLink`," +
			"ADD UNIQUE INDEX `productName_UNIQUE` (`productName` ASC);") //About

	}
}

func InitProductReferences() {
	conn, _ := db.GetDB()

	conn.Query("ALTER TABLE `commonwealthcocktails`.`product`" +
		"ADD CONSTRAINT product_producttype_id_fk FOREIGN KEY(productType) REFERENCES producttype(idProductType)," +
		"ADD CONSTRAINT product_productgrouptype_id_fk FOREIGN KEY(productGroupType) REFERENCES grouptype(idGroupType)," +
		"ADD CONSTRAINT product_productarticle_id_fk FOREIGN KEY(productArticle) REFERENCES post(idPost)," +
		"ADD CONSTRAINT product_productrecipe_id_fk FOREIGN KEY(productRecipe) REFERENCES recipe(idRecipe)," +
		"ADD CONSTRAINT product_productabout_id_fk FOREIGN KEY(productAbout) REFERENCES post(idPost);")

	addProductToMetasTables()
	addDerivedProductTables()
	addGroupProductTables()
}

func addProductToMetasTables() {
	conn, _ := db.GetDB()

	var temp string
	if err := conn.QueryRow("SHOW TABLES LIKE 'productToMetas';").Scan(&temp); err == nil {
		log.Println("productToMetas Table Exists")
	} else if err == sql.ErrNoRows {
		log.Println("Creating productToMetas Table")
		//Drink
		conn.Exec("CREATE TABLE `commonwealthcocktails`.`productToMetas` (`idProductToMetas` INT NOT NULL AUTO_INCREMENT," +
			" `idProduct` INT NOT NULL," +
			" `idMeta` INT NOT NULL," +
			" PRIMARY KEY (`idProductToMetas`)," +
			" CONSTRAINT productToMetas_idProduct_id_fk FOREIGN KEY(idProduct) REFERENCES product(idProduct)," +
			" CONSTRAINT productToMetas_idMeta_id_fk FOREIGN KEY(idMeta) REFERENCES meta(idMeta));")
	}
}

func addDerivedProductTables() {
	conn, _ := db.GetDB()

	var temp string
	if err := conn.QueryRow("SHOW TABLES LIKE 'derivedProduct';").Scan(&temp); err == nil {
		log.Println("derivedProduct Table Exists")
	} else if err == sql.ErrNoRows {
		log.Println("Creating derivedProduct Table")
		//Drink
		conn.Exec("CREATE TABLE `commonwealthcocktails`.`derivedProduct` (`idDerivedProduct` INT NOT NULL AUTO_INCREMENT," +
			" `idBaseProduct` INT NOT NULL," +
			" `idProduct` INT NOT NULL," +
			" PRIMARY KEY (`idDerivedProduct`)," +
			" CONSTRAINT derivedProduct_idBaseProduct_id_fk FOREIGN KEY(idBaseProduct) REFERENCES product(idProduct)," +
			" CONSTRAINT derivedProduct_idProduct_id_fk FOREIGN KEY(idProduct) REFERENCES product(idProduct));")
	}
}

func addGroupProductTables() {
	conn, _ := db.GetDB()

	var temp string
	if err := conn.QueryRow("SHOW TABLES LIKE 'groupProduct';").Scan(&temp); err == nil {
		log.Println("groupProduct Table Exists")
	} else if err == sql.ErrNoRows {
		log.Println("Creating groupProduct Table")
		//Drink
		conn.Exec("CREATE TABLE `commonwealthcocktails`.`groupProduct` (`idGroupProduct` INT NOT NULL AUTO_INCREMENT," +
			" `idBaseProduct` INT NOT NULL," +
			" `idProduct` INT NOT NULL," +
			" PRIMARY KEY (`idGroupProduct`)," +
			" CONSTRAINT groupProduct_idGroup_id_fk FOREIGN KEY(idGroupProduct) REFERENCES product(idProduct)," +
			" CONSTRAINT groupProduct_idProduct_id_fk FOREIGN KEY(idProduct) REFERENCES product(idProduct));")
	}
}

func ProcessProducts() {
	conn, _ := db.GetDB()

	for _, product := range Products {
		log.Println(product.ProductName)
		var buffer bytes.Buffer
		buffer.WriteString("INSERT INTO `commonwealthcocktails`.`product` SET ")
		if product.ProductName != "" {
			sqlString := strings.Replace(string(product.ProductName), "\\", "\\\\", -1)
			sqlString = strings.Replace(sqlString, "\"", "\\\"", -1)
			buffer.WriteString("`productName`=\"" + sqlString + "\",")
		}
		buffer.WriteString(" `productType`=" + strconv.Itoa(int(product.ProductType)) + ",")
		buffer.WriteString(" `productGroupType`=" + strconv.Itoa(int(product.ProductGroupType)) + ",")
		if product.Description != "" {
			sqlString := strings.Replace(string(product.Description), "\\", "\\\\", -1)
			sqlString = strings.Replace(sqlString, "\"", "\\\"", -1)
			buffer.WriteString("`productDescription`=\"" + sqlString + "\",")
		}
		if product.Details != "" {
			sqlString := strings.Replace(string(product.Details), "\\", "\\\\", -1)
			sqlString = strings.Replace(sqlString, "\"", "\\\"", -1)
			buffer.WriteString("`productDetails`=\"" + sqlString + "\",")
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
		query = query + ";"
		log.Println(query)
		conn.Exec(query)
	}
}

//going to have to update this a bit TCH
func ProcessProductGroups() {
	conn, _ := db.GetDB()

	for _, productgroup := range ProductGroups {
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

}

func ProcessDerivedProducts() {
	conn, _ := db.GetDB()

	for _, derivedproduct := range DerivedProducts {
		baseproduct := SelectProduct(derivedproduct.BaseProduct)
		product := SelectProduct(derivedproduct.Product)
		if len(baseproduct) > 0 && len(product) > 0 {
			query := "INSERT INTO `commonwealthcocktails`.`derivedProduct` (`idBaseProduct`, `idProduct`) VALUES ('" + strconv.Itoa(baseproduct[0].ID) + "', '" + strconv.Itoa(product[0].ID) + "');"
			log.Println(query)
			conn.Exec(query)
		}
	}

}

func SelectProduct(product Product) []Product {
	var ret []Product
	conn, _ := db.GetDB()

	var canQuery = false
	log.Println(product.ProductName)
	var buffer bytes.Buffer
	buffer.WriteString("SELECT `idProduct`, `productName`, `productType`, `productGroupType` FROM `commonwealthcocktails`.`product` WHERE ")
	if product.ID != 0 {
		buffer.WriteString(" `idProduct`=" + strconv.Itoa(product.ID) + " AND")
		canQuery = true
	}
	if product.ProductName != "" {
		sqlString := strings.Replace(string(product.ProductName), "\\", "\\\\", -1)
		sqlString = strings.Replace(sqlString, "\"", "\\\"", -1)
		buffer.WriteString("`productName`=\"" + sqlString + "\" AND")
		canQuery = true
	}
	if int(product.ProductType) != 0 {
		buffer.WriteString(" `productType`=" + strconv.Itoa(int(product.ProductType)) + " AND")
		canQuery = true
	}
	if int(product.ProductGroupType) != 0 {
		buffer.WriteString(" `productGroupType`=" + strconv.Itoa(int(product.ProductGroupType)) + " AND")
		canQuery = true
	}
	if product.Description != "" {
		sqlString := strings.Replace(string(product.Description), "\\", "\\\\", -1)
		sqlString = strings.Replace(sqlString, "\"", "\\\"", -1)
		buffer.WriteString("`productDescription`=\"" + sqlString + "\" AND")
		canQuery = true
	}
	if product.Details != "" {
		sqlString := strings.Replace(string(product.Details), "\\", "\\\\", -1)
		sqlString = strings.Replace(sqlString, "\"", "\\\"", -1)
		buffer.WriteString("`productDetails`=\"" + sqlString + "\" AND")
		canQuery = true
	}
	if product.PreText != "" {
		buffer.WriteString("`productPreText`=\"" + product.PreText + "\" AND")
		canQuery = true
	}
	if product.PostText != "" {
		buffer.WriteString("`productPostText`=\"" + product.PostText + "\" AND")
		canQuery = true
	}
	if product.Rating != 0 {
		buffer.WriteString(" `productRating`=" + strconv.Itoa(product.Rating) + " AND")
		canQuery = true
	}
	if product.ImagePath != "" {
		buffer.WriteString("`productImagePath`=\"" + product.ImagePath + "\" AND")
		canQuery = true
	}
	if product.Image != "" {
		buffer.WriteString("`productImage`=\"" + product.Image + "\" AND")
		canQuery = true
	}
	if product.ImageSourceName != "" {
		buffer.WriteString("`productImageSourceName`=\"" + product.ImageSourceName + "\" AND")
		canQuery = true
	}
	if product.ImageSourceLink != "" {
		buffer.WriteString("`productImageSourceLink`=\"" + product.ImageSourceLink + "\" AND")
		canQuery = true
	}
	if product.SourceName != "" {
		buffer.WriteString("`productSourceName`=\"" + product.SourceName + "\" AND")
		canQuery = true
	}
	if product.SourceLink != "" {
		buffer.WriteString("`productSourceLink`=\"" + product.SourceLink + "\" AND")
		canQuery = true
	}

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
			var prod Product
			err := rows.Scan(&prod.ID, &prod.ProductName, &prod.ProductType, &prod.ProductGroupType)
			if err != nil {
				log.Fatal(err)
			}
			ret = append(ret, prod)
			log.Println(prod.ID, prod.ProductName, prod.ProductType, prod.ProductGroupType)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
	}
	return ret

}

func GetBaseProductByTypes() ProductByTypes {
	var ret ProductByTypes
	conn, _ := db.GetDB()

	rows, _ := conn.Query("SELECT COUNT(*) as count FROM  `commonwealthcocktails`.`producttype`;")
	count, err := checkCount(rows)
	log.Println("Product Types Found " + strconv.Itoa(count))
	rows.Close()
	for i := 0; i < count; i++ {
		var pbt ProductByType
		pbt_rows, _ := conn.Query("SELECT `idProductType`, `productTypeName` FROM  `commonwealthcocktails`.`producttype` WHERE idProductType='" + strconv.Itoa(i+1) + "';")
		defer pbt_rows.Close()
		for pbt_rows.Next() {
			err = pbt_rows.Scan(&pbt.ProductType, &pbt.ProductName)
			if err != nil {
				log.Fatal(err)
			}
			var inProduct Product
			inProduct.ProductType = pbt.ProductType
			inProduct.ProductGroupType = Base
			outProduct := SelectProduct(inProduct)
			pbt.Products = outProduct
			for _, element := range outProduct {
				GetBaseProductByIDWithBDFromDB(element.ID)
			}
		}
		ret.pbt = append(ret.pbt, pbt)
	}

	return ret
}

func GetBaseProductByIDWithBDFromDB(ID int) *BaseProductWithBD {
	conn, _ := db.GetDB()

	var bpwbd BaseProductWithBD
	var inProduct Product
	inProduct.ID = ID
	p := SelectProduct(inProduct)
	bpwbd.Product = p[0]
	var dp []Product
	var bp Product
	log.Println("Product With ID for BD return " + strconv.Itoa(ID))
	if p[0].ProductGroupType == Base {
		bd_rows, _ := conn.Query("SELECT `idProduct` FROM  `commonwealthcocktails`.`derivedProduct` WHERE idBaseProduct='" + strconv.Itoa(p[0].ID) + "';")
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
			dp = append(dp, derivedProduct[0])
		}
		bpwbd.DerivedProducts = dp
	} else {
		bd_rows, _ := conn.Query("SELECT `idBaseProduct` FROM  `commonwealthcocktails`.`derivedProduct` WHERE idProduct='" + strconv.Itoa(p[0].ID) + "';")
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
	}
	return &bpwbd
}

//func SelectProductsByCocktail(cocktail Cocktail) []Product{
//SELECT * FROM commonwealthcocktails.product
//JOIN commonwealthcocktails.cocktailToProducts ON cocktailToProducts.idProduct=product.idProduct
//JOIN  commonwealthcocktails.cocktail ON cocktailToProducts.idCocktail=cocktail.idCocktail
//WHERE cocktail.idCocktail=2;
//}

func SelectProductByID(ID int) Product {
	var ret Product
	conn, _ := db.GetDB()

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
			ret.ProductType = ProductType(pt)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(ret.ID, ret.ProductName, int(ret.ProductType), ret.Description, ret.ImagePath, ret.Image)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
	}
	return ret
}

func SelectProductsByCocktailAndProductType(ID int, pt int) []Product {
	var ret []Product
	conn, _ := db.GetDB()

	var buffer bytes.Buffer
	var canQuery = false
	buffer.WriteString("SELECT product.idProduct, product.productName, product.productType, COALESCE(product.productDescription, ''), COALESCE(product.productImagePath, '')," +
		" COALESCE(product.productImage, '') FROM commonwealthcocktails.product" +
		" JOIN commonwealthcocktails.cocktailToProducts ON product.idProduct=cocktailToProducts.idProduct" +
		" JOIN commonwealthcocktails.cocktail ON cocktailToProducts.idCocktail=cocktail.idCocktail" +
		" WHERE cocktail.idCocktail=" + strconv.Itoa(ID) + " AND product.productType=" + strconv.Itoa(pt) + ";")
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
			var prod Product
			var desc string
			var pt int
			err := rows.Scan(&prod.ID, &prod.ProductName, &pt, &desc, &prod.ImagePath, &prod.Image)
			prod.Description = template.HTML(desc)
			prod.ProductType = ProductType(pt)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(prod.ID, prod.ProductName, int(prod.ProductType), prod.Description, prod.ImagePath, prod.Image)
			ret = append(ret, prod)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
	}
	return ret
}
