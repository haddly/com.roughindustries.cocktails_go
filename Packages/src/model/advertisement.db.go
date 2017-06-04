//model/advertisement.db.go
package model

import (
	"database/sql"
	"db"
	"log"
)

func InitAdvertisementTables() {
	conn, _ := db.GetDB()
	var temp string
	if err := conn.QueryRow("SHOW TABLES LIKE 'adtype';").Scan(&temp); err == nil {
		log.Println("adtype Table Exists")
	} else if err == sql.ErrNoRows {
		log.Println("Creating adtype Table")
		conn.Query("CREATE TABLE `commonwealthcocktails`.`adtype` (`idAdType` INT NOT NULL AUTO_INCREMENT,  PRIMARY KEY (`idAdType`));")
		conn.Query("ALTER TABLE `commonwealthcocktails`.`adtype` ADD COLUMN `adTypeName` VARCHAR(150) NOT NULL AFTER `idAdType`;")
		conn.Query("INSERT INTO `commonwealthcocktails`.`adtype` (`idAdType`, `adTypeName`) VALUES ('1', 'ProductAds');")
		conn.Query("INSERT INTO `commonwealthcocktails`.`adtype` (`idAdType`, `adTypeName`) VALUES ('2', 'CocktailAds');")
		conn.Query("INSERT INTO `commonwealthcocktails`.`adtype` (`idAdType`, `adTypeName`) VALUES ('3', 'ProductPageAds');")
		conn.Query("INSERT INTO `commonwealthcocktails`.`adtype` (`idAdType`, `adTypeName`) VALUES ('4', 'CocktailPageAds');")
	}

	if err := conn.QueryRow("SHOW TABLES LIKE 'advertisement';").Scan(&temp); err == nil {
		log.Println("advertisement Table Exists")
	} else if err == sql.ErrNoRows {
		log.Println("Creating advertisement Table")
		conn.Query("CREATE TABLE `commonwealthcocktails`.`advertisement` (`idAdvertisement` INT NOT NULL AUTO_INCREMENT,PRIMARY KEY (`idAdvertisement`));")
		conn.Query("ALTER TABLE `commonwealthcocktails`.`advertisement`" +
			"ADD COLUMN `advertiserCompany` VARCHAR(250) NULL AFTER `idAdvertisement`," +
			"ADD COLUMN `advertiserName` VARCHAR(250) NULL AFTER `advertiserCompany`," +
			"ADD COLUMN `advertiserLink` VARCHAR(250) NULL AFTER `advertiserName`," +
			"ADD COLUMN `largeHorSnippet` VARCHAR(500) NULL AFTER `advertiserLink`," +
			"ADD COLUMN `mediumHorSnippet` VARCHAR(500) NULL AFTER `largeHorSnippet`," +
			"ADD COLUMN `smallHorSnippet` VARCHAR(500) NULL AFTER `mediumHorSnippet`," +
			"ADD COLUMN `bannerAdSnippet` VARCHAR(500) NULL AFTER `smallHorSnippet`," +
			"ADD COLUMN `largeVertSnippet` VARCHAR(500) NULL AFTER `bannerAdSnippet`," +
			"ADD COLUMN `mediumVertSnippet` VARCHAR(500) NULL AFTER `largeVertSnippet`," +
			"ADD COLUMN `smallVertSnippet` VARCHAR(500) NULL AFTER `mediumVertSnippet`," +
			"ADD COLUMN `idAdType` INT NOT NULL, ADD CONSTRAINT idAdType_id_fk FOREIGN KEY(idAdType) REFERENCES adtype(idAdType)," +
			"ADD COLUMN `Page` VARCHAR(150) NULL AFTER `smallVertSnippet`;")
	}

	if err := conn.QueryRow("SHOW TABLES LIKE 'advertiser';").Scan(&temp); err == nil {
		log.Println("advertiser Table Exists")
	} else if err == sql.ErrNoRows {
		log.Println("Creating advertiser Table")
		conn.Query("CREATE TABLE `commonwealthcocktails`.`advertiser` (`idAdvertiser` INT NOT NULL AUTO_INCREMENT,PRIMARY KEY (`idAdvertiser`));")
		conn.Query("ALTER TABLE `commonwealthcocktails`.`advertiser`" +
			"ADD COLUMN `advertiserCompany` VARCHAR(250) NULL AFTER `idAdvertiser`," +
			"ADD COLUMN `advertiserName` VARCHAR(250) NULL AFTER `advertiserCompany`," +
			"ADD COLUMN `advertiserLink` VARCHAR(250) NULL AFTER `advertiserName`;")
	}
	conn = nil
}

func InitAdvertisementReferences() {
	addAdToCocktailsTables()
	addAdProductToProductTables()
	addAdToPostsTables()
}

func addAdToCocktailsTables() {
	conn, _ := db.GetDB()
	var temp string
	if err := conn.QueryRow("SHOW TABLES LIKE 'adToCocktails';").Scan(&temp); err == nil {
		log.Println("adToCocktails Table Exists")
	} else if err == sql.ErrNoRows {
		log.Println("Creating adToCocktails Table")
		conn.Exec("CREATE TABLE `commonwealthcocktails`.`adToCocktails` (`idAdToCocktails` INT NOT NULL AUTO_INCREMENT," +
			" `idAdvertisement` INT NOT NULL," +
			" `idCocktail` INT NOT NULL," +
			" PRIMARY KEY (`idAdToCocktails`)," +
			" CONSTRAINT adToCocktails_idAdvertisement_id_fk FOREIGN KEY(idAdvertisement) REFERENCES advertisement(idAdvertisement)," +
			" CONSTRAINT adToCocktails_idCocktail_id_fk FOREIGN KEY(idCocktail) REFERENCES cocktail(idCocktail));")

	}
	conn = nil
}

func addAdProductToProductTables() {
	conn, _ := db.GetDB()
	var temp string
	if err := conn.QueryRow("SHOW TABLES LIKE 'adProductToProduct';").Scan(&temp); err == nil {
		log.Println("adProductToProduct Table Exists")
	} else if err == sql.ErrNoRows {
		log.Println("Creating adProductToProduct Table")
		//ad product, base product, advertiser, start date, end date, ....
		conn.Exec("CREATE TABLE `commonwealthcocktails`.`adProductToProduct` (`idAdProductToProduct` INT NOT NULL AUTO_INCREMENT," +
			" `idAdProduct` INT NOT NULL," +
			" `idProduct` INT NOT NULL," +
			" PRIMARY KEY (`idAdToPoducts`)," +
			" CONSTRAINT adProductToProduct_idAdProduct_id_fk FOREIGN KEY(idAdProduct) REFERENCES product(idProduct)," +
			" CONSTRAINT adProductToProduct_idProduct_id_fk FOREIGN KEY(idProduct) REFERENCES product(idProduct));")
	}
	conn = nil
}

func addAdToPostsTables() {
	conn, _ := db.GetDB()
	var temp string
	if err := conn.QueryRow("SHOW TABLES LIKE 'adToPosts';").Scan(&temp); err == nil {
		log.Println("adToPosts Table Exists")
	} else if err == sql.ErrNoRows {
		log.Println("Creating adToPosts Table")
		conn.Exec("CREATE TABLE `commonwealthcocktails`.`adToPosts` (`idAdToPosts` INT NOT NULL AUTO_INCREMENT," +
			" `idAdvertisement` INT NOT NULL," +
			" `idPost` INT NOT NULL," +
			" PRIMARY KEY (`idAdToPosts`)," +
			" CONSTRAINT adToPosts_idAdvertisement_id_fk FOREIGN KEY(idAdvertisement) REFERENCES advertisement(idAdvertisement)," +
			" CONSTRAINT adToPosts_idPost_id_fk FOREIGN KEY(idPost) REFERENCES post(idPost));")

	}
	conn = nil
}