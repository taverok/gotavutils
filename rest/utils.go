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
)

var (
	Modifier  = modifiers.New()
	Validator = validator.New()

	errSyntax    *json.SyntaxError
	errUnmarshal *json.UnmarshalTypeError
)

func Map(r io.Reader, target any) error {
	err := json.NewDecoder(r).Decode(target)
	if err != nil {
		switch {
		case errors.As(err, &errSyntax), errors.Is(err, io.ErrUnexpectedEOF), errors.Is(err, io.EOF):
			return NewError(err, err.Error())
		case errors.As(err, &errUnmarshal):
			return NewError(err, fmt.Sprintf(
				"incorrect JSON type for field %q at %d",
				errUnmarshal.Field,
				errUnmarshal.Offset,
			))
		default:
			return ErrGeneric
		}
	}

	if err = Modifier.Struct(context.Background(), target); err != nil {
		return NewError(nil, err.Error())
	}

	if err = Validator.Struct(target); err != nil {
		return NewError(nil, err.Error())
	}

	return nil
}

func GetQueryParamString(r *http.Request, name string, required bool) (string, error) {
	param := r.URL.Query().Get(name)
	if param == "" && required {
		return "", NewError(nil, fmt.Sprintf("expected query param %s not found", name))
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
