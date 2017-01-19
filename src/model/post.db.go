//model/post.db.go
package model

import (
	"database/sql"
	"db"
	"log"
)

func InitPostTables() {
	conn, _ := db.GetDB()
	var temp string
	if err := conn.QueryRow("SHOW TABLES LIKE 'posttype';").Scan(&temp); err == nil {
		log.Println("PostType Table Exists")
	} else if err == sql.ErrNoRows {
		log.Println("Creating PostType Table")
		conn.Query("CREATE TABLE `commonwealthcocktails`.`posttype` (`idPostType` INT NOT NULL AUTO_INCREMENT,  PRIMARY KEY (`idPostType`));")
	}

	if err := conn.QueryRow("SHOW TABLES LIKE 'poststatus';").Scan(&temp); err == nil {
		log.Println("PostStatus Table Exists")
	} else if err == sql.ErrNoRows {
		log.Println("Creating PostStatus Table")
		conn.Query("CREATE TABLE `commonwealthcocktails`.`poststatus` (`idPostStatus` INT NOT NULL AUTO_INCREMENT,PRIMARY KEY (`idPostStatus`));")
	}

	if err := conn.QueryRow("SHOW TABLES LIKE 'post';").Scan(&temp); err == nil {
		log.Println("Post Table Exists")
	} else if err == sql.ErrNoRows {
		log.Println("Creating Post Table")
		conn.Query("CREATE TABLE `commonwealthcocktails`.`post` (`idPost` INT NOT NULL AUTO_INCREMENT,PRIMARY KEY (`idPost`));")
	}
}
