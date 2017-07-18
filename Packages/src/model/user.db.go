//model/user.connectors.go
package model

import (
	"bytes"
	"connectors"
	"database/sql"
	"html"
	"log"
	"strings"
)

func InitUserTables() {
	conn, _ := connectors.GetDB()

	var temp string
	if err := conn.QueryRow("SHOW TABLES LIKE 'users';").Scan(&temp); err == nil {
		log.Println("Users Table Exists")
	} else if err == sql.ErrNoRows {
		log.Println("Creating Users Table")
		conn.Query("CREATE TABLE `commonwealthcocktails`.`users` (`idUser` INT NOT NULL AUTO_INCREMENT,PRIMARY KEY (`idUser`));")
		conn.Query("ALTER TABLE `commonwealthcocktails`.`users`" +
			"ADD COLUMN `userName` VARCHAR(150) NOT NULL AFTER `idUser`," + //userName
			"ADD COLUMN `userPassword` VARCHAR(250) NOT NULL AFTER `userName`," + //userPassword
			"ADD COLUMN `userEmail` VARCHAR(250) NOT NULL AFTER `userPassword`," + //userEmail
			"ADD COLUMN `userLastLogin` DATETIME NOT NULL AFTER `userEmail`;") //userLastLogin

	}
}

func ProcessUsers() {
	conn, _ := connectors.GetDB()

	for _, user := range Users {
		log.Println(user.Username)
		var buffer bytes.Buffer
		var args []interface{}
		buffer.WriteString("INSERT INTO `commonwealthcocktails`.`users` SET ")
		if user.Username != "" {
			buffer.WriteString("`userName`=?,")
			args = append(args, html.EscapeString(user.Username))
		}
		if user.Password != "" {
			buffer.WriteString("`userPassword`=?,")
			args = append(args, user.Password)
		}
		if user.Email != "" {
			buffer.WriteString("`userEmail`=?,")
			args = append(args, user.Email)
		}
		query := buffer.String()
		query = strings.TrimRight(query, ",")
		query = query + ";"
		log.Println(query)
		conn.Exec(query, args...)
	}
}

func SelectUserForLogin(user User, isOauth bool) *User {
	var ret User
	conn, _ := connectors.GetDB()

	log.Println(user.Username)
	var buffer bytes.Buffer
	var canQuery = false
	var args []interface{}
	buffer.WriteString("SELECT `idUser`, `userName`, `userPassword`, `userEmail` FROM `commonwealthcocktails`.`users` WHERE ")
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
