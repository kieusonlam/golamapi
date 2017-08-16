package models

import "aahframework.org/log.v0"

type (
	// Post struct hold the post details
	Post struct {
		ID         int    `json:"id"`
		Title      string `json:"title"`
		Content    string `json:"content"`
		Caterories string `json:"categories"`
	}
)

// CreatePost use to create new post.
func CreatePost(title *string, content *string) *Post {
	post := &Post{
		Title:      *title,
		Content:    *content,
		Caterories: "1",
	}
	err := db.Insert(post)
	if err != nil {
		log.Error(err)
	}
	return post
}

// GetPosts use to get all posts.
func GetPosts() []Post {
	var posts []Post
	err := db.Model(&posts).Select()
	if err != nil {
		log.Error(err)
	}
	return posts
}

// GetPost use to get single post.
func GetPost(id int) Post {
	var post Post
	err := db.Model(&post).Where("id = ?", id).Select()
	if err != nil {
		log.Error(err)
	}
	return post
}
