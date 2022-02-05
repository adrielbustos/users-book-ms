package services

import (
	"github.com/adrielbustos/users-book-ms/domain/users"
	cryptoutils "github.com/adrielbustos/users-book-ms/utils/crypto_utils"
	datesutils "github.com/adrielbustos/users-book-ms/utils/datesUtils"
	"github.com/adrielbustos/users-book-ms/utils/restErrors"
)

var (
	UserService userServiceInterface = &userService{}
)

type userService struct {
}

type userServiceInterface interface {
	CreateUser(users.User) (*users.User, *restErrors.RestErr)
	GetUser(int64) (*users.User, *restErrors.RestErr)
	UpdateUser(bool, users.User) (*users.User, *restErrors.RestErr)
	DeleteUser(int64) *restErrors.RestErr
	Search(string) (users.Users, *restErrors.RestErr)
}

func (s *userService) CreateUser(user users.User) (*users.User, *restErrors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.Status = users.StatusActive
	user.DateCreated = datesutils.GetDbFormat()
	user.Password = cryptoutils.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *userService) GetUser(ui int64) (*users.User, *restErrors.RestErr) {
	r := &users.User{
		Id: ui,
	}
	if err := r.Get(); err != nil {
		return nil, err
	}
	return r, nil
}

func (s *userService) UpdateUser(isPartial bool, user users.User) (*users.User, *restErrors.RestErr) {
	current, err := s.GetUser(user.Id)
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

func (s *userService) DeleteUser(ui int64) *restErrors.RestErr {
	user := &users.User{
		Id: ui,
	}
	return user.Delete()
}

func (s *userService) Search(status string) (users.Users, *restErrors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}
