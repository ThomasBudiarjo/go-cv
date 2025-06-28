
# Micro-SaaS Platform: CV-to-Cover-Letter Generator

## 1. Vision

A web-based application that helps job applicants quickly generate personalized cover letters and application emails by analyzing their CV and the corresponding job description.

## 2. Core Features (MVP)

*   **User Authentication:** Secure sign-up and login for users.
*   **CV Upload:** Users can upload their CV in common formats (PDF, DOCX).
*   **Job Description Input:** A text area for users to paste the job description.
*   **Document Generation:** With a single click, the platform will generate:
    *   A tailored Cover Letter.
    *   A professional Application Email.
*   **Editor:** A simple rich-text editor to review and modify the generated documents before downloading or copying.

## 3. Target Audience

*   Job seekers who want to save time and effort in writing application documents.
*   Students and recent graduates applying for their first jobs.
*   Professionals looking to switch careers.

## 4. Technology Stack

*   **Frontend:** HTMX.
*   **Backend:** Golang with Chi.
*   **Database:** SQLite.
*   **AI/LLM:** Gemma 3 27B.
*   **Deployment:** Docker for containerization.

## 5. High-Level Architecture

1.  **Frontend:**
    *   Handles user interactions (sign-up, login, file uploads).
    *   Communicates with the backend via a REST API.
    *   Displays the generated documents.

2.  **Backend (Go):**
    *   **API Endpoints:**
        *   `/auth/` for user registration and login (using JWT).
        *   `/cv/upload` to handle CV file uploads and parsing.
        *   `/generate/` to trigger the document generation process.
    *   **Services:**
        *   `UserService`: Manages user data.
        *   `CVParserService`: Extracts text and key information from CVs.
        *   `GenerationService`: Constructs prompts for the LLM and processes the output.

3.  **Database (Sqlite):**
    *   `users` table (id, email, password_hash, created_at).
    *   `cvs` table (id, user_id, original_filename, parsed_text, created_at).
    *   `generated_documents` table (id, user_id, type, content, created_at).

## 6. User Flow

1.  The user uploads their CV.
2.  The user pastes a job description to text area section.
3.  The user check generate Cover Letter and/or Email.
4.  The user clicks generate.
5.  The backend sends CV to LLM and get the information.
6.  The backend processes the request, sends a prompt to the LLM, and receives the generated text.
7.  The frontend displays the generated text in an text area.

## 7. Monetization Strategy

*   **Free Tier:** 3 free document generations per month.
*   **Pro Tier ($5/month):**
    *   Unlimited document generations.
    *   Access to premium templates.
    *   Ability to save and manage multiple CVs.

## 8. Project Milestones (MVP)

1.  **Phase 1: Mock/Basic Frontend**
    *   Set up the basic project structure.
    *   Create a simple UI with placeholders for CV upload, job description input, and generated content display.

2.  **Phase 2: API Development**
    *   Define and implement the core API endpoints for:
        *   CV upload and parsing.
        *   Triggering the generation of cover letters and emails.

3.  **Phase 3: Service Implementation**
    *   Implement the `CVParserService` to extract text from uploaded CVs.
    *   Implement the `GenerationService` to:
        *   Construct prompts for the LLM.
        *   Process the LLM's output.
    *   Integrate with the chosen LLM API.

4.  **Phase 4: Frontend Integration and Refinement**
    *   Connect the frontend to the backend API.
    *   Replace UI mockups with functional components.
    *   Refine the user interface and experience.
