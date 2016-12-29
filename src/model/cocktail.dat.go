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
		Rating:          5,
		AKA: []Name{
			Name{Name: "Jamaican Milkshake"},
		},
		Tool: []Product{
			Products[14],
		},
		Strength: []Meta{
			Metadata[4],
		},
		Difficulty: []Meta{
			Metadata[0],
		},
		TOD: []Meta{
			Metadata[10],
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
			Metadata[14],
		},
		BaseSpirit: []Meta{},
		Garnish: []Product{
			Products[15],
			Products[16],
		},
		Drinkware: []Product{
			Products[17],
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
		Comment:         "A little more OJ won't hurt it.",
		Rating:          5,
		AKA:             []Name{},
		Tool: []Product{
			Products[14],
			Products[18],
		},
		Strength: []Meta{
			Metadata[4],
		},
		Difficulty: []Meta{
			Metadata[0],
		},
		TOD: []Meta{
			Metadata[7],
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
			Metadata[12],
		},
		Technique: []Meta{
			Metadata[14],
			Metadata[15],
		},
		BaseSpirit: []Meta{
			Metadata[16],
		},
		Garnish: []Product{
			Products[21],
		},
		Drinkware: []Product{
			Products[17],
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
			Products[14],
		},
		Strength: []Meta{
			Metadata[4],
		},
		Difficulty: []Meta{
			Metadata[0],
		},
		TOD: []Meta{
			Metadata[9],
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
			Metadata[12],
		},
		Technique: []Meta{
			Metadata[14],
		},
		BaseSpirit: []Meta{},
		Garnish: []Product{
			Products[15],
		},
		Drinkware: []Product{
			Products[20],
			Products[17],
			Products[19],
		},
		Recipe: Recipes[2],
	},
	Cocktail{
		ID:              4,
		Title:           "Colorado Bulldog",
		Name:            "Colorado Bulldog",
		Description:     "The combination of coffee liqueur, vodka, cream and the cola creates a sweet and refreshing taste. We know it looks like a White Russian but don’t be misled; the Colorado Bulldog is bubbly.",
		Image:           "mad_monk_milkshake.jpg",
		ImageSourceName: "Difford's Guide",
		ImageSourceLink: "https://www.diffordsguide.com/cocktails/recipe/1215/mad-monk-milkshake",
		SourceName:      "Kahlúa Website",
		SourceLink:      "http://www.kahlua.com/en/drinks/classic/colorado-bulldog/",
		Rating:          4,
		AKA:             []Name{},
		Tool: []Product{
			Products[14],
		},
		Strength: []Meta{
			Metadata[4],
		},
		Difficulty: []Meta{
			Metadata[0],
		},
		TOD: []Meta{
			Metadata[8],
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
			Metadata[12],
		},
		Technique: []Meta{
			Metadata[14],
		},
		BaseSpirit: []Meta{},
		Garnish: []Product{
			Products[15],
		},
		Drinkware: []Product{
			Products[17],
		},
		Recipe: Recipes[3],
	},
}
