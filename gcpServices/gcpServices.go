package sharedServices

import (
	"encoding/json"
	"os"

	ctv "github.com/sty-holdings/sharedServices/v2025/constantsTypesVars"
	errs "github.com/sty-holdings/sharedServices/v2025/errorServices"
	hlps "github.com/sty-holdings/sharedServices/v2025/helpers"
	vals "github.com/sty-holdings/sharedServices/v2025/validators"
)

// NewGCPService - creates a new GCP Service
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func NewGCPService(geminiConfigFilename string) (gcpServicePtr *GCPService, errorInfo errs.ErrorInfo) {

	var (
		tGCPConfig GCPConfig
	)

	if errorInfo = hlps.CheckValueNotEmpty(ctv.LBL_GCP_SERVIER, geminiConfigFilename, errs.ErrEmptyRequiredParameter, ctv.FN_SERVICE_CONFIG_FILENAME); errorInfo.Error != nil {
		return
	}

	if tGCPConfig, errorInfo = loadGCPConfig(geminiConfigFilename); errorInfo.Error != nil {
		return
	}

	if errorInfo = validateGCPConfig(tGCPConfig); errorInfo.Error != nil {
		return
	}

	gcpServicePtr = &GCPService{GCPConfig: tGCPConfig}

	return
}

// Private methods below here

// loadGCPConfig - reads, validates, and returns a gcp service configuration
//
//	Customer Messages: None
//	Errors: error returned by ReadConfigFile or validateConfiguration
//	Verifications: validateConfiguration
func loadGCPConfig(gcpConfigFilename string) (gcpConfig GCPConfig, errorInfo errs.ErrorInfo) {

	var (
		tConfigData []byte
	)

	if errorInfo = hlps.CheckValueNotEmpty(ctv.LBL_GCP_SERVIER, gcpConfigFilename, errs.ErrEmptyRequiredParameter, ctv.FN_CONFIG_FILENAME); errorInfo.Error != nil {
		return
	}

	if tConfigData, errorInfo.Error = os.ReadFile(hlps.PrependWorkingDirectory(gcpConfigFilename)); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_GCP_SERVIER, ctv.LBL_CONFIG_EXTENSION_FILENAME, gcpConfigFilename))
		return
	}

	if errorInfo.Error = json.Unmarshal(tConfigData, &gcpConfig); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_GCP_SERVIER, ctv.LBL_CONFIG_EXTENSION_FILENAME, gcpConfigFilename))
		return
	}

	return
}

// validateGCPConfig - reads, validates, and returns a gcp service configuration
//
//	Customer Messages: None
//	Errors: error returned by ReadConfigFile or validateConfiguration
//	Verifications: validateConfiguration
func validateGCPConfig(gcpConfig GCPConfig) (errorInfo errs.ErrorInfo) {

	if errorInfo = hlps.CheckValueNotEmpty(ctv.LBL_GCP_SERVIER, gcpConfig.GCPCredentialFilename, errs.ErrEmptyRequiredParameter, ctv.FN_GCP_CREDENTIAL_FILENAME); errorInfo.Error != nil {
		return
	}
	if errorInfo = vals.DoesFileExistsAndReadable(gcpConfig.GCPCredentialFilename, ctv.FN_GCP_CREDENTIAL_FILENAME); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckValueNotEmpty(ctv.LBL_GCP_SERVIER, gcpConfig.GCPLocation, errs.ErrEmptyRequiredParameter, ctv.FN_GCP_LOCATION); errorInfo.Error != nil {
		return
	}
	errorInfo = hlps.CheckValueNotEmpty(ctv.LBL_GCP_SERVIER, gcpConfig.GCPProjectID, errs.ErrEmptyRequiredParameter, ctv.FN_GCP_PROJECT_ID)

	return
}
