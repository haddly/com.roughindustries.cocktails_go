//model/user.go
package model

import (
	"time"
)

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

type UserSession struct {
	SessionKey   string
	Username     string
	UserID       int
	LoginTime    time.Time
	LastSeenTime time.Time
}
