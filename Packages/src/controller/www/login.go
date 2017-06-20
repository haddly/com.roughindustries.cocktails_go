package www

import (
	"fmt"
	"github.com/gorilla/securecookie"
	"log"
	"model"
	"net/http"
)

// cookie handling

type Login struct {
}

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

// server main method

func (login *Login) Init() {
	log.Println("Init in www/login.go")
	http.HandleFunc("/loginIndex", login.loginIndexHandler)
	http.HandleFunc("/internal", login.internalPageHandler)
	http.HandleFunc("/second", login.secondPageHandler)
	http.HandleFunc("/login", login.loginHandler)
	http.HandleFunc("/logout", login.logoutHandler)
}
