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

- [ ] Implement the `CVParserService` to extract text from uploaded CVs.
  - [ ] Construct prompts for The LLM.
  - [ ] Process the LLM's output.
- [ ] Implement the `GenerationService` to:
    - [ ] Construct prompts for the LLM.
    - [ ] Process the LLM's output.
- [ ] Integrate with the chosen LLM API.

### Phase 4: Frontend Integration and Refinement

- [ ] Connect the frontend to the backend API.
- [ ] Replace UI mockups with functional components.
- [ ] Refine the user interface and experience.
