package users

import (
	"strings"

	"github.com/adrielbustos/users-book-ms/utils/restErrors"
)

const (
	StatusActive = "active"
)

func (user *User) Validate() *restErrors.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return restErrors.NewBadRequest("invalid email address")
	}
	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return restErrors.NewBadRequest("invalid Password")
	}
	return nil
}
