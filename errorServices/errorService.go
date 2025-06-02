package sharedServices

import (
	// Add imports here

	"errors"
	"fmt"
	"log"
	"runtime"

	ctv "github.com/sty-holdings/sharedServices/v2025/constantsTypesVars"
)

// NewErrorInfo - will return an ErrorInfo object. It does not send anything to the log or the console.
//
//	Customer Messages: None
//	Errors: Missing values will be filled with 'MISSING'.
//	Verifications: None
func NewErrorInfo(
	myError error,
	additionalInfo string,
) (errorInfo ErrorInfo) {

	if myError == nil {
		return
	}

	errorInfo.AdditionalInfo = additionalInfo
	errorInfo.Error = myError
	errorInfo.Message = myError.Error()
	errorInfo.StackTrace = string(getStackTrace())

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

	return fmt.Sprintf("%s %s %s.", extensionName, label, value)
}

// BuildLabelValueMessage - builds a string using the label, value, and message. This can be used for ErrorInfo additional Info field. fmt.Sprintf("%s%s",...).
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func BuildLabelValueMessage(extensionName string, label string, value string, message string) (additionalInfo string) {

	return fmt.Sprintf("%s %s %s %s.", extensionName, label, value, message)
}

// BuildLabelSubLabelValue - builds a string using the label, sublabel, and value. This can be used for ErrorInfo additional Info field. fmt.Sprintf("%s%s",...).
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func BuildLabelSubLabelValue(extensionName string, label string, subLabel string, value string) (additionalInfo string) {

	return fmt.Sprintf("%s %s %s %s.", extensionName, label, subLabel, value)
}

// BuildLabelSubLabelValueMessage - builds a string using the label, sub-label, value, and message. This can be used for ErrorInfo additional Info field. fmt.Sprintf("%s%s",...).
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func BuildLabelSubLabelValueMessage(extensionName string, label string, subLabel string, value string, message string) (additionalInfo string) {

	return fmt.Sprintf("%s %s %s %s %s.", extensionName, label, subLabel, value, message)
}

// BuildInternalUserIDLabelValue - builds a string using the UID, label, and value. This can be used for ErrorInfo additional Info field. fmt.Sprintf("%s%s",...).
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func BuildInternalUserIDLabelValue(extensionName string, internalUserID string, label string, value string) (additionalInfo string) {

	return fmt.Sprintf("%s STYH Internal User Id: %s %s %s.", extensionName, internalUserID, label, value)
}

// BuildInternalUserIDLabelValueMessage - builds a string using the UID, label, value, and message. This can be used for ErrorInfo additional Info field. fmt.Sprintf("%s%s",...).
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func BuildInternalUserIDLabelValueMessage(extensionName string, internalUserID string, label string, value string, message string) (additionalInfo string) {

	return fmt.Sprintf("%s STYH Internal User Id: %s %s %s %s.", extensionName, internalUserID, label, value, message)
}

// BuildSystemActionLabelValue - builds a string using the UID, system action, label, and value. This can be used for ErrorInfo additional Info field. fmt.Sprintf("%s%s",...).
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func BuildSystemActionLabelValue(extensionName string, systemAction string, label string, value string) (additionalInfo string) {

	return fmt.Sprintf("%s System Action: %s %s %s.", extensionName, systemAction, label, value)
}

// BuildSystemActionLabelValueMessage - builds a string using the UID, system action, label, value, and message. This can be used for ErrorInfo additional Info field. fmt.Sprintf("%s%s",...).
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func BuildSystemActionLabelValueMessage(extensionName string, systemAction string, label string, value string, message string) (additionalInfo string) {

	return fmt.Sprintf("%s System Action: %s %s %s %s.", extensionName, systemAction, label, value, message)
}

// BuildSystemActionInternalUserIDLabelValue - builds a string using the UID, system action, label, and value. This can be used for ErrorInfo additional Info field. fmt.Sprintf("%s%s",...).
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func BuildSystemActionInternalUserIDLabelValue(extensionName string, internalUserID string, systemAction string, label string, value string) (additionalInfo string) {

	return fmt.Sprintf("%s System Action: %s STYH Internal User Id: %s %s %s.", extensionName, internalUserID, systemAction, label, value)
}

// BuildSystemActionInternalUserIDLabelValueMessage - builds a string using the UID, system action, label, value, and message.
// This can be used for ErrorInfo additional Info field. fmt.Sprintf("%s%s", ...).
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func BuildSystemActionInternalUserIDLabelValueMessage(extensionName string, internalUserID string, systemAction string, label string, value string, message string) (additionalInfo string) {

	return fmt.Sprintf("%s System Action: %s STYH Internal User Id: %s %s %s %s.", extensionName, systemAction, internalUserID, label, value, message)
}

func GetErrorInfoString(errorInfo ErrorInfo) string {

	return fmt.Sprintf("Error: %s Additional Info: '%s' \nStackTrace: %s\n", errorInfo.Error.Error(), errorInfo.AdditionalInfo, string(getStackTrace()))
}

// PrintError - processes and outputs error information with stack trace.
//
//	Customer Messages: None
//	Errors: errs.ErrEmptyError if myError is nil
//	Verifications: None
func PrintError(
	myError error,
	additionalInfo string,
) {

	var (
		errorInfo ErrorInfo
	)

	if myError == nil {
		errorInfo = newErrorInfoFromError(getStackTrace(), ErrEmptyError)
	} else {
		errorInfo = newErrorInfoFromError(getStackTrace(), myError)
	}

	errorInfo.AdditionalInfo = additionalInfo

	outputError(errorInfo)
}

// PrintErrorInfo - will output error information using this format:
// "[ERROR] {Error Message} Additional Info: '{Additional Info}' File: {Filename} Near Line Number: {Line Number}\n"
// If the outputMode is displayed, the color will be red. The default is to output to the log.
//
//	Customer Messages: None
//	Errors: ErrEmptyError
//	Verifications: None
func PrintErrorInfo(myErrorInfo ErrorInfo) {

	var (
		errorInfo ErrorInfo
	)

	errorInfo.AdditionalInfo = myErrorInfo.AdditionalInfo
	errorInfo.Error = myErrorInfo.Error
	errorInfo.Message = myErrorInfo.Message
	errorInfo.StackTrace = string(getStackTrace())

	outputError(errorInfo)
}

// Private Functions

func getStackTrace() (stackTrace []byte) {

	stackTrace = make([]byte, 2048)

	runtime.Stack(stackTrace, false)

	return
}

func outputError(errorInfo ErrorInfo) {

	log.Printf(
		"[ERROR] %s Additional Info: '%s' \nStackTrace: %s\n",
		errorInfo.Error.Error(),
		errorInfo.AdditionalInfo,
		errorInfo.StackTrace,
	)
}

func newErrorInfoFromError(stackTrace []byte, myError error) (errorInfo ErrorInfo) {

	errorInfo.StackTrace = string(stackTrace)
	errorInfo.Error = myError

	return
}
