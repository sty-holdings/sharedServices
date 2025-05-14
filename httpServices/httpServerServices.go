package sharedServices

import (
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-yaml"

	ctv "github.com/sty-holdings/sharedServices/v2025/constantsTypesVars"
	errs "github.com/sty-holdings/sharedServices/v2025/errorServices"
	hlps "github.com/sty-holdings/sharedServices/v2025/helpers"
	vals "github.com/sty-holdings/sharedServices/v2025/validators"
)

// NewHTTPServer - creates a new httpServices service using the provided extension values.
//
//	Customer Messages: None
//	Errors: error returned by validateConfiguration
//	Verifications: validateConfiguration
func NewHTTPServer(configFilename string) (servicePtr *HTTPServerService, errorInfo errs.ErrorInfo) {

	var (
		tConfig HTTPConfiguration
	)

	if tConfig, errorInfo = loadHTTPServerConfig(configFilename); errorInfo.Error != nil {
		return
	}

	if errorInfo = validateConfiguration(tConfig); errorInfo.Error != nil {
		return
	}

	gin.SetMode(tConfig.GinMode)
	servicePtr = &HTTPServerService{
		Config: tConfig,
		Secure: false,
	}
	servicePtr.Config = tConfig
	if tConfig.TLSInfo.TLSCertFQN == ctv.VAL_EMPTY ||
		tConfig.TLSInfo.TLSPrivateKeyFQN == ctv.VAL_EMPTY {
		servicePtr.Secure = false
	} else {
		servicePtr.Secure = true
	}

	for _, port := range tConfig.Ports {
		if port < ctv.VAL_ZERO {
			port = int(math.Abs(float64(port)))
		}
		servicePtr.GinEnginePtr[uint(port)] = gin.Default()
		servicePtr.GinEnginePtr[uint(port)].LoadHTMLGlob(fmt.Sprintf("%s/*", tConfig.TemplateDirectory))
	}

	return
}

func (servicePtr *HTTPServerService) GenerateAddressWithPort(dnsOrIP string, port uint) string {

	return fmt.Sprintf("%s:%d", dnsOrIP, port)
}

// Private methods

//  Private Functions

// loadHTTPServerConfig - loads and parses an HTTP server configuration file.
//
//	Customer Messages: None
//	Errors: errs.Err if a config file is missing, empty, or contains invalid YAML.
//	Verifications: None
func loadHTTPServerConfig(configFilename string) (config HTTPConfiguration, errorInfo errs.ErrorInfo) {

	var (
		tConfigData []byte
	)

	if errorInfo = hlps.CheckValueNotEmpty(ctv.LBL_SERVICE_HTTP_SERVER, configFilename, errs.ErrEmptyRequiredParameter, ctv.FN_CONFIG_FILENAME); errorInfo.Error != nil {
		return
	}

	if tConfigData, errorInfo.Error = os.ReadFile(hlps.PrependWorkingDirectory(configFilename)); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_SERVICE_HTTP_SERVER, ctv.LBL_CONFIG_HTTP_SERVER_FILENAME, configFilename))
		return
	}

	if errorInfo.Error = yaml.Unmarshal(tConfigData, &config); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_SERVICE_HTTP_SERVER, ctv.LBL_CONFIG_HTTP_SERVER_FILENAME, configFilename))
		return
	}

	return
}

// validateConfiguration - validates the HTTP server configuration values.
//
//	Customer Messages: None
//	Errors: errs.ErrEmptyRequiredParameter, errs.ErrInvalidBase64, errs.ErrOSFileDoesntExist, errs.ErrOSFileUnreadable
//	Verifications: vals.IsGinModeValid, vals.DoesFileExistsAndReadable
func validateConfiguration(config HTTPConfiguration) (errorInfo errs.ErrorInfo) {

	if errorInfo = hlps.CheckArrayLengthGTZero(ctv.VAL_SERVICE_HTTP_SERVER, config.DeepLinks, errs.ErrEmptyRequiredParameter, ctv.LBL_HTTP_SERVER_REGISTRY); errorInfo.Error != nil {
		return
	}
	if vals.IsGinModeValid(config.GinMode) == false {
		errs.NewErrorInfo(errs.ErrInvalidBase64, fmt.Sprintf("%v%v", ctv.LBL_DIRECTORY, config.GinMode))
		return
	}
	if errorInfo = hlps.CheckValueNotEmpty(ctv.VAL_SERVICE_HTTP_SERVER, config.Host, errs.ErrEmptyRequiredParameter, ctv.LBL_HOST); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckValueNotEmpty(ctv.VAL_SERVICE_HTTP_SERVER, strconv.Itoa(config.Port), errs.ErrEmptyRequiredParameter, ctv.LBL_HTTP_PORT); errorInfo.Error != nil {
		return
	}
	if config.TLSInfo.TLSCertFQN != ctv.VAL_EMPTY && config.TLSInfo.TLSPrivateKeyFQN != ctv.VAL_EMPTY {
		if errorInfo = vals.DoesFileExistsAndReadable(config.TLSInfo.TLSCertFQN, ctv.LBL_TLS_CERTIFICATE_FILENAME); errorInfo.Error != nil {
			return
		}
		if errorInfo = vals.DoesFileExistsAndReadable(config.TLSInfo.TLSPrivateKeyFQN, ctv.LBL_TLS_PRIVATE_KEY_FILENAME); errorInfo.Error != nil {
			return
		}
	}

	return
}
