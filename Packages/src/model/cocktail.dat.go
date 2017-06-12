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
		Rating:          4,
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
			Metadata[22],
		},
		Type: []Meta{
			Metadata[23],
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
			Metadata[24],
			Metadata[25],
		},
		Type: []Meta{
			Metadata[26],
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
			Metadata[22],
		},
		Type: []Meta{
			Metadata[26],
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
			Metadata[25],
			Metadata[27],
		},
		Type: []Meta{
			Metadata[26],
		},
		Served: []Meta{
			Metadata[12],
		},
		Technique: []Meta{
			Metadata[14],
		},
		Ratio: []Meta{
			Metadata[29],
		},
		Family: []Meta{
			Metadata[28],
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
			Metadata[25],
			Metadata[27],
		},
		Type: []Meta{
			Metadata[26],
		},
		Served: []Meta{
			Metadata[12],
		},
		Technique: []Meta{
			Metadata[14],
		},
		Ratio: []Meta{},
		Family: []Meta{
			Metadata[28],
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
	Cocktail{
		ID:              6,
		Title:           "Alexander",
		Name:            "Alexander",
		Description:     "Originally known as <i>Alexander #2</i>, the Alexander is thought to have been created sometime during the 1930s, certainly prior to 1937 when it first appears in print. This classic blend of cognac and chocolate smoothed with cream is based on the original Alexander calling for gin as its base. As to whom substituted cognac in place of gin is lost in time.",
		ImagePath:       "https://s3.ca-central-1.amazonaws.com/commonwealthcocktailsimages",
		Image:           "alexander.jpg",
		ImageSourceName: "Difford's Guide",
		ImageSourceLink: "https://cdn.diffordsguide.com/contrib/stock-images/2015/12/36/201504ca8aabbe67df989ccc2ef013606206.jpg",
		SourceName:      "International Bartenders Association",
		SourceLink:      "http://iba-world.com/iba-official-cocktails/alexander/",
		Rating:          3,
		AKA:             []Name{},
		Tool: []Product{
			Products[14],
			Products[50],
		},
		Strength: []Meta{
			Metadata[4],
		},
		Difficulty: []Meta{
			Metadata[0],
		},
		TOD: []Meta{
			Metadata[9],
			Metadata[10],
		},
		Flavor: []Meta{
			Metadata[22],
		},
		Type: []Meta{
			Metadata[30],
		},
		Served: []Meta{
			Metadata[13],
		},
		Technique: []Meta{
			Metadata[14],
		},
		Ratio:        []Meta{},
		Family:       []Meta{},
		IsFamilyRoot: false,
		BaseSpirit: []Meta{
			Metadata[20],
		},
		Garnish: []Product{
			Products[48],
		},
		Drinkware: []Product{
			Products[49],
		},
		Recipe: Recipes[5],
	},
	Cocktail{
		ID:              7,
		Title:           "Americano",
		Name:            "Americano",
		Description:     "First served in the 1860s in Gaspare Campari's bar in Milan, this was originally known as the 'Milano-Torino' as Campari came from Milano (Milan) and sweet vermouth from Torino (Turin). It was not until Prohibition that the Italians noticed an influx of Americans who enjoyed the drink and so dubbed it Americano.",
		ImagePath:       "https://s3.ca-central-1.amazonaws.com/commonwealthcocktailsimages",
		Image:           "americano.jpg",
		ImageSourceName: "Branca",
		ImageSourceLink: "http://www.branca.it/dati/menu/cocktail_americano_pag_foto.jpg",
		SourceName:      "International Bartenders Association",
		SourceLink:      "http://iba-world.com/iba-official-cocktails/americano/",
		Rating:          3,
		AKA:             []Name{},
		Tool: []Product{
			Products[54],
			Products[55],
		},
		Strength: []Meta{
			Metadata[4],
		},
		Difficulty: []Meta{
			Metadata[0],
		},
		TOD: []Meta{
			Metadata[7],
			Metadata[8],
		},
		Flavor: []Meta{
			Metadata[31],
		},
		Type: []Meta{
			Metadata[30],
		},
		Served: []Meta{
			Metadata[12],
		},
		Technique: []Meta{
			Metadata[21],
		},
		Ratio:        []Meta{},
		Family:       []Meta{},
		IsFamilyRoot: false,
		BaseSpirit:   []Meta{},
		Garnish: []Product{
			Products[56],
		},
		Drinkware: []Product{
			Products[17],
		},
		Recipe: Recipes[6],
	},
	Cocktail{
		ID:              8,
		Title:           "Angel Face",
		Name:            "Angel Face",
		Description:     "Rich apricot and apple with a backbone of botanical gin. Balanced rather than dry or sweet.<br><br>This drink looks better when stirred but the original 1930 recipe is shaken and we think it tastes better for it - that is unless you add some water to increase the dilution of the stirred recipe.",
		ImagePath:       "https://s3.ca-central-1.amazonaws.com/commonwealthcocktailsimages",
		Image:           "angel_face.jpg",
		ImageSourceName: "Difford's Guide",
		ImageSourceLink: "https://cdn.diffordsguide.com/contrib/stock-images/2015/8/02/2015646aef827a2f03d90f7f520ed97dc295.jpg",
		SourceName:      "International Bartenders Association",
		SourceLink:      "http://iba-world.com/iba-official-cocktails/angel-face/",
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
			Metadata[6],
			Metadata[7],
			Metadata[10],
		},
		Flavor: []Meta{
			Metadata[32],
		},
		Type: []Meta{
			Metadata[30],
		},
		Served: []Meta{
			Metadata[13],
		},
		Technique: []Meta{
			Metadata[14],
		},
		Ratio:        []Meta{},
		Family:       []Meta{},
		IsFamilyRoot: false,
		BaseSpirit:   []Meta{},
		Garnish: []Product{
			Products[59],
		},
		Drinkware: []Product{
			Products[49],
		},
		Recipe: Recipes[7],
	},
	Cocktail{
		ID:              9,
		Title:           "Aviation",
		Name:            "Aviation",
		Description:     "This is a fantastic, tangy cocktail and dangerously easy to drink - too many of these and you really will be flying.",
		ImagePath:       "https://s3.ca-central-1.amazonaws.com/commonwealthcocktailsimages",
		Image:           "aviation.jpg",
		ImageSourceName: "Secret Gin Club",
		ImageSourceLink: "https://secretginclub.files.wordpress.com/2012/11/secret-gin-club-no-3-aviation-cocktail.jpg",
		SourceName:      "International Bartenders Association",
		SourceLink:      "http://iba-world.com/iba-official-cocktails/aviation/",
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
			Metadata[6],
			Metadata[7],
			Metadata[10],
		},
		Flavor: []Meta{
			Metadata[32],
		},
		Type: []Meta{
			Metadata[30],
		},
		Served: []Meta{
			Metadata[13],
		},
		Technique: []Meta{
			Metadata[14],
		},
		Ratio:        []Meta{},
		Family:       []Meta{},
		IsFamilyRoot: false,
		BaseSpirit:   []Meta{},
		Garnish: []Product{
			Products[15],
		},
		Drinkware: []Product{
			Products[60],
		},
		Recipe: Recipes[8],
	},
}

var CS = []CocktailSet{
	CocktailSet{
		ChildCocktails: []Cocktail{
			Cocktails[4],
		},
		RootCocktail: Cocktails[3],
	},
}
