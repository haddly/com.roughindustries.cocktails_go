// Copyright 2017 Rough Industries LLC. All rights reserved.
//model/post.db.go:package model
package model

import (
	"bytes"
	"github.com/CommonwealthCocktails/connectors"
	log "github.com/sirupsen/logrus"
	"html"
	"html/template"
	"strconv"
	"strings"
	"time"
)

//CREATE, UPDATE, DELETE
//Insert a post record into the database
func (post *Post) InsertPost(site string) int {
	//set the ID to zero to indicate an insert
	post.ID = 0
	post.PostCreateDate = time.Now()
	post.PostModifiedDate = time.Now()
	return post.processPost(site)
}

//Update a post record in the database based on ID
func (post *Post) UpdatePost(site string) int {
	p := post.SelectPost(site)
	post.PostCreateDate = p[0].PostCreateDate
	post.PostModifiedDate = time.Now()
	return post.processPost(site)
}

//Process an insert or an update
func (post *Post) processPost(site string) int {
	conn, _ := connectors.GetDBFromMap(site) //get db connection
	var args []interface{}                   //arguments for variables in the data struct
	var buffer bytes.Buffer                  //buffer for the query

	//If the ID is zero then do an insert else do an update based on the ID
	if post.ID == 0 {
		buffer.WriteString("INSERT INTO `post` ( ")
	} else {
		buffer.WriteString("UPDATE `post` SET ")
	}

	//Append the correct columns to be added based on data available in the
	//data structure
	if post.PostTitle != "" {
		if post.ID == 0 {
			buffer.WriteString("`postTitle`,")
		} else {
			buffer.WriteString("`postTitle`=?,")
		}
		args = append(args, string(post.PostTitle))
	}
	if post.PostExcerpt != "" {
		if post.ID == 0 {
			buffer.WriteString("`postExcerpt`,")
		} else {
			buffer.WriteString("`postExcerpt`=?,")
		}
		args = append(args, string(post.PostExcerpt))
	}
	if post.PostContent != "" {
		if post.ID == 0 {
			buffer.WriteString("`postContent`,")
		} else {
			buffer.WriteString("`postContent`=?,")
		}
		args = append(args, string(post.PostContent))
	}
	if post.PostImage != "" {
		if post.ID == 0 {
			buffer.WriteString("`postImage`,")
		} else {
			buffer.WriteString("`postImage`=?,")
		}
		args = append(args, string(post.PostImage))
	}
	if post.ID == 0 {
		buffer.WriteString("`postAuthor`,")
	} else {
		buffer.WriteString("`postAuthor`=?,")
	}
	args = append(args, post.PostAuthor.ID)
	if post.ID == 0 {
		buffer.WriteString("`postStatus`,")
	} else {
		buffer.WriteString("`postStatus`=?,")
	}
	args = append(args, post.PostStatus.String())
	if post.ID == 0 {
		buffer.WriteString("`postCreateDate`,")
	} else {
		buffer.WriteString("`postCreateDate`=?,")
	}
	args = append(args, post.PostCreateDate)
	if post.ID == 0 {
		buffer.WriteString("`postModifiedDate`,")
	} else {
		buffer.WriteString("`postModifiedDate`=?,")
	}
	args = append(args, post.PostModifiedDate)
	//Cleanup the query and append where if it is an update
	query := buffer.String()
	query = strings.TrimRight(query, ",")
	if post.ID == 0 {
		vals := strings.Repeat("?,", len(args))
		vals = strings.TrimRight(vals, ",")
		query = query + ") VALUES (" + vals + ");"
	} else {
		query = query + " WHERE `idPost`=?;"
		args = append(args, post.ID)
	}

	//Lets do this thing
	log.Infoln(query)
	log.Infoln(args)
	r, _ := conn.Exec(query, args...)
	id, _ := r.LastInsertId()
	ret := int(id)
	return ret
}

//SELECTS
//Select a post by an id number
func (post *Post) SelectPostByID(ID int, site string) Post {
	var inPost Post
	inPost.ID = ID
	p := inPost.SelectPost(site)
	return p[0]
}

//Select from the post table based on the attributes set in the post object.
func (post *Post) SelectPost(site string) []Post {
	var ret []Post
	conn, _ := connectors.GetDBFromMap(site)
	var args []interface{} //arguments for variables in the data struct
	var buffer bytes.Buffer
	buffer.WriteString("SELECT idPost, postAuthor, postCreateDate, postContent, postTitle, postExcerpt, postStatus, postModifiedDate, COALESCE(postImage, '') FROM `post` WHERE ")
	if post.ID != 0 {
		buffer.WriteString("`idPost`=? AND")
		args = append(args, strconv.Itoa(post.ID))
	}
	query := buffer.String()
	query = strings.TrimRight(query, " WHERE")
	query = strings.TrimRight(query, " AND")
	query = query + ";"
	log.Infoln(query)
	rows, err := conn.Query(query, args...)
	if err != nil {
		log.Error(err)
	}
	defer rows.Close()
	for rows.Next() {
		var post Post
		var content string
		var title string
		var excerpt string
		var status string
		err := rows.Scan(&post.ID, &post.PostAuthor.ID, &post.PostCreateDate, &content, &title, &excerpt, &status, &post.PostModifiedDate, &post.PostImage)
		if err != nil {
			log.Error(err)
		}
		post.PostContent = template.HTML(html.UnescapeString(content))
		post.PostTitle = template.HTML(html.UnescapeString(title))
		post.PostExcerpt = template.HTML(html.UnescapeString(excerpt))
		post.PostStatus = PostStatusConst(PostStatusStringToInt(status))
		log.Infoln(post.ID, post.PostAuthor.ID, post.PostCreateDate, post.PostContent, post.PostTitle, post.PostExcerpt, post.PostStatus, post.PostModifiedDate, post.PostImage)
		post.PostAuthor = *post.PostAuthor.SelectUser(site)
		ret = append(ret, post)
	}
	err = rows.Err()
	if err != nil {
		log.Error(err)
	}

	return ret
}

//Select all posts in the database
func (post *Post) SelectAllPosts(site string) []Post {
	var ret []Post
	conn, _ := connectors.GetDBFromMap(site)
	log.Infoln(site)
	var buffer bytes.Buffer
	buffer.WriteString("SELECT idPost, postAuthor, postCreateDate, postContent, postTitle, postExcerpt, postStatus, postModifiedDate, COALESCE(postImage, '') FROM commonwealthcocktails.post;")
	query := buffer.String()
	log.Infoln(query)
	rows, err := conn.Query(query)
	if err != nil {
		log.Error(err)
	}
	defer rows.Close()
	for rows.Next() {
		var post Post
		var content string
		var title string
		var excerpt string
		var status string
		err := rows.Scan(&post.ID, &post.PostAuthor.ID, &post.PostCreateDate, &content, &title, &excerpt, &status, &post.PostModifiedDate, &post.PostImage)
		if err != nil {
			log.Error(err)
		}
		post.PostContent = template.HTML(html.UnescapeString(content))
		post.PostTitle = template.HTML(html.UnescapeString(title))
		post.PostExcerpt = template.HTML(html.UnescapeString(excerpt))
		post.PostStatus = PostStatusConst(PostStatusStringToInt(status))
		log.Infoln(post.ID, post.PostAuthor.ID, post.PostCreateDate, post.PostContent, post.PostTitle, post.PostExcerpt, post.PostStatus, post.PostModifiedDate, post.PostImage)
		post.PostAuthor = *post.PostAuthor.SelectUser(site)
		ret = append(ret, post)
	}
	err = rows.Err()
	if err != nil {
		log.Error(err)
	}

	return ret
}
