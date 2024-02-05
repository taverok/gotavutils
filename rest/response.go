package rest

import (
	"encoding/json"
	"net/http"
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
