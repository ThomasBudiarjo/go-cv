// cmd/web/main.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/thomasbudiarjo/go-cv/internal/handlers" // <-- Import your handlers
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Route for the home page, using our new handler
	r.Get("/", handlers.HomePage)

	// **New API Endpoint**
	// This route will receive the POST request from our HTMX form
	r.Post("/generate", handlers.GenerateDocuments)

	port := ":8080"
	fmt.Printf("Starting server on http://localhost%s\n", port)

	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}
}
