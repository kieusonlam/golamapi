package controllers

import (
	"golamapi/app/models"

	"aahframework.org/aah.v0"
	"aahframework.org/essentials.v0"
	"aahframework.org/log.v0"
)

// CategoryController is to demostrate the REST API endpoints for Category
type CategoryController struct {
	AppController
}

// CreateCategory create new category in database and return data
func (c *CategoryController) CreateCategory(category *models.Category) {
	// Applying validation
	// Validation is upcoming feature in aah framework
	if ess.IsStrEmpty(category.Name) {
		c.Reply().BadRequest().JSON(aah.Data{
			"message": "Category name is missing",
		})
		return
	}

	cat, err := models.CreateCategory(category)
	if err != nil {
		log.Error(err)
		c.Reply().InternalServerError().JSON(aah.Data{
			"message": "Error occurred while creating category",
		})
		return
	}

	c.Reply().Ok().JSON(aah.Data{
		"data": cat,
	})
}

// GetCategories get all categories data
func (c *CategoryController) GetCategories() {
	cats := models.GetCategories()

	c.Reply().Ok().JSON(aah.Data{
		"data": cats,
	})
}

// GetCategory get single category
func (c *CategoryController) GetCategory(id int) {
	cat := models.GetCategory(id)
	if cat.ID == 0 {
		c.Reply().NotFound().JSON(aah.Data{
			"message": "Category is not found!",
		})
		return
	}

	c.Reply().Ok().JSON(aah.Data{
		"data": cat,
	})
}

// UpdateCategory update category in database and return data
func (c *CategoryController) UpdateCategory(id int, category *models.Category) {
	category.ID = id
	cat, err := models.UpdateCategory(category)
	if err != nil {
		log.Error(err)
		c.Reply().InternalServerError().JSON(aah.Data{
			"message": "Error occurred while updating category",
		})
		return
	}

	c.Reply().Ok().JSON(aah.Data{
		"data": cat,
	})
}

// DeleteCategory delete category by id
func (c *CategoryController) DeleteCategory(id int) {
	_, err := models.DeleteCategory(id)
	if err != nil {
		log.Error(err)
		c.Reply().InternalServerError().JSON(aah.Data{
			"message": "Error occurred while deleting category",
		})
		return
	}

	c.Reply().NoContent()
}
