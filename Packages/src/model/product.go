//model/product.go
package model

import (
	"html/template"
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
	Errors           map[string]string
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

type BaseProductWithBDG struct {
	Product         Product
	DerivedProducts []Product
	GroupProducts   []Product
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
