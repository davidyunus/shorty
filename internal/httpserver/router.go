package httpserver

import (
	"github.com/go-chi/chi"
	"github.com/rs/cors"

	"github.com/davidyunus/shorty/internal/app"
	"github.com/davidyunus/shorty/internal/httpserver/handler/url"
)

func (hs *HTTPServer) compileRouter() chi.Router {
	r := chi.NewRouter()

	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	r.Use(cors.Handler)
	r.Use(app.InjectorApp(hs.App))

	r.Route("/v1", func(r chi.Router) {
		r.Get("/{short}", url.GetURL())
		r.Get("/{short}/stats", url.GetURLStats())
		r.Post("/{short}", url.CreateURL())
	})
	return r
}
