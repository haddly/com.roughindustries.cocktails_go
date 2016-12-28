//model/recipe.dat.go
package model

var Recipes = []Recipe{
	Recipe{
		ID:     0,
		Method: "Combine all of the ingredients in an ice filled cocktail shaker.  Cover, shake well, and pour into a Rocks glass.  Add a couple of sipping straws, garnish accordingly.",
		RecipeSteps: []RecipeStep{
			//1 oz. Kahlua
			RecipeStep{
				Ingredient:           Products[6],
				RecipeCardinalFloat:  1.0,
				RecipeCardinalString: "1",
				RecipeDoze:           Ounce,
				RecipeOrdinal:        0,
			},
			//1 oz. Coconut Rum
			RecipeStep{
				Ingredient:           Products[9],
				RecipeCardinalFloat:  1.0,
				RecipeCardinalString: "1",
				RecipeDoze:           Ounce,
				RecipeOrdinal:        1,
			},
			//1 oz. Baileys Irish Cream
			RecipeStep{
				Ingredient:           Products[3],
				RecipeCardinalFloat:  1.0,
				RecipeCardinalString: "1",
				RecipeDoze:           Ounce,
				RecipeOrdinal:        2,
			},
			//.5 oz Amaretto
			RecipeStep{
				Ingredient:           Products[10],
				RecipeCardinalFloat:  0.5,
				RecipeCardinalString: "½",
				RecipeDoze:           Ounce,
				RecipeOrdinal:        3,
			},
			//.5 oz Frangelico
			RecipeStep{
				Ingredient:           Products[4],
				RecipeCardinalFloat:  0.5,
				RecipeCardinalString: "½",
				RecipeDoze:           Ounce,
				RecipeOrdinal:        4,
			},
			//1 oz Cream
			RecipeStep{
				Ingredient:           Products[8],
				RecipeCardinalFloat:  1.0,
				RecipeCardinalString: "1",
				RecipeDoze:           Ounce,
				RecipeOrdinal:        5,
			},
		},
	},
	Recipe{
		Method: "Muddle ginger in base of shaker. Add other ingredients, shake with ice and fine strain into ice-filled glass.",
		RecipeSteps: []RecipeStep{
			//2 slice Fresh root ginger
			RecipeStep{
				Ingredient:           Products[11],
				RecipeCardinalFloat:  2.0,
				RecipeCardinalString: "2",
				RecipeDoze:           Slice,
				RecipeOrdinal:        0,
			},
			//2 oz. Bourbon Whiskey
			RecipeStep{
				Ingredient:           Products[2],
				RecipeCardinalFloat:  2.0,
				RecipeCardinalString: "2",
				RecipeDoze:           Ounce,
				RecipeOrdinal:        1,
			},
			//.75 oz. Orange Juice
			RecipeStep{
				Ingredient:           Products[0],
				RecipeCardinalFloat:  .75,
				RecipeCardinalString: "¾",
				RecipeDoze:           Ounce,
				RecipeOrdinal:        2,
			},
			//1 oz Lemon Juice
			RecipeStep{
				Ingredient:           Products[1],
				RecipeCardinalFloat:  1.0,
				RecipeCardinalString: "1",
				RecipeDoze:           Ounce,
				RecipeOrdinal:        3,
			},
			//.5 oz Maple Syrup
			RecipeStep{
				Ingredient:           Products[5],
				RecipeCardinalFloat:  0.5,
				RecipeCardinalString: "½",
				RecipeDoze:           Ounce,
				RecipeOrdinal:        4,
			},
		},
	},
	Recipe{
		Method: "Shake all ingredients with ice and strain into ice-filled glass.",
		RecipeSteps: []RecipeStep{
			//1 ounce irish cream
			RecipeStep{
				Ingredient:           Products[3],
				RecipeCardinalFloat:  1.0,
				RecipeCardinalString: "1",
				RecipeDoze:           Liqueur,
				RecipeOrdinal:        0,
			},
			//2 oz. Frangelico
			RecipeStep{
				Ingredient:           Products[4],
				RecipeCardinalFloat:  2.0,
				RecipeCardinalString: "2",
				RecipeDoze:           Ounce,
				RecipeOrdinal:        1,
			},
			//.5 oz. coffee liqueur
			RecipeStep{
				Ingredient:           Products[6],
				RecipeCardinalFloat:  .5,
				RecipeCardinalString: "½",
				RecipeDoze:           Ounce,
				RecipeOrdinal:        2,
			},
			//2 oz Milk
			RecipeStep{
				Ingredient:           Products[7],
				RecipeCardinalFloat:  2.0,
				RecipeCardinalString: "2",
				RecipeDoze:           Ounce,
				RecipeOrdinal:        3,
			},
			//1 oz Cream
			RecipeStep{
				Ingredient:           Products[8],
				RecipeCardinalFloat:  1.0,
				RecipeCardinalString: "1",
				RecipeDoze:           Ounce,
				RecipeOrdinal:        4,
			},
		},
	},
	Recipe{
		Method: "Shake the booze and cream with ice. Strain into a short glass with lots of ice, but make sure not to fill it to the brim (we know you want too). Leave some space to top it off with cola.",
		RecipeSteps: []RecipeStep{
			//1 oz Kahlua
			RecipeStep{
				Ingredient:           Products[6],
				RecipeCardinalFloat:  1.0,
				RecipeCardinalString: "1",
				RecipeDoze:           Ounce,
				RecipeOrdinal:        0,
			},
			//1 oz Vodka
			RecipeStep{
				Ingredient:           Products[12],
				RecipeCardinalFloat:  1.0,
				RecipeCardinalString: "1",
				RecipeDoze:           Ounce,
				RecipeOrdinal:        0,
			},
			//1 oz Milk
			RecipeStep{
				Ingredient:           Products[7],
				RecipeCardinalFloat:  1.0,
				RecipeCardinalString: "1",
				RecipeDoze:           Ounce,
				RecipeOrdinal:        2,
			},
			//1 oz Cream
			RecipeStep{
				Ingredient:           Products[8],
				RecipeCardinalFloat:  1.0,
				RecipeCardinalString: "1",
				RecipeDoze:           Ounce,
				RecipeOrdinal:        3,
			},
			//Top off with Cola
			RecipeStep{
				Ingredient:    Products[13],
				RecipeDoze:    TopOffWith,
				RecipeOrdinal: 4,
			},
		},
	},
}
