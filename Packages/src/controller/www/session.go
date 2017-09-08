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
	"net/http"
)

//the sessions are stored here in coockies
var store = sessions.NewCookieStore([]byte(randSeq(64)))

//memcache session mapping
var managed_sessions = make(map[string]string)

//get the session from the cookies and cross reference it with the
//memcache
//TODO: setup database session mapping store
func GetSession(r *http.Request) (string, bool) {
	session, err := store.Get(r, "cocktails")
	if err != nil {
		panic("Bad GetSession Request!!!")
		return "", false
	}
	// Retrieve our struct and type-assert it
	if session_id, ok := session.Values["session_id"].(string); !ok {
		log.Println("No Session ID")
		return "", false
	} else {
		log.Println("Good Session ID")

		//MEMCACHE SESSION GET
		mc, _ := connectors.GetMC()
		if mc != nil {
			item := new(memcache.Item)
			item, _ = mc.Get("managed_sessions")
			managed_sessions = make(map[string]string)
			if item != nil {
				if len(item.Value) > 0 {
					read := bytes.NewReader(item.Value)
					dec := gob.NewDecoder(read)
					dec.Decode(&managed_sessions)
				}
			}
			data := managed_sessions[session_id+"_u"]
			if data == "" {
				log.Println("Not Authenticated")
				return data, false
			} else {
				log.Println("Authenticated")
				return data, true

			}
		} else {
			//Try the database here
			// if db connection is good {
			// } else {
			// 	panic
			// }
		}
		//MEMCACHE SESSION SET
	}
	return "", false
}

//set the session for the cookies and cross reference it to the
//memcache
//TODO: setup database session mapping store
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

	//MEMCACHE SESSION SET
	mc, _ := connectors.GetMC()
	if mc != nil {
		item := new(memcache.Item)
		item, _ = mc.Get("managed_sessions")
		managed_sessions = make(map[string]string)
		if item != nil {
			if len(item.Value) > 0 {
				read := bytes.NewReader(item.Value)
				dec := gob.NewDecoder(read)
				dec.Decode(&managed_sessions)
			}
		}
		managed_sessions[session_id+"_u"] = data
		buf := new(bytes.Buffer)
		enc := gob.NewEncoder(buf)
		enc.Encode(managed_sessions)

		mc.Set(&memcache.Item{Key: "managed_sessions", Value: buf.Bytes()})
	} else {
		//Try the database here
		// if db connection is good {
		// } else {
		// 	panic
		// }
	}
	//MEMCACHE SESSION SET

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
			item := new(memcache.Item)
			item, _ = mc.Get("managed_sessions")
			managed_sessions = make(map[string]string)
			if item != nil {
				if len(item.Value) > 0 {
					read := bytes.NewReader(item.Value)
					dec := gob.NewDecoder(read)
					dec.Decode(&managed_sessions)
				}
				delete(managed_sessions, session_id+"_u")
				buf := new(bytes.Buffer)
				enc := gob.NewEncoder(buf)
				enc.Encode(managed_sessions)

				mc.Set(&memcache.Item{Key: "managed_sessions", Value: buf.Bytes()})
			}
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
