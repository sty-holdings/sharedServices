package sharedServices

import (
	"cloud.google.com/go/vertexai/genai"

	errs "github.com/sty-holdings/sharedServices/v2025/errorServices"
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
	// SI_TOPIC_AI_QUESTION
	SI_KEY_SIMPLE_ANSWER  = "simple-answer"
	SI_KEY_COMPLEX_ANSWER = "complex-answer"
	// SI_TOPIC_ANALYZE_QUESTION
	SI_KEY_CATEGORY_PROMPY_COMPARISON        = "category-prompt-comparison"
	SI_KEY_TIME_PERIOD_SPECIAL_WORDS_PRESENT = "time-period-special-words-present"
	SI_KEY_TIME_PERIOD_WORDS_PRESENT         = "time-period-words-present"
	SI_KEY_TIME_PERIOD_VALUES                = "time-period-values"
	// SI_TOPIC_DETERMINE_API
	SI_KEY_DETEMINE_API = "determine-api"
	// SI_TOPIC_GENERATE_ANSWER
	SI_KEY_BUSINESS_ANALYST = "business-analyst"
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
	debugOn   bool
	modelPtrs map[string]*genai.GenerativeModel
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

type GeminiResponse struct {
	Response   string
	SIKey      string
	TokenCount genai.UsageMetadata
	ErrorInfo  errs.ErrorInfo
}

var (
	modelPoolNames = []string{
		"pool-0",
		"pool-1",
		"pool-2",
		"pool-3",
	}

	SITopicAnalyzeQuestionKeys = []string{
		SI_KEY_CATEGORY_PROMPY_COMPARISON,
		SI_KEY_TIME_PERIOD_SPECIAL_WORDS_PRESENT,
		SI_KEY_TIME_PERIOD_WORDS_PRESENT,
		SI_KEY_TIME_PERIOD_VALUES,
	}

	siTopicKeyPoolAssignment = map[string]string{
		// SI_TOPIC_AI_QUESTION
		SI_KEY_SIMPLE_ANSWER:  "pool-0",
		SI_KEY_COMPLEX_ANSWER: "pool-1",
		// SI_TOPIC_ANALYZE_QUESTION
		SI_KEY_CATEGORY_PROMPY_COMPARISON:        "pool-0",
		SI_KEY_TIME_PERIOD_SPECIAL_WORDS_PRESENT: "pool-1",
		SI_KEY_TIME_PERIOD_WORDS_PRESENT:         "pool-2",
		SI_KEY_TIME_PERIOD_VALUES:                "pool-3",
		// SI_TOPIC_DETERMINE_API
		SI_KEY_DETEMINE_API: "pool-0",
		// SI_TOPIC_GENERATE_ANSWER
		SI_KEY_BUSINESS_ANALYST: "pool-0",
	}
)
