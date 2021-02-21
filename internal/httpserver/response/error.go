package response

import (
	"encoding/json"
	"net/http"
)

//ErrorResponse represents error message
//swagger:model
type ErrorResponse struct {
	Code    string      `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// WithError emits proper error response
func WithError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	var errorCode string
	switch status {
	case http.StatusBadRequest:
		errorCode = "Bad Request"
	case http.StatusConflict:
		errorCode = "Conflict"
	case http.StatusUnprocessableEntity:
		errorCode = "Unprocessable Entity"
	case http.StatusNotFound:
		errorCode = "Not Found"
	default:
		errorCode = "Service Error"
	}

	json.NewEncoder(w).Encode(ErrorResponse{
		Code:    errorCode,
		Message: message,
	})
}
