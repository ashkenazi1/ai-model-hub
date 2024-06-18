package chatgpt

import (
	"context"

	"github.com/ashkenazi1/ai-model-hub/models"
	openai "github.com/sashabaranov/go-openai"
)

type OpenAiModel struct {
	apiKey string
	model  string
}

func New(key, model string) *OpenAiModel {
	return &OpenAiModel{
		apiKey: key,
		model:  model, //"gpt-3.5-turbo",
	}
}

func (o OpenAiModel) ExecutePrompt(prompt string) (*models.AiModelResponse, error) {
	client := openai.NewClient(o.apiKey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: o.model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)

	res := new(models.AiModelResponse)

	if err != nil {
		res.Result = err.Error()
		return res, err
	}

	res.Result = resp.Choices[0].Message.Content
	return res, err
}

func (o OpenAiModel) Close() error {
	return nil
}
