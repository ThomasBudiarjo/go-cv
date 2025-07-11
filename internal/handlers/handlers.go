package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/thomasbudiarjo/go-cv/internal/services"
)

type Handlers struct {
	llmClient *services.LLMClient
}

func NewHandlers(llmClient *services.LLMClient) *Handlers {
	return &Handlers{
		llmClient: llmClient,
	}
}

// HomePage serves the main HTML page.
func (h *Handlers) HomePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./ui/html/home.page.html")
}

// GenerateDocuments handles the form submission from the frontend.
func (h *Handlers) GenerateDocuments(w http.ResponseWriter, r *http.Request) {
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

	// Initialize basic CV parser for text extraction
	cvParser := services.NewCVParserService()

	// Extract raw text from CV file
	cvText, err := cvParser.ParseCV(file, header)
	if err != nil {
		log.Printf("Failed to parse CV file: %v", err)
		http.Error(w, "Failed to parse CV file", http.StatusInternalServerError)
		return
	}

	// Step 1: Use LLM to parse CV and extract structured information
	log.Println("Parsing CV with LLM...")
	candidateJSON, err := h.llmClient.ParseCV(cvText)
	if err != nil {
		log.Printf("Failed to parse CV with LLM: %v", err)
		http.Error(w, "Failed to analyze CV content", http.StatusInternalServerError)
		return
	}

	// Step 2: Generate cover letter using LLM
	log.Println("Generating cover letter...")
	coverLetter, err := h.llmClient.GenerateCoverLetter(candidateJSON, jobDescription)
	if err != nil {
		log.Printf("Failed to generate cover letter: %v", err)
		http.Error(w, "Failed to generate cover letter", http.StatusInternalServerError)
		return
	}

	// Step 3: Generate application email using LLM
	log.Println("Generating application email...")
	email, err := h.llmClient.GenerateEmail(candidateJSON, jobDescription)
	if err != nil {
		log.Printf("Failed to generate email: %v", err)
		http.Error(w, "Failed to generate email", http.StatusInternalServerError)
		return
	}

	// Return HTML response for HTMX
	html := fmt.Sprintf(`
		<h2>Generated Documents</h2>
		<article id="cover-letter-result">
			<h3>Cover Letter</h3>
			<textarea rows="15" style="width: 100%%; font-family: monospace; resize: vertical;">%s</textarea>
		</article>
		<article id="email-result">
			<h3>Application Email</h3>
			<textarea rows="10" style="width: 100%%; font-family: monospace; resize: vertical;">%s</textarea>
		</article>
		<details style="margin-top: 20px;">
			<summary><strong>Debug: Extracted Candidate Information</strong></summary>
			<pre style="background: #f5f5f5; padding: 10px; border-radius: 5px; overflow-x: auto;">%s</pre>
		</details>
	`, coverLetter, email, candidateJSON)

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}
