package chatgpt

import (
	"context"
	"fmt"
	"log"

	"github.com/ashkenazi1/ai-model-hub/models"
	openai "github.com/sashabaranov/go-openai"
)

type OpenAiModel struct {
	apiKey         string
	Model          string
	SystemPrompt   string
	UseChatHistory bool
	chatHistory    []string
}

func New(key, model string, systemPrompt string, useChatHistory bool) *OpenAiModel {
	return &OpenAiModel{
		apiKey:         key,
		Model:          model,
		SystemPrompt:   systemPrompt,
		UseChatHistory: useChatHistory,
	}
}

func (o OpenAiModel) ExecutePrompt(prompt string) (*models.AiModelResponse, error) {
	client := openai.NewClient(o.apiKey)

	req := openai.ChatCompletionRequest{
		Model: o.Model,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: o.SystemPrompt,
			},
		},
	}

	var aiPrompt string
	if o.UseChatHistory {
		aiPrompt = fmt.Sprintf("Chat History: %s, Current Message: %s", o.chatHistory, prompt)
	} else {
		aiPrompt = prompt
	}

	req.Messages = append(req.Messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: aiPrompt,
	})

	res := new(models.AiModelResponse)
	resp, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		res.Result = err.Error()
		return res, err
	}

	if o.UseChatHistory {
		o.chatHistory = append(o.chatHistory, fmt.Sprintf("User: %s", prompt))
		o.chatHistory = append(o.chatHistory, fmt.Sprintf("AI: %s", resp.Choices[0].Message.Content))

		if len(o.chatHistory) > 20 {
			o.chatHistory = o.chatHistory[2:]
		}
	}

	log.Println(o.chatHistory)

	res.Result = resp.Choices[0].Message.Content
	return res, err
}

func (o OpenAiModel) Close() error {
	return nil
}
