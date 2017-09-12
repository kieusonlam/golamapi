// Copyright (c) Jeevanandam M. (https://github.com/jeevatkm)
// go-aah/tutorials source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package models

import (
	"aahframework.org/log.v0"
	"golang.org/x/crypto/bcrypt"
)

// This is for demo purpose, set of users
// Typically you will be using Database, API calls, LDAP, etc to get the Authentication
// Information.

type (
	// User struct hold the user details
	User struct {
		ID          int      `json:"id"`
		FirstName   string   `json:"first_name"`
		LastName    string   `json:"last_name"`
		Email       string   `json:"email"`
		Password    []byte   `json:"password"`
		IsLocked    bool     `json:"is_locked"`
		IsExpried   bool     `json:"is_expried"`
		Roles       []string `json:"roles,omitempty"`
		Permissions []string `json:"permission,omitempty"`
	}

	// NewUser struct used to create new user
	NewUser struct {
		FirstName   string   `json:"first_name"`
		LastName    string   `json:"last_name"`
		Email       string   `json:"email"`
		Password    string   `json:"password"`
	}

	// UserToken struct used to token creation request
	UserToken struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
)

// FindUserByEmail find user by email
func FindUserByEmail(email string) User {
	var user User
	err := db.Model(&user).Where("email = ?", email).Select()
	if err != nil {
		log.Error(err)
	}
	return user
}

// CreateUser usee to create new user
func CreateUser(newuser *NewUser) (*User, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(newuser.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error(err)
	}
	user := &User{
		FirstName: 		newuser.FirstName,
		LastName: 		newuser.LastName,
		Email: 			newuser.Email,
		Password: 		hashPassword,
	}
	err = db.Insert(user)
	return user, err
}

// GetUsers use to get all posts.
func GetUsers() []User {
	var users []User
	err := db.Model(&users).Order("id ASC").Select()
	if err != nil {
		log.Error(err)
	}
	return users
}

// UpdateUser use to update user
func UpdateUser(user *User) User {
	_, err := db.Model(user).
		Column("first_name").
		Column("last_name").
		Column("password").
		Column("is_locked").
		Column("is_expried").
		Column("roles").
		Column("permission").Returning("*").Update()
	if err != nil {
		log.Error(err)
	}
	return FindUserByEmail(user.Email)
}

// DeleteUser use to delete post
func DeleteUser(email string) (string, error) {
	err := db.Delete(&User{
		Email: email,
	})
	if err != nil {
		return "0", err
	}
	return email, nil
}

// CreateTestUsers use to create test users
func CreateTestUsers() []User {
	users := []User{{
		FirstName:   "East",
		LastName:    "Corner",
		Password:    []byte("$2y$10$2A4GsJ6SmLAMvDe8XmTam.MSkKojdobBVJfIU7GiyoM.lWt.XV3H6"), // welcome123
		Email:       "user1@example.com",
		Roles:       []string{"employee", "manager"},
		Permissions: []string{"user:read,edit:reportee"},
	}, {
		FirstName:   "West",
		LastName:    "Corner",
		Password:    []byte("$2y$10$2A4GsJ6SmLAMvDe8XmTam.MSkKojdobBVJfIU7GiyoM.lWt.XV3H6"), // welcome123
		Email:       "user2@example.com",
		Roles:       []string{"employee"},
		Permissions: []string{},
	}, {
		FirstName: "South",
		LastName:  "Corner",
		Password:  []byte("$2y$10$2A4GsJ6SmLAMvDe8XmTam.MSkKojdobBVJfIU7GiyoM.lWt.XV3H6"), // welcome123
		Email:     "user3@example.com",
		IsLocked:  true,
	}, {
		FirstName:   "Admin",
		LastName:    "Corner",
		Password:    []byte("$2y$10$2A4GsJ6SmLAMvDe8XmTam.MSkKojdobBVJfIU7GiyoM.lWt.XV3H6"), // welcome123
		Email:       "admin@example.com",
		Roles:       []string{"employee", "manager", "admin"},
		Permissions: []string{"user:read,edit,delete:reportee"},
	}}
	err := db.Insert(&users)
	if err != nil {
		log.Error(err)
	}
	return users
}
