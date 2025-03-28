package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ChatRequest struct {
	Model    string        `json:"model"`
	Messages []MessageData `json:"messages"`
}

type MessageData struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatResponse struct {
	Choices []struct {
		Message MessageData `json:"message"`
	} `json:"choices"`
}

func CallLLM(prompt string, apiKey string) ([]string, error) {
	url := "https://api.groq.com/openai/v1/chat/completions"

	reqBody := ChatRequest{
		Model: "llama3-8b-8192",
		Messages: []MessageData{
			{Role: "system", Content: `
You are an AI task planner, not an analyst.

Given a user request, your job is to return ONLY a JSON array listing the containerized tasks to run.

 Allowed task names are:
- "clean_data"
- "sentiment_analysis"

Your response MUST:
- Be a single JSON array like ["clean_data", "sentiment_analysis"]
- NOT contain any explanations or text
- NOT perform the actual analysis

Examples:

User: Clean the data only  
→ ["clean_data"]

User: Analyze sentiment of a review  
→ ["clean_data", "sentiment_analysis"]

User: Just sentiment analysis  
→ ["sentiment_analysis"]

If the request is unclear, return ["clean_data"]
`},

			{Role: "user", Content: prompt},
		},
	}

	jsonData, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+apiKey)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	bodyBytes, _ := io.ReadAll(res.Body)

	var response ChatResponse
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		fmt.Println(" Error parsing response JSON:", string(bodyBytes))
		return nil, err
	}

	
	if len(response.Choices) == 0 {
		fmt.Println(" LLM response had no choices. Raw response:", string(bodyBytes))
		return nil, fmt.Errorf("empty response from LLM")
	}

	 
	var tasks []string
	err = json.Unmarshal([]byte(response.Choices[0].Message.Content), &tasks)
	if err != nil {
		fmt.Println(" Failed to parse LLM output as JSON array:", response.Choices[0].Message.Content)
		return nil, err
	}

	return tasks, nil
}
