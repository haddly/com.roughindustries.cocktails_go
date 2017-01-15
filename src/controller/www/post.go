//www/post.go
package www

import (
	"log"
	"model"
	"net/http"
)

type Post struct {
}

//render the page based on the name of the file provided
func (post *Post) RenderPostsTemplate(w http.ResponseWriter, tmpl string, p []model.Post) {
	t, err := parseTempFiles(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	err = t.ExecuteTemplate(w, "base", p)
	if err != nil {
		log.Fatal(err)
	}
}

func (post *Post) PostsHandler(w http.ResponseWriter, r *http.Request) {
	var p []model.Post
	p = model.GetPosts()
	post.RenderPostsTemplate(w, "articles", p)
}

func (post *Post) Init() {
	log.Println("Init in www/post.go")
	http.HandleFunc("/posts", post.PostsHandler)
}
