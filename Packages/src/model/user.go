//model/user.go
package model

import (
	"strings"
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

func (user *User) Validate() bool {
	user.Errors = make(map[string]string)

	if strings.TrimSpace(user.Username) == "" {
		user.Errors["Username"] = "Please enter a valid username"
	}

	if len(user.Password) == 0 {
		user.Errors["Password"] = "Please enter a valid password"
	}

	return len(user.Errors) == 0
}
