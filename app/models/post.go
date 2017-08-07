package models

import "github.com/jinzhu/gorm"

type (
	// Posts struct hold the post details
	Posts struct {
		gorm.Model
		Title   string `gorm:"not null" json:"title"`
		Content string `json:"content"`
	}
	// Users struct hold the users details
	Users struct {
		gorm.Model
		Firstname string `gorm:"not null" form:"firstname" json:"firstname"`
		Lastname  string `gorm:"not null" form:"lastname" json:"lastname"`
	}
)

func initDbUser() {
	// Creating the table
	if !db.HasTable(&Posts{}) {
		db.CreateTable(&Posts{})
	}
}

// CreateNewPost ust to create new post.
func CreateNewPost(title *string, content *string) Posts {
	var post Posts
	post.Title = *title
	post.Content = *content
	db.Create(&post)
	return post
}

// GetPostsData use to get all posts.
func GetPostsData() []Posts {
	var posts []Posts
	db.Find(&posts)
	return posts
}
