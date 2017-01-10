package model

var Products = []Product{
	Product{
		ID:              1,
		ProductName:     "Orange Juice",
		ProductType:     Mixer,
		BDG:             Base,
		PreText:         "Fresh Squeezed",
		Description:     "On the top ten important things to have in a bar, orange juice is a must have. Fresh squeezed such is preferable.",
		ImagePath:       "https://s3.ca-central-1.amazonaws.com/commonwealthcocktailsimages",
		Image:           "orange-juice-glass-bottle-1-pint.jpg",
		ImageSourceName: "Mortons Dairies",
		ImageSourceLink: "http://www.mortonsdairies.co.uk/media/products/orange-juice-glass-bottle-1-pint.jpg",
	},
	Product{
		ID:              2,
		ProductName:     "Lemon Juice",
		ProductType:     Mixer,
		BDG:             Base,
		Description:     "Oh boy, this mixer is about as essential as it gets. Try to always use fresh squeezed juice! The one excuse for not using it is… No, we can't think of any.",
		PreText:         "Fresh Squeezed",
		ImagePath:       "https://s3.ca-central-1.amazonaws.com/commonwealthcocktailsimages",
		Image:           "lemon.png",
		ImageSourceName: "Dr Deborah Baker",
		ImageSourceLink: "http://www.drdeborahbaker.com/wp-content/uploads/2014/07/lemon.png",
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
		ID:              6,
		ProductName:     "Maple Syrup",
		ProductType:     Liqueur,
		BDG:             Base,
		ImagePath:       "https://s3.ca-central-1.amazonaws.com/commonwealthcocktailsimages",
		Image:           "Maple_Leaf_Bottle.jpg",
		ImageSourceName: "McCamus Maple Syrup",
		ImageSourceLink: "http://cdn.shopify.com/s/files/1/0628/8453/products/Leaf_Bottle.jpg?v=1415141353",
	},
	Product{
		ID:          7,
		ProductName: "Coffee Liqueur",
		ProductType: Liqueur,
		BDG:         Base,
	},
	Product{
		ID:              8,
		ProductName:     "Milk",
		ProductType:     Mixer,
		BDG:             Base,
		ImagePath:       "https://s3.ca-central-1.amazonaws.com/commonwealthcocktailsimages",
		Image:           "milk.png",
		ImageSourceName: "PngImg.com",
		ImageSourceLink: "http://pngimg.com/upload/milk_PNG12746.png",
	},
	Product{
		ID:              9,
		ProductName:     "Cream",
		ProductType:     Mixer,
		BDG:             Base,
		ImagePath:       "https://s3.ca-central-1.amazonaws.com/commonwealthcocktailsimages",
		Image:           "cream.jpg",
		ImageSourceName: "PngImg.com",
		ImageSourceLink: "http://pngimg.com/upload/bottle_PNG2934.png",
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
		ID:              12,
		ProductName:     "Ginger Root",
		ProductType:     Mixer,
		BDG:             Base,
		PostText:        "(thumbnail size)",
		ImagePath:       "https://s3.ca-central-1.amazonaws.com/commonwealthcocktailsimages",
		Image:           "Ginger.jpg",
		ImageSourceName: "LoveTheGarden.com",
		ImageSourceLink: "https://www.lovethegarden.com/sites/default/files/styles/full_width_700/public/images_and_media/Ginger.jpg?itok=sIbshAWY",
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
		ID:              25,
		ProductName:     "Tropicana™ Orange Juice",
		ProductType:     Mixer,
		BDG:             Derived,
		ImagePath:       "https://s3.ca-central-1.amazonaws.com/commonwealthcocktailsimages",
		Image:           "tropicana-orange-juice.jpg",
		ImageSourceName: "BreadBerry",
		ImageSourceLink: "https://jpg.breadberry.com/content/images/thumbs/0094209_tropicana-orange-juice-no-pulp-59-oz.jpeg",
	},
	Product{
		ID:              26,
		ProductName:     "ReaLemon™ Lemon Juice",
		ProductType:     Mixer,
		BDG:             Derived,
		Description:     "<p>Just about everyone recognizes ReaLemon and ReaLime – the distinctive, fruit-shaped squeeze bottles are a familiar sight to anyone who frequents grocery stores.</p>ReaLemon and ReaLime juices are made by concentrating the juice of high-quality fresh lemons and limes, and then adding back just the right amount of water to create regular-strength juice.</p>",
		Details:         "<p>ReaLemon Lemon Juice From Concentrate was first introduced in 1934 by Irving Swartzburg. ReaLime Juice From Concentrate followed in 1947. Since then, people have come to trust ReaLemon and ReaLime because they provide all the great taste of premium-quality lemons and limes, but are more convenient, more economical and more consistent in taste and strength than fresh fruit. </p><p>By the year 2000, these two brands had grown to dominate their category. In August 2001, ReaLemon and ReaLime became part of the Mott's family when they were acquired from Eagle Family Foods.</p><p>ReaLemon Lemon Juice From Concentrate is regular-strength juice and convenient to use with no slicing or squeezing required. ReaLemon is an easy way to add perfect lemon flavor to all your favorite dishes with consistent taste from bottle to bottle.</p><p>Similar to ReaLemon, ReaLime Lime Juice From Concentrate is regular-strength juice made from fresh, quality limes. ReaLime adds a perfect citrusy zing to beverages, marinades, meats, seafood and salads.</p><p>Today, ReaLemon and ReaLime are part of Plano, Texas-based Dr Pepper Snapple Group, an integrated refreshment beverage business marketing more than 50 beverage brands throughout North America.</p>",
		ImagePath:       "https://s3.ca-central-1.amazonaws.com/commonwealthcocktailsimages",
		Image:           "realemon.png",
		ImageSourceName: "Walmart",
		ImageSourceLink: "https://i5.walmartimages.com/asr/970ff170-2686-4083-97d2-592b034f893a_1.f688f66fe61a5574f1240c711867bb65.jpeg",
	},
	Product{
		ID:          27,
		ProductName: "Pappy Van Winkle's™ Bourbon Whiskey",
		ProductType: Spirit,
		BDG:         Derived,
	},
	Product{
		ID:              28,
		ProductName:     "Breckenridge™ Bourbon Whiskey",
		ProductType:     Spirit,
		BDG:             Derived,
		ImagePath:       "https://s3.ca-central-1.amazonaws.com/commonwealthcocktailsimages",
		Image:           "breckenridge-bourbon-whiskey.jpg",
		ImageSourceName: "Casker.com",
		ImageSourceLink: "https://media2.caskers.com/media/catalog/product/cache/1/thumbnail/1000x/9df78eab33525d08d6e5fb8d27136e95/b/r/breckenridge-bourbon-whiskey-1_2.jpg",
	},
	Product{
		ID:              29,
		ProductName:     "Bailey's Irish Cream™",
		ProductType:     Liqueur,
		BDG:             Derived,
		ImagePath:       "https://s3.ca-central-1.amazonaws.com/commonwealthcocktailsimages",
		Image:           "baileys.jpg",
		ImageSourceName: "The Whiskey Exchange",
		ImageSourceLink: "https://img.thewhiskyexchange.com/900/liq_bai1.jpg",
	},
	Product{
		ID:              30,
		ProductName:     "Frangelico™",
		ProductType:     Liqueur,
		BDG:             Derived,
		ImagePath:       "https://s3.ca-central-1.amazonaws.com/commonwealthcocktailsimages",
		Image:           "frangelico.jpg",
		ImageSourceName: "The Whiskey Exchange",
		ImageSourceLink: "https://img.thewhiskyexchange.com/540/liq_fra1.jpg",
	},
	Product{
		ID:              31,
		ProductName:     "Stonewall Kitchen™ Maine Maple Syrup",
		ProductType:     Liqueur,
		BDG:             Derived,
		ImagePath:       "https://s3.ca-central-1.amazonaws.com/commonwealthcocktailsimages",
		Image:           "swk_maple_syrup.jpg",
		ImageSourceName: "Stonewall Kitchen™",
		ImageSourceLink: "https://sits-pod32.demandware.net/dw/image/v2/AAYB_PRD/on/demandware.static/-/Sites-swk-catalog/default/dw760e9d78/images/170801.jpg?sw=500",
	},
	Product{
		ID:              32,
		ProductName:     "Kahlúa™",
		ProductType:     Liqueur,
		BDG:             Derived,
		ImagePath:       "https://s3.ca-central-1.amazonaws.com/commonwealthcocktailsimages",
		Image:           "kahlua.jpg",
		ImageSourceName: "The Whiskey Exchange",
		ImageSourceLink: "https://img.thewhiskyexchange.com/540/liq_kah1.jpg",
	},
	Product{
		ID:              33,
		ProductName:     "Malibu™ Coconut Rum",
		ProductType:     Spirit,
		BDG:             Derived,
		ImagePath:       "https://s3.ca-central-1.amazonaws.com/commonwealthcocktailsimages",
		Image:           "malibu_cr.jpg",
		ImageSourceName: "The Whiskey Exchange",
		ImageSourceLink: "https://img.thewhiskyexchange.com/540/liq_mal1.jpg",
	},
	Product{
		ID:              34,
		ProductName:     "Disaronno™",
		ProductType:     Liqueur,
		BDG:             Derived,
		ImagePath:       "https://s3.ca-central-1.amazonaws.com/commonwealthcocktailsimages",
		Image:           "disaronno.jpg",
		ImageSourceName: "The Whiskey Exchange",
		ImageSourceLink: "https://img.thewhiskyexchange.com/900/liq_ama1.jpg",
	},
	Product{
		ID:              35,
		ProductName:     "Taylor'd Milestones \"No.1 Classic\" Whiskey Glass",
		ProductType:     Drinkware,
		BDG:             Derived,
		Description:     "<p>10 oz capacity for enjoying a whiskey on the rocks or a vintage scotch served neat.  Each scotch glass arrives marked with our signature \"diamond\" laser etched on the base of the glass.  A symbol of quality and uniqueness in all Taylor'd Milestones Glassware.</p><p>This Whiskey Glass is the perfect size to partner with a large size ice ball or large ice cubes, still leaving room for your favorite spirit.</p><p>Crystal Clear, Strong & Durable Glasses For Your Home Bar</p><p>MADE IN THE USA</p><p>MEETS SERVICE INDUSTRY QUALITY STANDARDS</p><p>RESISTANT TO BREAKAGE AND SCRATCHING</p><p>DISHWASHER SAFE ( hand washing recommended for prolonged glass clarity )</p>",
		ImagePath:       "https://s3.ca-central-1.amazonaws.com/commonwealthcocktailsimages",
		Image:           "taylord_milstones_classic_10_oz_scotch_glass.jpg",
		ImageSourceName: "Taylor'd Milestones",
		ImageSourceLink: "http://cdn.shopify.com/s/files/1/1140/5810/products/scotch_rocks.jpg?v=1465840383",
		SourceName:      "Taylor'd Milestones",
		SourceLink:      "http://taylordmilestones.com/collections/scotch-whiskey-collection/products/10-oz-scotch-whiskey-glasses",
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
		BDG:         Group,
	},
	Product{
		ID:          40,
		ProductName: "Simple Syrup",
		ProductType: Mixer,
		BDG:         Base,
	},
	Product{
		ID:              41,
		ProductName:     "Bitters",
		ProductType:     Mixer,
		BDG:             Base,
		Description:     "A bitters is traditionally an alcoholic preparation flavored with botanical matter such that the end result is characterized by a bitter, sour, or bittersweet flavor. Numerous longstanding brands of bitters were originally developed as patent medicines, but are now sold as digestifs and cocktail flavorings.",
		ImagePath:       "https://s3.ca-central-1.amazonaws.com/commonwealthcocktailsimages",
		Image:           "bitters.jpg",
		ImageSourceName: "KegWorks",
		ImageSourceLink: "https://pull01-kegworks.netdna-ssl.com/media/catalog/product/cache/1/image/9df78eab33525d08d6e5fb8d27136e95/4/0/4000-18-bitters-bottle-b1_1.jpg",
	},
	Product{
		ID:              42,
		ProductName:     "Egg White",
		ProductType:     Mixer,
		BDG:             Base,
		ImagePath:       "https://s3.ca-central-1.amazonaws.com/commonwealthcocktailsimages",
		Image:           "eggs.jpg",
		ImageSourceName: "EatingWorks.com",
		ImageSourceLink: "EatingWorkshttp://eatingworks.com/wp-content/uploads/2016/02/eggs.jpg",
	},
	Product{
		ID:          43,
		ProductName: "Gin",
		ProductType: Spirit,
		BDG:         Base,
	},
	Product{
		ID:              44,
		ProductName:     "Angostura Aromatic Bitters",
		ProductType:     Mixer,
		BDG:             Derived,
		Description:     "Angostura bitters is a concentrated bitters, or botanically infused alcoholic mixture, made of water, 44.7% ethanol, gentian, herbs and spices, by House of Angostura in Trinidad and Tobago. It is typically used for flavouring beverages or food.",
		ImagePath:       "https://s3.ca-central-1.amazonaws.com/commonwealthcocktailsimages",
		Image:           "angostura_aromatic_bitters.jpg",
		ImageSourceName: "The Whisky Exchange",
		ImageSourceLink: "https://img.thewhiskyexchange.com/540/liq_bit1.jpg",
		Rating:          4,
	},
	Product{
		ID:          45,
		ProductName: "Spiced Rum",
		ProductType: Spirit,
		BDG:         Base,
	},
	Product{
		ID:              46,
		ProductName:     "The Kraken Black Spiced Rum",
		ProductType:     Spirit,
		BDG:             Derived,
		Description:     "A dark spiced Caribbean rum introduced to the UK in Spring 2010, Kraken's old-style bottle and superlative packaging seem to earmark it out for greatness, and perhaps it will help some of those many thousands of punters nursing an old-recipe-Sailor-Jerry-sized hole in their hearts.",
		Details:         "<ul><li>Kraken takes its name from a mythological sea beast that is said to attack ships sailing the Atlantic.</li><li>The rum in the spirit is made from molasses and is aged for 12 to 24 months.</li><li>The liquor is then flavored with a number of spices, including cinnamon, ginger and clove.</li></ul>",
		ImagePath:       "https://s3.ca-central-1.amazonaws.com/commonwealthcocktailsimages",
		Image:           "rum_kraken.jpg",
		ImageSourceName: "The Whisky Exchange",
		ImageSourceLink: "https://img.thewhiskyexchange.com/900/rum_kra3.jpg",
		SourceName:      "The Whisky Exchange",
		SourceLink:      "https://www.thewhiskyexchange.com/p/12021/kraken-black-spiced-rum",
		Rating:          5,
		Drink: []Meta{
			Metadata[17],
			Metadata[18],
			Metadata[19],
		},
	},
}

var ProductGroups = []GroupProduct{
	GroupProduct{
		Products: []Product{
			Products[2],
			Products[22],
			Products[23],
			Products[42],
			Products[44],
		},
		GroupProduct: Products[38],
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
	DerivedProduct{
		Product:     Products[45],
		BaseProduct: Products[44],
	},
}
