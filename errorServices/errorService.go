package sharedServices

import (
	// Add imports here

	"errors"
	"fmt"
	"log"
	"runtime"

	ctv "github.com/sty-holdings/sharedServices/v2025/constantsTypesVars"
)

// NewErrorInfo - will return an ErrorInfo object.
//
//	Customer Messages: None
//	Errors: Missing values will be filled with 'MISSING'.
//	Verifications: None
func NewErrorInfo(
	myError error,
	additionalInfo string,
) (errorInfo ErrorInfo) {

	var (
		buf = make([]byte, 2048)
	)

	if myError == nil {
		return
	} else {
		runtime.Stack(buf, false)
		errorInfo = newError(buf, myError)
	}

	if additionalInfo == ctv.VAL_EMPTY {
		errorInfo.AdditionalInfo = ctv.TXT_EMPTY
	} else {
		errorInfo.AdditionalInfo = additionalInfo
	}
	errorInfo.Message = myError.Error()

	outputError(errorInfo)

	return
}

// NewGRPCErrorInfo - will return an ErrorInfo object with the Error containing both the error and additional info
// combined. Only the errorInfo.Error property will be returned. All other properties are empty.
//
//	Customer Messages: None
//	Errors: Missing values will be filled with 'MISSING'.
//	Verifications: None
func NewGRPCErrorInfo(
	myError error,
	additionalInfo string,
) (errorInfo ErrorInfo) {

	if myError == nil {
		return
	}

	if additionalInfo == ctv.VAL_EMPTY {
		errorInfo.AdditionalInfo = ctv.TXT_EMPTY
	} else {
		errorInfo.AdditionalInfo = additionalInfo
	}
	errorInfo.Error = errors.New(fmt.Sprintf("%s - Additional Info: %s", myError.Error(), additionalInfo))

	outputError(errorInfo)

	return
}

// BuildLabelValue - builds a string using the label and value. This can be used for ErrorInfo additional Info field. fmt.Sprintf("%s%s",...).
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func BuildLabelValue(extensionName string, label string, value string) (additionalInfo string) {

	return fmt.Sprintf("%s %s%s.", extensionName, label, value)
}

// BuildLabelSubLabelValueMessage - builds a string using the label, sub-label, value and message. This can be used for ErrorInfo additional Info field. fmt.Sprintf("%s%s",...).
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func BuildLabelSubLabelValueMessage(extensionName string, label string, subLabel string, value string, message string) (additionalInfo string) {

	return fmt.Sprintf("%s %s%s%s %s.", extensionName, label, subLabel, value, message)
}

// BuildSTYHUserIdLabelValue - builds a string using the UID, label, and value. This can be used for ErrorInfo additional Info field. fmt.Sprintf("%s%s",...).
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func BuildSTYHUserIdLabelValue(extensionName string, styhUserId string, label string, value string) (additionalInfo string) {

	return fmt.Sprintf("%s STYH User Id: %s %s %s.", extensionName, styhUserId, label, value)
}

// BuildSystemActionLabelValue - builds a string using the UID, system action, label, and value. This can be used for ErrorInfo additional Info field. fmt.Sprintf("%s%s",...).
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func BuildSystemActionLabelValue(extensionName string, systemAction string, label string, value string) (additionalInfo string) {

	return fmt.Sprintf("%s System Action: %s %s%s.", extensionName, systemAction, label, value)
}

// BuildSystemActionSTYHUserIdLabelValue - builds a string using the UID, system action, label, and value. This can be used for ErrorInfo additional Info field. fmt.Sprintf("%s%s",...).
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func BuildSystemActionSTYHUserIdLabelValue(extensionName string, styhUserId string, systemAction string, label string, value string) (additionalInfo string) {

	return fmt.Sprintf("%s System Action: %s STYH User Id: %s %s%s.", extensionName, styhUserId, systemAction, label, value)
}

// PrintError - will output error information using this format:
// "[ERROR] {Error Message} Additional Info: '{Additional Info}' File: {Filename} Near Line Number: {Line Number}\n"
// If the outputMode is display, the color will be red. The default is to output to the log.
//
//	Customer Messages: None
//	Errors: Missing values will be filled with 'MISSING'.
//	Verifications: None
func PrintError(
	myError error,
	additionalInfo string,
) {

	var (
		buf       = make([]byte, 2048)
		errorInfo ErrorInfo
	)

	runtime.Stack(buf, false)
	if myError == nil {
		errorInfo = newError(buf, ErrEmptyError)
	} else {
		errorInfo = newError(buf, myError)
	}
	if additionalInfo == ctv.VAL_EMPTY {
		errorInfo.AdditionalInfo = ctv.TXT_EMPTY
	} else {
		errorInfo.AdditionalInfo = additionalInfo
	}

	outputError(errorInfo)
}

// PrintErrorInfo - will output error information using this format:
// "[ERROR] {Error Message} Additional Info: '{Additional Info}' File: {Filename} Near Line Number: {Line Number}\n"
// If the outputMode is display, the color will be red. The default is to output to the log.
//
//	Customer Messages: None
//	Errors: ErrEmptyError
//	Verifications: None
func PrintErrorInfo(errorInfo ErrorInfo) {

	var (
		buf = make([]byte, 2048)
	)

	runtime.Stack(buf, false)
	if errorInfo.Error == nil {
		errorInfo = newError(buf, ErrEmptyError)
	}

	outputError(errorInfo)
}

// Private Functions
func outputError(errorInfo ErrorInfo) {

	log.Printf(
		"[ERROR] %s Additional Info: '%s' \nStackTrace: %s\n",
		errorInfo.Error.Error(),
		errorInfo.AdditionalInfo,
		errorInfo.StackTrace,
	)
}

func newError(stackTrace []byte, myError error) (errorInfo ErrorInfo) {

	errorInfo.StackTrace = string(stackTrace)
	errorInfo.Error = myError

	return
}
