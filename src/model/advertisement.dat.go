package model

var Advertisements = []Advertisement{
	Advertisement{
		ID:             1,
		AdvertiserName: "Breckenridge™ Distillery",
		AdvertiserLink: "http://www.breckenridgedistillery.com/",
		Article:        "",
		Blurb:          "",
		Image:          "breckenridgedistillery-02.png",
		Cocktails: []Cocktail{
			Cocktails[1],
		},
		Products: []ProdcutAdvertisement{
			ProdcutAdvertisement{
				BaseProduct:       Products[2],
				AdvertisedProduct: DerivedProducts[3],
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
				BaseProduct:       Products[2],
				AdvertisedProduct: DerivedProducts[2],
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
				BaseProduct:       Products[1],
				AdvertisedProduct: DerivedProducts[1],
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
				BaseProduct:       Products[0],
				AdvertisedProduct: DerivedProducts[0],
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
				BaseProduct:       Products[3],
				AdvertisedProduct: DerivedProducts[4],
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
		Cocktails: []Cocktail{
			Cocktails[2],
		},
		Products: []ProdcutAdvertisement{
			ProdcutAdvertisement{
				BaseProduct:       Products[4],
				AdvertisedProduct: DerivedProducts[5],
			},
		},
	},
	Advertisement{
		ID:             7,
		AdvertiserName: "Frangelico™",
		Article:        "",
		Blurb:          "",
		Image:          "",
		Products: []ProdcutAdvertisement{
			ProdcutAdvertisement{
				BaseProduct:       Products[4],
				AdvertisedProduct: DerivedProducts[5],
			},
		},
	},
	Advertisement{
		ID:             8,
		AdvertiserName: "Stonewall Kitchen™",
		Article:        "",
		Blurb:          "",
		Image:          "",
		Products: []ProdcutAdvertisement{
			ProdcutAdvertisement{
				BaseProduct:       Products[5],
				AdvertisedProduct: DerivedProducts[6],
			},
		},
	},
	Advertisement{
		ID:             9,
		AdvertiserName: "Pernod Ricard",
		Article:        "",
		Blurb:          "",
		Image:          "",
		Products: []ProdcutAdvertisement{
			ProdcutAdvertisement{
				BaseProduct:       Products[6],
				AdvertisedProduct: DerivedProducts[7],
			},
			ProdcutAdvertisement{
				BaseProduct:       Products[9],
				AdvertisedProduct: DerivedProducts[8],
			},
		},
	},
	Advertisement{
		ID:             10,
		AdvertiserName: "The Kahlúa™ Company",
		Article:        "",
		Blurb:          "",
		Image:          "kahlua.png",
		Cocktails: []Cocktail{
			Cocktails[0],
		},
		Products: []ProdcutAdvertisement{
			ProdcutAdvertisement{
				BaseProduct:       Products[6],
				AdvertisedProduct: DerivedProducts[7],
			},
		},
	},
	Advertisement{
		ID:             11,
		AdvertiserName: "ILLVA SARONNO S.p.A",
		Article:        "",
		Blurb:          "",
		Image:          "",
		Products: []ProdcutAdvertisement{
			ProdcutAdvertisement{
				BaseProduct:       Products[10],
				AdvertisedProduct: DerivedProducts[9],
			},
		},
	},
	Advertisement{
		ID:             12,
		AdvertiserName: "Taylor'd Milestones",
		Article:        "",
		Blurb:          "",
		Image:          "",
		Products: []ProdcutAdvertisement{
			ProdcutAdvertisement{
				BaseProduct:       Products[17],
				AdvertisedProduct: DerivedProducts[10],
			},
		},
	},
	Advertisement{
		ID:             12,
		AdvertiserName: "OXO",
		Article:        "",
		Blurb:          "",
		Image:          "",
		Products: []ProdcutAdvertisement{
			ProdcutAdvertisement{
				BaseProduct:       Products[14],
				AdvertisedProduct: DerivedProducts[11],
			},
		},
	},
	Advertisement{
		ID:             13,
		AdvertiserName: "The Absolut™ Company",
		Article:        "",
		Blurb:          "",
		Image:          "",
		Products: []ProdcutAdvertisement{
			ProdcutAdvertisement{
				BaseProduct:       Products[12],
				AdvertisedProduct: DerivedProducts[12],
			},
		},
	},
	Advertisement{
		ID:             14,
		AdvertiserName: "Anchor Distilling Company",
		Article:        "",
		Blurb:          "",
		Image:          "",
		Products: []ProdcutAdvertisement{
			ProdcutAdvertisement{
				BaseProduct:       Products[15],
				AdvertisedProduct: DerivedProducts[13],
			},
		},
	},
}
