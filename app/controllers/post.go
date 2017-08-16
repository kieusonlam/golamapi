package controllers

import (
	"encoding/json"
	"strconv"

	"golamapi/app/models"

	"aahframework.org/aah.v0"
)

// PostController is to demostrate the REST API endpoints for Post.
type PostController struct {
	AppController
}

// CreatePost create new post in database and return data,
func (a *PostController) CreatePost() {
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

	title := &reqValues.Title
	content := &reqValues.Content

	post := models.CreatePost(title, content)

	a.Reply().Ok().JSON(aah.Data{
		"data": post,
	})
}

// GetPosts get all post data
func (a *PostController) GetPosts() {
	posts := models.GetPosts()

	a.Reply().Ok().JSON(aah.Data{
		"data": posts,
	})
}

// GetPost get single post
func (a *PostController) GetPost() {
	id, _ := strconv.Atoi(a.Req.PathValue("id"))
	post := models.GetPost(id)
	if post.ID != 0 {
		a.Reply().Ok().JSON(aah.Data{
			"data": post,
		})
	}
	a.Reply().NotFound().JSON(aah.Data{
		"message": "Post not found!",
	})
}
