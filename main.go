package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", handleIndexRequest).Methods("GET", "OPTIONS")
	router.HandleFunc("/hello", handleFunc).Methods("GET", "OPTIONS")

	// Start the HTTPS server
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("failed to start https server: %v", err)
	}
}

func handleIndexRequest(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("pages/home.html")
	if err != nil {
		http.Error(w, "Could not parse template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func handleFunc(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Hello from the server!"))
}
