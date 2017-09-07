package models

import "aahframework.org/log.v0"

// Category hold category details
type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Posts       []Post `json:"posts" pg:"many2many:post_categories"`
}

// CreateCategory use to create new category
func CreateCategory(cat *Category) (*Category, error) {
	err := db.Insert(cat)
	return cat, err
}

// GetCategories use to get all categories
func GetCategories() []Category {
	var categories []Category
	err := db.Model(&categories).Column("category.*", "Posts").Order("id ASC").Select()
	if err != nil {
		log.Error(err)
	}
	return categories
}

// GetCategory use to get single category
func GetCategory(id int) *Category {
	var cat Category
	err := db.Model(&cat).Column("category.*", "Posts").Where("id = ?", id).Select()
	if err != nil {
		log.Error(err)
	}
	return &cat
}

// UpdateCategory use to create new category
func UpdateCategory(cat *Category) (*Category, error) {
	err := db.Update(cat)
	if err != nil {
		return nil, err
	}
	return GetCategory(cat.ID), nil
}

// DeleteCategory use to delete category
func DeleteCategory(id int) (int, error) {
	err := db.Delete(&Category{
		ID: id,
	})
	if err != nil {
		return 0, err
	}
	return id, nil
}
