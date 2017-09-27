// Copyright 2017 Rough Industries LLC. All rights reserved.
//model/user.go:package model
package model

import (
	"time"
)

//DATA STRUCTURES
//User data structure.
type User struct {
	ID         int
	Username   string
	FullName   string
	Password   string
	IsDisabled bool

	Email         string
	Authenticated bool
	LastLogin     time.Time
	Errors        map[string]string
}

//User session data structure.  Keeps track of users logged in or accessing the
//site.
type UserSession struct {
	SessionKey        string
	User              User
	CSRF              string
	CSRFGenTime       time.Time
	LoginTime         time.Time
	LastSeenTime      time.Time
	LastRemoteAddr    string
	LastXForwardedFor string
}
