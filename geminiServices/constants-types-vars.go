package sharedServices

import (
	"cloud.google.com/go/vertexai/genai"

	ctv "github.com/sty-holdings/sharedServices/v2025/constantsTypesVars"
)

//goland:noinspection ALL
const ()

type GeminiConfig struct {
	GeminiMaxOutputTokens    string                           `json:"gemini_max_output_tokens"`
	GeminiModelName          string                           `json:"gemini_model_name"`
	GeminiSetTopProbability  string                           `json:"gemini_set_top_probability"`
	GeminiSystemInstructions map[string]ctv.SystemInstruction `json:"gemini_system_instructions"`
	GeminiTemperature        string                           `json:"gemini_temperature"`
}

type GeminiService struct {
	GeminiClientPtr *genai.Client
	geminiConfig    GeminiConfig
}
