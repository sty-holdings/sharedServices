package sharedServices

import (
	"cloud.google.com/go/vertexai/genai"
)

//goland:noinspection All
const (
	// SI = System Instruction
	SI_TOPIC_AI_QUESTION      = "ai-question"
	SI_TOPIC_ANALYZE_QUESTION = "analyze-question"
	SI_TOPIC_DETERMINE_API    = "determine-api"
	SI_TOPIC_GENERATE_ANSWER  = "generate-answer"
)

//goland:noinspection All
const (
	// SI = System Instruction
	SI_KEY_SIMPLE_ANSWER                     = "simple-answer"
	SI_KEY_COMPLEX_ANSWER                    = "complex-answer"
	SI_KEY_CATEGORY_PROMPY_COMPARISON        = "category-prompt-Comparison"
	SI_KEY_TIME_PERIOD_SPECIAL_WORDS_PRESENT = "time-period-special-words-present"
	SI_KEY_TIME_PERIOD_WORDS_PRESENT         = "time-period-words-present"
	SI_KEY_TIME_PERIOD_VALUES                = "time-period-values"
	SI_KEY_DETEMINE_API                      = "determine-api"
	SI_KEY_BUSINESS_ANALYST                  = "business-analyst"
)

type GeminiConfig struct {
	MaxOutputTokens    string             `json:"max_output_tokens"`
	ModelName          string             `json:"model_name"`
	SetTopProbability  string             `json:"set_top_probability"`
	SystemInstructions SystemInstructions `json:"system_instructions"`
	Temperature        string             `json:"temperature"`
}

type GeminiService struct {
	clientPtr *genai.Client
	modelPtr  *genai.GenerativeModel
	config    GeminiConfig
}

type InstructionSet struct {
	Instruction  string `json:"instruction"`
	OutputFormat string `json:"output_format"`
	SetDate      bool   `json:"set_date"`
}

type SystemInstructions struct {
	AIQuestion      map[string]InstructionSet `json:"ai-question"`
	AnalyzeQuestion map[string]InstructionSet `json:"analyze-question"`
	DetermineAPI    map[string]InstructionSet `json:"determine-api"`
	GenerateAnswer  map[string]InstructionSet `json:"generate-answer"`
}
