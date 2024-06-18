package ai

import (
	"errors"

	"github.com/ashkenazi1/ai-model-hub/chatgpt"
	"github.com/ashkenazi1/ai-model-hub/claude"
	"github.com/ashkenazi1/ai-model-hub/gemini"
	"github.com/ashkenazi1/ai-model-hub/models"
)

type AiModel struct {
	model models.Models
}

func New(modelName, apiKey, choosenModel string) (*AiModel, error) {
	var model models.Models

	switch modelName {
	case "claude":
		model = claude.New(apiKey, choosenModel)
		return &AiModel{model: model}, nil

	case "chatgpt":
		model = chatgpt.New(apiKey, choosenModel)
		return &AiModel{model: model}, nil

	case "gemini":
		model, err := gemini.New(apiKey, choosenModel)
		if err != nil {
			return nil, err
		}
		return &AiModel{model: model}, nil
	}

	return nil, errors.New("model not found")
}

func (a AiModel) ExecutePrompt(prompt string) (*models.AiModelResponse, error) {
	return a.model.ExecutePrompt(prompt)
}

func (a AiModel) Close() error {
	return a.model.Close()
}
