package controllers

import (
	"encoding/json"
	"golamapi/app/models"

	aah "aahframework.org/aah.v0"
)

// PostCatRelationController is to demostrate the REST API endpoints for Post.
type PostCatRelationController struct {
	AppController
}

// PostPostCatRelation create new post in database and return data,
func (a *PostCatRelationController) PostPostCatRelation() {
	var reqValues struct {
		PostID    int `json:"post_id"`
		PostCatID int `json:"postcat_id"`
	}
	if err := json.Unmarshal(a.Req.Payload, &reqValues); err != nil {
		a.Reply().BadRequest().JSON(aah.Data{
			"message": "bad request",
		})
		return
	}

	postid := &reqValues.PostID
	postcatid := &reqValues.PostCatID

	postcat := models.PostCatRelation(*postid, *postcatid)

	a.Reply().Ok().JSON(aah.Data{
		"data": postcat,
	})
}

// GetPostCatRels get single post
func (a *PostCatRelationController) GetPostCatRels() {
	postcats := models.GetPostCatRelations()

	a.Reply().Ok().JSON(aah.Data{
		"data": postcats,
	})
}
