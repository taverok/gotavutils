package rest

import (
	"encoding/json"
	"net/http"

	"github.com/taverok/gotavutils/rest/puberr"
)

func OK(w http.ResponseWriter, body any) error {
	w.WriteHeader(http.StatusOK)

	jsonData, err := json.Marshal(&body)
	if err != nil {
		return err
	}

	if _, err := w.Write(jsonData); err != nil {
		return err
	}

	return nil
}

type AcceptedResponse struct {
	Success bool `json:"success"`
}

func Accepted(w http.ResponseWriter) error {
	w.WriteHeader(http.StatusAccepted)
	jsonData, err := json.Marshal(&AcceptedResponse{true})
	if err != nil {
		return err
	}

	if _, err = w.Write(jsonData); err != nil {
		return err
	}

	return err
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func HandleError(w http.ResponseWriter, err error) {
	pubError := puberr.Parse(err)
	w.WriteHeader(pubError.HTTPCode)
	_ = json.NewEncoder(w).Encode(ErrorResponse{Error: pubError.Error()})
}
