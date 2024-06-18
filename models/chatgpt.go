package models

type ChatGptModel struct{}

func (c ChatGptModel) Gpt3_5Turbo() string {
	return "gpt-3.5-turbo"
}

func (c ChatGptModel) Gpt4o() string {
	return "gpt-4o"
}

func (c ChatGptModel) Gpt4Turbo() string {
	return "gpt-4-turbo"
}
