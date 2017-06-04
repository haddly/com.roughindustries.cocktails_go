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
		conn.Query("ALTER TABLE `commonwealthcocktails`.`metatype`" +
			"ADD COLUMN `metatypeShowInCocktailsIndex` BOOLEAN AFTER `idMetaType`," + //ShowInCocktailsIndex
			"ADD COLUMN `metatypeName`  VARCHAR(150) NOT NULL AFTER `metatypeShowInCocktailsIndex`," + //MetaTypeName
			"ADD COLUMN `metatypeOrdinal` INT AFTER `metatypeName`;") //Ordinal
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

	if err := conn.QueryRow("SHOW TABLES LIKE 'grouptype';").Scan(&temp); err == nil {
		log.Println("grouptype Table Exists")
	} else if err == sql.ErrNoRows {
		log.Println("Creating grouptype Table")
		conn.Query("CREATE TABLE `commonwealthcocktails`.`grouptype` (`idGroupType` INT NOT NULL AUTO_INCREMENT,PRIMARY KEY (`idGroupType`));")
		conn.Query("ALTER TABLE `commonwealthcocktails`.`grouptype` ADD COLUMN `groupTypeName` VARCHAR(150) NOT NULL AFTER `idGroupType`;")
		conn.Exec("INSERT INTO `commonwealthcocktails`.`grouptype` (`idGroupType`, `groupTypeName`) VALUES ('1', 'Base');")
		conn.Exec("INSERT INTO `commonwealthcocktails`.`grouptype` (`idGroupType`, `groupTypeName`) VALUES ('2', 'Derived');")
		conn.Exec("INSERT INTO `commonwealthcocktails`.`grouptype` (`idGroupType`, `groupTypeName`) VALUES ('3', 'Group');")
		conn.Exec("INSERT INTO `commonwealthcocktails`.`grouptype` (`idGroupType`, `groupTypeName`) VALUES ('4', 'Substitute');")
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

func ProcessMetaTypes() {
	conn, _ := db.GetDB()

	for _, metatype := range MetaTypes {
		log.Println(metatype.MetaTypeName)
		var buffer bytes.Buffer
		buffer.WriteString("INSERT INTO `commonwealthcocktails`.`metatype` SET ")
		if metatype.MetaTypeName != "" {
			buffer.WriteString("`metatypeName`=\"" + metatype.MetaTypeName + "\",")
		}
		if metatype.Ordinal != 0 {
			buffer.WriteString(" `metatypeOrdinal`=" + strconv.Itoa(metatype.Ordinal) + ",")
		}
		if metatype.ShowInCocktailsIndex {
			buffer.WriteString(" `metatypeShowInCocktailsIndex`='1',")
		} else {
			buffer.WriteString(" `metatypeShowInCocktailsIndex`='0',")
		}
		query := buffer.String()
		query = strings.TrimRight(query, ",")
		query = query + ";"
		log.Println(query)
		conn.Exec(query)
	}
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
		//TO DO: Select by MetaType to get ID
		metatype := SelectMetaType(meta.MetaType)
		buffer.WriteString(" `metaType`=" + strconv.Itoa(metatype[0].ID) + ",")

		query := buffer.String()
		query = strings.TrimRight(query, ",")
		query = query + ";"
		log.Println(query)
		conn.Exec(query)
	}
}

func SelectMetaType(metatype MetaType) []MetaType {
	var ret []MetaType

	conn, _ := db.GetDB()

	log.Println(metatype.MetaTypeName)
	var buffer bytes.Buffer
	var canQuery = false
	buffer.WriteString("SELECT `idMetaType`, `metatypeName`, `metatypeOrdinal`, `metatypeShowInCocktailsIndex` FROM `commonwealthcocktails`.`metatype` WHERE ")
	if metatype.ID != 0 {
		buffer.WriteString("`idMetaType`=" + strconv.Itoa(metatype.ID) + " AND")
		canQuery = true
	}
	if metatype.MetaTypeName != "" {
		buffer.WriteString("`metatypeName`=\"" + metatype.MetaTypeName + "\" AND")
		canQuery = true
	}
	if metatype.Ordinal != 0 {
		buffer.WriteString(" `metatypeOrdinal`=" + strconv.Itoa(metatype.Ordinal) + " AND")
		canQuery = true
	}
	if metatype.ShowInCocktailsIndex {
		buffer.WriteString(" `metatypeShowInCocktailsIndex`='1' AND")
	} else {
		buffer.WriteString(" `metatypeShowInCocktailsIndex`='0' AND")
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
			var metatype MetaType
			err := rows.Scan(&metatype.ID, &metatype.MetaTypeName, &metatype.Ordinal, &metatype.ShowInCocktailsIndex)
			if err != nil {
				log.Fatal(err)
			}
			ret = append(ret, metatype)
			log.Println(metatype.ID, metatype.MetaTypeName, metatype.Ordinal, metatype.ShowInCocktailsIndex)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
	}
	return ret
}

func SelectMeta(meta Meta) []Meta {
	var ret []Meta
	conn, _ := db.GetDB()

	log.Println(meta.MetaName)
	var buffer bytes.Buffer
	var canQuery = false
	buffer.WriteString("SELECT `idMeta`, `metaName`, `metaType` FROM `commonwealthcocktails`.`meta` WHERE ")
	if meta.ID != 0 {
		buffer.WriteString("`idMeta`=" + strconv.Itoa(meta.ID) + " AND")
		canQuery = true
	}
	if meta.MetaName != "" {
		buffer.WriteString("`metaName`=\"" + meta.MetaName + "\" AND")
		canQuery = true
	}
	if meta.MetaType.ID != 0 {
		buffer.WriteString(" `metaType`=" + strconv.Itoa(meta.MetaType.ID) + " AND")
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
			var meta Meta
			err := rows.Scan(&meta.ID, &meta.MetaName, &meta.MetaType.ID)
			if err != nil {
				log.Fatal(err)
			}
			ret = append(ret, meta)
			log.Println(meta.ID, meta.MetaName, meta.MetaType.ID)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
	}
	return ret
}

func GetMetaByTypes(byShowInCocktailsIndex bool, orderBy bool) MetasByTypes {
	var ret MetasByTypes
	var mtList []int
	conn, _ := db.GetDB()

	query := "SELECT `idMetaType` FROM  `commonwealthcocktails`.`metatype`"

	if byShowInCocktailsIndex {
		query += " WHERE metatypeShowInCocktailsIndex=1"
	}
	if orderBy {
		query += " ORDER BY metatypeOrdinal"
	}
	query += ";"
	rows, _ := conn.Query(query)
	//var count int
	var err error
	if rows == nil {
		//count = 0
	} else {
		//count, err = checkCount(rows)
		for rows.Next() {
			var item int
			rows.Scan(&item)
			log.Println(item)
			mtList = append(mtList, item)
		}
	}
	//log.Println("Meta Types Found " + strconv.Itoa(count))
	if rows != nil {
		rows.Close()
	}
	for _, i := range mtList {
		var mbt MetasByType
		mbt_rows, _ := conn.Query("SELECT `idMetaType`, `metatypeName`, `metatypeShowInCocktailsIndex`, `metatypeOrdinal` FROM  `commonwealthcocktails`.`metatype` WHERE idMetaType='" + strconv.Itoa(i) + "';")
		defer mbt_rows.Close()
		for mbt_rows.Next() {
			err = mbt_rows.Scan(&mbt.MetaType.ID, &mbt.MetaType.MetaTypeName, &mbt.MetaType.ShowInCocktailsIndex, &mbt.MetaType.Ordinal)
			if err != nil {
				log.Fatal(err)
			}
			var inMeta Meta
			inMeta.MetaType = mbt.MetaType
			outMeta := SelectMeta(inMeta)
			mbt.Metas = outMeta
		}
		ret.MBT = append(ret.MBT, mbt)
	}
	return ret

}

func checkCount(rows *sql.Rows) (count int, err error) {
	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			log.Fatal(err)
			return 0, err
		}
	}
	return count, nil
}

func SelectMetasByCocktailAndMetaType(ID int, mt int) []Meta {
	var ret []Meta
	conn, _ := db.GetDB()

	var buffer bytes.Buffer
	var canQuery = false
	buffer.WriteString("SELECT meta.idMeta, meta.metaName, meta.metaType" +
		" FROM commonwealthcocktails.meta" +
		" JOIN commonwealthcocktails.cocktailToMetas ON meta.idMeta=cocktailToMetas.idMeta" +
		" JOIN commonwealthcocktails.cocktail ON cocktailToMetas.idCocktail=cocktail.idCocktail" +
		" WHERE cocktail.idCocktail=" + strconv.Itoa(ID) + " AND meta.metaType=" + strconv.Itoa(mt) + "")
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
			var meta Meta
			err := rows.Scan(&meta.ID, &meta.MetaName, &meta.MetaType.ID)
			if err != nil {
				log.Fatal(err)
			}
			ret = append(ret, meta)
			log.Println(meta.ID, meta.MetaName, meta.MetaType.ID)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
	}
	return ret
}
