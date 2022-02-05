package restErrors

import (
	"errors"
	"net/http"
)

type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewBadRequest(m string) *RestErr {
	return &RestErr{
		Message: m,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NewNotFound(m string) *RestErr {
	return &RestErr{
		Message: m,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

func NewInternalServerError(m string) *RestErr {
	return &RestErr{
		Message: m,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
}

func NewError(msg string) error {
	return errors.New(msg)
}
