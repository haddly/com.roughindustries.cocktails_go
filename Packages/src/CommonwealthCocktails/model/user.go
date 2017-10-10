// Copyright 2017 Rough Industries LLC. All rights reserved.
//model/user.go:package model
package model

import (
	"time"
)

//DATA STRUCTURES
//User data structure.
type User struct {
	ID            int
	Username      string
	FullName      string
	Password      string
	IsDisabled    bool
	Roles         string
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

//ENUMERATIONS - These must match the database one for one in both ID and order
//The integer values for the producttype enumeration
type RolesConst int

const (
	Administrator = 1 + iota
	Contributor
)

//The string values for the producttype enumeration
var RolesStrings = [...]string{
	"Administrator",
	"Contributor",
}

// String returns the English name of the Roles ("Administrator", "Contributor", ...).
func (rt RolesConst) String() string { return RolesStrings[rt-1] }
