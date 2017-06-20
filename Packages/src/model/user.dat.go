//model/user.dat.go
package model

import (
	"time"
)

var Users = []User{
	User{
		ID:        0,
		Username:  "hestert",
		Password:  "password",
		LastLogin: time.Now(),
	},
}
