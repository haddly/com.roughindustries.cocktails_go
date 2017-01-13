//model/post.go
package model

import (
	"html/template"
	"time"
)

type PostType int

const (
	AboutPost = 1 + iota
	ArticlePost
	Blurb
)

var PostTypeStrings = [...]string{
	"AboutPost",
	"ArticlePost",
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
	ID           int
	post_author  string
	post_date    time.Time
	post_content template.HTML
	post_title   string
	post_excerpt template.HTML
	post_status  PostStatus //status of the post, e.g. ‘draft’, ‘pending’, ‘private’, ‘publish’. Also a great WordPress news site.
	post_name    string
	post_type    PostType
}
