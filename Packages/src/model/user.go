//model/user.go
package model

import (
	"time"
)

type User struct {
	ID          int
	Username    string
	Password    string
	LastLogin   time.Time
}

func GetUser(username string) *User {
    for _, element := range Users {
        if element.Username == username {
            return &element
        }
    }
	return nil
}
