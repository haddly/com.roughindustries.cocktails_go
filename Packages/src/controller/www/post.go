//www/post.go
package www

import (
	"log"
	"model"
	"net/http"
	"net/url"
	"strconv"
)

type Post struct {
}

func (post *Post) Init() {
	log.Println("Init in www/post.go")
	http.HandleFunc("/post", post.PostHandler)
	http.HandleFunc("/posts", post.PostsHandler)
}

func (post *Post) RenderPostTemplate(w http.ResponseWriter, tmpl string, p *model.Post) {
	t, err := parseTempFiles(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	err = t.ExecuteTemplate(w, "base", p)
	if err != nil {
		log.Fatal(err)
	}
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

func (post *Post) PostHandler(w http.ResponseWriter, r *http.Request) {
	var p *model.Post
	u, err := url.Parse(r.URL.String())
	if err != nil {
		post.RenderPostTemplate(w, "404", p)
	}
	//log.Println("Product: " + r.URL.Path[1:])
	m, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		post.RenderPostTemplate(w, "404", p)
	}
	if len(m["ID"]) == 0 {
		post.RenderPostTemplate(w, "404", p)
	} else {
		//apply the template page info to the index page
		id, _ := strconv.Atoi(m["ID"][0])
		if len(model.Posts) <= id-1 {
			post.RenderPostTemplate(w, "404", p)
		} else {
			p = model.GetPost(id)
			post.RenderPostTemplate(w, "article", p)
		}
	}
}

func (post *Post) PostsHandler(w http.ResponseWriter, r *http.Request) {
	var p []model.Post
	p = model.GetPosts()
	post.RenderPostsTemplate(w, "articles", p)
}
