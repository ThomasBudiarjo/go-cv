// main.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Initialize a new Chi router
	r := chi.NewRouter()

	// Use a logger middleware for handy request logging
	r.Use(middleware.Logger)

	// A route to serve our HTML page
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./ui/html/home.page.html")
	})

	port := ":8080"
	fmt.Printf("Starting server on http://localhost%s\n", port)

	// Start the server with the Chi router
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}
}
