// Copyright (c) Jeevanandam M. (https://github.com/jeevatkm)
// go-aah/tutorials source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package controllers

import (
	"encoding/json"
	"time"

	"golamapi/app/models"
	"golamapi/app/security"

	"aahframework.org/aah.v0"
	"aahframework.org/log.v0"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// AppController struct application controller
type AppController struct {
	*aah.Context
}

// Index method is application root API endpoint.
func (a *AppController) Index() {
	a.Reply().Ok().JSON(models.Greet{
		Message: "Welcome to aah framework - REST API JWT Auth using Generic Auth scheme",
	})
}

// Token method validates the given username and password the generates the
// JWT token.
func (a *AppController) Token() {
	// In upcoming release auto binding feature is coming :)
	// username := a.Req.FormValue("username")
	// password := a.Req.FormValue("password")

	// If you want JSON payload instead of form-data submission
	// uncomment below lines and comment above two lines
	var reqValues struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.Unmarshal(a.Req.Payload, &reqValues); err != nil {
		a.Reply().BadRequest().JSON(aah.Data{
			"message": "bad request",
		})
		return
	}

	username := reqValues.Username
	password := reqValues.Password

	// get the user details by username
	user := models.FindUserByEmail(username)
	if user.ID == 0 || user.IsExpried || user.IsLocked {
		a.Reply().Unauthorized().JSON(aah.Data{
			"message": "invalid credentials",
		})
		return
	}

	// validate password
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(password)); err != nil {
		a.Reply().Unauthorized().JSON(aah.Data{
			"message": "invalid credentials",
		})
		return
	}

	// Generate JWT token
	token := security.CreateJWTToken()

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	signedToken, err := token.SignedString(security.JWTSigningKey)
	if err != nil {
		log.Error(err)
		a.Reply().InternalServerError().JSON(aah.Data{
			"message": "Whoops! something went wrong...",
		})
		return
	}

	// everything went good, respond token
	a.Reply().Ok().JSON(aah.Data{
		"token": signedToken,
	})
}
