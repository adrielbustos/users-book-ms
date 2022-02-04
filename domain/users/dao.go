package users

import (
	"fmt"

	usersdb "github.com/adrielbustos/users-book-ms/datasources/mysql/users_db"

	datesutils "github.com/adrielbustos/users-book-ms/utils/datesUtils"
	mysqlutils "github.com/adrielbustos/users-book-ms/utils/mysql_utils"
	"github.com/adrielbustos/users-book-ms/utils/restErrors"
)

const (
	queryInsertUsers = "INSERT INTO users (first_name, last_name, email, date_created) VALUES (?,?,?,?)"
	queryUpdateUsers = "UPDATE users SET first_name = ?, last_name = ?, email = ? WHERE id = ?;"
	queryGetUser     = "SELECT * FROM users"
	queryDeleteUser  = "DELETE FROM users WHERE id = ?"
)

func (user *User) Get() *restErrors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryGetUser)
	if err != nil {
		return restErrors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	result := stmt.QueryRow()
	if getErr := result.Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.DateCreated,
	); getErr != nil {
		return mysqlutils.ParseError(getErr)
	}
	return nil
}

func (user *User) Save() *restErrors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryInsertUsers)
	if err != nil {
		return restErrors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	user.DateCreated = datesutils.GetNowString()
	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if saveErr != nil {
		return mysqlutils.ParseError(saveErr)
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return restErrors.NewInternalServerError(fmt.Sprintf("erro to trying to get las insert ID: %s", err.Error()))
	}
	user.Id = userId
	return nil
}

func (user *User) Update() *restErrors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryUpdateUsers)
	if err != nil {
		return restErrors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return mysqlutils.ParseError(err)
	}
	return nil
}

func (user *User) Delete() *restErrors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryDeleteUser)
	if err != nil {
		return restErrors.NewInternalServerError(err.Error())
	}
	_, err = stmt.Exec(user.Id)
	if err != nil {
		return mysqlutils.ParseError(err)
	}
	defer stmt.Close()
	return nil
}
