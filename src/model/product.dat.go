package model

var Products = []Product{
	Product{
		ID:          1,
		ProductName: "Orange Juice",
		ProductType: Mixer,
		IsBase:      true,
		PreText:     "Fresh Squeezed",
	},
	Product{
		ID:          2,
		ProductName: "Lemon Juice",
		ProductType: Mixer,
		IsBase:      true,
	},
	Product{
		ID:          3,
		ProductName: "Bourbon Whiskey",
		ProductType: Spirit,
		IsBase:      true,
	},
	Product{
		ID:          4,
		ProductName: "Irish Cream Liqueur",
		ProductType: Liqueur,
		IsBase:      true,
	},
	Product{
		ID:          5,
		ProductName: "Hazelnut Liqueur",
		ProductType: Liqueur,
		IsBase:      true,
	},
	Product{
		ID:          6,
		ProductName: "Maple Syrup",
		ProductType: Liqueur,
		IsBase:      true,
	},
	Product{
		ID:          7,
		ProductName: "Coffee Liqueur",
		ProductType: Liqueur,
		IsBase:      true,
	},
	Product{
		ID:          8,
		ProductName: "Milk",
		ProductType: Mixer,
		IsBase:      true,
	},
	Product{
		ID:          9,
		ProductName: "Cream",
		ProductType: Mixer,
		IsBase:      true,
	},
	Product{
		ID:          10,
		ProductName: "Coconut Rum",
		ProductType: Spirit,
		IsBase:      true,
	},
	Product{
		ID:          11,
		ProductName: "Amaretto",
		ProductType: Liqueur,
		IsBase:      true,
	},
	Product{
		ID:          12,
		ProductName: "Ginger Root",
		ProductType: Mixer,
		IsBase:      true,
		PostText:    "(thumbnail size)",
	},
	Product{
		ID:          13,
		ProductName: "Vodka",
		ProductType: Spirit,
		IsBase:      true,
	},
	Product{
		ID:          14,
		ProductName: "Cola",
		ProductType: Mixer,
		IsBase:      true,
	},
	Product{
		ID:          15,
		ProductName: "Shaker",
		ProductType: Tool,
		IsBase:      true,
	},
	Product{
		ID:          16,
		ProductName: "Cherry",
		ProductType: Garnish,
		IsBase:      true,
	},
	Product{
		ID:          17,
		ProductName: "Starfruit",
		ProductType: Garnish,
		IsBase:      true,
		PreText:     "Slice of",
	},
	Product{
		ID:          18,
		ProductName: "Old Fashioned",
		ProductType: Drinkware,
		IsBase:      true,
	},
	Product{
		ID:          19,
		ProductName: "Muddler",
		ProductType: Tool,
		IsBase:      true,
	},
	Product{
		ID:          20,
		ProductName: "Collins",
		ProductType: Drinkware,
		IsBase:      true,
	},
	Product{
		ID:          21,
		ProductName: "Highball",
		ProductType: Drinkware,
		IsBase:      true,
	},
	Product{
		ID:          22,
		ProductName: "Orange",
		ProductType: Garnish,
		IsBase:      true,
		PostText:    "Zest Twist",
	},
	Product{
		ID:          23,
		ProductName: "Rye Whiskey",
		ProductType: Spirit,
		IsBase:      true,
	},
	Product{
		ID:          24,
		ProductName: "Tennessee Whiskey",
		ProductType: Spirit,
		IsBase:      true,
	},
}

var DerivedProducts = []DerivedProduct{
	DerivedProduct{
		Product: Product{
			ID:          1,
			ProductName: "Tropicana™ Orange Juice",
			ProductType: Mixer,
		},
		BaseProduct: Products[0],
	},
	DerivedProduct{
		Product: Product{
			ID:          2,
			ProductName: "ReaLemon™ Lemon Juice",
			ProductType: Mixer,
		},
		BaseProduct: Products[1],
	},
	DerivedProduct{
		Product: Product{
			ID:          3,
			ProductName: "Pappy Van Winkle's™ Bourbon Whiskey",
			ProductType: Spirit,
		},
		BaseProduct: Products[2],
	},
	DerivedProduct{
		Product: Product{
			ID:          4,
			ProductName: "Breckenridge™ Bourbon Whiskey",
			ProductType: Spirit,
		},
		BaseProduct: Products[2],
	},
	DerivedProduct{
		Product: Product{
			ID:          5,
			ProductName: "Bailey's Irish Cream™",
			ProductType: Liqueur,
		},
		BaseProduct: Products[3],
	},
	DerivedProduct{
		Product: Product{
			ID:          6,
			ProductName: "Frangelico™",
			ProductType: Liqueur,
		},
		BaseProduct: Products[4],
	},
	DerivedProduct{
		Product: Product{
			ID:          7,
			ProductName: "Stonewall Kitchen™ Maine Maple Syrup",
			ProductType: Liqueur,
		},
		BaseProduct: Products[5],
	},
	DerivedProduct{
		Product: Product{
			ID:          8,
			ProductName: "Kahlúa™",
			ProductType: Liqueur,
		},
		BaseProduct: Products[6],
	},
	DerivedProduct{
		Product: Product{
			ID:          9,
			ProductName: "Malibu™ Coconut Rum",
			ProductType: Spirit,
		},
		BaseProduct: Products[9],
	},
	DerivedProduct{
		Product: Product{
			ID:          10,
			ProductName: "Disaronno™",
			ProductType: Liqueur,
		},
		BaseProduct: Products[10],
	},
	DerivedProduct{
		Product: Product{
			ID:          11,
			ProductName: "Taylor'd Milestones \"No.1 Classic\" Whiskey Glass",
			ProductType: Drinkware,
		},
		BaseProduct: Products[17],
	},
	DerivedProduct{
		Product: Product{
			ID:          12,
			ProductName: "OXO™ SteeL Cocktail Shaker",
			ProductType: Tool,
		},
		BaseProduct: Products[14],
	},
	DerivedProduct{
		Product: Product{
			ID:          13,
			ProductName: "Absolut™ Vodka",
			ProductType: Spirit,
		},
		BaseProduct: Products[12],
	},
	DerivedProduct{
		Product: Product{
			ID:          14,
			ProductName: "Luxardo™ Original Maraschino Cherries",
			ProductType: Garnish,
		},
		BaseProduct: Products[15],
	},
}
