//model/advertisement.go
package model

type Advertisement struct {
	ID             int
	AdvertiserName string
	Article        string
	Blurb          string
	Image          string
	Cocktails      []int
	Products       []ProdcutAdvertisement
}

type ProdcutAdvertisement struct {
	BaseProductID       int
	AdvertisedProductID int
}
