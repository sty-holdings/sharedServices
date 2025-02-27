package sharedServices

import (
	"cloud.google.com/go/vertexai/genai"
)

//goland:noinspection ALL
const ()

type GeminiConfig struct {
	MaxOutputTokens    string             `json:"max_output_tokens"`
	ModelName          string             `json:"model_name"`
	SetTopProbability  string             `json:"set_top_probability"`
	SystemInstructions SystemInstructions `json:"system_instructions"`
	Temperature        string             `json:"temperature"`
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
