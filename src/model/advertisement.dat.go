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
				BaseProduct:       Products[4],
				AdvertisedProduct: Products[6],
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
				BaseProduct:       Products[4],
				AdvertisedProduct: Products[5],
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
				BaseProduct:       Products[2],
				AdvertisedProduct: Products[3],
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
				AdvertisedProduct: Products[1],
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
				BaseProduct:       Products[7],
				AdvertisedProduct: Products[8],
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
				BaseProduct:       Products[9],
				AdvertisedProduct: Products[10],
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
				BaseProduct:       Products[9],
				AdvertisedProduct: Products[10],
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
				BaseProduct:       Products[11],
				AdvertisedProduct: Products[12],
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
				BaseProduct:       Products[13],
				AdvertisedProduct: Products[14],
			},
			ProdcutAdvertisement{
				BaseProduct:       Products[17],
				AdvertisedProduct: Products[18],
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
				BaseProduct:       Products[13],
				AdvertisedProduct: Products[14],
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
				BaseProduct:       Products[19],
				AdvertisedProduct: Products[20],
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
				BaseProduct:       Products[27],
				AdvertisedProduct: Products[28],
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
				BaseProduct:       Products[24],
				AdvertisedProduct: Products[30],
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
				BaseProduct:       Products[22],
				AdvertisedProduct: Products[34],
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
				BaseProduct:       Products[25],
				AdvertisedProduct: Products[35],
			},
		},
	},
}
