package main

import (
	_ "github.com/lib/pq"

	"github.com/davidyunus/shorty/internal/httpserver"
)

func main() {
	server := httpserver.NewServer()

	server.Serve()
}
