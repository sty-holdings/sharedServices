package sharedServices

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	// "net/httpServices"
	// "os"
	// "time"

	ctv "github.com/sty-holdings/sharedServices/v2024/constantsTypesVars"
	errs "github.com/sty-holdings/sharedServices/v2024/errorServices"
	hlp "github.com/sty-holdings/sharedServices/v2024/helpers"
	jwts "github.com/sty-holdings/sharedServices/v2024/jwtServices"
	pi "github.com/sty-holdings/sharedServices/v2024/programInfo"
	vals "github.com/sty-holdings/sharedServices/v2024/validators"
)

type HTTPConfiguration struct {
	CredentialsFilename string       `json:"credentials_filename"`
	GinMode             string       `json:"gin_mode"`
	HTTPDomain          string       `json:"http_domain"`
	MessageEnvironment  string       `json:"message_environment"`
	Port                int          `json:"port"`
	RequestedThreads    uint         `json:"requested_threads"`
	RouteRegistry       []RouteInfo  `json:"route_registry"`
	TLSInfo             jwts.TLSInfo `json:"tls_info"`
}

type RouteInfo struct {
	Namespace   string `json:"namespace"`
	Subject     string `json:"subject"`
	Description string `json:"description"`
}

type HTTPService struct {
	Config         HTTPConfiguration
	CredentialsFQN string
	HTTPServerPtr  *http.Server
	Secure         bool
}

// NewHTTP - creates a new httpServices service using the provided extension values.
//
//	Customer Messages: None
//	Errors: error returned by validateConfiguration
//	Verifications: validateConfiguration
func NewHTTP(configFilename string) (
	service HTTPService,
	errorInfo errs.ErrorInfo,
) {

	var (
		tAdditionalInfo = fmt.Sprintf("%v%v", ctv.LBL_FILENAME, configFilename)
		tConfig         HTTPConfiguration
		tConfigData     []byte
	)

	//if tConfigData, errorInfo = config.ReadConfigFile(hlp.PrependWorkingDirectory(configFilename)); errorInfo.Error != nil {
	//	return
	//}

	if errorInfo.Error = json.Unmarshal(tConfigData, &tConfig); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, tAdditionalInfo)
		return
	}

	if errorInfo = validateConfiguration(tConfig); errorInfo.Error != nil {
		return
	}

	service.Config = tConfig
	service.CredentialsFQN = hlp.PrependWorkingDirectory(tConfig.CredentialsFilename)

	if tConfig.TLSInfo.TLSCert == ctv.VAL_EMPTY ||
		tConfig.TLSInfo.TLSPrivateKey == ctv.VAL_EMPTY ||
		tConfig.TLSInfo.TLSCABundle == ctv.VAL_EMPTY {
		service.Secure = false
	} else {
		service.Secure = true
	}

	return
}

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
		errs.NewErrorInfo(pi.ErrBase64Invalid, fmt.Sprintf("%v%v", ctv.LBL_DIRECTORY, config.CredentialsFilename))
		return
	}
	if vals.IsGinModeValid(config.GinMode) == false {
		errs.NewErrorInfo(pi.ErrBase64Invalid, fmt.Sprintf("%v%v", ctv.LBL_DIRECTORY, config.CredentialsFilename))
		return
	}
	if vals.IsEnvironmentValid(config.MessageEnvironment) == false {
		errorInfo = errs.NewErrorInfo(pi.ErrEnvironmentInvalid, fmt.Sprintf("%v%v", ctv.LBL_ENVIRONMENT, config.MessageEnvironment))
		return
	}
	if vals.IsGinModeValid(config.GinMode) {
		config.GinMode = strings.ToLower(config.GinMode)
	} else {
		errorInfo = errs.NewErrorInfo(pi.ErrGinModeInvalid, fmt.Sprintf("%v%v", ctv.LBL_GIN_MODE, config.GinMode))
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
		errs.NewErrorInfo(pi.ErrSubjectsMissing, ctv.VAL_EMPTY)
	}

	return
}
