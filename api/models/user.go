package models

import "github.com/Kamva/mgm/v2"

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Email            string `json:"email" bson:"email"`
	Username         string `json:"username" bson:"username"`
	Password         string `json:"password" bson:"password"`
}

func Register(email, username, password string) *User {
	return &User{
		Email:    email,
		Username: username,
		Password: password,
	}
}
