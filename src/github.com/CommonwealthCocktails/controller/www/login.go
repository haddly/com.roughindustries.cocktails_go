// Copyright 2017 Rough Industries LLC. All rights reserved.
//controller/www/login.go: Functions and handlers for dealing with logins.  This
//includes standard page login and OAuth
package www

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"github.com/CommonwealthCocktails/connectors"
	"github.com/CommonwealthCocktails/model"
	"github.com/asaskevich/govalidator"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/golang/glog"
	"github.com/microcosm-cc/bluemonday"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/google"
	"html"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

//Variables for use within the login package
var (
	//SET THESE LINES AND ADD #gitignore to the end of the line as a comment to ignore your info
	//Google OAuth variables
	//googleOauthConfig = &oauth2.Config{
	//RedirectURL:  ??,
	//ClientID:     ??,
	//ClientSecret: ??,
	//Scopes: []string{"https://www.googleapis.com/auth/userinfo.profile",
	//	"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/plus.me"},
	//Endpoint: google.Endpoint,
	//}
	//Facebook OAuth variables
	//facebookOauthConfig = &oauth2.Config{
	//ClientID: ??,
	//ClientSecret: ??,
	//RedirectURL: ??,
	//Scopes:   []string{"public_profile", "email"},
	//Endpoint: facebook.Endpoint,
	//}

	// Some random string, random for each request
	// this way could create a memory leak sense I don't clear out the map ever, just a heads up
	oauthStateString = make(map[string]bool)

	//Default user is the user you can get into the system with at all times
	//allowDefault = ??
	//defaultUser = ??
	//defaultPassword = ??

	//sitekey = "{Your site key here}"
	//re      = recaptcha.R{
	//	Secret: "{Your secret here}",
	//}

)

//Login page handler which displays the standard login page.
func loginIndexHandler(w http.ResponseWriter, r *http.Request, page *page) {
	page.RenderPageTemplate(w, r, "loginindex")
}

//Login page request handler which process the standard login request.  This
//will after verifying the user and password create a user session
func loginHandler(w http.ResponseWriter, r *http.Request, page *page) {

	//check the recaptcha to make sure we don't have bot or something
	isValid := re.Verify(*r)
	if !isValid {
		glog.Infoln("Invalid! These errors ocurred: %v", re.LastError())
		page.Errors["loginErrors"] = "You failed the reCAPTCHA test.  You might be bot or something"
		page.RenderPageTemplate(w, r, "/loginindex")
		return
	}

	//this is in case you need to perform DB actions before the DB is setup
	//otherwise you wouldn't have an users
	if allowDefault && page.User.Username == defaultUser {
		//if page.User.Password == defaultPassword {
		if bcrypt.CompareHashAndPassword([]byte(defaultPassword), []byte(page.User.Password)) == nil {
			page.UserSession.IsDefaultUser = true
			page.UserSession.User.Username = defaultUser
			page.UserSession.LoginTime = time.Now()
			page.UserSession.LastSeenTime = time.Now()
			ClearSession(w, r)
			SetSession(w, r, &page.UserSession, true)
			http.Redirect(w, r, "/", 302)
			return
		} else {
			page.UserSession.IsDefaultUser = false
		}
	} else {
		page.UserSession.IsDefaultUser = false
	}
	//Confirm the username is in DB and password after getting user from DB
	usr := page.User.SelectUserForLogin(false)
	glog.Infoln(usr.Password)
	glog.Infoln(page.User.Password)
	if bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(page.User.Password)) == nil {
		glog.Infoln("Passed password")
		if len(usr.Username) > 0 {
			page.UserSession.User = *usr
			page.UserSession.LoginTime = time.Now()
			page.UserSession.LastSeenTime = time.Now()
			SetSession(w, r, &page.UserSession, true)
			http.Redirect(w, r, "/", 302)
			return
		} else {
			glog.Infoln("Bad username or password: " + page.User.Username)
			page.Errors["loginErrors"] = "Bad Username and/or Password!"
			page.RenderPageTemplate(w, r, "/loginindex")
			return
		}
	} else {
		glog.Infoln("Bad username or password: " + page.User.Username)
		page.Errors["loginErrors"] = "Bad Username and/or Password!"
		page.RenderPageTemplate(w, r, "/loginindex")
		return
	}

}

//Logout page request handler which process the standard logout request.  This
//will close the user's session
func logoutHandler(w http.ResponseWriter, r *http.Request, page *page) {
	ClearSession(w, r)
	http.Redirect(w, r, "/", 302)
}

//Initial request from the website that then submits the request to google
func handleGoogleLogin(w http.ResponseWriter, r *http.Request, page *page) {
	str := randSeq(64)
	mc, _ := connectors.GetMC()
	if mc != nil {
		buf := new(bytes.Buffer)
		enc := gob.NewEncoder(buf)
		//change to a date time to see if the seq is out of date so we don't accept it
		enc.Encode(time.Now())
		mc.Set(&memcache.Item{Key: str, Value: buf.Bytes()})
	} else {
		glog.Infoln("Bad memcache handleGoogleLogin")
		//Try the database here
		// if db connection is good {
		// } else {
		// 	panic
		// }
	}
	url := googleOauthConfig.AuthCodeURL(str)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

//Handler for the response to the Google OAuth request from handleGoogleLogin
func handleGoogleCallback(w http.ResponseWriter, r *http.Request, page *page) {
	state := r.FormValue("state")
	//change to a date time to see if the seq is out of date so we don't accept it
	var timeForState time.Time
	mc, _ := connectors.GetMC()
	if mc != nil {
		item := new(memcache.Item)
		item, _ = mc.Get(state)
		if item != nil {
			mc.Delete(state)
			if len(item.Value) > 0 {
				read := bytes.NewReader(item.Value)
				dec := gob.NewDecoder(read)
				dec.Decode(&timeForState)
				if time.Since(timeForState).Seconds() > 30 {
					glog.Infoln("invalid oauth state")
					http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
					return
				}
			} else {
				glog.Infoln("invalid oauth state")
				return
			}
		} else {
			glog.Infoln("invalid oauth state")
			return
		}
	} else {
		glog.Infoln("Bad memcache handleGoogleCallback")
		//Try the database here
		// if db connection is good {
		// } else {
		// 	panic
		// }
	}

	code := r.FormValue("code")
	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		glog.Infoln("Code exchange failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	//glog.Infoln("Content: %s\n", contents)
	s := string(contents[:])
	//Get the email address
	email := strings.Replace(strings.Replace(strings.Split(strings.Split(s, ",")[1], ":")[1], "\"", "", -1), " ", "", -1)
	glog.Infoln(email)
	var user model.User
	user.Email = email
	usr := user.SelectUserForLogin(true)
	us := new(model.UserSession)
	us.LoginTime = time.Now()
	us.LastSeenTime = time.Now()
	us.User = *usr
	SetSession(w, r, us, true)
	http.Redirect(w, r, "/", 302)
}

//Initial request from the website that then submits the request to facebook
func handleFacebookLogin(w http.ResponseWriter, r *http.Request, page *page) {
	str := randSeq(64)
	mc, _ := connectors.GetMC()
	if mc != nil {
		buf := new(bytes.Buffer)
		enc := gob.NewEncoder(buf)
		//change to a date time to see if the seq is out of date so we don't accept it
		enc.Encode(time.Now())
		mc.Set(&memcache.Item{Key: str, Value: buf.Bytes()})
	} else {
		glog.Infoln("Bad memcache handleGoogleLogin")
		//Try the database here
		// if db connection is good {
		// } else {
		// 	panic
		// }
	}
	oauthStateString[str] = true
	url := facebookOauthConfig.AuthCodeURL(str)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

//Handler for the response to the Facebook OAuth request from handleFacebookLogin
func handleFacebookCallback(w http.ResponseWriter, r *http.Request, page *page) {
	state := r.FormValue("state")
	//change to a date time to see if the seq is out of date so we don't accept it
	var timeForState time.Time
	mc, _ := connectors.GetMC()
	if mc != nil {
		item := new(memcache.Item)
		item, _ = mc.Get(state)
		if item != nil {
			mc.Delete(state)
			if len(item.Value) > 0 {
				read := bytes.NewReader(item.Value)
				dec := gob.NewDecoder(read)
				dec.Decode(&timeForState)
				if time.Since(timeForState).Seconds() > 30 {
					glog.Infoln("invalid oauth state")
					http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
					return
				}
			} else {
				glog.Infoln("invalid oauth state")
				return
			}
		} else {
			glog.Infoln("invalid oauth state")
			return
		}
	} else {
		glog.Infoln("Bad memcache handleFacebookCallback")
		//Try the database here
		// if db connection is good {
		// } else {
		// 	panic
		// }
	}

	code := r.FormValue("code")
	//_, err := facebookOauthConfig.Exchange(oauth2.NoContext, code)
	token, err := facebookOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		glog.Infoln("Code exchange failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	response, err := http.Get("https://graph.facebook.com/v2.9/me?fields=id%2Cemail%2Cname&access_token=" + token.AccessToken)

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	var dat map[string]interface{}
	json.Unmarshal([]byte(contents), &dat)
	glog.Infoln(dat)
	// //Get the email address
	email := dat["email"]
	var user model.User
	user.Email = email.(string)
	usr := user.SelectUserForLogin(true)
	us := new(model.UserSession)
	us.LoginTime = time.Now()
	us.LastSeenTime = time.Now()
	us.User = *usr
	SetSession(w, r, us, true)
	http.Redirect(w, r, "/", 302)
}

//Parses the form and then validates the login form request
//and populates the user struct
func ValidateLogin(w http.ResponseWriter, r *http.Request, page *page) bool {
	page.User.Errors = make(map[string]string)
	r.ParseForm() // Required if you don't call r.FormValue()
	pSP := bluemonday.StrictPolicy()
	if len(r.Form["name"]) > 0 && strings.TrimSpace(r.Form["name"][0]) != "" && govalidator.IsPrintableASCII(r.Form["name"][0]) {
		page.User.Username = html.EscapeString(pSP.Sanitize(r.Form["name"][0]))
	} else {
		page.User.Errors["Username"] = "Please enter a valid username"
	}
	if len(r.Form["password"]) > 0 && strings.TrimSpace(r.Form["password"][0]) != "" && govalidator.IsPrintableASCII(r.Form["password"][0]) {
		page.User.Password = html.EscapeString(pSP.Sanitize(r.Form["password"][0]))
	} else {
		page.User.Errors["Password"] = "Please enter a valid password"
	}
	if len(page.User.Errors) > 0 {
		page.Errors["loginErrors"] = "Invalid Username and/or Password"
	}
	return len(page.User.Errors) == 0
}
