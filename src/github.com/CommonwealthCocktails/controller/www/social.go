// Copyright 2017 Rough Industries LLC. All rights reserved.
//controller/www/social.go: Functions and handlers for interacting with the
//social media websites.
package www

import (
	fb "github.com/huandu/facebook"
	log "github.com/sirupsen/logrus"
	"net/http"
	//"strconv"
)

//delete all the memcache entries
func SocialPostHandler(w http.ResponseWriter, r *http.Request, page *page) {
	log.Infoln("SocialPostHandler")
	page_res, err := fb.Get("/774837096005400", fb.Params{
		"fields":       "access_token",
		"access_token": page.UserSession.User.FBAccessToken,
	})
	//page_res, err := http.Get("https://graph.facebook.com/v2.9/me/accounts?access_token=" + token.AccessToken)
	if err == nil {
		log.Infoln(page_res)
		res, _ := fb.Post("/774837096005400/feed?", fb.Params{
			"link":         page.BaseURL + "/" + page.SocialSource,
			"access_token": page_res["access_token"],
		})
		log.Infoln("here is the facebook results:", res)
	} else {
		log.Errorln(err)
	}
	return
}

//Validates the cocktail form request and populates the Cocktail
//struct
func ValidateSocialPost(w http.ResponseWriter, r *http.Request, page *page) bool {
	r.ParseForm() // Required if you don't call r.FormValue()
	if len(r.Form["URLExt"]) > 0 {
		page.SocialSource = r.Form["URLExt"][0]
		log.Infoln(page.SocialSource)
	} else {
		log.Errorln("Invalid Source")
		return false
	}
	return true
}
