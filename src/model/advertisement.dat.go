package model

var Advertisements = []Advertisement{
	Advertisement{
		ID:             0,
		AdvertiserName: "Breckenridge™ Distillery",
		Article:        "",
		Blurb:          "",
		Image:          "",
		Cocktails: []int{
			2,
		},
		Products: []ProdcutAdvertisement{
			ProdcutAdvertisement{
				BaseProductID:       5,
				AdvertisedProductID: 7,
			},
		},
	},
	Advertisement{
		ID:             1,
		AdvertiserName: "Breckenridge™ Distillery",
		Article:        "",
		Blurb:          "",
		Image:          "",
		Products: []ProdcutAdvertisement{
			ProdcutAdvertisement{
				BaseProductID:       5,
				AdvertisedProductID: 6,
			},
		},
	},
	Advertisement{
		ID:             2,
		AdvertiserName: "ReaLemon™ Distillery",
		Article:        "",
		Blurb:          "",
		Image:          "",
		Products: []ProdcutAdvertisement{
			ProdcutAdvertisement{
				BaseProductID:       3,
				AdvertisedProductID: 4,
			},
		},
	},
	Advertisement{
		ID:             3,
		AdvertiserName: "Tropicana™ Distillery",
		Article:        "",
		Blurb:          "",
		Image:          "",
		Products: []ProdcutAdvertisement{
			ProdcutAdvertisement{
				BaseProductID:       1,
				AdvertisedProductID: 2,
			},
		},
	},
}
