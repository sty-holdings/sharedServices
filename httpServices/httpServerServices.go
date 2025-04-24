package sharedServices

import (
	"encoding/json"
	"fmt"
	"strings"

	ctv "github.com/sty-holdings/sharedServices/v2025/constantsTypesVars"
	errs "github.com/sty-holdings/sharedServices/v2025/errorServices"
	hlp "github.com/sty-holdings/sharedServices/v2025/helpers"
	vals "github.com/sty-holdings/sharedServices/v2025/validators"
)

// NewHTTPServer - creates a new httpServices service using the provided extension values.
//
//	Customer Messages: None
//	Errors: error returned by validateConfiguration
//	Verifications: validateConfiguration
func NewHTTPServer(configFilename string) (httpServicePtr *HTTPServerService, errorInfo errs.ErrorInfo) {

	var (
		tAdditionalInfo = fmt.Sprintf("%v%v", ctv.LBL_FILENAME, configFilename)
		tConfig         HTTPConfiguration
		tConfigData     []byte
	)

	// if tConfigData, errorInfo = config.ReadConfigFile(hlp.PrependWorkingDirectory(configFilename)); errorInfo.Error != nil {
	//	return
	// }

	if errorInfo.Error = json.Unmarshal(tConfigData, &tConfig); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, tAdditionalInfo)
		return
	}

	if errorInfo = validateConfiguration(tConfig); errorInfo.Error != nil {
		return
	}

	httpServicePtr.Config = tConfig
	httpServicePtr.CredentialsFQN = hlp.PrependWorkingDirectory(tConfig.CredentialsFilename)

	if tConfig.TLSInfo.TLSCert == ctv.VAL_EMPTY ||
		tConfig.TLSInfo.TLSPrivateKey == ctv.VAL_EMPTY ||
		tConfig.TLSInfo.TLSCABundle == ctv.VAL_EMPTY {
		httpServicePtr.Secure = false
	} else {
		httpServicePtr.Secure = true
	}

	return
}

// Private methods

//  Private Functions

// validateConfiguration - checks the NATS service configuration is valid.
//
//	Customer Messages: None
//	Errors: ErrEnvironmentInvalid, ErrMessageNamespaceInvalid, ErrDomainInvalid, error returned from DoesFileExistsAndReadable, ErrSubjectsMissing
//	Verifications: None
func validateConfiguration(config HTTPConfiguration) (errorInfo errs.ErrorInfo) {

	if errorInfo = vals.DoesFileExistsAndReadable(config.CredentialsFilename, ctv.LBL_FILENAME); errorInfo.Error != nil {
		errs.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", ctv.LBL_DIRECTORY, config.CredentialsFilename))
		return
	}
	if vals.IsBase64Encode(config.CredentialsFilename) == false {
		errs.NewErrorInfo(errs.ErrInvalidBase64, fmt.Sprintf("%v%v", ctv.LBL_DIRECTORY, config.CredentialsFilename))
		return
	}
	if vals.IsGinModeValid(config.GinMode) == false {
		errs.NewErrorInfo(errs.ErrInvalidBase64, fmt.Sprintf("%v%v", ctv.LBL_DIRECTORY, config.CredentialsFilename))
		return
	}
	if vals.IsEnvironmentValid(config.MessageEnvironment) == false {
		errorInfo = errs.NewErrorInfo(errs.ErrInvalidEnvirnoment, fmt.Sprintf("%v%v", ctv.LBL_ENVIRONMENT, config.MessageEnvironment))
		return
	}
	if vals.IsGinModeValid(config.GinMode) {
		config.GinMode = strings.ToLower(config.GinMode)
	} else {
		errorInfo = errs.NewErrorInfo(errs.ErrInvalidGinMode, fmt.Sprintf("%v%v", ctv.LBL_GIN_MODE, config.GinMode))
		return
	}
	if config.TLSInfo.TLSCert != ctv.VAL_EMPTY && config.TLSInfo.TLSPrivateKey != ctv.VAL_EMPTY && config.TLSInfo.TLSCABundle != ctv.VAL_EMPTY {
		if errorInfo = vals.DoesFileExistsAndReadable(config.TLSInfo.TLSCert, ctv.LBL_FILENAME); errorInfo.Error != nil {
			errs.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", ctv.LBL_DIRECTORY, config.TLSInfo.TLSCert))
			return
		}
		if errorInfo = vals.DoesFileExistsAndReadable(config.TLSInfo.TLSPrivateKey, ctv.LBL_FILENAME); errorInfo.Error != nil {
			errs.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", ctv.LBL_DIRECTORY, config.TLSInfo.TLSPrivateKey))
			return
		}
		if errorInfo = vals.DoesFileExistsAndReadable(config.TLSInfo.TLSCABundle, ctv.LBL_FILENAME); errorInfo.Error != nil {
			errs.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", ctv.LBL_DIRECTORY, config.TLSInfo.TLSCABundle))
			return
		}
	}
	if len(config.RouteRegistry) == ctv.VAL_ZERO {
		// errs.NewErrorInfo(errs.ErrSubjectsMissing, ctv.VAL_EMPTY)
	}

	return
}
