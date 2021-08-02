package config

import (
	"os"

	"github.com/joho/godotenv"
)

func get(variable string, defaultVar string) string {
	godotenv.Load()
	envar := os.Getenv(variable)
	if envar == "" {
		return defaultVar
	}
	return envar
}

var (
	DBUser   = get("SCHOTORI_DATABASE_USER", "root")
	DBHost   = get("SCHOTORI_DATABASE_HOST", "localhost")
	DBPasswd = get("SCHOTORI_DATABAE_PASSWORD", "root")
	DBName   = get("SCHOTORI_DATABASE_NAME", "schotori")
	SSLMode  = get("SCHOTORI_DATABSE_SSLMODE", "disable")
	HTTPAddr = get("SCHOTORI_HTTP_ADDR", "127.0.0.1:5000")
)
