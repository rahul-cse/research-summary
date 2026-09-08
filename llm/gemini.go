package llm

import (
	"os"

	"context"

	"encoding/json"

	"google.golang.org/genai"
)

func AnalyzeText(text string) (Analysis, error) {

	prompt := ` You are analyzing a research paper. 
	Read the research paper text below and classify it into four dimensions:
	1. domain: The main research domain or domains.
	2. Methods: The main research methods used.
	3. Study_application: What the research actually studies, does, or applies.
	4. Summary: Give a very concise summary of the research paper.

	A paper can belong to multiple domains, use multiple methods, and have multiple study/application types.
	Return ONLY valid JSON in exactly this structure: { "domain": ["..."], "methods": ["..."], "study_application": ["..."], "summary": "..." }
	Research paper:` + text

	ctx := context.Background()

	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})

	if err != nil {
		return Analysis{}, err
	}

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-3.5-flash-lite",
		genai.Text(prompt),
		nil,
	)

	var analysis Analysis

	err = json.Unmarshal([]byte(result.Text()), &analysis)
	if err != nil {
		return Analysis{}, err
	}

	return analysis, nil
}
