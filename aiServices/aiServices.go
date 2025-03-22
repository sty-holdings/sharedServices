package sharedServices

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"cloud.google.com/go/vertexai/genai"
	"google.golang.org/api/option"
	"gopkg.in/yaml.v3"

	ctv "github.com/sty-holdings/sharedServices/v2025/constantsTypesVars"
	errs "github.com/sty-holdings/sharedServices/v2025/errorServices"
	hlps "github.com/sty-holdings/sharedServices/v2025/helpers"
)

var (
	CTXBackground = context.Background()
)

// NewAIService - creates a new AI client
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func NewAIService(gcpCredentialsFilename string, gcpProjectId string, gcpLocation string, AIConfigFilename string, debugOn bool) (aiServicePtr *AIService, errorInfo errs.ErrorInfo) {

	var (
		tAIConfig AIConfig
	)

	if errorInfo = hlps.CheckValueNotEmpty(ctv.LBL_SERVICE_AI, gcpCredentialsFilename, errs.ErrEmptyRequiredParameter, ctv.FN_GCP_CREDENTIAL_FILENAME); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckValueNotEmpty(ctv.LBL_SERVICE_AI, AIConfigFilename, errs.ErrEmptyRequiredParameter, ctv.FN_SERVICE_CONFIG_FILENAME); errorInfo.Error != nil {
		return
	}

	if tAIConfig, errorInfo = loadAIConfig(AIConfigFilename); errorInfo.Error != nil {
		return
	}

	if errorInfo = validateAIConfig(tAIConfig); errorInfo.Error != nil {
		return
	}

	aiServicePtr = &AIService{
		config:  tAIConfig,
		debugOn: debugOn,
	}

	if aiServicePtr.clientPtr, errorInfo.Error = genai.NewClient(
		CTXBackground,
		gcpProjectId,
		gcpLocation,
		option.WithCredentialsFile(gcpCredentialsFilename),
	); errorInfo.Error != nil {
		return
	}

	aiServicePtr.modelPtrs = make(map[string]*genai.GenerativeModel, len(modelPoolNames))
	errorInfo = aiServicePtr.buildModelPool()

	return
}

// buildModelPool - will configure and create a model pool based on the number of entries in SITopicAnalyzeQuestionKeys.
//
//	Customer Message: none
//	Errors: none
//	Verifications: none
func (aiServicePtr *AIService) buildModelPool() (errorInfo errs.ErrorInfo) {

	var (
		tFloat32 float32
		tFloat64 float64
		tInt32   int32
		tInt64   int64
	)

	for _, worker := range modelPoolNames {
		aiServicePtr.modelPtrs[worker] = aiServicePtr.clientPtr.GenerativeModel(aiServicePtr.config.ModelName)
		if tInt64, errorInfo.Error = strconv.ParseInt(aiServicePtr.config.MaxOutputTokens, 10, 32); errorInfo.Error != nil {
			errorInfo = errs.NewErrorInfo(errs.ErrInvalidInteger, fmt.Sprintf("%s%s\n", ctv.LBL_AI_MAX_OUTPUT_TOKENS, aiServicePtr.config.MaxOutputTokens))
			return
		}
		tInt32 = int32(tInt64)
		aiServicePtr.modelPtrs[worker].MaxOutputTokens = &tInt32

		if tFloat64, errorInfo.Error = strconv.ParseFloat(aiServicePtr.config.SetTopProbability, 64); errorInfo.Error != nil {
			errorInfo = errs.NewErrorInfo(errs.ErrInvalidFloat, fmt.Sprintf("%s%s\n", ctv.LBL_AI_SET_TOP_PROBABILITY, aiServicePtr.config.SetTopProbability))
			return
		}
		tFloat32 = float32(tFloat64)
		aiServicePtr.modelPtrs[worker].SetTopP(tFloat32)

		if tFloat64, errorInfo.Error = strconv.ParseFloat(aiServicePtr.config.Temperature, 64); errorInfo.Error != nil {
			errorInfo = errs.NewErrorInfo(errs.ErrInvalidFloat, fmt.Sprintf("%s%s\n", ctv.LBL_AI_TEMPERATURE, aiServicePtr.config.Temperature))
			return
		}
		tFloat32 = float32(tFloat64)
		aiServicePtr.modelPtrs[worker].Temperature = &tFloat32
		tFloat32 = float32(tFloat64)
		aiServicePtr.modelPtrs[worker].SafetySettings = []*genai.SafetySetting{
			{
				Category:  genai.HarmCategoryHarassment,
				Threshold: genai.HarmBlockLowAndAbove,
			},
			{
				Category:  genai.HarmCategoryHateSpeech,
				Threshold: genai.HarmBlockLowAndAbove,
			},
			{
				Category:  genai.HarmCategoryDangerousContent,
				Threshold: genai.HarmBlockLowAndAbove,
			},
			{
				Category:  genai.HarmCategorySexuallyExplicit,
				Threshold: genai.HarmBlockLowAndAbove,
			},
		}
	}

	return
}

// GenerateContent - takes inputs and submits them to the AI engine, parses the output, and returns the results and token counts
//
//	Customer Messages: None
//	Errors: returned by GenerateContent, returned by loadSystemInstruction
//	Verifications: None
func (aiServicePtr *AIService) GenerateContent(
	extensionName string,
	locationPtr *time.Location,
	prompt string,
	promptData map[string]string,
	systemInstructionTopic string,
	systemInstructionKey string,
	additionalInstructions string,
) (aiResponse AIResponse) {

	var (
		tGenerateContentResponsePtr *genai.GenerateContentResponse
		tInstruction                string
		tPool                       = siTopicKeyPoolAssignment[systemInstructionKey]
		tPromptData                 string
		tResponseParts              []genai.Part
	)

	for source, data := range promptData {
		tPromptData += fmt.Sprintf("%s %s ", source, data)
	}

	if tInstruction, aiResponse.ErrorInfo = aiServicePtr.loadSystemInstruction(extensionName, locationPtr, systemInstructionTopic, systemInstructionKey); aiResponse.ErrorInfo.Error != nil {
		return
	}
	tInstruction = fmt.Sprintf("%s %s", tInstruction, additionalInstructions)
	aiServicePtr.modelPtrs[tPool].SystemInstruction = &genai.Content{Parts: []genai.Part{genai.Text(tInstruction)}}

	if tGenerateContentResponsePtr, aiResponse.ErrorInfo.Error = aiServicePtr.modelPtrs[tPool].GenerateContent(
		CTXBackground, genai.Text(fmt.Sprintf("%s %s", prompt, tPromptData)),
	); aiResponse.ErrorInfo.Error != nil {
		aiResponse.ErrorInfo = errs.NewErrorInfo(aiResponse.ErrorInfo.Error, ctv.VAL_EMPTY)
		return
	}

	tResponseParts = tGenerateContentResponsePtr.Candidates[0].Content.Parts
	for _, part := range tResponseParts {
		aiResponse.Response = strings.ReplaceAll(fmt.Sprintf("%s", part), "\n", "")
		aiResponse.Response = strings.ReplaceAll(aiResponse.Response, "json", "")
		aiResponse.Response = strings.ReplaceAll(aiResponse.Response, "\n", "")
		aiResponse.Response = strings.ReplaceAll(aiResponse.Response, "  ", " ")
		aiResponse.Response = strings.ReplaceAll(aiResponse.Response, "```", "")
	}

	aiResponse.SIKey = systemInstructionKey
	aiResponse.TokenCount = *tGenerateContentResponsePtr.UsageMetadata

	if aiServicePtr.debugOn {
		log.Printf("Pool: %s\n", tPool)
		log.Printf("SI Key: %s\n", aiResponse.SIKey)
		log.Printf("Response: %s\n", aiResponse.Response)
		log.Printf("token count: %d\n", aiResponse.TokenCount)
		log.Printf("Error: %d\n", aiResponse.ErrorInfo)
	}

	return
}

// loadSystemInstruction - using the system instruction topic and key, the instruction will be returned. If SetDate is true and locationPtr is nil,
// the default loc, err := time.LoadLocation("America/Los_Angeles") will be used.
//
//	Customer Messages: None
//	Errors: ErrSystemInstructionKeyInvalid, ErrSystemInstructionTopicInvalid
//	Verifications: None
func (aiServicePtr *AIService) loadSystemInstruction(extensionName string, locationPtr *time.Location, topic string, key string) (systemInstruction string, errorInfo errs.ErrorInfo) {

	var (
		tSetDate      bool
		tOutputFormat string
	)

	switch topic {
	case SI_TOPIC_ANALYZE_QUESTION:
		switch key {
		case SI_KEY_CATEGORY_SENTENCE:
			fallthrough
		case SI_KEY_SPECIAL_WORDS:
			fallthrough
		case SI_KEY_TIME_PERIOD_VALUES:
			systemInstruction = aiServicePtr.config.SystemInstructions[key].Instruction
			tOutputFormat = aiServicePtr.config.SystemInstructions[key].OutputFormat
			tSetDate = aiServicePtr.config.SystemInstructions[key].SetDate
		default:
			errorInfo = errs.NewErrorInfo(errs.ErrInvalidSystemInstructionKey, errs.BuildLabelValue(ctv.LBL_SERVICE_AI, ctv.LBL_AI_SYSTEM_INSTRUCTION_KEY, topic))
			return
		}
	case SI_TOPIC_GENERATE_ANSWER:
		switch key {
		case SI_KEY_BUSINESS_ANALYST:
			fallthrough
		case SI_KEY_MARKETING_ANALYST:
			fallthrough
		case SI_KEY_NOT_SUPPORTED:
			systemInstruction = aiServicePtr.config.SystemInstructions[key].Instruction
			tOutputFormat = aiServicePtr.config.SystemInstructions[key].OutputFormat
			tSetDate = aiServicePtr.config.SystemInstructions[key].SetDate
		default:
			errorInfo = errs.NewErrorInfo(errs.ErrInvalidSystemInstructionKey, errs.BuildLabelValue(ctv.LBL_SERVICE_AI, ctv.LBL_AI_SYSTEM_INSTRUCTION_KEY, topic))
			return
		}
	default:
		errorInfo = errs.NewErrorInfo(errs.ErrInvalidSystemInstructionTopic, errs.BuildLabelValue(ctv.LBL_SERVICE_AI, ctv.LBL_AI_SYSTEM_INSTRUCTION_TOPIC, topic))
		return
	}

	if tOutputFormat != ctv.VAL_EMPTY {
		systemInstruction = fmt.Sprintf("%s OUTPUT: %s", systemInstruction, tOutputFormat)
	}
	if tSetDate {
		if locationPtr == nil {
			if locationPtr, errorInfo.Error = time.LoadLocation("America/Los_Angeles"); errorInfo.Error != nil {
				errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(extensionName, ctv.LBL_TIMEZONE_LOCATION, ctv.TXT_FAILED))
			}
		}
		systemInstruction = fmt.Sprintf("%s %v", systemInstruction, fmt.Sprintf("TODAY: %s TIMEZONE: %s", time.Now().In(locationPtr).Format("2006-01-02"), locationPtr.String()))
	}

	return
}

// Private methods below here

// loadAIConfig - reads, and returns a ai service pointer
//
//	Customer Messages: None
//	Errors: error returned by ReadConfigFile or validateConfiguration
//	Verifications: validateConfiguration
func loadAIConfig(AIConfigFilename string) (AIConfig AIConfig, errorInfo errs.ErrorInfo) {

	var (
		tConfigData []byte
	)

	if errorInfo = hlps.CheckValueNotEmpty(ctv.LBL_SERVICE_AI, AIConfigFilename, errs.ErrEmptyRequiredParameter, ctv.LBL_CONFIG_AI_FILENAME); errorInfo.Error != nil {
		return
	}

	if tConfigData, errorInfo.Error = os.ReadFile(hlps.PrependWorkingDirectory(AIConfigFilename)); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(
			errorInfo.Error,
			errs.BuildLabelSubLabelValueMessage(ctv.LBL_SERVICE_AI, ctv.LBL_CONFIG_AI, ctv.LBL_CONFIG_EXTENSION_FILENAME, AIConfigFilename, ctv.TXT_READ_FAILED),
		)
		return
	}

	if errorInfo.Error = yaml.Unmarshal(tConfigData, &AIConfig); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(
			errorInfo.Error,
			errs.BuildLabelSubLabelValueMessage(ctv.LBL_SERVICE_AI, ctv.LBL_CONFIG_AI, ctv.LBL_CONFIG_EXTENSION_FILENAME, AIConfigFilename, ctv.TXT_UNMARSHAL_FAILED),
		)
		return
	}

	return
}

// validateAIConfig - validates the ai service configuration
//
//	Customer Messages: None
//	Errors: error returned by ReadConfigFile or validateConfiguration
//	Verifications: validateConfiguration
func validateAIConfig(AIConfig AIConfig) (errorInfo errs.ErrorInfo) {

	if errorInfo = hlps.CheckValueNotEmpty(ctv.LBL_SERVICE_AI, AIConfig.MaxOutputTokens, errs.ErrEmptyRequiredParameter, ctv.FN_AI_MAX_OUTPUT_TOKENS); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckValueNotEmpty(ctv.LBL_SERVICE_AI, AIConfig.ModelName, errs.ErrEmptyRequiredParameter, ctv.FN_AI_MODEL_NAME); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckValueNotEmpty(ctv.LBL_SERVICE_AI, AIConfig.SetTopProbability, errs.ErrEmptyRequiredParameter, ctv.FN_AI_SET_TOP_K); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckMapLengthGTZero(
		ctv.LBL_SERVICE_AI,
		AIConfig.SystemInstructions,
		errs.ErrEmptyRequiredParameter,
		ctv.FN_ANALYZE_QUESTION,
	); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckMapLengthGTZero(
		ctv.LBL_SERVICE_AI,
		AIConfig.SystemInstructions,
		errs.ErrEmptyRequiredParameter,
		ctv.FN_GENERATE_ANSWER,
	); errorInfo.Error != nil {
		return
	}
	errorInfo = hlps.CheckValueNotEmpty(ctv.LBL_SERVICE_AI, AIConfig.Temperature, errs.ErrEmptyRequiredParameter, ctv.FN_AI_TEMPERATURE)

	return
}
