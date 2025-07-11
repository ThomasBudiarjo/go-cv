package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/thomasbudiarjo/go-cv/internal/handlers"
	"github.com/thomasbudiarjo/go-cv/internal/services"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	// Get API key from environment
	apiKey := os.Getenv("GOOGLE_AI_STUDIO_API_KEY")
	if apiKey == "" {
		log.Fatal("GOOGLE_AI_STUDIO_API_KEY environment variable is required")
	}

	// Initialize LLM client
	llmClient := services.NewLLMClient(apiKey)

	// Initialize handlers with dependencies
	appHandlers := handlers.NewHandlers(llmClient)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Route for the home page
	r.Get("/", appHandlers.HomePage)

	// API endpoint for document generation
	r.Post("/generate", appHandlers.GenerateDocuments)

	port := ":8082"
	fmt.Printf("Starting server on http://localhost%s\n", port)

	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}
}
