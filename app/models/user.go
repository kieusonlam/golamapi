// Copyright (c) Jeevanandam M. (https://github.com/jeevatkm)
// go-aah/tutorials source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package models

import "aahframework.org/log.v0"

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
		Password    []byte   `json:"-"`
		IsLocked    bool     `json:"is_locked"`
		IsExpried   bool     `json:"is_expried"`
		Roles       []string `json:"roles,omitempty"`
		Permissions []string `json:"permission,omitempty"`
	}
)

// FindUserByEmail use to get single post.
func FindUserByEmail(email string) User {
	var user User
	err := db.Model(&user).Where("email = ?", email).Select()
	if err != nil {
		log.Error(err)
	}
	return user
}

// CreateTestUsers use to create new post.
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
