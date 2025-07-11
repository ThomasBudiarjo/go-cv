package services

import (
	"fmt"
	"strings"
)

type GenerationService struct{}

type GenerationRequest struct {
	CVText         string
	JobDescription string
}

type GenerationResponse struct {
	CoverLetter string
	Email       string
}

func NewGenerationService() *GenerationService {
	return &GenerationService{}
}

func (gs *GenerationService) GenerateDocuments(req GenerationRequest) (*GenerationResponse, error) {
	if req.CVText == "" {
		return nil, fmt.Errorf("CV text is required")
	}
	if req.JobDescription == "" {
		return nil, fmt.Errorf("job description is required")
	}

	// TODO: Integrate with actual LLM (Gemma 3 27B or similar)
	// For now, generate mock responses based on the inputs
	
	coverLetter := gs.generateMockCoverLetter(req.CVText, req.JobDescription)
	email := gs.generateMockEmail(req.CVText, req.JobDescription)

	return &GenerationResponse{
		CoverLetter: coverLetter,
		Email:       email,
	}, nil
}

func (gs *GenerationService) generateMockCoverLetter(cvText, jobDesc string) string {

	jobLines := strings.Split(jobDesc, "\n")
	jobTitle := ""
	company := ""
	
	// Try to extract job title and company from job description
	for _, line := range jobLines {
		line = strings.TrimSpace(line)
		if line != "" {
			if jobTitle == "" {
				jobTitle = line
			} else if company == "" && strings.Contains(strings.ToLower(line), "company") {
				company = line
			}
			break
		}
	}

	return fmt.Sprintf(`Dear Hiring Manager,

I am writing to express my strong interest in the %s position. Based on my background and experience outlined in my CV, I believe I would be an excellent fit for this role.

Key qualifications from my CV:
%s

I am particularly excited about this opportunity because it aligns perfectly with my career goals and expertise. The job description resonates with my experience and I am confident I can contribute meaningfully to your team.

Thank you for considering my application. I look forward to discussing how my skills and experience can benefit your organization.

Sincerely,
[Your Name]

---
Note: This is a mock cover letter. LLM integration is pending.`, 
		jobTitle, 
		gs.extractKeyPoints(cvText))
}

func (gs *GenerationService) generateMockEmail(cvText, jobDesc string) string {
	lines := strings.Split(jobDesc, "\n")
	jobTitle := "the position"
	if len(lines) > 0 && strings.TrimSpace(lines[0]) != "" {
		jobTitle = strings.TrimSpace(lines[0])
	}

	return fmt.Sprintf(`Subject: Application for %s

Dear Hiring Manager,

I hope this email finds you well. I am writing to submit my application for %s as advertised.

I have attached my CV and cover letter for your review. My background includes:
%s

I am excited about the opportunity to contribute to your team and would welcome the chance to discuss my qualifications further.

Thank you for your time and consideration.

Best regards,
[Your Name]
[Your Phone Number]
[Your Email Address]

---
Note: This is a mock email. LLM integration is pending.`,
		jobTitle,
		jobTitle, 
		gs.extractKeyPoints(cvText))
}

func (gs *GenerationService) extractKeyPoints(cvText string) string {
	// Simple extraction of first few meaningful lines
	lines := strings.Split(cvText, "\n")
	var keyPoints []string
	
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" && len(line) > 10 {
			keyPoints = append(keyPoints, "• "+line)
			if len(keyPoints) >= 3 {
				break
			}
		}
	}
	
	if len(keyPoints) == 0 {
		return "• Professional experience and qualifications as detailed in CV"
	}
	
	return strings.Join(keyPoints, "\n")
}