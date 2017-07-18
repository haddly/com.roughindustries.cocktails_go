//model/post.connectors.go
package model

import (
	"bytes"
	"connectors"
	"database/sql"
	"log"
	"strconv"
	"strings"
	"time"
)

func InitPostTables() {
	conn, _ := connectors.GetDB()

	var temp string
	if err := conn.QueryRow("SHOW TABLES LIKE 'posttype';").Scan(&temp); err == nil {
		log.Println("PostType Table Exists")
	} else if err == sql.ErrNoRows {
		log.Println("Creating PostType Table")
		conn.Query("CREATE TABLE `commonwealthcocktails`.`posttype` (`idPostType` INT NOT NULL AUTO_INCREMENT,  PRIMARY KEY (`idPostType`));")
		conn.Query("ALTER TABLE `commonwealthcocktails`.`posttype` ADD COLUMN `postTypeName` VARCHAR(150) NOT NULL AFTER `idPostType`;")
		conn.Exec("INSERT INTO `commonwealthcocktails`.`posttype` (`idPostType`, `postTypeName`) VALUES ('1', 'InnerPost');")
		conn.Exec("INSERT INTO `commonwealthcocktails`.`posttype` (`idPostType`, `postTypeName`) VALUES ('2', 'FullPost');")
		conn.Exec("INSERT INTO `commonwealthcocktails`.`posttype` (`idPostType`, `postTypeName`) VALUES ('3', 'Blurb');")
	}

	if err := conn.QueryRow("SHOW TABLES LIKE 'poststatus';").Scan(&temp); err == nil {
		log.Println("PostStatus Table Exists")
	} else if err == sql.ErrNoRows {
		log.Println("Creating PostStatus Table")
		conn.Query("CREATE TABLE `commonwealthcocktails`.`poststatus` (`idPostStatus` INT NOT NULL AUTO_INCREMENT,PRIMARY KEY (`idPostStatus`));")
		conn.Query("ALTER TABLE `commonwealthcocktails`.`poststatus` ADD COLUMN `postStatusName` VARCHAR(150) NOT NULL AFTER `idPostStatus`;")
		conn.Exec("INSERT INTO `commonwealthcocktails`.`poststatus` (`idPostStatus`, `postStatusName`) VALUES ('1', 'Draft');")
		conn.Exec("INSERT INTO `commonwealthcocktails`.`poststatus` (`idPostStatus`, `postStatusName`) VALUES ('2', 'Pending');")
		conn.Exec("INSERT INTO `commonwealthcocktails`.`poststatus` (`idPostStatus`, `postStatusName`) VALUES ('3', 'Private');")
		conn.Exec("INSERT INTO `commonwealthcocktails`.`poststatus` (`idPostStatus`, `postStatusName`) VALUES ('4', 'Publish');")

	}

	if err := conn.QueryRow("SHOW TABLES LIKE 'post';").Scan(&temp); err == nil {
		log.Println("Post Table Exists")
	} else if err == sql.ErrNoRows {
		log.Println("Creating Post Table")
		conn.Query("CREATE TABLE `commonwealthcocktails`.`post` (`idPost` INT NOT NULL AUTO_INCREMENT,PRIMARY KEY (`idPost`));")
		conn.Query("ALTER TABLE `commonwealthcocktails`.`post`" +
			"ADD COLUMN `postAuthor` VARCHAR(150) NOT NULL AFTER `idPost`," + //PostAuthor
			"ADD COLUMN `postTitle` VARCHAR(250) NOT NULL AFTER `postAuthor`," + //PostTitle
			"ADD COLUMN `postContent` VARCHAR(1500) AFTER `postTitle`," + //PostContent
			"ADD COLUMN `postExcerpt` VARCHAR(750) AFTER `postContent`," + //PostExcerpt
			"ADD COLUMN `postName` VARCHAR(750) AFTER `postExcerpt`," + //PostName
			"ADD COLUMN `postExcerptImagePath` VARCHAR(750) AFTER `postName`," + //PostExcerptImagePath
			"ADD COLUMN `postExcerptImage` VARCHAR(750) AFTER `postExcerptImagePath`," + //PostExcerptImage
			"ADD COLUMN `postExcerptImageSourceName` VARCHAR(750) AFTER `postExcerptImage`," + //PostExcerptImageSourceName
			"ADD COLUMN `postExcerptImageSourceLink` VARCHAR(750) AFTER `postExcerptImageSourceName`," + //PostExcerptImageSourceLink
			"ADD COLUMN `postDate` DATETIME NOT NULL AFTER `postExcerptImageSourceLink`," + //PostDate
			"ADD COLUMN `postStatus` INT NOT NULL, ADD CONSTRAINT post_poststatus_id_fk FOREIGN KEY(postStatus) REFERENCES poststatus(idPostStatus)," + //PostStatus
			"ADD COLUMN `postType` INT NOT NULL, ADD CONSTRAINT post_posttype_id_fk FOREIGN KEY(postType) REFERENCES posttype(idPostType);") //PostType

	}
}

func ProcessPosts() {
	conn, _ := connectors.GetDB()

	for _, post := range Posts {
		log.Println(post.PostTitle)
		var buffer bytes.Buffer
		buffer.WriteString("INSERT INTO `commonwealthcocktails`.`post` SET ")
		if post.PostTitle != "" {
			buffer.WriteString("`postTitle`=\"" + post.PostTitle + "\",")
		}
		if post.PostName != "" {
			buffer.WriteString("`postName`=\"" + post.PostName + "\",")
		}
		if post.PostAuthor != "" {
			buffer.WriteString("`postAuthor`=\"" + post.PostAuthor + "\",")
		}
		if post.PostContent != "" {
			sqlString := strings.Replace(string(post.PostContent), "\\", "\\\\", -1)
			sqlString = strings.Replace(sqlString, "\"", "\\\"", -1)
			buffer.WriteString("`postContent`=\"" + sqlString + "\",")
		}
		if post.PostExcerpt != "" {
			sqlString := strings.Replace(string(post.PostExcerpt), "\\", "\\\\", -1)
			sqlString = strings.Replace(sqlString, "\"", "\\\"", -1)
			buffer.WriteString("`postExcerpt`=\"" + sqlString + "\",")
		}
		buffer.WriteString(" `postDate`=\"" + post.PostDate.Format(time.RFC3339) + "\",")
		buffer.WriteString(" `postStatus`=" + strconv.Itoa(int(post.PostStatus)) + ",")
		buffer.WriteString(" `postType`=" + strconv.Itoa(int(post.PostType)) + ",")
		if post.PostExcerptImagePath != "" {
			buffer.WriteString("`postExcerptImagePath`=\"" + post.PostExcerptImagePath + "\",")
		}
		if post.PostExcerptImage != "" {
			buffer.WriteString("`postExcerptImage`=\"" + post.PostExcerptImage + "\",")
		}
		if post.PostExcerptImageSourceName != "" {
			buffer.WriteString("`postExcerptImageSourceName`=\"" + post.PostExcerptImageSourceName + "\",")
		}
		if post.PostExcerptImageSourceLink != "" {
			buffer.WriteString("`postExcerptImageSourceLink`=\"" + post.PostExcerptImageSourceLink + "\",")
		}
		query := buffer.String()
		query = strings.TrimRight(query, ",")
		query = query + ";"
		log.Println(query)
		conn.Exec(query)
	}
}
