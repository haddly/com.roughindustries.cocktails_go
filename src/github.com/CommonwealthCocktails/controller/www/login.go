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
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/google"
	"html"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/smtp"
	"strings"
	"time"
)

//Variables for use within the login package
var (
	//OAuth
	googleOauthConfig   *oauth2.Config
	facebookOauthConfig *oauth2.Config

	// Some random string, random for each request
	// this way could create a memory leak sense I don't clear out the map ever, just a heads up
	oauthStateString = make(map[string]bool)

	//Default user is the user you can get into the system with at all times
	allowDefault    bool
	defaultUser     string
	defaultPassword string

	//reCAPTCHA
	sitekey    string
	re         reCAPTCHA
	sitekeyInv string
	reInv      reCAPTCHA
)

func LoginInit() {
	glog.Infoln("Login Init")
	//default user
	allowDefault = viper.GetBool("allowDefault")
	defaultUser = viper.GetString("defaultUser")
	//hash is = password
	defaultPassword = viper.GetString("defaultPassword")

	//reCAPTCHA
	sitekey = viper.GetString("reCAPTCHASiteKey")
	re = reCAPTCHA{
		Secret: viper.GetString("reCAPTCHASecret"),
	}
	sitekeyInv = viper.GetString("reCAPTCHASiteKeyInv")
	reInv = reCAPTCHA{
		Secret: viper.GetString("reCAPTCHASecretInv"),
	}

	googleOauthConfig = &oauth2.Config{
		ClientID:     viper.GetString("googleOauthConfigClientID"),
		ClientSecret: viper.GetString("googleOauthConfigClientSecret"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	facebookOauthConfig = &oauth2.Config{
		ClientID:     viper.GetString("facebookOauthConfigClientID"),
		ClientSecret: viper.GetString("facebookOauthConfigClientSecret"),
		Scopes:       []string{"public_profile", "email", "pages_show_list", "manage_pages", "publish_pages"},
		Endpoint:     facebook.Endpoint,
	}
}

//Login page handler which displays the standard login page.
func loginIndexHandler(w http.ResponseWriter, r *http.Request, page *page) {
	page.RenderPageTemplate(w, r, "loginindex")
}

//Register page handler which displays the standard register page.
func registerHandler(w http.ResponseWriter, r *http.Request, page *page) {
	page.RenderPageTemplate(w, r, "registerForm")
}

//Forgot Password page handler which displays the standard forgot passwd form page.
func forgotPasswdHandler(w http.ResponseWriter, r *http.Request, page *page) {
	page.RenderPageTemplate(w, r, "forgotpasswdform")
}

//Reset Password page handler which displays the standard reset passwd form page.
func resetPasswdHandler(w http.ResponseWriter, r *http.Request, page *page) {
	page.RenderPageTemplate(w, r, "passwdresetform")
}

//Login page request handler which process the standard login request.  This
//will after verifying the user and password create a user session
func loginHandler(w http.ResponseWriter, r *http.Request, page *page) {

	//check the recaptcha to make sure we don't have bot or something
	isValid := re.Verify(*r)
	if !isValid {
		glog.Infoln("Invalid! These errors ocurred: %v", re.LastError())
		page.Errors["loginErrors"] = "You failed the reCAPTCHA test.  You might be bot or something"
		page.RenderPageTemplate(w, r, "loginindex")
		return
	}

	//this is in case you need to perform DB actions before the DB is setup
	//otherwise you wouldn't have an users
	glog.Infoln(allowDefault)
	if allowDefault && page.User.Username == defaultUser {
		//if page.User.Password == defaultPassword {
		if bcrypt.CompareHashAndPassword([]byte(defaultPassword), []byte(page.User.Password)) == nil {
			page.UserSession.IsDefaultUser = true
			page.UserSession.User.Username = defaultUser
			page.UserSession.LoginTime = time.Now()
			page.UserSession.LastSeenTime = time.Now()
			page.UserSession.User.VerificationComplete = true
			page.UserSession.User.Role = 1
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
			if !usr.VerificationComplete {
				glog.Infoln("User has not finished verification process: " + page.User.Username)
				page.Errors["loginErrors"] = "Unverified Account! Please either signup or check your email to verify it if you have already registered."
				page.RenderPageTemplate(w, r, "/loginindex")
				return
			}
			page.UserSession.User = *usr
			glog.Infoln(page.UserSession.User.VerificationComplete)
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

//Register page request handler which process the standard register request.
func registerFormHandler(w http.ResponseWriter, r *http.Request, page *page) {
	pSP := bluemonday.StrictPolicy()
	//check the recaptcha to make sure we don't have bot or something
	isValid := reInv.Verify(*r)
	if !isValid {
		glog.Infoln("Invalid! These errors ocurred: %v", re.LastError())
		page.Errors["loginErrors"] = "You failed the reCAPTCHA test.  You might be bot or something"
		page.RenderPageTemplate(w, r, "registerForm")
		return
	}

	//I am using gmail smtp.  If you have 2 step authentication get an app
	//password that corresponds to the from email account you use.
	//from := ""
	//pass := ""

	//check username and email against database, if it exists return to
	//registration page with an error else add user to the database
	users := page.User.SelectUsersByIdORUsernameOREmail()
	if len(users) > 0 {
		for _, user := range users {
			if user.Username == page.User.Username {
				page.User.Errors["Username"] = "Username already exists"
			}
			if user.Email == page.User.Email {
				page.User.Errors["Email"] = "Email already exists"
			}
		}
		registerHandler(w, r, page)
		return
	} else {
		//update the password so that it is salted and hashed with bcrypt
		glog.Infoln("Password:", page.User.Password)
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(page.User.Password), bcrypt.DefaultCost)
		page.User.Password = string(hashedPassword)
		glog.Infoln("Hash:    ", hashedPassword)
		//set the verification code and time created
		code, _ := GenerateRandomString(64)
		page.User.VerificationCode = html.EscapeString(pSP.Sanitize(code))
		page.User.VerificationInitTime = time.Now()
		page.User.VerificationComplete = false
		//add user to database
		page.User.InsertUser()
	}

	//load up the verification email
	t, err := template.ParseFiles("./view/webcontent/www/templates/verificationemail.html")
	if err != nil {
		glog.Errorln(err)
		Error404(w, err)
		return
	}
	buf := new(bytes.Buffer)
	err = t.ExecuteTemplate(buf, "base", page)
	body := buf.String()
	if err != nil {
		glog.Errorln(err)
		Error404(w, err)
		return
	}

	//Email setup.  This is designed ot send an html enabled email
	msg := "From: " + from + "\n" +
		"To: " + page.User.Email + "\n" +
		"MIME-Version: 1.0" + "\r\n" +
		"Content-type: text/html" + "\r\n" +
		"Subject: Registration for Commonwealth Cocktails\n\n" +
		body

	//gmail smtp
	err = smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{page.User.Email}, []byte(msg))

	if err != nil {
		glog.Infoln("smtp error: %s", err)
		Error404(w, err)
		return
	}
	page.RenderPageTemplate(w, r, "verifyregisternotice")
}

//Register Verification page request handler which process the standard
//register verificationrequest.
func verifyRegisterHandler(w http.ResponseWriter, r *http.Request, page *page) {
	pSP := bluemonday.StrictPolicy()
	if len(page.User.Errors) > 0 {
		page.User.VerificationComplete = false
		glog.Errorln("Someone tried to verify with a bad link, potential attack!")
	} else {
		//confirm the user's email and code
		usr := page.User.SelectUserForVerification()
		if usr.ID != 0 {
			//check that the time sense the user requested addition hasn't been
			//more the 24 hours
			delta := time.Now().Sub(usr.VerificationInitTime)
			if delta.Hours() > 24 {
				page.User.Errors["User"] = "You have waited to long for verification."
				glog.Errorln("Someone tried to verify with a verification that is to old, potential attack!")
			} else {
				//update the verification code and time of init and complete
				page.User = *usr
				code, _ := GenerateRandomString(64)
				page.User.VerificationCode = html.EscapeString(pSP.Sanitize(code))
				page.User.VerificationInitTime = time.Time{}
				page.User.VerificationComplete = true
				//update the user in the database
				page.User.UpdateUser()
			}
		} else {
			page.User.Errors["User"] = "You have provided bad information for verification."
			glog.Errorln("Someone tried to verify with a bad code and email, potential attack!")
		}
	}
	page.RenderPageTemplate(w, r, "verifyregisternotice")
}

//Forgot password page request handler which process the standard forgot password request.
func forgotPasswdFormHandler(w http.ResponseWriter, r *http.Request, page *page) {
	pSP := bluemonday.StrictPolicy()
	//check the recaptcha to make sure we don't have bot or something
	isValid := reInv.Verify(*r)
	if !isValid {
		glog.Infoln("Invalid! These errors ocurred: %v", re.LastError())
		page.Errors["loginErrors"] = "You failed the reCAPTCHA test.  You might be bot or something"
		page.RenderPageTemplate(w, r, "forgotpasswdform")
		return
	}

	//check username and email against database, if it exists return to
	//registration page with an error else add user to the database
	users := page.User.SelectUsersByIdORUsernameOREmail()
	if len(users) == 0 {
		page.User.Errors["Email"] = "No account with that email found."
		forgotPasswdHandler(w, r, page)
		return
	} else if len(users) == 1 {
		page.User = users[0]
		tempPasswd, _ := GenerateRandomString(15)
		glog.Infoln("Passwd: ", tempPasswd)
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(tempPasswd), bcrypt.DefaultCost)
		page.User.Password = string(hashedPassword)
		glog.Infoln("Hash:    ", hashedPassword)
		//update the verification code for a password reset
		code, _ := GenerateRandomString(64)
		page.User.VerificationCode = html.EscapeString(pSP.Sanitize(code))
		page.User.VerificationInitTime = time.Now()
		//update the user in the db
		page.User.UpdateUser()
		page.User.Password = tempPasswd
		//load up the verification email
		t, err := template.ParseFiles("./view/webcontent/www/templates/forgotpasswdemail.html")
		if err != nil {
			glog.Errorln(err)
			Error404(w, err)
			return
		}
		buf := new(bytes.Buffer)
		err = t.ExecuteTemplate(buf, "base", page)
		body := buf.String()
		if err != nil {
			glog.Errorln(err)
			Error404(w, err)
			return
		}

		//I am using gmail smtp.  If you have 2 step authentication get an app
		//password that corresponds to the from email account you use.
		//from := ""
		//pass := ""

		//Email setup.  This is designed ot send an html enabled email
		msg := "From: " + from + "\n" +
			"To: " + page.User.Email + "\n" +
			"MIME-Version: 1.0" + "\r\n" +
			"Content-type: text/html" + "\r\n" +
			"Subject: Password Reset for Commonwealth Cocktails\n\n" +
			body

		//gmail smtp
		err = smtp.SendMail("smtp.gmail.com:587",
			smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
			from, []string{page.User.Email}, []byte(msg))

		if err != nil {
			glog.Infoln("smtp error: %s", err)
			Error404(w, err)
			return
		}
	} else {
		glog.Errorln("We have more then one account with the same email address.")
		Error404(w, nil)
		return
	}
	page.RenderPageTemplate(w, r, "forgotpasswdnotice")
}

//Reset password page request handler which process the standard reset password request.
func resetPasswdFormHandler(w http.ResponseWriter, r *http.Request, page *page) {
	pSP := bluemonday.StrictPolicy()
	//check the recaptcha to make sure we don't have bot or something
	isValid := reInv.Verify(*r)
	if !isValid {
		glog.Infoln("Invalid! These errors ocurred: %v", re.LastError())
		page.Errors["loginErrors"] = "You failed the reCAPTCHA test.  You might be bot or something"
		page.RenderPageTemplate(w, r, "index")
		return
	}

	//check password and code against database
	users := page.User.SelectUsersByIdORUsernameOREmail()
	if len(users) == 1 {
		//check that the time sense the user requested addition hasn't been
		//more the 24 hours
		delta := time.Now().Sub(users[0].VerificationInitTime)
		if delta.Hours() > 24 {
			page.Errors["passwdResetErrors"] = "You have waited to long to reset your password.  Please request another reset."
			glog.Errorln("Someone tried to reset a password that is to old, potential attack!")
			page.RenderPageTemplate(w, r, "passwdresetform")
			return
		} else {
			if (bcrypt.CompareHashAndPassword([]byte(users[0].Password), []byte(page.User.Password)) == nil) && (page.User.VerificationCode == users[0].VerificationCode) {
				//update the password so that it is salted and hashed with bcrypt
				glog.Infoln("Password:", page.User.NewPassword)
				hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(page.User.NewPassword), bcrypt.DefaultCost)
				users[0].Password = string(hashedPassword)
				glog.Infoln("Hash:    ", hashedPassword)
				//set the verification code and time created
				code, _ := GenerateRandomString(64)
				users[0].VerificationCode = html.EscapeString(pSP.Sanitize(code))
				users[0].VerificationInitTime = time.Time{}
				page.User = users[0]
				//update the user in the db
				page.User.UpdateUser()
			} else {
				glog.Errorln("Someone is trying to reset passwords with bad data.  This could be an attack.")
				Error404(w, nil)
				return
			}
		}
	} else {
		glog.Errorln("We have zero or more then one accounts with this email address.")
		Error404(w, nil)
		return
	}
	page.RenderPageTemplate(w, r, "loginindex")
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
	googleOauthConfig.RedirectURL = page.BaseURL + "/GoogleCallback"
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
	if !usr.VerificationComplete {
		glog.Infoln("User has not finished verification process: " + page.User.Username)
		page.Errors["loginErrors"] = "Unverified Account! Please either signup or check your email to verify it if you have already registered."
		page.RenderPageTemplate(w, r, "/loginindex")
		return
	}
	us := new(model.UserSession)
	us.LoginTime = time.Now()
	us.LastSeenTime = time.Now()
	usr.GoogleAccessToken = token.AccessToken
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
	facebookOauthConfig.RedirectURL = page.BaseURL + "/FacebookCallback"
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
	if !usr.VerificationComplete {
		glog.Infoln("User has not finished verification process: " + page.User.Username)
		page.Errors["loginErrors"] = "Unverified Account! Please either signup or check your email to verify it if you have already registered."
		page.RenderPageTemplate(w, r, "/loginindex")
		return
	}
	us := new(model.UserSession)
	us.LoginTime = time.Now()
	us.LastSeenTime = time.Now()
	usr.FBAccessToken = token.AccessToken
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

//Parses the form and then validates the register form request
//and populates the user struct
func ValidateRegister(w http.ResponseWriter, r *http.Request, page *page) bool {
	page.User.Errors = make(map[string]string)
	r.ParseForm() // Required if you don't call r.FormValue()
	pSP := bluemonday.StrictPolicy()
	if len(r.Form["firstname"]) > 0 && govalidator.IsPrintableASCII(r.Form["firstname"][0]) {
		page.User.FirstName = html.EscapeString(pSP.Sanitize(r.Form["firstname"][0]))
	} else {
		page.User.Errors["FirstName"] = "Please enter a valid first name"
	}
	if len(r.Form["lastname"]) > 0 && govalidator.IsPrintableASCII(r.Form["lastname"][0]) {
		page.User.LastName = html.EscapeString(pSP.Sanitize(r.Form["lastname"][0]))
	} else {
		page.User.Errors["LastName"] = "Please enter a valid last name"
	}
	if len(r.Form["username"]) > 0 && govalidator.IsPrintableASCII(r.Form["username"][0]) {
		page.User.Username = html.EscapeString(pSP.Sanitize(r.Form["username"][0]))
	} else {
		page.User.Errors["Username"] = "Please enter a valid username"
	}
	if len(r.Form["email"]) > 0 && govalidator.IsEmail(r.Form["email"][0]) {
		page.User.Email = html.EscapeString(pSP.Sanitize(r.Form["email"][0]))
	} else {
		page.User.Errors["Email"] = "Please enter a valid email"
	}
	if len(r.Form["password"]) > 0 && govalidator.IsPrintableASCII(r.Form["password"][0]) {
		if len(r.Form["passwordconfirm"]) > 0 && govalidator.IsPrintableASCII(r.Form["passwordconfirm"][0]) {
			if r.Form["passwordconfirm"][0] == r.Form["password"][0] {
				if len(r.Form["password"][0]) >= 12 {
					page.User.Password = html.EscapeString(pSP.Sanitize(r.Form["password"][0]))
				} else {
					page.User.Errors["Password"] = "Your password must be at least 12 characters."
				}
			} else {
				page.User.Errors["Password"] = "Your password and confirmation password do not match."
			}
		} else {
			page.User.Errors["Password"] = "Please enter a valid password confirmation"
		}
	} else {
		page.User.Errors["Password"] = "Please enter a valid password"
	}
	if len(page.User.Errors) > 0 {
		page.Errors["loginErrors"] = "Invalid Username and/or Password"
	}
	return len(page.User.Errors) == 0
}

//Checks the page meta struct that required fields are filled in.
func RequiredRegister(w http.ResponseWriter, r *http.Request, page *page) bool {
	page.User.Errors = make(map[string]string)
	r.ParseForm() // Required if you don't call r.FormValue()
	missingRequired := false
	if r.Form["firstname"] == nil || len(r.Form["firstname"]) == 0 || strings.TrimSpace(r.Form["firstname"][0]) == "" {
		page.User.Errors["FirstName"] = "First name is required"
		missingRequired = true
	}
	if r.Form["lastname"] == nil || len(r.Form["lastname"]) == 0 || strings.TrimSpace(r.Form["lastname"][0]) == "" {
		page.User.Errors["LastName"] = "Last name is required"
		missingRequired = true
	}
	if r.Form["username"] == nil || len(r.Form["username"]) == 0 || strings.TrimSpace(r.Form["username"][0]) == "" {
		page.User.Errors["Username"] = "Username is required"
		missingRequired = true
	}
	if r.Form["email"] == nil || len(r.Form["email"]) == 0 || strings.TrimSpace(r.Form["email"][0]) == "" {
		page.User.Errors["Email"] = "Email is required"
		missingRequired = true
	}
	if r.Form["password"] == nil || len(r.Form["password"]) == 0 || strings.TrimSpace(r.Form["password"][0]) == "" {
		page.User.Errors["Password"] = "Password is required"
		missingRequired = true
	}
	if r.Form["passwordconfirm"] == nil || len(r.Form["passwordconfirm"]) == 0 || strings.TrimSpace(r.Form["passwordconfirm"][0]) == "" {
		page.User.Errors["Password"] = "Password Confirmation is required"
		missingRequired = true
	}
	return missingRequired
}

//Parses the form and then validates the register verification form request
//and populates the user struct
func ValidateEmailCode(w http.ResponseWriter, r *http.Request, page *page) bool {
	page.User.Errors = make(map[string]string)
	r.ParseForm() // Required if you don't call r.FormValue()
	pSP := bluemonday.StrictPolicy()
	if len(r.Form["code"]) > 0 && govalidator.IsPrintableASCII(r.Form["code"][0]) {
		page.User.VerificationCode = html.EscapeString(pSP.Sanitize(r.Form["code"][0]))
	} else {
		page.User.Errors["Code"] = "Please enter a valid verification code"
	}
	if len(r.Form["email"]) > 0 && govalidator.IsEmail(r.Form["email"][0]) {
		page.User.Email = html.EscapeString(pSP.Sanitize(r.Form["email"][0]))
	} else {
		page.User.Errors["Email"] = "Please enter a valid email address"
	}
	return len(page.User.Errors) == 0
}

//Checks the page meta struct that required fields are filled in.
func RequiredEmailCode(w http.ResponseWriter, r *http.Request, page *page) bool {
	page.User.Errors = make(map[string]string)
	r.ParseForm() // Required if you don't call r.FormValue()
	missingRequired := false
	if r.Form["code"] == nil || len(r.Form["code"]) == 0 || strings.TrimSpace(r.Form["code"][0]) == "" {
		page.User.Errors["Code"] = "Verification Code is required"
		missingRequired = true
	}
	if r.Form["email"] == nil || len(r.Form["email"]) == 0 || strings.TrimSpace(r.Form["email"][0]) == "" {
		page.User.Errors["Email"] = "Email is required"
		missingRequired = true
	}
	return missingRequired
}

//Parses the form and then validates the register verification form request
//and populates the user struct
func ValidateForgotPasswd(w http.ResponseWriter, r *http.Request, page *page) bool {
	page.User.Errors = make(map[string]string)
	r.ParseForm() // Required if you don't call r.FormValue()
	pSP := bluemonday.StrictPolicy()
	if len(r.Form["email"]) > 0 && govalidator.IsEmail(r.Form["email"][0]) {
		page.User.Email = html.EscapeString(pSP.Sanitize(r.Form["email"][0]))
	} else {
		page.User.Errors["Email"] = "Please enter a valid email address"
	}
	return len(page.User.Errors) == 0
}

//Checks the page meta struct that required fields are filled in.
func RequiredForgotPasswd(w http.ResponseWriter, r *http.Request, page *page) bool {
	page.User.Errors = make(map[string]string)
	r.ParseForm() // Required if you don't call r.FormValue()
	missingRequired := false
	if r.Form["email"] == nil || len(r.Form["email"]) == 0 || strings.TrimSpace(r.Form["email"][0]) == "" {
		page.User.Errors["Email"] = "Email is required"
		missingRequired = true
	}
	return missingRequired
}

//Parses the form and then validates the password reset form request
//and populates the user struct
func ValidateResetPasswd(w http.ResponseWriter, r *http.Request, page *page) bool {
	page.User.Errors = make(map[string]string)
	r.ParseForm() // Required if you don't call r.FormValue()
	pSP := bluemonday.StrictPolicy()
	if len(r.Form["code"]) > 0 && govalidator.IsPrintableASCII(r.Form["code"][0]) {
		page.User.VerificationCode = html.EscapeString(pSP.Sanitize(r.Form["code"][0]))
	} else {
		page.User.Errors["Code"] = "Please enter a valid verification code"
	}
	if len(r.Form["email"]) > 0 && govalidator.IsEmail(r.Form["email"][0]) {
		page.User.Email = html.EscapeString(pSP.Sanitize(r.Form["email"][0]))
	} else {
		page.User.Errors["Email"] = "Please enter a valid email address"
	}
	if len(r.Form["currentpassword"]) > 0 && govalidator.IsPrintableASCII(r.Form["currentpassword"][0]) {
		page.User.Password = html.EscapeString(pSP.Sanitize(r.Form["currentpassword"][0]))
	} else {
		page.User.Errors["CurrentPassword"] = "Please enter a valid password"
	}
	if len(r.Form["password"]) > 0 && govalidator.IsPrintableASCII(r.Form["password"][0]) {
		if len(r.Form["passwordconfirm"]) > 0 && govalidator.IsPrintableASCII(r.Form["passwordconfirm"][0]) {
			if r.Form["passwordconfirm"][0] == r.Form["password"][0] {
				if len(r.Form["password"][0]) >= 12 {
					page.User.NewPassword = html.EscapeString(pSP.Sanitize(r.Form["password"][0]))
				} else {
					page.User.Errors["Password"] = "Your password must be at least 12 characters."
				}
			} else {
				page.User.Errors["Password"] = "Your password and confirmation password do not match."
			}
		} else {
			page.User.Errors["Password"] = "Please enter a valid password confirmation"
		}
	} else {
		page.User.Errors["Password"] = "Please enter a valid password"
	}
	if len(page.User.Errors) > 0 {
		page.Errors["passwdResetErrors"] = "Invalid Password"
	}
	return len(page.User.Errors) == 0
}

//Checks the page meta struct that required fields are filled in.
func RequiredResetPasswd(w http.ResponseWriter, r *http.Request, page *page) bool {
	page.User.Errors = make(map[string]string)
	r.ParseForm() // Required if you don't call r.FormValue()
	missingRequired := false
	if r.Form["code"] == nil || len(r.Form["code"]) == 0 || strings.TrimSpace(r.Form["code"][0]) == "" {
		page.User.Errors["Code"] = "Verification Code is required"
		missingRequired = true
	}
	if r.Form["email"] == nil || len(r.Form["email"]) == 0 || strings.TrimSpace(r.Form["email"][0]) == "" {
		page.User.Errors["Email"] = "Email is required"
		missingRequired = true
	}
	if r.Form["currentpassword"] == nil || len(r.Form["currentpassword"]) == 0 || strings.TrimSpace(r.Form["currentpassword"][0]) == "" {
		page.User.Errors["CurrentPassword"] = "Current Password is required"
		missingRequired = true
	}
	if r.Form["password"] == nil || len(r.Form["password"]) == 0 || strings.TrimSpace(r.Form["password"][0]) == "" {
		page.User.Errors["Password"] = "New Password is required"
		missingRequired = true
	}
	if r.Form["passwordconfirm"] == nil || len(r.Form["passwordconfirm"]) == 0 || strings.TrimSpace(r.Form["passwordconfirm"][0]) == "" {
		page.User.Errors["Password"] = "Confirmation Password is required"
		missingRequired = true
	}
	if missingRequired {
		page.Errors["passwdResetErrors"] = "Invalid Data Entered"
	}
	return missingRequired
}
