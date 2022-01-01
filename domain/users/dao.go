package users

import (
	"fmt"

	"github.com/adrielbustos/users-book-ms/utils/restErrors"
)

var (
	usersDb = make(map[int64]*User)
)

func (user *User) Get() *restErrors.RestErr {
	r := usersDb[user.Id]
	if r == nil {
		return restErrors.NewNotFound(fmt.Sprintf("user %d not found", user.Id))
	}
	// user = *r // TODO: this work?
	user.Id = r.Id
	user.FirstName = r.FirstName
	user.LastName = r.LastName
	user.Email = r.Email
	user.DateCreated = r.DateCreated
	return nil
}

func (user *User) Save() *restErrors.RestErr {
	current := usersDb[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return restErrors.NewBadRequest(fmt.Sprintf("email %s already registered", user.Email))
		}
		return restErrors.NewBadRequest(fmt.Sprintf("user %d already exists", user.Id))
	}
	usersDb[user.Id] = user
	return nil
}
