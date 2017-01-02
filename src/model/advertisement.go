//model/advertisement.go
package model

type Advertisement struct {
	ID             int
	AdvertiserName string
	AdvertiserLink string
	Article        string
	Blurb          string
	Image          string
	Cocktails      []Cocktail
	Products       []ProdcutAdvertisement
}

type ProdcutAdvertisement struct {
	BaseProduct       Product
	AdvertisedProduct Product
}
