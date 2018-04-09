// Copyright 2017 Rough Industries LLC. All rights reserved.
//model/memcache.go:package model
package model

import (
	"bytes"
	"encoding/gob"
	"github.com/CommonwealthCocktails/connectors"
	"github.com/bradfitz/gomemcache/memcache"
	log "github.com/sirupsen/logrus"
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
func LoadAllMemcache(site string) {
	LoadMCWithProductData(site)
	LoadMCWithMetaData(site)
	LoadMCWithCocktailByAlphaNumsData(site)
	LoadMCWithCocktailData(site)
	//LoadMCWithPostData()
}

//Reinitialize the memcache with all the cocktails in alpha numeric order
func LoadMCWithCocktailByAlphaNumsData(site string) {
	mc, _ := connectors.GetMC()
	if mc != nil {
		cocktail := new(Cocktail)
		mc.Delete("cba")

		buf := new(bytes.Buffer)
		enc := gob.NewEncoder(buf)
		var cba CocktailsByAlphaNums
		cba = cocktail.SelectCocktailsByAlphaNums(true, site)
		enc.Encode(cba)

		mc.Set(&memcache.Item{Key: "cba", Value: buf.Bytes()})
	}
}

//Reinitialize the products lists in the memcache for both ingredients and
//non-ingredients, ingredients only, and non-ingredients only lists
func LoadMCWithProductData(site string) {
	mc, _ := connectors.GetMC()
	if mc != nil {
		product := new(Product)

		mc.Delete("pbt_tt")
		mc.Delete("pbt_tf")
		mc.Delete("pbt_ft")

		buf := new(bytes.Buffer)
		enc := gob.NewEncoder(buf)
		var pbt ProductsByTypes
		pbt = product.SelectProductsByTypes(true, true, true, site)
		enc.Encode(pbt)

		mc.Set(&memcache.Item{Key: "pbt_tt", Value: buf.Bytes()})

		buf = new(bytes.Buffer)
		enc = gob.NewEncoder(buf)
		pbt = product.SelectProductsByTypes(true, false, true, site)
		enc.Encode(pbt)

		mc.Set(&memcache.Item{Key: "pbt_tf", Value: buf.Bytes()})

		buf = new(bytes.Buffer)
		enc = gob.NewEncoder(buf)
		pbt = product.SelectProductsByTypes(false, true, true, site)
		enc.Encode(pbt)

		mc.Set(&memcache.Item{Key: "pbt_ft", Value: buf.Bytes()})
	}
}

//Reinitialize the meta lists in the memcache for both show in cocktail
//index ordered by metatype ordinal, just the list of meta values to show
//in the cocktail index with no order by, and all the metas and ordered by
//ordinal
func LoadMCWithMetaData(site string) {
	mc, _ := connectors.GetMC()
	if mc != nil {
		meta := new(Meta)

		mc.Delete("mbt_tt")
		mc.Delete("mbt_tf")
		mc.Delete("mbt_ft")

		buf := new(bytes.Buffer)
		enc := gob.NewEncoder(buf)
		var mbt MetasByTypes
		mbt = meta.SelectMetaByTypes(true, true, true, site)
		enc.Encode(mbt)

		mc.Set(&memcache.Item{Key: "mbt_tt", Value: buf.Bytes()})

		buf = new(bytes.Buffer)
		enc = gob.NewEncoder(buf)
		mbt = meta.SelectMetaByTypes(true, false, true, site)
		enc.Encode(mbt)

		mc.Set(&memcache.Item{Key: "mbt_tf", Value: buf.Bytes()})

		buf = new(bytes.Buffer)
		enc = gob.NewEncoder(buf)
		mbt = meta.SelectMetaByTypes(false, true, true, site)
		enc.Encode(mbt)

		mc.Set(&memcache.Item{Key: "mbt_ft", Value: buf.Bytes()})
	}
}

//Reinitialize the post lists in the memcache
func LoadMCWithPostData(site string) {
	mc, _ := connectors.GetMC()
	if mc != nil {
		post := new(Post)
		mc.Delete("posts")

		buf := new(bytes.Buffer)
		enc := gob.NewEncoder(buf)
		var p []Post
		p = post.SelectAllPosts(site)
		enc.Encode(p)

		mc.Set(&memcache.Item{Key: "posts", Value: buf.Bytes()})
	}
}

//Reinitialize the cocktail lists in the memcache
func LoadMCWithCocktailData(site string) {
	mc, _ := connectors.GetMC()
	if mc != nil {
		cocktail := new(Cocktail)
		//mc.Delete("cocktails")

		buf := new(bytes.Buffer)
		enc := gob.NewEncoder(buf)
		var c []Cocktail
		c = cocktail.SelectAllCocktails(site)
		c = c[0:25]
		for i, _ := range c {
			var cs CocktailSet
			cs = c[i].SelectCocktailsByID(c[i].ID, true, site)
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
		log.Infoln("\n\n\n")
		log.Infoln(cocktails)
		for i, _ := range c {
			log.Infoln(c[i].ID)
		}
	}
}
