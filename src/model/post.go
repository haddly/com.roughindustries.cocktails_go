//model/post.go
package model

import (
	"html/template"
	"strconv"
	"time"
)

type PostType int

const (
	InnerPost = 1 + iota
	FullPost
	Blurb
)

var PostTypeStrings = [...]string{
	"InnerPost",
	"FullPost",
	"Blurb",
}

func (pt PostType) String() string { return PostTypeStrings[pt-1] }

type PostStatus int

const (
	Draft = 1 + iota
	Pending
	Private
	Publish
)

var PostStatusStrings = [...]string{
	"Draft",
	"Pending",
	"Private",
	"Publish",
}

func (ps PostStatus) String() string { return PostStatusStrings[ps-1] }

type Post struct {
	ID                         int
	PostAuthor                 string
	PostDate                   time.Time
	PostContent                template.HTML
	PostTitle                  string
	PostExcerpt                template.HTML
	PostStatus                 PostStatus //status of the post, e.g. ‘draft’, ‘pending’, ‘private’, ‘publish’. Also a great WordPress news site.
	PostName                   string
	PostType                   PostType
	PostExcerptImagePath       string
	PostExcerptImage           string
	PostExcerptImageSourceName string
	PostExcerptImageSourceLink string

	//Advertiser Info
	Advertisement Advertisement
}

func (post *Post) FormattedDate() string {
	return post.PostDate.Month().String()[:3] + " " + strconv.Itoa(post.PostDate.Day()) + " " + strconv.Itoa(post.PostDate.Year())
}

func GetPost(ID int) *Post {
	p := Posts[ID]
	return &p
}

func GetPosts() []Post {
	var p []Post
	for _, element := range Posts {
		if element.PostType == FullPost {
			p = append(p, element)
		}
	}
	return p
}
