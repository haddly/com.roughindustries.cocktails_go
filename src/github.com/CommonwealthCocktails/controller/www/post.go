// Copyright 2017 Rough Industries LLC. All rights reserved.
package www

import (
	"github.com/CommonwealthCocktails/model"
	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	"github.com/microcosm-cc/bluemonday"
	log "github.com/sirupsen/logrus"
	"html"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

//Post page handler which displays the standard post page.
func PostHandler(w http.ResponseWriter, r *http.Request, page *page) {
	//Process Form gets an ID if it was passed
	params := mux.Vars(r)
	if len(params["postID"]) == 0 {
		page.RenderPageTemplate(w, r, "404")
	} else {
		//apply the template page info to the index page
		id, _ := strconv.Atoi(params["postID"])
		page.Post = page.Post.SelectPostByID(id, page.View)
		page.RenderPageTemplate(w, r, "post")
	}
}

//Posts page (i.e. all the posts) request handler which
//displays the all the posts page.
func PostsHandler(w http.ResponseWriter, r *http.Request, page *page) {
	var p []model.Post
	p = page.Post.SelectAllPosts(page.View)
	totalP := len(p)
	diff := len(p) - ((page.Pagination.CurrentPage - 1) * 25)
	if diff > 25 {
		p = p[(page.Pagination.CurrentPage-1)*25 : ((page.Pagination.CurrentPage-1)*25)+25]
	} else {
		p = p[(page.Pagination.CurrentPage-1)*25 : ((page.Pagination.CurrentPage-1)*25)+diff]
	}
	page.Posts = p
	PaginationCalculate(page, page.Pagination.CurrentPage, 25, totalP, 2)
	page.SubrouteURL = "posts"
	page.RenderPageTemplate(w, r, "posts")
}

//Post Modification Form page handler which displays the Post Modification
//Form page.
func PostModFormHandler(w http.ResponseWriter, r *http.Request, page *page) {
	page.Posts = page.Post.SelectAllPosts(page.View)
	page.IsForm = true
	if page.Post.ID == 0 {
		page.Post.PostAuthor = page.UserSession.User
		//apply the template page info to the index page
		page.RenderPageTemplate(w, r, "postmodform")
	} else {
		var in model.Post
		in.ID = page.Post.ID
		out := in.SelectPost(page.View)
		page.Post = out[0]
		page.RenderPageTemplate(w, r, "postmodform")
	}
}

//Post modification form page request handler which process the post
//modification request.  This will after verifying a valid user session,
//modify the post data based on the request.
func PostModHandler(w http.ResponseWriter, r *http.Request, page *page) {
	//did we get an add, update, or something else request
	page.IsForm = true
	if page.SubmitButtonString == "add" {
		ret_id := page.Post.InsertPost(page.View)
		model.LoadMCWithPostData(page.View)
		page.Post.ID = ret_id
		outPost := page.Post.SelectPost(page.View)
		page.Post = outPost[0]
		page.Messages["postModifySuccess"] = "Post data modified successfully and memcache updated!"
		page.Posts = page.Post.SelectAllPosts(page.View)
		page.RenderPageTemplate(w, r, "postmodform")
		return
	} else if page.SubmitButtonString == "update" {
		rows_updated := page.Post.UpdatePost(page.View)
		model.LoadMCWithPostData(page.View)
		log.Infoln("Updated " + strconv.Itoa(rows_updated) + " rows")
		outPost := page.Post.SelectPost(page.View)
		page.Post = outPost[0]
		page.Messages["postModifySuccess"] = "Post data modified successfully and memcache updated!"
		page.Posts = page.Post.SelectAllPosts(page.View)
		page.RenderPageTemplate(w, r, "postmodform")
		return
	} else {
		//we only allow add and update right now
		page.Messages["postModifyFail"] = "Post data modification failed.  You tried to perform an unknown operation!"
		page.RenderPageTemplate(w, r, "postmodform")
		return
	}
}

//Parses the form and then validates the post form request and
//populates the Post struct
func ValidatePost(w http.ResponseWriter, r *http.Request, page *page) bool {
	page.Post.Errors = make(map[string]string)
	r.ParseForm() // Required if you don't call r.FormValue()
	params := mux.Vars(r)
	pUGCP := bluemonday.UGCPolicy()
	pUGCP.AllowElements("img")
	pSP := bluemonday.StrictPolicy()
	log.Infoln(r.Form)
	if len(r.Form["postID"]) > 0 && strings.TrimSpace(r.Form["postID"][0]) != "" {
		if govalidator.IsInt(r.Form["postID"][0]) {
			page.Post.ID, _ = strconv.Atoi(r.Form["postID"][0])
		} else {
			page.Post.Errors["PostID"] = "Please enter a valid post id. "
		}
	} else {
		if govalidator.IsInt(params["postID"]) {
			page.Post.ID, _ = strconv.Atoi(params["postID"])
		} else {
			page.Post.Errors["PostID"] = "Please enter a valid post id. "
		}
	}
	if len(r.Form["postTitle"]) > 0 && strings.TrimSpace(r.Form["postTitle"][0]) != "" {
		if govalidator.IsPrintableASCII(r.Form["postTitle"][0]) {
			page.Post.PostTitle = template.HTML(pSP.Sanitize(r.Form["postTitle"][0]))
		} else {
			page.Post.Errors["PostTitle"] = "Please enter a valid post title. "
		}
	}
	if len(r.Form["postStatus"]) > 0 && strings.TrimSpace(r.Form["postStatus"][0]) != "" {
		if govalidator.IsInt(r.Form["postStatus"][0]) {
			status, _ := strconv.Atoi(r.Form["postStatus"][0])
			page.Post.PostStatus = model.PostStatusConst(status)
		} else {
			page.Post.Errors["PostStatus"] = "Please enter a valid post status. "
		}
	}
	if len(r.Form["postAuthor"]) > 0 && strings.TrimSpace(r.Form["postAuthor"][0]) != "" {
		if govalidator.IsInt(r.Form["postAuthor"][0]) {
			page.Post.PostAuthor.ID, _ = strconv.Atoi(r.Form["postAuthor"][0])
		} else {
			page.Post.Errors["PostAuthor"] = "Please enter a valid post author. "
		}
	}
	if len(r.Form["postExcerpt"]) > 0 && strings.TrimSpace(r.Form["postExcerpt"][0]) != "" {
		if govalidator.IsASCII(r.Form["postExcerpt"][0]) {
			//sanitize the input, we don't want XSS
			page.Post.PostExcerpt = template.HTML(pSP.Sanitize(html.EscapeString(r.Form["postExcerpt"][0])))
		} else {
			page.Post.Errors["PostExcerpt"] = "Please enter a valid post excerpt. "
		}
	} else {
		page.Post.PostExcerpt = ""
	}
	if len(r.Form["postContent"]) > 0 && strings.TrimSpace(r.Form["postContent"][0]) != "" {
		if govalidator.IsASCII(r.Form["postContent"][0]) {
			//sanitize the input, we don't want XSS
			page.Post.PostContent = template.HTML(pSP.Sanitize(html.EscapeString(r.Form["postContent"][0])))
		} else {
			page.Post.Errors["PostContent"] = "Please enter a valid post excerpt. "
		}
	} else {
		page.Post.PostContent = ""
	}
	if len(r.Form["postImage"]) > 0 {
		page.Post.PostImage = r.Form["postImage"][0]
	}
	if len(r.Form["button"]) > 0 && strings.TrimSpace(r.Form["button"][0]) != "" {
		if govalidator.IsAlpha(r.Form["button"][0]) {
			page.SubmitButtonString = pSP.Sanitize(r.Form["button"][0])
		} else {
			page.Post.Errors["button"] = "Please click a valid button. "
		}
	}
	if len(page.Post.Errors) > 0 {
		page.Errors["postErrors"] = "You have errors in your input. "
	}
	log.Errorln(page.Post)
	return len(page.Post.Errors) == 0
}

//Checks the page post struct that required fields are filled in.
func RequiredPostMod(w http.ResponseWriter, r *http.Request, page *page) bool {
	page.Post.Errors = make(map[string]string)
	r.ParseForm() // Required if you don't call r.FormValue()
	missingRequired := false
	if r.Form["postTitle"] == nil || len(r.Form["postTitle"]) == 0 || strings.TrimSpace(r.Form["postTitle"][0]) == "" {
		page.Post.Errors["PostTitle"] = "Post title is required."
		missingRequired = true
	}
	if r.Form["postStatus"] == nil || len(r.Form["postStatus"]) == 0 {
		page.Post.Errors["PostStatus"] = "Post status is required."
		missingRequired = true
	}
	if r.Form["postAuthor"] == nil || len(r.Form["postAuthor"]) == 0 {
		page.Post.Errors["PostAuthor"] = "Post author is required."
		missingRequired = true
	}
	log.Errorln(page.Post)
	return missingRequired
}
