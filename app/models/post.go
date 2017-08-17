package models

import (
	"time"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type (
	// Post struct hold the post details
	Post struct {
		ID         int            `json:"id"`
		Title      string         `json:"title"`
		Content    string         `json:"content"`
		Categories []PostCategory `json:"categories"`
		CreatedAt  time.Time      `json:"created_at"`
		UpdatedAt  time.Time      `json:"updated_at"`
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
	b.UpdatedAt = time.Now()
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

// UpdatePost use to create new post.
func UpdatePost(id int, title string, content string) interface{} {
	post := &Post{
		ID:      id,
		Title:   title,
		Content: content,
	}
	_, err := db.Model(post).Column("title").Column("content").Column("updated_at").Returning("*").Update()
	if err != nil {
		panic(err)
	}
	return GetPost(id)
}

// DeletePost use to create new post.
func DeletePost(id int) interface{} {
	err := db.Delete(&Post{
		ID: id,
	})
	if err != nil {
		panic(err)
	}
	return id
}
