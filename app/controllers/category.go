package controllers

import (
	"encoding/json"
	"strconv"

	"golamapi/app/models"

	"aahframework.org/aah.v0"
)

// CategoryController is to demostrate the REST API endpoints for Post.
type CategoryController struct {
	AppController
}

// CreateCategory create new category in database and return data,
func (a *CategoryController) CreateCategory() {
	var reqValues struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	if err := json.Unmarshal(a.Req.Payload, &reqValues); err != nil {
		a.Reply().BadRequest().JSON(aah.Data{
			"message": "bad request",
		})
		return
	}

	name := &reqValues.Name
	description := &reqValues.Description

	cat := models.CreateCategory(*name, *description)

	a.Reply().Ok().JSON(aah.Data{
		"data": cat,
	})
}

// PostCategories get all category data
func (a *CategoryController) GetCategories() {
	cats := models.GetCategories()

	a.Reply().Ok().JSON(aah.Data{
		"data": cats,
	})
}

// GetCategory get single post
func (a *CategoryController) GetCategory() {
	id, _ := strconv.Atoi(a.Req.PathValue("id"))
	cat := models.GetCategory(id)
	if cat.ID != 0 {
		a.Reply().Ok().JSON(aah.Data{
			"data": cat,
		})
	}
	a.Reply().NotFound().JSON(aah.Data{
		"data": cat,
	})
}

// UpdateCategory update category in database and return data,
func (a *CategoryController) UpdateCategory() {
	var reqValues struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	if err := json.Unmarshal(a.Req.Payload, &reqValues); err != nil {
		a.Reply().BadRequest().JSON(aah.Data{
			"message": "bad request",
		})
		return
	}

	id, _ := strconv.Atoi(a.Req.PathValue("id"))

	title := &reqValues.Title
	content := &reqValues.Content

	post := models.UpdateCategory(id, *title, *content)

	a.Reply().Ok().JSON(aah.Data{
		"data": post,
	})
}

// DeleteCategory create new post in database and return data,
func (a *CategoryController) DeleteCategory() {
	id, _ := strconv.Atoi(a.Req.PathValue("id"))

	post := models.DeleteCategory(id)

	a.Reply().Ok().JSON(aah.Data{
		"data": post,
	})
}
