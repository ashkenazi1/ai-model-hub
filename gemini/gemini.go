package gemini

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/ashkenazi1/ai-model-hub/models"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type GeminiAIModel struct {
	usedModel string
	model     *genai.GenerativeModel
	client    *genai.Client
}

func New(apiKey, model string) (*GeminiAIModel, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}

	aiModel := client.GenerativeModel(model)
	return &GeminiAIModel{model: aiModel, client: client, usedModel: model}, nil
}

func (g *GeminiAIModel) QueryImage(imgPath, prompt string) (string, error) {
	if g.usedModel != "gemini-pro-vision" {
		return "", fmt.Errorf("this model does not support image generation")
	}

	path, err := filepath.Abs(imgPath)
	if err != nil {
		return "", err
	}

	imgData, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	imgSuffix := strings.Replace(filepath.Ext(imgPath), ".", "", 1)

	geminiPrompt := []genai.Part{
		genai.ImageData(imgSuffix, imgData),
		genai.Text(prompt),
	}

	ctx := context.Background()
	resp, err := g.model.GenerateContent(ctx, geminiPrompt...)
	if err != nil {
		log.Fatal(err)
	}

	byteData, err := json.Marshal(resp.Candidates[0].Content.Parts)
	if err != nil {
		return "", err
	}

	textRes := string(byteData)
	return textRes, nil
}

func (g *GeminiAIModel) ExecutePrompt(prompt string) (*models.AiModelResponse, error) {

	if g.usedModel == "gemini-pro-vision" {
		return nil, fmt.Errorf("this model does not support text generation")
	}

	geminiPrompt := []genai.Part{
		genai.Text(prompt),
	}

	ctx := context.Background()
	resp, err := g.model.GenerateContent(ctx, geminiPrompt...)
	if err != nil {
		return nil, err
	}

	byteData, err := json.Marshal(resp.Candidates[0].Content.Parts[0])
	if err != nil {
		return nil, err
	}

	res := new(models.AiModelResponse)
	res.Result = string(byteData)
	return res, nil
}

func (g *GeminiAIModel) Close() error {
	return g.client.Close()
}
