package mysqlutils

import (
	"strings"

	"github.com/adrielbustos/users-book-ms/utils/restErrors"
	"github.com/go-sql-driver/mysql"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *restErrors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return restErrors.NewNotFound("not recod matchin given id")
		}
		return restErrors.NewInternalServerError("error parsing mysql response")
	}
	switch sqlErr.Number {
	case 1062:
		return restErrors.NewBadRequest("invalid data")
	}
	return restErrors.NewInternalServerError("error processing request")
}
