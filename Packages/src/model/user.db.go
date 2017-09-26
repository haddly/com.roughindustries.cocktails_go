// Copyright 2017 Rough Industries LLC. All rights reserved.
//model/user.db.go:package model
package model

import (
	"bytes"
	"connectors"
	"html"
	"log"
	"strings"
)

//CREATE, UPDATE, DELETE

//SELECTS
//Check the database for a user based on user name and password, if Ouath is
//set then check the user based on email
func (user *User) SelectUserForLogin(isOauth bool) *User {
	var ret User
	conn, _ := connectors.GetDB()
	log.Println(user.Username)
	var buffer bytes.Buffer
	var canQuery = false
	var args []interface{}
	buffer.WriteString("SELECT `idUser`, `userName`, `userPassword`, `userEmail` FROM `users` WHERE ")
	if user.Username != "" {
		buffer.WriteString(" `userName`=? AND")
		args = append(args, html.EscapeString(user.Username))
		canQuery = true
	}
	if isOauth {
		if user.Email != "" {
			buffer.WriteString("`userEmail`=? AND")
			args = append(args, user.Email)
			canQuery = true
		}
	}

	if canQuery {
		query := buffer.String()
		query = strings.TrimRight(query, " AND")
		query = query + ";"
		log.Println(query)
		rows, err := conn.Query(query, args...)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&ret.ID, &ret.Username, &ret.Password, &ret.Email)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(ret.ID, ret.Username, ret.Password, ret.Email)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
	}
	return &ret
}
