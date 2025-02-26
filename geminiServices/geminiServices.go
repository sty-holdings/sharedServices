package sharedServices

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"cloud.google.com/go/vertexai/genai"
	"google.golang.org/api/option"

	ctv "github.com/sty-holdings/sharedServices/v2025/constantsTypesVars"
	errs "github.com/sty-holdings/sharedServices/v2025/errorServices"
	hlps "github.com/sty-holdings/sharedServices/v2025/helpers"
	vals "github.com/sty-holdings/sharedServices/v2025/validators"
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
		geminiConfig: tGeminiConfig,
	}

	if geminiServicePtr.GeminiClientPtr, errorInfo.Error = genai.NewClient(
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

	geminiServicePtr.GeminiModelPtr = geminiServicePtr.GeminiClientPtr.GenerativeModel(geminiServicePtr.geminiConfig.GeminiModelName)

	if tInt64, errorInfo.Error = strconv.ParseInt(geminiServicePtr.geminiConfig.GeminiMaxOutputTokens, 10, 32); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errs.ErrIntegerInvalid, fmt.Sprintf("%s%s\n", ctv.LBL_GEMINI_MAX_OUTPUT_TOKENS, geminiServicePtr.geminiConfig.GeminiMaxOutputTokens))
		return
	}
	tInt32 = int32(tInt64)
	geminiServicePtr.GeminiModelPtr.MaxOutputTokens = &tInt32

	if tFloat64, errorInfo.Error = strconv.ParseFloat(geminiServicePtr.geminiConfig.GeminiSetTopProbability, 64); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errs.ErrFloatInvalid, fmt.Sprintf("%s%s\n", ctv.LBL_GEMINI_SET_TOP_PROBABILITY, geminiServicePtr.geminiConfig.GeminiSetTopProbability))
		return
	}
	tFloat32 = float32(tFloat64)
	geminiServicePtr.GeminiModelPtr.SetTopP(tFloat32)

	if tFloat64, errorInfo.Error = strconv.ParseFloat(geminiServicePtr.geminiConfig.GeminiTemperature, 64); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errs.ErrFloatInvalid, fmt.Sprintf("%s%s\n", ctv.LBL_GEMINI_TEMPERATURE, geminiServicePtr.geminiConfig.GeminiTemperature))
		return
	}
	tFloat32 = float32(tFloat64)
	geminiServicePtr.GeminiModelPtr.Temperature = &tFloat32

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

	if errorInfo = hlps.CheckValueNotEmpty(geminiConfig.GeminiMaxOutputTokens, errs.ErrRequiredParameterMissing, ctv.FN_GEMINI_MAX_OUTPUT_TOKENS); errorInfo.Error != nil {
		return
	}
	if errorInfo = vals.DoesFileExistsAndReadable(geminiConfig.GeminiModelName, ctv.FN_GEMINI_MODEL_NAME); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckValueNotEmpty(geminiConfig.GeminiSetTopProbability, errs.ErrRequiredParameterMissing, ctv.FN_GEMINI_SET_TOP_K); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckMapLengthGTZero(geminiConfig.GeminiSystemInstructions, errs.ErrRequiredParameterMissing, ctv.FN_GEMINI_SYSTEM_INSTRUCTION); errorInfo.Error != nil {
		return
	}
	errorInfo = hlps.CheckValueNotEmpty(geminiConfig.GeminiTemperature, errs.ErrRequiredParameterMissing, ctv.FN_GEMINI_TEMPERATURE)

	return
}
