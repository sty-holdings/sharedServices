package sharedServices

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"os"
	"reflect"
	"regexp"
	"strings"
	"time"

	ctv "github.com/sty-holdings/sharedServices/v2025/constantsTypesVars"
	errs "github.com/sty-holdings/sharedServices/v2025/errorServices"
	pi "github.com/sty-holdings/sharedServices/v2025/programInfo"
)

// AreMapKeysPopulated - will test to make sure all map keys are set to anything other than nil or empty.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func AreMapKeysPopulated(myMap map[any]interface{}) bool {

	if IsMapPopulated(myMap) {
		for key, _ := range myMap {
			if key == nil || key == ctv.TXT_EMPTY {
				return false
			}
		}
	} else {
		return false
	}

	return true
}

// AreMapValuesPopulated - will test to make sure all map values are set to anything other than nil or empty.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func AreMapValuesPopulated(myMap map[any]interface{}) bool {

	if IsMapPopulated(myMap) {
		for _, value := range myMap {
			if value == nil || value == ctv.VAL_EMPTY {
				return false
			}
		}
	} else {
		return false
	}

	return true
}

// AreMapKeysValuesPopulated - check keys and value for missing values. Findings are ctv.GOOD, ctv.MISSING_VALUE,
// ctv.MISSING_KEY, or ctv.VAL_EMPTY_WORD.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: IsMapPopulated, AreMapKeysPopulated, AreMapValuesPopulated
func AreMapKeysValuesPopulated(myMap map[any]interface{}) (finding string) {

	if IsMapPopulated(myMap) {
		if AreMapKeysPopulated(myMap) {
			if AreMapValuesPopulated(myMap) {
				finding = ctv.TXT_GOOD
			} else {
				finding = ctv.TXT_MISSING_VALUE
			}
		} else {
			finding = ctv.TXT_MISSING_KEY
		}
	} else {
		finding = ctv.TXT_EMPTY
	}

	return
}

// DoesFileExistsAndReadable - works on any file. If the filename is not fully qualified
// the working directory will be prepended to the filename.
//
//	Customer Messages: None
//	Errors: ErrFileMissing, ErrFileUnreadable
//	Verifications: None
func DoesFileExistsAndReadable(filename, fileLabel string) (errorInfo errs.ErrorInfo) {

	var (
		fqn = PrependWorkingDirectory(filename)
	)

	if fileLabel == ctv.VAL_EMPTY {
		fileLabel = ctv.TXT_NO_LABEL_PROVIDED
	}

	if filename == ctv.VAL_EMPTY {
		errorInfo = errs.NewErrorInfo(errs.ErrRequiredParameterMissing, errs.BuildLabelValue(fileLabel, ctv.TXT_IS_EMPTY))
		return
	}
	if DoesFileExist(fqn) == false {
		errorInfo = errs.NewErrorInfo(errs.ErrFileDoesntExist, errs.BuildLabelValue(fileLabel, filename))
		return
	}
	if IsFileReadable(fqn) == false { // File is not readable
		errorInfo = errs.NewErrorInfo(errs.ErrFileUnreadable, errs.BuildLabelValue(fileLabel, filename))
	}

	return
}

// CheckFileValidJSON - reads the file and checks the contents
// func CheckFileValidJSON(FQN, fileLabel string) (errorInfo errs.ErrorInfo) {
//
// 	var (
// 		jsonData           []byte
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	errs.PrintDebugTrail(tFunctionName)
// 	errorInfo = errs.GetFunctionInfo()
// 	errorInfo.AdditionalInfo = fmt.Sprintf("File: %v  Config File Label: %v", FQN, fileLabel)
//
// 	if jsonData, errorInfo.Error = os.ReadFile(FQN); errorInfo.Error != nil {
// 		errorInfo.Error = errs.ErrFileUnreadable
// 		errorInfo.AdditionalInfo = fmt.Sprintf("FQN: %v File Label: %v", FQN, fileLabel)
// 		errs.PrintError(errorInfo)
// 	} else {
// 		if _isJSON := IsJSONValid(jsonData); _isJSON == false {
// 			errorInfo.Error = errs.ErrFileUnreadable
// 			errorInfo.AdditionalInfo = fmt.Sprintf("FQN: %v File Label: %v", FQN, fileLabel)
// 			errs.PrintError(errorInfo)
// 		}
// 	}
//
// 	return
// }

// DoesDirectoryExist - checks is the directory exists
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func DoesDirectoryExist(directoryName string) bool {

	return DoesFileExist(directoryName)
}

// DoesFileExist - does the value exist on the file system
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func DoesFileExist(fileName string) bool {

	if _, err := os.Stat(fileName); err == nil {
		return true
	}

	return false
}

// IsBase64Encode - will check if string is a valid base64 string.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func IsBase64Encode(base64Value string) bool {

	var (
		errorInfo errs.ErrorInfo
	)

	if _, errorInfo = Base64Decode(base64Value); errorInfo.Error == nil {
		return true
	}

	return false
}

// IsDirectoryFullyQualified - checks to see if the directory starts and ends with a slash.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func IsDirectoryFullyQualified(directory string) bool {

	if strings.HasPrefix(directory, ctv.FORWARD_SLASH) {
		if strings.HasSuffix(directory, ctv.FORWARD_SLASH) {
			return true
		}
	}

	return false

}

// IsDirectoryValid - checks to see if:
// - parameter is populated
// - directory exists and readable
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func IsDirectoryValid(directory string) bool {

	var (
		errorInfo errs.ErrorInfo
	)

	if errorInfo = ValidateDirectory(directory); errorInfo.Error == nil {
		return true
	}

	return false

}

// IsDomainValid - checks if domain naming is followed
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func IsDomainValid(domain string) bool {

	if strings.ToLower(domain) == ctv.LOCAL_HOST {
		return true
	} else {
		regex := regexp.MustCompile(`^(([a-zA-Z]{1})|([a-zA-Z]{1}[a-zA-Z]{1})|([a-zA-Z]{1}[0-9]{1})|([0-9]{1}[a-zA-Z]{1})|([a-zA-Z0-9][a-zA-Z0-9-_]{1,61}[a-zA-Z0-9]))\.([a-zA-Z]{2,6}|[a-zA-Z0-9-]{2,30}\.[a-zA-Z]{2,3})$`)
		if regex.MatchString(domain) {
			return true
		}
	}

	return false
}

// IsEmailAddressValid - checks the following:
// - length is > 2 and < 255
// - matches Regex validation
// - domain passes net.LookupMX call
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func IsEmailAddressValid(emailAddress string) (errorInfo errs.ErrorInfo) {

	var (
		mx []*net.MX
	)

	if len(emailAddress) < 3 || len(emailAddress) > 254 {
		errorInfo.Error = errors.New("The email address length must be greater than 2 and less than 255.")
	} else {
		if emailRegex.MatchString(emailAddress) {
			parts := strings.Split(emailAddress, "@")
			if mx, errorInfo.Error = net.LookupMX(parts[1]); errorInfo.Error != nil || len(mx) == 0 {
				errorInfo.Error = errors.New(fmt.Sprintf("The email address failed the Domain: '%v' lookup.", parts[1]))
			}
		} else {
			errorInfo.Error = errors.New(fmt.Sprintf("The email address '%v' is invalid.", emailAddress))
		}
	}

	return
}

// IsEnvironmentValid - checks that the value is valid. This function input is case-sensitive. Valid
// values are 'local', 'development', 'demo', and 'production'.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func IsEnvironmentValid(environment string) bool {

	switch environment {
	case ctv.ENVIRONMENT_LOCAL:
	case ctv.ENVIRONMENT_DEVELOPMENT:
	case ctv.ENVIRONMENT_DEMO:
	case ctv.ENVIRONMENT_PRODUCTION:
	default:
		return false
	}

	return true
}

// IsEmpty - checks that the value is empty.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func IsEmpty(value interface{}) bool {
	switch v := value.(type) {
	case string:
		return isEmptyString(v)
	case []interface{}: // Check for slices and arrays
		return isEmptyCollection(v)
	case map[interface{}]interface{}: // Check for maps
		return isEmptyCollection(v)
	case chan interface{}: // Check for channels
		return isEmptyCollection(v)
	case *interface{}: // Check for pointers
		return isEmptyPointer(v)
	default:
		// For other types, consider if they have an "empty" equivalent
		return false
	}
}

// IsExtensionValid - checks if the value is a valid extension
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func IsExtensionValid(extension string) bool {
	switch extension {
	case ctv.EXTENSION_ADMIN:
		fallthrough
	case ctv.EXTENSION_DK_CLIENT:
		fallthrough
	case ctv.EXTENSION_HAL:
		fallthrough
	case ctv.EXTENSION_SAAS_PROFILE:
		fallthrough
	default:
		return false
	}
}

// IsFileReadable - tries to open the file using 0644 permissions
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func IsFileReadable(fileName string) bool {

	if _, err := os.OpenFile(fileName, os.O_RDONLY, 0644); err == nil {
		return true
	}

	return false
}

// IsFutureYear - determines if the year is in the future
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func IsFutureMonth(year int, month int, location string) bool {

	var (
		tLocationPtr *time.Location
		tYear        int
		tMonth       time.Month
	)

	tLocationPtr, _ = time.LoadLocation(location)

	now := time.Now().In(tLocationPtr)
	tYear, tMonth, _ = now.Date()

	if year > tYear {
		return true
	}
	if year < tYear {
		return false
	}
	if month > int(tMonth) {
		return true
	}

	return false
}

// IsFutureYear - determines if the year is in the future
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func IsFutureYear(year int, location string) bool {

	var (
		tLocationPtr *time.Location
		tYear        int
	)

	tLocationPtr, _ = time.LoadLocation(location)

	now := time.Now().In(tLocationPtr)
	tYear, _, _ = now.Date()

	if year > tYear {
		return true
	}

	return false
}

// IsGinModeValid validates that the Gin httpServices framework mode is correctly set.
func IsGinModeValid(mode string) bool {

	switch strings.ToLower(mode) {
	case ctv.MODE_DEBUG:
	case ctv.MODE_RELEASE:
	default:
		return false
	}

	return true
}

// IsPopulated - checks that the value is populated.
func IsPopulated(value interface{}) bool {

	if IsEmpty(value) {
		return false
	}

	return true
}

// IsIPAddressValid - checks if the data provide is a valid IP address
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func IsIPAddressValid(ipAddress any) bool {

	// Checking if it is a valid IP addresses
	if IsIPv4Valid(ipAddress.(string)) || IsIPv6Valid(ipAddress.(string)) {
		return true
	}

	return false
}

// IsIPv4Valid - checks if the data provide is a valid IPv4 address
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func IsIPv4Valid(ipAddress any) bool {

	var (
		tIPv4Regex = regexp.MustCompile(`^(((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.|$)){4})`)
	)

	// Checking if it is a valid IPv4 addresses
	if tIPv4Regex.MatchString(ipAddress.(string)) {
		return true
	}

	return false
}

// IsIPv6Valid - checks if the data provide is a valid IPv6 address
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func IsIPv6Valid(ipAddress any) bool {

	var (
		tIPv6Regex = regexp.MustCompile(`^(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))$`)
	)

	// Checking if it is a valid IPv4 addresses
	if tIPv6Regex.MatchString(ipAddress.(string)) {
		return true
	}

	return false
}

// IsJSONValid - checks if the data provide is valid JSON.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func IsJSONValid(jsonIn []byte) bool {

	var (
		jsonString map[string]interface{}
	)

	return json.Unmarshal(jsonIn, &jsonString) == nil
}

// IsMapPopulated - will determine if the map is populated.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func IsMapPopulated(myMap map[any]interface{}) bool {

	return len(myMap) > 0
}

// IsStruct - will determine if the variable is a struct.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func IsStruct(v interface{}) bool {

	return reflect.TypeOf(v).Kind() == reflect.Struct
}

// IsMessagePrefixValid - is case-insensitive
// func IsMessagePrefixValid(messagePrefix string) bool {
//
// 	switch strings.ToUpper(messagePrefix) {
// 	case ctv.MESSAGE_PREFIX_SAVUPPROD:
// 	case ctv.MESSAGE_PREFIX_SAVUPDEV:
// 	case ctv.MESSAGE_PREFIX_SAVUPLOCAL:
// 	default:
// 		return false
// 	}
//
// 	return true
// }

// IsPeriodValid
// func IsPeriodValid(period string) bool {
//
// 	switch strings.ToUpper(period) {
// 	case ctv.YEAR:
// 	case ctv.MONTH:
// 	case ctv.DAY:
// 	default:
// 		return false
// 	}
//
// 	return true
// }

// This will set the connection values so GetConnection can be executed.
// func IsPostgresSSLModeValid(sslMode string) bool {
//
// 	switch sslMode {
// 	case ctv.POSTGRES_SSL_MODE_ALLOW:
// 	case ctv.POSTGRES_SSL_MODE_DISABLE:
// 	case ctv.POSTGRES_SSL_MODE_PREFER:
// 	case ctv.POSTGRES_SSL_MODE_REQUIRED:
// 	default:
// 		return false
// 	}
//
// 	return true
// }

// IsUserRegisterTypedValid
// func IsUserRegisterTypedValid(period string) bool {
//
// 	switch strings.ToUpper(period) {
// 	case ctv.COLLECTION_USER_TO_DO_LIST:
// 	case ctv.COLLECTION_USER_GOALS:
// 	default:
// 		return false
// 	}
//
// 	return true
// }

// IsURLValid
// func IsURLValid(URL string) bool {
//
// 	if _, err := url.ParseRequestURI(URL); err == nil {
// 		return true
// 	}
//
// 	return false
// }

// IsUUIDValid
// func IsUUIDValid(uuid string) bool {
//
// 	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9aAbB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
// 	return r.MatchString(uuid)
// }

// ValidateAuthenticatorService - Firebase is not support at this time
// func ValidateAuthenticatorService(authenticatorService string) (errorInfo errs.ErrorInfo) {
//
// 	var (
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	errs.PrintDebugTrail(tFunctionName)
//
// 	switch strings.ToUpper(authenticatorService) {
// 	case ctv.AUTH_COGNITO:
// 	case ctv.AUTH_FIREBASE:
// 		fallthrough // ToDo This is because AUTH_FIREBASE is not supported right now
// 	default:
// 		errorInfo.Error = errors.New(fmt.Sprintf("The supplied authenticator service is not supported! Authenticator Service: %v (Authenticator Service is case insensitive)", authenticatorService))
// 		if authenticatorService == ctv.VAL_EMPTY {
// 			errorInfo.AdditionalInfo = "Authenticator Service parameter is empty"
// 		} else {
// 			errorInfo.AdditionalInfo = "Authenticator Service: " + authenticatorService
// 		}
// 	}
//
// 	return
// }

// ValidateDirectory - validates that the directory value is not empty and the value exists on the file system
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func ValidateDirectory(directory string) (errorInfo errs.ErrorInfo) {

	if directory == ctv.VAL_EMPTY {
		errorInfo = errs.NewErrorInfo(errs.ErrRequiredArgumentMissing, ctv.TXT_DIRECTORY_PARAM_EMPTY)
		return
	}
	if DoesDirectoryExist(directory) == false {
		errorInfo = errs.NewErrorInfo(errs.ErrRequiredArgumentMissing, fmt.Sprintf("%v%v", ctv.LBL_DIRECTORY, directory))
	}

	return
}

// ValidateTransferMethod
// func ValidateTransferMethod(transferMethod string) (errorInfo errs.ErrorInfo) {
//
// 	var (
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	errs.PrintDebugTrail(tFunctionName)
//
// 	switch strings.ToUpper(transferMethod) {
// 	case ctv.TRANFER_STRIPE:
// 	case ctv.TRANFER_WIRE:
// 	case ctv.TRANFER_CHECK:
// 	case ctv.TRANFER_ZELLE:
// 	default:
// 		errorInfo.Error = errs.ErrTransferMethodInvalid
// 		if transferMethod == ctv.VAL_EMPTY {
// 			errorInfo.AdditionalInfo = "Transfer Method parameter is empty"
// 		} else {
// 			errorInfo.AdditionalInfo = "Transfer Method: " + transferMethod
// 		}
// 	}
//
// 	return
// }

// Private methods below here

func isEmptyString(value string) bool {
	return value == ""
}

func isEmptyCollection(value interface{}) bool {
	v := reflect.ValueOf(value)
	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array || v.Kind() == reflect.Map || v.Kind() == reflect.Chan {
		return v.IsNil() || v.Len() == 0
	}
	return false
}

func isEmptyPointer(value interface{}) bool {
	v := reflect.ValueOf(value)
	if v.Kind() == reflect.Ptr {
		return v.IsNil()
	}
	return false
}

// GetMapKeyPopulatedError - builds ...
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetMapKeyPopulatedError(finding string) (errorInfo errs.ErrorInfo) {

	pi.GetFunctionInfo(1)

	switch strings.ToLower(finding) {
	case ctv.TXT_EMPTY:
		errorInfo = errs.NewErrorInfo(errs.ErrMapIsEmpty, ctv.VAL_EMPTY)
	case ctv.TXT_MISSING_KEY:
		errorInfo = errs.NewErrorInfo(errs.ErrMapIsMissingKey, ctv.VAL_EMPTY)
	case ctv.TXT_MISSING_VALUE:
		errorInfo = errs.NewErrorInfo(errs.ErrMapIsMissingValue, ctv.VAL_EMPTY)
	case ctv.VAL_EMPTY:
		fallthrough
	default:
		errs.NewErrorInfo(errs.ErrMapIsMissingValue, "The 'finding' argument is empty.")
	}

	return
}
