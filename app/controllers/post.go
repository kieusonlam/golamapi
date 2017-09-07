package controllers

import (
	"golamapi/app/models"

	"aahframework.org/aah.v0"
	"aahframework.org/essentials.v0"
	"aahframework.org/log.v0"
)

// PostController is to demostrate the REST API endpoints for Post
type PostController struct {
	AppController
}

// CreatePost create new post in database and return data
func (p *PostController) CreatePost(post *models.Post) {
	// Applying validation
	// Validation is upcoming feature in aah framework
	if ess.IsStrEmpty(post.Title) {
		p.Reply().BadRequest().JSON(aah.Data{
			"message": "post title is missing",
		})
		return
	}

	createdPost, err := models.CreatePost(post)
	if err != nil {
		log.Error(err)
		p.Reply().InternalServerError().JSON(aah.Data{
			"message": "error occurred while creating post",
		})
		return
	}

	p.Reply().Ok().JSON(aah.Data{
		"data": createdPost,
	})
}

// GetPosts get all post data
func (p *PostController) GetPosts() {
	posts := models.GetPosts()

	p.Reply().Ok().JSON(aah.Data{
		"data": posts,
	})
}

// GetPost get single post
func (p *PostController) GetPost(id int) {
	post := models.GetPost(id)
	if post.ID == 0 {
		p.Reply().NotFound().JSON(aah.Data{
			"message": "Post is not found!",
		})
		return
	}

	p.Reply().Ok().JSON(aah.Data{
		"data": post,
	})
}

// UpdatePost update post in database and return data
func (p *PostController) UpdatePost(id int, post *models.Post) {
	// Applying validation
	// Validation is upcoming feature in aah framework
	if id == 0 {
		p.Reply().BadRequest().JSON(aah.Data{
			"message": "bad request",
		})
		return
	}

	post.ID = id
	updatedPost := models.UpdatePost(post)

	p.Reply().Ok().JSON(aah.Data{
		"data": updatedPost,
	})
}

// DeletePost delete post by id
func (p *PostController) DeletePost(id int) {
	_, err := models.DeletePost(id)
	if err != nil {
		log.Error(err)
		p.Reply().InternalServerError().JSON(aah.Data{
			"message": "Error occurred while deleting post",
		})
		return
	}

	p.Reply().NoContent()
}
