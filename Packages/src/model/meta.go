//model/meta.go
package model

import (
	"html/template"
)

type MetaType struct {
	ID                   int
	ShowInCocktailsIndex bool
	MetaTypeName         string
	MetaTypeNameNoSpaces string
	Ordinal              int
	HasRoot              bool
	IsOneToMany          bool
}

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

var MetaTypeCount = 13

//family can be used to duplicate liquor.coms pages likr
//http://www.liquor.com/mosaic/margarita-recipes/#gs.0BEty3o
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

type GroupType int

const (
	Base = 1 + iota
	Derived
	Group
	Substitute
)

var GroupTypeStrings = [...]string{
	"Base",
	"Derived",
	"Group",
	"Substitute",
}

// String returns the English name of the Grouptype ("Base", "Derived", ...).
func (ct GroupType) String() string { return GroupTypeStrings[ct-1] }

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

type Meta struct {
	ID       int               //valid string in validator map key is metaID
	MetaName string            //valid string in validator map key is metaName
	MetaType MetaType          //valid string in validator map key is metaType
	Article  Post              //TBD
	Blurb    template.HTML     //valid string in validator map key is metaBlurb
	Errors   map[string]string //N/A for validator
}

type MetasByTypes struct {
	MBT []MetasByType
}

type MetasByType struct {
	MetaType MetaType
	Metas    []Meta
}
