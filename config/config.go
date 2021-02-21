package config

import "os"

const (
	dBConnectionString = "DB_CONNECTION_STRING"
)

type Config struct {
}

func DBConnectionString() string {
	return getStringOrDefault(dBConnectionString, "postgres://postgres@localhost:5432/url?sslmode=disable")
}

func getStringOrDefault(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}
