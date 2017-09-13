//model/user.connectors.go
package model

import (
	"bytes"
	"connectors"
	"html"
	"log"
	"strings"
)

func SelectUserForLogin(user User, isOauth bool) *User {
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
		if user.Password != "" {
			buffer.WriteString(" `userPassword`=? AND")
			args = append(args, user.Password)
			canQuery = true
		}
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
