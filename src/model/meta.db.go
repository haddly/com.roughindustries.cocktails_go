//model/meta.db.go
package model

import (
	"bytes"
	"database/sql"
	"db"
	"log"
	"strconv"
	"strings"
)

func InitMetaTables() {
	conn, _ := db.GetDB()
	var temp string
	if err := conn.QueryRow("SHOW TABLES LIKE 'metatype';").Scan(&temp); err == nil {
		log.Println("metatype Table Exists")
	} else if err == sql.ErrNoRows {
		log.Println("Creating metatype Table")
		conn.Query("CREATE TABLE `commonwealthcocktails`.`metatype` (`idMetaType` INT NOT NULL AUTO_INCREMENT,  PRIMARY KEY (`idMetaType`));")
		conn.Query("ALTER TABLE `commonwealthcocktails`.`metatype` ADD COLUMN `metaTypeName` VARCHAR(150) NOT NULL AFTER `idMetaType`;")
		conn.Exec("INSERT INTO `commonwealthcocktails`.`metatype` (`idMetaType`, `metaTypeName`) VALUES ('1', 'Flavor');")
		conn.Exec("INSERT INTO `commonwealthcocktails`.`metatype` (`idMetaType`, `metaTypeName`) VALUES ('2', 'Base Spirit');")
		conn.Exec("INSERT INTO `commonwealthcocktails`.`metatype` (`idMetaType`, `metaTypeName`) VALUES ('3', 'Type');")
		conn.Exec("INSERT INTO `commonwealthcocktails`.`metatype` (`idMetaType`, `metaTypeName`) VALUES ('4', 'Occasion');")
		conn.Exec("INSERT INTO `commonwealthcocktails`.`metatype` (`idMetaType`, `metaTypeName`) VALUES ('5', 'Family');")
		conn.Exec("INSERT INTO `commonwealthcocktails`.`metatype` (`idMetaType`, `metaTypeName`) VALUES ('6', 'Formula');")
		conn.Exec("INSERT INTO `commonwealthcocktails`.`metatype` (`idMetaType`, `metaTypeName`) VALUES ('7', 'Served');")
		conn.Exec("INSERT INTO `commonwealthcocktails`.`metatype` (`idMetaType`, `metaTypeName`) VALUES ('8', 'Technique');")
		conn.Exec("INSERT INTO `commonwealthcocktails`.`metatype` (`idMetaType`, `metaTypeName`) VALUES ('9', 'Strength');")
		conn.Exec("INSERT INTO `commonwealthcocktails`.`metatype` (`idMetaType`, `metaTypeName`) VALUES ('10', 'Difficulty');")
		conn.Exec("INSERT INTO `commonwealthcocktails`.`metatype` (`idMetaType`, `metaTypeName`) VALUES ('11', 'Time of Day');")
		conn.Exec("INSERT INTO `commonwealthcocktails`.`metatype` (`idMetaType`, `metaTypeName`) VALUES ('12', 'Ratio');")
		conn.Exec("INSERT INTO `commonwealthcocktails`.`metatype` (`idMetaType`, `metaTypeName`) VALUES ('13', 'Drink');")
	}

	if err := conn.QueryRow("SHOW TABLES LIKE 'meta';").Scan(&temp); err == nil {
		log.Println("Meta Table Exists")
	} else if err == sql.ErrNoRows {
		log.Println("Creating Meta Table")
		conn.Query("CREATE TABLE `commonwealthcocktails`.`meta` (`idMeta` INT NOT NULL AUTO_INCREMENT,PRIMARY KEY (`idMeta`));") //ID
		conn.Query("ALTER TABLE `commonwealthcocktails`.`meta`" +
			"ADD COLUMN `metaName` VARCHAR(150) NOT NULL AFTER `idMeta`," + //MetaName
			"ADD COLUMN `metaType`  INT NOT NULL AFTER `metaName`," + //MetaType
			"ADD COLUMN `metaArticle` INT AFTER `metaType`," + //Article
			"ADD COLUMN `metaBlurb` INT AFTER `metaArticle`;") //Blurb
	}
}

func InitMetaReferences() {
	conn, _ := db.GetDB()
	log.Println("Creating Meta References")
	conn.Query("ALTER TABLE `commonwealthcocktails`.`meta`" +
		" ADD CONSTRAINT meta_metatype_id_fk FOREIGN KEY(metaType) REFERENCES metatype(idMetaType)," +
		" ADD CONSTRAINT meta_metaarticle_id_fk FOREIGN KEY(metaArticle) REFERENCES post(idPost)," +
		" ADD CONSTRAINT meta_metablurb_id_fk FOREIGN KEY(metaBlurb) REFERENCES post(idPost);")
}

func ProcessMetas() {
	conn, _ := db.GetDB()

	for _, meta := range Metadata {
		log.Println(meta.MetaName)
		var buffer bytes.Buffer
		buffer.WriteString("INSERT INTO `commonwealthcocktails`.`meta` SET ")
		if meta.MetaName != "" {
			buffer.WriteString("`metaName`=\"" + meta.MetaName + "\",")
		}
		buffer.WriteString(" `metaType`=" + strconv.Itoa(int(meta.MetaType)) + ",")

		query := buffer.String()
		query = strings.TrimRight(query, ",")
		query = query + ";"
		log.Println(query)
		conn.Exec(query)
	}
}

func SelectMeta(meta Meta) Meta {
	var ret Meta
	conn, _ := db.GetDB()

	log.Println(meta.MetaName)
	var buffer bytes.Buffer
	buffer.WriteString("SELECT `idMeta`, `metaName`, `metaType` FROM `commonwealthcocktails`.`meta` WHERE ")
	if meta.MetaName != "" {
		buffer.WriteString("`metaName`=\"" + meta.MetaName + "\" AND")
	}
	buffer.WriteString(" `metaType`=" + strconv.Itoa(int(meta.MetaType)) + " AND")

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
		err := rows.Scan(&ret.ID, &ret.MetaName, &ret.MetaType)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(ret.ID, ret.MetaName, ret.MetaType)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return ret
}
