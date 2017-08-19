package models

import "aahframework.org/log.v0"

// Category hold post category details
type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Posts       []Post `json:"posts"  pg:"many2many:post_categories"`
}

// CreateCategory use to create new post.
func CreateCategory(name string, description string) *Category {
	cat := &Category{
		Name:        name,
		Description: description,
	}
	err := db.Insert(cat)
	if err != nil {
		log.Error(err)
	}
	return cat
}

// GetCategories use to get all posts.
func GetCategories() []Category {
	var categories []Category
	err := db.Model(&categories).Column("category.*", "Posts").Select()
	if err != nil {
		log.Error(err)
	}
	return categories
}

// GetCategory use to get single post.
func GetCategory(id int) Category {
	var cat Category
	err := db.Model(&cat).Column("category.*", "Posts").Where("id = ?", id).Select()
	if err != nil {
		log.Error(err)
	}
	return cat
}

// UpdateCategory use to create new post.
func UpdateCategory(id int, name string, description string) interface{} {
	err := db.Update(&Category{
		ID:          id,
		Name:        name,
		Description: description,
	})
	if err != nil {
		log.Error(err)
	}
	return GetCategory(id)
}

// DeleteCategory use to create new post.
func DeleteCategory(id int) interface{} {
	err := db.Delete(&Category{
		ID: id,
	})
	if err != nil {
		log.Error(err)
	}
	return id
}