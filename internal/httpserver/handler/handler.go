package handler

import (
	"net/http"

	"github.com/davidyunus/shorty/internal/app"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

// Create returns http handler function
func Create(fn HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-r.Context().Done():
			return
		default:
			err := fn(w, r)
			if err != nil {
				_ = app.FromContext(r.Context())

			}
			return
		}
	}
}
