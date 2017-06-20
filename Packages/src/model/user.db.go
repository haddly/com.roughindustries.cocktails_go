//model/user.db.go
package model

import (
	"bytes"
	"database/sql"
	"db"
	"log"
	"strings"
)

func InitUserTables() {
	conn, _ := db.GetDB()

	var temp string
	if err := conn.QueryRow("SHOW TABLES LIKE 'users';").Scan(&temp); err == nil {
		log.Println("Users Table Exists")
	} else if err == sql.ErrNoRows {
		log.Println("Creating Users Table")
		conn.Query("CREATE TABLE `commonwealthcocktails`.`users` (`idUser` INT NOT NULL AUTO_INCREMENT,PRIMARY KEY (`idUser`));")
		conn.Query("ALTER TABLE `commonwealthcocktails`.`users`" +
			"ADD COLUMN `userName` VARCHAR(150) NOT NULL AFTER `idUser`," + //userName
			"ADD COLUMN `userPassword` VARCHAR(250) NOT NULL AFTER `userName`," + //userPassword
			"ADD COLUMN `userLastLogin` DATETIME NOT NULL AFTER `userPassword`;") //userLastLogin

	}
}

func ProcessUsers() {
	conn, _ := db.GetDB()

	for _, user := range Users {
		log.Println(user.Username)
		var buffer bytes.Buffer
		buffer.WriteString("INSERT INTO `commonwealthcocktails`.`users` SET ")
		if user.Username != "" {
			buffer.WriteString("`userName`=\"" + user.Username + "\",")
		}
		if user.Password != "" {
			buffer.WriteString("`userPassword`=\"" + user.Password + "\",")
		}
		query := buffer.String()
		query = strings.TrimRight(query, ",")
		query = query + ";"
		log.Println(query)
		conn.Exec(query)
	}
}

func SelectUserForLogin(user User) *User {
	var ret User
	conn, _ := db.GetDB()

	log.Println(user.Username)
	var buffer bytes.Buffer
	var canQuery = false
	buffer.WriteString("SELECT `idUser`, `userName`, `userPassword` FROM `commonwealthcocktails`.`users` WHERE ")
	if user.Username != "" {
		buffer.WriteString("`userName`=\"" + user.Username + "\" AND")
		if user.Password != "" {
			buffer.WriteString("`userPassword`=\"" + user.Password + "\" AND")
			canQuery = true
		}
	}

	if canQuery {
		query := buffer.String()
		query = strings.TrimRight(query, " AND")
		query = query + ";"
		log.Println(query)
		rows, err := conn.Query(query)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&ret.ID, &ret.Username, &ret.Password)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(ret.ID, ret.Username, ret.Password)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
	}
	return &ret
}
