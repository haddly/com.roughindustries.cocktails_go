//model/product.go
package model

import (
	"html/template"
	"math/rand"
)

type ProductType struct {
	ID                      int
	IsIngredient            bool
	ProductTypeName         string
	ProductTypeNameNoSpaces string
}

type ProductTypeConst int

const (
	Spirit = 1 + iota
	Liqueur
	Wine
	Mixer
	Beer
	Garnish
	Drinkware
	Tool
)

var ProductTypeStrings = [...]string{
	"Spirit",
	"Liqueur",
	"Wine",
	"Mixer",
	"Beer",
	"Garnish",
	"Drinkware",
	"Tool",
}

// String returns the English name of the Producttype ("Spirit", "Liqueur", ...).
func (pt ProductTypeConst) String() string { return ProductTypeStrings[pt-1] }

func ProductTypeStringToInt(a string) int {
	var i = 1
	for _, b := range ProductTypeStrings {
		if b == a {
			return i
		}
		i++
	}
	return 0
}

type Product struct {
	ID               int
	ProductName      string
	ProductType      ProductType
	Description      template.HTML
	Details          template.HTML
	ImagePath        string
	Image            string
	ImageSourceName  string
	ImageSourceLink  string
	Article          Post
	Recipe           Recipe
	ProductGroupType GroupType
	PreText          string
	PostText         string
	Drink            []Meta
	Rating           int
	SourceName       string
	SourceLink       string
	About            Post
	//Advertiser Info
	Advertisement Advertisement
}

type DerivedProduct struct {
	Product     Product
	BaseProduct Product
}

type GroupProduct struct {
	Products     []Product
	GroupProduct Product
}
type BaseProductWithBD struct {
	Product         Product
	DerivedProducts []Product
	BaseProduct     Product
}

type ProductsByTypes struct {
	PBT []ProductsByType
}

type ProductsByType struct {
	ProductType ProductType
	Products    []Product
}

func GetProducts() []Product {
	var p []Product
	p = Products
	return p
}

func GetBaseProductWithBD() *BaseProductWithBD {
	p := Products[rand.Intn(len(Products))]
	ID := p.ID
	var bpwbd BaseProductWithBD
	bpwbd.Product = p
	var dp []Product
	var bp Product
	if p.ProductGroupType == Base {
		for _, dps_element := range DerivedProducts {
			if dps_element.BaseProduct.ID == ID {
				dp = append(dp, dps_element.Product)
			}
		}
		bpwbd.DerivedProducts = dp
	} else {
		for _, dps_element := range DerivedProducts {
			if dps_element.Product.ID == ID {
				bp = dps_element.BaseProduct
				break
			}
		}
		bpwbd.BaseProduct = bp
	}
	return &bpwbd
}

func GetBaseProductByIDWithBD(ID int) *BaseProductWithBD {
	p := Products[ID-1]
	var bpwbd BaseProductWithBD
	bpwbd.Product = p
	var dp []Product
	var bp Product
	if p.ProductGroupType == Base {
		for _, dps_element := range DerivedProducts {
			if dps_element.BaseProduct.ID == ID {
				dp = append(dp, dps_element.Product)
			}
		}
		bpwbd.DerivedProducts = dp
	} else {
		for _, dps_element := range DerivedProducts {
			if dps_element.Product.ID == ID {
				bp = dps_element.BaseProduct
				break
			}
		}
		bpwbd.BaseProduct = bp
	}

	//need to check the ad type
	for ad_index, ad_element := range Advertisements {
		if ad_element.AdType == ProductAds {
			for _, adprodcuts_element := range ad_element.Products {
				if p.ID == adprodcuts_element.BaseProduct.ID {
					bpwbd.Product.Advertisement = Advertisements[ad_index]
				}
			}
		}
	}

	return &bpwbd
}
