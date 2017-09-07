package controllers

import (
	"golamapi/app/models"

	"aahframework.org/aah.v0"
	"aahframework.org/essentials.v0"
	"aahframework.org/log.v0"
	"golang.org/x/crypto/bcrypt"
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

// CreateUser create new User
func (u *UserController) CreateUser(user *models.NewUser) {
	// Applying validation
	// Validation is upcoming feature in aah framework
	if ess.IsStrEmpty(user.Email) {
		u.Reply().BadRequest().JSON(aah.Data{
			"message": "User email is missing",
		})
		return
	}

	existedUser := models.FindUserByEmail(user.Email)
	if existedUser.ID != 0 {
		u.Reply().BadRequest().JSON(aah.Data{
			"message": "This email adready exists!",
		})
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error(err)
	}
	user.Password = []byte(hashPassword)

	createdUser, err := models.CreateUser(user)
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
