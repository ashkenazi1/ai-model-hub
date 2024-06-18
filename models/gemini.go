package models

type GeminiModel struct{}

func (g GeminiModel) GeminiPro() string {
	return "gemini-1.5-pro"
}

func (g GeminiModel) GeminiProFlash() string {
	return "gemini-1.5-flash"
}

func (g GeminiModel) GeminiProVision() string {
	return "gemini-pro-vision"
}

func (g GeminiModel) GeminiProVisionTextV2() string {
	return "gemini-pro-vision-text-v2"
}
