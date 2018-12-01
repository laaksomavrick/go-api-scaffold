package core

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse defines the shape of the default error response served by the application
type ErrorResponse struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}

// JSONError issues an ErrorResponse payload to the client
func JSONError(w http.ResponseWriter, err error, status int) {
	errResp := &ErrorResponse{
		Status: status,
		Error:  err.Error(),
	}
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(errResp)
}
