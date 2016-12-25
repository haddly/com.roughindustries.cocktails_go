package model

var Products = []Product{
	Product{
		ID:          0,
		ProductName: "Orange Juice",
		ProductType: Mixer,
		IsBase:      true,
		Overrider:   1,
	},
	Product{
		ID:          1,
		ProductName: "Tropicana Orange Juice",
		ProductType: Mixer,
		IsBase:      false,
	},
	Product{
		ID:          2,
		ProductName: "Lemon Juice",
		ProductType: Mixer,
		IsBase:      true,
		Overrider:   3,
	},
	Product{
		ID:          3,
		ProductName: "ReaLemon™ Lemon Juice",
		ProductType: Mixer,
		IsBase:      false,
	},
}
