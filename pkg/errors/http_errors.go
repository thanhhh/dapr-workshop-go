package errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const (
	ErrBadRequest         = "Bad request"
	ErrEmailAlreadyExists = "User with given email already exists"
	ErrNoSuchUser         = "User not found"
	ErrWrongCredentials   = "Wrong Credentials"
	ErrNotFound           = "Not Found"
	ErrUnauthorized       = "Unauthorized"
	ErrForbidden          = "Forbidden"
	ErrBadQueryParams     = "Invalid query params"
)

var (
	BadRequest          = errors.New("Bad request")
	NotFound            = errors.New("Not Found")
	Forbidden           = errors.New("Forbidden")
	InternalServerError = errors.New("Internal Server Error")
)

type RestErr interface {
	Status() int
	Error() string
	Causes() interface{}
}

type RestError struct {
	ErrStatus int         `json:"status,omitempty"`
	ErrError  string      `json:"error,omitempty"`
	ErrCauses interface{} `json:"-"`
}

func (e RestError) Error() string {
	return fmt.Sprintf("status: %d - errors: %s - causes: %v", e.ErrStatus, e.ErrError, e.ErrCauses)
}

func (e RestError) Status() int {
	return e.ErrStatus
}

func (e RestError) Causes() interface{} {
	return e.ErrCauses
}

func NewRestError(status int, err string, causes interface{}) RestErr {
	return RestError{
		ErrStatus: status,
		ErrError:  err,
		ErrCauses: causes,
	}
}

func NewRestErrorWithMessage(status int, err string, causes interface{}) RestErr {
	return RestError{
		ErrStatus: status,
		ErrError:  err,
		ErrCauses: causes,
	}
}

func NewRestErrorFromBytes(bytes []byte) (RestErr, error) {
	var apiErr RestError
	if err := json.Unmarshal(bytes, &apiErr); err != nil {
		return nil, errors.New("invalid json")
	}
	return apiErr, nil
}

func NewBadRequestError(causes interface{}) RestErr {
	return RestError{
		ErrStatus: http.StatusBadRequest,
		ErrError:  BadRequest.Error(),
		ErrCauses: causes,
	}
}

func NewNotFoundError(causes interface{}) RestErr {
	return RestError{
		ErrStatus: http.StatusNotFound,
		ErrError:  NotFound.Error(),
		ErrCauses: causes,
	}
}

func NewForbiddenError(causes interface{}) RestErr {
	return RestError{
		ErrStatus: http.StatusForbidden,
		ErrError:  Forbidden.Error(),
		ErrCauses: causes,
	}
}

func NewInternalServerError(causes interface{}) RestErr {
	result := RestError{
		ErrStatus: http.StatusInternalServerError,
		ErrError:  InternalServerError.Error(),
		ErrCauses: causes,
	}
	return result
}

// Parser of error string messages returns RestError
// func ParseErrors(err error) RestErr {
// 	switch {
// 	case errors.Is(err, sql.ErrNoRows):
// 		return NewRestError(http.StatusNotFound, NotFound.Error(), err)
// 	case errors.Is(err, context.DeadlineExceeded):
// 		return NewRestError(http.StatusRequestTimeout, RequestTimeoutError.Error(), err)
// 	case strings.Contains(err.Error(), "SQLSTATE"):
// 		return parseSqlErrors(err)
// 	case strings.Contains(err.Error(), "Field validation"):
// 		return parseValidatorError(err)
// 	case strings.Contains(err.Error(), "Unmarshal"):
// 		return NewRestError(http.StatusBadRequest, BadRequest.Error(), err)
// 	case strings.Contains(err.Error(), "UUID"):
// 		return NewRestError(http.StatusBadRequest, err.Error(), err)
// 	case strings.Contains(strings.ToLower(err.Error()), "cookie"):
// 		return NewRestError(http.StatusUnauthorized, Unauthorized.Error(), err)
// 	case strings.Contains(strings.ToLower(err.Error()), "token"):
// 		return NewRestError(http.StatusUnauthorized, Unauthorized.Error(), err)
// 	case strings.Contains(strings.ToLower(err.Error()), "bcrypt"):
// 		return NewRestError(http.StatusBadRequest, BadRequest.Error(), err)
// 	default:
// 		if restErr, ok := err.(RestErr); ok {
// 			return restErr
// 		}
// 		return NewInternalServerError(err)
// 	}
// }

// func parseSqlErrors(err error) RestErr {
// 	if strings.Contains(err.Error(), "23505") {
// 		return NewRestError(http.StatusBadRequest, ExistsEmailError.Error(), err)
// 	}

// 	return NewRestError(http.StatusBadRequest, BadRequest.Error(), err)
// }

// func parseValidatorError(err error) RestErr {
// 	if strings.Contains(err.Error(), "Password") {
// 		return NewRestError(http.StatusBadRequest, "Invalid password, min length 6", err)
// 	}

// 	if strings.Contains(err.Error(), "Email") {
// 		return NewRestError(http.StatusBadRequest, "Invalid email", err)
// 	}

// 	return NewRestError(http.StatusBadRequest, BadRequest.Error(), err)
// }

// // Error response
// func ErrorResponse(err error) (int, interface{}) {
// 	return ParseErrors(err).Status(), ParseErrors(err)
// }
