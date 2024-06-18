package models

type Models interface {
	ExecutePrompt(prompt string) (*AiModelResponse, error)
	Close() error
}

type AiModelResponse struct {
	Result string
}
