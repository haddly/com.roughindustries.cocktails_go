//model/meta.go
package model

type MetaType int

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

var MetaTypeStrings = [...]string{
	"Flavor",      //Fruity, Bitter, Creamy, ...
	"Base Spirit", //Vodka, Gin, Bourbon, ...
	"Type",        //Tiki,
	"Occasion",    //Christmas, 4th of July, Halloween, ...
	"Family",      //Margarita, Martini, ...
	"Formula",
	"Served",      //Highball, Martini, Old Fashioned, ...
	"Technique",   //Shaking, Stirring, Straining
	"Strength",    //Weak, Medium, Strong
	"Difficulty",  //Easy, Medium, Hard
	"Time of Day", //Evening, Dessert, Brunch, ...
	"Ratio",       //3:4:8, ...
	"Drink",
}

// String returns the English name of the metatype ("Flavor", "Base Spirit", ...).
func (mt MetaType) String() string { return MetaTypeStrings[mt-1] }

type Meta struct {
	ID       int
	MetaName string
	MetaType MetaType
	Article  Post
	Blurb    Post
}
