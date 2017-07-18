//model/user.dat.go
package model

import (
	"time"
)

var Users = []User{
	User{
		ID:        0,
		Username:  "test",
		Password:  "password",
		Email:     "",
		LastLogin: time.Now(),
	},
}
