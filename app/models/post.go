package models

import (
	"github.com/go-pg/pg"
)

type (
	// Post struct hold the post details
	Post struct {
		ID         int
		Title      string
		Content    string
		Caterories string
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
		panic(err)
	}
	return post
}

// GetPosts use to get all posts.
func GetPosts() []Post {
	var posts []Post
	err := db.Model(&posts).Select()
	if err != nil {
		panic(err)
	}
	return posts
}

// GetPost use to get single post.
func GetPost(id int) interface{} {
	var post Post
	err := db.Model(&post).Where("id = ?", id).Select()
	if err == pg.ErrNoRows {
		return map[string]interface{}{
			"error": "Post not found",
		}
	} else if err != nil {
		panic(err)
	}
	return post
}
