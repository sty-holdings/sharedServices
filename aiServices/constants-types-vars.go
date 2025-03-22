package sharedServices

import (
	"cloud.google.com/go/vertexai/genai"

	errs "github.com/sty-holdings/sharedServices/v2025/errorServices"
)

//goland:noinspection All
const (
	// SI = System Instruction
	SI_TOPIC_ANALYZE_QUESTION = "analyze-question"
	SI_TOPIC_GENERATE_ANSWER  = "generate-answer"
)

//goland:noinspection All
const (
	// SI = System Instruction
	SI_KEY_CATEGORY_SENTENCE  = "category_sentence"
	SI_KEY_SPECIAL_WORDS      = "special_words"
	SI_KEY_TIME_PERIOD_VALUES = "time_period_values"
	SI_KEY_BUSINESS_ANALYST   = "business_analyst"
	SI_KEY_MARKETING_ANALYST  = "marketing_analyst"
	SI_KEY_NOT_SUPPORTED      = "not_supported"
)

type AIConfig struct {
	MaxOutputTokens    string             `json:"max_output_tokens"`
	ModelName          string             `json:"model_name"`
	SetTopProbability  string             `json:"set_top_probability"`
	SystemInstructions SystemInstructions `yaml:"system_instructions"`
	Temperature        string             `json:"temperature"`
}

type AIResponse struct {
	Response   string
	SIKey      string
	TokenCount genai.UsageMetadata
	ErrorInfo  errs.ErrorInfo
}

type AIService struct {
	clientPtr *genai.Client
	debugOn   bool
	modelPtrs map[string]*genai.GenerativeModel
	config    AIConfig
}

type AnalyzeQuestions struct {
	CategorySentence InstructionSet `yaml:"category_sentence"`
	SpecialWords     InstructionSet `yaml:"special_words"`
	TimePeriodValues InstructionSet `yaml:"time_period_values"`
}

type GenerateAnswer struct {
	BusinessAnalyst  InstructionSet `yaml:"business_analyst"`
	MarketingAnalyst InstructionSet `yaml:"marketing_analyst"`
	NotSupported     InstructionSet `yaml:"not_supported"`
}

type InstructionSet struct {
	Instruction  string `json:"instruction"`
	OutputFormat string `json:"output_format"`
	SetDate      bool   `json:"set_date"`
}

type SystemInstructions struct {
	AnalyzeQuestions AnalyzeQuestions `yaml:"analyze_questions"`
	GenerateAnswer   GenerateAnswer   `yaml:"generate_answer"`
}

var (
	modelPoolNames = []string{
		"pool-0",
		"pool-1",
		"pool-2",
	}

	SITopicAnalyzeQuestionKeys = []string{
		SI_KEY_CATEGORY_SENTENCE,
		SI_KEY_SPECIAL_WORDS,
		SI_KEY_TIME_PERIOD_VALUES,
	}

	SITopicGenerateAnswerKeys = []string{
		SI_KEY_BUSINESS_ANALYST,
		SI_KEY_MARKETING_ANALYST,
		SI_KEY_NOT_SUPPORTED,
	}

	siTopicKeyPoolAssignment = map[string]string{
		// SI_TOPIC_ANALYZE_QUESTION
		SI_KEY_CATEGORY_SENTENCE:  "pool-0",
		SI_KEY_SPECIAL_WORDS:      "pool-1",
		SI_KEY_TIME_PERIOD_VALUES: "pool-2",
		// SI_TOPIC_GENERATE_ANSWER
		SI_KEY_BUSINESS_ANALYST:  "pool-0",
		SI_KEY_MARKETING_ANALYST: "pool-0",
		SI_KEY_NOT_SUPPORTED:     "pool-0",
	}
)
