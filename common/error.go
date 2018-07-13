package common

import (
	"github.com/juju/errors"
	"net/http"
)

type AppError struct {
	error
	HttpStatusCode int
}

func ErrInternalServer(err error) error {
	if err == nil {
		return nil
	}
	ae := ErrTraceCode(http.StatusInternalServerError, err).(*AppError)
	ae.error.(*errors.Err).SetLocation(1)
	return ae
}

func ErrInternalServerS(format string) error {
	ae := Err(http.StatusInternalServerError, format).(*AppError)
	ae.error.(*errors.Err).SetLocation(1)
	return ae
}

func ErrInternalServerf(format string, args ...interface{}) error {
	ae := Errf(http.StatusInternalServerError, format, args).(*AppError)
	ae.error.(*errors.Err).SetLocation(1)
	return ae
}

func ErrForbiddenf(format string, args ...interface{}) error {
	ae := Errf(http.StatusForbidden, format, args).(*AppError)
	ae.error.(*errors.Err).SetLocation(1)
	return ae
}

func ErrForbidden(format string) error {
	ae := Err(http.StatusForbidden, format).(*AppError)
	ae.error.(*errors.Err).SetLocation(1)
	return ae
}

func ErrNotFoundf(format string, args ...interface{}) error {
	ae := Errf(http.StatusNotFound, format, args).(*AppError)
	ae.error.(*errors.Err).SetLocation(1)
	return ae
}

func ErrNotFound(format string) error {
	ae := Err(http.StatusNotFound, format).(*AppError)
	ae.error.(*errors.Err).SetLocation(1)
	return ae
}

func ErrBadRequestf(format string, args ...interface{}) error {
	ae := Errf(http.StatusBadRequest, format, args).(*AppError)
	ae.error.(*errors.Err).SetLocation(1)
	return ae
}

func ErrBadRequest(format string) error {
	ae := Err(http.StatusBadRequest, format).(*AppError)
	ae.error.(*errors.Err).SetLocation(1)
	return ae
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////


func ErrAnnotatef(code int, other error, format string, args ...interface{}) error {
	cause := other
	if ae, ok := other.(*AppError); ok {
		cause = ae.error
	}
	err := errors.Annotatef(cause, format, args)
	err.(*errors.Err).SetLocation(1)
	return &AppError{err, code}
}

func ErrAnnotate(code int, other error, format string) error {
	cause := other
	if ae, ok := other.(*AppError); ok {
		cause = ae.error
	}
	err := errors.Annotate(cause, format)
	err.(*errors.Err).SetLocation(1)
	return &AppError{err, code}
}

func ErrTraceCode(code int, other error) error {
	cause := other
	if ae, ok := other.(*AppError); ok {
		cause = ae.error
	}
	err := errors.Trace(cause)
	err.(*errors.Err).SetLocation(1)
	return &AppError{err, code}
}

func ErrTrace(other error) error {
	code := http.StatusInternalServerError
	cause := other
	if ae, ok := other.(*AppError); ok {
		code = ae.HttpStatusCode
		cause = ae.error
	}
	err := errors.Trace(cause)
	err.(*errors.Err).SetLocation(1)
	return &AppError{err, code}
}

func Errf(code int, format string, args ...interface{}) error {
	err := errors.Errorf(format, args)
	err.(*errors.Err).SetLocation(1)
	return &AppError{err, code}
}

func Err(code int, format string) error {
	err := errors.New(format)
	err.(*errors.Err).SetLocation(1)
	return &AppError{err, code}
}
