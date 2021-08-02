package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DBUser    string `envconfig:"SCHOTORI_DATABASE_USER" default:"root"`
	DBHost    string `envconfig:"SCHOTORI_DATABASE_HOST" default:"localhost"`
	DBPasswd  string `envconfig:"SCHOTORI_DATABAE_PASSWORD" default:"root"`
	DBName    string `envconfig:"SCHOTORI_DATABASE_NAME" default:"schotori"`
	DBSSLMode string `envconfig:"SCHOTORI_DATABSE_SSLMODE" default:"disable"`
	HTTPAddr  string `envconfig:"SCHOTORI_HTTP_ADDR" default:"127.0.0.1:5000"`
}

var Cfg Config

func init() {
	godotenv.Load()
	err := envconfig.Process("", &Cfg)
	if err != nil {
		panic(err)
	}
}
