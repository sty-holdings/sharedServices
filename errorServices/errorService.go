package sharedServices

import (
	// Add imports here

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

	if myError == nil {
		return
	} else {
		errorInfo = newError(myError)
	}

	if additionalInfo == ctv.VAL_EMPTY {
		errorInfo.AdditionalInfo = ctv.TXT_EMPTY
	} else {
		errorInfo.AdditionalInfo = additionalInfo
	}
	errorInfo.Message = myError.Error()

	return
}

// BuildLabelValue - builds a string using the label and value. This can be used for ErrorInfo additional Info field. fmt.Sprintf("%s%s",...).
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func BuildLabelValue(label string, value string) (additionalInfo string) {

	return fmt.Sprintf("%s%s.", label, value)
}

// BuildUIdLabelValue - builds a string using the uId, label, and value. This can be used for ErrorInfo additional Info field. fmt.Sprintf("%s%s",...).
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func BuildUIdLabelValue(uId string, label string, value string) (additionalInfo string) {

	return fmt.Sprintf("UId: %s %s%s.", uId, label, value)
}

// BuildUIdSubjectLabelValue - builds a string using the uId, subject, label, and value. This can be used for ErrorInfo additional Info field. fmt.Sprintf("%s%s",...).
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func BuildUIdSubjectLabelValue(uId string, subject string, label string, value string) (additionalInfo string) {

	return fmt.Sprintf("UId: %s Subject: %s %s%s.", uId, subject, label, value)
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
		errorInfo ErrorInfo
	)

	if myError == nil {
		errorInfo = newError(ErrErrorMissing)
	} else {
		errorInfo = newError(myError)
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
//	Errors: ErrErrorMissing
//	Verifications: None
func PrintErrorInfo(errorInfo ErrorInfo) {

	if errorInfo.Error == nil {
		errorInfo = newError(ErrErrorMissing)
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

func newError(myError error) (errorInfo ErrorInfo) {

	errorInfo.StackTrace = getErrorFunctionFileNameLineNumber()
	errorInfo.Error = myError

	return
}

func getErrorFunctionFileNameLineNumber() string {

	var (
		buf = make([]byte, 1024)
	)

	runtime.Stack(buf, false)

	return string(buf)
}

// DumpErrorInfos - outputs multiple error messages
//
//	func DumpErrorInfos(ErrorInfos []ErrorInfo) {
//		for _, info := range ErrorInfos {
//			PrintError(info)
//		}
//	}
