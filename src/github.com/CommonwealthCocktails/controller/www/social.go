// Copyright 2017 Rough Industries LLC. All rights reserved.
//controller/www/social.go: Functions and handlers for interacting with the
//social media websites.
package www

import (
	"github.com/golang/glog"
	fb "github.com/huandu/facebook"
	"net/http"
	//"strconv"
)

//delete all the memcache entries
func SocialPostHandler(w http.ResponseWriter, r *http.Request, page *page) {
	page_res, err := fb.Get("/774837096005400", fb.Params{
		"fields":       "access_token",
		"access_token": page.UserSession.User.FBAccessToken,
	})
	//page_res, err := http.Get("https://graph.facebook.com/v2.9/me/accounts?access_token=" + token.AccessToken)
	if err == nil {
		glog.Infoln(page_res)
		res, _ := fb.Post("/774837096005400/feed?", fb.Params{
			"link":         page.BaseURL + "/" + page.SocialSource,
			"access_token": page_res["access_token"],
		})
		glog.Infoln("here is the facebook results:", res)
	}
	return
}

//Validates the cocktail form request and populates the Cocktail
//struct
func ValidateSocialPost(w http.ResponseWriter, r *http.Request, page *page) bool {
	r.ParseForm() // Required if you don't call r.FormValue()
	if len(r.Form["URLExt"]) > 0 {
		page.SocialSource = r.Form["URLExt"][0]
		glog.Infoln(page.SocialSource)
	} else {
		glog.Errorln("Invalid Source")
		return false
	}
	return true
}
