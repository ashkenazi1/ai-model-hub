package models

type AiModel struct{}

func (a AiModel) Claude() string {
	return "claude"
}

func (a AiModel) ChatGpt() string {
	return "chatgpt"
}

func (a AiModel) Gemini() string {
	return "gemini"
}
