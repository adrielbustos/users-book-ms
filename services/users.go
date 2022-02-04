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

func UpdateUser(isPartial bool, user users.User) (*users.User, *restErrors.RestErr) {
	current, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}
	// if err := user.Validate(); err != nil {
	// 	return nil, err
	// }
	if isPartial {
		if current.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if current.LastName != "" {
			current.LastName = user.LastName
		}
		if current.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}
	err = current.Update()
	if err != nil {
		return nil, err
	}
	return current, nil
}

func DeleteUser(ui int64) *restErrors.RestErr {
	user := &users.User{
		Id: ui,
	}
	return user.Delete()
}
