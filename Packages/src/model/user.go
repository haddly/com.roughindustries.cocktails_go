//model/user.go
package model

import (
	"time"
)

type User struct {
	ID            int
	Username      string
	Password      string
	Email         string
	Authenticated bool
	LastLogin     time.Time
	Errors        map[string]string
}
