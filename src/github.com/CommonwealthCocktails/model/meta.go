// Copyright 2017 Rough Industries LLC. All rights reserved.
//model/meta.go:package model
package model

import (
	"html/template"
)

//DATA STRUCTURES
//Metatype data structure
type MetaType struct {
	ID                   int
	ShowInCocktailsIndex bool
	MetaTypeName         string
	MetaTypeNameNoSpaces string
	Ordinal              int
	HasRoot              bool
	IsOneToMany          bool
}

//Meta data structure
type Meta struct {
	ID       int
	MetaName string
	MetaType MetaType
	Blurb    template.HTML
	Errors   map[string]string
}

//A list of a set of metas grouped by metatype
type MetasByTypes struct {
	MBT []MetasByType
}

//A set of metas for a metatype
type MetasByType struct {
	MetaType MetaType
	Metas    []Meta
}

//ENUMERATIONS - These must match the database one for one in both ID and order
//The integer values for the metatype enumeration
type MetaTypesConst int

const (
	Flavor = 1 + iota
	BaseSpirit
	Type
	Occasion
	Family
	Formula
	Served
	Technique
	Strength
	Difficulty
	TOD
	Ratio
	Drink
)

//The string values for the metatype enumeration
//family can be used to duplicate liquor.coms pages like
//https://www.liquor.com/mosaic/margarita-recipes/#gs.0BEty3o
//it can include a post with a excerpt and multiple base
//cocktails that are shown first, i.e. frozen margarita and
//margarita on the rocks.  Then derived margaritas.  Mules,
//Daiquiris, martinis, bloddy marys, are all examples of this
//idea
var MetaTypeStrings = [...]string{
	"Flavor",      //Fruity, Bitter, Creamy, ...
	"Base Spirit", //Vodka, Gin, Bourbon, ...
	"Type",        //Tiki,
	"Occasion",    //Christmas, 4th of July, Halloween, ...
	"Family",      //Margarita, Martini, Sour, ...
	"Formula",
	"Served",      //Highball, Martini, Old Fashioned, ...
	"Technique",   //Shaking, Stirring, Straining
	"Strength",    //Weak, Medium, Strong
	"Difficulty",  //Easy, Medium, Hard
	"Time of Day", //Evening, Dessert, Brunch, ...
	"Ratio",       //3:4:8, ...
	"Drink",       //On The Rocks, Neat, ...
}

// String returns the English name of the metatype ("Flavor", "Base Spirit", ...).
func (mt MetaTypesConst) String() string { return MetaTypeStrings[mt-1] }

//The integer values for the grouptype enumeration
type GroupType int

const (
	Base = 1 + iota
	Derived
	Group
	Substitute
)

//The string values for the grouptype enumeration
var GroupTypeStrings = [...]string{
	"Base",
	"Derived",
	"Group",
	"Substitute",
}

// String returns the English name of the Grouptype ("Base", "Derived", ...).
func (ct GroupType) String() string { return GroupTypeStrings[ct-1] }

//Helper function for templates so that they can take a grouptype string
//and convert it to the int const
func GroupTypeStringToInt(a string) int {
	var i = 1
	for _, b := range GroupTypeStrings {
		if b == a {
			return i
		}
		i++
	}
	return 0
}
