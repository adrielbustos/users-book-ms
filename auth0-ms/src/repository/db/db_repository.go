package db

import "github.com/adrielbustos/users-book-ms/utils/restErrors"

type DbRepository interface {
	GetById(string) (*AccessToken, *restErrors.RestErr)
}

type dbRepository struct {
}

func NewRepository() DbRepository {
	return &dbRepository{}
}

func (r *dbRepository) GetById(id string) (*AccessToken, *restErrors.RestErr) {
	session, err := cassandra.GetSession()
	if err != nil {
		panic(err)
	}
	return nil, &restErrors.NewInternalServerError("not implemented")
}
