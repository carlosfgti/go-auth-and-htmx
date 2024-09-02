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
	router.HandleFunc("/login", loginPage).Methods("GET", "OPTIONS")
	router.HandleFunc("/login", loginRequest).Methods("POST", "OPTIONS")

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

func loginPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("pages/login.html")
	if err != nil {
		http.Error(w, "Could not parse template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func loginRequest(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Basic authentication logic (replace with your own logic)
	if username == "user" && password == "password" {
		w.Write([]byte("<p>Login successful! Welcome, " + username + ".</p>"))
	} else {
		w.Write([]byte("<p>Invalid username or password. Please try again.</p>"))
	}
}
