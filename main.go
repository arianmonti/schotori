package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"schotori/config"

	"github.com/gorilla/handlers"
	_ "github.com/lib/pq"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func main() {
	conn, _ := sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
		config.Cfg.DBUser,
		config.Cfg.DBPasswd,
		config.Cfg.DBName,
		config.Cfg.DBSSLMode,
	))

	if err := conn.Ping(); err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", index)

	file, _ := os.OpenFile("server.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	multiLogger := io.MultiWriter(os.Stdout, file)
	logged := handlers.LoggingHandler(multiLogger, mux)
	server := &http.Server{
		Addr:         config.Cfg.HTTPAddr,
		Handler:      logged,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	server.ListenAndServe()
}
