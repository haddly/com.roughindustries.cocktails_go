//model/recipe.dat.go
package model

var Recipes = []Recipe{
	Recipe{
		ID:     1,
		Method: "Combine all of the OriginalIngredients in an ice filled cocktail shaker.  Cover, shake well, and pour into a Rocks glass.  Add a couple of sipping straws, garnish accordingly.",
		RecipeSteps: []RecipeStep{
			//1 oz. Kahlua
			RecipeStep{
				OriginalIngredient:   Products[6],
				RecipeCardinalFloat:  1.0,
				RecipeCardinalString: "1",
				RecipeDoze:           Ounce,
				RecipeOrdinal:        0,
			},
			//1 oz. Coconut Rum
			RecipeStep{
				OriginalIngredient:   Products[9],
				RecipeCardinalFloat:  1.0,
				RecipeCardinalString: "1",
				RecipeDoze:           Ounce,
				RecipeOrdinal:        1,
			},
			//1 oz. Baileys Irish Cream
			RecipeStep{
				OriginalIngredient:   Products[3],
				RecipeCardinalFloat:  1.0,
				RecipeCardinalString: "1",
				RecipeDoze:           Ounce,
				RecipeOrdinal:        2,
			},
			//.5 oz Amaretto
			RecipeStep{
				OriginalIngredient:   Products[10],
				RecipeCardinalFloat:  0.5,
				RecipeCardinalString: "½",
				RecipeDoze:           Ounce,
				RecipeOrdinal:        3,
			},
			//.5 oz Frangelico
			RecipeStep{
				OriginalIngredient:   Products[4],
				RecipeCardinalFloat:  0.5,
				RecipeCardinalString: "½",
				RecipeDoze:           Ounce,
				RecipeOrdinal:        4,
			},
			//1 oz Cream
			RecipeStep{
				OriginalIngredient:   Products[8],
				RecipeCardinalFloat:  1.0,
				RecipeCardinalString: "1",
				RecipeDoze:           Ounce,
				RecipeOrdinal:        5,
			},
		},
	},
	Recipe{
		ID:     2,
		Method: "Muddle ginger in base of shaker. Add other OriginalIngredients, shake with ice and fine strain into ice-filled glass.",
		RecipeSteps: []RecipeStep{
			//2 slice Fresh root ginger
			RecipeStep{
				OriginalIngredient:   Products[11],
				RecipeCardinalFloat:  2.0,
				RecipeCardinalString: "2",
				RecipeDoze:           Slice,
				RecipeOrdinal:        0,
			},
			//2 oz. Bourbon Whiskey
			RecipeStep{
				OriginalIngredient: Products[2],
				AltIngredient: []Product{
					Products[22], Products[23],
				},
				RecipeCardinalFloat:  2.0,
				RecipeCardinalString: "2",
				RecipeDoze:           Ounce,
				RecipeOrdinal:        1,
			},
			//.75 oz. Orange Juice
			RecipeStep{
				OriginalIngredient:   Products[0],
				RecipeCardinalFloat:  .75,
				RecipeCardinalString: "¾",
				RecipeDoze:           Ounce,
				RecipeOrdinal:        2,
			},
			//1 oz Lemon Juice
			RecipeStep{
				OriginalIngredient:   Products[1],
				RecipeCardinalFloat:  1.0,
				RecipeCardinalString: "1",
				RecipeDoze:           Ounce,
				RecipeOrdinal:        3,
			},
			//.5 oz Maple Syrup
			RecipeStep{
				OriginalIngredient:   Products[5],
				RecipeCardinalFloat:  0.5,
				RecipeCardinalString: "½",
				RecipeDoze:           Ounce,
				RecipeOrdinal:        4,
			},
		},
	},
	Recipe{
		ID:     3,
		Method: "Shake all OriginalIngredients with ice and strain into ice-filled glass.",
		RecipeSteps: []RecipeStep{
			//1 ounce irish cream
			RecipeStep{
				OriginalIngredient:   Products[3],
				RecipeCardinalFloat:  1.0,
				RecipeCardinalString: "1",
				RecipeDoze:           Liqueur,
				RecipeOrdinal:        0,
			},
			//2 oz. Frangelico
			RecipeStep{
				OriginalIngredient:   Products[4],
				RecipeCardinalFloat:  2.0,
				RecipeCardinalString: "2",
				RecipeDoze:           Ounce,
				RecipeOrdinal:        1,
			},
			//.5 oz. coffee liqueur
			RecipeStep{
				OriginalIngredient:   Products[6],
				RecipeCardinalFloat:  .5,
				RecipeCardinalString: "½",
				RecipeDoze:           Ounce,
				RecipeOrdinal:        2,
			},
			//2 oz Milk
			RecipeStep{
				OriginalIngredient:   Products[7],
				RecipeCardinalFloat:  2.0,
				RecipeCardinalString: "2",
				RecipeDoze:           Ounce,
				RecipeOrdinal:        3,
			},
			//1 oz Cream
			RecipeStep{
				OriginalIngredient:   Products[8],
				RecipeCardinalFloat:  1.0,
				RecipeCardinalString: "1",
				RecipeDoze:           Ounce,
				RecipeOrdinal:        4,
			},
		},
	},
	Recipe{
		ID:     4,
		Method: "Shake all ingredients with ice and strain into ice-filled glass.",
		RecipeSteps: []RecipeStep{
			//1 oz generic
			RecipeStep{
				OriginalIngredient:   Products[38],
				RecipeCardinalFloat:  2.0,
				RecipeCardinalString: "2",
				RecipeDoze:           Ounce,
				RecipeOrdinal:        0,
			},
			//1 oz lemon juice
			RecipeStep{
				OriginalIngredient:   Products[1],
				RecipeCardinalFloat:  1.0,
				RecipeCardinalString: "1",
				RecipeDoze:           Ounce,
				RecipeOrdinal:        1,
			},
			//1/2 oz Simple Syrup
			RecipeStep{
				OriginalIngredient:   Products[39],
				RecipeCardinalFloat:  0.5,
				RecipeCardinalString: "½",
				RecipeDoze:           Ounce,
				RecipeOrdinal:        2,
			},
			//3 dashes bittes
			RecipeStep{
				OriginalIngredient:   Products[40],
				RecipeCardinalFloat:  3.0,
				RecipeCardinalString: "3",
				RecipeDoze:           Dash,
				RecipeOrdinal:        3,
			},
			//.5 fresh egg white
			RecipeStep{
				OriginalIngredient:   Products[41],
				RecipeCardinalFloat:  0.5,
				RecipeCardinalString: "½",
				RecipeDoze:           Fresh,
				RecipeOrdinal:        4,
			},
		},
	},
	Recipe{
		ID:     5,
		Method: "Dry Shake (without ice) all ingredients to emulsify. Add ice, shake again and strain into ice-filled glass.",
		RecipeSteps: []RecipeStep{
			//2 oz Amaretto
			RecipeStep{
				OriginalIngredient:   Products[10],
				RecipeCardinalFloat:  2.0,
				RecipeCardinalString: "2",
				RecipeDoze:           Ounce,
				RecipeOrdinal:        0,
			},
			//1 oz lemon juice
			RecipeStep{
				OriginalIngredient:   Products[1],
				RecipeCardinalFloat:  1.0,
				RecipeCardinalString: "1",
				RecipeDoze:           Ounce,
				RecipeOrdinal:        1,
			},
			//1 dashes bittes
			RecipeStep{
				OriginalIngredient:   Products[40],
				RecipeCardinalFloat:  1.0,
				RecipeCardinalString: "1",
				RecipeDoze:           Dash,
				RecipeOrdinal:        2,
			},
			//.5 fresh egg white
			RecipeStep{
				OriginalIngredient:   Products[41],
				RecipeCardinalFloat:  0.5,
				RecipeCardinalString: "½",
				RecipeDoze:           Fresh,
				RecipeOrdinal:        3,
			},
		},
	},
	Recipe{
		ID:     6,
		Method: "Shake and strain into a chilled cocktail glass.",
		RecipeSteps: []RecipeStep{
			//1 oz cognac
			RecipeStep{
				OriginalIngredient:   Products[46],
				RecipeCardinalFloat:  1.0,
				RecipeCardinalString: "1",
				RecipeDoze:           Ounce,
				RecipeOrdinal:        0,
			},
			//1 oz Créme de Cacao
			RecipeStep{
				OriginalIngredient:   Products[47],
				RecipeCardinalFloat:  1.0,
				RecipeCardinalString: "1",
				RecipeDoze:           Ounce,
				RecipeOrdinal:        1,
			},
			//1 oz cream
			RecipeStep{
				OriginalIngredient:   Products[8],
				RecipeCardinalFloat:  1.0,
				RecipeCardinalString: "1",
				RecipeDoze:           Ounce,
				RecipeOrdinal:        2,
			},
		},
	},
	Recipe{
		ID:     7,
		Method: "Mix the first two ingredients directly in an old-fashioned glass filled with ice-cubes, then add a splash of soda water.",
		RecipeSteps: []RecipeStep{
			//1 oz campari bitter
			RecipeStep{
				OriginalIngredient:   Products[51],
				RecipeCardinalFloat:  1.0,
				RecipeCardinalString: "1",
				RecipeDoze:           Ounce,
				RecipeOrdinal:        0,
			},
			//1 oz sweet vermouth
			RecipeStep{
				OriginalIngredient:   Products[52],
				RecipeCardinalFloat:  1.0,
				RecipeCardinalString: "1",
				RecipeDoze:           Ounce,
				RecipeOrdinal:        1,
			},
			//1 oz cream
			RecipeStep{
				OriginalIngredient: Products[53],
				RecipeDoze:         TopOffWith,
				RecipeOrdinal:      2,
			},
		},
	},
}
