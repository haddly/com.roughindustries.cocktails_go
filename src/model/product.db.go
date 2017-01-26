//model/product.db.go
package model

import (
	"bytes"
	"database/sql"
	"db"
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

	if err := conn.QueryRow("SHOW TABLES LIKE 'bdgtype';").Scan(&temp); err == nil {
		log.Println("bdgtype Table Exists")
	} else if err == sql.ErrNoRows {
		log.Println("Creating bdgtype Table")
		conn.Query("CREATE TABLE `commonwealthcocktails`.`bdgtype` (`idBDGType` INT NOT NULL AUTO_INCREMENT,PRIMARY KEY (`idBDGType`));")
		conn.Query("ALTER TABLE `commonwealthcocktails`.`bdgtype` ADD COLUMN `bdgTypeName` VARCHAR(150) NOT NULL AFTER `idBDGType`;")
		conn.Exec("INSERT INTO `commonwealthcocktails`.`bdgtype` (`idBDGType`, `bdgTypeName`) VALUES ('1', 'Base');")
		conn.Exec("INSERT INTO `commonwealthcocktails`.`bdgtype` (`idBDGType`, `bdgTypeName`) VALUES ('2', 'Derived');")
		conn.Exec("INSERT INTO `commonwealthcocktails`.`bdgtype` (`idBDGType`, `bdgTypeName`) VALUES ('3', 'Group');")

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
			"ADD COLUMN `productBDGType` INT NOT NULL AFTER `productRecipe`," + //BDG
			"ADD COLUMN `productPreText` VARCHAR(250) AFTER `productBDGType`," + //PreText
			"ADD COLUMN `productPostText` VARCHAR(250) AFTER `productPreText`," + //PostText
			"ADD COLUMN `productRating` INT(1) AFTER `productPostText`," + //Rating
			"ADD COLUMN `productSourceName` VARCHAR(1500) AFTER `productRating`," + //SourceName
			"ADD COLUMN `productSourceLink` VARCHAR(1500) AFTER `productSourceName`," + //SourceLink
			"ADD COLUMN `productAbout` INT AFTER `productSourceLink`;") //About

	}
}

func InitProductReferences() {
	conn, _ := db.GetDB()
	conn.Query("ALTER TABLE `commonwealthcocktails`.`product`" +
		"ADD CONSTRAINT product_producttype_id_fk FOREIGN KEY(productType) REFERENCES producttype(idProductType)," +
		"ADD CONSTRAINT product_productarticle_id_fk FOREIGN KEY(productArticle) REFERENCES post(idPost)," +
		"ADD CONSTRAINT product_productrecipe_id_fk FOREIGN KEY(productRecipe) REFERENCES recipe(idRecipe)," +
		"ADD CONSTRAINT product_productbdgtype_id_fk FOREIGN KEY(productBDGType) REFERENCES bdgtype(idBDGType)," +
		"ADD CONSTRAINT product_productabout_id_fk FOREIGN KEY(productAbout) REFERENCES post(idPost);")
	addProductToMetasTables()
	addDerivedProductTables()
	addProductGroupTables()
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

func addProductGroupTables() {
	conn, _ := db.GetDB()
	var temp string
	if err := conn.QueryRow("SHOW TABLES LIKE 'groupProduct';").Scan(&temp); err == nil {
		log.Println("groupProduct Table Exists")
	} else if err == sql.ErrNoRows {
		log.Println("Creating groupProduct Table")
		//Drink
		conn.Exec("CREATE TABLE `commonwealthcocktails`.`groupProduct` (`idGroupProduct` INT NOT NULL AUTO_INCREMENT," +
			" `idProductToGroup` INT NOT NULL," +
			" `idProduct` INT NOT NULL," +
			" PRIMARY KEY (`idGroupProduct`)," +
			" CONSTRAINT groupProduct_idGroupProduct_id_fk FOREIGN KEY(idProductToGroup) REFERENCES product(idProduct)," +
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
		buffer.WriteString(" `productBDGType`=" + strconv.Itoa(int(product.BDG)) + ",")
		if product.Description != "" {
			sqlString := strings.Replace(string(product.Description), "\\", "\\\\", -1)
			sqlString = strings.Replace(sqlString, "\"", "\\\"", -1)
			buffer.WriteString("`productDescription`=\"" + sqlString + "\",")
		}
		if product.Details != "" {
			sqlString := strings.Replace(string(product.Description), "\\", "\\\\", -1)
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
