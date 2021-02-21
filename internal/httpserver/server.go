package httpserver

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/davidyunus/shorty/config"
	"github.com/davidyunus/shorty/internal/app"
	storages_url "github.com/davidyunus/shorty/internal/data/url"
	"github.com/davidyunus/shorty/internal/url"
)

// HTTPServer ...
type HTTPServer struct {
	App *app.App
}

func buildService(storages *app.Storages) *app.Services {
	u := url.NewService(storages.URL)
	return &app.Services{
		URL: u,
	}
}

func buildStorage(db *sql.DB) *app.Storages {
	u := storages_url.NewStorage(db)
	return &app.Storages{
		URL: u,
	}
}

func buildApp() *app.App {
	db, err := sql.Open("postgres", config.DBConnectionString())
	if err != nil {
		panic(err)
	}

	storages := buildStorage(db)
	services := buildService(storages)
	return &app.App{
		Services: services,
		Storages: storages,
	}
}

// NewServer ...
func NewServer() *HTTPServer {
	app := buildApp()
	return &HTTPServer{
		App: app,
	}
}

// Serve ...
func (hs *HTTPServer) Serve() {
	r := hs.compileRouter()

	log.Printf("Listen to port 8080. Go to : http://127.0.0.1:8080")
	srv := http.Server{Addr: ":8080", Handler: r}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen: %v\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	log.Println("Shutdown server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown: ", err)
	}
	log.Println("Server exiting")
}
