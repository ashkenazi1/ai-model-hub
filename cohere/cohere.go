package cohere

import (
	"context"
	"log"

	"github.com/ashkenazi1/ai-model-hub/models"
	cohere "github.com/cohere-ai/cohere-go/v2"
	client "github.com/cohere-ai/cohere-go/v2/client"
)

type Cohere struct {
	apiKey      string
	Model       string
	client      *client.Client
	UseHistory  bool
	ChatHistory []*cohere.Message
}

func New(key string, model string, useHistory bool) *Cohere {
	co := client.NewClient(client.WithToken(key))

	return &Cohere{
		apiKey:     key,
		Model:      model,
		client:     co,
		UseHistory: useHistory,
	}
}

func (c Cohere) ExecutePrompt(prompt string) (*models.AiModelResponse, error) {
	resp, err := c.client.Chat(
		context.TODO(),
		&cohere.ChatRequest{
			ChatHistory: c.ChatHistory,
			Message:     prompt,
			Connectors: []*cohere.ChatConnector{
				{Id: "web-search"},
			},
		},
	)

	if c.UseHistory {
		c.ChatHistory = append(c.ChatHistory, &cohere.Message{
			Role: "USER",
			User: &cohere.ChatMessage{
				Message: prompt,
			},
		})
		c.ChatHistory = append(c.ChatHistory, &cohere.Message{
			Role: "CHATBOT",
			Chatbot: &cohere.ChatMessage{
				Message: resp.Text,
			},
		})
	}

	if err != nil {
		log.Fatal(err)
	}

	// log.Printf("%+v", resp)
	return &models.AiModelResponse{
		Result: resp.Text,
	}, nil

}

func (c Cohere) Close() error {
	return nil
}
