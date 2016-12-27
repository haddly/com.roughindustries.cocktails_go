package model

var Advertisements = []Advertisement{
	Advertisement{
		ID:             1,
		AdvertiserName: "Breckenridge™ Distillery",
		AdvertiserLink: "http://www.breckenridgedistillery.com/",
		Article:        "",
		Blurb:          "",
		Image:          "breckenridgedistillery-02.png",
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
		ID:             2,
		AdvertiserName: "Pappy Van Winkle's™ Distillery",
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
		ID:             3,
		AdvertiserName: "ReaLemon™",
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
		ID:             4,
		AdvertiserName: "Tropicana™",
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
	Advertisement{
		ID:             5,
		AdvertiserName: "Bailey and Co™",
		Article:        "",
		Blurb:          "",
		Image:          "",
		Products: []ProdcutAdvertisement{
			ProdcutAdvertisement{
				BaseProductID:       8,
				AdvertisedProductID: 9,
			},
		},
	},
	Advertisement{
		ID:             6,
		AdvertiserName: "Frangelico™",
		AdvertiserLink: "http://frangelico.com/",
		Article:        "",
		Blurb:          "",
		Image:          "frangelico.jpg",
		Cocktails: []int{
			3,
		},
	},
}
