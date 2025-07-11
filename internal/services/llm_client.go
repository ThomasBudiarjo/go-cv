package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	GEMINI_FLASH_API_URL = "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.5-flash:generateContent"
	REQUEST_TIMEOUT      = 30 * time.Second
)

type LLMClient struct {
	apiKey     string
	httpClient *http.Client
}

type GeminiRequest struct {
	Contents []Content `json:"contents"`
}

type Content struct {
	Parts []Part `json:"parts"`
}

type Part struct {
	Text string `json:"text"`
}

type GeminiResponse struct {
	Candidates []Candidate `json:"candidates"`
	Error      *APIError   `json:"error,omitempty"`
}

type Candidate struct {
	Content Content `json:"content"`
}

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewLLMClient(apiKey string) *LLMClient {
	return &LLMClient{
		apiKey: apiKey,
		httpClient: &http.Client{
			Timeout: REQUEST_TIMEOUT,
		},
	}
}

func (llm *LLMClient) GenerateText(prompt string) (string, error) {
	if prompt == "" {
		return "", fmt.Errorf("prompt cannot be empty")
	}

	requestBody := GeminiRequest{
		Contents: []Content{
			{
				Parts: []Part{
					{Text: prompt},
				},
			},
		},
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	url := fmt.Sprintf("%s?key=%s", GEMINI_FLASH_API_URL, llm.apiKey)
	
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := llm.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var geminiResp GeminiResponse
	if err := json.Unmarshal(body, &geminiResp); err != nil {
		return "", fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if geminiResp.Error != nil {
		return "", fmt.Errorf("API error %d: %s", geminiResp.Error.Code, geminiResp.Error.Message)
	}

	if len(geminiResp.Candidates) == 0 || len(geminiResp.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("no content generated")
	}

	return geminiResp.Candidates[0].Content.Parts[0].Text, nil
}

func (llm *LLMClient) ParseCV(cvText string) (string, error) {
	prompt := fmt.Sprintf(`You are an expert CV analyzer. Your task is to extract and structure information from CVs with high accuracy.

Analyze the following CV and extract information into a structured JSON format. Be thorough and precise in your extraction.

IMPORTANT: Return ONLY valid JSON, no explanations or markdown formatting.

Required JSON structure:
{
  "personal_info": {
    "name": "Full name",
    "email": "Email address", 
    "phone": "Phone number",
    "location": "City, State/Country",
    "linkedin": "LinkedIn URL if available",
    "website": "Personal website if available"
  },
  "professional_summary": "2-3 sentence professional summary highlighting key strengths and experience",
  "key_skills": {
    "technical": ["technical skill 1", "technical skill 2"],
    "soft_skills": ["communication", "leadership", "problem-solving"],
    "tools_and_technologies": ["specific tools", "software", "frameworks"]
  },
  "work_experience": [
    {
      "position": "Job title",
      "company": "Company name",
      "location": "City, State",
      "duration": "Start date - End date",
      "key_achievements": [
        "Specific achievement with metrics if available",
        "Another significant accomplishment"
      ],
      "responsibilities": "Main responsibilities summary"
    }
  ],
  "education": [
    {
      "degree": "Degree type and field",
      "institution": "University/School name",
      "location": "City, State",
      "graduation_year": "Year",
      "gpa": "GPA if mentioned",
      "relevant_coursework": ["course1", "course2"],
      "honors": "Any honors/awards"
    }
  ],
  "certifications": [
    {
      "name": "Certification name",
      "issuer": "Issuing organization",
      "year": "Year obtained",
      "expiry": "Expiry date if applicable"
    }
  ],
  "languages": [
    {
      "language": "Language name",
      "proficiency": "Native/Fluent/Conversational/Basic"
    }
  ],
  "projects": [
    {
      "name": "Project name",
      "description": "Brief description",
      "technologies": ["tech1", "tech2"],
      "duration": "Project duration"
    }
  ]
}

Rules:
- If a field is not found, use null for strings, [] for arrays, or {} for objects
- Extract ALL skills mentioned, categorized appropriately
- Include quantifiable achievements where available
- Be precise with dates and formatting
- Capture the most relevant and impressive aspects

CV Content:
%s`, cvText)

	return llm.GenerateText(prompt)
}

func (llm *LLMClient) GenerateCoverLetter(candidateJSON, jobDescription string) (string, error) {
	prompt := fmt.Sprintf(`You are an expert career counselor and professional writer specializing in creating compelling cover letters.

Task: Write a highly personalized, professional cover letter that will capture the hiring manager's attention.

Candidate Profile (JSON):
%s

Job Description:
%s

Requirements for the cover letter:

STRUCTURE:
1. Professional header with proper salutation
2. Opening paragraph: Hook the reader with a strong opening that mentions the specific role and company
3. Body paragraph 1: Match your most relevant experience/achievements to job requirements 
4. Body paragraph 2: Demonstrate cultural fit and specific value you'll bring
5. Closing paragraph: Call to action and professional sign-off

CONTENT GUIDELINES:
- Use the candidate's actual name and contact information from the JSON
- Reference specific requirements from the job description
- Quantify achievements where possible (use metrics from the candidate's experience)
- Show genuine knowledge about the company/role (extract details from job description)
- Use active voice and strong action verbs
- Avoid generic phrases like "I am writing to apply" or "Please find my resume attached"
- Make each sentence add value and demonstrate fit
- Tone: Professional yet personable, confident but not arrogant

PERSONALIZATION:
- Match specific technical skills from candidate to job requirements
- Highlight relevant experience that directly relates to the role
- Reference years of experience and career progression
- Include relevant certifications or education if they match job needs

FORMATTING:
- Use proper business letter format
- Include candidate's contact information at the top
- Professional salutation (use "Hiring Manager" if no name provided)
- Single-spaced paragraphs with line breaks between
- Professional closing ("Best regards" or "Sincerely")

Write ONLY the cover letter content. Do not include explanations, notes, or formatting instructions.`, candidateJSON, jobDescription)

	return llm.GenerateText(prompt)
}

func (llm *LLMClient) GenerateEmail(candidateJSON, jobDescription string) (string, error) {
	prompt := fmt.Sprintf(`You are an expert recruiter and career coach specializing in crafting compelling application emails that get responses.

Task: Write a professional application email that will make the hiring manager want to learn more about the candidate.

Candidate Profile (JSON):
%s

Job Description:
%s

Requirements for the application email:

SUBJECT LINE:
- Create a compelling, specific subject line that includes the role title
- Should grab attention and indicate strong fit
- Examples: "Experienced [Role] | [Key Skill/Achievement]" or "[Years] Years [Relevant Experience] - [Role Title] Application"

EMAIL STRUCTURE:
1. Professional greeting with proper salutation
2. Opening: Strong hook that immediately shows relevance and value
3. Body paragraph: 2-3 most compelling qualifications that match job requirements
4. Value proposition: What specific value you bring to the organization
5. Call to action: Request for interview/next steps
6. Professional closing with contact information

CONTENT GUIDELINES:
- Keep it concise (under 200 words in body)
- Lead with your strongest, most relevant qualification
- Quantify achievements where possible (use metrics from candidate profile)
- Reference specific requirements from the job posting
- Show knowledge of the company/role
- Use confident, professional tone
- Avoid repeating everything from CV - highlight the best parts
- Make every sentence count

PERSONALIZATION:
- Extract company name and role title from job description
- Match candidate's experience level to position requirements
- Highlight most relevant technical skills and achievements
- Reference specific qualifications that make candidate stand out

FORMAT:
Subject: [compelling subject line]

Dear [Hiring Manager/Team],

[Email body content]

Best regards,
[Candidate Name]
[Phone Number]
[Email Address]

Write ONLY the email content with subject line. Do not include explanations or additional notes.`, candidateJSON, jobDescription)

	return llm.GenerateText(prompt)
}