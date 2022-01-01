package users

import (
	"strings"

	"github.com/adrielbustos/users-book-ms/utils/restErrors"
)

func (user *User) Validate() *restErrors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return restErrors.NewBadRequest("invalid email address")
	}
	return nil
}
