//model/post.dat.go
package model

import (
	"time"
)

var Posts = []Post{
	Post{
		ID:          0,
		PostAuthor:  "The Commonwealth",
		PostDate:    time.Date(2017, 01, 17, 20, 34, 58, 651387237, time.UTC),
		PostContent: "First Article - Cotent Test",
		PostTitle:   "First Article - Title Test",
		PostExcerpt: "First Article - Excerpt Test",
		PostName:    "Internal Name",
		PostType:    FullPost,
		PostStatus:  Publish,
	},
}
