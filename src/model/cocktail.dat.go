//model/cocktail.dat.go
package model

var Cocktails = []Cocktail{
	Cocktail{
		Title:           "Jamaican Quaalude",
		Name:            "Jamaican Quaalude",
		Description:     "I'll assume that this delicious cocktail's name is derived from its tropical flavors (Jamaican), and its mind numbing effects (Quaalude). With five spirits, and a bit of cream to blend it all together, this rich drink is a great dessert cocktail that will definitely keep the evening going. We hope you'll try our featured cocktail, the Jamaican Quaalude!",
		Method:          "Combine all of the ingredients in an ice filled cocktail shaker.  Cover, shake well, and pour into a Rocks glass.  Add a couple of sipping straws, garnish accordingly.",
		Image:           "jamaican_quaalude_750_750.png",
		ImageSourceName: "Unknown",
		ImageSourceLink: "",
		SourceName:      "Hampton Roads Happy Hour",
		SourceLink:      "http://hamptonroadshappyhour.com/jamaican-quaalude",
		Rating:          2.5 / 5.0 * 100,
		AKA: []AKA{
			AKA{
				Name: "Jamaican Quaalude",
			},
		},
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
		Recipe: Recipe{
			RecipeSteps: []RecipeStep{
				//1 oz. Kahlua
				RecipeStep{
					Ingredient: Product{
						ProductName: "Kahlua",
						ProductType: Liqueur,
					},
					RecipeCardinal: 1.0,
					RecipeDoze:     Ounce,
					RecipeOrdinal:  0,
				},
				//1 oz. Coconut Rum
				RecipeStep{
					Ingredient: Product{
						ProductName: "Coconut Rum",
						ProductType: Spirit,
					},
					RecipeCardinal: 1.0,
					RecipeDoze:     Ounce,
					RecipeOrdinal:  1,
				},
				//1 oz. Baileys Irish Cream
				RecipeStep{
					Ingredient: Product{
						ProductName: "Irish Cream",
						ProductType: Liqueur,
					},
					RecipeCardinal: 1.0,
					RecipeDoze:     Ounce,
					RecipeOrdinal:  2,
				},
				//.5 oz Amaretto
				RecipeStep{
					Ingredient: Product{
						ProductName: "Amaretto",
						ProductType: Liqueur,
					},
					RecipeCardinal: 0.5,
					RecipeDoze:     Ounce,
					RecipeOrdinal:  3,
				},
				//.5 oz Frangelico
				RecipeStep{
					Ingredient: Product{
						ProductName: "Frangelico",
						ProductType: Liqueur,
					},
					RecipeCardinal: 0.5,
					RecipeDoze:     Ounce,
					RecipeOrdinal:  4,
				},
				//1 oz Cream
				RecipeStep{
					Ingredient: Product{
						ProductName: "Cream",
						ProductType: Mixer,
					},
					RecipeCardinal: 1.0,
					RecipeDoze:     Ounce,
					RecipeOrdinal:  5,
				},
			},
		},
	},
}
