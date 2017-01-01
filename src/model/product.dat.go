package model

var Products = []Product{
	Product{
		ID:          1,
		ProductName: "Orange Juice",
		ProductType: Mixer,
		BDG:         Base,
		PreText:     "Fresh Squeezed",
	},
	Product{
		ID:          2,
		ProductName: "Lemon Juice",
		ProductType: Mixer,
		BDG:         Base,
	},
	Product{
		ID:          3,
		ProductName: "Bourbon Whiskey",
		ProductType: Spirit,
		BDG:         Base,
	},
	Product{
		ID:          4,
		ProductName: "Irish Cream Liqueur",
		ProductType: Liqueur,
		BDG:         Base,
	},
	Product{
		ID:          5,
		ProductName: "Hazelnut Liqueur",
		ProductType: Liqueur,
		BDG:         Base,
	},
	Product{
		ID:          6,
		ProductName: "Maple Syrup",
		ProductType: Liqueur,
		BDG:         Base,
	},
	Product{
		ID:          7,
		ProductName: "Coffee Liqueur",
		ProductType: Liqueur,
		BDG:         Base,
	},
	Product{
		ID:          8,
		ProductName: "Milk",
		ProductType: Mixer,
		BDG:         Base,
	},
	Product{
		ID:          9,
		ProductName: "Cream",
		ProductType: Mixer,
		BDG:         Base,
	},
	Product{
		ID:          10,
		ProductName: "Coconut Rum",
		ProductType: Spirit,
		BDG:         Base,
	},
	Product{
		ID:          11,
		ProductName: "Amaretto",
		ProductType: Liqueur,
		BDG:         Base,
	},
	Product{
		ID:          12,
		ProductName: "Ginger Root",
		ProductType: Mixer,
		BDG:         Base,
		PostText:    "(thumbnail size)",
	},
	Product{
		ID:          13,
		ProductName: "Vodka",
		ProductType: Spirit,
		BDG:         Base,
	},
	Product{
		ID:          14,
		ProductName: "Cola",
		ProductType: Mixer,
		BDG:         Base,
	},
	Product{
		ID:          15,
		ProductName: "Shaker",
		ProductType: Tool,
		BDG:         Base,
	},
	Product{
		ID:          16,
		ProductName: "Cherry",
		ProductType: Garnish,
		BDG:         Base,
	},
	Product{
		ID:          17,
		ProductName: "Starfruit",
		ProductType: Garnish,
		BDG:         Base,
		PreText:     "Slice of",
	},
	Product{
		ID:          18,
		ProductName: "Old Fashioned",
		ProductType: Drinkware,
		BDG:         Base,
	},
	Product{
		ID:          19,
		ProductName: "Muddler",
		ProductType: Tool,
		BDG:         Base,
	},
	Product{
		ID:          20,
		ProductName: "Collins",
		ProductType: Drinkware,
		BDG:         Base,
	},
	Product{
		ID:          21,
		ProductName: "Highball",
		ProductType: Drinkware,
		BDG:         Base,
	},
	Product{
		ID:          22,
		ProductName: "Orange",
		ProductType: Garnish,
		BDG:         Base,
		PostText:    "Zest Twist",
	},
	Product{
		ID:          23,
		ProductName: "Rye Whiskey",
		ProductType: Spirit,
		BDG:         Base,
	},
	Product{
		ID:          24,
		ProductName: "Tennessee Whiskey",
		ProductType: Spirit,
		BDG:         Base,
	},
	Product{
		ID:          25,
		ProductName: "Tropicana™ Orange Juice",
		ProductType: Mixer,
		BDG:         Derived,
	},
	Product{
		ID:          26,
		ProductName: "ReaLemon™ Lemon Juice",
		ProductType: Mixer,
		BDG:         Derived,
	},
	Product{
		ID:          27,
		ProductName: "Pappy Van Winkle's™ Bourbon Whiskey",
		ProductType: Spirit,
		BDG:         Derived,
	},
	Product{
		ID:          28,
		ProductName: "Breckenridge™ Bourbon Whiskey",
		ProductType: Spirit,
		BDG:         Derived,
	},
	Product{
		ID:          29,
		ProductName: "Bailey's Irish Cream™",
		ProductType: Liqueur,
		BDG:         Derived,
	},
	Product{
		ID:          30,
		ProductName: "Frangelico™",
		ProductType: Liqueur,
		BDG:         Derived,
	},
	Product{
		ID:          31,
		ProductName: "Stonewall Kitchen™ Maine Maple Syrup",
		ProductType: Liqueur,
		BDG:         Derived,
	},
	Product{
		ID:          32,
		ProductName: "Kahlúa™",
		ProductType: Liqueur,
		BDG:         Derived,
	},
	Product{
		ID:          33,
		ProductName: "Malibu™ Coconut Rum",
		ProductType: Spirit,
		BDG:         Derived,
	},
	Product{
		ID:          34,
		ProductName: "Disaronno™",
		ProductType: Liqueur,
		BDG:         Derived,
	},
	Product{
		ID:          35,
		ProductName: "Taylor'd Milestones \"No.1 Classic\" Whiskey Glass",
		ProductType: Drinkware,
		BDG:         Derived,
	},
	Product{
		ID:          36,
		ProductName: "OXO™ SteeL Cocktail Shaker",
		ProductType: Tool,
		BDG:         Derived,
	},
	Product{
		ID:          37,
		ProductName: "Absolut™ Vodka",
		ProductType: Spirit,
		BDG:         Derived,
	},
	Product{
		ID:          38,
		ProductName: "Luxardo™ Original Maraschino Cherries",
		ProductType: Garnish,
		BDG:         Derived,
	},
	Product{
		ID:          39,
		ProductName: "Brandy, whisk(e)y, gin, rum etc.",
		ProductType: Spirit,
		BDG:         Generic,
	},
	Product{
		ID:          40,
		ProductName: "Simple Syrup",
		ProductType: Mixer,
		BDG:         Base,
	},
	Product{
		ID:          41,
		ProductName: "Bitters",
		ProductType: Mixer,
		BDG:         Base,
	},
	Product{
		ID:          42,
		ProductName: "Egg White",
		ProductType: Mixer,
		BDG:         Base,
	},
	Product{
		ID:          43,
		ProductName: "Gin",
		ProductType: Spirit,
		BDG:         Base,
	},
	Product{
		ID:          44,
		ProductName: "Angostura Aromatic Bitters",
		ProductType: Mixer,
		BDG:         Derived,
	},
}

var GenericProducts = []GenericProduct{
	GenericProduct{
		Products: []Product{
			Products[2],
			Products[22],
			Products[23],
			Products[42],
		},
		GenericProduct: Products[38],
	},
}

var DerivedProducts = []DerivedProduct{
	DerivedProduct{
		Product:     Products[24],
		BaseProduct: Products[0],
	},
	DerivedProduct{
		Product:     Products[25],
		BaseProduct: Products[1],
	},
	DerivedProduct{
		Product:     Products[26],
		BaseProduct: Products[2],
	},
	DerivedProduct{
		Product:     Products[27],
		BaseProduct: Products[2],
	},
	DerivedProduct{
		Product:     Products[28],
		BaseProduct: Products[3],
	},
	DerivedProduct{
		Product:     Products[29],
		BaseProduct: Products[4],
	},
	DerivedProduct{
		Product:     Products[30],
		BaseProduct: Products[5],
	},
	DerivedProduct{
		Product:     Products[31],
		BaseProduct: Products[6],
	},
	DerivedProduct{
		Product:     Products[32],
		BaseProduct: Products[9],
	},
	DerivedProduct{
		Product:     Products[33],
		BaseProduct: Products[10],
	},
	DerivedProduct{
		Product:     Products[34],
		BaseProduct: Products[17],
	},
	DerivedProduct{
		Product:     Products[35],
		BaseProduct: Products[14],
	},
	DerivedProduct{
		Product:     Products[36],
		BaseProduct: Products[12],
	},
	DerivedProduct{
		Product:     Products[37],
		BaseProduct: Products[15],
	},
	DerivedProduct{
		Product:     Products[43],
		BaseProduct: Products[40],
	},
}
