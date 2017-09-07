package models

import "aahframework.org/log.v0"

// Category hold post category details
type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Posts       []Post `json:"posts" pg:"many2many:post_categories"`
}

// CreateCategory use to create new post.
func CreateCategory(cat *Category) (*Category, error) {
	err := db.Insert(cat)
	return cat, err
}

// GetCategories use to get all posts.
func GetCategories() []Category {
	var categories []Category
	err := db.Model(&categories).Column("category.*", "Posts").Order("id ASC").Select()
	if err != nil {
		log.Error(err)
	}
	return categories
}

// GetCategory use to get single post.
func GetCategory(id int) *Category {
	var cat Category
	err := db.Model(&cat).Column("category.*", "Posts").Where("id = ?", id).Select()
	if err != nil {
		log.Error(err)
	}
	return &cat
}

// UpdateCategory use to create new post.
func UpdateCategory(cat *Category) (*Category, error) {
	err := db.Update(cat)
	if err != nil {
		return nil, err
	}
	return GetCategory(cat.ID), nil
}

// DeleteCategory use to create new post.
func DeleteCategory(id int) (int, error) {
	err := db.Delete(&Category{
		ID: id,
	})
	if err != nil {
		return 0, err
	}
	return id, nil
}
