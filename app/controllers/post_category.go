package controllers

import (
	"encoding/json"
	"golamapi/app/models"
	"strconv"

	aah "aahframework.org/aah.v0"
)

// PostCategoryController is to demostrate the REST API endpoints for Post.
type PostCategoryController struct {
	AppController
}

// PostPostCatRelation create new post in database and return data,
func (a *PostCategoryController) PostPostCatRelation() {
	var reqValues struct {
		PostID     int `json:"post_id"`
		CategoryID int `json:"category_id"`
	}
	if err := json.Unmarshal(a.Req.Payload, &reqValues); err != nil {
		a.Reply().BadRequest().JSON(aah.Data{
			"message": "bad request",
		})
		return
	}

	postid := &reqValues.PostID
	catid := &reqValues.CategoryID

	postcat := models.PostCatRelation(*postid, *catid)

	a.Reply().Ok().JSON(aah.Data{
		"data": postcat,
	})
}

// GetPostCatRels get single post
func (a *PostCategoryController) GetPostCatRels() {
	postcats := models.GetPostCatRelations()

	a.Reply().Ok().JSON(aah.Data{
		"data": postcats,
	})
}

// DeletePostCatRel create new post in database and return data,
func (a *PostCategoryController) DeletePostCatRel() {
	id, _ := strconv.Atoi(a.Req.PathValue("id"))

	post := models.DelPostCatRel(id)

	a.Reply().Ok().JSON(aah.Data{
		"data": post,
	})
}
