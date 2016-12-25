//model/recipe.dat.go
package model

var Recipes = []Recipe{
	Recipe{
		ID:     0,
		Method: "Combine all of the ingredients in an ice filled cocktail shaker.  Cover, shake well, and pour into a Rocks glass.  Add a couple of sipping straws, garnish accordingly.",
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
	Recipe{
		Method: "Muddle ginger in base of shaker. Add other ingredients, shake with ice and fine strain into ice-filled glass.",
		RecipeSteps: []RecipeStep{
			//2 slice Fresh root ginger
			RecipeStep{
				Ingredient: Product{
					ProductName: "Ginger Root",
					ProductType: Mixer,
				},
				RecipeCardinal: 2.0,
				RecipeDoze:     Slice,
				RecipeOrdinal:  0,
			},
			//2 oz. Bourbon Whiskey
			RecipeStep{
				Ingredient: Product{
					ProductName: "Bourbon Whiskey",
					ProductType: Spirit,
				},
				RecipeCardinal: 2.0,
				RecipeDoze:     Ounce,
				RecipeOrdinal:  1,
			},
			//.75 oz. Orange Juice
			RecipeStep{
				Ingredient:     Products[0],
				RecipeCardinal: .75,
				RecipeDoze:     Ounce,
				RecipeOrdinal:  2,
			},
			//1 oz Lemon Juice
			RecipeStep{
				Ingredient:     Products[2],
				RecipeCardinal: 1.0,
				RecipeDoze:     Ounce,
				RecipeOrdinal:  3,
			},
			//.5 oz Maple Syrup
			RecipeStep{
				Ingredient: Product{
					ProductName: "Maple Syrup",
					ProductType: Mixer,
				},
				RecipeCardinal: 0.5,
				RecipeDoze:     Ounce,
				RecipeOrdinal:  4,
			},
		},
	},
}
