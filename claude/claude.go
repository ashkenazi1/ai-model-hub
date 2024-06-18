package claude

import (
	"bytes"
	"encoding/json"

	"github.com/ashkenazi1/ai-model-hub/models"
	"github.com/ashkenazi1/go_requester"
)

type ClaudeAi struct {
	apiKey string
	Model  string
}

func New(apiKey string, model string) *ClaudeAi {
	return &ClaudeAi{
		apiKey: apiKey,
		Model:  model,
	}
}

func (c ClaudeAi) ExecutePrompt(prompt string) (*models.AiModelResponse, error) {
	requester := go_requester.New()
	res := new(models.AiModelResponse)

	messages := []map[string]interface{}{
		{"role": "user", "content": prompt},
	}

	body := map[string]interface{}{
		"model":      c.Model,
		"max_tokens": 1024,
		"messages":   messages,
	}
	jsonBytes, _ := json.Marshal(body)
	bodyReader := bytes.NewReader(jsonBytes)

	requester.Headers["x-api-key"] = c.apiKey
	requester.Headers["Content-Type"] = "application/json"
	requester.Headers["anthropic-version"] = "2023-06-01"
	aiRes, err := requester.Post("https://api.anthropic.com/v1/messages", bodyReader)
	if err != nil {
		return nil, err
	}
	res.Result = string(aiRes)
	return res, nil

}

func (c ClaudeAi) Close() error {
	return nil
}
