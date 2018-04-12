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
	"encoding/gob"
	"github.com/CommonwealthCocktails/connectors"
	"github.com/CommonwealthCocktails/model"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/gorilla/sessions"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

//get the session from the cookies and cross reference it with the
//memcache
//TODO: setup database session mapping store
func GetSession(w http.ResponseWriter, r *http.Request, site string) (model.UserSession, bool) {
	store := sessions.NewCookieStore([]byte(store_key))
	session, err := store.Get(r, "cocktails")
	if err != nil {
		log.Errorln("ERROR: Bad GetSession Request!!!")
		log.Errorln(err)
		return *new(model.UserSession), false
	}
	// Retrieve our struct and type-assert it
	if session_id, ok := session.Values["session_id"].(string); !ok {
		log.Infoln("No Session ID")
		us := new(model.UserSession)
		us.IsDefaultUser = false
		us.LoginTime = time.Now()
		us.LastSeenTime = time.Now()
		SetSession(w, r, us, true, site)
		return *us, false
	} else if csrf, ok := session.Values["csrf"].(string); !ok {
		log.Errorln("ERROR: No CSRF, possible attack!")
		us := new(model.UserSession)
		us.IsDefaultUser = false
		us.LoginTime = time.Now()
		us.LastSeenTime = time.Now()
		SetSession(w, r, us, true, site)
		return *us, false
	} else {
		mc, _ := connectors.GetMC()
		var userSession model.UserSession
		if mc != nil {
			item := new(memcache.Item)
			item, _ = mc.Get(session_id)
			if item != nil {
				if len(item.Value) > 0 {
					read := bytes.NewReader(item.Value)
					dec := gob.NewDecoder(read)
					dec.Decode(&userSession)
				}
			}
		} else {
			userSession.SessionKey = session_id
			sessions := userSession.SelectUserSession(site)
			if len(sessions) > 0 {
				userSession = sessions[0]
			}
		}
		//load the cookies csrf into the csrf for the session so if we
		//need to check it later we can.  Whatever csrf value was the last
		//one issued not the one from this page
		userSession.CSRF = csrf
		//log.Print("LastSeenTime")
		//log.Print(userSession.LastSeenTime)
		//Authenticate
		if !userSession.IsDefaultUser && userSession.User.Username == "" {
			//The user hasn't logged in so this is to track the session
			//with no user attached.
			return userSession, false
		} else if time.Since(userSession.LoginTime).Hours() >= 24 {
			//Has it been to long sense the user logged in so they have to
			//login again, i.e. even if they have been using the site for
			//24hours we cap the total time they can be logged into one
			//session
			log.Infoln("User has been logged in for over 24 hours")
			log.Infoln("Looks like this is an old session.")
			ClearSession(w, r, site)
			return userSession, false
		} else {
			//Has it been to log sense they last touched the website.  This
			//prevents idle web clients from keeping a session open.  So
			//if they went to get a cup of coffee and got distracted it will
			//close the session so no one hopes on if it has been more then
			//x seconds
			if time.Since(userSession.LastSeenTime).Minutes() >= 10 {
				log.Infoln("User has been inactive in for over 10 minutes")
				log.Infoln("Looks like this is an old session.")
				ClearSession(w, r, site)
				return userSession, false
			} else {
				if !userSession.User.VerificationComplete {
					log.Infoln("User has not completed verification process")
					ClearSession(w, r, site)
					return userSession, false
				} else {
					log.Infoln("Authenticated")
					return userSession, true
				}
			}
		}
	}
	us := new(model.UserSession)
	us.LoginTime = time.Now()
	us.LastSeenTime = time.Now()
	SetSession(w, r, us, true, site)
	return *us, false
}

//set the session for the cookies and cross reference it to the
//memcache
//TODO: setup database session mapping store
func SetSession(w http.ResponseWriter, r *http.Request, us *model.UserSession, regenSessionID bool, site string) string {
	store := sessions.NewCookieStore([]byte(store_key))
	session, err := store.Get(r, "cocktails")
	if err != nil {
		Error404(w, err.Error(), site)
		return ""
	}
	// Set some session values.
	//check for session_id, if none then create it
	var session_id string
	var ok bool
	session.Options.MaxAge = 0
	session.Options.HttpOnly = true
	session.Options.Secure = true
	prev_session_id := ""
	session_id, ok = session.Values["session_id"].(string)
	if !ok || regenSessionID {
		if ok {
			prev_session_id = session_id
		}
		session_id, _ = GenerateRandomString(32)
		session.Values["session_id"] = session_id
		us.CSRFBase, _ = GenerateRandomString(32)
		us.CSRFKey, _ = GenerateRandomBytes(32)
		us.CSRFGenTime = time.Now()
	} else {
		if us.CSRFBase == "" || len(us.CSRFKey) == 0 {
			log.Infoln("Bad CSRF Base or Key")
			us.CSRFBase, _ = GenerateRandomString(32)
			us.CSRFKey, _ = GenerateRandomBytes(32)
			us.CSRFGenTime = time.Now()
		}
		prev_session_id = session_id
	}
	us.SessionKey = session_id
	//set the csrf value
	session.Values["csrf"] = us.CSRF
	session.Save(r, w)
	mc, _ := connectors.GetMC()
	if mc != nil {
		buf := new(bytes.Buffer)
		enc := gob.NewEncoder(buf)
		enc.Encode(*us)
		//log.Infoln("Updating session")
		if prev_session_id != "" {
			mc.Delete(prev_session_id)
		}
		mc.Set(&memcache.Item{Key: session_id, Value: buf.Bytes()})
	} else {
		if prev_session_id != "" {
			var prevus model.UserSession
			prevus.SessionKey = prev_session_id
			sessions := us.SelectUserSession(site)
			if len(sessions) > 0 {
				us.UpdateUserSession(prev_session_id, site)
			} else {
				us.InsertUserSession(site)
			}
		} else {
			us.InsertUserSession(site)
		}
	}
	return session_id
}

//clear the session for the cookies and the memcache
//TODO: setup database session mapping store
func ClearSession(w http.ResponseWriter, r *http.Request, site string) {
	store := sessions.NewCookieStore([]byte(store_key))
	session, err := store.Get(r, "cocktails")
	if err != nil {
		Error404(w, err.Error(), site)
		return
	}
	if session_id, ok := session.Values["session_id"].(string); !ok {
		log.Infoln("Bad Session ID")
	} else {
		log.Infoln("Clearing session")
		mc, _ := connectors.GetMC()
		var userSession model.UserSession
		if mc != nil {
			item := new(memcache.Item)
			item, _ = mc.Get(session_id)
			if item != nil {
				if len(item.Value) > 0 {
					read := bytes.NewReader(item.Value)
					dec := gob.NewDecoder(read)
					dec.Decode(&userSession)
				}
			}
		} else {
			userSession.SessionKey = session_id
			sessions := userSession.SelectUserSession(site)
			if len(sessions) > 0 {
				userSession = sessions[0]
			}
		}
		userSession.User = model.User{}
		userSession.IsDefaultUser = false
		userSession.LoginTime = time.Time{}
		SetSession(w, r, &userSession, true, site)
	}
}
