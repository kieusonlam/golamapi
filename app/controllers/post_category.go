package controllers

import (
	"golamapi/app/models"

	"aahframework.org/aah.v0"
	"aahframework.org/log.v0"
)

// PostCategoryController is to demostrate the REST API endpoints for Post.
type PostCategoryController struct {
	AppController
}

// PostPostCatRelation create new post in database and return data,
func (p *PostCategoryController) PostPostCatRelation(postCategory *models.PostCategory) {
	postcat, err := models.CreatePostCatRelation(postCategory)
	if err != nil {
		log.Error(err)
		p.Reply().InternalServerError().JSON(aah.Data{
			"message": "Error occurred while creating post and category",
		})
		return
	}

	p.Reply().Ok().JSON(aah.Data{
		"data": postcat,
	})
}

// GetPostCatRels get single post
func (p *PostCategoryController) GetPostCatRels() {
	postcats := models.GetPostCatRelations()

	p.Reply().Ok().JSON(aah.Data{
		"data": postcats,
	})
}

// DeletePostCatRel create new post in database and return data,
func (p *PostCategoryController) DeletePostCatRel(id int) {
	_, err := models.DelPostCatRel(id)
	if err != nil {
		log.Error(err)
		p.Reply().InternalServerError().JSON(aah.Data{
			"message": "Error occurred while deleting relation of post and category",
		})
		return
	}

	p.Reply().NoContent()
}
