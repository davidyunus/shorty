package app

import (
	"context"
	"net/http"

	"github.com/davidyunus/shorty/internal/url"

	storages_url "github.com/davidyunus/shorty/internal/data/url"
)

type App struct {
	Services *Services
	Storages *Storages
}

// Services collection of services
type Services struct {
	URL *url.Service
}

// Storages collection of storages
type Storages struct {
	URL storages_url.IStorage
}

type k string

const key = k("app")

func FromContext(ctx context.Context) *App {
	app, ok := ctx.Value(key).(*App)
	if !ok {
		return nil
	}
	return app
}

// InjectorApp ...
func InjectorApp(app *App) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), key, app)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
