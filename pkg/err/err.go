package pkgerr

import (
	"fmt"
)

type Status int64

const (
	InvalidJWT Status = iota
	Forbidden
	Internal
	CorruptedData
	NotFound
	BadRequest
)

type Error struct {
	Status  Status
	Message string
}

func getDetails(err error) string {
	detail := ""
	if nil != err {
		detail = fmt.Sprintf(" details: %s", err.Error())
	}
	return detail
}

func NewError(code int, status Status, message string) *Error {
	return &Error{
		Status:  status,
		Message: message,
	}
}
func NewInvalidJWTError(err error) *Error {
	return &Error{
		Status:  InvalidJWT,
		Message: fmt.Sprintf("the provided JWT is invalid : %s", err.Error()),
	}
}
func NewForbiddenError(err error) *Error {
	return &Error{
		Status:  Forbidden,
		Message: fmt.Sprintf("you don't have the required permissions : %s", err.Error()),
	}
}
func NewInternalError(err error) *Error {
	return &Error{
		Status:  Internal,
		Message: fmt.Sprintf("an internal error occured : %s", err.Error()),
	}
}
func NewCorruptedDataError(name string, id string, err error) *Error {
	return &Error{
		Status:  CorruptedData,
		Message: fmt.Sprintf("the retrieved %s row with id (%s) is corrupted.%s", name, id, getDetails(err)),
	}
}
func NewNotFoundError(name string, id string) *Error {
	return &Error{
		Status:  NotFound,
		Message: fmt.Sprintf("%s with id (%s) could not be found.", name, id),
	}
}
func NewBadRequestError(obj interface{}, name string, err error) *Error {
	return &Error{
		Status:  BadRequest,
		Message: fmt.Sprintf("%+v is not a valid %s.%s", obj, name, getDetails(err)),
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("Error %v: %s", e.Status, e.Message)
}
