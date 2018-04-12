// Copyright 2017 Rough Industries LLC. All rights reserved.
//controller/www/images.go:

package www

import (
	"github.com/asaskevich/govalidator"
	"github.com/fogleman/gg"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

//Image Modification Form page handler which displays the Image Modification
//Form page.
func ImageModFormHandler(w http.ResponseWriter, r *http.Request, page *page) {
	cba := page.Cocktail.SelectCocktailsByAlphaNums(false, page.View)
	page.CocktailsByAlphaNums = cba
	pbt := page.Product.SelectProductsByTypes(true, true, false, page.View)
	page.ProductsByTypes = pbt
	if page.Cocktail.ID == 0 {
		if page.Product.ID == 0 {
			//apply the template page info to the index page
			page.RenderPageTemplate(w, r, "imagemodform")
		} else {
			page.Product = page.Product.SelectProductByID(page.Product.ID, page.View)
			page.RenderPageTemplate(w, r, "imagemodform")
		}
	} else {
		out := page.Cocktail.SelectCocktailsByID(page.Cocktail.ID, false, page.View)
		page.Cocktail = out.Cocktail
		page.RenderPageTemplate(w, r, "imagemodform")
	}

}

//Image modification form page request handler which process the image
//modification request.
func ImageModHandler(w http.ResponseWriter, r *http.Request, page *page) {
	url := page.Image.ImageSource
	// don't worry about errors
	response, e := http.Get(url)
	if e != nil {
		log.Errorln(e)
		page.RenderPageTemplate(w, r, "404")
		return
	}

	cba := page.Cocktail.SelectCocktailsByAlphaNums(false, page.View)
	page.CocktailsByAlphaNums = cba
	pbt := page.Product.SelectProductsByTypes(true, true, false, page.View)
	page.ProductsByTypes = pbt

	defer response.Body.Close()

	//open a file for writing
	t, _ := GenerateRandomString(16)
	file, err := os.Create("./view/webcontent/www/tmp/" + t + "-" + page.Image.Keywords + "-commonwealth-cocktails.jpg")
	if err != nil {
		log.Errorln(err)
		page.RenderPageTemplate(w, r, "404")
		return
	}
	// Use io.Copy to just dump the response body to the file. This supports huge files
	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Errorln(err)
		page.RenderPageTemplate(w, r, "404")
		return
	}
	file.Close()
	log.Infoln("Success!")
	log.Infoln("Image Test")

	im, err1 := gg.LoadImage("./view/webcontent/www/tmp/" + t + "-" + page.Image.Keywords + "-commonwealth-cocktails.jpg")
	if err1 != nil {
		log.Errorln(err1)
		page.RenderPageTemplate(w, r, "404")
		return
	}
	dc := gg.NewContext(im.Bounds().Size().X, im.Bounds().Size().Y)
	font_size := 48
	if err1 := dc.LoadFontFace("./view/webcontent/www/fonts/PoiretOne-Regular.ttf", float64(font_size)); err1 != nil {
		panic(err1)
	}
	lines := len(dc.WordWrap(page.Image.Text, float64(dc.Width())))
	log.Infoln(lines)
	dc.DrawImage(im, 0, 0)
	dc.SavePNG("./view/webcontent/www/tmp/" + t + "-" + page.Image.Keywords + "-commonwealth-cocktails.png")
	dc = gg.NewContext(im.Bounds().Size().X, im.Bounds().Size().Y+(lines*font_size)+36)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	if err1 := dc.LoadFontFace("./view/webcontent/www/fonts/PoiretOne-Regular.ttf", float64(font_size)); err1 != nil {
		panic(err1)
	}
	dc.DrawImage(im, 0, 0)
	dc.DrawStringWrapped(page.Image.Text, float64(dc.Width())/2, float64(dc.Height()-36)-(2.0*float64(lines)*float64(font_size)), 0.5, -1.0, float64(dc.Width()), 1.2, gg.AlignCenter)
	if err1 = dc.LoadFontFace("./view/webcontent/www/fonts/PoiretOne-Regular.ttf", float64(18)); err1 != nil {
		panic(err1)
	}
	dc.DrawStringWrapped("CommonwealthCocktails.com", float64(dc.Width())/2, float64(dc.Height()-48), 0.5, -1.0, float64(dc.Width()), 0.0, gg.AlignCenter)
	dc.Clip()
	dc.SavePNG("./view/webcontent/www/tmp/" + t + "-" + page.Image.Keywords + "-commonwealth-cocktails_labeled.png")

	if page.Image.ImageGen == "test" {
		LoadFileToS3("./view/webcontent/www/tmp/"+t+"-"+page.Image.Keywords+"-commonwealth-cocktails.png", "commonwealthcocktailstmp")
		LoadFileToS3("./view/webcontent/www/tmp/"+t+"-"+page.Image.Keywords+"-commonwealth-cocktails_labeled.png", "commonwealthcocktailstmp")
		page.Image.ImageSource = "https://s3.amazonaws.com/commonwealthcocktailstmp/" + t + "-" + page.Image.Keywords + "-commonwealth-cocktails.png"
		page.Image.LabeledImageSource = "https://s3.amazonaws.com/commonwealthcocktailstmp/" + t + "-" + page.Image.Keywords + "-commonwealth-cocktails_labeled.png"
	} else if page.Image.ImageGen == "deploy" {
		LoadFileToS3("./view/webcontent/www/tmp/"+t+"-"+page.Image.Keywords+"-commonwealth-cocktails.png", "commonwealthcocktailsimages")
		LoadFileToS3("./view/webcontent/www/tmp/"+t+"-"+page.Image.Keywords+"-commonwealth-cocktails_labeled.png", "commonwealthcocktailsimages")
		page.Image.ImageSource = "https://s3.amazonaws.com/commonwealthcocktailsimages/" + t + "-" + page.Image.Keywords + "-commonwealth-cocktails.png"
		page.Image.LabeledImageSource = "https://s3.amazonaws.com/commonwealthcocktailsimages/" + t + "-" + page.Image.Keywords + "-commonwealth-cocktails_labeled.png"

	}

	err = os.Remove("./view/webcontent/www/tmp/" + t + "-" + page.Image.Keywords + "-commonwealth-cocktails.jpg")
	if err != nil {
		log.Errorln(err)
		page.RenderPageTemplate(w, r, "404")
		return
	}
	err = os.Remove("./view/webcontent/www/tmp/" + t + "-" + page.Image.Keywords + "-commonwealth-cocktails.png")
	if err != nil {
		log.Errorln(err)
		page.RenderPageTemplate(w, r, "404")
		return
	}
	err = os.Remove("./view/webcontent/www/tmp/" + t + "-" + page.Image.Keywords + "-commonwealth-cocktails_labeled.png")
	if err != nil {
		log.Errorln(err)
		page.RenderPageTemplate(w, r, "404")
		return
	}

	page.RenderPageTemplate(w, r, "image")

}

//Image update form page request handler which process the image
//update request.
func ImageUpdateHandler(w http.ResponseWriter, r *http.Request, page *page) {
	if page.Image.ImageUpdate == "cocktail" {
		id := strconv.FormatInt(int64(page.Cocktail.UpdateCocktailImages(page.View)), 10)
		http.Redirect(w, r, "/cocktail?cocktailID="+id, 302)
	} else if page.Image.ImageUpdate == "product" {
		id := strconv.FormatInt(int64(page.Product.UpdateProductImages(page.View)), 10)
		http.Redirect(w, r, "/product?ID="+id, 302)
	} else if page.Image.ImageUpdate == "restart" {
		cba := page.Cocktail.SelectCocktailsByAlphaNums(false, page.View)
		page.CocktailsByAlphaNums = cba
		pbt := page.Product.SelectProductsByTypes(true, true, false, page.View)
		page.ProductsByTypes = pbt
		page.RenderPageTemplate(w, r, "imagemodform")
	}
}

//Parses the form and then validates the image form request and
//populates the Image struct
func ValidateImage(w http.ResponseWriter, r *http.Request, page *page) bool {
	r.ParseForm() // Required if you don't call r.FormValue()
	if len(r.Form["imageText"]) > 0 {
		page.Image.Text = r.Form["imageText"][0]
	}
	if len(r.Form["imageSource"]) > 0 {
		page.Image.ImageSource = r.Form["imageSource"][0]
	}
	if len(r.Form["imageKeywords"]) > 0 {
		page.Image.Keywords = r.Form["imageKeywords"][0]
	}
	if len(r.Form["imageGen"]) > 0 {
		page.Image.ImageGen = r.Form["imageGen"][0]
	}
	if len(r.Form["cocktailID"]) > 0 && strings.TrimSpace(r.Form["cocktailID"][0]) != "" {
		if govalidator.IsInt(r.Form["cocktailID"][0]) {
			page.Cocktail.ID, _ = strconv.Atoi(r.Form["cocktailID"][0])
		} else {
			page.Cocktail.Errors["CocktailID"] = "Please enter a valid cocktail id. "
		}
	}
	if len(r.Form["productID"]) > 0 && strings.TrimSpace(r.Form["productID"][0]) != "" {
		if govalidator.IsInt(r.Form["productID"][0]) {
			page.Product.ID, _ = strconv.Atoi(r.Form["productID"][0])
		} else {
			page.Product.Errors["ProductID"] = "Please enter a valid cocktail id. "
		}
	}
	if len(r.Form["imageUpdate"]) > 0 {
		page.Image.ImageUpdate = r.Form["imageUpdate"][0]
		if page.Image.ImageUpdate == "cocktail" {
			if len(r.Form["cocktailSelect"]) > 0 {
				page.Cocktail.ID, _ = strconv.Atoi(r.Form["cocktailSelect"][0])
				if len(r.Form["imageSource"]) > 0 {
					page.Cocktail.ImagePath, page.Cocktail.Image = filepath.Split(r.Form["imageSource"][0])
					page.Cocktail.ImagePath = strings.TrimSuffix(page.Cocktail.ImagePath, "/")
				}
				if len(r.Form["imageLabeledImageSource"]) > 0 {
					page.Cocktail.LabeledImageLink = r.Form["imageLabeledImageSource"][0]
				}
			}
		} else if page.Image.ImageUpdate == "product" {
			if len(r.Form["productSelect"]) > 0 {
				page.Product.ID, _ = strconv.Atoi(r.Form["productSelect"][0])
				if len(r.Form["imageSource"]) > 0 {
					page.Product.ImagePath, page.Product.Image = filepath.Split(r.Form["imageSource"][0])
					page.Product.ImagePath = strings.TrimSuffix(page.Product.ImagePath, "/")
				}
				if len(r.Form["imageLabeledImageSource"]) > 0 {
					page.Product.LabeledImageLink = r.Form["imageLabeledImageSource"][0]
				}
			}
		}
	}

	return true
}

//Checks the page meta struct that required fields are filled in.
func RequiredImageMod(w http.ResponseWriter, r *http.Request, page *page) bool {
	r.ParseForm() // Required if you don't call r.FormValue()
	missingRequired := false
	return missingRequired
}
