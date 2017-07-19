//model/database.go
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
		pbt = GetProductsByTypes(true, true, true)
		enc.Encode(pbt)

		mc.Set(&memcache.Item{Key: "pbt_tt", Value: buf.Bytes()})

		buf = new(bytes.Buffer)
		enc = gob.NewEncoder(buf)
		pbt = GetProductsByTypes(true, false, true)
		enc.Encode(pbt)

		mc.Set(&memcache.Item{Key: "pbt_tf", Value: buf.Bytes()})

		buf = new(bytes.Buffer)
		enc = gob.NewEncoder(buf)
		pbt = GetProductsByTypes(false, true, true)
		enc.Encode(pbt)

		mc.Set(&memcache.Item{Key: "pbt_ft", Value: buf.Bytes()})

		// item, _ := mc.Get("pbt")
		// var ret ProductsByTypes
		// read := bytes.NewReader(item.Value)
		// dec := gob.NewDecoder(read)
		// dec.Decode(&ret)
		// log.Println(ret)
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
		mbt = GetMetaByTypes(true, true, true)
		enc.Encode(mbt)

		mc.Set(&memcache.Item{Key: "mbt_tt", Value: buf.Bytes()})

		buf = new(bytes.Buffer)
		enc = gob.NewEncoder(buf)
		mbt = GetMetaByTypes(true, false, true)
		enc.Encode(mbt)

		mc.Set(&memcache.Item{Key: "mbt_tf", Value: buf.Bytes()})

		buf = new(bytes.Buffer)
		enc = gob.NewEncoder(buf)
		mbt = GetMetaByTypes(false, true, true)
		enc.Encode(mbt)

		mc.Set(&memcache.Item{Key: "mbt_ft", Value: buf.Bytes()})
	}
}
