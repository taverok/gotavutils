package rest

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
)

var (
	ErrGeneric          = NewError("error occurred")
	ErrNotImplemented   = NewError("not implemented")
	ErrExists           = NewError("already exists")
	ErrNoElements       = NewError("no more elements")
	ErrNotAuthorized    = NewError("not authorized").SetHTTPCode(http.StatusUnauthorized)
	ErrForbidden        = NewError("forbidden").SetHTTPCode(http.StatusForbidden)
	ErrNotFound         = NewError("not found").SetHTTPCode(http.StatusNotFound)
	ErrResourceNotFound = NewError("resource not found").SetHTTPCode(http.StatusNotFound)
	ErrInvalidParams    = NewError("invalid params")
	ErrInternal         = NewError("internal error").SetHTTPCode(http.StatusInternalServerError)
	ErrNotOwnedResource = NewError("this resource is not owned by this author")
	ErrInvalidRequest   = NewError("request format is not valid")
	ErrAuth             = NewError("wrong password or email")
)

type PublicError struct {
	Cause     error
	PublicMsg string
	HTTPCode  int
}

func NewError(publicMsg string) *PublicError {
	return &PublicError{
		PublicMsg: publicMsg,
		HTTPCode:  http.StatusBadRequest,
	}
}

func (it *PublicError) SetHTTPCode(code int) *PublicError {
	it.HTTPCode = code
	return it
}

func (it *PublicError) SetCause(cause error) *PublicError {
	it.Cause = cause
	return it
}

func (it *PublicError) Error() string {
	if it.Cause != nil {
		return it.Cause.Error()
	}

	return it.PublicMsg
}

func (it *PublicError) Unwrap() error {
	return it.Cause
}

func (it *PublicError) PublicError() string {
	return it.PublicMsg
}

func Parse(err error) *PublicError {
	var appErr *PublicError
	ok := errors.As(err, &appErr)
	if ok {
		return appErr
	}

	if errors.Is(err, sql.ErrNoRows) {
		return ErrNotFound
	}

	slog.Error(fmt.Sprintf("%v", err))
	return NewError("generic error").
		SetHTTPCode(http.StatusInternalServerError).
		SetCause(err)
}
