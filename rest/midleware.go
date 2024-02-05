package rest

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/taverok/gotavutils/rest/puberr"
)

func Json[T any](f func(w http.ResponseWriter, r *http.Request) (T, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := f(w, r)
		if err != nil {
			handleError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		err = json.NewEncoder(w).Encode(body)
		if err != nil {
			slog.Error(fmt.Sprintf("%v", err))
		}
	}
}

func handleError(w http.ResponseWriter, err error) {
	pubErr := puberr.Parse(err)

	w.WriteHeader(pubErr.HTTPCode)
	err = json.NewEncoder(w).Encode(pubErr)
	if err != nil {
		slog.Error(fmt.Sprintf("%v", err))
	}
}
