package model

var Advertisers = []Advertiser{
	Advertiser{
		ID:             1,
		AdvertiserName: "Breckenridge™ Distillery",
		AdvertiserLink: "http://www.breckenridgedistillery.com/",
	},
	Advertiser{
		ID:             2,
		AdvertiserName: "Pappy Van Winkle's™ Distillery",
	},
	Advertiser{
		ID:             3,
		AdvertiserName: "ReaLemon™",
	},
	Advertiser{
		ID:             4,
		AdvertiserName: "Tropicana™",
	},
	Advertiser{
		ID:             5,
		AdvertiserName: "Bailey and Co™",
	},
	Advertiser{
		ID:             6,
		AdvertiserName: "Frangelico™",
		AdvertiserLink: "http://frangelico.com/",
	},
	Advertiser{
		ID:             7,
		AdvertiserName: "Stonewall Kitchen™",
	},
	Advertiser{
		ID:             8,
		AdvertiserName: "Pernod Ricard",
	},
	Advertiser{
		ID:             9,
		AdvertiserName: "The Kahlúa™ Company",
		AdvertiserLink: "http://www.kahlua.com/",
	},
	Advertiser{
		ID:                10,
		AdvertiserCompany: "ILLVA SARONNO S.p.A",
		AdvertiserName:    "Disaronno",
		AdvertiserLink:    "http://www.disaronno.com/",
	},
	Advertiser{
		ID:             11,
		AdvertiserName: "Taylor'd Milestones",
	},
	Advertiser{
		ID:             12,
		AdvertiserName: "OXO",
	},
	Advertiser{
		ID:             13,
		AdvertiserName: "The Absolut™ Company",
	},
	Advertiser{
		ID:             14,
		AdvertiserName: "Anchor Distilling Company",
	},
	Advertiser{
		ID:             15,
		AdvertiserName: "Commonwealth Cocktails",
	},
}

var ProductAdvertisements = []ProductAdvertisement{
	ProductAdvertisement{
		BaseProduct:       Products[2],
		AdvertisedProduct: DerivedProducts[3].Product,
	},
	ProductAdvertisement{
		BaseProduct:       Products[2],
		AdvertisedProduct: DerivedProducts[2].Product,
	},
	ProductAdvertisement{
		BaseProduct:       Products[1],
		AdvertisedProduct: DerivedProducts[1].Product,
	},
	ProductAdvertisement{
		BaseProduct:       Products[0],
		AdvertisedProduct: DerivedProducts[0].Product,
	},
	ProductAdvertisement{
		BaseProduct:       Products[3],
		AdvertisedProduct: DerivedProducts[4].Product,
	},
	ProductAdvertisement{
		BaseProduct:       Products[4],
		AdvertisedProduct: DerivedProducts[5].Product,
	},
	ProductAdvertisement{
		BaseProduct:       Products[5],
		AdvertisedProduct: DerivedProducts[6].Product,
	},
	ProductAdvertisement{
		BaseProduct:       Products[6],
		AdvertisedProduct: DerivedProducts[7].Product,
	},
	ProductAdvertisement{
		BaseProduct:       Products[9],
		AdvertisedProduct: DerivedProducts[8].Product,
	},
	ProductAdvertisement{
		BaseProduct:       Products[10],
		AdvertisedProduct: DerivedProducts[9].Product,
	},
}

var Advertisements = []Advertisement{
	Advertisement{
		ID:                1,
		Advertiser:        Advertisers[0],
		MediumHorSnippet:  "<a target=\"_blank\" href=\"http://www.breckenridgedistillery.com/\"><img class=\"img-responsive\" src=\"https://s3.ca-central-1.amazonaws.com/commonwealthcocktailsimages/breckenridgedistillery-02.png\" alt=\"Breckenridge™ Distillery\" /></a>",
		MediumVertSnippet: "<img class=\"img-responsive\" src=\"https://s3.ca-central-1.amazonaws.com/commonwealthcocktailsimages/BreckDistillery_ad.png\"/>",
		Cocktails: []Cocktail{
			Cocktails[1],
		},
		Products: []ProductAdvertisement{
			ProductAdvertisement{
				BaseProduct:       Products[2],
				AdvertisedProduct: DerivedProducts[3].Product,
			},
		},
		AdType: CocktailAds,
	},
	Advertisement{
		ID:         2,
		Advertiser: Advertisers[1],
		Products: []ProductAdvertisement{
			ProductAdvertisement{
				BaseProduct:       Products[2],
				AdvertisedProduct: DerivedProducts[2].Product,
			},
		},
	},
	Advertisement{
		ID:         3,
		Advertiser: Advertisers[2],
		Products: []ProductAdvertisement{
			ProductAdvertisement{
				BaseProduct:       Products[1],
				AdvertisedProduct: DerivedProducts[1].Product,
			},
		},
	},
	Advertisement{
		ID:         4,
		Advertiser: Advertisers[3],
		Products: []ProductAdvertisement{
			ProductAdvertisement{
				BaseProduct:       Products[0],
				AdvertisedProduct: DerivedProducts[0].Product,
			},
		},
	},
	Advertisement{
		ID:         5,
		Advertiser: Advertisers[4],
		Products: []ProductAdvertisement{
			ProductAdvertisement{
				BaseProduct:       Products[3],
				AdvertisedProduct: DerivedProducts[4].Product,
			},
		},
	},
	Advertisement{
		ID:                6,
		Advertiser:        Advertisers[5],
		MediumHorSnippet:  "<a target=\"_blank\" href=\"http://frangelico.com/\"><img class=\"img-responsive\" src=\"https://s3.ca-central-1.amazonaws.com/commonwealthcocktailsimages/frangelico_logo.jpg\" alt=\"Frangelico™\" /></a>",
		MediumVertSnippet: "<img class=\"img-responsive\" src=\"https://s3.ca-central-1.amazonaws.com/commonwealthcocktailsimages/frangelico_ad.jpg\"/>",
		Cocktails: []Cocktail{
			Cocktails[2],
		},
		Products: []ProductAdvertisement{
			ProductAdvertisement{
				BaseProduct:       Products[4],
				AdvertisedProduct: DerivedProducts[5].Product,
			},
		},
		AdType: CocktailAds,
	},
	Advertisement{
		ID:         7,
		Advertiser: Advertisers[5],
		Products: []ProductAdvertisement{
			ProductAdvertisement{
				BaseProduct:       Products[4],
				AdvertisedProduct: DerivedProducts[5].Product,
			},
		},
	},
	Advertisement{
		ID:                8,
		Advertiser:        Advertisers[6],
		MediumHorSnippet:  "<a target=\"_blank\" href=\"http://www.stonewallkitchen.com/\"><img class=\"img-responsive\" src=\"https://s3.ca-central-1.amazonaws.com/commonwealthcocktailsimages/stonewall-kitchen-logo_header.png\" alt=\"Stonewall Kitchen™\" /></a>",
		MediumVertSnippet: "<img class=\"img-responsive\" src=\"https://s3.ca-central-1.amazonaws.com/commonwealthcocktailsimages/stonewall-kitchen-logo_vertrec.png\"/>",
		Products: []ProductAdvertisement{
			ProductAdvertisement{
				BaseProduct:       Products[5],
				AdvertisedProduct: DerivedProducts[6].Product,
			},
			ProductAdvertisement{
				BaseProduct: Products[30],
			},
		},
		AdType: ProductAds,
	},
	Advertisement{
		ID:         9,
		Advertiser: Advertisers[7],
		Products: []ProductAdvertisement{
			ProductAdvertisement{
				BaseProduct:       Products[6],
				AdvertisedProduct: DerivedProducts[7].Product,
			},
			ProductAdvertisement{
				BaseProduct:       Products[9],
				AdvertisedProduct: DerivedProducts[8].Product,
			},
		},
		AdType: ProductAds,
	},
	Advertisement{
		ID:               10,
		Advertiser:       Advertisers[8],
		MediumHorSnippet: "<a target=\"_blank\" href=\"http://www.kahlua.com/\"><img class=\"img-responsive\" src=\"https://s3.ca-central-1.amazonaws.com/commonwealthcocktailsimages/kahlua_logo.png\" alt=\"The Kahlúa™ Company\" /></a>",
		Cocktails: []Cocktail{
			Cocktails[0],
		},
		Products: []ProductAdvertisement{
			ProductAdvertisement{
				BaseProduct:       Products[6],
				AdvertisedProduct: DerivedProducts[7].Product,
			},
		},
		AdType: CocktailAds,
	},
	Advertisement{
		ID:                11,
		Advertiser:        Advertisers[9],
		MediumHorSnippet:  "<a target=\"_blank\" href=\"http://www.disaronno.com/\"><img class=\"img-responsive\" src=\"https://s3.ca-central-1.amazonaws.com/commonwealthcocktailsimages/Disaronno-Logo.png\" alt=\"Disaronno\" /></a>",
		MediumVertSnippet: "<img class=\"img-responsive\" src=\"https://s3.ca-central-1.amazonaws.com/commonwealthcocktailsimages/disaronno_ad.jpg\"/>",
		Cocktails: []Cocktail{
			Cocktails[4],
		},
		Products: []ProductAdvertisement{
			ProductAdvertisement{
				BaseProduct:       Products[10],
				AdvertisedProduct: DerivedProducts[9].Product,
			},
		},
		AdType: CocktailAds,
	},
	Advertisement{
		ID:         12,
		Advertiser: Advertisers[10],
		Products: []ProductAdvertisement{
			ProductAdvertisement{
				BaseProduct:       Products[17],
				AdvertisedProduct: DerivedProducts[10].Product,
			},
		},
	},
	Advertisement{
		ID:         13,
		Advertiser: Advertisers[11],
		Products: []ProductAdvertisement{
			ProductAdvertisement{
				BaseProduct:       Products[14],
				AdvertisedProduct: DerivedProducts[11].Product,
			},
		},
	},
	Advertisement{
		ID:         14,
		Advertiser: Advertisers[12],
		Products: []ProductAdvertisement{
			ProductAdvertisement{
				BaseProduct:       Products[12],
				AdvertisedProduct: DerivedProducts[12].Product,
			},
		},
	},
	Advertisement{
		ID:         15,
		Advertiser: Advertisers[13],
		Products: []ProductAdvertisement{
			ProductAdvertisement{
				BaseProduct:       Products[15],
				AdvertisedProduct: DerivedProducts[13].Product,
			},
		},
	},
	Advertisement{
		ID: 16,
		Products: []ProductAdvertisement{
			ProductAdvertisement{
				BaseProduct:       Products[40],
				AdvertisedProduct: DerivedProducts[14].Product,
			},
		},
	},
	Advertisement{
		ID:                17,
		Advertiser:        Advertisers[9],
		MediumHorSnippet:  "<a target=\"_blank\" href=\"http://www.disaronno.com/\"><img class=\"img-responsive\" src=\"https://s3.ca-central-1.amazonaws.com/commonwealthcocktailsimages/Disaronno-Logo.png\" alt=\"Disaronno\" /></a>",
		MediumVertSnippet: "<img class=\"img-responsive\" src=\"https://s3.ca-central-1.amazonaws.com/commonwealthcocktailsimages/disaronno_ad.jpg\"/>",
		Products: []ProductAdvertisement{
			ProductAdvertisement{
				BaseProduct: Products[33],
			},
			ProductAdvertisement{
				BaseProduct: Products[10],
			},
		},
		AdType: ProductAds,
	},
}
