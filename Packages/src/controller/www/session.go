// Package session
package www

import (
	"github.com/gorilla/sessions"
	"log"
	"net/http"
)

var store = sessions.NewCookieStore([]byte(randSeq(64)))

var managed_sessions = make(map[string]string)

func GetSession(r *http.Request) (string, bool) {
	session, err := store.Get(r, "cocktails")
	if err != nil {
		panic("Bad GetSession Request!!!")
		return "", false
	}
	// Retrieve our struct and type-assert it
	if session_id, ok := session.Values["session_id"].(string); !ok {
		log.Println("Bad Session ID")
		return "", false
	} else {
		log.Println("Good Session ID")
		data := managed_sessions[session_id+"_u"]
		if data == "" {
			log.Println("Not Authenticated")
			return data, false
		} else {
			log.Println("Authenticated")
			return data, true

		}
	}
	return "", false
}

func SetSession(w http.ResponseWriter, r *http.Request, data string) string {
	session, err := store.Get(r, "cocktails")
	if err != nil {
		Error404(w, err.Error())
		return ""
	}
	session.Options.MaxAge = 0
	session.Options.HttpOnly = true
	session.Options.Secure = true
	log.Println(session.Options)
	// Set some session values.
	session_id := randSeq(64)
	session.Values["session_id"] = session_id
	session.Save(r, w)
	managed_sessions[session_id+"_u"] = data
	return session_id
}

func ClearSession(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "cocktails")
	if err != nil {
		Error404(w, err.Error())
		return
	}
	session.Options.MaxAge = -1
	session.Save(r, w)
}

func init() {

}
