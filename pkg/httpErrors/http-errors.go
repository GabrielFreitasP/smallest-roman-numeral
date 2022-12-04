package httpErrors

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

var (
	BadRequest          = errors.New("bad request")
	NotFound            = errors.New("not found")
	RequestTimeoutError = errors.New("request timeout")
	InternalServerError = errors.New("internal Server Error")
)

// Rest error interface
type RestErr interface {
	Status() int
	Error() string
	Causes() interface{}
}

// Rest error struct
type RestError struct {
	ErrStatus int         `json:"status,omitempty"`
	ErrError  string      `json:"error,omitempty"`
	ErrCauses interface{} `json:"-"`
}

// REST error formatted
func (e RestError) Error() string {
	return fmt.Sprintf("status: %d - errors: %s - causes: %v", e.ErrStatus, e.ErrError, e.ErrCauses)
}

// REST error status
func (e RestError) Status() int {
	return e.ErrStatus
}

// REST error causes
func (e RestError) Causes() interface{} {
	return e.ErrCauses
}

// REST error construct
func NewRestError(status int, err string, causes interface{}) RestErr {
	return RestError{
		ErrStatus: status,
		ErrError:  err,
		ErrCauses: causes,
	}
}

// Parser of error string messages returns RestError
func ParseErrors(err error) RestErr {
	switch {
	case strings.Contains(err.Error(), "unmarshal"):
		return NewRestError(http.StatusBadRequest, BadRequest.Error(), err)
	case strings.Contains(err.Error(), "EOF"):
		return NewRestError(http.StatusNotFound, NotFound.Error(), err)
	case errors.Is(err, context.DeadlineExceeded):
		return NewRestError(http.StatusRequestTimeout, RequestTimeoutError.Error(), err)
	default:
		return NewRestError(http.StatusInternalServerError, InternalServerError.Error(), err)
	}
}

// Error response
func ErrorResponse(err error) (int, interface{}) {
	restErr := ParseErrors(err)
	return restErr.Status(), restErr
}
