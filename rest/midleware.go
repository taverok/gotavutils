package rest

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
)

type handler func(w http.ResponseWriter, r *http.Request) error

// Do is middleware that brings common app logic, headers, error handling
func Do(h handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		err := h(w, r)
		if err == nil {
			return
		}

		HandleError(w, err)
	}
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func HandleError(w http.ResponseWriter, err error) {
	pubError := toPublicError(err)
	w.WriteHeader(pubError.HTTPCode)
	_ = json.NewEncoder(w).Encode(ErrorResponse{Error: pubError.Error()})

	slog.Error(pubError.PublicError(), "err", err)
}

func toPublicError(err error) *PublicError {
	appErr, ok := err.(*PublicError)
	if ok {
		return appErr
	}

	if errors.Is(err, sql.ErrNoRows) {
		return ErrNotFound
	}

	return ErrGeneric
}
