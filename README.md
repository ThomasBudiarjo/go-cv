# Go-CV: Cover Letter and Application Email Generator

This project is a micro-SaaS platform that generates cover letters and application emails based on a user's CV and a job description.

## Project Plan

The project plan is detailed in [app_plan.md](app_plan.md).

## Progress Tracker

### Phase 1: Mock/Basic Frontend

- [x] Set up the basic project structure.
- [x] Create a simple UI with placeholders for CV upload, job description input, and generated content display.

### Phase 2: API Development

- [x] Define and implement the core API endpoints for:
    - [x] CV upload and parsing.
    - [x] Triggering the generation of cover letters and emails.

### Phase 3: Service Implementation

- [x] Implement the `CVParserService` to extract text from uploaded CVs.
  - [x] Support for PDF, DOCX, and TXT file formats.
  - [x] Basic text extraction functionality.
- [x] Implement LLM integration with Google AI Studio:
    - [x] **Gemini Flash 2.5** API client implementation.
    - [x] Expert-level CV parsing prompts for structured data extraction.
    - [x] Professional cover letter generation prompts.
    - [x] Compelling application email generation prompts.
- [x] Two-step AI workflow:
    - [x] Step 1: LLM analyzes CV → extracts structured candidate profile (JSON).
    - [x] Step 2: LLM generates personalized documents using structured data + job description.

### Phase 4: Frontend Integration and Refinement

- [x] Connect the frontend to the backend API.
- [x] Replace UI mockups with functional components.
- [x] Real-time document generation with HTMX.
- [x] Enhanced UI with resizable text areas and debug information.
- [x] Error handling and loading indicators.

## Current Status

✅ **COMPLETED** - The application is now fully functional with real AI integration!

### What's Working:
- **File Upload**: Support for CV files (TXT, PDF*, DOCX*)
- **AI-Powered Analysis**: Gemini Flash 2.5 extracts structured candidate information
- **Smart Document Generation**: Personalized cover letters and application emails
- **Real-time Results**: Instant generation and display of professional documents
- **Debug Tools**: View extracted candidate data for transparency

*PDF and DOCX parsing shows helpful messages - TXT files work fully

### How to Run:

1. **Get Google AI Studio API Key:**
   - Visit [Google AI Studio](https://aistudio.google.com/)
   - Sign in with your Google account
   - Click "Get API Key" and create a new key

2. **Setup environment variables:**
   ```bash
   cp .env.example .env
   ```
   Edit the `.env` file and replace `your_api_key_here` with your actual API key:
   ```
   GOOGLE_AI_STUDIO_API_KEY=your_api_key_here
   ```

3. **Install dependencies:**
   ```bash
   go mod tidy
   ```

4. **Run the application:**
   ```bash
   go run cmd/web/main.go
   ```

5. **Visit:** http://localhost:8082

### Environment Setup:
- **Required**: `GOOGLE_AI_STUDIO_API_KEY` - Your Google AI Studio API key
- **File**: Copy `.env.example` to `.env` and add your credentials

### Technology Stack:
- **Backend**: Go with Chi router
- **Frontend**: HTMX with Pico CSS
- **AI**: Google AI Studio - Gemini Flash 2.5
- **File Processing**: Multi-format CV parsing
