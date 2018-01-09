// Copyright 2017 Rough Industries LLC. All rights reserved.
//model/user.go:package model
package model

import (
	"time"
)

//DATA STRUCTURES
//User data structure.
type User struct {
	ID                   int
	Username             string
	FirstName            string
	LastName             string
	Password             string
	NewPassword          string
	IsDisabled           bool
	Role                 UserRole
	Email                string
	VerificationCode     string
	VerificationInitTime time.Time
	VerificationComplete bool
	Authenticated        bool
	LastLogin            time.Time
	GoogleAccessToken    string
	FBAccessToken        string
	Errors               map[string]string
}

//User session data structure.  Keeps track of users logged in or accessing the
//site.
type UserSession struct {
	SessionKey        string
	User              User
	CSRFBase          string
	CSRFKey           []byte
	CSRF              string
	CSRFGenTime       time.Time
	LoginTime         time.Time
	LastSeenTime      time.Time
	IsDefaultUser     bool
	LastRemoteAddr    string
	LastXForwardedFor string
}

type OAuth struct {
	Key  string
	Time time.Time
}

//ENUMERATIONS - These must match the database one for one in both ID and order
//The integer values for the userrole enumeration
type UserRole int

const (
	SuperAdmin = 1 + iota
	Admin
	Viewer
	Editor
)

//The string values for the userrole enumeration
var UserRoleStrings = [...]string{
	"Super Administrator",
	"Administrator",
	"Viewer",
	"Editor",
}

// String returns the English name of the Roles ("Administrator", "Viewer", ...).
func (ur UserRole) String() string { return UserRoleStrings[ur-1] }
