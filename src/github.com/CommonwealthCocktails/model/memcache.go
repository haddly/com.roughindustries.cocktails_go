// Copyright 2017 Rough Industries LLC. All rights reserved.
//model/memcache.go:package model
package model

import (
	"bytes"
	"encoding/gob"
	"github.com/CommonwealthCocktails/connectors"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/golang/glog"
	"strconv"
)

//Clears out the whole memcache
func DeleteAllMemcache() {
	mc, _ := connectors.GetMC()
	if mc != nil {
		mc.DeleteAll()
	}
}

//Calls all the loaders to reinitialize the memcache
func LoadAllMemcache() {
	LoadMCWithProductData()
	LoadMCWithMetaData()
	LoadMCWithCocktailByAlphaNumsData()
	LoadMCWithCocktailData()
	//LoadMCWithPostData()
}

//Reinitialize the memcache with all the cocktails in alpha numeric order
func LoadMCWithCocktailByAlphaNumsData() {
	mc, _ := connectors.GetMC()
	if mc != nil {
		cocktail := new(Cocktail)
		mc.Delete("cba")

		buf := new(bytes.Buffer)
		enc := gob.NewEncoder(buf)
		var cba CocktailsByAlphaNums
		cba = cocktail.SelectCocktailsByAlphaNums(true)
		enc.Encode(cba)

		mc.Set(&memcache.Item{Key: "cba", Value: buf.Bytes()})
	}
}

//Reinitialize the products lists in the memcache for both ingredients and
//non-ingredients, ingredients only, and non-ingredients only lists
func LoadMCWithProductData() {
	mc, _ := connectors.GetMC()
	if mc != nil {
		product := new(Product)

		mc.Delete("pbt_tt")
		mc.Delete("pbt_tf")
		mc.Delete("pbt_ft")

		buf := new(bytes.Buffer)
		enc := gob.NewEncoder(buf)
		var pbt ProductsByTypes
		pbt = product.SelectProductsByTypes(true, true, true)
		enc.Encode(pbt)

		mc.Set(&memcache.Item{Key: "pbt_tt", Value: buf.Bytes()})

		buf = new(bytes.Buffer)
		enc = gob.NewEncoder(buf)
		pbt = product.SelectProductsByTypes(true, false, true)
		enc.Encode(pbt)

		mc.Set(&memcache.Item{Key: "pbt_tf", Value: buf.Bytes()})

		buf = new(bytes.Buffer)
		enc = gob.NewEncoder(buf)
		pbt = product.SelectProductsByTypes(false, true, true)
		enc.Encode(pbt)

		mc.Set(&memcache.Item{Key: "pbt_ft", Value: buf.Bytes()})
	}
}

//Reinitialize the meta lists in the memcache for both show in cocktail
//index ordered by metatype ordinal, just the list of meta values to show
//in the cocktail index with no order by, and all the metas and ordered by
//ordinal
func LoadMCWithMetaData() {
	mc, _ := connectors.GetMC()
	if mc != nil {
		meta := new(Meta)

		mc.Delete("mbt_tt")
		mc.Delete("mbt_tf")
		mc.Delete("mbt_ft")

		buf := new(bytes.Buffer)
		enc := gob.NewEncoder(buf)
		var mbt MetasByTypes
		mbt = meta.SelectMetaByTypes(true, true, true)
		enc.Encode(mbt)

		mc.Set(&memcache.Item{Key: "mbt_tt", Value: buf.Bytes()})

		buf = new(bytes.Buffer)
		enc = gob.NewEncoder(buf)
		mbt = meta.SelectMetaByTypes(true, false, true)
		enc.Encode(mbt)

		mc.Set(&memcache.Item{Key: "mbt_tf", Value: buf.Bytes()})

		buf = new(bytes.Buffer)
		enc = gob.NewEncoder(buf)
		mbt = meta.SelectMetaByTypes(false, true, true)
		enc.Encode(mbt)

		mc.Set(&memcache.Item{Key: "mbt_ft", Value: buf.Bytes()})
	}
}

//Reinitialize the post lists in the memcache
func LoadMCWithPostData() {
	mc, _ := connectors.GetMC()
	if mc != nil {
		post := new(Post)
		mc.Delete("posts")

		buf := new(bytes.Buffer)
		enc := gob.NewEncoder(buf)
		var p []Post
		p = post.SelectAllPosts()
		enc.Encode(p)

		mc.Set(&memcache.Item{Key: "posts", Value: buf.Bytes()})
	}
}

//Reinitialize the cocktail lists in the memcache
func LoadMCWithCocktailData() {
	mc, _ := connectors.GetMC()
	if mc != nil {
		cocktail := new(Cocktail)
		//mc.Delete("cocktails")

		buf := new(bytes.Buffer)
		enc := gob.NewEncoder(buf)
		var c []Cocktail
		c = cocktail.SelectAllCocktails(false)
		c = c[0:25]
		for i, _ := range c {
			var cs CocktailSet
			cs = c[i].SelectCocktailsByID(c[i].ID, true)
			c[i] = cs.Cocktail
			var temp []string
			for _, e := range c[i].Recipe.RecipeSteps {
				temp = append(temp, strconv.Itoa(e.OriginalIngredient.ID))
			}
			//c[i].RelatedCocktails = c[i].SelectCocktailsByIngredientIDs(temp)
			enc.Encode(c[i])
			mc.Set(&memcache.Item{Key: "cocktail_" + strconv.Itoa(c[i].ID), Value: buf.Bytes()})
		}
		var cocktails Cocktail
		//cocktails.List = c
		// enc.Encode(cocktails)
		// mc.Set(&memcache.Item{Key: "cocktails", Value: buf.Bytes()})

		item := new(memcache.Item)
		item, _ = mc.Get("cocktail_6")
		read := bytes.NewReader(item.Value)
		dec := gob.NewDecoder(read)
		dec.Decode(&cocktails)
		glog.Infoln("\n\n\n")
		glog.Infoln(cocktails)
		for i, _ := range c {
			glog.Infoln(c[i].ID)
		}
	}
}
