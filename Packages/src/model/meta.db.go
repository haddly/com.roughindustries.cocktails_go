//model/meta.connectors.go
package model

import (
	"bytes"
	"connectors"
	"database/sql"
	"encoding/gob"
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"html"
	"html/template"
	"log"
	"strconv"
	"strings"
)

func InitMetaTables() {
	conn, _ := connectors.GetDB()

	var temp string
	if err := conn.QueryRow("SHOW TABLES LIKE 'metatype';").Scan(&temp); err == nil {
		log.Println("metatype Table Exists")
	} else if err == sql.ErrNoRows {
		log.Println("Creating metatype Table")
		conn.Query("CREATE TABLE `commonwealthcocktails`.`metatype` (`idMetaType` INT NOT NULL AUTO_INCREMENT,  PRIMARY KEY (`idMetaType`));")
		conn.Query("ALTER TABLE `commonwealthcocktails`.`metatype`" +
			"ADD COLUMN `metatypeShowInCocktailsIndex` BOOLEAN AFTER `idMetaType`," + //ShowInCocktailsIndex
			"ADD COLUMN `metatypeName`  VARCHAR(150) NOT NULL AFTER `metatypeShowInCocktailsIndex`," + //IsSetType
			"ADD COLUMN `metatypeOrdinal` INT AFTER `metatypeName`," + //Ordinal
			"ADD COLUMN `metatypeHasRoot` TINYINT(1) NULL AFTER `metatypeOrdinal`," + //HasRoot
			"ADD COLUMN `metatypeIsOneToMany` TINYINT(1) NULL AFTER `metatypeHasRoot`;") //IsOneToMany
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
			"ADD COLUMN `metaBlurb` VARCHAR(500) AFTER `metaArticle`;") //Blurb
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
	conn, _ := connectors.GetDB()

	log.Println("Creating Meta References")
	conn.Query("ALTER TABLE `commonwealthcocktails`.`meta`" +
		" ADD CONSTRAINT meta_metatype_id_fk FOREIGN KEY(metaType) REFERENCES metatype(idMetaType)," +
		" ADD CONSTRAINT meta_metaarticle_id_fk FOREIGN KEY(metaArticle) REFERENCES post(idPost)," +
		" ADD CONSTRAINT meta_metablurb_id_fk FOREIGN KEY(metaBlurb) REFERENCES post(idPost);")
}

func ProcessMetaTypes() {
	conn, _ := connectors.GetDB()

	for _, metatype := range MetaTypes {
		log.Println(metatype.MetaTypeName)
		var buffer bytes.Buffer
		var args []interface{}
		buffer.WriteString("INSERT INTO `commonwealthcocktails`.`metatype` SET ")
		if metatype.MetaTypeName != "" {
			buffer.WriteString("`metatypeName`=?,")
			args = append(args, metatype.MetaTypeName)
		}
		if metatype.Ordinal != 0 {
			buffer.WriteString(" `metatypeOrdinal`=?,")
			args = append(args, metatype.Ordinal)
		}
		if metatype.ShowInCocktailsIndex {
			buffer.WriteString(" `metatypeShowInCocktailsIndex`=?,")
			args = append(args, 1)
		} else {
			buffer.WriteString(" `metatypeShowInCocktailsIndex`=?,")
			args = append(args, 0)
		}
		if metatype.HasRoot {
			buffer.WriteString(" `metatypeHasRoot`=?,")
			args = append(args, 1)
		} else {
			buffer.WriteString(" `metatypeHasRoot`=?,")
			args = append(args, 0)
		}
		if metatype.IsOneToMany {
			buffer.WriteString(" `metatypeIsOneToMany`=?,")
			args = append(args, 1)
		} else {
			buffer.WriteString(" `metatypeIsOneToMany`=?,")
			args = append(args, 0)
		}
		query := buffer.String()
		query = strings.TrimRight(query, ",")
		query = query + ";"
		log.Println(query)
		conn.Exec(query, args...)
	}
}

func ProcessMetas() {
	for _, meta := range Metadata {
		log.Println(meta.MetaName)
		meta.ID = 0
		ProcessMeta(meta)
	}
}

func ProcessMeta(meta Meta) int {
	conn, _ := connectors.GetDB()
	var args []interface{}

	var buffer bytes.Buffer
	if meta.ID == 0 {
		buffer.WriteString("INSERT INTO `commonwealthcocktails`.`meta` SET ")
	} else {
		buffer.WriteString("UPDATE `commonwealthcocktails`.`meta` SET ")
	}
	if meta.MetaName != "" {
		buffer.WriteString("`metaName`=?,")
		args = append(args, html.EscapeString(meta.MetaName))
	}
	if meta.Blurb != "" {
		buffer.WriteString(" `metaBlurb`=?,")
		args = append(args, html.EscapeString(string(meta.Blurb)))
	}
	metatype := SelectMetaType(meta.MetaType, true, true, true)
	buffer.WriteString(" `metaType`=?,")
	args = append(args, strconv.Itoa(metatype[0].ID))
	query := buffer.String()
	query = strings.TrimRight(query, ",")
	if meta.ID == 0 {
		query = query + ";"
	} else {
		query = query + " WHERE `idMeta`=?;"
		args = append(args, meta.ID)
	}

	log.Println(query)
	r, _ := conn.Exec(query, args...)
	id, _ := r.LastInsertId()
	ret := int(id)
	return ret
}

func InsertMeta(meta Meta) int {
	meta.ID = 0
	return ProcessMeta(meta)
}

func UpdateMeta(meta Meta) int {
	return ProcessMeta(meta)
}

func SelectMetaType(metatype MetaType, ignoreShowInCocktailsIndex bool, ignoreHasRoot bool, ignoreIsOneToMany bool) []MetaType {
	var ret []MetaType

	conn, _ := connectors.GetDB()

	//log.Println(metatype.MetaTypeName)
	var buffer bytes.Buffer
	var canQuery = false
	buffer.WriteString("SELECT `idMetaType`, `metatypeName`, `metatypeOrdinal`, `metatypeShowInCocktailsIndex`, `metatypeHasRoot`, `metatypeIsOneToMany` FROM `commonwealthcocktails`.`metatype` WHERE ")
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
	if !ignoreShowInCocktailsIndex {
		if metatype.ShowInCocktailsIndex {
			buffer.WriteString(" `metatypeShowInCocktailsIndex`='1' AND")
		} else {
			buffer.WriteString(" `metatypeShowInCocktailsIndex`='0' AND")
		}
	}

	if !ignoreHasRoot {
		if metatype.HasRoot {
			buffer.WriteString(" `metatypeHasRoot`='1' AND")
		} else {
			buffer.WriteString(" `metatypeHasRoot`='0' AND")
		}
	}

	if !ignoreIsOneToMany {
		if metatype.IsOneToMany {
			buffer.WriteString(" `metatypeIsOneToMany`='1' AND")
		} else {
			buffer.WriteString(" `metatypeIsOneToMany`='0' AND")
		}
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
			err := rows.Scan(&metatype.ID, &metatype.MetaTypeName, &metatype.Ordinal, &metatype.ShowInCocktailsIndex, &metatype.HasRoot, &metatype.IsOneToMany)
			if err != nil {
				log.Fatal(err)
			}
			ret = append(ret, metatype)
			log.Println(metatype.ID, metatype.MetaTypeName, metatype.Ordinal, metatype.ShowInCocktailsIndex, metatype.HasRoot, metatype.IsOneToMany)
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
	conn, _ := connectors.GetDB()

	log.Println(meta.MetaName)
	var buffer bytes.Buffer
	buffer.WriteString("SELECT `idMeta`, `metaName`, `metaType`, COALESCE(`metaBlurb`, '') FROM `commonwealthcocktails`.`meta` WHERE ")
	if meta.ID != 0 {
		buffer.WriteString("`idMeta`=" + strconv.Itoa(meta.ID) + " AND")
	}
	if meta.MetaName != "" {
		buffer.WriteString("`metaName`=\"" + meta.MetaName + "\" AND")
	}
	if meta.MetaType.ID != 0 {
		buffer.WriteString(" `metaType` IN (" + strconv.Itoa(meta.MetaType.ID) + ") AND")
	}

	query := buffer.String()
	query = strings.TrimRight(query, " WHERE")
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
		var blurb string
		err := rows.Scan(&meta.ID, &meta.MetaName, &meta.MetaType.ID, &blurb)
		if err != nil {
			log.Fatal(err)
		}
		meta.Blurb = template.HTML(blurb)
		ret = append(ret, meta)
		log.Println(meta.ID, meta.MetaName, meta.MetaType.ID, meta.Blurb)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return ret
}

func GetMetaByTypes(byShowInCocktailsIndex bool, orderBy bool, ignoreCache bool) MetasByTypes {
	ret := new(MetasByTypes)
	ret = nil
	if !ignoreCache {
		mc, _ := connectors.GetMC()
		if mc != nil {
			item := new(memcache.Item)
			item = nil
			if byShowInCocktailsIndex && orderBy {
				item, _ = mc.Get("mbt_tt")
			} else if byShowInCocktailsIndex && !orderBy {
				item, _ = mc.Get("mbt_tf")
			} else if !byShowInCocktailsIndex && orderBy {
				item, _ = mc.Get("mbt_ft")
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
		ret = new(MetasByTypes)
		var mtList []int
		conn, _ := connectors.GetDB()

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
		mtListString := strings.Trim(strings.Replace(fmt.Sprint(mtList), " ", ",", -1), "[]")
		//for _, i := range mtList {
		var mbt MetasByType
		mbt_rows, _ := conn.Query("SELECT `idMetaType`, `metatypeName`, `metatypeShowInCocktailsIndex`, `metatypeOrdinal`, `metatypeHasRoot`, `metatypeIsOneToMany` FROM  `commonwealthcocktails`.`metatype` WHERE idMetaType IN (" + mtListString + ");")
		defer mbt_rows.Close()
		for mbt_rows.Next() {
			err = mbt_rows.Scan(&mbt.MetaType.ID, &mbt.MetaType.MetaTypeName, &mbt.MetaType.ShowInCocktailsIndex, &mbt.MetaType.Ordinal, &mbt.MetaType.HasRoot, &mbt.MetaType.IsOneToMany)
			if err != nil {
				log.Fatal(err)
			}
			mbt.MetaType.MetaTypeNameNoSpaces = strings.Join(strings.Fields(mbt.MetaType.MetaTypeName), "")
			var inMeta Meta
			inMeta.MetaType = mbt.MetaType
			outMeta := SelectMeta(inMeta)
			mbt.Metas = outMeta
			ret.MBT = append(ret.MBT, mbt)
		}

	}
	return *ret

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
	conn, _ := connectors.GetDB()

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
