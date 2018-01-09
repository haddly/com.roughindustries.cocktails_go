// Copyright 2017 Rough Industries LLC. All rights reserved.
//model/post.db.go:package model
package model

import (
	"bytes"
	"github.com/CommonwealthCocktails/connectors"
	"github.com/golang/glog"
	"html"
	"html/template"
	"strconv"
	"strings"
	"time"
)

//CREATE, UPDATE, DELETE
//Insert a post record into the database
func (post *Post) InsertPost() int {
	//set the ID to zero to indicate an insert
	post.ID = 0
	post.PostCreateDate = time.Now()
	post.PostModifiedDate = time.Now()
	return post.processPost()
}

//Update a post record in the database based on ID
func (post *Post) UpdatePost() int {
	p := post.SelectPost()
	post.PostCreateDate = p[0].PostCreateDate
	post.PostModifiedDate = time.Now()
	return post.processPost()
}

//Process an insert or an update
func (post *Post) processPost() int {
	conn, _ := connectors.GetDB() //get db connection
	var args []interface{}        //arguments for variables in the data struct
	var buffer bytes.Buffer       //buffer for the query

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
	glog.Infoln(query)
	glog.Infoln(args)
	r, _ := conn.Exec(query, args...)
	id, _ := r.LastInsertId()
	ret := int(id)
	return ret
}

//SELECTS
//Select a post by an id number
func (post *Post) SelectPostByID(ID int) Post {
	var inPost Post
	inPost.ID = ID
	p := inPost.SelectPost()
	return p[0]
}

//Select from the post table based on the attributes set in the post object.
func (post *Post) SelectPost() []Post {
	var ret []Post
	conn, _ := connectors.GetDB()
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
	glog.Infoln(query)
	rows, err := conn.Query(query, args...)
	if err != nil {
		glog.Error(err)
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
			glog.Error(err)
		}
		post.PostContent = template.HTML(html.UnescapeString(content))
		post.PostTitle = template.HTML(html.UnescapeString(title))
		post.PostExcerpt = template.HTML(html.UnescapeString(excerpt))
		post.PostStatus = PostStatusConst(PostStatusStringToInt(status))
		glog.Infoln(post.ID, post.PostAuthor.ID, post.PostCreateDate, post.PostContent, post.PostTitle, post.PostExcerpt, post.PostStatus, post.PostModifiedDate, post.PostImage)
		post.PostAuthor = *post.PostAuthor.SelectUser()
		ret = append(ret, post)
	}
	err = rows.Err()
	if err != nil {
		glog.Error(err)
	}

	return ret
}

//Select all posts in the database
func (post *Post) SelectAllPosts() []Post {
	var ret []Post
	conn, _ := connectors.GetDB()
	var buffer bytes.Buffer
	buffer.WriteString("SELECT idPost, postAuthor, postCreateDate, postContent, postTitle, postExcerpt, postStatus, postModifiedDate, COALESCE(postImage, '') FROM commonwealthcocktails.post;")
	query := buffer.String()
	glog.Infoln(query)
	rows, err := conn.Query(query)
	if err != nil {
		glog.Error(err)
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
			glog.Error(err)
		}
		post.PostContent = template.HTML(html.UnescapeString(content))
		post.PostTitle = template.HTML(html.UnescapeString(title))
		post.PostExcerpt = template.HTML(html.UnescapeString(excerpt))
		post.PostStatus = PostStatusConst(PostStatusStringToInt(status))
		glog.Infoln(post.ID, post.PostAuthor.ID, post.PostCreateDate, post.PostContent, post.PostTitle, post.PostExcerpt, post.PostStatus, post.PostModifiedDate, post.PostImage)
		post.PostAuthor = *post.PostAuthor.SelectUser()
		ret = append(ret, post)
	}
	err = rows.Err()
	if err != nil {
		glog.Error(err)
	}

	return ret
}
