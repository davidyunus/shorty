package response

import (
	"encoding/json"
	"net/http"
)

// JSON writes json http response
func JSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	var code string
	switch status {
	case http.StatusFound:
		code = "Found"
	case http.StatusCreated:
		code = "Created"
	default:
		code = "OK"
	}

	json.NewEncoder(w).Encode(ErrorResponse{
		Code: code,
		Data: data,
	})
}
