package url

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/davidyunus/shorty/internal/app"
	"github.com/davidyunus/shorty/internal/httpserver/handler"
	"github.com/davidyunus/shorty/internal/httpserver/response"
	"github.com/go-chi/chi"
)

// URL request param url
type URL struct {
	URL       string `json:"url"`
	Shortcode string `json:"shortcode"`
}

// CreateURL create url handler
func CreateURL() http.HandlerFunc {
	return handler.Create(func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()
		app := app.FromContext(ctx)

		var u *URL
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&u)
		if err != nil {
			return err
		}
		defer r.Body.Close()

		if u.URL == "" {
			response.WithError(w, http.StatusBadRequest, "url is not present")
			return nil
		}

		match, _ := regexp.MatchString("[0-9a-zA-Z_]{6}$", u.Shortcode)
		if match == false && u.Shortcode != "" {
			response.WithError(w, http.StatusUnprocessableEntity, "The shortcode fails to meet the following regexp: ^[0-9a-zA-Z_]{6}$.")
			return nil
		}

		if u.Shortcode != "" {
			uri, err := app.Services.URL.GetURL(ctx, u.Shortcode)
			if err != nil {
				return err
			}
			if uri != nil {
				response.WithError(w, http.StatusConflict, "The the desired shortcode is already in use. Shortcodes are case-sensitive.")
				return nil
			}
		}

		uri, err := app.Services.URL.CreateURL(ctx, u.URL, u.Shortcode)
		if err != nil {
			return err
		}

		response.JSON(w, http.StatusCreated, map[string]string{"shortcode": uri.Shortcode})
		return nil
	})
}

// GetURL get url by shortcode
func GetURL() http.HandlerFunc {
	return handler.Create(func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()
		app := app.FromContext(ctx)

		short := chi.URLParam(r, "short")
		uri, err := app.Services.URL.GetURLandAddCount(ctx, short)
		if err != nil {
			return nil
		}

		if uri == nil {
			response.WithError(w, http.StatusNotFound, "The shortcode cannot be found in the system")
			return nil
		}

		response.JSON(w, http.StatusFound, uri.URL)
		return nil
	})
}

// GetURLStats get url stats by shortcode
func GetURLStats() http.HandlerFunc {
	return handler.Create(func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()
		app := app.FromContext(ctx)
		short := chi.URLParam(r, "short")
		uri, err := app.Services.URL.GetURL(ctx, short)
		if err != nil {
			return err
		}

		if uri == nil {
			response.WithError(w, http.StatusNotFound, "The shortcode cannot be found in the system")
			return nil
		}

		response.JSON(w, http.StatusOK, uri)
		return nil
	})
}
