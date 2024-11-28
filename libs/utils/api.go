package utils

import (
	"encoding/json"
	"net/http"
)

type ApiError struct {
	Message string `json:"message"`
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func ErrorResponse(w http.ResponseWriter, _ *http.Request, status int, err string) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	
	return json.NewEncoder(w).Encode(ApiError{
		Message: err,
	})
}