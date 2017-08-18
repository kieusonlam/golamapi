package models

import (
	"time"

	"aahframework.org/log.v0"
	"github.com/go-pg/pg/orm"
)

type (
	// Post struct hold the post details
	Post struct {
		ID         int        `json:"id"`
		Title      string     `json:"title"`
		Content    string     `json:"content"`
		Categories []Category `json:"categories" pg:",many_to_many:posts_categories,fk:Post,joinFK:Category"`
		CreatedAt  time.Time  `json:"created_at"`
		UpdatedAt  time.Time  `json:"updated_at"`
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
		log.Error(err)
	}
	return post
}

// GetPosts use to get all posts.
func GetPosts() []Post {
	var posts []Post
	err := db.Model(&posts).Column("post.*", "categories").Select()
	if err != nil {
		log.Error(err)
	}
	return posts
}

// GetPost use to get single post.
func GetPost(id int) Post {
	var post Post
	err := db.Model(&post).Column("post.*", "categories").Where("id = ?", id).Select()
	if err != nil {
		log.Error(err)
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
