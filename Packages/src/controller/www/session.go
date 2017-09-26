// Copyright 2017 Rough Industries LLC. All rights reserved.
//controller/www/session.go: Because I am using Google Cloud Platform
//and Golang I need this to be scalable which requires persistance
//outside of the httpd golang instance running.  If I just used
//a single instance this wouldn't be an issue sense I could use local
//memory.  With multiple instances this is not possible. GCP
//automatically creates 2 instances with the basic flex env which if I am trying
//to share managed sessions over the 2 instances requires remote persistance
//storage for the sessions.  The idea is to save the managed sessions in a
//memcache and/or a database but try to use the memcache for the bulk of the gets
//for the manages sessions, because of performance.
package www

import (
	"bytes"
	"connectors"
	"encoding/gob"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/gorilla/sessions"
	"log"
	"model"
	"net/http"
	"time"
)

//the sessions are stored here in coockies
var store = sessions.NewCookieStore([]byte(randSeq(64)))

//get the session from the cookies and cross reference it with the
//memcache
//TODO: setup database session mapping store
func GetSession(w http.ResponseWriter, r *http.Request) (model.UserSession, bool) {
	session, err := store.Get(r, "cocktails")
	if err != nil {
		panic("Bad GetSession Request!!!")
		return *new(model.UserSession), false
	}
	// Retrieve our struct and type-assert it
	if session_id, ok := session.Values["session_id"].(string); !ok {
		log.Println("No Session ID")
		return *new(model.UserSession), false
	} else {
		mc, _ := connectors.GetMC()
		if mc != nil {
			item := new(memcache.Item)
			item, _ = mc.Get(session_id)
			var userSession model.UserSession
			if item != nil {
				if len(item.Value) > 0 {
					read := bytes.NewReader(item.Value)
					dec := gob.NewDecoder(read)
					dec.Decode(&userSession)
				}
			}
			//Authenticate
			//Has it been to long sense the user logged in so they have to
			//login again, i.e. even if they have been using the site for
			//24hours we cap the total time they can be logged into one
			//session
			if time.Since(userSession.LoginTime).Hours() >= 24 {
				log.Println("User has been logged in for over 24 hours")
				log.Println("Looks like this is an old session.")
				ClearSession(w, r)
				return userSession, false
			} else {
				//Has it been to log sense they last touched the website.  This
				//prevents idle web clients from keeping a session open.  So
				//if they went to get a cup of coffee and got distracted it will
				//close the session so no one hopes on if it has been more then
				//x seconds
				if time.Since(userSession.LastSeenTime).Minutes() >= 10 {
					log.Println("User has been inactive in for over 10 minutes")
					log.Println("Looks like this is an old session.")
					ClearSession(w, r)
					return userSession, false
				} else {
					log.Println("Authenticated")
					//reset the last seen time
					userSession.LastSeenTime = time.Now()
					return userSession, true
				}
			}

		} else {
			//Try the database here
			// if db connection is good {
			// } else {
			// 	panic
			// }
		}
	}
	return *new(model.UserSession), false
}

//set the session for the cookies and cross reference it to the
//memcache
//TODO: setup database session mapping store
func SetSession(w http.ResponseWriter, r *http.Request, us *model.UserSession) string {
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
	mc, _ := connectors.GetMC()
	if mc != nil {
		us.SessionKey = session_id
		buf := new(bytes.Buffer)
		enc := gob.NewEncoder(buf)
		enc.Encode(*us)
		log.Println("Adding session")
		mc.Set(&memcache.Item{Key: session_id, Value: buf.Bytes()})
	} else {
		//Try the database here
		// if db connection is good {
		// } else {
		// 	panic
		// }
	}
	return session_id
}

//clear the session for the cookies and the memcache
//TODO: setup database session mapping store
func ClearSession(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "cocktails")
	if err != nil {
		Error404(w, err.Error())
		return
	}
	session.Options.MaxAge = -1
	session.Save(r, w)

	if session_id, ok := session.Values["session_id"].(string); !ok {
		log.Println("Bad Session ID")
	} else {
		//MEMCACHE SESSION DELETE
		mc, _ := connectors.GetMC()
		if mc != nil {
			log.Println("Deleteing session")
			mc.Delete(session_id)
		} else {
			//Try the database here
			// if db connection is good {
			// } else {
			// 	panic
			// }
		}
		//MEMCACHE SESSION DELETE
	}
}
