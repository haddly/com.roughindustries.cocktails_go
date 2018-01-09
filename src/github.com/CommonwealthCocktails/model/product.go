// Copyright 2017 Rough Industries LLC. All rights reserved.
//model/product.go:package model
package model

import (
	"html/template"
)

//DATA STRUCTURES
//Product data structure
type Product struct {
	ID               int
	ProductName      template.HTML
	ProductType      ProductType
	Description      template.HTML
	Details          template.HTML
	ImagePath        string
	Image            string
	LabeledImageLink string
	ImageSourceName  string
	ImageSourceLink  string
	Recipe           Recipe
	ProductGroupType GroupType
	PreText          string
	PostText         string
	Drink            []Meta
	Rating           int
	SourceName       string
	SourceLink       string
	Errors           map[string]string
	//Affiliates
	AmazonLink string
}

//Data struct to represent a product type.  IsIngredient determines if the
//product is an ingredent in cocktails versus a garnish or say drinkware.
type ProductType struct {
	ID                      int
	IsIngredient            bool
	ProductTypeName         string
	ProductTypeNameNoSpaces string
}

//A derived product has a base product that the product is derived from.
//This is a one to one relationaship.
type DerivedProduct struct {
	Product     Product
	BaseProduct Product
}

//A group product has a group base product that the products are apart of.
//This is a many to one relationaship.
type GroupProduct struct {
	Products     []Product
	GroupProduct Product
}

//The BDG data struct relates a product to either derived products, group
//products or a base product.  It is a full picture of the product with
//respect to other products it is associated to
type BaseProductWithBDG struct {
	Product         Product
	DerivedProducts []Product
	GroupProducts   []Product
	BaseProduct     Product
}

//A set of products based on the product types they are apart of.
type ProductsByTypes struct {
	PBT []ProductsByType
}

//A product type that has the products that are related to it in a data
//structure
type ProductsByType struct {
	ProductType ProductType
	Products    []Product
}

//ENUMERATIONS - These must match the database one for one in both ID and order
//The integer values for the producttype enumeration
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

//The string values for the producttype enumeration
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

//Helper function to convert a product type string to it's int value.
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
