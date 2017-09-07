package models

import (
	"time"

	"aahframework.org/log.v0"
	"github.com/go-pg/pg/orm"
)

// Post struct hold the post details
type Post struct {
	ID         int        `json:"id"`
	Title      string     `json:"title"`
	Content    string     `json:"content"`
	Categories []Category `json:"categories" pg:"many2many:post_categories"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

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
func CreatePost(post *Post) (*Post, error) {
	err := db.Insert(post)
	return post, err
}

// GetPosts use to get all posts.
func GetPosts() []Post {
	var posts []Post
	err := db.Model(&posts).Column("post.*", "Categories").Order("id ASC").Select()
	if err != nil {
		log.Error(err)
	}
	return posts
}

// GetPost use to get single post.
func GetPost(id int) *Post {
	var post Post
	err := db.Model(&post).Column("post.*", "Categories").Where("id = ?", id).Select()
	if err != nil {
		log.Error(err)
	}
	return &post
}

// UpdatePost use to update post.
func UpdatePost(post *Post) *Post {
	_, err := db.Model(post).Column("title").Column("content").Column("updated_at").Returning("*").Update()
	if err != nil {
		log.Error(err)
	}
	return GetPost(post.ID)
}

// DeletePost use to delete post.
func DeletePost(id int) (int, error) {
	err := db.Delete(&Post{
		ID: id,
	})
	if err != nil {
		return 0, err
	}
	return id, nil
}
