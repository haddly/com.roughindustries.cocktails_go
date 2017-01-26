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
			"ADD COLUMN `postAuthor` VARCHAR(150) NULL AFTER `idPost`," + //PostAuthor
			"ADD COLUMN `postTitle` VARCHAR(250) NOT NULL AFTER `postAuthor`," + //PostTitle
			"ADD COLUMN `postContent` VARCHAR(1500) NOT NULL AFTER `postTitle`," + //PostContent
			"ADD COLUMN `postExcerpt` VARCHAR(750) NOT NULL AFTER `postContent`," + //PostExcerpt
			"ADD COLUMN `PostName` VARCHAR(750) NOT NULL AFTER `postExcerpt`," + //PostName
			"ADD COLUMN `PostExcerptImagePath` VARCHAR(750) NOT NULL AFTER `PostName`," + //PostExcerptImagePath
			"ADD COLUMN `PostExcerptImage` VARCHAR(750) NOT NULL AFTER `PostExcerptImagePath`," + //PostExcerptImage
			"ADD COLUMN `PostExcerptImageSourceName` VARCHAR(750) NOT NULL AFTER `PostExcerptImage`," + //PostExcerptImageSourceName
			"ADD COLUMN `PostExcerptImageSourceLink` VARCHAR(750) NOT NULL AFTER `PostExcerptImageSourceName`," + //PostExcerptImageSourceLink
			"ADD COLUMN `postDate` DATETIME NOT NULL AFTER `PostExcerptImageSourceLink`," + //PostDate
			"ADD COLUMN `postStatus` INT NOT NULL, ADD CONSTRAINT post_poststatus_id_fk FOREIGN KEY(postStatus) REFERENCES poststatus(idPostStatus)," + //PostStatus
			"ADD COLUMN `postType` INT NOT NULL, ADD CONSTRAINT post_posttype_id_fk FOREIGN KEY(postType) REFERENCES posttype(idPostType);") //PostType

	}
}
