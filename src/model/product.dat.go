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
	Product{
		ID:          25,
		ProductName: "Tropicana™ Orange Juice",
		ProductType: Mixer,
	},
	Product{
		ID:          26,
		ProductName: "ReaLemon™ Lemon Juice",
		ProductType: Mixer,
	},
	Product{
		ID:          27,
		ProductName: "Pappy Van Winkle's™ Bourbon Whiskey",
		ProductType: Spirit,
	},
	Product{
		ID:          28,
		ProductName: "Breckenridge™ Bourbon Whiskey",
		ProductType: Spirit,
	},
	Product{
		ID:          29,
		ProductName: "Bailey's Irish Cream™",
		ProductType: Liqueur,
	},
	Product{
		ID:          30,
		ProductName: "Frangelico™",
		ProductType: Liqueur,
	},
	Product{
		ID:          31,
		ProductName: "Stonewall Kitchen™ Maine Maple Syrup",
		ProductType: Liqueur,
	},
	Product{
		ID:          32,
		ProductName: "Kahlúa™",
		ProductType: Liqueur,
	},
	Product{
		ID:          33,
		ProductName: "Malibu™ Coconut Rum",
		ProductType: Spirit,
	},
	Product{
		ID:          34,
		ProductName: "Disaronno™",
		ProductType: Liqueur,
	},
	Product{
		ID:          35,
		ProductName: "Taylor'd Milestones \"No.1 Classic\" Whiskey Glass",
		ProductType: Drinkware,
	},
	Product{
		ID:          36,
		ProductName: "OXO™ SteeL Cocktail Shaker",
		ProductType: Tool,
	},
	Product{
		ID:          37,
		ProductName: "Absolut™ Vodka",
		ProductType: Spirit,
	},
	Product{
		ID:          38,
		ProductName: "Luxardo™ Original Maraschino Cherries",
		ProductType: Garnish,
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
}
