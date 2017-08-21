//controller/www/login.go: Functions and handlers for dealing with logins.  This
//includes standard page login and OAuth
package www

import (
	"bytes"
	"connectors"
	"encoding/gob"
	"encoding/json"
	"github.com/bradfitz/gomemcache/memcache"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"log"
	"model"
	"net/http"
	"strings"
)

//Login struct for defining methods to
type Login struct {
}

var (
	//SET THESE LINES AND ADD #gitignore to the end of the line as a comment to ignore your info
	//googleOauthConfig = &oauth2.Config{
	//RedirectURL:  ??,
	//ClientID:     ??,
	//ClientSecret: ??,
	//Scopes: []string{"https://www.googleapis.com/auth/userinfo.profile",
	//	"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/plus.me"},
	//Endpoint: google.Endpoint,
	//}
	}
	// Some random string, random for each request
	// this way could create a memory leak sense I don't clear out the map ever, just a heads up
	oauthStateString = make(map[string]bool)

	//defaultUser = ??
	//defaultPassword = ??
)

//Init to setup the http handlers
func (login *Login) Init() {
	log.Println("Init in www/login.go")
	http.HandleFunc("/loginIndex", login.loginIndexHandler)
	http.HandleFunc("/login", login.loginHandler)
	http.HandleFunc("/logout", login.logoutHandler)
	http.HandleFunc("/GoogleLogin", login.handleGoogleLogin)
	http.HandleFunc("/GoogleCallback", login.handleGoogleCallback)
	http.HandleFunc("/FacebookLogin", login.handleFacebookLogin)
	http.HandleFunc("/FacebookCallback", login.handleFacebookCallback)
}

//Login page handler which displays the standard login page.
func (login *Login) loginIndexHandler(w http.ResponseWriter, r *http.Request) {
	// STANDARD HANDLER HEADER START
	// catch all errors and return 404
	defer func() {
		// recover from panic if one occured. Set err to nil otherwise.
		if rec := recover(); rec != nil {
			Error404(w, rec)
		}
	}()
	page := NewPage()
	page.Username, page.Authenticated = GetSession(r)
	// STANDARD HANLDER HEADER END
	page.RenderPageTemplate(w, "loginindex")
}

//Login page request handler which process the standard login request.  This
//will after verifying the user and password create a user session
func (login *Login) loginHandler(w http.ResponseWriter, r *http.Request) {
	// STANDARD HANDLER HEADER START
	// catch all errors and return 404
	defer func() {
		// recover from panic if one occured. Set err to nil otherwise.
		if rec := recover(); rec != nil {
			Error404(w, rec)
		}
	}()
	page := NewPage()
	page.Username, page.Authenticated = GetSession(r)
	// STANDARD HANLDER HEADER END
	if ValidateLogin(&page.User, r) {
		//this is in case you need to perform DB actions before the DB is setup
		//otherwise you wouldn't have an users
		if page.User.Username == defaultUser {
			if page.User.Password == defaultPassword {
				SetSession(w, r, page.User.Username)
				http.Redirect(w, r, "/", 302)
				return
			}
		}
		//Confirm the username and password
		usr := model.SelectUserForLogin(page.User, false)
		if len(usr.Username) > 0 {
			SetSession(w, r, usr.Username)
			http.Redirect(w, r, "/", 302)
			return
		} else {
			log.Println("Bad username or password: " + page.User.Username)
			page.Errors["loginErrors"] = "Bad Username and/or Password!"
			page.RenderPageTemplate(w, "/loginindex")
			return
		}
	} else {
		log.Println("Bad username or password: " + page.User.Username)
		page.Errors["loginErrors"] = "Invalid Username and/or Password"
		page.RenderPageTemplate(w, "/loginindex")
		return
	}

}

//Logout page request handler which process the standard logout request.  This
//will close the user's session
func (login *Login) logoutHandler(w http.ResponseWriter, r *http.Request) {
	ClearSession(w, r)
	http.Redirect(w, r, "/", 302)
}

func (login *Login) handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	// CATCH ONLY HEADER START
	defer func() {
		// recover from panic if one occured. Set err to nil otherwise.
		if rec := recover(); rec != nil {
			Error404(w, rec)
		}
	}()
	// CATCH ONLY HEADER START
	str := randSeq(64)
	//MEMCACHE OAUTH SET
	mc, _ := connectors.GetMC()
	if mc != nil {
		item := new(memcache.Item)
		item, _ = mc.Get("oauthStateString")
		if item != nil {
			if len(item.Value) > 0 {
				read := bytes.NewReader(item.Value)
				dec := gob.NewDecoder(read)
				dec.Decode(&oauthStateString)
			}
		}
		oauthStateString[str] = true
		buf := new(bytes.Buffer)
		enc := gob.NewEncoder(buf)
		enc.Encode(oauthStateString)

		mc.Set(&memcache.Item{Key: "oauthStateString", Value: buf.Bytes()})
	} else {
		log.Println("Bad memcache handleGoogleLogin")
		//Try the database here
		// if db connection is good {
		// } else {
		// 	panic
		// }
	}
	//MEMCACHE OAUTH SET

	url := googleOauthConfig.AuthCodeURL(str)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (login *Login) handleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	// CATCH ONLY HEADER START
	defer func() {
		// recover from panic if one occured. Set err to nil otherwise.
		if rec := recover(); rec != nil {
			Error404(w, rec)
		}
	}()
	// CATCH ONLY HEADER START

	//MEMCACHE OAUTH GET
	mc, _ := connectors.GetMC()
	if mc != nil {
		item := new(memcache.Item)
		item, _ = mc.Get("oauthStateString")
		if item != nil {
			if len(item.Value) > 0 {
				read := bytes.NewReader(item.Value)
				dec := gob.NewDecoder(read)
				dec.Decode(&oauthStateString)
			}
		}
	} else {
		log.Println("Bad memcache handleGoogleCallback")
		//Try the database here
		// if db connection is good {
		// } else {
		// 	panic
		// }
	}
	//MEMCACHE OAUTH GET

	state := r.FormValue("state")
	if !oauthStateString[state] {
		log.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	} else {
		delete(oauthStateString, state)
	}

	//MEMCACHE OAUTH SET
	if mc != nil {
		buf := new(bytes.Buffer)
		enc := gob.NewEncoder(buf)
		enc.Encode(oauthStateString)

		mc.Set(&memcache.Item{Key: "oauthStateString", Value: buf.Bytes()})
	} else {
		//Try the database here
		// if db connection is good {
		// } else {
		// 	panic
		// }
	}
	//MEMCACHE OAUTH SET

	code := r.FormValue("code")
	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Println("Code exchange failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	//log.Printf("Content: %s\n", contents)
	s := string(contents[:])
	//Get the email address
	email := strings.Replace(strings.Replace(strings.Split(strings.Split(s, ",")[1], ":")[1], "\"", "", -1), " ", "", -1)
	log.Println(email)
	var user model.User
	user.Email = email
	usr := model.SelectUserForLogin(user, true)
	SetSession(w, r, usr.Username)
	http.Redirect(w, r, "/", 302)
}

func (login *Login) handleFacebookLogin(w http.ResponseWriter, r *http.Request) {
	// CATCH ONLY HEADER START
	defer func() {
		// recover from panic if one occured. Set err to nil otherwise.
		if rec := recover(); rec != nil {
			Error404(w, rec)
		}
	}()
	// CATCH ONLY HEADER START

	str := randSeq(64)

	//MEMCACHE OAUTH SET
	mc, _ := connectors.GetMC()
	if mc != nil {
		item := new(memcache.Item)
		item, _ = mc.Get("oauthStateString")
		if item != nil {
			if len(item.Value) > 0 {
				read := bytes.NewReader(item.Value)
				dec := gob.NewDecoder(read)
				dec.Decode(&oauthStateString)
			}
		}
		oauthStateString[str] = true
		buf := new(bytes.Buffer)
		enc := gob.NewEncoder(buf)
		enc.Encode(oauthStateString)

		mc.Set(&memcache.Item{Key: "oauthStateString", Value: buf.Bytes()})
	} else {
		log.Println("Bad memcache handleGoogleLogin")
		//Try the database here
		// if db connection is good {
		// } else {
		// 	panic
		// }
	}
	//MEMCACHE OAUTH SET
	oauthStateString[str] = true
	url := facebookOauthConfig.AuthCodeURL(str)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (login *Login) handleFacebookCallback(w http.ResponseWriter, r *http.Request) {
	// CATCH ONLY HEADER START
	defer func() {
		// recover from panic if one occured. Set err to nil otherwise.
		if rec := recover(); rec != nil {
			Error404(w, rec)
		}
	}()
	// CATCH ONLY HEADER START

	//MEMCACHE OAUTH GET
	mc, _ := connectors.GetMC()
	if mc != nil {
		item := new(memcache.Item)
		item, _ = mc.Get("oauthStateString")
		if item != nil {
			if len(item.Value) > 0 {
				read := bytes.NewReader(item.Value)
				dec := gob.NewDecoder(read)
				dec.Decode(&oauthStateString)
			}
		}
	} else {
		log.Println("Bad memcache handleGoogleCallback")
		//Try the database here
		// if db connection is good {
		// } else {
		// 	panic
		// }
	}
	//MEMCACHE OAUTH GET

	state := r.FormValue("state")
	if !oauthStateString[state] {
		log.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	} else {
		delete(oauthStateString, state)
	}

	//MEMCACHE OAUTH SET
	if mc != nil {
		buf := new(bytes.Buffer)
		enc := gob.NewEncoder(buf)
		enc.Encode(oauthStateString)

		mc.Set(&memcache.Item{Key: "oauthStateString", Value: buf.Bytes()})
	} else {
		//Try the database here
		// if db connection is good {
		// } else {
		// 	panic
		// }
	}
	//MEMCACHE OAUTH SET

	code := r.FormValue("code")
	//_, err := facebookOauthConfig.Exchange(oauth2.NoContext, code)
	token, err := facebookOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Println("Code exchange failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	response, err := http.Get("https://graph.facebook.com/v2.9/me?fields=id%2Cemail%2Cname&access_token=" + token.AccessToken)

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	var dat map[string]interface{}
	json.Unmarshal([]byte(contents), &dat)
	log.Println(dat)
	// //Get the email address
	email := dat["email"]
	var user model.User
	user.Email = email.(string)
	usr := model.SelectUserForLogin(user, true)
	SetSession(w, r, usr.Username)
	http.Redirect(w, r, "/", 302)
}

//Validates the login form request and populates the user
//struct
func ValidateLogin(user *model.User, r *http.Request) bool {
	user.Errors = make(map[string]string)

	if len(r.FormValue("name")) > 0 {
		user.Username = r.FormValue("name")
	} else {
		user.Errors["Username"] = "Please enter a valid username"
	}
	if len(r.FormValue("password")) > 0 {
		user.Password = r.FormValue("password")
	} else {
		user.Errors["Password"] = "Please enter a valid password"
	}
	return len(user.Errors) == 0
}
