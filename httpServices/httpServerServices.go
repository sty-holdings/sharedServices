package sharedServices

import (
	"fmt"
	"math"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-yaml"

	ctv "github.com/sty-holdings/sharedServices/v2025/constantsTypesVars"
	errs "github.com/sty-holdings/sharedServices/v2025/errorServices"
	hlps "github.com/sty-holdings/sharedServices/v2025/helpers"
	vldts "github.com/sty-holdings/sharedServices/v2025/validators"
)

// NewHTTPServer - creates and initializes an HTTPServerService instance.
//
//	Customer Messages: None
//	Errors: errs.ErrEmptyRequiredParameter, errs.ErrInvalidBase64, errs.ErrOSFileDoesntExist, errs.ErrOSFileUnreadable
//	Verifications: vals.IsGinModeValid, vals.DoesFileExistsAndReadable
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
		Config:         tConfig,
		GinEnginePtrs:  make(map[uint]*gin.Engine),
		HTTPServerPtrs: make(map[uint]*http.Server),
		Secure:         false,
	}
	if tConfig.TLSInfo.TLSCertFQN != ctv.VAL_EMPTY ||
		tConfig.TLSInfo.TLSPrivateKeyFQN != ctv.VAL_EMPTY {
		servicePtr.Secure = true
	}

	for _, port := range tConfig.Ports {
		if port != ctv.VAL_ZERO {
			if port < ctv.VAL_ZERO {
				port = int(math.Abs(float64(port)))
			}
			servicePtr.Ports = append(servicePtr.Ports, uint(port))
			servicePtr.GinEnginePtrs[uint(port)] = gin.Default()
			if tConfig.TemplateDirectory != ctv.VAL_EMPTY {
				servicePtr.GinEnginePtrs[uint(port)].LoadHTMLGlob(fmt.Sprintf("%s/*", tConfig.TemplateDirectory))
			}
			servicePtr.HTTPServerPtrs[uint(port)] = &http.Server{
				Addr:    servicePtr.GenerateAddressWithPort(tConfig.Host, uint(port)),
				Handler: servicePtr.GinEnginePtrs[uint(port)],
			}
		}
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

	if errorInfo = vldts.CheckValueNotEmpty(ctv.LBL_SERVICE_HTTP_SERVER, configFilename, ctv.FN_CONFIG_FILENAME); errorInfo.Error != nil {
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

	if errorInfo = vldts.CheckArrayLengthGTZero(ctv.VAL_SERVICE_HTTP_SERVER, config.DeepLinks, ctv.LBL_HTTP_SERVER_REGISTRY); errorInfo.Error != nil {
		return
	}
	if vals.IsGinModeValid(config.GinMode) == false {
		errs.NewErrorInfo(errs.ErrInvalidBase64, fmt.Sprintf("%v%v", ctv.LBL_DIRECTORY, config.GinMode))
		return
	}
	if errorInfo = vldts.CheckValueNotEmpty(ctv.VAL_SERVICE_HTTP_SERVER, config.Host, ctv.LBL_HOST); errorInfo.Error != nil {
		return
	}
	if errorInfo = vldts.CheckArrayLengthGTZero(ctv.VAL_SERVICE_HTTP_SERVER, config.Ports, ctv.LBL_HTTP_PORT); errorInfo.Error != nil {
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
