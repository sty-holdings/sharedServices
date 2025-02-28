package sharedServices

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"cloud.google.com/go/vertexai/genai"
	"google.golang.org/api/option"

	ctv "github.com/sty-holdings/sharedServices/v2025/constantsTypesVars"
	errs "github.com/sty-holdings/sharedServices/v2025/errorServices"
	hlps "github.com/sty-holdings/sharedServices/v2025/helpers"
)

var (
	CTXBackground = context.Background()
)

// NewGeminiService - creates a new gemini client
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func NewGeminiService(gcpCredentialsFilename string, gcpProjectId string, gcpLocation string, geminiConfigFilename string) (geminiServicePtr *GeminiService, errorInfo errs.ErrorInfo) {

	var (
		tGeminiConfig GeminiConfig
	)

	if errorInfo = hlps.CheckValueNotEmpty(gcpCredentialsFilename, errs.ErrRequiredParameterMissing, ctv.FN_GCP_CREDENTIAL_FILENAME); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckValueNotEmpty(geminiConfigFilename, errs.ErrRequiredParameterMissing, ctv.FN_SERVICE_CONFIG_FILENAME); errorInfo.Error != nil {
		return
	}

	if tGeminiConfig, errorInfo = loadGeminiConfig(geminiConfigFilename); errorInfo.Error != nil {
		return
	}

	if errorInfo = validateGeminiConfig(tGeminiConfig); errorInfo.Error != nil {
		return
	}

	geminiServicePtr = &GeminiService{
		config: tGeminiConfig,
	}

	if geminiServicePtr.clientPtr, errorInfo.Error = genai.NewClient(
		context.Background(),
		gcpProjectId,
		gcpLocation,
		option.WithCredentialsFile(gcpCredentialsFilename),
	); errorInfo.Error != nil {
		return
	}

	errorInfo = geminiServicePtr.buildModel()

	return
}

// buildModel - will create the model instance and configure the settings
//
//	Customer Message: none
//	Errors: none
//	Verifications: none
func (geminiServicePtr *GeminiService) buildModel() (errorInfo errs.ErrorInfo) {

	var (
		tFloat32 float32
		tFloat64 float64
		tInt32   int32
		tInt64   int64
	)

	geminiServicePtr.modelPtr = geminiServicePtr.clientPtr.GenerativeModel(geminiServicePtr.config.ModelName)

	if tInt64, errorInfo.Error = strconv.ParseInt(geminiServicePtr.config.MaxOutputTokens, 10, 32); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errs.ErrIntegerInvalid, fmt.Sprintf("%s%s\n", ctv.LBL_GEMINI_MAX_OUTPUT_TOKENS, geminiServicePtr.config.MaxOutputTokens))
		return
	}
	tInt32 = int32(tInt64)
	geminiServicePtr.modelPtr.MaxOutputTokens = &tInt32

	if tFloat64, errorInfo.Error = strconv.ParseFloat(geminiServicePtr.config.SetTopProbability, 64); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errs.ErrFloatInvalid, fmt.Sprintf("%s%s\n", ctv.LBL_GEMINI_SET_TOP_PROBABILITY, geminiServicePtr.config.SetTopProbability))
		return
	}
	tFloat32 = float32(tFloat64)
	geminiServicePtr.modelPtr.SetTopP(tFloat32)

	if tFloat64, errorInfo.Error = strconv.ParseFloat(geminiServicePtr.config.Temperature, 64); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errs.ErrFloatInvalid, fmt.Sprintf("%s%s\n", ctv.LBL_GEMINI_TEMPERATURE, geminiServicePtr.config.Temperature))
		return
	}
	tFloat32 = float32(tFloat64)
	geminiServicePtr.modelPtr.Temperature = &tFloat32

	return
}

// GenerateContent - takes inputs and submits them to the AI engine, parses the output, and returns the results and token counts
//
//	Customer Messages: None
//	Errors: returned by GenerateContent, returned by loadSystemInstruction
//	Verifications: None
func (geminiServicePtr *GeminiService) GenerateContent(
	locationPtr *time.Location, prompt string, promptData map[string]string, systemInstructionTopic string,
	systemInstructionKey string,
) (geminiResponse GeminiResponse) {

	var (
		tGenerateContentResponsePtr *genai.GenerateContentResponse
		tInstruction                string
		tPromptData                 string
		tResponseParts              []genai.Part
	)

	for source, data := range promptData {
		tPromptData += fmt.Sprintf("%s %s ", source, data)
	}

	if tInstruction, geminiResponse.errorInfo = geminiServicePtr.loadSystemInstruction(locationPtr, systemInstructionTopic, systemInstructionKey); errorInfo.Error != nil {
		return
	}
	geminiServicePtr.modelPtr.SystemInstruction = &genai.Content{Parts: []genai.Part{genai.Text(tInstruction)}}

	if tGenerateContentResponsePtr, geminiResponse.errorInfo.Error = geminiServicePtr.modelPtr.GenerateContent(
		context.Background(), genai.Text(fmt.Sprintf("%s %s", prompt, tPromptData)),
	); geminiResponse.errorInfo.Error != nil {
		geminiResponse.errorInfo = errs.NewErrorInfo(geminiResponse.errorInfo.Error, ctv.VAL_EMPTY)
		return
	}

	tResponseParts = tGenerateContentResponsePtr.Candidates[0].Content.Parts
	for _, part := range tResponseParts {
		geminiResponse.response = strings.ReplaceAll(fmt.Sprintf("%s", part), "\n", "")
		geminiResponse.response = strings.ReplaceAll(geminiResponse.response, "json", "")
		geminiResponse.response = strings.ReplaceAll(geminiResponse.response, "\n", "")
		geminiResponse.response = strings.ReplaceAll(geminiResponse.response, "  ", " ")
		geminiResponse.response = strings.ReplaceAll(geminiResponse.response, "```", "")
	}

	geminiResponse.tokenCount = *tGenerateContentResponsePtr.UsageMetadata

	return
}

// loadSystemInstruction - using the system instruction topic and key, the instruction will be returned
//
//	Customer Messages: None
//	Errors: ErrSystemInstructionKeyInvalid, ErrSystemInstructionTopicInvalid
//	Verifications: None
func (geminiServicePtr *GeminiService) loadSystemInstruction(locationPtr *time.Location, topic string, key string) (systemInstruction string, errorInfo errs.ErrorInfo) {

	var (
		tSetDate      bool
		tOutputFormat string
	)

	switch topic {
	case SI_TOPIC_AI_QUESTION:
		switch key {
		case SI_KEY_SIMPLE_ANSWER:
			fallthrough
		case SI_KEY_COMPLEX_ANSWER:
			systemInstruction = geminiServicePtr.config.SystemInstructions.AIQuestion[key].Instruction
			tOutputFormat = geminiServicePtr.config.SystemInstructions.AIQuestion[key].OutputFormat
			tSetDate = geminiServicePtr.config.SystemInstructions.AIQuestion[key].SetDate
		default:
			errorInfo = errs.NewErrorInfo(errs.ErrSystemInstructionKeyInvalid, errs.BuildLabelValue(ctv.LBL_GEMINI_SYSTEM_INSTRUCTION_KEY, topic))
		}
	case SI_TOPIC_ANALYZE_QUESTION:
		switch key {
		case SI_KEY_CATEGORY_PROMPY_COMPARISON:
			fallthrough
		case SI_KEY_TIME_PERIOD_SPECIAL_WORDS_PRESENT:
			fallthrough
		case SI_KEY_TIME_PERIOD_WORDS_PRESENT:
			fallthrough
		case SI_KEY_TIME_PERIOD_VALUES:
			systemInstruction = geminiServicePtr.config.SystemInstructions.AnalyzeQuestion[key].Instruction
			tOutputFormat = geminiServicePtr.config.SystemInstructions.AnalyzeQuestion[key].OutputFormat
			tSetDate = geminiServicePtr.config.SystemInstructions.AnalyzeQuestion[key].SetDate
		default:
			errorInfo = errs.NewErrorInfo(errs.ErrSystemInstructionKeyInvalid, errs.BuildLabelValue(ctv.LBL_GEMINI_SYSTEM_INSTRUCTION_KEY, topic))
		}
	case SI_TOPIC_DETERMINE_API:
		systemInstruction = geminiServicePtr.config.SystemInstructions.DetermineAPI[SI_KEY_DETEMINE_API].Instruction
		tOutputFormat = geminiServicePtr.config.SystemInstructions.DetermineAPI[key].OutputFormat
		tSetDate = geminiServicePtr.config.SystemInstructions.DetermineAPI[key].SetDate
	case SI_TOPIC_GENERATE_ANSWER:
		systemInstruction = geminiServicePtr.config.SystemInstructions.GenerateAnswer[SI_KEY_BUSINESS_ANALYST].Instruction
		tOutputFormat = geminiServicePtr.config.SystemInstructions.GenerateAnswer[key].OutputFormat
		tSetDate = geminiServicePtr.config.SystemInstructions.GenerateAnswer[key].SetDate
	default:
		errorInfo = errs.NewErrorInfo(errs.ErrSystemInstructionTopicInvalid, errs.BuildLabelValue(ctv.LBL_GEMINI_SYSTEM_INSTRUCTION_TOPIC, topic))
	}

	if tOutputFormat != ctv.VAL_EMPTY {
		systemInstruction = fmt.Sprintf("%s %s", systemInstruction, tOutputFormat)
	}
	if tSetDate {
		systemInstruction = fmt.Sprintf("%s %v", systemInstruction, fmt.Sprintf("today %s timezone: %s", time.Now().In(locationPtr).Format("2006-01-02"), locationPtr.String()))
	}

	return
}

// Private methods below here

// loadGeminiConfig - reads, and returns a gemini service pointer
//
//	Customer Messages: None
//	Errors: error returned by ReadConfigFile or validateConfiguration
//	Verifications: validateConfiguration
func loadGeminiConfig(geminiConfigFilename string) (geminiConfig GeminiConfig, errorInfo errs.ErrorInfo) {

	var (
		tConfigData []byte
	)

	if errorInfo = hlps.CheckValueNotEmpty(geminiConfigFilename, errs.ErrRequiredParameterMissing, ctv.FN_CONFIG_FILENAME); errorInfo.Error != nil {
		return
	}

	if tConfigData, errorInfo.Error = os.ReadFile(hlps.PrependWorkingDirectory(geminiConfigFilename)); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_EXTENSION_CONFIG_FILENAME, geminiConfigFilename))
		return
	}

	if errorInfo.Error = json.Unmarshal(tConfigData, &geminiConfig); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_EXTENSION_CONFIG_FILENAME, geminiConfigFilename))
		return
	}

	return
}

// validateGeminiConfig - validates the gemini service configuration
//
//	Customer Messages: None
//	Errors: error returned by ReadConfigFile or validateConfiguration
//	Verifications: validateConfiguration
func validateGeminiConfig(geminiConfig GeminiConfig) (errorInfo errs.ErrorInfo) {

	if errorInfo = hlps.CheckValueNotEmpty(geminiConfig.MaxOutputTokens, errs.ErrRequiredParameterMissing, ctv.FN_GEMINI_MAX_OUTPUT_TOKENS); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckValueNotEmpty(geminiConfig.ModelName, errs.ErrRequiredParameterMissing, ctv.FN_GEMINI_MODEL_NAME); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckValueNotEmpty(geminiConfig.SetTopProbability, errs.ErrRequiredParameterMissing, ctv.FN_GEMINI_SET_TOP_K); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckMapLengthGTZero(geminiConfig.SystemInstructions.AIQuestion, errs.ErrRequiredParameterMissing, ctv.FN_SI_AI_QUESTION); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckMapLengthGTZero(geminiConfig.SystemInstructions.AnalyzeQuestion, errs.ErrRequiredParameterMissing, ctv.FN_SI_ANALYZE_QUESTION); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckMapLengthGTZero(geminiConfig.SystemInstructions.DetermineAPI, errs.ErrRequiredParameterMissing, ctv.FN_SI_DETERMINE_API); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckMapLengthGTZero(geminiConfig.SystemInstructions.GenerateAnswer, errs.ErrRequiredParameterMissing, ctv.FN_SI_GENERATE_ANSWER); errorInfo.Error != nil {
		return
	}
	errorInfo = hlps.CheckValueNotEmpty(geminiConfig.Temperature, errs.ErrRequiredParameterMissing, ctv.FN_GEMINI_TEMPERATURE)

	return
}
