//model/advertisement.go
package model

import (
	"html/template"
)

type AdType int

const (
	ProductAds = 1 + iota
	CocktailAds
	ProductPageAds
	CocktailPageAds
)

var AdTypeStrings = [...]string{
	"ProductAds",      //
	"CocktailAds",     //
	"ProductPageAds",  //
	"CocktailPageAds", //
}

// String returns the English name of the adtype ("Product", "Cocktail", ...).
func (at AdType) String() string { return AdTypeStrings[at-1] }

type Advertisement struct {
	ID                int
	AdvertiserCompany string
	AdvertiserName    string
	AdvertiserLink    string
	HeaderSnippet     template.HTML
	BannerAdSnippet   template.HTML
	VertRecSnippet    template.HTML
	Cocktails         []Cocktail
	Products          []ProductAdvertisement
	AdType            AdType
	Page              string
}

type ProductAdvertisement struct {
	BaseProduct       Product
	AdvertisedProduct Product
}
