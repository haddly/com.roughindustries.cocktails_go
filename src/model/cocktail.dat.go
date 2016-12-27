//model/cocktail.dat.go
package model

var Cocktails = []Cocktail{
	Cocktail{
		ID:              1,
		Title:           "Jamaican Quaalude",
		Name:            "Jamaican Quaalude",
		Description:     "I'll assume that this delicious cocktail's name is derived from its tropical flavors (Jamaican), and its mind numbing effects (Quaalude). With five spirits, and a bit of cream to blend it all together, this rich drink is a great dessert cocktail that will definitely keep the evening going. We hope you'll try our featured cocktail, the Jamaican Quaalude!",
		Image:           "jamaican_quaalude_750_750.png",
		ImageSourceName: "Unknown",
		ImageSourceLink: "",
		SourceName:      "Hampton Roads Happy Hour",
		SourceLink:      "http://hamptonroadshappyhour.com/jamaican-quaalude",
		// Rating:          2.5 / 5.0 * 100,
		Rating: 5,
		AKA:    []Name{},
		Tool: []Product{
			Product{
				ProductName: "Shaker",
				ProductType: Tool,
			},
		},
		Strength: []Meta{
			Meta{
				MetaName: "Medium",
				MetaType: Flavor,
			},
		},
		Difficulty: []Meta{
			Metadata[0],
		},
		TOD: []Meta{
			Meta{
				MetaName: "Evening",
				MetaType: TOD,
			},
		},
		Flavor: []Meta{
			Meta{
				MetaName: "Creamy",
				MetaType: Flavor,
			},
		},
		Type: []Meta{
			Meta{
				MetaName: "Creamy",
				MetaType: Type,
			},
		},
		Served: []Meta{
			Meta{
				MetaName: "On the Rocks",
				MetaType: Served,
			},
		},
		Technique: []Meta{
			Meta{
				MetaName: "Shaken",
				MetaType: Technique,
			},
		},
		BaseSpirit: []Meta{},
		Garnish: []Product{
			Product{
				ProductName: "Cherry",
				ProductType: Garnish,
			},
			Product{
				ProductName: "Slice of Starfruit",
				ProductType: Garnish,
			},
		},
		Drinkware: []Product{
			Product{
				ProductName: "Old Fashioned",
				ProductType: Drinkware,
			},
		},
		Recipe: Recipes[0],
	},
	Cocktail{
		ID:              2,
		Title:           "Devil's Share",
		Name:            "Devil's Share",
		Description:     "This cocktail is really smooth on the toungue because of the maple syrup.  The citrus from the lemon and sugar from the maple syrup give a great sour flavor profile.  One of my favorites, mostly because of the maple syrup.",
		Image:           "devils_share.jpg",
		ImageSourceName: "Difford's Guide",
		ImageSourceLink: "https://www.diffordsguide.com/cocktails/recipe/2376/devils-share",
		SourceName:      "Difford's Guide",
		SourceLink:      "https://www.diffordsguide.com/cocktails/recipe/2376/devils-share",
		Rating:          5,
		AKA:             []Name{},
		Tool: []Product{
			Product{
				ProductName: "Shaker",
				ProductType: Tool,
			},
			Product{
				ProductName: "Muddler",
				ProductType: Tool,
			},
		},
		Strength: []Meta{
			Meta{
				MetaName: "Medium",
				MetaType: Flavor,
			},
		},
		Difficulty: []Meta{
			Meta{
				MetaName: "Easy",
				MetaType: Difficulty,
			},
		},
		TOD: []Meta{
			Meta{
				MetaName: "Afternoon",
				MetaType: TOD,
			},
		},
		Flavor: []Meta{
			Meta{
				MetaName: "Fruity/Citrus",
				MetaType: Flavor,
			},
			Meta{
				MetaName: "Sour",
				MetaType: Flavor,
			},
		},
		Type: []Meta{
			Meta{
				MetaName: "Recommended",
				MetaType: Type,
			},
		},
		Served: []Meta{
			Meta{
				MetaName: "On the Rocks",
				MetaType: Served,
			},
		},
		Technique: []Meta{
			Meta{
				MetaName: "Shaken",
				MetaType: Technique,
			},
			Meta{
				MetaName: "Muddled",
				MetaType: Technique,
			},
		},
		BaseSpirit: []Meta{
			Meta{
				MetaName: "Bourbon Whiskey",
				MetaType: BaseSpirit,
			},
		},
		Garnish: []Product{
			Product{
				ProductName: "Orange Zest Twist",
				ProductType: Garnish,
			},
		},
		Drinkware: []Product{
			Product{
				ProductName: "Old Fashioned",
				ProductType: Drinkware,
			},
		},
		Recipe: Recipes[1],
	},
	Cocktail{
		ID:              3,
		Title:           "Mad Monk Milkshake",
		Name:            "Mad Monk Milkshake",
		Description:     "This was one of my really first adventures into creamy cocktails.  It was and still is one of my favorite cocktails ever.  It is smooth, really easy ot drink, and frankly just fun.",
		Image:           "mad_monk_milkshake.jpg",
		ImageSourceName: "Difford's Guide",
		ImageSourceLink: "https://www.diffordsguide.com/cocktails/recipe/1215/mad-monk-milkshake",
		SourceName:      "Difford's Guide",
		SourceLink:      "https://www.diffordsguide.com/cocktails/recipe/1215/mad-monk-milkshake",
		Rating:          5,
		AKA:             []Name{},
		Tool: []Product{
			Product{
				ProductName: "Shaker",
				ProductType: Tool,
			},
		},
		Strength: []Meta{
			Meta{
				MetaName: "Medium",
				MetaType: Flavor,
			},
		},
		Difficulty: []Meta{
			Meta{
				MetaName: "Easy",
				MetaType: Difficulty,
			},
		},
		TOD: []Meta{
			Meta{
				MetaName: "Dessert",
				MetaType: TOD,
			},
		},
		Flavor: []Meta{
			Meta{
				MetaName: "Creamy",
				MetaType: Flavor,
			},
		},
		Type: []Meta{
			Meta{
				MetaName: "Recommended",
				MetaType: Type,
			},
		},
		Served: []Meta{
			Meta{
				MetaName: "On the Rocks",
				MetaType: Served,
			},
		},
		Technique: []Meta{
			Meta{
				MetaName: "Shaken",
				MetaType: Technique,
			},
		},
		BaseSpirit: []Meta{},
		Garnish: []Product{
			Product{
				ProductName: "Cherry",
				ProductType: Garnish,
			},
		},
		Drinkware: []Product{
			Product{
				ProductName: "Collins",
				ProductType: Drinkware,
			},
			Product{
				ProductName: "Old Fashioned",
				ProductType: Drinkware,
			},
			Product{
				ProductName: "Highball",
				ProductType: Drinkware,
			},
		},
		Recipe: Recipes[2],
	},
}
