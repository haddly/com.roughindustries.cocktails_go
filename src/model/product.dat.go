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
}
