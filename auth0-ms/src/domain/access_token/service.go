package access_token

import "github.com/adrielbustos/users-book-ms/utils/restErrors"

type Repository interface {
	GetById(string) (*AccessToken, *restErrors.RestErr)
}

type Service interface {
	GetById(string) (*AccessToken, *restErrors.RestErr)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(id string) (*AccessToken, *restErrors.RestErr) {
	accessToken, err := s.repository.GetById(id)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}
