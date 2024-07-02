package claude

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/ashkenazi1/ai-model-hub/models"
	"github.com/ashkenazi1/go_requester"
)

type ClaudeAi struct {
	apiKey         string
	Model          string
	SystemPrompt   string
	UseChatHistory bool
	chatHistory    []string
}

func New(apiKey, model, systemPrompt string, useChatHistory bool) *ClaudeAi {
	return &ClaudeAi{
		apiKey:         apiKey,
		Model:          model,
		SystemPrompt:   systemPrompt,
		UseChatHistory: useChatHistory,
	}
}

func (c ClaudeAi) ExecutePrompt(prompt string) (*models.AiModelResponse, error) {
	requester := go_requester.New()
	res := new(models.AiModelResponse)

	var aiPrompt string
	if c.UseChatHistory {
		aiPrompt = fmt.Sprintf("Chat History: %s, Current Message: %s", c.chatHistory, prompt)
	} else {
		aiPrompt = prompt
	}

	messages := []map[string]interface{}{
		{"role": "user", "content": aiPrompt},
	}

	body := map[string]interface{}{
		"model":      c.Model,
		"max_tokens": 1024,
		"messages":   messages,
	}

	if c.SystemPrompt != "" {
		body["system"] = c.SystemPrompt
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

	if c.UseChatHistory {
		c.chatHistory = append(c.chatHistory, fmt.Sprintf("User: %s", prompt))
		c.chatHistory = append(c.chatHistory, fmt.Sprintf("AI: %s", string(aiRes)))

		if len(c.chatHistory) > 20 {
			c.chatHistory = c.chatHistory[2:]
		}
	}

	res.Result = string(aiRes)
	return res, nil

}

func (c ClaudeAi) Close() error {
	return nil
}
