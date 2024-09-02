package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", handleIndexRequest).Methods("GET", "OPTIONS")

	// Start the HTTPS server
	server := &http.Server{
		Addr:    ":8443",
		Handler: router,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("failed to start https server: %v", err)
	}
}

func handleIndexRequest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
