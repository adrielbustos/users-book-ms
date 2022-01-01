package services

import (
	"github.com/adrielbustos/users-book-ms/domain/users"
	"github.com/adrielbustos/users-book-ms/utils/restErrors"
)

func CreateUser(user users.User) (*users.User, *restErrors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUser(ui int64) (*users.User, *restErrors.RestErr) {
	r := &users.User{
		Id: ui,
	}
	if err := r.Get(); err != nil {
		return nil, err
	}
	return r, nil
}
