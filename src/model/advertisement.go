//model/advertisement.go
package model

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
	Image             string
	ImagePath         string
	Image600x150      string
	ImagePath600x150  string
	ImageBanner       string
	ImagePathBanner   string
	Cocktails         []Cocktail
	Products          []ProductAdvertisement
	AdType            AdType
	Page              string
}

type ProductAdvertisement struct {
	BaseProduct       Product
	AdvertisedProduct Product
}
