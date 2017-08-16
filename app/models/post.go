package models

import (
	"time"

	"aahframework.org/log.v0"
	"github.com/go-pg/pg/orm"
)

type (
	// Post struct hold the post details
	Post struct {
		ID        int       `json:"id"`
		Title     string    `json:"title"`
		Content   string    `json:"content"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)

// BeforeInsert add current time to create_at
func (b *Post) BeforeInsert(db orm.DB) error {
	if b.CreatedAt.IsZero() {
		b.CreatedAt = time.Now()
	}
	if b.UpdatedAt.IsZero() {
		b.UpdatedAt = time.Now()
	}
	return nil
}

// BeforeUpdate add current time to create_at
func (b *Post) BeforeUpdate(db orm.DB) error {
	if b.CreatedAt.IsZero() {
		b.CreatedAt = time.Now()
	}
	if b.UpdatedAt.IsZero() {
		b.UpdatedAt = time.Now()
	}
	return nil
}

// CreatePost use to create new post.
func CreatePost(title string, content string) *Post {
	post := &Post{
		Title:   title,
		Content: content,
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

// UpdatePost use to create new post.
func UpdatePost(id int, title string, content string) interface{} {
	err := db.Update(&Post{
		ID:      id,
		Title:   title,
		Content: content,
	})
	if err != nil {
		log.Error(err)
	}
	return GetPost(id)
}

// DeletePost use to create new post.
func DeletePost(id int) interface{} {
	err := db.Delete(&Post{
		ID: id,
	})
	if err != nil {
		log.Error(err)
	}
	return id
}
