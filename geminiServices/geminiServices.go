package sharedServices

import (
	"context"
	"encoding/json"
	"os"

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
func NewGeminiService(gcpCredentialsFilename string, geminiConfigFilename string) (geminiServicePtr *GeminiConfig, errorInfo errs.ErrorInfo) {

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

	geminiServicePtr = &tGeminiConfig

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
	if errorInfo = hlps.CheckValueNotEmpty(geminiConfig.GeminiTemperature, errs.ErrRequiredParameterMissing, ctv.FN_GEMINI_TEMPERATURE); errorInfo.Error != nil {
		return
	}

	return
}
