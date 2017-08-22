package controllers

import (
	"golamapi/app/models"

	"aahframework.org/aah.v0"
)

// UserController is to demostrate the REST API endpoints for User.
type UserController struct {
	AppController
}

// CreateTestUsers create new post in database and return data,
func (a *UserController) CreateTestUsers() {
	users := models.CreateTestUsers()

	a.Reply().Ok().JSON(aah.Data{
		"data": users,
	})
}
