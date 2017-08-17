package models

import (
	"github.com/go-pg/pg"
)

type (
	// PostCategory hold post category details
	PostCategory struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Posts       []Post `pg:"many_to_many:posts_categories"`
	}
)

// CreatePostCategory use to create new post.
func CreatePostCategory(name string, description string) *PostCategory {
	postcat := &PostCategory{
		Name:        name,
		Description: description,
	}
	err := db.Insert(postcat)
	if err != nil {
		panic(err)
	}
	return postcat
}

// GetPostCategories use to get all posts.
func GetPostCategories() []PostCategory {
	var postcats []PostCategory
	err := db.Model(&postcats).Select()
	if err != nil {
		panic(err)
	}
	return postcats
}

// GetPostCategory use to get single post.
func GetPostCategory(id int) interface{} {
	var postcat PostCategory
	err := db.Model(&postcat).Where("id = ?", id).Select()
	if err == pg.ErrNoRows {
		return map[string]interface{}{
			"error": "Post not found",
		}
	} else if err != nil {
		panic(err)
	}
	return postcat
}

// UpdatePostCategory use to create new post.
func UpdatePostCategory(id int, name string, description string) interface{} {
	err := db.Update(&PostCategory{
		ID:          id,
		Name:        name,
		Description: description,
	})
	if err != nil {
		panic(err)
	}
	return GetPostCategory(id)
}

// DeletePostCategory use to create new post.
func DeletePostCategory(id int) interface{} {
	err := db.Delete(&PostCategory{
		ID: id,
	})
	if err != nil {
		panic(err)
	}
	return id
}
