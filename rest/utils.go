package rest

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/go-playground/mold/v4/modifiers"
	"github.com/go-playground/validator/v10"
	"github.com/taverok/gotavutils/rest/puberr"
)

var (
	modifier = modifiers.New()
	validate = validator.New()

	errSyntax    *json.SyntaxError
	errUnmarshal *json.UnmarshalTypeError
)

func Map(r io.Reader, target any) error {
	err := json.NewDecoder(r).Decode(target)
	if err != nil {
		switch {
		case errors.Is(err, io.EOF):
			return puberr.NewClientError("more content expected").SetCause(err)
		case errors.As(err, &errSyntax), errors.Is(err, io.ErrUnexpectedEOF):
			return puberr.NewClientError(err.Error()).SetCause(err)
		case errors.As(err, &errUnmarshal):
			msg := fmt.Sprintf("incorrect JSON type for field %q at %d", errUnmarshal.Field, errUnmarshal.Offset)
			return puberr.NewClientError(msg).SetCause(err)
		default:
			return puberr.ErrGeneric
		}
	}

	if err = modifier.Struct(context.Background(), target); err != nil {
		return puberr.NewClientError(err.Error())
	}

	if err = validate.Struct(target); err != nil {
		return puberr.NewClientError(err.Error())
	}

	return nil
}

func XAuthToken(r *http.Request) (string, error) {
	t := r.Header.Get("X-Auth-Token")
	if t == "" {
		return "", puberr.ErrNotAuthorized
	}

	return t, nil
}

func GetQueryParamString(r *http.Request, name string, required bool) (string, error) {
	param := r.URL.Query().Get(name)
	if param == "" && required {
		return "", puberr.NewClientError(fmt.Sprintf("expected query param %s not found", name))
	}

	return param, nil
}

func GetQueryParamInt(r *http.Request, name string, required bool) (int, error) {
	paramStr, err := GetQueryParamString(r, name, required)
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(paramStr)
}
