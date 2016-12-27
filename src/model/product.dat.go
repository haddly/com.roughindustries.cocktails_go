package model

var Products = []Product{
	Product{
		ID:          1,
		ProductName: "Orange Juice",
		ProductType: Mixer,
		IsBase:      true,
	},
	Product{
		ID:          2,
		ProductName: "Tropicana™ Orange Juice",
		ProductType: Mixer,
		IsBase:      false,
	},
	Product{
		ID:          3,
		ProductName: "Lemon Juice",
		ProductType: Mixer,
		IsBase:      true,
	},
	Product{
		ID:          4,
		ProductName: "ReaLemon™ Lemon Juice",
		ProductType: Mixer,
		IsBase:      false,
	},
	Product{
		ID:          5,
		ProductName: "Bourbon Whiskey",
		ProductType: Spirit,
		IsBase:      true,
	},
	Product{
		ID:          6,
		ProductName: "Pappy Van Winkle's™ Bourbon Whiskey",
		ProductType: Spirit,
		IsBase:      false,
	},
	Product{
		ID:          7,
		ProductName: "Breckenridge™ Bourbon Whiskey",
		ProductType: Spirit,
		IsBase:      false,
	},
	Product{
		ID:          8,
		ProductName: "Irish Cream Liqueur",
		ProductType: Liqueur,
		IsBase:      true,
	},
	Product{
		ID:          9,
		ProductName: "Bailey's Irish Cream™",
		ProductType: Liqueur,
		IsBase:      false,
	},
	Product{
		ID:          10,
		ProductName: "Hazelnut Liqueur",
		ProductType: Liqueur,
		IsBase:      true,
	},
	Product{
		ID:          11,
		ProductName: "Frangelico™",
		ProductType: Liqueur,
		IsBase:      false,
	},
	Product{
		ID:          12,
		ProductName: "Maple Syrup",
		ProductType: Liqueur,
		IsBase:      true,
	},
	Product{
		ID:          13,
		ProductName: "Stonewall Kitchen™ Maine Maple Syrup",
		ProductType: Liqueur,
		IsBase:      false,
	},
}
