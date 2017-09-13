// Copyright 2017 Rough Industries LLC. All rights reserved.
//model/memcache.go:package model
package model

import (
	"bytes"
	"connectors"
	"encoding/gob"
	"github.com/bradfitz/gomemcache/memcache"
)

func DeleteAllMemcache() {
	mc, _ := connectors.GetMC()
	if mc != nil {
		mc.DeleteAll()
	}
}

func LoadAllMemcache() {
	LoadMCWithProductData()
	LoadMCWithMetaData()
	LoadMCWithCocktailByAlphaNumsData()
}

func LoadMCWithCocktailByAlphaNumsData() {
	mc, _ := connectors.GetMC()
	if mc != nil {
		mc.Delete("cba")

		buf := new(bytes.Buffer)
		enc := gob.NewEncoder(buf)
		var cba CocktailsByAlphaNums
		cba = GetCocktailsByAlphaNums(true)
		enc.Encode(cba)

		mc.Set(&memcache.Item{Key: "cba", Value: buf.Bytes()})
	}
}

func LoadMCWithProductData() {
	mc, _ := connectors.GetMC()
	if mc != nil {
		mc.Delete("pbt_tt")
		mc.Delete("pbt_tf")
		mc.Delete("pbt_ft")

		buf := new(bytes.Buffer)
		enc := gob.NewEncoder(buf)
		var pbt ProductsByTypes
		pbt = SelectProductsByTypes(true, true, true)
		enc.Encode(pbt)

		mc.Set(&memcache.Item{Key: "pbt_tt", Value: buf.Bytes()})

		buf = new(bytes.Buffer)
		enc = gob.NewEncoder(buf)
		pbt = SelectProductsByTypes(true, false, true)
		enc.Encode(pbt)

		mc.Set(&memcache.Item{Key: "pbt_tf", Value: buf.Bytes()})

		buf = new(bytes.Buffer)
		enc = gob.NewEncoder(buf)
		pbt = SelectProductsByTypes(false, true, true)
		enc.Encode(pbt)

		mc.Set(&memcache.Item{Key: "pbt_ft", Value: buf.Bytes()})
	}
}

func LoadMCWithMetaData() {
	mc, _ := connectors.GetMC()
	if mc != nil {
		mc.Delete("mbt_tt")
		mc.Delete("mbt_tf")
		mc.Delete("mbt_ft")

		buf := new(bytes.Buffer)
		enc := gob.NewEncoder(buf)
		var mbt MetasByTypes
		mbt = SelectMetaByTypes(true, true, true)
		enc.Encode(mbt)

		mc.Set(&memcache.Item{Key: "mbt_tt", Value: buf.Bytes()})

		buf = new(bytes.Buffer)
		enc = gob.NewEncoder(buf)
		mbt = SelectMetaByTypes(true, false, true)
		enc.Encode(mbt)

		mc.Set(&memcache.Item{Key: "mbt_tf", Value: buf.Bytes()})

		buf = new(bytes.Buffer)
		enc = gob.NewEncoder(buf)
		mbt = SelectMetaByTypes(false, true, true)
		enc.Encode(mbt)

		mc.Set(&memcache.Item{Key: "mbt_ft", Value: buf.Bytes()})
	}
}
