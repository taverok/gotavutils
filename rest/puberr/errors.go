package puberr

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
)

var (
	ErrGeneric          = NewClientError("error occurred")
	ErrNotImplemented   = NewClientError("not implemented")
	ErrExists           = NewClientError("already exists")
	ErrNoElements       = NewClientError("no more elements")
	ErrNotAuthorized    = NewClientError("not authorized").SetHTTPCode(http.StatusUnauthorized)
	ErrForbidden        = NewClientError("forbidden").SetHTTPCode(http.StatusForbidden)
	ErrNotFound         = NewClientError("not found").SetHTTPCode(http.StatusNotFound)
	ErrResourceNotFound = NewClientError("resource not found").SetHTTPCode(http.StatusNotFound)
	ErrInvalidParams    = NewClientError("invalid params")
	ErrInternal         = NewClientError("internal error").SetHTTPCode(http.StatusInternalServerError)
	ErrNotOwnedResource = NewClientError("this resource is not owned by this author")
	ErrInvalidRequest   = NewClientError("request format is not valid")
	ErrAuth             = NewClientError("wrong password or email")
)

type Error struct {
	Cause     error
	PublicMsg string
	HTTPCode  int
}

func NewClientError(msg string) *Error {
	return &Error{
		PublicMsg: msg,
		HTTPCode:  http.StatusBadRequest,
	}
}

func (it Error) SetHTTPCode(code int) Error {
	it.HTTPCode = code
	return it
}

func (it Error) SetCause(cause error) Error {
	it.Cause = cause
	return it
}

func (it Error) Error() string {
	if it.Cause != nil {
		return it.Cause.Error()
	}

	return it.PublicMsg
}

func (it Error) Unwrap() error {
	return it.Cause
}

func (it Error) PublicError() string {
	return it.PublicMsg
}

func Parse(err error) Error {
	var appErr *Error
	ok := errors.As(err, &appErr)
	if ok {
		return *appErr
	}

	if errors.Is(err, sql.ErrNoRows) {
		return ErrNotFound
	}

	slog.Error(fmt.Sprintf("%v", err))
	return NewClientError("generic error").
		SetHTTPCode(http.StatusInternalServerError).
		SetCause(err)
}
