// Copyright 2017 Rough Industries LLC. All rights reserved.
//model/meta.db.go:package model
package model

import (
	"bytes"
	"github.com/CommonwealthCocktails/connectors"
	"encoding/gob"
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"html"
	"html/template"
	"github.com/golang/glog"
	"strconv"
	"strings"
)

//CREATE, UPDATE, DELETE
//Insert a meta record into the database
func (meta *Meta) InsertMeta() int {
	//set the ID to zero to indicate an insert
	meta.ID = 0
	return meta.processMeta()
}

//Update a meta record in the database based on ID
func (meta *Meta) UpdateMeta() int {
	return meta.processMeta()
}

//Process an insert or an update
func (meta *Meta) processMeta() int {
	conn, _ := connectors.GetDB() //get db connection
	var args []interface{}        //arguments for variables in the data struct
	var buffer bytes.Buffer       //buffer for the query

	//If the ID is zero then do an insert else do an update based on the ID
	if meta.ID == 0 {
		buffer.WriteString("INSERT INTO `meta` ( ")
	} else {
		buffer.WriteString("UPDATE `meta` SET ")
	}

	//Append the correct columns to be added based on data available in the
	//data structure
	if meta.MetaName != "" {
		if meta.ID == 0 {
			buffer.WriteString("`metaName`,")
		} else {
			buffer.WriteString("`metaName`=?,")
		}
		args = append(args, html.EscapeString(meta.MetaName))
	}
	if meta.ID == 0 {
		buffer.WriteString("`metaBlurb`,")
	} else {
		buffer.WriteString(" `metaBlurb`=?,")
	}
	args = append(args, html.EscapeString(string(meta.Blurb)))
	metatype := meta.MetaType.SelectMetaType(true, true, true)
	if meta.ID == 0 {
		buffer.WriteString("`metaType`,")
	} else {
		buffer.WriteString(" `metaType`=?,")
	}
	args = append(args, strconv.Itoa(metatype[0].ID))

	//Cleanup the query and append where if it is an update
	query := buffer.String()
	query = strings.TrimRight(query, ",")
	if meta.ID == 0 {
		vals := strings.Repeat("?,", len(args))
		vals = strings.TrimRight(vals, ",")
		query = query + ") VALUES (" + vals + ");"
	} else {
		query = query + " WHERE `idMeta`=?;"
		args = append(args, meta.ID)
	}

	//Lets do this thing
	glog.Infoln(query)
	r, _ := conn.Exec(query, args...)
	id, _ := r.LastInsertId()
	ret := int(id)
	return ret
}

//SELECTS
//Select from the metatype table based on the attributes set in the metatype
//object.  Also ignores the showincocktailsindex column or ignoreHasRoot column
//or the ignoreIsOneToMany column based on the flags provided
func (metatype *MetaType) SelectMetaType(ignoreShowInCocktailsIndex bool, ignoreHasRoot bool, ignoreIsOneToMany bool) []MetaType {
	var ret []MetaType
	conn, _ := connectors.GetDB()
	var args []interface{} //arguments for variables in the data struct
	var buffer bytes.Buffer
	var canQuery = false
	buffer.WriteString("SELECT `idMetaType`, `metatypeName`, `metatypeOrdinal`, `metatypeShowInCocktailsIndex`, `metatypeHasRoot`, `metatypeIsOneToMany` FROM `metatype` WHERE ")
	if metatype.ID != 0 {
		buffer.WriteString("`idMetaType`=? AND")
		args = append(args, strconv.Itoa(metatype.ID))
		canQuery = true
	}
	if metatype.MetaTypeName != "" {
		buffer.WriteString("`metatypeName`=? AND")
		args = append(args, metatype.MetaTypeName)
		canQuery = true
	}
	if metatype.Ordinal != 0 {
		buffer.WriteString(" `metatypeOrdinal`=? AND")
		args = append(args, strconv.Itoa(metatype.Ordinal))
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
		glog.Infoln(query)
		rows, err := conn.Query(query, args...)
		if err != nil {
			glog.Error(err)
		}
		defer rows.Close()
		for rows.Next() {
			var metatype MetaType
			err := rows.Scan(&metatype.ID, &metatype.MetaTypeName, &metatype.Ordinal, &metatype.ShowInCocktailsIndex, &metatype.HasRoot, &metatype.IsOneToMany)
			if err != nil {
				glog.Error(err)
			}
			ret = append(ret, metatype)
			glog.Infoln(metatype.ID, metatype.MetaTypeName, metatype.Ordinal, metatype.ShowInCocktailsIndex, metatype.HasRoot, metatype.IsOneToMany)
		}
		err = rows.Err()
		if err != nil {
			glog.Error(err)
		}
	}
	return ret
}

//Select from the meta table based on the attributes set in the meta object.
func (meta *Meta) SelectMeta() []Meta {
	var ret []Meta
	conn, _ := connectors.GetDB()
	var args []interface{} //arguments for variables in the data struct
	var buffer bytes.Buffer
	buffer.WriteString("SELECT `idMeta`, `metaName`, `metaType`, COALESCE(`metaBlurb`, '') FROM `meta` WHERE ")
	if meta.ID != 0 {
		buffer.WriteString("`idMeta`=? AND")
		args = append(args, strconv.Itoa(meta.ID))
	}
	if meta.MetaName != "" {
		buffer.WriteString("`metaName`=? AND")
		args = append(args, meta.MetaName)
	}
	if meta.MetaType.ID != 0 {
		buffer.WriteString(" `metaType` IN (?) AND")
		args = append(args, strconv.Itoa(meta.MetaType.ID))
	}

	query := buffer.String()
	query = strings.TrimRight(query, " WHERE")
	query = strings.TrimRight(query, " AND")
	query = query + ";"
	glog.Infoln(query)
	rows, err := conn.Query(query, args...)
	if err != nil {
		glog.Error(err)
	}
	defer rows.Close()
	for rows.Next() {
		var meta Meta
		var blurb string
		err := rows.Scan(&meta.ID, &meta.MetaName, &meta.MetaType.ID, &blurb)
		if err != nil {
			glog.Error(err)
		}
		meta.Blurb = template.HTML(html.UnescapeString(blurb))
		ret = append(ret, meta)
		glog.Infoln(meta.ID, meta.MetaName, meta.MetaType.ID, meta.Blurb)
	}
	err = rows.Err()
	if err != nil {
		glog.Error(err)
	}

	return ret
}

//Select a set of meta records based on the flags passed in via the metatypes
//table. If ignore cache is true then the database query is run otherwise the
//cache is checked first.
func (meta *Meta) SelectMetaByTypes(byShowInCocktailsIndex bool, orderBy bool, ignoreCache bool) MetasByTypes {
	ret := new(MetasByTypes)
	ret = nil
	if !ignoreCache {
		ret = meta.memcachedMetaByTypes(byShowInCocktailsIndex, orderBy)
	}
	if ret == nil {
		ret = new(MetasByTypes)
		var mtList []int
		conn, _ := connectors.GetDB()
		query := "SELECT `idMetaType` FROM  `metatype`"
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
				glog.Infoln(item)
				mtList = append(mtList, item)
			}
		}
		//glog.Infoln("Meta Types Found " + strconv.Itoa(count))
		if rows != nil {
			rows.Close()
		}
		mtListString := strings.Trim(strings.Replace(fmt.Sprint(mtList), " ", ",", -1), "[]")
		//for _, i := range mtList {
		var mbt MetasByType
		mbt_rows, _ := conn.Query("SELECT `idMetaType`, `metatypeName`, `metatypeShowInCocktailsIndex`, `metatypeOrdinal`, `metatypeHasRoot`, `metatypeIsOneToMany` FROM  `metatype` WHERE idMetaType IN (" + mtListString + ");")
		defer mbt_rows.Close()
		for mbt_rows.Next() {
			err = mbt_rows.Scan(&mbt.MetaType.ID, &mbt.MetaType.MetaTypeName, &mbt.MetaType.ShowInCocktailsIndex, &mbt.MetaType.Ordinal, &mbt.MetaType.HasRoot, &mbt.MetaType.IsOneToMany)
			if err != nil {
				glog.Error(err)
			}
			mbt.MetaType.MetaTypeNameNoSpaces = strings.Join(strings.Fields(mbt.MetaType.MetaTypeName), "")
			var inMeta Meta
			inMeta.MetaType = mbt.MetaType
			outMeta := inMeta.SelectMeta()
			mbt.Metas = outMeta
			ret.MBT = append(ret.MBT, mbt)
		}

	}
	return *ret

}

//Memcache retrieval of metas by types
func (meta *Meta) memcachedMetaByTypes(byShowInCocktailsIndex bool, orderBy bool) *MetasByTypes {
	ret := new(MetasByTypes)
	ret = nil
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
	return ret
}

//Select a set of meta records based on a cocktail id nad metatype id
func (meta *Meta) SelectMetasByCocktailAndMetaType(ID int, mt int) ([]Meta, bool) {
	var ret []Meta
	var isRoot bool
	isRoot = false
	conn, _ := connectors.GetDB()
	var args []interface{} //arguments for variables in the data struct
	var buffer bytes.Buffer
	buffer.WriteString("SELECT meta.idMeta, meta.metaName, meta.metaType, cocktailToMetas.isRootCocktailForMeta" +
		" FROM meta" +
		" JOIN cocktailToMetas ON meta.idMeta=cocktailToMetas.idMeta" +
		" JOIN cocktail ON cocktailToMetas.idCocktail=cocktail.idCocktail" +
		" WHERE cocktail.idCocktail=? AND meta.metaType=?;")
	args = append(args, strconv.Itoa(ID))
	args = append(args, strconv.Itoa(mt))
	query := buffer.String()
	glog.Infoln(query)
	rows, err := conn.Query(query, args...)
	if err != nil {
		glog.Error(err)
	}
	defer rows.Close()
	for rows.Next() {
		var meta Meta
		var root string
		err := rows.Scan(&meta.ID, &meta.MetaName, &meta.MetaType.ID, &root)
		if err != nil {
			glog.Error(err)
		}
		if !isRoot {
			isRoot, _ = strconv.ParseBool(root)
		}
		ret = append(ret, meta)
		glog.Infoln(meta.ID, meta.MetaName, meta.MetaType.ID, isRoot)
	}
	err = rows.Err()
	if err != nil {
		glog.Error(err)
	}

	return ret, isRoot
}
