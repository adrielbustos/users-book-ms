package users

import (
	"fmt"

	usersdb "github.com/adrielbustos/users-book-ms/datasources/mysql/users_db"
	"github.com/adrielbustos/users-book-ms/logger"

	"github.com/adrielbustos/users-book-ms/utils/restErrors"
)

const (
	queryInsertUsers      = "INSERT INTO users (first_name, last_name, email, date_created) VALUES (?,?,?,?,?,?)"
	queryUpdateUsers      = "UPDATE users SET first_name = ?, last_name = ?, email = ? WHERE id = ?;"
	queryGetUser          = "SELECT * FROM users"
	queryDeleteUser       = "DELETE FROM users WHERE id = ?"
	queryFindUserByStatus = "SELECT * FROM users WHERE status = ?"
)

func (user *User) Get() *restErrors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("Error when trying to prepare get user", err)
		return restErrors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	result := stmt.QueryRow()
	if getErr := result.Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.DateCreated,
		&user.Status,
	); getErr != nil {
		logger.Error("Error when trying to prepare get user by id", err)
		return restErrors.NewInternalServerError("database error")
		// return mysqlutils.ParseError(getErr)
	}
	return nil
}

func (user *User) Save() *restErrors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryInsertUsers)
	if err != nil {
		logger.Error("Error when trying to prepare save user", err)
		return restErrors.NewInternalServerError("db error")
	}
	defer stmt.Close()
	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.DateCreated, user.Password)
	if saveErr != nil {
		logger.Error("Error on save user", saveErr)
		// return mysqlutils.ParseError(saveErr)
		return restErrors.NewInternalServerError("db error")
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error(fmt.Sprintf("erro to trying to get las insert ID: %s", err.Error()), err)
		return restErrors.NewInternalServerError("db error")
		// return restErrors.NewInternalServerError(fmt.Sprintf("erro to trying to get las insert ID: %s", err.Error()))
	}
	user.Id = userId
	return nil
}

func (user *User) Update() *restErrors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryUpdateUsers)
	if err != nil {
		logger.Error("Error when trying to prepare update user", err)
		return restErrors.NewInternalServerError("db error")
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		logger.Error("Erroron update user", err)
		return restErrors.NewInternalServerError("db error")
		// return mysqlutils.ParseError(err)
	}
	return nil
}

func (user *User) Delete() *restErrors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("Erroron prepare delete user", err)
		return restErrors.NewInternalServerError("db error")
	}
	_, err = stmt.Exec(user.Id)
	if err != nil {
		logger.Error("Error on delete user", err)
		return restErrors.NewInternalServerError("db error")
		// return mysqlutils.ParseError(err)
	}
	defer stmt.Close()
	return nil
}

func (user *User) FindByStatus(status string) ([]User, *restErrors.RestErr) {
	stmt, err := usersdb.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		logger.Error("Error on prepare find user", err)
		return nil, restErrors.NewInternalServerError("db error")
		// return nil, restErrors.NewInternalServerError(err.Error())
	}
	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("Error on execute find user", err)
		return nil, restErrors.NewInternalServerError("db error")
	}
	defer rows.Close()
	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status, &user.Password); err != nil {
			logger.Error("Error on parse find user", err)
			return nil, restErrors.NewInternalServerError("db error")
			// return nil, mysqlutils.ParseError(err)
		}
		results = append(results, user)
	}
	if len(results) == 0 {
		return nil, restErrors.NewNotFound("no users matching by status: " + status)
	}
	return results, nil
}
