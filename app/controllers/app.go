package controllers

import (
	"time"

	"golamapi/app/models"
	"golamapi/app/security"

	"aahframework.org/aah.v0"
	"aahframework.org/essentials.v0"
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
func (a *AppController) Token(tokenReq *models.UserToken) {
	if ess.IsStrEmpty(tokenReq.Username) || ess.IsStrEmpty(tokenReq.Password) {
		a.Reply().BadRequest().JSON(aah.Data{
			"message": "bad request",
		})
		return
	}

	// get the user details by username
	user := models.FindUserByEmail(tokenReq.Username)
	if user.ID == 0 || user.IsExpried || user.IsLocked {
		a.Reply().Unauthorized().JSON(aah.Data{
			"message": "invalid credentials",
		})
		return
	}

	// validate password
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(tokenReq.Password)); err != nil {
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
