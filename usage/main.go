package main

import (
	"log"
	"os"

	"github.com/ashkenazi1/ai-model-hub/ai"
	"github.com/ashkenazi1/ai-model-hub/models"
)

func main() {
	OpenAiKey := os.Getenv("API_KEY")
	// ClaudeAPIKey := os.Getenv("API_KEY")
	// Gemini_Api_Key := os.Getenv("API_KEY")

	choosenAi := models.AiModel{}
	choosenModel := models.ChatGptModel{}

	ai, err := ai.New(choosenAi.ChatGpt(), OpenAiKey, choosenModel.Gpt4o())
	if err != nil {
		panic(err)
	}
	res, err := ai.ExecutePrompt("Hello, how are you?")
	if err != nil {
		panic(err)
	}
	log.Println(res.Result)
}
