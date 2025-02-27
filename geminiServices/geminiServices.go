package sharedServices

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

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

	return
}

// buildModel - will create the model instance and configure the settings
//
//	Customer Message: none
//	Errors: none
//	Verifications: none
func (geminiServicePtr *GeminiService) BuildModel() (errorInfo errs.ErrorInfo) {

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
//	Errors: None
//	Verifications: None
func (geminiServicePtr *GeminiService) GenerateContent(prompt string, promptData map[string]string, systemInstruction string) (
	response string, tokenCount genai.UsageMetadata, errorInfo errs.ErrorInfo,
) {

	var (
		tGenerateContentResponsePtr *genai.GenerateContentResponse
		tPromptData                 string
		tResponseParts              []genai.Part
	)

	for source, data := range promptData {
		tPromptData += fmt.Sprintf("%s %s ", source, data)
	}

	geminiServicePtr.modelPtr.SystemInstruction = &genai.Content{Parts: []genai.Part{genai.Text(systemInstruction)}}

	if tGenerateContentResponsePtr, errorInfo.Error = geminiServicePtr.modelPtr.GenerateContent(
		context.Background(), genai.Text(fmt.Sprintf("%s %s", prompt, tPromptData)),
	); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, ctv.VAL_EMPTY)
		return
	}

	tResponseParts = tGenerateContentResponsePtr.Candidates[0].Content.Parts
	for _, part := range tResponseParts {
		response = strings.ReplaceAll(fmt.Sprintf("%s", part), "\n", "")
		response = strings.ReplaceAll(response, "json", "")
		response = strings.ReplaceAll(response, "\n", "")
		response = strings.ReplaceAll(response, "  ", " ")
	}

	tokenCount = *tGenerateContentResponsePtr.UsageMetadata
	
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
	if errorInfo = hlps.CheckMapLengthGTZero(geminiConfig.SystemInstructions.AnalyzeQuestion, errs.ErrRequiredParameterMissing, ctv.FN_GEMINI_SYSTEM_INSTRUCTION); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckMapLengthGTZero(geminiConfig.SystemInstructions.Hal, errs.ErrRequiredParameterMissing, ctv.FN_GEMINI_SYSTEM_INSTRUCTION); errorInfo.Error != nil {
		return
	}
	errorInfo = hlps.CheckValueNotEmpty(geminiConfig.Temperature, errs.ErrRequiredParameterMissing, ctv.FN_GEMINI_TEMPERATURE)

	return
}
