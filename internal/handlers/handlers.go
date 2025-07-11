package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/thomasbudiarjo/go-cv/internal/services"
)

// HomePage serves the main HTML page.
func HomePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./ui/html/home.page.html")
}

// GenerateDocuments handles the form submission from the frontend.
func GenerateDocuments(w http.ResponseWriter, r *http.Request) {
	// Parse multipart form
	err := r.ParseMultipartForm(10 << 20) // 10 MB max
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	// Get uploaded CV file
	file, header, err := r.FormFile("cv")
	if err != nil {
		http.Error(w, "Failed to get CV file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Get job description
	jobDescription := r.FormValue("job-description")
	if jobDescription == "" {
		http.Error(w, "Job description is required", http.StatusBadRequest)
		return
	}

	// Initialize services
	cvParser := services.NewCVParserService()
	generator := services.NewGenerationService()

	// Parse CV
	cvText, err := cvParser.ParseCV(file, header)
	if err != nil {
		log.Printf("Failed to parse CV: %v", err)
		http.Error(w, "Failed to parse CV file", http.StatusInternalServerError)
		return
	}

	// Generate documents
	genRequest := services.GenerationRequest{
		CVText:         cvText,
		JobDescription: jobDescription,
	}

	response, err := generator.GenerateDocuments(genRequest)
	if err != nil {
		log.Printf("Failed to generate documents: %v", err)
		http.Error(w, "Failed to generate documents", http.StatusInternalServerError)
		return
	}

	// Return HTML response for HTMX
	html := fmt.Sprintf(`
		<h2>Generated Documents</h2>
		<article id="cover-letter-result">
			<h3>Cover Letter</h3>
			<textarea rows="15" style="width: 100%%; font-family: monospace;">%s</textarea>
		</article>
		<article id="email-result">
			<h3>Application Email</h3>
			<textarea rows="10" style="width: 100%%; font-family: monospace;">%s</textarea>
		</article>
	`, response.CoverLetter, response.Email)

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}
