package controllers

import (
	"golamapi/app/models"

	"aahframework.org/aah.v0"
	"aahframework.org/essentials.v0"
	"aahframework.org/log.v0"
)

// UserController is to demostrate the REST API endpoints for User
type UserController struct {
	AppController
}

// CreateTestUsers create test users
func (u *UserController) CreateTestUsers() {
	users := models.CreateTestUsers()

	u.Reply().Ok().JSON(aah.Data{
		"data": users,
	})
}

// CreateUser create new user
func (u *UserController) CreateUser(newuser *models.NewUser) {
	// Applying validation
	// Validation is upcoming feature in aah framework
	if ess.IsStrEmpty(newuser.Email) {
		u.Reply().BadRequest().JSON(aah.Data{
			"message": "User email is missing",
		})
		return
	}

	existedUser := models.FindUserByEmail(newuser.Email)
	if existedUser.ID != 0 {
		u.Reply().BadRequest().JSON(aah.Data{
			"message": "This email adready exists!",
		})
		return
	}

	createdUser, err := models.CreateUser(newuser)
	if err != nil {
		log.Error(err)
		u.Reply().InternalServerError().JSON(aah.Data{
			"message": "Error occurred while creating user",
		})
		return
	}

	u.Reply().Ok().JSON(aah.Data{
		"data": createdUser,
	})
}

// GetUsers get all users data
func (u *UserController) GetUsers() {
	users := models.GetUsers()

	u.Reply().Ok().JSON(aah.Data{
		"data": users,
	})
}

// GetUser get user by email
func (u *UserController) GetUser(email string) {
	user := models.FindUserByEmail(email)

	if user.ID == 0 {
		u.Reply().NotFound().JSON(aah.Data{
			"message": "User is not found!",
		})
		return
	}

	u.Reply().Ok().JSON(aah.Data{
		"data": user,
	})
}

// UpdateUser update user in database and return data
func (u *UserController) UpdateUser(email string, user *models.User) {
	if email == "" {
		u.Reply().BadRequest().JSON(aah.Data{
			"message": "bad request",
		})
		return
	}

	user.Email = email
	updatedUser := models.UpdateUser(user)

	u.Reply().Ok().JSON(aah.Data{
		"data": updatedUser,
	})
}

// DeleteUser delete user by email
func (u *UserController) DeleteUser(email string) {
	_, err := models.DeleteUser(email)
	if err != nil {
		log.Error(err)
		u.Reply().InternalServerError().JSON(aah.Data{
			"message": "Error occurred while deleting post",
		})
		return
	}

	u.Reply().NoContent()
}
