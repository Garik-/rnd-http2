package main

import (
	"log"
	"net/http"
	"time"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

var helloWorld = []byte("Hello, HTTP/2.0 world!\n")

const address = ":8080"

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write(helloWorld)
}

func main() {
	h2s := &http2.Server{
		MaxConcurrentStreams: 250,
		IdleTimeout:          30 * time.Second,
	}

	server := &http.Server{
		Addr:           address,
		Handler:        h2c.NewHandler(http.HandlerFunc(handler), h2s),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Starting optimized HTTP/2 h2c server on " + address)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
