package www

import (
	//"bytes"
	"fmt"
	"github.com/gorilla/securecookie"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"log"
	"model"
	"net/http"
	"strings"
)

// cookie handling

type Login struct {
}

var ()

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func (login *Login) RenderLoginIndexTemplate(w http.ResponseWriter, tmpl string, page *Page) {
	t, err := parseTempFiles(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	err = t.ExecuteTemplate(w, "base", page)
	if err != nil {
		log.Fatal(err)
	}
}

func GetUserName(request *http.Request) (userName string) {
	if cookie, err := request.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["name"]
		}
	}
	return userName
}

func SetSession(userName string, response http.ResponseWriter) {
	value := map[string]string{
		"name": userName,
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}

func ClearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}

// login handler

func (login *Login) loginHandler(response http.ResponseWriter, request *http.Request) {
	name := request.FormValue("name")
	pass := request.FormValue("password")
	if name != "" && pass != "" {
		user := model.User{
			Username: name,
			Password: pass,
		}
		usr := model.SelectUserForLogin(user)
		// usr := model.GetUser(name)
		if usr != nil {
			if usr.Password == pass {
				SetSession(name, response)
				http.Redirect(response, request, "/", 302)
			}
		}
	}
	login.RenderLoginIndexTemplate(response, "404", nil)
}

// logout handler

func (login *Login) logoutHandler(response http.ResponseWriter, request *http.Request) {
	ClearSession(response)
	http.Redirect(response, request, "/", 302)
}

func (login *Login) loginIndexHandler(response http.ResponseWriter, request *http.Request) {
	//fmt.Fprintf(response, indexPage)
	var page Page
	login.RenderLoginIndexTemplate(response, "loginindex", &page)
}

// internal page

const internalPage = `
<h1>Internal</h1>
<hr>
<small>User: %s</small>
<form method="post" action="/logout">
    <button type="submit">Logout</button>
</form>
`

func (login *Login) internalPageHandler(response http.ResponseWriter, request *http.Request) {
	userName := GetUserName(request)
	if userName != "" {
		fmt.Fprintf(response, internalPage, userName)
	} else {
		login.RenderLoginIndexTemplate(response, "404", nil)
	}
}

// internal page

const secondPage = `
<h1>Second Internal</h1>
<hr>
<small>User: %s</small>
<form method="post" action="/logout">
    <button type="submit">Logout</button>
</form>
`

func (login *Login) secondPageHandler(response http.ResponseWriter, request *http.Request) {
	userName := GetUserName(request)
	if userName != "" {
		fmt.Fprintf(response, secondPage, userName)
	} else {
		http.Redirect(response, request, "/", 404)
	}
}

func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	if state != oauthStateString {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	code := r.FormValue("code")
	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Println("Code exchange failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	//log.Printf("Content: %s\n", contents)
	s := string(contents[:])
	results := strings.Replace(strings.Replace(strings.Split(strings.Split(s, ",")[1], ":")[1], "\"", "", -1), " ", "", -1)
	//results := strings.Split(s, ",")
	log.Println(results)

	SetSession("hestert", w)
	http.Redirect(w, r, "/", 302)
}

// server main method

func (login *Login) Init() {
	log.Println("Init in www/login.go")
	http.HandleFunc("/loginIndex", login.loginIndexHandler)
	http.HandleFunc("/internal", login.internalPageHandler)
	http.HandleFunc("/second", login.secondPageHandler)
	http.HandleFunc("/login", login.loginHandler)
	http.HandleFunc("/logout", login.logoutHandler)
	http.HandleFunc("/GoogleLogin", handleGoogleLogin)
	http.HandleFunc("/GoogleCallback", handleGoogleCallback)
}
