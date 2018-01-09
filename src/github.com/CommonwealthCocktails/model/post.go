// Copyright 2017 Rough Industries LLC. All rights reserved.
//model/post.go:package model
package model

import (
	"html/template"
	"time"
)

//DATA STRUCTURES
//Post data structure
type Post struct {
	ID               int
	PostAuthor       User
	PostCreateDate   time.Time
	PostContent      template.HTML
	PostTitle        template.HTML
	PostExcerpt      template.HTML
	PostStatus       PostStatusConst
	PostModifiedDate time.Time
	PostImage        string
	Errors           map[string]string
}

//ENUMERATIONS
//The integer values for the poststatus enumeration
type PostStatusConst int

const (
	Draft = 1 + iota
	Publish
)

//The string values for the poststatus enumeration
var PostStatusStrings = [...]string{
	"Draft",
	"Publish",
}

// String returns the English name of the poststatus ("Draft", "Publish", ...).
func (ps PostStatusConst) String() string { return PostStatusStrings[ps-1] }

//Helper function to convert a post status string to it's int value.
func PostStatusStringToInt(a string) int {
	var i = 1
	for _, b := range PostStatusStrings {
		if b == a {
			return i
		}
		i++
	}
	return 0
}
