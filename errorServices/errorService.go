package sharedServices

import (
	// Add imports here

	"fmt"
	"log"
	"runtime"

	ctv "github.com/sty-holdings/sharedServices/v2024/constantsTypesVars"
)

type ErrorInfo struct {
	AdditionalInfo string `json:"error_additional_info"`
	Error          error  `json:"-"`
	FileName       string `json:"error_filename"`
	FunctionName   string `json:"error_function_name"`
	LineNumber     int    `json:"error_line_number"`
	Message        string `json:"error_message"`
}

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

// BuildLabelValue - builds a string using a label and value. This can be used for ErrorInfo additional Info field. fmt.Sprintf("%s%s",...).
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func BuildLabelValue(label string, value string) (additionalInfo string) {

	return fmt.Sprintf("%s%s.", label, value)
}

// BuildUIdLabelValue - builds a string using a label and value. This can be used for ErrorInfo additional Info field. fmt.Sprintf("%s%s",...).
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func BuildUIdLabelValue(uId string, label string, value string) (additionalInfo string) {

	return fmt.Sprintf("UId: %s %s%s.", uId, label, value)
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
		"[ERROR] %v Additional Info: '%v' File: %v Near Line Number: %v\n",
		errorInfo.Error.Error(),
		errorInfo.AdditionalInfo,
		errorInfo.FileName,
		errorInfo.LineNumber,
	)
}

func newError(myError error) (errorInfo ErrorInfo) {

	errorInfo = getErrorFunctionFileNameLineNumber(3)
	errorInfo.Error = myError

	return
}

func getErrorFunctionFileNameLineNumber(level int) (errorInfo ErrorInfo) {

	var (
		tFunction, _, _, _ = runtime.Caller(level)
	)

	errorInfo.FunctionName = runtime.FuncForPC(tFunction).Name()
	_, errorInfo.FileName, errorInfo.LineNumber, _ = runtime.Caller(level)

	return
}

// DumpErrorInfos - outputs multiple error messages
//
//	func DumpErrorInfos(ErrorInfos []ErrorInfo) {
//		for _, info := range ErrorInfos {
//			PrintError(info)
//		}
//	}
