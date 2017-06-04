//model/meta.dat.go
package model

var MetaTypes = []MetaType{
	MetaType{
		ID:                   Flavor,
		MetaTypeName:         "Flavor",
		ShowInCocktailsIndex: true,
		Ordinal:              2,
	},
	MetaType{
		ID:                   BaseSpirit,
		MetaTypeName:         "Base Spirit",
		ShowInCocktailsIndex: true,
		Ordinal:              1,
	},
	MetaType{
		ID:                   Type,
		MetaTypeName:         "Type",
		ShowInCocktailsIndex: false,
		Ordinal:              3,
	},
	MetaType{
		ID:                   Occasion,
		MetaTypeName:         "Occasion",
		ShowInCocktailsIndex: false,
		Ordinal:              4,
	},
	MetaType{
		ID:                   Family,
		MetaTypeName:         "Family",
		ShowInCocktailsIndex: false,
		Ordinal:              5,
	},
	MetaType{
		ID:                   Formula,
		MetaTypeName:         "Formula",
		ShowInCocktailsIndex: false,
		Ordinal:              6,
	},
	MetaType{
		ID:                   Served,
		MetaTypeName:         "Served",
		ShowInCocktailsIndex: false,
		Ordinal:              7,
	},
	MetaType{
		ID:                   Technique,
		MetaTypeName:         "Technique",
		ShowInCocktailsIndex: false,
		Ordinal:              8,
	},
	MetaType{
		ID:                   Strength,
		MetaTypeName:         "Strength",
		ShowInCocktailsIndex: false,
		Ordinal:              9,
	},
	MetaType{
		ID:                   Difficulty,
		MetaTypeName:         "Difficulty",
		ShowInCocktailsIndex: false,
		Ordinal:              10,
	},
	MetaType{
		ID:                   TOD,
		MetaTypeName:         "Time of Day",
		ShowInCocktailsIndex: false,
		Ordinal:              11,
	},
	MetaType{
		ID:                   Ratio,
		MetaTypeName:         "Ratio",
		ShowInCocktailsIndex: false,
		Ordinal:              12,
	},
	MetaType{
		ID:                   Drink,
		MetaTypeName:         "Drink",
		ShowInCocktailsIndex: false,
		Ordinal:              13,
	},
}

var Metadata = []Meta{
	Meta{
		ID:       1,
		MetaName: "Easy",
		MetaType: MetaTypes[Difficulty-1],
	},
	Meta{
		ID:       2,
		MetaName: "Medium",
		MetaType: MetaTypes[Difficulty-1],
	},
	Meta{
		ID:       3,
		MetaName: "Hard",
		MetaType: MetaTypes[Difficulty-1],
	},
	Meta{
		ID:       4,
		MetaName: "Weak",
		MetaType: MetaTypes[Strength-1],
	},
	Meta{
		ID:       5,
		MetaName: "Medium",
		MetaType: MetaTypes[Strength-1],
	},
	Meta{
		ID:       6,
		MetaName: "Strong",
		MetaType: MetaTypes[Strength-1],
	},
	Meta{
		ID:       7,
		MetaName: "Morning/Brunch",
		MetaType: MetaTypes[TOD-1],
	},
	Meta{
		ID:       8,
		MetaName: "Afternoon",
		MetaType: MetaTypes[TOD-1],
	},
	Meta{
		ID:       9,
		MetaName: "Dinner",
		MetaType: MetaTypes[TOD-1],
	},
	Meta{
		ID:       10,
		MetaName: "Dessert",
		MetaType: MetaTypes[TOD-1],
	},
	Meta{
		ID:       11,
		MetaName: "Evening",
		MetaType: MetaTypes[TOD-1],
	},
	Meta{
		ID:       12,
		MetaName: "Nightcap",
		MetaType: MetaTypes[TOD-1],
	},
	Meta{
		ID:       13,
		MetaName: "On the Rocks",
		MetaType: MetaTypes[Served-1],
	},
	Meta{
		ID:       14,
		MetaName: "Neat/Straight Up",
		MetaType: MetaTypes[Served-1],
	},
	Meta{
		ID:       15,
		MetaName: "Shaken",
		MetaType: MetaTypes[Technique-1],
	},
	Meta{
		ID:       16,
		MetaName: "Muddled",
		MetaType: MetaTypes[Technique-1],
	},
	Meta{
		ID:       17,
		MetaName: "Bourbon Whiskey",
		MetaType: MetaTypes[BaseSpirit-1],
	},
	Meta{
		ID:       18,
		MetaName: "Neat/Straight Up",
		MetaType: MetaTypes[Drink-1],
	},
	Meta{
		ID:       19,
		MetaName: "With Cola",
		MetaType: MetaTypes[Drink-1],
	},
	Meta{
		ID:       20,
		MetaName: "In Cocktails",
		MetaType: MetaTypes[Drink-1],
	},
	Meta{
		ID:       21,
		MetaName: "Cognac",
		MetaType: MetaTypes[BaseSpirit-1],
	},
	Meta{
		ID:       22,
		MetaName: "Stirred",
		MetaType: MetaTypes[Technique-1],
	},
	Meta{
		ID:       23,
		MetaName: "Creamy",
		MetaType: MetaTypes[Flavor-1],
	},
	Meta{
		ID:       24,
		MetaName: "Creamy",
		MetaType: MetaTypes[Type-1],
	},
	Meta{
		ID:       25,
		MetaName: "Fruity/Citrus",
		MetaType: MetaTypes[Flavor-1],
	},
	Meta{
		ID:       26,
		MetaName: "Sour",
		MetaType: MetaTypes[Flavor-1],
	},
	Meta{
		ID:       27,
		MetaName: "Recommended",
		MetaType: MetaTypes[Type-1],
	},
	Meta{
		ID:       28,
		MetaName: "Sweet",
		MetaType: MetaTypes[Flavor-1],
	},
	Meta{
		ID:       29,
		MetaName: "Sour",
		MetaType: MetaTypes[Family-1],
	},
	Meta{
		ID:       30,
		MetaName: "4:2:8",
		MetaType: MetaTypes[Ratio-1],
	},
	Meta{
		ID:       31,
		MetaName: "Top 100",
		MetaType: MetaTypes[Type-1],
	},
	Meta{
		ID:       32,
		MetaName: "Bitter",
		MetaType: MetaTypes[Flavor-1],
	},
	Meta{
		ID:       33,
		MetaName: "Fruity",
		MetaType: MetaTypes[Flavor-1],
	},
}
