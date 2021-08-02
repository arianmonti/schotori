package main

import (
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)

	file, _ := os.OpenFile("server.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	multiLogger := io.MultiWriter(os.Stdout, file)
	logged := handlers.LoggingHandler(multiLogger, mux)
	server := &http.Server{
		Addr:         "127.0.0.1:5000",
		Handler:      logged,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	server.ListenAndServe()
}