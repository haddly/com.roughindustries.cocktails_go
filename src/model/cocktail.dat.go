//model/cocktail.dat.go
package model

var Cocktails = []Cocktail{
	Cocktail{
		ID:              1,
		Title:           "Jamaican Quaalude",
		Name:            "Jamaican Quaalude",
		Description:     "I'll assume that this delicious cocktail's name is derived from its tropical flavors (Jamaican), and its mind numbing effects (Quaalude). With five spirits, and a bit of cream to blend it all together, this rich drink is a great dessert cocktail that will definitely keep the evening going. We hope you'll try our featured cocktail, the Jamaican Quaalude!",
		ImagePath:       "https://s3.ca-central-1.amazonaws.com/commonwealthcocktailsimages",
		Image:           "jamaican_quaalude_750_750.png",
		ImageSourceName: "Unknown",
		ImageSourceLink: "",
		SourceName:      "Hampton Roads Happy Hour",
		SourceLink:      "http://hamptonroadshappyhour.com/jamaican-quaalude",
		Rating:          5,
		IsFamilyRoot:    false,
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
		ImagePath:       "https://s3.ca-central-1.amazonaws.com/commonwealthcocktailsimages",
		Image:           "devils_share.jpg",
		ImageSourceName: "Difford's Guide",
		ImageSourceLink: "https://www.diffordsguide.com/cocktails/recipe/2376/devils-share",
		SourceName:      "Difford's Guide",
		SourceLink:      "https://www.diffordsguide.com/cocktails/recipe/2376/devils-share",
		Comment:         "A little more OJ won't hurt it.",
		Rating:          5,
		IsFamilyRoot:    false,
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
		ImagePath:       "https://s3.ca-central-1.amazonaws.com/commonwealthcocktailsimages",
		Image:           "mad_monk_milkshake.jpg",
		ImageSourceName: "Difford's Guide",
		ImageSourceLink: "https://www.diffordsguide.com/cocktails/recipe/1215/mad-monk-milkshake",
		SourceName:      "Difford's Guide",
		SourceLink:      "https://www.diffordsguide.com/cocktails/recipe/1215/mad-monk-milkshake",
		Rating:          5,
		IsFamilyRoot:    false,
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
		Title:           "Sour",
		Name:            "Sour",
		Description:     "Sours are aptly named drinks. Their flavour comes from either lemon or lime juice, which is balanced with sugar.<br><br>Sours can be based on practically any spirit but the bourbon based Whiskey Sour is by far the most popular. Many (including us) believe this drink is only properly made when smoothed with a little egg white.<br><br>Sours are served either straight-up in a sour glass (rather like a small flute) or on the rocks in an old-fashioned glass. They are traditionally garnished with a cherry and an orange slice, or sometimes a lemon slice.",
		Image:           "24-sour.jpg",
		ImagePath:       "https://s3.ca-central-1.amazonaws.com/commonwealthcocktailsimages",
		ImageSourceName: "Liquor.com",
		ImageSourceLink: "http://liquor.s3.amazonaws.com/wp-content/uploads/2010/02/24-sour.jpeg",
		SourceName:      "Difford's Guide",
		SourceLink:      "https://www.diffordsguide.com/cocktails/recipe/1795/sour-generic-name",
		Comment:         "This 4:2:8 formula is a tad sourer than the classic 3:4:8 which translates as: three quarter part of the sour ingredient (lemon juice), one part of the sweet ingredient (sugar syrup) and two parts of the strong ingredient (spirit). So if you find my formula too sour than best follow the classic proportions in future.",
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
				MetaName: "Sour",
				MetaType: Flavor,
			},
			Meta{
				MetaName: "Sweet",
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
		Ratio: []Meta{
			Meta{
				MetaName: "4:2:8",
				MetaType: Ratio,
			},
		},
		Family: Meta{
			MetaName: "Sour",
			MetaType: Family,
		},
		IsFamilyRoot: true,
		BaseSpirit:   []Meta{},
		Garnish: []Product{
			Products[15],
		},
		Drinkware: []Product{
			Products[17],
		},
		Recipe: Recipes[3],
	},
	Cocktail{
		ID:              5,
		Title:           "Amaretto Sour",
		Name:            "Amaretto Sour",
		Description:     "Sweet 'n' sour - frothy with an almond buzz. Three dashes (12 drops) of Angostura bitters help balance the drink and add an extra burst of flavour",
		ImagePath:       "https://s3.ca-central-1.amazonaws.com/commonwealthcocktailsimages",
		Image:           "amaretto_sour.jpg",
		ImageSourceName: "Difford's Guide",
		ImageSourceLink: "https://d3d9kvjpisqjte.cloudfront.net/stock-images/2015/8/55/2015bed5b504e209d7e292dceaaf470b64b4.jpg",
		SourceName:      "Difford's Guide",
		SourceLink:      "https://www.diffordsguide.com/cocktails/recipe/53/amaretto-sour",
		Rating:          3,
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
				MetaName: "Sour",
				MetaType: Flavor,
			},
			Meta{
				MetaName: "Sweet",
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
		Ratio: []Meta{},
		Family: Meta{
			MetaName: "Sour",
			MetaType: Family,
		},
		IsFamilyRoot: false,
		BaseSpirit:   []Meta{},
		Garnish: []Product{
			Products[15],
		},
		Drinkware: []Product{
			Products[17],
		},
		Recipe: Recipes[4],
	},
}

var FamilyCocktails = []FamilyCocktail{
	FamilyCocktail{
		ChildCocktails: []Cocktail{
			Cocktails[4],
		},
		RootCocktail: Cocktails[3],
	},
}
