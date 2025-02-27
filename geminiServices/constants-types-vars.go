package sharedServices

import (
	"cloud.google.com/go/vertexai/genai"
)

//goland:noinspection ALL
const ()

type GeminiConfig struct {
	GeminiMaxOutputTokens    string                        `json:"gemini_max_output_tokens"`
	GeminiModelName          string                        `json:"gemini_model_name"`
	GeminiSetTopProbability  string                        `json:"gemini_set_top_probability"`
	GeminiSystemInstructions map[string]SystemInstructions `json:"gemini_system_instructions"`
	GeminiTemperature        string                        `json:"gemini_temperature"`
}

type GeminiService struct {
	GeminiClientPtr      *genai.Client
	GeminiModelPtr       *genai.GenerativeModel
	geminiConfig         GeminiConfig
	DKSystemInstructions map[string]SystemInstructions
}

type InstructionSet struct {
	Instruction  string `json:"instruction"`
	OutputFormat string `json:"output_format"`
	SetDate      bool   `json:"set_date"`
}

type SystemInstructions struct {
	AnalyzeQuestion map[string]InstructionSet `json:"analyze-question"`
	Hal             map[string]InstructionSet `json:"hal"`
}
